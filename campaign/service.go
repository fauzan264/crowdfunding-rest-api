package campaign

import "github.com/google/uuid"

type Service interface {
	FindCampaigns(userId string) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaigns(userId string) ([]Campaign, error) {
	if userId != "" {
		campaigns, err := s.repository.FindByUserId(uuid.MustParse(userId))

		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
