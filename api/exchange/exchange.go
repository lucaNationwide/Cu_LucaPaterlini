package exchange

import (
	"Cu_LucaPaterlini/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const emptyResponse = "empty response"
const layoutISO = "2006-01-02"

// ResponseHandlerExchange implements the response used by GetRateByDate.
type ResponseHandlerExchange struct {
	Val        float64
	Suggestion string
}

// GetRateByDate returns the last value of the currency with euro as base in a specific interval
// and with the advice to buy or sell depending on the average of the days in the range.
func GetRateByDate(ticker, base string, startDate, endDate time.Time) (res ResponseHandlerExchange, err error) {
	// checking supported tickers
	found := false
	for _, b := range config.SupportedTickers {
		if b == base {
			found = true
		}
	}
	if !found {
		err = fmt.Errorf("ticker %s not supported", ticker)
		return
	}
	// collection the response from the third party api
	requestURL := config.FullThirdPartyAPIPath + "history?start_at=" + startDate.Format(layoutISO) + "&end_at=" +
		endDate.Format(layoutISO) + "&symbols=" + ticker + "&base=" + base
	client := &http.Client{Timeout: config.DefaultRequestsTimeout}
	resp, err := client.Get(requestURL)
	if err != nil {
		return
	}
	if resp == nil {
		err = errors.New(emptyResponse)
		return
	}
	// converting the json response in an array with the values and the average
	response := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}
	tmp := response["rates"].(map[string]interface{})

	var n int
	var sum float64
	var tmpDate time.Time
	for key, item := range tmp {
		x := item.(map[string]interface{})[ticker].(float64)
		sum += x
		n++
		if tmpDate, err = time.Parse(layoutISO, key); err != nil {
			return
		}
		if tmpDate.After(startDate) {
			startDate = tmpDate
			res.Val = x
		}
	}

	// giving an advice considering the average of the previous 7 days.
	res.Suggestion = "sell"
	if sum/float64(n) > res.Val {
		res.Suggestion = "buy"
	}
	res.Suggestion += fmt.Sprintf(" %s/%s", base, ticker)
	return
}
