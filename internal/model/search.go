// model/search.go

package model

import (
	"github.com/rsdel2007/proj/contract"
	"github.com/sahilm/fuzzy"
	"gorm.io/gorm"
)

func GetPaginatedVideos(db *gorm.DB, offset, pageSize int) []*contract.Video {
	var videos []*contract.Video
	db.Order("published_at DESC").Offset(offset).Limit(pageSize).Find(&videos)
	return videos
}

func SearchVideos(db *gorm.DB, query string, offset, pageSize int) ([]*contract.Video, error) {
	var videos []*contract.Video

	var allVideos []*contract.Video
	db.Order("published_at DESC").Find(&allVideos)
	//fmt.Println(len(allVideos))
	matches := fuzzy.Find(query, getSearchableStrings(allVideos))
	//fmt.Println("len: ", len(matches))
	var matchingVideoIDs []int
	for _, match := range matches {
		matchingVideoIDs = append(matchingVideoIDs, allVideos[(match.Index-1)/2].ID)
	}

	// Retrieve the matching videos based on fuzzy search
	db.Find(&videos, matchingVideoIDs)

	return videos, nil
}

func getSearchableStrings(videos []*contract.Video) []string {
	var searchableStrings []string
	for _, video := range videos {
		searchableStrings = append(searchableStrings, video.Title, video.Description)
	}
	return searchableStrings
}
