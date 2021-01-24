package currency

func ConvertCedisToPessewas(amount int64) int64 {
	return (amount * 100)
}

func ConvertPessewasToCedis(amount int64) int64 {
	return (amount / 100)
}
