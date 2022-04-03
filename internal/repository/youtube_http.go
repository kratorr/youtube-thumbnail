package repository

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://img.youtube.com/vi/%s/maxresdefault.jpg"

type ThumbnailYoutubeHTTP struct{}

func NewThumbnail() *ThumbnailYoutubeHTTP {
	return &ThumbnailYoutubeHTTP{}
}

func (r *ThumbnailYoutubeHTTP) Download(id string) ([]byte, error) {
	respone, _ := http.Get(fmt.Sprintf(url, id))
	defer respone.Body.Close()
	// TODO обработать ошибку, контекст наверно прихерачить
	body, _ := io.ReadAll(respone.Body)

	return body, nil
}
