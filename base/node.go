package base

import (
	"fmt"
	"io"
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
	rw io.ReadWriteCloser
	rpc_server *rpc.Server
	rpc_client *rpc.Client //TODO: deal with dis later
}

func NewNode(host string, port int) *Node {

	rpc_server := rpc.NewServer()
	//rpc_client := rpc.NewClient()
	rpc_message := new(RPCMessage)
	rpc_server.Register(rpc_message)

	return &Node {
		host: host,
		port: port,
		rpc_server: rpc_server,
	}
}

// HTTP Handler -> Handles http connections to this host:port
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

//used to send messages to other nodes
//TODO: multiple type of messages (if needed)
func (node *Node) Send(to_addr string, msg *RPCMessageRequest) (*RPCMessageResponse, error){
	reply := new(RPCMessageResponse)
	client, err := rpc.DialHTTPPath("tcp", to_addr, "/rpc")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = client.Call("RPCMessage.RpcReceive", msg, reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

