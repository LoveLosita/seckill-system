package init_client

import (
	itemservice2 "client/kitex-gens/items/kitex_gen/items/itemservice"
	"github.com/cloudwego/kitex/client"
)

var NewItemClient itemservice2.Client

func InitItemSvClient() error {
	var err error
	NewItemClient, err = itemservice2.NewClient("itemservice", client.WithHostPorts("0.0.0.0:8890"))
	if err != nil {
		return err
	}
	return nil
}
