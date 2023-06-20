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
	} else if len(flags) == 2 && flags[1] == "list" {
		cdc := keychain.GenerateCodec(flags[0], nil)
		fmt.Println(cdc.List())
	} else if len(flags) < 3 {
		panic("not enough arguments\nexpected: <password> <flag> [ARGS]")
	}

	cdc := keychain.GenerateCodec(flags[0], nil)

	if flags[1] == "set" {
		if len(flags) < 4 {
			panic("not enough arguments\nexpected: <password> set <domain> <username>")
		}
		fmt.Println(cdc.Set(flags[2], flags[3]))
	} else if flags[1] == "get" {
		fmt.Println(cdc.Get(flags[2]))
	} else if flags[1] == "update" {
		if len(flags) < 4 {
			panic("not enough arguments\nexpected: <password> update <domain> <password>")
		}
		fmt.Println(cdc.Update(flags[2], flags[3]))
	} else if flags[1] == "copy" {
		if len(flags) < 4 {
			panic("not enough arguments\nexpected: <password> copy <domain> <username>")
		}
		fmt.Println(cdc.Copy(flags[2], flags[3]))
	} else if flags[1] == "delete" {
		cdc.Delete(flags[2])
		fmt.Println("deleted")
	} else {
		fmt.Println("invalid flag")
		fmt.Println(keychain.HelpMessage())
	}
}
