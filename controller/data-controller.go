package controller

import (
	"Lineman_project/client"
	"Lineman_project/entity"
	"net/http"
)

type DataController interface {
	ShowSummary() entity.SummaryResponse
}

type dataController struct {
	patients    []entity.Patient
	patientData client.FetchData
}

func New(patientData client.FetchData) DataController {
	return &dataController{
		patientData: patientData,
	}
}

func (service *dataController) ShowSummary() entity.SummaryResponse {
	patients, err := service.patientData.FetchPatients("https://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		return err
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

	summary := Summary{
		Province: provinceCount,
		AgeGroup: ageGroupCount,
	}

	ctx.JSON(http.StatusOK, summary)
})
}

