package entities

type CarPart struct {
	Id           int64  `json:"id"`
	WinCode      string `json:"winCode"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	CarBand      string `json:"carBrand"`
	MainPhoto    string `json:"mainPhoto"`
}
