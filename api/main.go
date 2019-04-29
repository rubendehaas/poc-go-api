package main

import (
	"app/database"
	"app/router"
)

func init() {

	database.Init()
	router.Init()
}

func main() {

}
