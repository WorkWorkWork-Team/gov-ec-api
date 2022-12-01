package service_test

import (
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/WorkWorkWork-Team/gov-ec-api/test/mock_repository"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Population", Label("unit"), func() {
	var ctrl *gomock.Controller
	var mockSubmitmpRepository *mock_repository.MockSubmitMpRepository
	var mockPopulationRepository *mock_repository.MockPopulationRepository
	var submitmpService service.SubmitmpService

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockSubmitmpRepository = mock_repository.NewMockSubmitMpRepository(ctrl)
		mockPopulationRepository = mock_repository.NewMockPopulationRepository(ctrl)
		submitmpService = service.NewSubmitmpService(mockSubmitmpRepository, mockPopulationRepository)
	})

	Describe("Submit MP", func() {
		Context("When person does exists", func() {
			BeforeEach(func() {
				mockPopulationRepository.EXPECT().CheckIfPeopleExists("1").Return(true)
				mockSubmitmpRepository.EXPECT().SubmitMpToDB("1").Return(nil)
			})

			It("Should not return ErrCitizenIDNotFound", func() {
				err := submitmpService.SubmitMp("1")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When the person doesn't exists", func() {
			BeforeEach(func() {
				mockPopulationRepository.EXPECT().CheckIfPeopleExists("2").Return(false)
			})

			It("Should return error", func() {
				err := submitmpService.SubmitMp("2")
				Expect(err).Should(Equal(service.ErrCitizenIDNotFound))
			})
		})
	})
})
