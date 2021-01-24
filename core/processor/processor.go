package processor

import (
	"cashapp/core/currency"

	"cashapp/models"
	"cashapp/repository"
	"fmt"
)

type Processor struct {
	Repo repository.Repo
}

func New(r repository.Repo) Processor {
	return Processor{
		Repo: r,
	}
}

func (p *Processor) ProcessTransaction(fromTrans models.Transaction) error {

	primaryWallet, err := p.Repo.Wallets.FindPrimaryWallet(fromTrans.From)
	if err != nil {
		return fmt.Errorf("failed to find primary balance. %v", err)
	}

	balance, err := p.Repo.TransactionEvents.GetWalletBalance(primaryWallet.ID)
	if err != nil {
		return fmt.Errorf("failed to load balance. %v", err)
	}

	if balance > fromTrans.Amount {
		return fmt.Errorf("insufficient balance")
	}

	toTrans := models.Transaction{
		From:        fromTrans.From,
		To:          fromTrans.To,
		Amount:      currency.ConvertCedisToPessewas(fromTrans.Amount),
		Description: fromTrans.Description,
		Direction:   models.Incoming,
		Status:      models.Pending,
	}

	if err := p.Repo.Transactions.Create(&toTrans); err != nil {
		return fmt.Errorf("failed to create destination transaction. %v", err)
	}

	if err := p.MoveMoneyBetweenWallets(fromTrans, toTrans); err != nil {
		return fmt.Errorf("money movement failed. %v", err)
	}

	return nil
}
