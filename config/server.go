package config

type Server struct {
	Port        int
	ContextPath string
}

var DefaultServer = &Server{
<<<<<<< Updated upstream
	Port:        8080,
	ContextPath: "/",
=======
	Port:        7070,
	ContextPath: "/douyin",
>>>>>>> Stashed changes
}
