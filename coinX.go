package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"math"
)

func main() {
	requestBundle()
}

func requestBundle() Display {
	var display Display
	btcKoinim, err := getBtcKoinim()
	ltcKoinim, err := getLtcKoinim()
	btcParibu, err := getBtcParibu()
	allKoineks, err := getAllKoineks()
	allBtcturk, err := getAllBtcturk()
	btcBitfinex, err := getBtcBitfinex()
	ltcBitfinex, err := getLtcBitfinex()
	ethBitfinex, err := getEthBitfinex()
	btcBitstamp, err := getBtcBitstamp()
	ltcBitstamp, err := getLtcBitstamp()
	ethBitstamp, err := getEthBitstamp()
	rate, err := getUsdTry()
	if err != nil {
		return display
	}

	currency := rate.Rates.TRY

	btcKoinimTRY := btcKoinim.Last_order
	btcParibuTRY := btcParibu.BTC_TL.Last
	btcKoineksTRY, _ := strconv.ParseFloat(allKoineks.BTC.Current, 64)
	btcBtcturkTRY := allBtcturk[0].Last
	btcBitfinexUSD, _ := strconv.ParseFloat(btcBitfinex.Last_price, 64)
	btcBitstampUSD, _ := strconv.ParseFloat(btcBitstamp.Last, 64)

	btcKoinimUSD := btcKoinimTRY / currency
	btcParibuUSD := btcParibuTRY / currency
	btcKoineksUSD := btcKoineksTRY / currency
	btcBtcturkUSD := btcBtcturkTRY / currency

	btcBitfinexKoinimGap := btcKoinimUSD - btcBitfinexUSD
	btcBitfinexParibuGap := btcParibuUSD - btcBitfinexUSD
	btcBitfinexKoineksGap := btcKoineksUSD - btcBitfinexUSD
	btcBitfinexBtcturkGap := btcBtcturkUSD - btcBitfinexUSD
	btcBitstampKoinimGap := btcKoinimUSD - btcBitstampUSD
	btcBitstampParibuGap := btcParibuUSD - btcBitstampUSD
	btcBitstampKoineksGap := btcKoineksUSD - btcBitstampUSD
	btcBitstampBtcturkGap := btcBtcturkUSD - btcBitstampUSD
	
	btcBitfinexKoinimGapPercentage := btcBitfinexKoinimGap / math.Min(btcBitfinexUSD, btcKoinimUSD) * 100
	btcBitfinexParibuGapPercentage := btcBitfinexParibuGap / math.Min(btcBitfinexUSD, btcParibuUSD) * 100
	btcBitfinexKoineksGapPercentage := btcBitfinexKoineksGap / math.Min(btcBitfinexUSD, btcKoineksUSD) * 100
	btcBitfinexBtcturkGapPercentage := btcBitfinexBtcturkGap / math.Min(btcBitfinexUSD, btcBtcturkUSD) * 100
	btcBitstampKoinimGapPercentage := btcBitstampKoinimGap / math.Min(btcBitstampUSD, btcKoinimUSD) * 100
	btcBitstampParibuGapPercentage := btcBitstampParibuGap / math.Min(btcBitstampUSD, btcParibuUSD) * 100
	btcBitstampKoineksGapPercentage := btcBitstampKoineksGap / math.Min(btcBitstampUSD, btcKoineksUSD) * 100
	btcBitstampBtcturkGapPercentage := btcBitstampBtcturkGap / math.Min(btcBitstampUSD, btcBtcturkUSD) * 100
	
	btcKoinimParibuGapTRY := btcKoinimTRY - btcParibuTRY
	btcKoinimKoineksGapTRY := btcKoinimTRY - btcKoineksTRY
	btcKoinimBtcturkGapTRY := btcKoinimTRY - btcBtcturkTRY
	btcParibuKoineksGapTRY := btcParibuTRY - btcKoineksTRY
	btcParibuBtcturkGapTRY := btcParibuTRY - btcBtcturkTRY	
	btcKoineksBtcturkGapTRY := btcKoineksTRY - btcBtcturkTRY	
	
	btcKoinimParibuGapPercentageTRY := btcKoinimParibuGapTRY / math.Min(btcKoinimTRY, btcParibuTRY) * 100
	btcKoinimKoineksGapPercentageTRY := btcKoinimKoineksGapTRY / math.Min(btcKoinimTRY, btcKoineksTRY) * 100
	btcKoinimBtcturkGapPercentageTRY := btcKoinimBtcturkGapTRY / math.Min(btcKoinimTRY, btcBtcturkTRY) * 100
	btcParibuKoineksGapPercentageTRY := btcParibuKoineksGapTRY / math.Min(btcParibuTRY, btcKoineksTRY) * 100
	btcParibuBtcturkGapPercentageTRY := btcParibuBtcturkGapTRY / math.Min(btcParibuTRY, btcBtcturkTRY) * 100
	btcKoineksBtcturkGapPercentageTRY := btcKoineksBtcturkGapTRY / math.Min(btcKoineksTRY, btcBtcturkTRY) * 100

	ltcKoinimTRY := ltcKoinim.Last_order
	ltcKoineksTRY, _ := strconv.ParseFloat(allKoineks.LTC.Current, 64)
	ltcBitfinexUSD, _ := strconv.ParseFloat(ltcBitfinex.Last_price, 64)
	ltcBitstampUSD, _ := strconv.ParseFloat(ltcBitstamp.Last, 64)

	ltcKoinimUSD := ltcKoinimTRY / currency
	ltcKoineksUSD := ltcKoineksTRY / currency

	ltcBitfinexKoinimGap := ltcKoinimUSD - ltcBitfinexUSD
	ltcBitfinexKoineksGap := ltcKoineksUSD - ltcBitfinexUSD
	ltcBitstampKoinimGap := ltcKoinimUSD - ltcBitstampUSD
	ltcBitstampKoineksGap := ltcKoineksUSD - ltcBitstampUSD

	ltcBitfinexKoinimGapPercentage := ltcBitfinexKoinimGap / math.Min(ltcBitfinexUSD, ltcKoinimUSD) * 100
	ltcBitfinexKoineksGapPercentage := ltcBitfinexKoineksGap / math.Min(ltcBitfinexUSD, ltcKoineksUSD) * 100
	ltcBitstampKoinimGapPercentage := ltcBitstampKoinimGap / math.Min(ltcBitstampUSD, ltcKoinimUSD) * 100
	ltcBitstampKoineksGapPercentage := ltcBitstampKoineksGap / math.Min(ltcBitstampUSD, ltcKoineksUSD) * 100

	ltcKoinimKoineksGapTRY := ltcKoinimTRY - ltcKoineksTRY
	ltcKoinimKoineksGapPercentageTRY := ltcKoinimKoineksGapTRY / math.Min(ltcKoinimTRY, ltcKoineksTRY) * 100

	ethKoineksTRY, _ := strconv.ParseFloat(allKoineks.ETH.Current, 64)
	ethBtcturkTRY := allBtcturk[2].Last
	ethBitfinexUSD, _ := strconv.ParseFloat(ethBitfinex.Last_price, 64)
	ethBitstampUSD, _ := strconv.ParseFloat(ethBitstamp.Last, 64)

	ethKoineksUSD := ethKoineksTRY / currency
	ethBtcturkUSD := ethBtcturkTRY / currency

	ethBitfinexKoineksGap := ethKoineksUSD - ethBitfinexUSD
	ethBitfinexBtcturkGap := ethBtcturkUSD - ethBitfinexUSD
	ethBitstampKoineksGap := ethKoineksUSD - ethBitstampUSD
	ethBitstampBtcturkGap := ethBtcturkUSD - ethBitstampUSD

	ethBitfinexKoineksGapPercentage := ethBitfinexKoineksGap / math.Min(ethBitfinexUSD, ethKoineksUSD) * 100
	ethBitfinexBtcturkGapPercentage := ethBitfinexBtcturkGap / math.Min(ethBitfinexUSD, ethBtcturkUSD) * 100
	ethBitstampKoineksGapPercentage := ethBitstampKoineksGap / math.Min(ethBitstampUSD, ethKoineksUSD) * 100
	ethBitstampBtcturkGapPercentage := ethBitstampBtcturkGap / math.Min(ethBitstampUSD, ethBtcturkUSD) * 100
	
	ethKoineksBtcturkGapTRY := ethKoineksTRY - ethBtcturkTRY
	ethKoineksBtcturkGapPercentageTRY := ethKoineksBtcturkGapTRY / math.Min(ethKoineksTRY, ethBtcturkTRY) * 100
		

	fmt.Println("***\n\t\t\t CURRENCY")
	fmt.Println("USD/TRY:\t\t", strconv.FormatFloat(currency, 'f', 4, 64), "₺")

	fmt.Println("***\n\t\t\t PRICE\t\t USD-EQUIVALENT")
	fmt.Println("BTC (Koinim):\t\t", strconv.FormatFloat(btcKoinimTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(btcKoinimUSD, 'f', 2, 64), "$")
	fmt.Println("BTC (Paribu):\t\t", strconv.FormatFloat(btcParibuTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(btcParibuUSD, 'f', 2, 64), "$")
	fmt.Println("BTC (Koineks):\t\t", strconv.FormatFloat(btcKoineksTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(btcKoineksUSD, 'f', 2, 64), "$")
	fmt.Println("BTC (BTCTurk):\t\t", strconv.FormatFloat(btcBtcturkTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(btcBtcturkUSD, 'f', 2, 64), "$")
	fmt.Println("BTC (Bitfinex):\t\t", strconv.FormatFloat(btcBitfinexUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitfinexUSD, 'f', 2, 64), "$")
	fmt.Println("BTC (Bitstamp):\t\t", strconv.FormatFloat(btcBitstampUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitstampUSD, 'f', 2, 64), "$")

	fmt.Println("---\n\t\t\t GAP\t\t PERCENTAGE")
	fmt.Println("BTC (Bitfinex/Koinim):\t", fmt.Sprintf("%+.2f", btcBitfinexKoinimGap), "$\t", fmt.Sprintf("%+.2f", btcBitfinexKoinimGapPercentage), "%")
	fmt.Println("BTC (Bitfinex/Paribu):\t", fmt.Sprintf("%+.2f", btcBitfinexParibuGap), "$\t", fmt.Sprintf("%+.2f", btcBitfinexParibuGapPercentage), "%")
	fmt.Println("BTC (Bitfinex/Koineks):\t", fmt.Sprintf("%+.2f", btcBitfinexKoineksGap), "$\t", fmt.Sprintf("%+.2f", btcBitfinexKoineksGapPercentage), "%")
	fmt.Println("BTC (Bitfinex/BTCTurk):\t", fmt.Sprintf("%+.2f", btcBitfinexBtcturkGap), "$\t", fmt.Sprintf("%+.2f", btcBitfinexBtcturkGapPercentage), "%")
	fmt.Println("BTC (Bitstamp/Koinim):\t", fmt.Sprintf("%+.2f", btcBitstampKoinimGap), "$\t", fmt.Sprintf("%+.2f", btcBitstampKoinimGapPercentage), "%")
	fmt.Println("BTC (Bitstamp/Paribu):\t", fmt.Sprintf("%+.2f", btcBitstampParibuGap), "$\t", fmt.Sprintf("%+.2f", btcBitstampParibuGapPercentage), "%")
	fmt.Println("BTC (Bitstamp/Koineks):\t", fmt.Sprintf("%+.2f", btcBitstampKoineksGap), "$\t", fmt.Sprintf("%+.2f", btcBitstampKoineksGapPercentage), "%")
	fmt.Println("BTC (Bitstamp/BTCTurk):\t", fmt.Sprintf("%+.2f", btcBitstampBtcturkGap), "$\t", fmt.Sprintf("%+.2f", btcBitstampBtcturkGapPercentage), "%")

	fmt.Println("---")
	fmt.Println("BTC (Koinim/Paribu):\t", fmt.Sprintf("%+.2f", btcKoinimParibuGapTRY), "₺\t", fmt.Sprintf("%+.2f", btcKoinimParibuGapPercentageTRY), "%")
	fmt.Println("BTC (Koinim/Koineks):\t", fmt.Sprintf("%+.2f", btcKoinimKoineksGapTRY), "₺\t", fmt.Sprintf("%+.2f", btcKoinimKoineksGapPercentageTRY), "%")
	fmt.Println("BTC (Koinim/BTCTurk):\t", fmt.Sprintf("%+.2f", btcKoinimBtcturkGapTRY), "₺\t", fmt.Sprintf("%+.2f", btcKoinimBtcturkGapPercentageTRY), "%")
	fmt.Println("BTC (Paribu/Koineks):\t", fmt.Sprintf("%+.2f", btcParibuKoineksGapTRY), "₺\t", fmt.Sprintf("%+.2f", btcParibuKoineksGapPercentageTRY), "%")
	fmt.Println("BTC (Paribu/BTCTurk):\t", fmt.Sprintf("%+.2f", btcParibuBtcturkGapTRY), "₺\t", fmt.Sprintf("%+.2f", btcParibuBtcturkGapPercentageTRY), "%")
	fmt.Println("BTC (Koineks/BTCTurk):\t", fmt.Sprintf("%+.2f", btcKoineksBtcturkGapTRY), "₺\t", fmt.Sprintf("%+.2f", btcKoineksBtcturkGapPercentageTRY), "%")

	fmt.Println("***\n\t\t\t PRICE\t\t USD-EQUIVALENT")
	fmt.Println("LTC (Koinim):\t\t", strconv.FormatFloat(ltcKoinimTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(ltcKoinimUSD, 'f', 2, 64), "$")
	fmt.Println("LTC (Koineks):\t\t", strconv.FormatFloat(ltcKoineksTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(ltcKoineksUSD, 'f', 2, 64), "$")
	fmt.Println("LTC (Bitfinex):\t\t", strconv.FormatFloat(ltcBitfinexUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(ltcBitfinexUSD, 'f', 2, 64), "$")
	fmt.Println("LTC (Bitstamp):\t\t", strconv.FormatFloat(ltcBitstampUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(ltcBitstampUSD, 'f', 2, 64), "$")

	fmt.Println("---\n\t\t\t GAP\t\t PERCENTAGE")
	fmt.Println("LTC (Bitfinex/Koinim):\t", fmt.Sprintf("%+.2f", ltcBitfinexKoinimGap), "$\t", fmt.Sprintf("%+.2f", ltcBitfinexKoinimGapPercentage), "%")
	fmt.Println("LTC (Bitfinex/Koineks):\t", fmt.Sprintf("%+.2f", ltcBitfinexKoineksGap), "$\t", fmt.Sprintf("%+.2f", ltcBitfinexKoineksGapPercentage), "%")
	fmt.Println("LTC (Bitstamp/Koinim):\t", fmt.Sprintf("%+.2f", ltcBitstampKoinimGap), "$\t", fmt.Sprintf("%+.2f", ltcBitstampKoinimGapPercentage), "%")
	fmt.Println("LTC (Bitstamp/Koineks):\t", fmt.Sprintf("%+.2f", ltcBitstampKoineksGap), "$\t", fmt.Sprintf("%+.2f", ltcBitstampKoineksGapPercentage), "%")

	fmt.Println("---")
	fmt.Println("LTC (Koinim/Koineks):\t", fmt.Sprintf("%+.2f", ltcKoinimKoineksGapTRY), "₺\t", fmt.Sprintf("%+.2f", ltcKoinimKoineksGapPercentageTRY), "%")

	fmt.Println("***\n\t\t\t PRICE\t\t USD-EQUIVALENT")
	fmt.Println("ETH (Koineks):\t\t", strconv.FormatFloat(ethKoineksTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(ethKoineksUSD, 'f', 2, 64), "$")
	fmt.Println("ETH (BTCTurk):\t\t", strconv.FormatFloat(ethBtcturkTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(ethBtcturkUSD, 'f', 2, 64), "$")
	fmt.Println("ETH (Bitfinex):\t\t", strconv.FormatFloat(ethBitfinexUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(ethBitfinexUSD, 'f', 2, 64), "$")
	fmt.Println("ETH (Bitstamp):\t\t", strconv.FormatFloat(ethBitstampUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(ethBitstampUSD, 'f', 2, 64), "$")

	fmt.Println("---\n\t\t\t GAP\t\t PERCENTAGE")
	fmt.Println("ETH (Bitfinex/Koineks):\t", fmt.Sprintf("%+.2f", ethBitfinexKoineksGap), "$\t", fmt.Sprintf("%+.2f", ethBitfinexKoineksGapPercentage), "%")
	fmt.Println("ETH (Bitfinex/BTCTurk):\t", fmt.Sprintf("%+.2f", ethBitfinexBtcturkGap), "$\t", fmt.Sprintf("%+.2f", ethBitfinexBtcturkGapPercentage), "%")
	fmt.Println("ETH (Bitstamp/Koineks):\t", fmt.Sprintf("%+.2f", ethBitstampKoineksGap), "$\t", fmt.Sprintf("%+.2f", ethBitstampKoineksGapPercentage), "%")
	fmt.Println("ETH (Bitstamp/BTCTurk):\t", fmt.Sprintf("%+.2f", ethBitstampBtcturkGap), "$\t", fmt.Sprintf("%+.2f", ethBitstampBtcturkGapPercentage), "%")
	
	fmt.Println("---")
	fmt.Println("ETH (Koineks/BTCTurk):\t", fmt.Sprintf("%+.2f", ethKoineksBtcturkGapTRY), "₺\t", fmt.Sprintf("%+.2f", ethKoineksBtcturkGapPercentageTRY), "%")
	fmt.Println("***")

	return display
}

func getBtcKoinim() (Koinim, error) {
	var koinim Koinim

	resp, err := http.Get("https://koinim.com/ticker")
	if err != nil {
		fmt.Println("KError1:", err)
		return koinim, err
	}
	//fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("KError2:", err)
		return koinim, err
	}
	//fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &koinim)
	if err != nil {
		fmt.Println("KError3:", err)
		return koinim, err
	}
	//fmt.Println("KJson:", koinim)
	return koinim, nil
}

func getLtcKoinim() (Koinim, error) {
	var koinim Koinim

	resp, err := http.Get("https://koinim.com/ticker/ltc/")
	if err != nil {
		fmt.Println("KError1:", err)
		return koinim, err
	}
	//fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("KError2:", err)
		return koinim, err
	}
	//fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &koinim)
	if err != nil {
		fmt.Println("KError3:", err)
		return koinim, err
	}
	//fmt.Println("KJson:", koinim)
	return koinim, nil
}

func getBtcParibu() (ParibuTop, error) {
	var paribu ParibuTop

	resp, err := http.Get("https://www.paribu.com/ticker")
	if err != nil {
		fmt.Println("PError1:", err)
		return paribu, err
	}
	//fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("PError2:", err)
		return paribu, err
	}
	//fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &paribu)
	if err != nil {
		fmt.Println("PError3:", err)
		return paribu, err
	}
	//fmt.Println("PJson:", paribu)
	return paribu, nil
}

func getBtcBitfinex() (Bitfinex, error) {
	var bitfinex Bitfinex

	resp, err := http.Get("https://api.bitfinex.com/v1/pubticker/btcusd")
	if err != nil {
		fmt.Println("BError1:", err)
		return bitfinex, err
	}
	//fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("BError2:", err)
		return bitfinex, err
	}
	//fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &bitfinex)
	if err != nil {
		fmt.Println("BError3:", err)
		return bitfinex, err
	}
	//fmt.Println("BJson:", bitfinex)
	return bitfinex, nil
}

func getLtcBitfinex() (Bitfinex, error) {
	var bitfinex Bitfinex

	resp, err := http.Get("https://api.bitfinex.com/v1/pubticker/ltcusd")
	if err != nil {
		fmt.Println("BError1:", err)
		return bitfinex, err
	}
	//fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("BError2:", err)
		return bitfinex, err
	}
	//fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &bitfinex)
	if err != nil {
		fmt.Println("BError3:", err)
		return bitfinex, err
	}
	//fmt.Println("BJson:", bitfinex)
	return bitfinex, nil
}

func getEthBitfinex() (Bitfinex, error) {
	var bitfinex Bitfinex

	resp, err := http.Get("https://api.bitfinex.com/v1/pubticker/ethusd")
	if err != nil {
		fmt.Println("BError1:", err)
		return bitfinex, err
	}
	//fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("BError2:", err)
		return bitfinex, err
	}
	//fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &bitfinex)
	if err != nil {
		fmt.Println("BError3:", err)
		return bitfinex, err
	}
	//fmt.Println("BJson:", bitfinex)
	return bitfinex, nil
}

func getAllKoineks() (KoineksTop, error) {
	var koineks KoineksTop

	resp, err := http.Get("https://koineks.com/ticker")
	if err != nil {
		fmt.Println("XError1:", err)
		return koineks, err
	}
	//fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("XError2:", err)
		return koineks, err
	}
	//fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &koineks)
	if err != nil {
		fmt.Println("XError3:", err)
		return koineks, err
	}
	//fmt.Println("XJson:", koineks)
	return koineks, nil
}

func getAllBtcturk() ([]Btcturk, error) {
	//var btcturk BtcturkTop
	btcturk := make([]Btcturk, 0)
	resp, err := http.Get("https://www.btcturk.com/api/ticker")
	if err != nil {
		fmt.Println("BTError1:", err)
		return btcturk, err
	}
	//fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("BTError2:", err)
		return btcturk, err
	}
	//fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &btcturk)
	if err != nil {
		fmt.Println("BTError3:", err)
		return btcturk, err
	}
	//fmt.Println("BTJson:", btcturk)
	return btcturk, nil
}

func getUsdTry() (UsdTry, error) {
	var rate UsdTry

	resp, err := http.Get("https://api.fixer.io/latest?base=USD&symbols=TRY")
	if err != nil {
		fmt.Println("RError1:", err)
		return rate, err
	}
	//fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("RError2:", err)
		return rate, err
	}
	//fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &rate)
	if err != nil {
		fmt.Println("RError3:", err)
		return rate, err
	}
	//fmt.Println("RJson:", rate)
	return rate, nil
}

func getBtcBitstamp() (Bitstamp, error) {
	var bitstamp Bitstamp

	resp, err := http.Get("https://www.bitstamp.net/api/v2/ticker/btcusd")
	if err != nil {
		fmt.Println("BError1:", err)
		return bitstamp, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("BError2:", err)
		return bitstamp, err
	}
	err = json.Unmarshal(bodyBytes, &bitstamp)
	if err != nil {
		fmt.Println("BError3:", err)
		return bitstamp, err
	}
	return bitstamp, nil
}

func getLtcBitstamp() (Bitstamp, error) {
	var bitstamp Bitstamp

	resp, err := http.Get("https://www.bitstamp.net/api/v2/ticker/ltcusd")
	if err != nil {
		fmt.Println("BError1:", err)
		return bitstamp, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("BError2:", err)
		return bitstamp, err
	}
	err = json.Unmarshal(bodyBytes, &bitstamp)
	if err != nil {
		fmt.Println("BError3:", err)
		return bitstamp, err
	}
	return bitstamp, nil
}

func getEthBitstamp() (Bitstamp, error) {
	var bitstamp Bitstamp

	resp, err := http.Get("https://www.bitstamp.net/api/v2/ticker/ethusd")
	if err != nil {
		fmt.Println("BError1:", err)
		return bitstamp, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("BError2:", err)
		return bitstamp, err
	}
	err = json.Unmarshal(bodyBytes, &bitstamp)
	if err != nil {
		fmt.Println("BError3:", err)
		return bitstamp, err
	}
	return bitstamp, nil
}
