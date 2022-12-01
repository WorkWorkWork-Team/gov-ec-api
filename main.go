package main

import (
	"fmt"
	"os"

	"github.com/WorkWorkWork-Team/common-go/databasemysql"
	"github.com/WorkWorkWork-Team/common-go/httpserver"
	"github.com/WorkWorkWork-Team/gov-ec-api/config"
	"github.com/WorkWorkWork-Team/gov-ec-api/handler"
	"github.com/WorkWorkWork-Team/gov-ec-api/repository"
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
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
	submitMpRepository := repository.NewSubmitMpRepository(mysql)

	// New Service
	populationService := service.NewPopulationService(populationRepository)
	submitmpService := service.NewSubmitmpService(submitMpRepository, populationRepository)

	// New Handler
	populationHandler := handler.NewPopulationHandler(populationService)
	submitmpHandler := handler.NewSubmitMpHandler(submitmpService)

	server := httpserver.NewHttpServer(appConfig.PROXY_URL)
	api := server.Group(appConfig.PROXY_URL)
	api.Use(handler.ValidateAPIKey(appConfig.API_KEY))
	{
		api.GET("/population/statistic/", populationHandler.GetPopulationStatistics)
		api.POST("/mp/submit/", submitmpHandler.SubmitMp)
		api.GET("/candidate/", populationHandler.GetAllCandidateInfo)
	}
	server.Run(fmt.Sprint(":", appConfig.LISTENING_PORT))
}
