package config

// CurrencyURL is URL to get exchange rates
var CurrencyURL string = "http://api.exchangeratesapi.io/v1/"

// PSQLConnection need for Postgres connection
type PSQLConnection struct {
	Host     string
	Port     string
	Database string
	Password string
	Username string
}

// ExchangeAPIConvert is request-model to ExchangeRates API
type ExchangeAPIConvert struct {
	AccessToken string
	From        string
	To          string
	Amount      float64
}
