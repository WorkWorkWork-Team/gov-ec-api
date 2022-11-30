package handler_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/WorkWorkWork-Team/gov-ec-api/handler"
	"github.com/WorkWorkWork-Team/gov-ec-api/model"
	"github.com/WorkWorkWork-Team/gov-ec-api/repository"
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	PopulationHander *handler.PopulationHandler
)

var _ = Describe("User Integration Test", Label("integration"), func() {
	BeforeEach(func() {
		populationRepository := repository.NewPopulationRepository(MySQLConnection)
		populationService := service.NewPopulationService(populationRepository)
		PopulationHandler = handler.NewPopulationHandler(populationService)
	})

	Context("Population Info API", func() {
		Context("Database have population data", func() {
			When("Voter EC want to get Population INFO", func() {
				It("should return success.", func() {
					// Expect no user in voted table
					var populationList []model.PopulationDatabaseRow
					row := MySQLConnection.Select(&populationList, "SELECT * FROM Population")
					Expect(row).ShouldNot(HaveOccurred())
					populationLength := len(populationList)
					Expect(populationLength).To(Equal(2))

					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					c.Request = httptest.NewRequest(http.MethodPost, "/population/statistic/", nil)
					PopulationHandler.GetPopulationStatistics(c)
					Expect(res.Result().StatusCode).To(Equal(http.StatusOK))
				})
			})
		})
	})
})
