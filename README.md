# go-example

##  概述
 
    这是一个web服务,使用的是gin框架。避免重复造轮子，把日常用到的一些组件或者中间件集成起来，方便下次直接使用。


### 一、引入go-gin框架，抽象出来启动器

- bootstrap/bootstrapper.go

	```
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
	
	```

- main.go

	```
	func init() {
		gredis.SetUp()
	}
	
	func newApp() *bootstrap.Bootstrapper {
		middlewares.SystemLogsSetUp()
		app := bootstrap.New("gin-web", "gin-web-example")
		app.Bootstrap()
		app.Configure(router.Configure)
		return app
	}
	
	//增加项目地址
	// @termsOfService https://github.com/likai1130/go-example
	func main() {
		app := newApp()
		conf := config.AppConfig
		port := conf.Server.Port
		listenPort := fmt.Sprintf(":%v", port) //格式化监听端口
		app.Listen(listenPort)
	
	}

	```


- 支持swagger. web/router/router.go

	```
	/**
	 * 路由配置，并根据配置文件设置根路径
	 * 参考url：https://github.com/gin-gonic/gin
	 */
	func Configure(r *bootstrap.Bootstrapper) {
		prefix := "/"
		//此处可以增加系统应用目录根路径
		pldConf := config.AppConfig
		contextPath := pldConf.Server.ContextPath
		if "" != contextPath && strings.HasPrefix(contextPath, "/") {
			//给拼接好的api ，增加前缀
			prefix = contextPath
		}
		rootRouter := r.Group(prefix) //设置访问的根目录
		concreteRouter(rootRouter)
		docs.SwaggerInfo.Title = "go-example:ONLINE API"
		docs.SwaggerInfo.Description = "This is Demo server online restFull api ."
		docs.SwaggerInfo.Version = "v0.1"
		rootRouter.GET("/api/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	
	
	```

### 二、存储

- go-redis

    ```
    这里使用的是goredis, 支持单机模式，集群模式，哨兵模式,根据配置文件判断使用某种模式。库：https://github.com/go-redis/redis
    ```
  example:
  
    ```
    func ClientTest(t *testing.T) {
        client, err := NewRedisClient()
        if err != nil {
            panic(err)
        }
    
        err = client.Set("app", "go-example", 0).Err()
        if err != nil {
            panic(err)
        }
    
        err = client.Get("app").Err()
        if err != nil {
            panic(err)
        }
    }
    ```
    
    
### 三、中间件

- 日志（golog）

	```
	使用第三方库：
	
	https://github.com/kataras/golog
	https://github.com/lestrrat-go/file-rotatelogs
	
	```
	- 支持日志分割，按天分割，按天保存，保存时间可配置。参考yaml文件
	
	```
	logger:
	  #gin模式
	  mode: debug
	  #日志级别
	  level: debug
	  #是否写入到文件
	  isOutPutFile: false
	  #保存期限
	  maxAgeDay: 7
	  #按天保存
	  rotationTime: 1
	
	```

	

- 监控
   - 支持pprof，端口9123
   - prometheus metric

    ```
    增加prometheus的metric需要的指标，指标在middlewares/ginprometheus.go
    
    /**
        从url上获取参数，放在MAP中，用来监控每个url的调用情况
    
    */
      
    func init() {
        UriParamsMap["agfid"] = ":agfid"
        UriParamsMap["queryType"] = ":queryType"
        UriParamsMap["paramValue"] = ":paramValue"
        UriParamsMap["dgst"] = ":dgst"
    }
    
    ```
- I18N

	支持国际化，自定义本地静态模板。tomls目录和common/e目录

### 四、配置文件

application.yaml

```
server:
  #若port不设置值，默认为8188，这个值必须
  port: 8188
  #应用的访问的根目录
  contextPath: /
  dataPath: data
  pprof: false
  pprofPort: 9123
	
redis:
  #redis地址
  addrs: ["localhost:6379"]
  #密码
  pwd: ""
  #线程池
  PoolSize: 1000
  #选择库，默认0
  db: 0
  #是否是哨兵模式
  isSentinel: false
  #哨兵模式master name
  masterName: ""
  #哨兵地址eg: ["127.0.0.1:26379","127.0.0.2:26379","127.0.0.3:26379"]
  sentinelAddrs: []
	
logger:
  #gin模式
  mode: debug
  #日志级别
  level: debug
  #是否写入到文件
  isOutPutFile: false
  #保存期限
  maxAgeDay: 7
  #按天保存
  rotationTime: 1
	
```

### 五、Docker

Dockerfile

```
FROM alpine:3.9
COPY ./go-example .
COPY ./application.yaml .
COPY ./tomls ./tomls
EXPOSE 8188
ENTRYPOINT ["/go-example"]
```

build

```
CGO_ENABLED=0 GOARCH=amd64  GOOS=linux go build -a -v -installsuffix cgo -o go-example .

#docker build --no-cache -t {镜像仓库}/{项目}/go-example:v0.0.1 .
#docker push {镜像仓库}/{项目}/go-example:v0.0.1
```
