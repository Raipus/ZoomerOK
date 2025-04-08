package caching

import (
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/patrickmn/go-cache"
)

func initCache() {
	c := cache.New(config.Config.TimeCache*time.Second, config.Config.TimeCache*time.Second)
	return c
}

var c *cache.Cache = initCache()

func SetCacheResetLink(username, resetLink string) {
	c.Set(config.Config.ResetCache+"_"+resetLink, username, cache.DefaultExpiration)
}

func SetCacheConfirmationLink(username, confirmationLink string) {
	c.Set(config.Config.ConfirmationCache+"_"+confirmationLink, username, cache.DefaultExpiration)
}

func GetCacheResetLink(resetLink string) string {
	resetLink, found := c.Get(config.Config.ConfirmationCache + "_" + resetLink)
	if found {
		return resetLink
	}

	return ""
}

func GetCacheConfirmationLink(confirmationLink string) string {
	confirmationLink, found := c.Get(config.Config.ConfirmationCache + "_" + confirmationLink)
	if found {
		return confirmationLink
	}

	return ""
}

func DeleteCacheResetLink(resetLink string) {
	c.Delete(config.Config.ResetCache + "_" + resetLink)
}

func DeleteCacheConfirmationLink(confirmationLink string) {
	c.Delete(config.Config.ConfirmationCache + "_" + confirmationLink)
}
