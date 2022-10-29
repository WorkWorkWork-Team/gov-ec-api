package service

import "github.com/WorkWorkWork-Team/gov-ec-api/repository"

type populationService struct {
	populationRepository repository.PopulationRepository
}

type PopulationService interface {
	GetTotalPopulation() bool
}

func NewPopulationService(populationRepository repository.PopulationRepository) PopulationService {
	return &populationService{
		populationRepository: populationRepository,
	}
}

func (p *populationService) GetTotalPopulation() bool {
	p.populationRepository.GetTotalNumPopulation()
	return false
}
