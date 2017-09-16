package main

import (
	_ "github.com/devbus/devbus/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"encoding/json"
	"github.com/astaxie/beego/context"
	"strings"
)

var globalSessions *session.Manager

func init() {
	conf := `
{
  "cookieName": "gosessionid",
  "enableSetCookie,omitempty": true,
  "gclifetime": 3600,
  "maxLifetime": 3600,
  "secure": false,
  "sessionIDHashFunc": "sha1",
  "sessionIDHashKey": "",
  "cookieLifeTime": 3600,
  "providerConfig": ""
}
	`
	mc := &session.ManagerConfig{}
	json.Unmarshal([]byte(conf), mc)
	globalSessions, _ = session.NewManager("memory", mc)
	go globalSessions.GC()
}

// handle
func init() {
	beego.InsertFilter("*", beego.BeforeExec, func(ctx *context.Context) {
		sess := ctx.Input.CruSession
		if sess == nil {
			uri := ctx.Request.RequestURI
			if !strings.HasPrefix(uri, "/api/session") &&
				!strings.HasSuffix(uri, ".html") &&
				!strings.HasSuffix(uri, ".js") &&
				!strings.HasSuffix(uri, ".css") &&
				!strings.HasSuffix(uri, "*font") {
				ctx.Redirect(302, "/app/session.html")
			}
		}
	})
}

func main() {
	beego.Run()
}
