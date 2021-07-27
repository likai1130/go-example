package middlewares

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/paradeum-team/gin-prometheus-ext"
	"go-example/config"
	"net/http"
	"strings"
)

var UriParamsMap = make(map[string]string)

/**
func init() {
	UriParamsMap["agfid"] = ":agfid"
	UriParamsMap["queryType"] = ":queryType"
	UriParamsMap["paramValue"] = ":paramValue"
	UriParamsMap["dgst"] = ":dgst"

}
*/

/**
Prometheus设置
*/
func PrometheusSetUp() *ginprometheusext.Prometheus {
	if config.AppConfig.Server.PProf {
		p := ginprometheusext.NewPrometheus("gin")
		p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
			url := c.Request.URL.Path
			for _, p := range c.Params {
				if v, ok := UriParamsMap[p.Key]; ok {
					url = strings.Replace(url, p.Value, v, 1)
				}
			}
			return url
		}

		monitorRouter := gin.New()
		p.SetListenAddressWithRouter(fmt.Sprintf(":%d", config.AppConfig.Server.PProfPort), monitorRouter)
		pprof.Register(monitorRouter)
		monitorRouter.GET("/env", func(c *gin.Context) {
			c.JSON(http.StatusOK, config.AppConfig)
		})
		return p
	}
	return nil
}
