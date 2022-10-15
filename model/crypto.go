package model

type RecentTrades struct {
	SymbolName string  `json:"symbolName"`
	Quantity   float32 `json:"quantity"`
	Time       uint64  `json:"tradeTime"`
}

func NewRecentTrade(symbol string, quantity float32, time uint64) *RecentTrades {
	return &RecentTrades{
		SymbolName: symbol,
		Quantity:   quantity,
		Time:       time,
	}
}
