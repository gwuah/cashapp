package processor

import (
	"cashapp/models"

	"gorm.io/gorm"
)

func (p *Processor) MoveMoneyBetweenWallets(fromTrans, toTrans models.Transaction) error {

	err := p.Repo.Transactions.SQLTransaction(func(tx *gorm.DB) error {
		debit := models.TransactionEvent{
			TransactionID: fromTrans.ID,
			Amount:        fromTrans.Amount,
			Type:          models.Debit,
		}

		if err := p.Repo.TransactionEvents.Save(tx, &debit); err != nil {
			return err
		}

		credit := models.TransactionEvent{
			TransactionID: toTrans.ID,
			Amount:        toTrans.Amount,
			Type:          models.Credit,
		}

		if err := p.Repo.TransactionEvents.Save(tx, &credit); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil
	}

	return nil
}
