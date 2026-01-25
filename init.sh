#!/bin/sh

OS=$(uname)

# Replace module name
read -p "What's your module name? " -r MODULE_NAME

echo "Setting module name to: $MODULE_NAME..."

# Update go.mod and source files
go mod edit -module $MODULE_NAME
if [ "$OS" == "Darwin" ]; then
    find . -type f -name "*.go" -print0 | xargs -0 sed -i '' "s/skeleton/$MODULE_NAME/g"
elif [ "$OS" == "Linux" ]; then
    find . -type f -name "*.go" -print0 | xargs -0 sed -i "s/skeleton/$MODULE_NAME/g"
fi

# Installing dependencies
read -p "Do you want to install dependencies now? (y/n) " -r INSTALL_DEPS
if [ "$INSTALL_DEPS" = "y" ] || [ "$INSTALL_DEPS" = "Y" ]; then
    echo "Installing dependencies..."
    go mod tidy
fi

# At the end, remove this init script
rm -f init.sh

echo "Initialization complete, you're all set!"
