package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-example/middlewares"
	"time"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*gin.Engine
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

// New returns a new Bootstrapper.
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Engine:       gin.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}
func (b *Bootstrapper) Bootstrap() *Bootstrapper {

	//设置业务日志级别或者中间件
	b.Use(middlewares.LoggerToFile(b.AppName))

	//设置监控
	if p := middlewares.PrometheusSetUp(); p != nil {
		p.Use(b.Engine)
	}
	b.Use(gin.Recovery())
	b.Use(middlewares.Cors())
	//b.Use(middlewares.Authentication())

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...Configurator) {
	b.Run(addr)
}
