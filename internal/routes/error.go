package routes

import errordelivery "skeleton/internal/api/error/delivery"

func (r *Route) mountErrorRoutes(handler *errordelivery.ErrorHandler) {
	r.srv.Router.NotFound(handler.NotFound)
	r.srv.Router.MethodNotAllowed(handler.MethodNotAllowed)
}
