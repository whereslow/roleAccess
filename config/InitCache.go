package config

import (
	"ValidStudio/global"
	"github.com/patrickmn/go-cache"
	"time"
)

func InitCache() {
	global.Cache = cache.New(60*time.Minute, 120*time.Minute)
}
