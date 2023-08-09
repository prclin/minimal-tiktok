package config

type Datasource struct {
	Mysql *Mysql
}

type Mysql struct {
	//DSN 为datasource name 格式为"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DSN string
	//最大空闲连接数
	MaxIdleConnections int
	//最大开启连接数
	MaxOpenConnections int
	//连接最大存活时间 秒
	ConnectionMaxLifetime int
}

var DefaultDataSource = &Datasource{
	Mysql: &Mysql{
		DSN:                   "root:20020716@tcp(192.168.200.129:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local",
		MaxIdleConnections:    10,
		MaxOpenConnections:    20,
		ConnectionMaxLifetime: 300,
	},
}
