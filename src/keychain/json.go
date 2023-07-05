package keychain

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func AppendJSON(jsonFile string, password Password) error {
	passwords := ReadJSON(jsonFile)

	passwords = append(passwords, password)

	jsonPasswords, err := json.MarshalIndent(passwords, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonFile, jsonPasswords, 0644)
	if err != nil {
		return err
	}

	return nil
}

// findFromJSON finds a password from a JSON file
func FindFromJSON(jsonFile, domain, username string) (Password, bool) {
	passwords := ReadJSON(jsonFile)

	for _, password := range passwords {
		if password.Domain == domain && password.Username == username {
			return password, true
		}
	}

	// Read the JSON file into a byte array
	return Password{}, false
}

// readJSON reads a JSON file and returns a slice of Passwords
func ReadJSON(jsonFile string) []Password {
	// Open the JSON file
	jsonData, err := ioutil.ReadFile(jsonFile)
	if err != nil || len(jsonData) == 0 {
		return []Password{}
	}

	var passwords []Password
	if err := json.Unmarshal(jsonData, &passwords); err != nil {
		panic(err)
	}

	return passwords
}