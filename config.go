package main

import(
	"database/sql"
	"github.com/obegarde/pressure/internal/database"

)

type apiConfig struct{
	db *database.Queries
	platform string
	secret string
	testApiKey string
}


