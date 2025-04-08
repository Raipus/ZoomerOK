package caching

import (
	"time"

	"github.com/Raipus/ZoomerOK/account/pkg/config"
	"github.com/patrickmn/go-cache"
)

// CachingInterface определяет методы для кэширования
type CachingInterface interface {
	SetCacheResetLink(username, resetLink string)
	SetCacheConfirmationLink(username, confirmationLink string)
	GetCacheResetLink(resetLink string) string
	GetCacheConfirmationLink(confirmationLink string) string
	DeleteCacheResetLink(resetLink string)
	DeleteCacheConfirmationLink(confirmationLink string)
}

// ProductionCachingInterface - экземпляр для использования в производстве
var ProductionCachingInterface CachingInterface = &RealCache{cache: initCache()}

// RealCache - реальная реализация кэша с использованием go-cache
type RealCache struct {
	cache *cache.Cache
}

// initCache - инициализация кэша
func initCache() *cache.Cache {
	return cache.New(time.Duration(config.Config.TimeCache)*time.Second, time.Duration(config.Config.TimeCache)*time.Second)
}

// Реализация методов интерфейса CachingInterface
func (r *RealCache) SetCacheResetLink(username, resetLink string) {
	key := config.Config.ResetCache + "_" + resetLink
	r.cache.Set(key, username, cache.DefaultExpiration)
}

func (r *RealCache) SetCacheConfirmationLink(username, confirmationLink string) {
	key := config.Config.ConfirmationCache + "_" + confirmationLink
	r.cache.Set(key, username, cache.DefaultExpiration)
}

func (r *RealCache) GetCacheResetLink(resetLink string) string {
	key := config.Config.ResetCache + "_" + resetLink
	value, found := r.cache.Get(key)
	if !found {
		return ""
	}
	strResetLink, ok := value.(string)
	if !ok {
		return ""
	}
	return strResetLink
}

func (r *RealCache) GetCacheConfirmationLink(confirmationLink string) string {
	key := config.Config.ConfirmationCache + "_" + confirmationLink
	value, found := r.cache.Get(key)
	if !found {
		return ""
	}
	strConfirmationLink, ok := value.(string)
	if !ok {
		return ""
	}
	return strConfirmationLink
}

func (r *RealCache) DeleteCacheResetLink(resetLink string) {
	key := config.Config.ResetCache + "_" + resetLink
	r.cache.Delete(key)
}

func (r *RealCache) DeleteCacheConfirmationLink(confirmationLink string) {
	key := config.Config.ConfirmationCache + "_" + confirmationLink
	r.cache.Delete(key)
}
