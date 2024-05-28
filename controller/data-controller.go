package controller


	 {
		patients, err := api.fetchPatients()
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


