package consumer

import (
	"ValidStudio/global"
	"math/rand/v2"
	"time"
)

func LoginReceive() {
	var channel = "login"
	subscribe := global.RDB.Subscribe(global.Background, channel)
	for msg := range subscribe.Channel() {
		username := msg.Payload
		token := global.RDB.Get(global.Background, username).Val()
		role := global.RDB.Get(global.Background, token).Val()
		t := time.Duration(2*3600000000000 + rand.IntN(10)*3600000000000)
		global.Cache.Set(username, token, t)
		global.Cache.Set(token, role, t)
	}
}
