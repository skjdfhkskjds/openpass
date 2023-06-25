package main

import (
	"fmt"
	"main/keychain"
	"os"
)

func main() {
	flags := os.Args[1:]
	if len(flags) == 1 && flags[0] == "help" {
		fmt.Println(keychain.HelpMessage())
		return
	} else if len(flags) == 2 && flags[1] == "list" {
		cdc := keychain.GenerateCodec(flags[0], nil)
		fmt.Println(cdc.List())
		return
	} else if len(flags) < 3 {
		panic("not enough arguments\nexpected: <password> <flag> [ARGS]")
	}

	cdc := keychain.GenerateCodec(flags[0], nil)
	var result string

	domain := keychain.ReduceDomain(flags[2])
	if flags[1] == "set" {
		if len(flags) < 4 {
			panic("not enough arguments\nexpected: <password> set <domain> <username>")
		}
		result = (cdc.Set(domain, flags[3]))
	} else if flags[1] == "get" {
		result = (cdc.Get(domain))
	} else if flags[1] == "update" {
		if len(flags) < 4 {
			panic("not enough arguments\nexpected: <password> update <domain> <password>")
		}
		result = (cdc.Update(domain, flags[3]))
	} else if flags[1] == "copy" {
		if len(flags) < 4 {
			panic("not enough arguments\nexpected: <password> copy <domain> <username>")
		}
		result = (cdc.Copy(domain, flags[3]))
	} else if flags[1] == "delete" {
		cdc.Delete(domain)
		result = ("deleted")
	} else {
		result = ("invalid flag")
		result = (keychain.HelpMessage())
	}

	fmt.Println(result)
}
