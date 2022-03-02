package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Bananenpro/tictactoe-cli/game"
	"github.com/Bananenpro/tictactoe-cli/network"
)

var tictactoe *game.Game
var useAI *bool

func handleCommand(con *network.ServerConnection, command string) {
	if strings.HasPrefix(command, "error:") {
		fmt.Println(command)
		os.Exit(1)
		return
	}

	if command == "disconnect" {
		if tictactoe.IsRunning() {
			fmt.Println("The server closed the connection.")
		}
		return
	}

	if command == "ping" {
		err := con.Send("pong")
		if err != nil {
			os.Exit(1)
		}
		return
	}

	if tictactoe == nil {
		if strings.HasPrefix(command, "match-found:") {
			parts := strings.Split(command, ":")
			if len(parts) == 2 {
				sign := parts[1]
				if sign != "x" && sign != "o" {
					fmt.Println("Invalid sign:", sign)
				} else {
					StartGame(con, sign)
				}
			}
		}
		return
	}

	tictactoe.HandleCommand(command)
}

func StartGame(con *network.ServerConnection, sign string) {
	fmt.Printf("Match found (sign: %s)!\n", sign)
	time.Sleep(1 * time.Second)

	tictactoe = game.New(con, sign, *useAI)

	tictactoe.Start()
}

func main() {
	useAI = flag.Bool("ai", false, "Let the computer play for you")
	flag.Parse()

	con, err := network.Connect("julianh.de:7531")
	if err != nil {
		fmt.Println("Failed to connect to server: ", err)
		return
	}
	defer con.Close()

	fmt.Println("Looking for opponents...")
	con.Read(handleCommand)
}
