package keychain

import (
	"encoding/json"
	"fmt"
	"os"
)

func (cdc *Codec) Set(domain, username string) string {
	_, found := FindFromJSON(jsonFile, domain, username)
	if !found {
		cdc.Delete(domain)
	}

	cdc.SetSalt(cdc.salt)
	password := cdc.generatePassword()
	if err := cdc.setPassword(domain, username, password); err != nil {
		panic(err)
	}

	return password
}

func (cdc *Codec) Get(domain, username string) (string, string) {
	password, found := FindFromJSON(jsonFile, domain, username)
	if !found {
		// TODO: so bad
		return "Password for " + domain + " not found", ""
	}

	cdc.SetSalt(password.Salt)
	pass, err := cdc.decrypt(password.Password)
	if err != nil {
		panic(err)
	}

	return username, pass
}

func (cdc *Codec) Update(domain, username, password string) string {
	cdc.Delete(domain)
	cdc.SetSalt(cdc.salt)
	if err := cdc.setPassword(domain, username, password); err != nil {
		panic(err)
	}

	return password
}

func (cdc *Codec) Copy(domain1, username1, domain2, username2 string) string {
	_, password := cdc.Get(domain1, username1)
	return cdc.Update(domain2, username2, password)
}

func (cdc *Codec) Delete(domain string) {
	passwords := ReadJSON(jsonFile)

	for i, password := range passwords {
		if password.Domain == domain {
			passwords = Remove(passwords, i)
			break
		}
	}

	jsonPasswords, err := json.Marshal(passwords)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(jsonFile, jsonPasswords, 0644)
	if err != nil {
		panic(err)
	}

	// reset the salt
	cdc.salt = nil
}

func (cdc *Codec) List() string {
	passwords := ReadJSON(jsonFile)
	output := ""
	for _, password := range passwords {
		pass, err := cdc.decrypt(password.Password)
		if err != nil {
			panic(err)
		}
		output += "-------------------\n"
		output += fmt.Sprintf("domain: %s\n", password.Domain)
		output += fmt.Sprintf("username: %s\n", password.Username)
		output += fmt.Sprintf("password: %s\n", pass)
		output += "-------------------\n\n"
	}
	return output
}
