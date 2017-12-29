package main


type Display struct {
	Bitfinex_Koinim  float64
	Bitfinex_Paribu  float64
	Bitfinex_Koineks float64
	Bitfinex_Btcturk float64
}

type Rate struct {
	TRY float64
}

type UsdTry struct {
	Base  string
	Date  string
	Rates Rate
}

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

type KoineksTop struct {
	BTC  Koineks
	ETH  Koineks
	LTC  Koineks
	DASH Koineks
	DOGE Koineks
}

type Koineks struct {
	Short_code        string
	Name              string
	Currency          string
	Current           string
	Change_amount     string
	Change_percentage string
	High              string
	Low               string
	Volume            float64
	Ask               string
	Bid               string
	Timestamp         string
}

type Btcturk struct {
	Pair      string
	High      float64
	Last      float64
	Timestamp float64
	Bid       float64
	Volume    float64
	Low       float64
	Ask       float64
	Open      float64
	Average   float64
}
