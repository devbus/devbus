package main

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/devbus/devbus/apis"
	"github.com/devbus/devbus/common"
	"github.com/devbus/devbus/router"
	"github.com/gin-gonic/gin"
)

var help = flag.Bool("h", false, "show usage")
var verbose = flag.Bool("verbose", false, "show verbose log")
var confPath = flag.String("conf", "", "devbus config file")

func main() {
	flag.Parse()
	if *help {
		flag.Usage()
		return
	}
	if *confPath == "" {
		flag.Usage()
		return
	}
	if err := common.Parse(*confPath); err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse devbus config, error: %+v", err)
		os.Exit(1)
	}

	if *verbose {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Run()
}
