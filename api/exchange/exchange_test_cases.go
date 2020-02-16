package exchange

import (
	"fmt"
)

var testCasesGetRateByDate = []struct {
	inputTicker      string
	inputStartDate   string
	inputEndDate     string
	expectedResponse ResponseHandlerExchange
	expectedError    error
	description      string
}{
	{
		inputTicker:    "USD",
		inputStartDate: "2018-12-10",
		inputEndDate:   "2018-12-20",
		expectedResponse: ResponseHandlerExchange{
			Val:        1.1451,
			Suggestion: "sell",
		},
		expectedError: nil,
		description:   "valid usd rate request",
	},
	{
		inputTicker:    "GBP",
		inputStartDate: "2018-12-10",
		inputEndDate:   "2018-12-20",
		expectedResponse: ResponseHandlerExchange{
			Val:        0.90335,
			Suggestion: "sell",
		},
		expectedError: nil,
		description:   "valid gbp rate request",
	},
	{
		inputTicker:      "EUR",
		inputStartDate:   "2020-12-10",
		inputEndDate:     "2018-12-20",
		expectedResponse: ResponseHandlerExchange{},
		expectedError:    fmt.Errorf("ticker EUR not supported"),
		description:      "unsupported ticker",
	},
}
