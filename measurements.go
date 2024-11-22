package main

import(
	"time"
	"net/http"
	"log"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/obegarde/pressureServer/internal/database"	
	"fmt"
)

type CreateMeasurementParamsJSON struct {
	MeasurementDate string `json:"measurement_date"`
	MeasurementTime string `json:"measurement_time"`
	Pressure1       float64	  `json:"pressure_1"`
	Pressure2      	float64	  `json:"pressure_2"`
	Temperature1    float64	`json:"temperature_1"`
	Temperature2    float64	  `json:"temperature_2"`
}
type MeasurementJSON struct {
	ID              uuid.UUID `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	MeasurementDate time.Time `json:"measurement_date" time_format:"2006-01-02"`
	MeasurementTime time.Time `json:"measurement_time" time_format:"15:04:05"`
	Pressure1       float64 `json:"pressure_1"`
	Pressure2       float64 `json:"pressure_2"`
	Temperature1    float64 `json:"temperature_1"`
	Temperature2    float64 `json:"temperature_2"`
}



func (cfg *apiConfig)handlerCreateMeasurements(w http.ResponseWriter, r *http.Request){
	//Verify Api key on file
	apiKey,err := GetApiKey(r.Header)
	if err != nil {
		respondWithError(w,http.StatusUnauthorized,"401 Unauthorized access", err)
		return
	}
	if apiKey != cfg.testApiKey{	
		respondWithError(w,http.StatusUnauthorized,"401 Unauthorized access", err)
		return
	}
	//Make a json decoder 
	decoder := json.NewDecoder(r.Body)
	// Create a list of parameters so we can feed in more than measurement at a time
	params := []CreateMeasurementParamsJSON{}
	// Decode into the params list
	err = decoder.Decode(&params)
	if err != nil{
		log.Println(err)
		respondWithError(w, http.StatusBadRequest,"Could not decode parameters", err)
		return
	}
	// Create a slice to hold the measurement responses from the database to respond with
	responseJSONList := []MeasurementJSON{}
	

	//Iterate over the list of params
	for _, param := range params{
		//Use timeFormatter to convert from string to time.Time
		dateTime, err := timeFormatter(param.MeasurementDate, param.MeasurementTime)
		if err != nil {
			log.Println(err)
			respondWithError(w, http.StatusInternalServerError,"Failed to convert timestrings to Time", err)
			return
		}
		convertedParams := database.CreateMeasurementParams{
			MeasurementDate:dateTime.Date,
			MeasurementTime:dateTime.Time,
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


func (cfg *apiConfig)handlerCreateBatchMeasurements(w http.ResponseWriter, r *http.Request){
	//Verify Api key on file
	apiKey,err := GetApiKey(r.Header)
	if err != nil {
		respondWithError(w,http.StatusUnauthorized,"401 Unauthorized access", err)
		return
	}
	if apiKey != cfg.testApiKey{	
		respondWithError(w,http.StatusUnauthorized,"401 Unauthorized access", err)
		return
	}
	//Make a json decoder 
	decoder := json.NewDecoder(r.Body)
	// Create a list of parameters so we can feed in more than measurement at a time
	params := []CreateMeasurementParamsJSON{}
	// Decode into the params list
	err = decoder.Decode(&params)
	if err != nil{
		log.Println(err)
		respondWithError(w, http.StatusBadRequest,"Could not decode parameters", err)
		return
	}
	fmt.Println(len(params))
	// Create slices for the batch insert
	dates := make([]time.Time, len(params))
	times := make([]time.Time, len(params))
	pressure1s := make([]float64, len(params))
	pressure2s := make([]float64, len(params))
	temperature1s := make([]float64, len(params))
	temperature2s := make([]float64, len(params))
	//Iterate over the params splitting them out 
	for i, p := range params{
		dateTime, err := timeFormatter(p.MeasurementDate, p.MeasurementTime)
		if err != nil{
			log.Println("Failed to convert time")
			continue
		}
		dates[i] = dateTime.Date
		times[i] = dateTime.Time
		pressure1s[i] = p.Pressure1
		pressure2s[i] = p.Pressure2
		temperature1s[i] = p.Temperature1
		temperature2s[i] = p.Temperature2

	}
	batchParams := database.CreateMeasurementsBatchParams{
		Column1: dates,
		Column2: times,
		Column3: pressure1s,
		Column4: pressure2s,
		Column5: temperature1s,
		Column6: temperature2s,
	}

	err = cfg.db.CreateMeasurementsBatch(r.Context(), batchParams)
	if err != nil{
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError,"Could not insert into database", err)
	}
	respondWithError(w, http.StatusNoContent,"Success, no return content", fmt.Errorf("No content"))
}
