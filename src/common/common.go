package common

// TODO: deprecated message info
func HelpMessage() string {
	return "'help': display this message\n" +
		"'<keychain password> set <domain> <username>': set a password for a domain\n" +
		"'<keychain password> get <domain>': get a password for a domain\n" +
		"'<keychain password> update <domain> <password>': update a password for a domain\n" +
		"'<keychain password> copy <domain> <username>': copy a password for a domain to the clipboard\n" +
		"'<keychain password> delete <domain>': delete a password for a domain\n" +
		"'<keychain password> list': lists all passwords"
}
