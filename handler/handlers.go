package handler

import (
	"URLshorter/shortener"
	"URLshorter/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

const serverHost = "http://localhost:8081"

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to URL Shortener",
	})
}

type UrlCreateRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreatShortURL(c *gin.Context) {
	var createRequest UrlCreateRequest
	errBind := c.BindJSON(&createRequest)
	if errBind != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while binding request",
			"error":   errBind.Error(),
		})
		return
	}

	shortUrl := shortener.GenerateShortURL(createRequest.LongUrl, createRequest.UserId)
	store.SaveUrlMapping(shortUrl, createRequest.LongUrl, createRequest.UserId)

	host := serverHost
	c.JSON(200, gin.H{
		"message":   "Short URL created successfully",
		"short_url": host + "/" + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("short_url")
	longUrl := store.RetrieveInitUrl(shortUrl)
	c.Redirect(http.StatusMovedPermanently, longUrl)
}
