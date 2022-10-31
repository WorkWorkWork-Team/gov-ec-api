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
	submitMpRepository := repository.NewSubmitMpRepository(mysql)

	submitmpService := service.NewSubmitmpService(submitMpRepository)

	submitmpHandler := handler.NewSubmitMpHandler(submitmpService)

	server := gin.Default()
	server.Use(handler.ValidateAPIKey(appConfig.API_KEY))
	{
		server.POST("/mp/submit/", submitmpHandler.SubmitMp)
	}
	server.Run(fmt.Sprint(":", appConfig.LISTENING_PORT))
}
