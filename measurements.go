package main

import(
	"time"
	"net/http"
	"log"
	"encoding/json"
)

type CreateMeasurementParamsJSON struct {
	MeasurementDate time.Time `json:"measurement_date"`
	MeasurementTime time.Time `json:"measurement_time"`
	Pressure1       string	  `json:"pressure_1"`
	Pressure2       string	  `json:"pressure_2"`
	Temperature1    string	  `json:"temperature_1"`
	Temperature2    string	  `json:"temperature_2"`
}

func (cfg *apiConfig)handlerCreateMeasurement(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	params := CreateMeasurementParamsJSON{}
	err := decoder.Decode(&params)
	if err != nil{
		log.Println(err)
		respondWithError(w, http.StatusBadRequest,"Could not decode parameters", err)
		return
	}
	measurement, err := cfg.db.CreateMeasurement(r.Context(), params)
	if err != nil{
		log.Println(err)
		respondWithError(w,http.StatusInternalServerError,"Could not create measurement entry", err)
		return
	}
	respondWithJSON(w, 201, CreateMeasurementParamsJSON{
		MeasurementDate: measurement.MeasurementDate,
		MeasurementTime: measurement.MeasurementTime, 
		Pressure1: measurement.Pressure1,
		Pressure2: measurement.Pressure2,
		Temperature1: measurement.Temperature1,
		Temperature2: measurement.Temperature2,
	})
}

