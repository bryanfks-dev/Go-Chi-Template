@echo off
setlocal enabledelayedexpansion

REM Ask for module name
set /p MODULE_NAME=What's your module name?

echo Setting module name to: %MODULE_NAME%...

REM Update go.mod module name
go mod edit -module %MODULE_NAME%

REM Replace "skeleton" with module name in all .go files
for /r %%f in (*.go) do (
    powershell -Command ^
        "(Get-Content '%%f') -replace 'skeleton', '%MODULE_NAME%' | Set-Content '%%f'"
)

REM Ask to install dependencies
set /p INSTALL_DEPS=Do you want to install dependencies now? (y/n)

if /i "%INSTALL_DEPS%"=="y" (
    echo Installing dependencies...
    go mod tidy
)

REM Remove this init script
del scripts\init.sh 2>nul
del scripts\init.bat 2>nul

echo Initialization complete, you're all set!
pause
