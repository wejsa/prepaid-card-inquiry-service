package ias

type BalanceInfo struct {
	Par     string `json:"par"`
	Balance int64  `json:"balance,omitempty"`
}

func (b *BalanceInfo) GetPar() string {
	return b.Par
}

func (b *BalanceInfo) SetPar(par string) {
	b.Par = par
}

func (b *BalanceInfo) GetBalance() int64 {
	return b.Balance
}

func (b *BalanceInfo) SetBalance(balance int64) {
	b.Balance = balance
}
