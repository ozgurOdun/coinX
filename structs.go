package main

type Koinim struct {
	Sell        float64
	High        float64
	Buy         float64
	Change_rate float64
	Bid         float64
	Wavg        float64
	Last_order  float64
	Volume      float64
	Low         float64
	Ask         float64
	Avg         float64
}

type ParibuTop struct {
	BTC_TL Paribu
}
type Paribu struct {
	Last          float64
	LowestAsk     float64
	HighestBid    float64
	PercentChange float64
	Volume        float64
	High24hr      float64
	Low24hr       float64
}

type Bitfinex struct {
	Mid        string
	Bid        string
	Ask        string
	Last_price string
	Low        string
	High       string
	Volume     string
	Timestamp  string
}

type Display struct {
	Bitfinex_Koinim float64
	Bitfinex_Paribu float64
}
type Rate struct {
	TRY float64
}
type UsdTry struct {
	Base  string
	Date  string
	Rates Rate
}
