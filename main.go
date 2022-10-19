package main

import (
	"os"
	"os/signal"
	"log"

	"github.com/Ahdeyyy/spotify2018Frontend/handlers"
	"github.com/Ahdeyyy/spotify2018Frontend/database"
	_ "github.com/Ahdeyyy/spotify2018Frontend/models"
)

func strToBool(str string) bool {
	if str == "true" || str == "True" {
		return true
	}
	return false
}

func main() {

	// read env file

	dB := db.OpenDatabase()
	// arr,_ := models.ParseCsv("data/top2018.csv")
	// sngs := models.ParseSong(arr)
	// db.InsertData(dB, sngs)


	defer dB.Close()

	server := NewCustomServer(os.Getenv("PORT"), false)
	server.addHandler("/", &handlers.Home{ Db : dB } )
	server.addHandler("/search", &handlers.Search{Db: dB} )
	server.addHandler("/api/",&handlers.Index{ Db: dB})
	log.Println("Server started on port", os.Getenv("PORT"))
	
	go server.Start()

	// Wait for a signal to quit:
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")


}
