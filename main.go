package main

import "assigment_5/restAPI/router"

func main() {
	var PORT = ":4000"

	router.StartServer().Run(PORT)
}
