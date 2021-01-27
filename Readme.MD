# Cashapp

A working version of a simple peer-to-peer payment application. This was my first iteration.
Lots of things changed in the v2 to make it more reliable. But this is more than enough to demonstrate the concept. 😏

# Setup

- create your own `.env` file based on `.sample.env`
- and then initialise using either `go run main.go` or `go build main.go && ./main`
- the port you set in your env file is the port on which it will run

# Features

- Users
- Wallets
- Transactions
- Balance

# Inspiration

- https://github.com/moov-io/accounts

## Notes

This is the initial version I built. It's not production ready. This system works assumes the world is perfect. <br/>
Also, no physical money is moved in this sytem. you'll need to integrate with a payment processor(expresspay, paystack, flutterwave)
The version 2 has alot of changes (retries, a queue, proper validations, notifications, ledger etc) but I don't think i'll release that.
