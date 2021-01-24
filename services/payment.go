package services

import (
	"cashapp/core"
	"cashapp/core/currency"

	"cashapp/core/processor"

	"cashapp/models"
	"cashapp/repository"
)

type paymentLayer struct {
	repository repository.Repo
	config     *core.Config
	processor  processor.Processor
}

func newPaymentLayer(r repository.Repo, c *core.Config) *paymentLayer {

	return &paymentLayer{
		repository: r,
		config:     c,
		processor:  processor.New(r),
	}
}

func (p *paymentLayer) SendMoney(req core.CreatePaymentRequest) core.Response {
	fromTrans := models.Transaction{
		From:        req.From,
		To:          req.To,
		Ref:         core.GenerateRef(),
		Amount:      currency.ConvertCedisToPessewas(req.Amount),
		Description: req.Description,
		Direction:   models.Outgoing,
		Status:      models.Pending,
		Purpose:     models.Transfer,
	}

	if err := p.repository.Transactions.Create(&fromTrans); err != nil {
		return core.Error(err, nil)
	}

	if err := p.processor.ProcessTransaction(fromTrans); err != nil {
		return core.Error(err, nil)
	}

	return core.Success(nil, nil)
}

func (p *paymentLayer) WithdrawMoney(req core.CreatePaymentRequest) core.Response {
	return core.Success(nil, nil)
}

func (p *paymentLayer) DepositMoney(req core.CreatePaymentRequest) core.Response {
	return core.Success(nil, nil)
}
