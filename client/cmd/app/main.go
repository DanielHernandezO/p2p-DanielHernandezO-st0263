package main

import "github.com/DanielHernandezO/p2p/client/internal/infrastructure/delivery/shell"

func main() {
	shellService := shell.NewShellCommands()
	shellService.Start()
}
