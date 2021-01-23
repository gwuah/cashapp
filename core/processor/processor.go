package processor

import (
	"cashapp/models"
	"cashapp/repository"
)

type Processor struct {
	Repo repository.Repo
}

func New(r repository.Repo) Processor {
	return Processor{
		Repo: r,
	}
}

func (p *Processor) ProcessTransaction(source models.Transaction) error {
	// check balance of origin to see if they can afford the transaction
	// call appropriate settlement formula to handle transaction

	// toTrans := models.Transaction{
	// 	From:        req.From,
	// 	To:          req.To,
	// 	Amount:      core.ConvertCedisToPessewas(req.Amount),
	// 	Description: req.Description,
	// 	Direction:   models.Incoming,
	// 	Status:      models.Pending,
	// }

	return nil
}
