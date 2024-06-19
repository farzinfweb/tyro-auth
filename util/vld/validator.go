package vld

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
)

func NewWithFa() (*validator.Validate, ut.Translator) {
	fa := fa.New()
	uni := ut.New(fa, fa)

	trans, _ := uni.GetTranslator("fa_IR")
	validate := validator.New(validator.WithRequiredStructEnabled())

	fa_translations.RegisterDefaultTranslations(validate, trans)

	return validate, trans
}

func NewWithEn() (*validator.Validate, ut.Translator) {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	validate := validator.New(validator.WithRequiredStructEnabled())

	en_translations.RegisterDefaultTranslations(validate, trans)

	return validate, trans
}
