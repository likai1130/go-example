package mongodb

import (
	"context"
	"fmt"
	"go-example/common/logger"
	"go-example/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

const MONGODB_CONNECT_TIMEOUT  = 20 * time.Second //连接超时时间

var mongoCli *mongo.Client
var once sync.Once

func NewMongoCliInstance() (mongoCli *mongo.Client, err error) {
	once.Do(func() (){
		mongoCli,err = setUp()
	})
	return mongoCli,err
}

func setUp() (mongoCli *mongo.Client, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mConf := config.AppConfig.MongoConf
	clientOptions := options.Client().SetHosts(mConf.Hosts).
		SetMaxPoolSize(mConf.MaxPoolSize).         //最大连接数量
		SetConnectTimeout(MONGODB_CONNECT_TIMEOUT) //连接超时20s
	if mConf.UserName != "" && mConf.Password != "" {
		clientOptions.SetAuth(options.Credential{Username: mConf.UserName, Password: mConf.Password})
	}

	mongoCli, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("MongoDB connect fail：%s", err.Error())
	}

	err = mongoCli.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("MongoDB ping fail：%s", err.Error())
	}

	logger.GetLogger().Info("MongoDB connect success !")
	return mongoCli, nil
}

func Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoCli.Disconnect(ctx)
}

func SetUp() {
	if _, err := NewMongoCliInstance(); err != nil {
		panic(err)
	}
}