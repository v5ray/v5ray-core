package socks

import "github.com/v4fly/v4ray-core/v0/common/protocol"

func (a *Account) Equals(another protocol.Account) bool {
	if account, ok := another.(*Account); ok {
		return a.Username == account.Username
	}
	return false
}

func (a *Account) AsAccount() (protocol.Account, error) {
	return a, nil
}

func (c *ServerConfig) HasAccount(username, password string) bool {
	if c.Accounts == nil {
		return false
	}
	storedPassed, found := c.Accounts[username]
	if !found {
		return false
	}
	return storedPassed == password
}
