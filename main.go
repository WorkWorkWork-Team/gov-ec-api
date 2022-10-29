package main

import (
	"fmt"
	"os"

	"github.com/WorkWorkWork-Team/common-go/databasemysql"
	"github.com/WorkWorkWork-Team/gov-ec-api/config"
	"github.com/WorkWorkWork-Team/gov-ec-api/handler"
	"github.com/WorkWorkWork-Team/gov-ec-api/repository"
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/gin-gonic/gin"
)

var appConfig config.Config

func init() {
	appConfig = config.Load()
}

func main() {
	mysql, err := databasemysql.NewDbConnection(databasemysql.Config{
		Hostname:     fmt.Sprint(appConfig.MYSQL_HOSTNAME, ":", appConfig.MYSQL_PORT),
		Username:     appConfig.MYSQL_USERNAME,
		Password:     appConfig.MYSQL_PASSWORD,
		DatabaseName: appConfig.MYSQL_DATABASE,
	})
	if err != nil {
		os.Exit(1)
		return
	}
	// New Repository
	populationRepository := repository.NewPopulationRepository(mysql)

	populationService := service.NewPopulationService(populationRepository)

	populationHandler := handler.NewPopulationHandler(populationService)

	server := gin.Default()
	server.GET("/population/statistic/", populationHandler.GetPopulationStatistics)
	server.Run(fmt.Sprint(":", appConfig.LISTENING_PORT))
}
