package middleware

import (
	"net/http"
	"skeleton/pkg/config"

	"github.com/go-chi/cors"
)

func CORSMiddleware(
	env config.Environment,
) func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			if env == config.EnvironmentDevelopment {
				return true
			}

			return origin == ProductionOrigin
		},
	})
}
