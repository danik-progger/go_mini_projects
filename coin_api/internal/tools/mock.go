package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"michael": {
		AuthToken: "456DEF",
		Username:  "michael",
	},
	"alice": {
		AuthToken: "789GHI",
		Username:  "alice",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    200,
		Username: "alex",
	},
	"michael": {
		Coins:    320,
		Username: "michael",
	},
	"alice": {
		Coins:    1000,
		Username: "alice",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
