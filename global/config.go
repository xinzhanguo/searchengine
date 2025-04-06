package global

// Config 服务器设置
type Config struct {
	Addr        string `yaml:"addr" json:"addr"`               // 监听地址
	Data        string `yaml:"data" json:"data"`               // 数据目录
	Debug       bool   `yaml:"debug" json:"debug"`             // 调试模式
	Dictionary  string `yaml:"dictionary" json:"dictionary"`   // 字典路径
	EnableAdmin bool   `yaml:"enableAdmin" json:"enableAdmin"` //启用admin
	Gomaxprocs  int    `yaml:"gomaxprocs" json:"gomaxprocs"`   //GOMAXPROCS
	Shard       int    `yaml:"shard" json:"shard"`             //分片数
	Auth        string `yaml:"auth" json:"auth"`               //认证
	EnableGzip  bool   `yaml:"enableGzip" json:"enableGzip"`   //是否开启gzip压缩
	Timeout     int64  `yaml:"timeout" json:"timeout"`         //超时时间
	BufferNum   int    `yaml:"bufferNum" json:"bufferNum"`     //分片缓冲数
}
