package helper

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

func RealTimeCurrencyCNY() (r float64, err error) {
	return RealTimeCurrency("CNY")
}

func RealTimeCurrency(currency string) (r float64, err error) {
	resp, err := http.Get("https://api.coinbase.com/v2/exchange-rates?currency=USD")
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	obj := string(body)
	value := gjson.Get(obj, "data.rates."+strings.ToUpper(currency))
	r = value.Float()
	return
}
