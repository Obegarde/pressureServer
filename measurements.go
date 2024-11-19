package main

import(
	"time"
	"net/http"
	"log"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/obegarde/pressureServer/internal/database"
)

type CreateMeasurementParamsJSON struct {
	MeasurementDate time.Time `json:"measurement_date"`
	MeasurementTime time.Time `json:"measurement_time"`
	Pressure1       string	  `json:"pressure_1"`
	Pressure2       string	  `json:"pressure_2"`
	Temperature1    string	  `json:"temperature_1"`
	Temperature2    string	  `json:"temperature_2"`
}
type MeasurementJSON struct {
	ID              uuid.UUID `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	MeasurementDate time.Time `json:"measurement_date"`
	MeasurementTime time.Time `json:"measurement_time"`
	Pressure1       string `json:"pressure_1"`
	Pressure2       string `json:"pressure_2"`
	Temperature1    string `json:"temperature_1"`
	Temperature2    string `json:"temperature_2"`
}



func (cfg *apiConfig)handlerCreateMeasurements(w http.ResponseWriter, r *http.Request){
	//Make a json decoder 
	decoder := json.NewDecoder(r.Body)
	// Create a list of parameters so we can feed in more than measurement at a time
	params := []CreateMeasurementParamsJSON{}
	// Decode into the params list
	err := decoder.Decode(&params)
	if err != nil{
		log.Println(err)
		respondWithError(w, http.StatusBadRequest,"Could not decode parameters", err)
		return
	}
	// Create a slice to hold the measurement responses from the database to respond withb
	responseJSONList := []MeasurementJSON{}
	//
	for _, param := range params{
		convertedParams := database.CreateMeasurementParams{
			MeasurementDate:param.MeasurementDate,
			MeasurementTime:param.MeasurementTime,
			Pressure1:param.Pressure1,
			Pressure2:param.Pressure2,
			Temperature1:param.Temperature1,
			Temperature2:param.Temperature2,
		}
		measurement, err := cfg.db.CreateMeasurement(r.Context(),convertedParams)
		if err != nil{
			log.Println(err)
			respondWithError(w,http.StatusInternalServerError,"Could not create measurement entry", err)
			return
		}	
		JSONmeasurement := MeasurementJSON{
				ID:measurement.ID,
				CreatedAt:measurement.CreatedAt,
				MeasurementDate:measurement.MeasurementDate,
				MeasurementTime:measurement.MeasurementTime,
				Pressure1:measurement.Pressure1,
				Pressure2:measurement.Pressure2,
				Temperature1:measurement.Temperature1,
				Temperature2:measurement.Temperature2,	
				
	}
		responseJSONList = append(responseJSONList,JSONmeasurement)	
	}
	
	respondWithJSON(w, 201, responseJSONList)
}

func (cfg *apiConfig)handlerGetMeasurements(w http.ResponseWriter, r *http.Request){
	allMeasurements, err := cfg.db.GetMeasurements(r.Context())
	if err != nil{
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError,"Failed to retrieve measurements",err)
		return
	}
	responseJSONList := []MeasurementJSON{}
	for _, measurement := range allMeasurements{	
		JSONmeasurement := MeasurementJSON{
				ID:measurement.ID,
				CreatedAt:measurement.CreatedAt,
				MeasurementDate:measurement.MeasurementDate,
				MeasurementTime:measurement.MeasurementTime,
				Pressure1:measurement.Pressure1,
				Pressure2:measurement.Pressure2,
				Temperature1:measurement.Temperature1,
				Temperature2:measurement.Temperature2,	
				
	}
		responseJSONList = append(responseJSONList,JSONmeasurement)	
	}
	respondWithJSON(w, 201, responseJSONList)
	} 

