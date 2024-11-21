package protocol

type RPCIPRequest struct{
	ToId int
	FromId int
	RequestNodeId int
}

type RPCIPResponse struct{
	ToId int
	FromId int
	RequestNodeIP string
}

type RPCNodeRegisterRequest struct {
	NodeId int
	NodeIp string
}

