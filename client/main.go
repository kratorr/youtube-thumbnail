package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
	"sync"

	"youtube-thumbnail/thumbnail"

	"google.golang.org/grpc"
)

func download(client thumbnail.ThubmnailDownloaderClient, inputUrl string) {
	address, err := url.ParseRequestURI(inputUrl)
	if err != nil {
		log.Println("failed to parse url")
	}

	videoID := address.Query().Get("v")
	ctx := context.Background()
	image, err := client.Download(ctx,
		&thumbnail.Request{
			URL: inputUrl,
		})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(fmt.Sprintf("./%s.jpg", videoID), image.Image, 0o644)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	serverIP := flag.String("ip", "127.0.0.1", "ip address grpc server")
	serverPort := flag.Int("port", 8081, "port grcp server")
	asyncPtr := flag.Bool("async", false, "enale async download")

	flag.Parse()

	urls := flag.Args()
	fmt.Println(urls)
	grcpConn, err := grpc.Dial(
		net.JoinHostPort(*serverIP, strconv.Itoa(*serverPort)), grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	client := thumbnail.NewThubmnailDownloaderClient(grcpConn)

	if *asyncPtr == true {
		wg := &sync.WaitGroup{}
		for _, url := range urls {
			wg.Add(1)
			url := url
			go func() {
				defer wg.Done()
				download(client, url)
			}()

		}
		wg.Wait()
	} else {
		for _, url := range urls {
			fmt.Println(url)
			download(client, url)
		}
	}
}
