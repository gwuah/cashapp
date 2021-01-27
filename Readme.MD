# Cashapp

A working version of a simple peer-to-peer payment application.

# Setup

- create your own `.env` file based on `.sample.env`
- and then initialise using either `go run main.go` or `go build main.go && ./main`
- the port you set in your env file is the port on which it will run

# Features

- Ledger
- Users
- Wallets
- Transactions
- Balance

## Notes

This is not production ready.
This system works assumes the world is perfect.
Also, no physical money is moved in this sytem. you'll need to integrate with a payment processor(expresspay, paystack, flutterwave)
When I write add an integration with a payment processor, expect lots of things in this to change. As it stands, there's no architecture
to retry transactions, no validations, etc etc
