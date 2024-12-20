package consumer

import (
	"ValidStudio/global"
)

func LogoutReceive() {
	var channel = "logout"
	subscribe := global.RDB.Subscribe(global.Background, channel)
	for msg := range subscribe.Channel() {
		username := msg.Payload
		token := global.RDB.Get(global.Background, username).Val()
		global.Cache.Delete(username)
		global.Cache.Delete(token)
	}
}
