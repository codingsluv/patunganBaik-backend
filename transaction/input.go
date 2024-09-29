package transaction

import "github.com/codingsluv/crowdfounding/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
