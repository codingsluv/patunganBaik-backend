package transaction

import (
	"time"

	"github.com/codingsluv/crowdfounding/user"
)

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	User       user.User
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
