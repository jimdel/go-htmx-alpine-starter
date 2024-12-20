package main

import "jimdel/pkg/server"

const PORT = ":3000"

func main() {
	err := server.Run(PORT)
	if err != nil {
		panic(err)
	}
}
