package main

import (
	"Lineman_project/entity"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func fetchPatients() ([]entity.Patient, error) {
	// URL of the public API
	url := "https://static.wongnai.com/devinterview/covid-cases.json"

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 status codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received non-200 status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Parse the JSON response
	var response entity.Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return response.Data, nil
}

func main() {
	router := gin.Default()

	router.GET("/patients", func(ctx *gin.Context) {
		patients, err := fetchPatients()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, patients)
	})

	router.GET("/covid/summary", func(ctx *gin.Context) {
		patients, err := fetchPatients()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		provinceCount := make(map[string]int)
		ageGroupCount := map[string]int{
			"0-30":  0,
			"31-60": 0,
			"61+":   0,
			"N/A":   0,
		}

		for _, patient := range patients {
			// Count provinces
			if patient.Province != "" {
				provinceCount[patient.Province]++
			}
			// Count age groups
			if patient.Age == nil {
				ageGroupCount["N/A"]++
			} else {
				switch {
				case *patient.Age <= 30:
					ageGroupCount["0-30"]++
				case *patient.Age <= 60:
					ageGroupCount["31-60"]++
				default:
					ageGroupCount["61+"]++
				}
			}
		}

		summary := entity.SummaryResponse{
			Province: provinceCount,
			AgeGroup: ageGroupCount,
		}

		ctx.JSON(http.StatusOK, summary)
	})

	router.Run(":8080")
}
