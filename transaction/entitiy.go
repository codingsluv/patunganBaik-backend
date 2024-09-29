package transaction

import "time"

type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
