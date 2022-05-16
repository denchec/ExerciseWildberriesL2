package pkg

import "fmt"

type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() Server {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (serv Server) GetType() string {
	return serv.Type
}

func (serv Server) PrintDetails() {
	fmt.Printf("%s Core: [%d], Memory: [%d] \n", serv.Type, serv.Core, serv.Memory)
}
