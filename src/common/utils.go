package common

import (
	"net/url"
	"os"
	"path/filepath"
	"runtime"
)

// ReduceDomain reduces a domain to its base domain
// e.g. https://www.google.com -> google.com
func ReduceDomain(domain string) string {
	parsedURL, err := url.Parse(domain)
	if err != nil || parsedURL.Host == "" {
		// Handle parsing error
		return domain
	}

	return parsedURL.Host
}

// GetPathOfCaller returns the absolute path of the file
// that called this function
func GetPathOfCaller(filePath string) string {
	_, absFilePath, _, ok := runtime.Caller(2)
	if !ok {
		PanicRed("Failed to get path of file %s", filePath)
	}
	joinedPath := filepath.Join(filepath.Dir(absFilePath), filePath)
	return filepath.Clean(joinedPath)
}

// WriteFile writes data to file at filePath without
// overwriting the file
func WriteFile(filePath string, data []byte) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
