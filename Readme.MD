# Cashapp

A working version of a simple peer-to-peer payment application.

# Setup

- create your own `.env` file based on `.sample.env`
- and then initialise using either `go run main.go` or `go build main.go && ./main`
- the port you set in your env file is the port on which it will run

# Features

- Ledger
- Accounts
- Wallets
- Transactions
- Balance

## Notes

No physical money is moved in this sysytem. if you want to move money, you need to integrate with a payment processor(expresspay, paystack, flutterwave)
This is just a core part pre-built for you to re-use or learn from
