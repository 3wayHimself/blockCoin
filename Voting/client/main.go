package main

import (
	"crypt/ecdsa",
	"fmt",
	"log",
	"strings"
)

func RunCmd(cmd string) {
	if strings.HasPrefix(cmd, "loadfile ") {
		
	} else if strings.HasPrefix(cmd, "send ") {
		SendTo(cmd[4:])
	}
	switch cmd {
	case "help":
		fmt.Println("Command list: help, loadfile {file}, send {ID}, status, genkey, pubkey, privkey")
	case "genkey":
		GenKey()
	case "status":
		fmt.Println( GetInfo() )
	case "pubkey":
		fmt.Println( SmallKey( key.PublicKey().X ) + "::" + SmallKey( key.PublicKey().Y ) )
	case "privkey":
		fmt.Println( SmallKey(key.D) )
	default:
		fmt.Println("Error: Invalid command")
	}
}

func main() {
	fmt.Println(`
		▒█▀▀█ █░░ █▀▀█ █▀▀ █░█ ▒█▀▀█ █▀▀█ ░▀░ █▀▀▄ 
		▒█▀▀▄ █░░ █░░█ █░░ █▀▄ ▒█░░░ █░░█ ▀█  █░░█ 
		▒█▄▄█ ▀▀▀ ▀▀▀▀ ▀▀▀ ▀░▀ ▒█▄▄█ ▀▀▀▀ ▀▀▀ ▀░░▀
	`)
	fmt.Println(`Enter a command.  Run 'help' if you aren't familiar with the BlockCoin wallet software.`)
	var input string
	for 1 {
		fmt.Print("> ")
		_, err = fmt.Scanln(&input)
		if err != nil {
			log.Fatal(err)
		}
		RunCmd(input)
	}
}