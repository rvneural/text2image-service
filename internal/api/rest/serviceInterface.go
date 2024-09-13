package rest

type Service interface {
	ConvertTextToImage(prompt, seed, widthRatio, heightRatio string) (string, string, error)
}
