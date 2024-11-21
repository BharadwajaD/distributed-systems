package base

import (
	"BharadwajaD/DistSys/pkg/base/protocol"
	"fmt"
)


type NodeManager struct{
	active_nodes map[int]string //node_id -> ip:port
	active_nodes_count int

	node *Connector
}

//Guarantee: this doesn't fail
func NewNodeManager(host string, port int) *NodeManager {

	nm := &NodeManager{
		active_nodes: make(map[int]string),
		active_nodes_count: 0,
		node: NewConnector(host, port, ""),
	}

	return nm
}

func (nm *NodeManager) RegisterNode(node_id int, node_ip string) {
	nm.active_nodes_count ++
	nm.active_nodes[node_id] = node_ip
}

func (nm *NodeManager) DeRegisterNode(node *Connector) {
	//TODO
}

func (nm *NodeManager) RPCNodeRegister(request *protocol.RPCNodeRegisterRequest, response *protocol.RPCSuccessResponse) error{

	nm.RegisterNode(request.NodeId, request.NodeIp)
	response.Body = fmt.Sprintf("Successfully added node: %d", request.NodeId)

	return nil
}

func (nm *NodeManager) RPCIdtoIP(request *protocol.RPCIPRequest, response *protocol.RPCIPResponse) error{
	response.ToId = request.FromId
	response.FromId = request.ToId

	response.RequestNodeIP = nm.active_nodes[request.RequestNodeId]

	return nil
}
