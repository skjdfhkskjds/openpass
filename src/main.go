package main

import (
	"fmt"
	"main/common"
	"main/keychain"
	"os"
)

func main() {
	flags := os.Args[1:]
	if len(flags) == 1 && flags[0] == "help" {
		fmt.Println(common.HelpMessage())
		return
	} else if len(flags) == 2 && flags[1] == "list" {
		cdc := keychain.GenerateCodec(flags[0], nil)
		fmt.Println(cdc.List())
		return
	} else if len(flags) < 3 {
		panic("not enough arguments\nexpected: <password> <flag> [ARGS]")
	}

	cdc := keychain.GenerateCodec(flags[0], nil)

	var (
		username string
		password string
	)

	domain := keychain.ReduceDomain(flags[2])
	username = flags[3]
	if flags[1] == "set" {
		if len(flags) < 4 {
			panic("not enough arguments\nexpected: <password> set <domain> <username>")
		}
		password = (cdc.Set(domain, flags[3]))
	} else if flags[1] == "get" {
		username, password = (cdc.Get(domain, username))
	} else if flags[1] == "update" {
		if len(flags) < 5 {
			panic("not enough arguments\nexpected: <password> update <domain> <password>")
		}
		password = (cdc.Update(domain, username, flags[4]))
	} else if flags[1] == "copy" {
		if len(flags) < 6 {
			panic("not enough arguments\nexpected: <password> copy <domain> <username>")
		}
		password = (cdc.Copy(domain, username, flags[4], flags[5]))
	} else if flags[1] == "delete" {
		cdc.Delete(domain)
		password = ("deleted")
	} else {
		password = ("invalid flag")
		password = (common.HelpMessage())
	}

	fmt.Println(username)
	fmt.Println(password)
}
