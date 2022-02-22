package main

import (
	"fmt"
	"time"

	"github.com/Bananenpro/tictactoe-cli/game"
	"github.com/Bananenpro/tictactoe-cli/network"
)

func main() {
	con, err := network.Connect("127.0.0.1:7531")
	if err != nil {
		fmt.Println("Failed to connect to server: ", err)
		return
	}
	defer con.Close()

	fmt.Println("Looking for opponents...")
	sign, err := con.WaitForMatch()
	if err != nil {
		fmt.Println("Failed join a match:", err)
		return
	}
	fmt.Printf("Match found (sign: %s)!\n", sign)
	time.Sleep(1 * time.Second)

	g := game.New(con)
	g.Start()
}
