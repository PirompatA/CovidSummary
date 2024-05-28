package controller

import (
	"Lineman_project/api"
	"Lineman_project/entity"
)

type DataController interface {
	GetCovidSummary() entity.SummaryResponse
}

type dataController struct {
	summary    entity.SummaryResponse
	apiService api.ApiService
}

func New() DataController {
	return &dataController{}
}

func (c *dataController) GetCovidSummary() entity.SummaryResponse {
	patients, _ := c.apiService.FetchPatients()

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

	return summary
}
