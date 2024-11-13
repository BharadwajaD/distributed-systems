package base

//using rpc to get ipc
type RPCMessageRequest struct {
	Content string
}

func NewRPCMessage(content string) *RPCMessageRequest {
	return &RPCMessageRequest {Content: content}
}

type RPCMessageResponse struct {
	Content string
}

type RPCMessage struct {
}

//sender node will call RpcReceive function of the receiver node
//call the respective function base on request message
func (rpc_msg *RPCMessage) RpcReceive(request *RPCMessageRequest, response *RPCMessageResponse) error {
	response.Content = request.Content
	return nil
}

