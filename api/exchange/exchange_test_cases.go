package exchange

import (
	"fmt"
	"time"
)

func testHelperDate(s string) time.Time {
	tmp, _ := time.Parse(layoutISO, s)
	return tmp
}

var testCasesGetRateByDate = []struct {
	inputTicker      string
	inputBase        string
	inputStartDate   time.Time
	inputEndDate     time.Time
	expectedResponse ResponseHandlerExchange
	expectedError    error
	description      string
}{
	{
		inputTicker:    "EUR",
		inputBase:      "USD",
		inputStartDate: testHelperDate("2018-12-10"),
		inputEndDate:   testHelperDate("2018-12-20"),
		expectedResponse: ResponseHandlerExchange{
			Val:        0.8732861759,
			Suggestion: "buy USD/EUR",
		},
		expectedError: nil,
		description:   "valid usd rate request",
	},
	{
		inputTicker:    "EUR",
		inputBase:      "GBP",
		inputStartDate: testHelperDate("2018-12-10"),
		inputEndDate:   testHelperDate("2018-12-20"),
		expectedResponse: ResponseHandlerExchange{
			Val:        1.1069906459,
			Suggestion: "buy GBP/EUR",
		},
		expectedError: nil,
		description:   "valid gbp rate request",
	},
	{
		inputTicker:      "EUR",
		inputBase:        "EUR",
		inputStartDate:   testHelperDate("2020-12-10"),
		inputEndDate:     testHelperDate("2018-12-20"),
		expectedResponse: ResponseHandlerExchange{},
		expectedError:    fmt.Errorf("ticker EUR not supported"),
		description:      "unsupported ticker",
	},
}
