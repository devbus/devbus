package main

import (
	"strings"

	"flag"
	"fmt"
	"os"

	"net/http"

	"github.com/devbus/devbus/common"
	_ "github.com/devbus/devbus/controllers"
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

func main() {
	flag.Parse()
	if *flag.Bool("h", false, "show usage") {
		flag.Usage()
		return
	}
	confPath := flag.String("conf", "", "devbus config file")
	if *confPath == "" {
		flag.Usage()
		return
	}
	if err := common.Parse(*confPath); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse devbus config, error: %+v", err)
		os.Exit(1)
	}

	router := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("devbus", store))

	router.Static("/app", "webapp/dist")
	router.Static("/static", "webapp/dist/static")

	router.Use(authFilter)

	router.Run(":8080")
}
