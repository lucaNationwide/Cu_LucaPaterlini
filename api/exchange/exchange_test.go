package exchange

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRateByDate(t *testing.T) {
	for _, tc := range testCasesGetRateByDate {
		got, err := GetRateByDate(tc.inputTicker, tc.inputStartDate, tc.inputEndDate)
		description := fmt.Sprintf("Test:%s GetRateByDate(%v,%v,%v)",
			tc.description, tc.inputTicker, tc.inputStartDate, tc.inputEndDate)
		fmt.Print(description)
		if got != tc.expectedResponse || !reflect.DeepEqual(err, tc.expectedError) {
			t.Fatalf("\nExpectedResponse: %v\nGot: %v\n"+
				"ExpectedError: %v\nGot: %v\n", tc.expectedResponse, got, tc.expectedError, err)
		} else {
			fmt.Println(":PASS")
		}
	}
}
