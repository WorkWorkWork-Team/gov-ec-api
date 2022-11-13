package handler

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type populationHandler struct {
	service service.PopulationService
}

func NewPopulationHandler(populationService service.PopulationService) populationHandler {
	return populationHandler{
		service: populationService,
	}
}

func (p *populationHandler) errorMessage(err error, functionName string) {
	logrus.WithFields(logrus.Fields{
		"Module":  "Handler",
		"Funtion": functionName,
	}).Error(err)
}

func (p populationHandler) GetPopulationStatistics(g *gin.Context) {
	popData, err := p.service.GetPopulationStatistics()
	if err == nil {
		g.JSON(http.StatusOK, popData)
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
}

func (p populationHandler) GetAllCandidateInfo(g *gin.Context) {
	popData, err := p.service.GetAllCandidateInfo()
	if err == nil {
		g.JSON(http.StatusOK, popData)
		return
	} else if errors.Is(err, sql.ErrNoRows) {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "Not matching data",
		})
		return
	}
	g.JSON(http.StatusInternalServerError, gin.H{
		"message": "Something went wrong.",
	})
}
