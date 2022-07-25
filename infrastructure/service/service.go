package service

import (
	"akawork.io/infrastructure/cache"
)

/**
 * Defines a BaseService
 */
type BaseService struct {
	cache.CacheManager
}
