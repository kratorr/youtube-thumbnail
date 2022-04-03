package service

import (
	"errors"
	"log"
	url "net/url"

	"youtube-thumbnail/internal/repository"
)

type ThumbnailDownloader interface {
	Download(r string) ([]byte, bool, error)
}

type Service struct {
	ThumbnailDownloader
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ThumbnailDownloader: NewThumbnailDownloader(repos.Thumbnail, repos.Cache),
	}
}

type ThumbnailDownloaderService struct {
	repo  repository.Thumbnail
	cache repository.Cache
}

func NewThumbnailDownloader(repo repository.Thumbnail, cache repository.Cache) *ThumbnailDownloaderService {
	return &ThumbnailDownloaderService{repo: repo, cache: cache}
}

func (s *ThumbnailDownloaderService) Download(request string) ([]byte, bool, error) {
	address, err := url.ParseRequestURI(request)
	if err != nil {
		log.Println("failed to parse url")

		return nil, false, errors.New("failed to parse url")
	}

	videoID := address.Query().Get("v")

	if videoID == "" {
		log.Println("Bad query params")

		return nil, false, errors.New("bad query params")
	}

	if s.cache.InCache(videoID) {
		log.Println("Cache hit", videoID)
		image := s.cache.Get(videoID)

		return image, true, nil
	}

	image, err := s.repo.Download(videoID)
	if err != nil {
		log.Println("Failed to download thumbnail")

		return nil, false, errors.New("failed to download thumbnail")
	}

	s.cache.Insert(videoID, image)

	return image, false, nil
}
