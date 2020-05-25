package api

import (
	"encoding/base64"
	"github.com/ybbus/jsonrpc"
)

var RPC jsonrpc.RPCClient

func InitRPC(url string, user string, password string) {
	RPC = jsonrpc.NewClientWithOpts(url, &jsonrpc.RPCClientOpts{
		CustomHeaders: map[string]string{
			"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+password)),
		},
	})
}
