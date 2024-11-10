package base

import (
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/rs/zerolog/log"
)

type NodeMaster struct{
	active_nodes map[string]*Node
	active_nodes_count int

	BASE_PORT int //42069
}

type Node struct {
	host string
	port int

	id int //later generate it from host and port
	rpc_server *rpc.Server
}

func NewNode(host string, port int) *Node {

	rpc_server := rpc.NewServer()

	return &Node {
		host: host,
		port: port,
		rpc_server: rpc_server,
	}
}

func (node *Node) Start() {

	listener, err := net.Listen("tcp", node.host+":"+strconv.Itoa(node.port))
	if err != nil {
		log.Fatal().Err(err)
	}
	defer listener.Close()

	node.rpc_server.HandleHTTP("/rpc", "/debug")
	log.Info().Msgf("Starting %s:%d\n", node.host, node.port)
	http.Serve(listener, nil)
}

func (node *Node) AttachObj(obj any) {

}
