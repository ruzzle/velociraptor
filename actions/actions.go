// Client actions are routines that run on the client and return a
// GrrMessage.
package actions

import (
	"www.velocidex.com/golang/velociraptor/context"
	crypto_proto "www.velocidex.com/golang/velociraptor/crypto/proto"
)

type ClientAction interface {
	Run(ctx *context.Context,
		args *crypto_proto.GrrMessage) []*crypto_proto.GrrMessage
}


func GetClientActionsMap() map[string]ClientAction {
	result := make(map[string]ClientAction)
	result["GetClientInfo"] = &GetClientInfo{}
	return result
}
