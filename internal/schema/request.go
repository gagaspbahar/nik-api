package schema

type SubmitIdsRequest struct {
	Data []User `json:"data"`
}

type ValidateDataRequest struct {
	Data []ValidateDataStruct `json:"data"`
}

type ValidateDataStruct struct {
	Id       string `json:"id"`
	Dob      string `json:"dob"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Gender   string `json:"gender"`
}
