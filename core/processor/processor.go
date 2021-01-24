package processor

import (
	"cashapp/core/currency"
	"cashapp/models"
	"errors"

	"cashapp/repository"
	"fmt"
	"log"

	"gorm.io/gorm"
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
	switch fromTrans.Purpose {
	case models.Transfer:
		if err := p.MoveMoneyBetweenWallets(fromTrans); err != nil {
			return fmt.Errorf("money transfer failed. %v", err)
		}
	case models.Transfer:
		if err := p.WithdrawMoneyFromWallet(fromTrans); err != nil {
			return fmt.Errorf("money withdrawal failed. %v", err)
		}
	case models.Deposit:
		if err := p.DepositMoneyIntoWallet(fromTrans); err != nil {
			return fmt.Errorf("money deposit failed. %v", err)
		}
	default:
		log.Println("no handler for purpose. purpose=", fromTrans.Purpose)
	}
	return nil
}

func (p *Processor) MoveMoneyBetweenWallets(fromTrans models.Transaction) error {

	originWallet, err := p.Repo.Wallets.FindPrimaryWallet(fromTrans.From)
	if err != nil {
		return fmt.Errorf("failed to find primary wallet for origin. %v", err)
	}

	destinationWallet, err := p.Repo.Wallets.FindPrimaryWallet(fromTrans.To)
	if err != nil {
		return fmt.Errorf("failed to find primary wallet for destination. %v", err)
	}

	balance, err := p.Repo.TransactionEvents.GetWalletBalance(originWallet.ID)
	if err != nil {
		return fmt.Errorf("failed to load balance. %v", err)
	}

	if balance > fromTrans.Amount {
		return errors.New("insufficient balance")
	}

	toTrans := models.Transaction{
		From:        fromTrans.From,
		To:          fromTrans.To,
		Ref:         fromTrans.Ref,
		Amount:      currency.ConvertCedisToPessewas(fromTrans.Amount),
		Description: fromTrans.Description,
		Direction:   models.Incoming,
		Status:      models.Pending,
		Purpose:     models.Transfer,
	}

	if err := p.Repo.Transactions.Create(&toTrans); err != nil {
		return fmt.Errorf("failed to create destination transaction. %v", err)
	}

	err = p.Repo.Transactions.SQLTransaction(func(tx *gorm.DB) error {
		debit := models.TransactionEvent{
			TransactionID: fromTrans.ID,
			WalletID:      originWallet.ID,
			Amount:        fromTrans.Amount,
			Type:          models.Debit,
		}

		if err := p.Repo.TransactionEvents.Save(tx, &debit); err != nil {
			return err
		}

		credit := models.TransactionEvent{
			TransactionID: toTrans.ID,
			WalletID:      destinationWallet.ID,
			Amount:        toTrans.Amount,
			Type:          models.Credit,
		}

		if err := p.Repo.TransactionEvents.Save(tx, &credit); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("money movement failed. err=%v", err)
	}

	return nil
}

func (p *Processor) DepositMoneyIntoWallet(fromTrans models.Transaction) error {
	return nil
}

func (p *Processor) WithdrawMoneyFromWallet(fromTrans models.Transaction) error {
	return nil
}
