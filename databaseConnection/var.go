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
	Id       string `json:"id"`
	Name     string `json:"name"`
	District string `json:"dsitrict"`
}

type Language struct {
	Language   string  `json:"language"`
	IsOfficial bool    `json:"isOfficial"`
	Percentage float32 `json:"percentage"`
}

var (
	country  = Country{}
	city     = City{}
	language = Language{}
)
