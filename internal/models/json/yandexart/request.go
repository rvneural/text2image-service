package yandexart

type Message struct {
	Weight string `json:"weight"`
	Text   string `json:"text"`
}

type Request struct {
	ModelUri          string `json:"modelUri"`
	GenerationOptions struct {
		Seed        string `json:"seed"`
		AspectRatio struct {
			WidthRatio  string `json:"widthRatio"`
			HeightRatio string `json:"heightRatio"`
		} `json:"aspectRatio"`
	} `json:"generationOptions"`
	Messages []Message `json:"messages"`
}
