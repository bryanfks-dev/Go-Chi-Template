package routes

import (
	"fmt"

	httpswagger "github.com/swaggo/http-swagger"
)

func (r *Route) mountDocsRoutes() {
	fmt := fmt.Sprintf(
		"API Docs will be available at %s%s",
		r.srv.Address(),
		APIDocsRoute,
	)
	r.logger.Info(fmt)

	r.srv.Router.Mount(APIDocsRoute, httpswagger.WrapHandler)
}
