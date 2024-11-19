package client

type Request struct {
	Operation_ID string `json:"operation_id" xml:"operation_id" form:"operation_id"`
	Prompt       string `json:"prompt" xml:"prompt" form:"prompt"`
	Seed         string `json:"seed" xml:"seed" form:"seed"`
	WidthRatio   string `json:"widthRatio" xml:"widthRatio" form:"widthRatio"`
	HeightRatio  string `json:"heightRatio" xml:"heightRatio" form:"heightRatio"`
}

type DBResult struct {
	Prompt    string `json:"prompt"`
	Seed      string `json:"seed"`
	B64string string `json:"image"`
	Name      string `json:"name"`
}
