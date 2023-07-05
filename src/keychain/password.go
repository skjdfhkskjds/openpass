package keychain

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	jsonFile = "./passwords.json"

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
func (cdc *Codec) generatePassword() string {
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

func (cdc *Codec) setPassword(domain, username, password string) error {
	encryptedPassword, err := cdc.encrypt(password)
	if err != nil {
		return err
	}

	result := Password{
		Domain:   domain,
		Username: username,
		Password: encryptedPassword,
		Salt:     cdc.salt,
	}

	if err := AppendJSON(jsonFile, result); err != nil {
		return err
	}

	return nil
}

func ReduceDomain(domain string) string {
	parsedURL, err := url.Parse(domain)
	if err != nil || parsedURL.Host == "" {
		// Handle parsing error
		return domain
	}

	return parsedURL.Host
}

func Remove(slice []Password, i int) []Password {
	slice[i] = slice[len(slice)-1]   // Copy last element to index i.
	slice[len(slice)-1] = Password{} // Erase last element (write zero value).
	return slice[:len(slice)-1]
}
