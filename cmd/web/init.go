package main

import (
	"fmt"
	configCommon "github.com/aryahmph/url-shortener/common/config"
	base62Manager "github.com/aryahmph/url-shortener/common/hash/base62"
	httpCommon "github.com/aryahmph/url-shortener/common/http"
	mongodbCommon "github.com/aryahmph/url-shortener/common/mongodb"
	zookeeperCommon "github.com/aryahmph/url-shortener/common/zookeeper"
	linkDelivery "github.com/aryahmph/url-shortener/internal/delivery/link/http"
	counterRepo "github.com/aryahmph/url-shortener/internal/repository/counter/zookeeper"
	linkRepo "github.com/aryahmph/url-shortener/internal/repository/link/mongodb"
	linkUCase "github.com/aryahmph/url-shortener/internal/usecase/link"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func main() {
	config := configCommon.LoadConfig()
	db := mongodbCommon.New(
		config.MongoUser, config.MongoPassword,
		config.MongoHost, config.MongoPort, config.MongoName,
	)
	zk := zookeeperCommon.New(config.ZookeeperHosts)
	defer zk.Close()

	hashManager := base62Manager.NewBase62Manager()

	counterRepository := counterRepo.NewZookeeperCounterRepository(zk)
	linkRepository := linkRepo.NewMongodbLinkRepository(db.Collection("links"))
	linkUsecase := linkUCase.NewLinkUsecase(linkRepository, counterRepository, hashManager)

	httpServer := httpCommon.NewHTTPServer()
	api := httpServer.E.Group("/api/v1", middleware.Logger(), middleware.CORS())

	linkDelivery.NewHTTPLinkDelivery(api, linkUsecase)

	httpServer.E.Logger.Fatal(httpServer.E.StartServer(&http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}))
}
