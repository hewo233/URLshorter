package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil, "Redis client should not be nil.")
}

func TestSaveAndRetrieveUrlMapping(t *testing.T) {

	//println("Begin : TestSaveAndRetrieveUrlMapping")

	originalUrl := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"

	userID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortUrl := "Jsz4k57oAX"

	SaveUrlMapping(shortUrl, originalUrl, userID)
	result := RetrieveInitUrl(shortUrl)

	assert.Equal(t, originalUrl, result, "The two urls should be the same.")
}
