package base

import (
	"net"
	"net/http"
	"net/rpc"
	"strconv"

	"github.com/rs/zerolog/log"

)

const NODE_MANAGER_ID = 0
const NODE_MANAGER_ADDR = "127.0.0.1:42069"

type Connector struct {
	host string
	port int
	id int //later generate it from host and port

	rpc_server *rpc.Server
	rpc_client *rpc.Client //TODO: deal with dis later

	node_manager_address string
}

func NewConnector(host string, port int, nm_address string) *Connector {

	rpc_server := rpc.NewServer()

	connector :=  &Connector {
		host: host,
		port: port,
		rpc_server: rpc_server,
		node_manager_address: nm_address,
	}

	//TODO: Register request node.Send()
	return connector
}

// HTTP Handler -> Handles http connections to this host:port
func (connector *Connector) Start() {

	listener, err := net.Listen("tcp", connector.host+":"+strconv.Itoa(connector.port))
	if err != nil {
		log.Fatal().Err(err)
	}
	defer listener.Close()

	connector.rpc_server.HandleHTTP("/rpc", "/debug")
	log.Info().Msgf("Starting %s:%d\n", connector.host, connector.port)
	http.Serve(listener, nil)
}

func (connector *Connector) RegisterService(service any) {
	connector.rpc_server.Register(service)
}
