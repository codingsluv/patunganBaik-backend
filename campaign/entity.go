package campaign

import (
	"time"

	"github.com/codingsluv/crowdfounding/user"
)

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmmount      int
	CurrentAmmount   int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImages
	User             user.User
}

type CampaignImages struct {
	ID         int
	CampaignID int
	Filename   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
