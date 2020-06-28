package main

import api "github.com/RiccardoBiosas/golang-ethereum-auth/api"

func main() {
	a := api.Api{}
	a.Mount()
	a.Run(":8080")
}
