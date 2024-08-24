package option

type Server struct {
	Host string
	Port int
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
	return func(server *Server) {
		server.Host = host
	}
}

func WithPort(port int) ServerOption {
	return func(server *Server) {
		server.Port = port
	}
}

func NewServer(options ...ServerOption) *Server {
	server := &Server{}
	for _, option := range options {
		option(server)
	}
	return server
}
