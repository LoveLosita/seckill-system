package items

import (
	items "kitex-server/items/kitex_gen/items/itemservice"
	"log"
)

func Start() {
	svr := items.NewServer(new(ItemServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
