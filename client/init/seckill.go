package init_client

import (
	seckillservice2 "client/kitex-gens/seckill/kitex_gen/seckill/seckillservice"
	"github.com/cloudwego/kitex/client"
)

var NewSecKillClient seckillservice2.Client

func InitSecKillSvClient() error {
	var err error
	NewSecKillClient, err = seckillservice2.NewClient("seckillservice", client.WithHostPorts("0.0.0.0:8891"))
	if err != nil {
		return err
	}
	return nil
}
