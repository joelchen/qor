package l10n

import "github.com/jinzhu/gorm"

type Interface interface {
	IsGlobal() bool
	SetLocale(locale string)
}

type Locale struct {
	LangaugeCode *string `sql:"size:6"`
}

func (l Locale) IsGlobal() bool {
	return l.LangaugeCode == nil
}

func (l Locale) SetLocale(locale string) {
	l.LangaugeCode = &locale
}

func Localize(scope *gorm.Scope, global Interface, locale string) {
	// find deleted locale -> reset deleted at
	// sync attrs from global to locale
}