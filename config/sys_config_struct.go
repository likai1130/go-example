package config

/***
 * 与yaml 对应的结构体
 */

type WebServer struct {
	Port        int    `profile:"port" profileDefault:"8188" `
	ContextPath string `profile:"contextPath" `
	DataPath    string `profile:"dataPath" `
	PProf       bool   `profile:"pprof" profileDefault:"true"`
	PProfPort   int    `profile:"pprofPort" profileDefault:"9123"`
}

type RedisConfig struct {
	Addrs         []string `json:"addrs"`
	Pwd           string   `json:"pwd" profileDefault:""`
	PoolSize      int      `json:"pool_size" profileDefault:"1000"`
	DB            int      `json:"db" profileDefault:"0"`
	IsSentinel    bool     `json:"is_sentinel" profileDefault:"false"`
	MasterName    string   `json:"master_name" profileDefault:""`
	SentinelAddrs []string `json:"sentinel_addrs"`
}

type Logger struct {
	Mode         string `prfile:"mode" profileDefault:"release" json:"mode"`
	Level        string `prfile:"level" profileDefault:"info" json:"level"`
	IsOutPutFile bool   `prfile:"isOutPutFile" profileDefault:"true" json:"isOutPutFile"`
	MaxAgeDay    int64  `profile:"maxAgeDay" profileDefault:"7" json:"maxAgeDay"`
	RotationTime int64  `profile:"rotationTime" profileDefault:"1" json:"rotationTime"`
}

/**
MongoDB 配置
*/
type MongoConf struct {
	Hosts       []string `profile:"hosts" profileDefault:"[127.0.0.1:27017]"`
	UserName    string   `profile:"userName"`
	Password    string   `profile:"password"`
	MaxPoolSize uint64      `profile:"maxPoolSize" profileDefault:"100" `
}

type ApplicationConfig struct {
	Server      WebServer   `profile:"server"`
	RedisConfig RedisConfig `profile:"redis"`
	Logger      Logger      `profile:"logger"`
	MongoConf MongoConf `profile:"mongodb"`
}
