package entities

type PartPhoto struct {
	Id        int64  `json:"id"`
	PartId    int64  `json:"partId"`
	PartPhoto string `json:"partPhoto"`
}
