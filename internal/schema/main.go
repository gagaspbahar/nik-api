package schema

type User struct {
	Id   string
	Name string
}

type UserSchema struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	DistrictId  string `json:"district_id"`
	CityId      string `json:"city_id"`
	ProvinceId  string `json:"province_id"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"dob"`
}

type NIKData struct {
	Id int `json:"id"`
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
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type City struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}

type District struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
}
