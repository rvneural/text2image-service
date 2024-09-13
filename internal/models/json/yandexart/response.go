package yandexart

type Response struct {
	Id       string `json:"id"`
	Done     bool   `json:"done"`
	Error    string `json:"error"`
	Response struct {
		Image string `json:"image"`
	} `json:"response"`
}
