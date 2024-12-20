package config

import (
	"ValidStudio/global"
	"github.com/patrickmn/go-cache"
	"time"
)

func InitCache() {
	global.Cache = cache.New(28*time.Minute, 30*time.Minute)
}
