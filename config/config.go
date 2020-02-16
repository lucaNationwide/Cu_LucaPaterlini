package config

import "time"

const (
	// FullThirdPartyAPIPath contain the main url of the path to call to access the 3rd party api.
	FullThirdPartyAPIPath = "https://api.exchangeratesapi.io/"
	// DefaultRequestsTimeout contains the timeout time for the request of the 3rd party api.
	DefaultRequestsTimeout = 2 * time.Second
	// DefaultAddr contains the default address to bind to run the api server.
	DefaultAddr = ":8123"
	// ServerWriteTimeout sets a max for each request
	ServerWriteTimeout = time.Second * 20
	// ServerReadTimeout sets a max for each write request
	ServerReadTimeout = time.Second * 20
	// DefaultCurrency to use as target currency
	DefaultCurrency = "EUR"
)

// SupportedTickers is a list of supported tickers
var SupportedTickers = [2]string{"USD", "GBP"}
