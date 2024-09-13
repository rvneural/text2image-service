package client

type Request struct {
	Prompt      string `json:"prompt" xml:"prompt" form:"prompt"`
	Seed        string `json:"seed" xml:"seed" form:"seed"`
	WidthRatio  string `json:"widthRatio" xml:"widthRatio" form:"widthRatio"`
	HeightRatio string `json:"heightRatio" xml:"heightRatio" form:"heightRatio"`
}
