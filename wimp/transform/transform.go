package transform

import "image"

type Transformer interface {
	Transform(image.Image) image.Image
}

var (
	transformers = make(map[string]Transformer)
)

func RegisterTransformer(name string, t Transformer) {
	transformers[name] = t
}

func GetTransformer(name string) Transformer {
	return transformers[name]
}