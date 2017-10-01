package router

import (
	"net/http"
	"strings"

	"github.com/devbus/devbus/apis"
	"github.com/devbus/devbus/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func authFilter(context *gin.Context) {
	uri := context.Request.URL.Path
	session := sessions.Default(context)
	if session.Get("uid") == nil {
		if !strings.HasPrefix(uri, "/api/session") &&
			!strings.HasPrefix(uri, "/app") &&
			!strings.HasPrefix(uri, "/static") {
			context.Redirect(http.StatusFound, "/app/session.html")
		}
	}
}

var router = gin.Default()

func routes() {
	group := router.Group("/api")

	// session
	group.POST("/session/login", apis.Login)
	group.POST("/session/register", apis.Register)
	group.GET("/session//activate", apis.Activate)

	// user
	group.PUT("/user/password", apis.ModifyPassword)

	// project
	group.GET("/project", apis.GetProject)
	group.POST("/project", apis.CreateProject)
	group.PUT("/project", apis.ModifyProject)

}

func Run() {
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("devbus", store))

	router.Static("/app", "webapp/dist")
	router.Static("/static", "webapp/dist/static")

	router.Use(authFilter)

	routes()

	router.Run(common.Config.Addr)
}
