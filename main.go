package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

var (
	BASE_IMAGE_FOLDER = os.Getenv("_BASE_IMAGE_FOLDER")
	IMAGE_SUFFIX      = ".jpg"
	ENCRYPTION_KEY    = os.Getenv("_ENCRYPTION_KEY")
	HOST              = os.Getenv("_HOST")
	BASE_URL          = os.Getenv("_BASE_URL")
)

func main() {
	connectMysql()
	router := gin.Default()

	store := persistence.NewInMemoryStore(time.Second)

	router.POST("/upload", upload)
	router.GET("/images/:image", cache.CachePage(store, time.Minute, serveImage))
	router.GET("/encrypt/:id", encryptId) // Only for test

	router.Run(HOST)
}
