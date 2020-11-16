package databaseConnection

type Country struct {
	Code             string     `json:"code"`
	Name             string     `json:"name"`
	Capital          string     `json:"capital"`
	OfficialLanguage string     `json:"officialLanguage"`
	Cities           []City     `json:"cities"`
	Languages        []Language `json:"languages"`
}

type City struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
	District    string `json:"dsitrict"`
	Info        string `json:"info"`
}

type Language struct {
	Language   string  `json:"language"`
	IsOfficial bool    `json:"isOfficial"`
	Percentage float32 `json:"percentage"`
}

//type Info struct {
//Population int `json:"Population"`
/*}*/

var (
	country  = Country{}
	city     = City{}
	language = Language{}
)
