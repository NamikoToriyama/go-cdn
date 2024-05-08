package main

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

var (
	BASE_IMAGE_FOLDER = "img/"
	IMAGE_SUFFIX      = ".jpg"
	ENCRYPTION_KEY    = "TESTKEY"
	HOST              = "0.0.0.0:8080"
	BASE_URL          = "http://localhost:8080/"
	MYSQL_HOST        = "127.0.0.1"
	MYSQL_PORT        = "3306"   //`#mysql config`
	MYSQL_DB          = "go_cdn" //`#mysql config`
	MYSQL_USER        = "root"   //`#mysql config`
	MYSQL_PASS        = "root"   //`#mysql config`
)

func main() {
	ConnectMysql()
	router := gin.Default()

	store := persistence.NewInMemoryStore(time.Second)

	router.POST("/upload", upload)
	router.GET("/images/:image", cache.CachePage(store, time.Minute, serveImage))
	router.GET("/encrypt/:id", encryptId) // Only for test

	router.Run(HOST)
}
