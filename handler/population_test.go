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

	Context("Get Candidates", func() {
		Context("Database have population data", func() {
			When("Population table exists", func() {
				It("Should return success", func() {
					var candidateList []model.PopulationDatabaseRow
					query := `
						SELECT p.CitizenID, LazerID, Name, Lastname, Birthday, Nationality, DistrictID
						FROM Population AS p
						JOIN Candidate AS c
						ON p.CitizenID = c.CitizenID
					`
					err := MySQLConnection.Select(&candidateList, query)
					Expect(err).ShouldNot(HaveOccurred())
					candidateListLength := len(candidateList)
					Expect(candidateListLength).Should(Equal(2))

					// call API
					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					PopulationHandler.GetAllCandidateInfo(c)

					// Expect 200
					Expect(c.Writer.Status()).To(Equal(http.StatusOK))
				})
			})
		})
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
					c.AddParam("CitizenID", TestUserCitizenID)
					PopulationHandler.GetPopulationStatistics(c)
					Expect(res.Result().StatusCode).To(Equal(http.StatusOK))
				})
			})
		})
	})
})
