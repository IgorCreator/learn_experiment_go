package main

import (
	"fmt"
	"strconv"
	"strings"
)

func menu() (int, error) {
	return fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > save   : exports the data to json and quits
  > quit   : quits

$> `)
}

func runCmd(in string, games []game, gamesByID map[int]game) bool {
	fmt.Println()

	cmd := strings.Fields(in)
	if len(cmd) == 0 {
		return true
	}

	switch cmd[0] {
	case "quit":
		return cmdQuit()
	case "list":
		return cmdList(games)
	case "id":
		return cmdByID(cmd, gamesByID)
	case "save":
		return cmdSave(games)
	}

	return true
}

func cmdByID(cmd []string, gamesByID map[int]game) bool {
	if len(cmd) != 2 {
		fmt.Println("wrong id")
		return true
	}

	id, err := strconv.Atoi(cmd[1])
	if err != nil {
		fmt.Println("wrong id")
		return true
	}

	g, ok := gamesByID[id]
	if !ok {
		fmt.Println("sorry. I don't have the game")
		return true
	}

	printGame(g)
	return true
}

func cmdList(games []game) bool {
	for _, g := range games {
		printGame(g)
	}
	return true
}

func cmdQuit() bool {
	fmt.Println("bye!")
	return false
}
