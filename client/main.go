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
	"log"
	"strings"
	//"strconv"
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

func Send(from, to, node, amount string, priv *ecdsa.PrivateKey) {
	var input string
	conn, err := net.Dial("tcp", node)
	Handle(err)

	fmt.Println("Connected to node")

	for {
		fmt.Print("Would you like to send " + amount + "ƁƇ to "+to + " (Y/N): ")
		_, err := fmt.Scanln(&input)
		Handle(err)
		
		if input == "n" || input == "N" || input == "no" || input == "No" {
	 		fmt.Println("Cancelled transaction")
	 		return
		}	else if input != "Y" && input != "y" && input != "yes" && input != "Yes" {
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
	fmt.Println("Signed: " + string(signed))
	fmt.Fprintf(conn, "[%s]:%s->[%s]  %s", from, amount, to, signed)
	fmt.Println("Sent!")
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