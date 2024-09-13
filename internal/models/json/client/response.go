package client

type Response struct {
	Error string `json:"error,omitempty"`
	Image struct {
		B64String string `json:"b64String"`
		Seed      string `json:"seed"`
	} `json:"image"`
}

type Error struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}
