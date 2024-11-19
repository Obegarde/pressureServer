package main

import(	
	"github.com/obegarde/pressureServer/internal/database"
	

)

type apiConfig struct{
	db *database.Queries
	platform string
	secret string
	testApiKey string
}


