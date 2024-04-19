package locale

import (
	"Shelex/tg-watermark-bot/env"
	"encoding/json"
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var loc *i18n.Localizer

func getLanguageFromLocale(locale string) language.Tag {
	switch locale {
	case "uk":
		return language.Ukrainian
	default:
		return language.English
	}

}

func Register(config *env.Config) {
	bundle := i18n.NewBundle(getLanguageFromLocale(config.Locale))
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile(fmt.Sprintf("./locale/%s.json", config.Locale))
	loc = i18n.NewLocalizer(bundle, config.Locale)
}

func Translate(messageID string) string {
	return loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})
}

func Error(messageID string) error {
	return fmt.Errorf(Translate(messageID))
}

func WithError(messageID string, err error) error {
	return fmt.Errorf("%s %s", Translate(messageID), err.Error())
}
