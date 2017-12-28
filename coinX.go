package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	koinim, err := getKoinim()
	if err != nil {
		return
	}
	fmt.Println("KJson2:", koinim)
	paribu, err := getParibu()
	if err != nil {
		return
	}
	fmt.Println("PJson2:", paribu)
	bitfinex, err := getBitfinex()
	if err != nil {
		return
	}
	fmt.Println("BJson2:", bitfinex)
	bitfinexlast, _ := strconv.ParseFloat(bitfinex.Last_price, 32)
	diff1 := koinim.Last_order - bitfinexlast
	fmt.Println("Diff:", diff1)

}

func getKoinim() (Koinim, error) {
	var koinim Koinim

	resp, err := http.Get("https://koinim.com/ticker")
	if err != nil {
		fmt.Println("KError1:", err)
		return koinim, err
	}
	fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("KError2:", err)
		return koinim, err
	}
	fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &koinim)
	if err != nil {
		fmt.Println("KError3:", err)
		return koinim, err
	}
	fmt.Println("KJson:", koinim)
	return koinim, nil
}
func getParibu() (ParibuTop, error) {
	var paribu ParibuTop

	resp, err := http.Get("https://www.paribu.com/ticker")
	if err != nil {
		fmt.Println("PError1:", err)
		return paribu, err
	}
	fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("PError2:", err)
		return paribu, err
	}
	fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &paribu)
	if err != nil {
		fmt.Println("PError3:", err)
		return paribu, err
	}
	fmt.Println("PJson:", paribu)
	return paribu, nil
}
func getBitfinex() (Bitfinex, error) {
	var bitfinex Bitfinex

	resp, err := http.Get("https://api.bitfinex.com/v1/pubticker/btcusd")
	if err != nil {
		fmt.Println("BError1:", err)
		return bitfinex, err
	}
	fmt.Println("Resp:", resp)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("BError2:", err)
		return bitfinex, err
	}
	fmt.Println("Body:", string(bodyBytes))
	err = json.Unmarshal(bodyBytes, &bitfinex)
	if err != nil {
		fmt.Println("BError3:", err)
		return bitfinex, err
	}
	fmt.Println("BJson:", bitfinex)
	return bitfinex, nil
}
