package service_test

import (
	"github.com/WorkWorkWork-Team/gov-ec-api/model"
	"github.com/WorkWorkWork-Team/gov-ec-api/service"
	"github.com/WorkWorkWork-Team/gov-ec-api/test/mock_repository"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Population", func() {
	var ctrl *gomock.Controller
	var mockPopulationRepository *mock_repository.MockPopulationRepository
	var populationService service.PopulationService
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockPopulationRepository = mock_repository.NewMockPopulationRepository(ctrl)
		populationService = service.NewPopulationService(mockPopulationRepository)
	})
	Describe("Get pop stat", func() {
		BeforeEach(func() {
			dis := []model.District{}
			dis_obj := model.District{
				DistrictID: 1,
				ProvinceID: 2,
				Name:       "hee",
			}
			var n int64 = 97
			dis = append(dis, dis_obj)
			mockPopulationRepository.EXPECT().
				QueryAllDistrict().
				Return(dis, nil)
			mockPopulationRepository.EXPECT().
				QueryPeopleRightToVote(gomock.Any()).
				Return(n, nil)
			mockPopulationRepository.EXPECT().
				QueryPeopleCommitedTheVote(gomock.Any()).
				Return(n, nil)
			mockPopulationRepository.EXPECT().
				QueryTotalPopulation(gomock.Any()).
				Return(n, nil)
		})
		It("Should not return error", func() {
			res := []model.PopulationResponseItem{}
			res = append(res, model.PopulationResponseItem{
				LocationID:            1,
				Location:              "hee",
				PeopleWithRightToVote: 97,
				PeopleCommitTheVote:   97,
				TotalPeople:           97,
			})
			Expect(populationService.GetPopulationStatistics()).Should(Equal(res))
		})
	})

	Describe("Get GetAllCandidateInfo", func() {
		BeforeEach(func() {
			out := []model.PopulationDatabaseRow{}
			mockPopulationRepository.EXPECT().
				QueryAllCandidate().
				Return(out, nil)

		})
		It("Should not return error", func() {
			res := []model.PopulationDatabaseRow{}
			Expect(populationService.GetAllCandidateInfo()).Should(Equal(res))
		})
	})
})
