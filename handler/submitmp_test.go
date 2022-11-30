package handler_test

import (
	"github.com/WorkWorkWork-Team/gov-ec-api/handler"
	"github.com/WorkWorkWork-Team/gov-ec-api/repository"
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
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

	Context("Validity API", func() {
		It("should pass", func() {
			Expect(1).To(Equal(1))
		})
	})
})
