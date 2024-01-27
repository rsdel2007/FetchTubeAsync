package youtube

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/rsdel2007/proj/contract"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"gorm.io/gorm"
)

var (
	searchQuery = "Cricket"
	apiKeys     []string
	apiKeysMu   sync.Mutex
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKeys = strings.Split(os.Getenv("APIKEY"), ",")
}

func FetchAndStoreVideos(db *gorm.DB) {
	// Initialize YouTube service
	service, err := initializeYouTubeService()
	if err != nil {
		log.Printf("Error initializing YouTube service: %v", err)
		return
	}

	// Fetch latest videos from YouTube API
	videos, err := fetchLatestVideos(service)
	if err != nil {
		log.Printf("Error fetching videos from YouTube API: %v", err)
		return
	}

	storeVideos(db, videos)
}

func initializeYouTubeService() (*youtube.Service, error) {
	apiKeysMu.Lock()
	defer apiKeysMu.Unlock()

	for _, apiKey := range apiKeys {
		client := &http.Client{
			Transport: &transport.APIKey{Key: apiKey},
		}

		service, err := youtube.New(client)
		if err == nil {
			return service, nil
		}

		log.Printf("Error creating YouTube client with API key %s: %v", apiKey, err)
	}

	return nil, fmt.Errorf("All API keys exhausted")
}

func fetchLatestVideos(service *youtube.Service) ([]*contract.Video, error) {
	ctx := context.Background()
	call := service.Search.List([]string{"snippet"}).
		Q(searchQuery).
		MaxResults(100)

	response, err := call.Context(ctx).Do()
	if err != nil {
		return nil, fmt.Errorf("Error fetching videos: %v", err)
	}

	var videos []*contract.Video
	for _, item := range response.Items {
		video := &contract.Video{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			PublishedAt: item.Snippet.PublishedAt,
			Thumbnail:   item.Snippet.Thumbnails.Default.Url,
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func storeVideos(db *gorm.DB, videos []*contract.Video) {
	for _, video := range videos {
		db.Create(video)
	}
}
