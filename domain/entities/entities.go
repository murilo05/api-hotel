package entities

type Error struct {
	ErrorMessage string `json:"errorMessage"`
}

type Env struct {
	HOST        string
	PORT        string
	DBHOST      string
	DBPORT      int
	DBUSER      string
	DBPASSAWORD string
	DBNAME      string
	TIMEOUT     int
}

type InputSchedule struct {
	Company Company `json:"empresa"`
	Hour    string  `json:"horario"`
}

type Company struct {
	CNPJ string `json:"cnpj"`
}

type ResponseGetSchudeles struct {
	Hour      string  `json:"horario"`
	Company   Company `json:"empresa"`
	FinalHour string  `json:"-"`
	AWS       AWS
}

type AWS interface {
}

type ResponseAvailability struct {
	StartingHour     string `json:"inicio"`
	FinalHour        string `json:"fim"`
	Available        bool   `json:"available"`
	AvailableFromSQL int    `json:"-"`
}
