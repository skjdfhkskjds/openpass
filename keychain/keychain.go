package keychain

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	jsonFile = "./keychain/passwords.json"

	passwordLength = 12
	digits         = "0123456789"
	specials       = "~=+%^*/()[]{}/!@#$?|"
	all            = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + digits + specials
)

type Password struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     []byte `json:"salt"`
}

// generatePassword generates a password with at least one
// digit and one special character
func (cdc Codec) generatePassword() string {
	rand.Seed(time.Now().UnixNano())
	passBytes := make([]byte, passwordLength)
	passBytes[0] = digits[rand.Intn(len(digits))]
	passBytes[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < passwordLength; i++ {
		passBytes[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(passBytes), func(i, j int) {
		passBytes[i], passBytes[j] = passBytes[j], passBytes[i]
	})
	password := string(passBytes)
	return password
}

// TODO: return error if pass already exists
func (cdc Codec) Set(domain, username string) string {
	cdc.SetSalt(cdc.salt)
	password := cdc.generatePassword()
	encryptedPassword, err := cdc.encrypt(password)
	if err != nil {
		panic(err)
	}

	result := Password{
		Domain:   domain,
		Username: username,
		Password: encryptedPassword,
		Salt:     cdc.salt,
	}

	if err := appendJSON(jsonFile, result); err != nil {
		panic(err)
	}

	return password
}

func (cdc Codec) Get(domain string) string {
	password := findFromJSON(jsonFile, domain)
	if password.Domain == "" {
		return "Password for " + domain + " not found"
	}

	cdc.SetSalt(password.Salt)
	pass, err := cdc.decrypt(password.Password)
	if err != nil {
		panic(err)
	}

	return pass
}

func (cdc Codec) Update(domain string, password string) string {
	username := findFromJSON(jsonFile, domain).Username

	cdc.Delete(domain)
	cdc.SetSalt(cdc.salt)
	encryptedPassword, err := cdc.encrypt(password)
	if err != nil {
		panic(err)
	}

	result := Password{
		Domain:   domain,
		Username: username,
		Password: encryptedPassword,
		Salt:     cdc.salt,
	}

	if err := appendJSON(jsonFile, result); err != nil {
		panic(err)
	}
	return password
}

func (cdc Codec) Copy(domain1, domain2 string) string {
	password := cdc.Get(domain1)
	return cdc.Update(domain2, password)
}

func (cdc Codec) Delete(domain string) {
	passwords := readJSON(jsonFile)

	for i, password := range passwords {
		if password.Domain == domain {
			passwords = remove(passwords, i)
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

func (cdc Codec) List() string {
	passwords := readJSON(jsonFile)
	output := ""
	for _, password := range passwords {
		pass, err := cdc.decrypt(password.Password)
		if err != nil {
			panic(err)
		}
		output += fmt.Sprintf("domain: %s\npassword: %s\n", password.Domain, pass)
	}
	return output
}
