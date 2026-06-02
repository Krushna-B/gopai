package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123456",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "SDJFY",
		Username:  "jaosn",
	},
	"marie": {
		AuthToken: "&@*$US",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    50,
		Username: "jaosn",
	},
	"marie": {
		Coins:    1000,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate database call
	time.Sleep(time.Second)

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate database call
	time.Sleep(time.Second)

	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
