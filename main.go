package main

import (
	"os"
	"os/signal"
	"log"
)

func strToBool(str string) bool {
	if str == "true" || str == "True" {
		return true
	}
	return false
}

func main() {

	// read env file

	db := openDatabase()
	// arr,_ := parseCsv("data/top2018.csv")
	// sngs := parseSong(arr)
	// insertData(db, sngs)


	defer db.Close()

	server := NewCustomServer(os.Getenv("PORT"), false)
	server.addHandler("/", &Home{ db : db } )
	server.addHandler("/search", &Search{db: db} )
	server.addHandler("/api/",&Index{db: db})
	log.Println("Server started on port", os.Getenv("PORT"))
	
	go server.Start()

	// Wait for a signal to quit:
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")


}
