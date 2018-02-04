package main

import (
	//"math/big"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
	//"io/ioutil"
	"log"
	"strings"
	//"encoding/base64"
)

var key *ecdsa.PrivateKey

/*
Dear Rob Pike,
	Go is a decent language.  However, for the sake of all the 'gophers' out there, ADD SOME NORMAL ERROR HANDELING.

Pls don't kill me,
	The_Sushi/Ender
*/
func Handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RunCmd(cmd string) {
	if strings.HasPrefix(cmd, "send ") {
		//Send(strings.Split(cmd, " ")[1], strings.Split(cmd, " ")[2])
	}
	
	switch cmd {
	case "help":
		fmt.Println("Command list: help, send {ID}, genkey, pubkey, privkey, exit")
		
	case "genkey":
		var err error
		key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		Handle(err)
		
	case "pubkey":
		fmt.Println( key.X.String() + "::" + key.Y.String() )
		
	case "privkey":
		fmt.Println( key.D.String() )
	case "exit":
		os.Exit(0)
		
	default:
		fmt.Println("Error: Invalid command")
		
	}
}

func main() {
	fmt.Printf("\033[1;35m%s\033[0m\n", `
		▒█▀▀█ █░░ █▀▀█ █▀▀ █░█ ▒█▀▀█ █▀▀█ ░▀░ █▀▀▄ 
		▒█▀▀▄ █░░ █░░█ █░░ █▀▄ ▒█░░░ █░░█ ▀█  █░░█ 
		▒█▄▄█ ▀▀▀ ▀▀▀▀ ▀▀▀ ▀░▀ ▒█▄▄█ ▀▀▀▀ ▀▀▀ ▀░░▀
	`)
	fmt.Println(`Enter a command.  Run 'help' if you aren't familiar with the BlockCoin wallet software.`)
	var input string
	for {
		fmt.Print("> ")
		_, err := fmt.Scanln(&input)
		Handle(err)
		
		RunCmd(input)
	}
}