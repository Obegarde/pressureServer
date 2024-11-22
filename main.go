package main

import (
	"net/http"
	"log"
	_"github.com/lib/pq"
	"github.com/joho/godotenv"
	"os"
	"database/sql"
	"github.com/obegarde/pressureServer/internal/database"
)

func main(){

	//Load env file
	godotenv.Load()
	//get the env variables
	dbURL := os.Getenv("DB_URL")
	platform := os.Getenv("PLATFORM")
	secret := os.Getenv("SECRET")
	testApiKey := os.Getenv("TEST_API_KEY")
	// Open a db connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil{
		log.Printf("DB error: %b",err)
	}
	defer db.Close()

	//Create a mux
	mux := http.NewServeMux()
	//Create api config struct
	apiCfg := apiConfig{
		db:database.New(db),
		platform:platform,
		secret:secret,
		testApiKey:testApiKey,
	}
	//mux.HandleFunc goes here
	mux.HandleFunc("POST /api/measurements",apiCfg.handlerCreateBatchMeasurements)	
	mux.HandleFunc("GET /api/measurements", apiCfg.handlerGetMeasurements)	
	//Create a ServerStruct
	server := &http.Server{
	Addr: ":8080",
	Handler:mux,
	}
	
	//Launch the server
	log.Fatal(server.ListenAndServe())


}
