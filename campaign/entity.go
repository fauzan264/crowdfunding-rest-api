package campaign

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	Id               uuid.UUID
	UserId           uuid.UUID
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CampaignImage struct {
	Id         uuid.UUID
	CampaignId uuid.UUID
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
