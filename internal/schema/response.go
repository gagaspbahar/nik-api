package schema

type ExtractResponse struct {
	Id       string `json:"id"`
	Dob      string `json:"dob"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Gender   string `json:"gender"`
}

type ValidateDataResponse struct {
	Id    string   `json:"id"`
	Error []string `json:"error"`
}
