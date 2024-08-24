package option

import "errors"

type Server2 struct {
	Host string
	Port int
}

type ServerOption2 func(*Server2) error

func WithHost2(host string) ServerOption2 {
	return func(server *Server2) error {
		server.Host = host
		return nil
	}
}

func WithPort2(port int) ServerOption2 {
	return func(server *Server2) error {
		if port <= 0 {
			return errors.New("invalid port number")
		}
		server.Port = port
		return nil
	}
}

func NewServer2(options ...ServerOption2) (*Server2, error) {
	server := &Server2{}
	for _, option := range options {
		err := option(server)
		if err != nil {
			return nil, err
		}
	}
	return server, nil
}
