package carPartDtos

import "database-api/app/core/contracts/dtos/partPhotoDtos"

type CarPartDto struct {
	Id           int64  `json:"id"`
	WinCode      string `json:"winCode"`
	Price        int64  `json:"price"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Manufacturer string `json:"manufacturer"`
	CarBand      string `json:"carBrand"`
	MainPhoto    string `json:"mainPhoto"`
}

type CounterDto struct {
	WinCode string `json:"winCode"`
	Count   int64  `json:"count"`
}

type PartDetailDto struct {
	Id           int64                        `json:"id"`
	WinCode      string                       `json:"winCode"`
	Price        int64                        `json:"price"`
	Title        string                       `json:"title"`
	Description  string                       `json:"description"`
	Manufacturer string                       `json:"manufacturer"`
	CarBand      string                       `json:"carBrand"`
	PartPhotos   []partPhotoDtos.PartPhotoDto `json:"partImages"`
}

type AddPartDto struct {
	WinCode      string `json:"winCode"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	CarBand      string `json:"carBrand"`
	MainPhoto    string `json:"mainPhoto"`
}
