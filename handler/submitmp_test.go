package handler_test

import (
	"bytes"
	"encoding/json"
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
	TestUserCitizenID string = "1234567891235"
	TestJWTSecretKey  string = "key"
	TestJWTIssuer     string = "Tester"
	TestJWTTTL        int    = 10
	PopulationHandler *handler.PopulationHandler
	SubmitmpHandler   *handler.SubmitmpHandler
)

var _ = Describe("SubmitMP Integration Test", Label("integration"), func() {
	BeforeEach(func() {
		// New Repository
		populationRepository := repository.NewPopulationRepository(MySQLConnection)
		submitMpRepository := repository.NewSubmitMpRepository(MySQLConnection)

		// New Service
		populationService := service.NewPopulationService(populationRepository)
		submitmpService := service.NewSubmitmpService(submitMpRepository, populationRepository)

		// New Handler
		PopulationHandler = handler.NewPopulationHandler(populationService)
		SubmitmpHandler = handler.NewSubmitMpHandler(submitmpService)
	})

	AfterEach(func() {
		MySQLConnection.Exec("DROP from District")
		MySQLConnection.Exec("DROP from Candidate")
		MySQLConnection.Exec("DROP from Mp")
		MySQLConnection.Exec("DROP from Population")
	})

	Context("Validity API", func() {
		Context("Database have population data", func() {
			When("The CitizenId is matching", func() {
				It("should return success.", func() {
					//Call Api
					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					body := model.SubmitMp{ID: 1234, CitizenID: "1234567891234"}
					jsonValue, _ := json.Marshal(body)
					c.Request = httptest.NewRequest(http.MethodPost, "/mp/submit/", bytes.NewBuffer(jsonValue))
					SubmitmpHandler.SubmitMp(c)

					Expect(res.Result().StatusCode).To(Equal(http.StatusOK))
				})
			})

			When("The CitizenId is not matching", func() {
				It("should return success.", func() {
					//Call Api
					res := httptest.NewRecorder()
					c, _ := gin.CreateTestContext(res)
					body := model.SubmitMp{ID: 1234, CitizenID: "0"}
					jsonValue, _ := json.Marshal(body)
					c.Request = httptest.NewRequest(http.MethodPost, "/mp/submit/", bytes.NewBuffer(jsonValue))
					SubmitmpHandler.SubmitMp(c)

					Expect(res.Result().StatusCode).To(Equal(http.StatusBadRequest))
				})
			})
		})
	})
})
