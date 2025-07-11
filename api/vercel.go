// Package vercel starts a web server to generate an SVG for the user's GitHub language statistics.

package api

import (
	"net/http"

	"go-readme-stats/app/routes"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.New()
	routes.Register(app)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
