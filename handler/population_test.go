package handler_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/WorkWorkWork-Team/gov-ec-api/handler"
	"github.com/WorkWorkWork-Team/gov-ec-api/model"
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/WorkWorkWork-Team/gov-voter-api/handler"
	model "github.com/WorkWorkWork-Team/gov-voter-api/models"
	"github.com/WorkWorkWork-Team/gov-voter-api/repository"
	"github.com/WorkWorkWork-Team/gov-voter-api/service"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	TestUserCitizenID   string = "1234567891235"
	TestJWTSecretKey    string = "key"
	TestJWTIssuer       string = "Tester"
	TestJWTTTL          int    = 10
	UserHandler         *handler.UserHandler
	AuthenticateHandler *handler.AuthenticateHandler
)

var _ = Describe("User Integration Test", Label("integration"), func() {
	BeforeEach(func() {
		populationRepository := repository.NewPopulationRepository(MySQLConnection)
		populationService := service.NewPopulationService(populationRepository)
		PopulationHandler := handler.NewPopulationHandler(populationService)
	})

	Context("Population Info API", func() {
		Context("Database have population data", func() {
			When("Voter EC want to get Population INFO", func() {
				It("should return success.", func() {
					// Expect no user in voted table
					var populationList []model.PopulationDatabaseRow
					err := MySQLConnection.Select(&populationList, "SELECT * FROM Population")
					Expect(err).ShouldNot(HaveOccurred())
					populationLength := len(populationList)
					Expect(populationLength).To(Equal(0))

					res := httptest.NewRecorder()
					c, r := gin.CreateTestContext(res)
					c.Request = httptest.NewRequest(http.MethodPost, "/api", nil)

					r.POST("/api", PopulationHandler.GetPopulationStatistics)

					r.ServeHTTP(res, c.Request)
				})
			})
		})
	})
})
