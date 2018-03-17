package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"os"
	"net"
	"strings"
	"bufio"
	"encoding/base64"
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
		panic(err)
	}
}

func Send(node, to, amount string, priv *ecdsa.PrivateKey) {
	var input string
	from := priv.X.String() + "::" + priv.Y.String()
	conn, err := net.Dial("tcp", node)
	Handle(err)

	fmt.Println("Connected to node")

	for {
		fmt.Print("Would you like to send " + amount + "BC to "+to + " (Y/N): ")
		_, err := fmt.Scanln(&input)
		Handle(err)
		
		if input == "n" || input == "N" || input == "no" || input == "No" {
			fmt.Println("Cancelled transaction")
			return
		} else if input != "Y" && input != "y" && input != "yes" && input != "Yes" {
			fmt.Println("Not a valid option")
		} else {
			break
		}
	}

	fmt.Println("Signing...")
	hashed := sha256.New()
	hashed.Write([]byte(amount + "->" + to))
	signed, err := priv.Sign( rand.Reader, hashed.Sum(nil), crypto.SHA256 )
	Handle(err)
	fmt.Println("Signed: " + base64.StdEncoding.EncodeToString(signed))
	fmt.Fprintf(conn, "%s||%s->%s##%s\n", from, amount, to, base64.StdEncoding.EncodeToString(signed))
	fmt.Println("Sent!")
}

func RunCmd(cmd string) {
	
	switch cmd {
	case "help":
		fmt.Println("Command list: help, send {node address} {ID} {amount}, genkey, pubkey, privkey, exit")
		
	case "genkey":
		var err error
		key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		Handle(err)
		
	case "pubkey":
		if key != nil {
			fmt.Println( key.X.String() + "::" + key.Y.String() )
		} else {
			println("Error: No key loaded!")
		}
		
	case "privkey":
		if key != nil {
			fmt.Println( key.D.String() )
		} else {
			println("Error: No key loaded!")
		}
	case "exit":
		os.Exit(0)
		
	default:
		if strings.HasPrefix(cmd, "send ") {
			if key != nil {
				args := strings.Split(cmd, " ")
				Send(args[1], args[2], args[3], key)
			} else {
				println("Error: No key loaded!")
			}
		} else {
			fmt.Println("Error: Invalid command")
		}
	}
}

func main() {
	fmt.Printf("\033[1;35m%s\033[0m\n", `
		▒█▀▀█ █░░ █▀▀█ █▀▀ █░█ ▒█▀▀█ █▀▀█ ░▀░ █▀▀▄ 
		▒█▀▀▄ █░░ █░░█ █░░ █▀▄ ▒█░░░ █░░█ ▀█░ █░░█ 
		▒█▄▄█ ▀▀▀ ▀▀▀▀ ▀▀▀ ▀ ▀ ▒█▄▄█ ▀▀▀▀ ▀▀▀ ▀  ▀
	`)
	fmt.Println(`Enter a command.  Run 'help' if you aren't familiar with the BlockCoin wallet software.`)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		Handle(scanner.Err())
		input := scanner.Text()
		
		RunCmd(input)
	}
}