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
	koinim, err := getKoinim()
	paribu, err := getParibu()
	bitfinex, err := getBitfinex()
	koineks, err := getKoineks()
	btcturk, err := getBtcturk()
	if err != nil {
		return display
	}
	bitfinexlast, _ := strconv.ParseFloat(bitfinex.Last_price, 32)
	diff1 := koinim.Last_order - (bitfinexlast * 3.8)
	display.Bitfinex_Koinim = diff1 * 100 / (bitfinexlast * 3.8)
	fmt.Println("KOINIM:", display.Bitfinex_Koinim)

	diff2 := paribu.BTC_TL.Last - (bitfinexlast * 3.8)
	display.Bitfinex_Paribu = diff2 * 100 / (bitfinexlast * 3.8)
	fmt.Println("PARIBU:", display.Bitfinex_Paribu)

	koineksCurrent, _ := strconv.ParseFloat(koineks.BTC.Current, 64)
	diff3 := koineksCurrent - (bitfinexlast * 3.8)
	display.Bitfinex_Koineks = diff3 * 100 / (bitfinexlast * 3.8)
	fmt.Println("KOINEKS:", display.Bitfinex_Koineks)

	diff4 := btcturk[0].Last - (bitfinexlast * 3.8)
	display.Bitfinex_Btcturk = diff4 * 100 / (bitfinexlast * 3.8)
	fmt.Println("BTCTURK:", display.Bitfinex_Btcturk)

	return display
}
func getKoinim() (Koinim, error) {
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
func getParibu() (ParibuTop, error) {
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
func getBitfinex() (Bitfinex, error) {
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
func getUsdTry() (Rate, error) {
	var rate Rate

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
	fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &rate)
	if err != nil {
		fmt.Println("RError3:", err)
		return rate, err
	}
	fmt.Println("RJson:", rate)
	return rate, nil
}

func getKoineks() (KoineksTop, error) {
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

func getBtcturk() ([]Btcturk, error) {
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
	fmt.Println("BTJson:", btcturk)
	return btcturk, nil
}
