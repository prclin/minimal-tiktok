package config

type Server struct {
	Port        int
	ContextPath string
}

var DefaultServer = &Server{
	Port:        8080,
	ContextPath: "/",
}
