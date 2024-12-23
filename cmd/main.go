package main

import server "github.com/geoffrey-anto/lugia/internal"

func main() {
	server := server.NewServer("0.0.0.0", 3000, "d15fd29f-3dfb-4211-994b-8969a3a4c5e7")

	server.Start()
}
