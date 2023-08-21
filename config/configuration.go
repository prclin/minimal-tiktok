package config

/*
Configuration 对应整个application.yaml配置文件
*/
type Configuration struct {
	Server     *Server
	Zap        *Zap
	Datasource *Datasource
	Jwt        *Jwt
	OSS        *OSS
}

var DefaultConfiguration = &Configuration{
	Server:     DefaultServer,
	Zap:        DefaultZap,
	Datasource: DefaultDataSource,
	Jwt:        DefaultJwt,
	OSS:        DefaultOSS,
}
