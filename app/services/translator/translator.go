package services

type TranslatorService struct {
	folder string
	locale string
}

func New(folder, locale string) TranslatorService {
	return TranslatorService{folder, locale}
}
