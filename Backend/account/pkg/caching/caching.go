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
	c.Set(config.Config.ResetCache+"_"+username, resetLink, cache.DefaultExpiration)
}

func SetCacheConfirmationLink(username, confirmationLink string) {
	c.Set(config.Config.ConfirmationCache+"_"+username, confirmationLink, cache.DefaultExpiration)
}

func GetCacheResetLink(username string) string {
	resetLink, found := c.Get(config.Config.ConfirmationCache + "_" + username)
	if found {
		return resetLink
	}

	return ""
}

func GetCacheConfirmationLink(username string) string {
	confirmationLink, found := c.Get(config.Config.ConfirmationCache + "_" + username)
	if found {
		return confirmationLink
	}

	return ""
}
