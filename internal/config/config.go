package config

// PSQLConnection need for Postgres connection
type PSQLConnection struct {
	Host     string
	Port     string
	Database string
	Password string
	Username string
}

// ExchangeAPIConvert is request-model for ExchangeRates API
type ExchangeAPIConvert struct {
	AccessToken string
	From        string
	To          string
	Amount      float64
}
