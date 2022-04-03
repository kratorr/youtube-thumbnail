package handler

import (
	"context"
	"fmt"

	"youtube-thumbnail/internal/service"
	"youtube-thumbnail/thumbnail"

	"google.golang.org/grpc"
)

type ThubmnailDownloaderServer struct {
	thumbnail.UnimplementedThubmnailDownloaderServer
	services *service.Service
}

func (ds *ThubmnailDownloaderServer) Download(ctx context.Context, r *thumbnail.Request) (*thumbnail.Response, error) {
	fmt.Println("Download", r.URL)
	image, cachtHit, err := ds.services.ThumbnailDownloader.Download(r.URL)
	if err != nil {
		return nil, err
	}

	response := thumbnail.Response{ID: r.URL, CacheHit: cachtHit, Image: image}
	return &response, nil
}

func (s *ThubmnailDownloaderServer) mustEmbedUnimplementedThubmnailDownloaderServer() {}

func NewHandler(server *grpc.Server, services *service.Service) {
	thumbnailGrpcServer := &ThubmnailDownloaderServer{services: services}
	thumbnail.RegisterThubmnailDownloaderServer(server, thumbnailGrpcServer)
}
