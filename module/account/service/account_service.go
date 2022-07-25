package service

import (
	"time"

	"akawork.io/dto"
	"akawork.io/infrastructure/cache"
)

// Defines object ICacheService
type ICacheService interface {
	CreatePostinCache() error
	GetPostinCache() string
}

// Defines object CacheService
type CacheService struct {
	Cache   cache.CacheManager
	Timeout time.Duration
}

// Return a new CacheService
func NewCacheService(cache cache.CacheManager, timeout time.Duration) ICacheService {
	return &CacheService{
		Cache:   cache,
		Timeout: timeout,
	}
}

/**
 * Return function to create room
 */
func (service *CacheService) CreatePostinCache() error {
	flag := dto.ColorDto{}
	flag.Flag = "2022"
	service.Cache.Set("DSC", flag, 1*time.Hour)
	return nil
}

func (service *CacheService) GetPostinCache() string {
	data := service.Cache.Get("DSC")
	return data
}
