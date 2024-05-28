package entity

// Patient represents the structure of each item in the "Data" array
type Patient struct {
	ConfirmDate    string  `json:"ConfirmDate"`
	No             *int    `json:"No"`  // Pointer to handle null values
	Age            *int    `json:"Age"` // Pointer to handle null values
	Gender         string  `json:"Gender"`
	GenderEn       string  `json:"GenderEn"`
	Nation         *string `json:"Nation"` // Pointer to handle null values
	NationEn       string  `json:"NationEn"`
	Province       string  `json:"Province"`
	ProvinceId     int     `json:"ProvinceId"`
	District       *string `json:"District"` // Pointer to handle null values
	ProvinceEn     string  `json:"ProvinceEn"`
	StatQuarantine int     `json:"StatQuarantine"`
}
