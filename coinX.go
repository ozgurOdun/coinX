package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

	btcBitfinexKoinimGapPercentage := btcBitfinexKoinimGap / btcKoineksUSD * 100
	btcBitfinexParibuGapPercentage := btcBitfinexParibuGap / btcKoineksUSD * 100
	btcBitfinexKoineksGapPercentage := btcBitfinexKoineksGap / btcKoineksUSD * 100
	btcBitfinexBtcturkGapPercentage := btcBitfinexBtcturkGap / btcKoineksUSD * 100
	btcBitstampKoinimGapPercentage := btcBitstampKoinimGap / btcKoineksUSD * 100
	btcBitstampParibuGapPercentage := btcBitstampParibuGap / btcKoineksUSD * 100
	btcBitstampKoineksGapPercentage := btcBitstampKoineksGap / btcKoineksUSD * 100
	btcBitstampBtcturkGapPercentage := btcBitstampBtcturkGap / btcKoineksUSD * 100

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

	ltcBitfinexKoinimGapPercentage := ltcBitfinexKoinimGap / ltcKoineksUSD * 100
	ltcBitfinexKoineksGapPercentage := ltcBitfinexKoineksGap / ltcKoineksUSD * 100
	ltcBitstampKoinimGapPercentage := ltcBitstampKoinimGap / ltcKoineksUSD * 100
	ltcBitstampKoineksGapPercentage := ltcBitstampKoineksGap / ltcKoineksUSD * 100

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

	ethBitfinexKoineksGapPercentage := ethBitfinexKoineksGap / ethBtcturkUSD * 100
	ethBitfinexBtcturkGapPercentage := ethBitfinexBtcturkGap / ethBtcturkUSD * 100
	ethBitstampKoineksGapPercentage := ethBitstampKoineksGap / ethBtcturkUSD * 100
	ethBitstampBtcturkGapPercentage := ethBitstampBtcturkGap / ethBtcturkUSD * 100

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
	fmt.Println("BTC (Bitfinex/Koinim):\t", strconv.FormatFloat(btcBitfinexKoinimGap, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitfinexKoinimGapPercentage, 'f', 2, 64), "%")
	fmt.Println("BTC (Bitfinex/Paribu):\t", strconv.FormatFloat(btcBitfinexParibuGap, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitfinexParibuGapPercentage, 'f', 2, 64), "%")
	fmt.Println("BTC (Bitfinex/Koineks):\t", strconv.FormatFloat(btcBitfinexKoineksGap, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitfinexKoineksGapPercentage, 'f', 2, 64), "%")
	fmt.Println("BTC (Bitfinex/BTCTurk):\t", strconv.FormatFloat(btcBitfinexBtcturkGap, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitfinexBtcturkGapPercentage, 'f', 2, 64), "%")
	fmt.Println("BTC (Bitstamp/Koinim):\t", strconv.FormatFloat(btcBitstampKoinimGap, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitstampKoinimGapPercentage, 'f', 2, 64), "%")
	fmt.Println("BTC (Bitstamp/Paribu):\t", strconv.FormatFloat(btcBitstampParibuGap, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitstampParibuGapPercentage, 'f', 2, 64), "%")
	fmt.Println("BTC (Bitstamp/Koineks):\t", strconv.FormatFloat(btcBitstampKoineksGap, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitstampKoineksGapPercentage, 'f', 2, 64), "%")
	fmt.Println("BTC (Bitstamp/BTCTurk):\t", strconv.FormatFloat(btcBitstampBtcturkGap, 'f', 2, 64), "$\t", strconv.FormatFloat(btcBitstampBtcturkGapPercentage, 'f', 2, 64), "%")

	fmt.Println("***\n\t\t\t PRICE\t\t USD-EQUIVALENT")
	fmt.Println("LTC (Koinim):\t\t", strconv.FormatFloat(ltcKoinimTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(ltcKoinimUSD, 'f', 2, 64), "$")
	fmt.Println("LTC (Koineks):\t\t", strconv.FormatFloat(ltcKoineksTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(ltcKoineksUSD, 'f', 2, 64), "$")
	fmt.Println("LTC (Bitfinex):\t\t", strconv.FormatFloat(ltcBitfinexUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(ltcBitfinexUSD, 'f', 2, 64), "$")
	fmt.Println("LTC (Bitstamp):\t\t", strconv.FormatFloat(ltcBitstampUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(ltcBitstampUSD, 'f', 2, 64), "$")

	fmt.Println("---\n\t\t\t GAP\t\t PERCENTAGE")
	fmt.Println("LTC (Bitfinex/Koinim):\t", strconv.FormatFloat(ltcBitfinexKoinimGap, 'f', 2, 64), "$\t", strconv.FormatFloat(ltcBitfinexKoinimGapPercentage, 'f', 2, 64), "%")
	fmt.Println("LTC (Bitfinex/Koineks):\t", strconv.FormatFloat(ltcBitfinexKoineksGap, 'f', 2, 64), "$\t", strconv.FormatFloat(ltcBitfinexKoineksGapPercentage, 'f', 2, 64), "%")
	fmt.Println("LTC (Bitstamp/Koinim):\t", strconv.FormatFloat(ltcBitstampKoinimGap, 'f', 2, 64), "$\t", strconv.FormatFloat(ltcBitstampKoinimGapPercentage, 'f', 2, 64), "%")
	fmt.Println("LTC (Bitstamp/Koineks):\t", strconv.FormatFloat(ltcBitstampKoineksGap, 'f', 2, 64), "$\t", strconv.FormatFloat(ltcBitstampKoineksGapPercentage, 'f', 2, 64), "%")

	fmt.Println("***\n\t\t\t PRICE\t\t USD-EQUIVALENT")
	fmt.Println("ETH (Koineks):\t\t", strconv.FormatFloat(ethKoineksTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(ethKoineksUSD, 'f', 2, 64), "$")
	fmt.Println("ETH (BTCTurk):\t\t", strconv.FormatFloat(ethBtcturkTRY, 'f', 2, 64), "₺\t", strconv.FormatFloat(ethBtcturkUSD, 'f', 2, 64), "$")
	fmt.Println("ETH (Bitfinex):\t\t", strconv.FormatFloat(ethBitfinexUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(ethBitfinexUSD, 'f', 2, 64), "$")
	fmt.Println("ETH (Bitstamp):\t\t", strconv.FormatFloat(ethBitstampUSD, 'f', 2, 64), "$\t", strconv.FormatFloat(ethBitstampUSD, 'f', 2, 64), "$")

	fmt.Println("---\n\t\t\t GAP\t\t PERCENTAGE")
	fmt.Println("ETH (Bitfinex/Koineks):\t", strconv.FormatFloat(ethBitfinexKoineksGap, 'f', 2, 64), "$\t", strconv.FormatFloat(ethBitfinexKoineksGapPercentage, 'f', 2, 64), "%")
	fmt.Println("ETH (Bitfinex/BTCTurk):\t", strconv.FormatFloat(ethBitfinexBtcturkGap, 'f', 2, 64), "$\t", strconv.FormatFloat(ethBitfinexBtcturkGapPercentage, 'f', 2, 64), "%")
	fmt.Println("ETH (Bitstamp/Koineks):\t", strconv.FormatFloat(ethBitstampKoineksGap, 'f', 2, 64), "$\t", strconv.FormatFloat(ethBitstampKoineksGapPercentage, 'f', 2, 64), "%")
	fmt.Println("ETH (Bitstamp/BTCTurk):\t", strconv.FormatFloat(ethBitstampBtcturkGap, 'f', 2, 64), "$\t", strconv.FormatFloat(ethBitstampBtcturkGapPercentage, 'f', 2, 64), "%")
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
