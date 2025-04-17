package init_client

import (
	"client/kitex-gens/users/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
)

var NewUserClient userservice.Client

func InitUserSvClient() error {
	var err error
	NewUserClient, err = userservice.NewClient("userservice", client.WithHostPorts("0.0.0.0:8889"))
	if err != nil {
		return err
	}
	return nil
}
