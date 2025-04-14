package sum

type Account struct {
	Name    string
	Balance float64
}

type Transaction struct {
	From Account
	To   Account
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from, To: to, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func applyTransaction(a Account, t Transaction) Account {
	if t.From.Name == a.Name {
		a.Balance -= t.Sum
	}
	if t.To.Name == a.Name {
		a.Balance += t.Sum
	}
	return a
}
