package operator

type Cq9ResponseBean struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
	Trance string      `json:"trance"`
}

type Status struct {
	Code     int    `json:"code"`
	Msg      string `json:"message"`
	DataTime string `json:"dataTime"`
}
