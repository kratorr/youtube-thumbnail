package repository

import "sync"

type Cache interface {
	InCache(string) bool
	Get(string) []byte
	Insert(string, []byte)
}

type Thumbnail interface {
	Download(videoID string) ([]byte, error)
}

type Repository struct {
	Thumbnail
	Cache
}

func NewRepository(cache map[string][]byte, mu *sync.RWMutex) *Repository {
	return &Repository{
		Thumbnail: NewThumbnail(),
		Cache:     NewCacheInmemory(cache, mu),
	}
}
