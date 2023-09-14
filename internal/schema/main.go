package schema

type User struct {
	Id   string
	Name string
}

type NIKData struct {
	Id int64 `json:"id"`
	AreaData
	DateOfBirth string `json:"dob"`
	Gender      string `json:"gender"`
}

type AreaData struct {
	Province Province `json:"province"`
	City     City     `json:"city"`
	District District `json:"district"`
}

type Province struct {
	Id   int64  `json:"id"`
	Nama string `json:"nama"`
}

type City struct {
	Id   int64  `json:"id"`
	Nama string `json:"nama"`
}

type District struct {
	Id   int64  `json:"id"`
	Nama string `json:"nama"`
}