package server

import (
	"forBlossem/adapter/route"
	"forBlossem/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func routes(engine *gin.Engine) {
	route.Route(engine, http.MethodGet, "/ping", handlers.PingHandler)
	route.Route(engine, http.MethodGet, "/get_access_token", handlers.AccessTokenGetHandler)
}
