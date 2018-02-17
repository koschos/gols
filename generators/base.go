package generators

type HashGeneratorInterface interface {
	GenerateHash(str string) string
}

type SlugGeneratorInterface interface {
	GenerateSlug() string
}
