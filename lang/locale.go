package lang

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Locale is a locale.
// It can be used to translate strings.
type Locale struct {
	// The language code.
	LangCode string
}

// NewLocale creates a new Locale.
func NewLocale(langCode string) Locale {
	return Locale{
		LangCode: langCode,
	}
}

// Tr translates a string using the specified key and parameters.
// If the specified key does not exist, it will return a placeholder using the key.
// Example placeholder: "{{something.that.doesnt.exist}}"
//
// If the desired language is not supported, it will use the default language.
func (l Locale) Tr(key string, params ...string) string {
	tr, has := Translations[key]
	if !has {
		// The key did not exist, so return a placeholder.
		return fmt.Sprintf("{{%s}}", key)
	}

	str, has := tr[l.LangCode]
	if !has {
		// Use the default language if the desired language is not supported.
		str = tr[DefaultLangCode]
	}

	// Fill in parameters.
	for i, param := range params {
		str = strings.ReplaceAll(str, "$"+strconv.Itoa(i+1), param)
	}

	return str
}

var templateRegex = regexp.MustCompile(`\{\{([a-zA-Z0-9._-]+)}}`)

// TrTemplate translates a template string.
// Template strings are strings that contain placeholders with translation keys.
// Example template string: "{{example.term_a}} + {{example.term_b}} = {{example.sum}}"
func (l Locale) TrTemplate(template string) string {
	matches := templateRegex.FindAllStringSubmatch(template, -1)

	for _, match := range matches {
		key := match[1]
		template = strings.ReplaceAll(template, match[0], l.Tr(key))
	}

	return template
}

// TrError translates an error message.
// It is shorthand for `TrTemplate(err.Error())`.
func (l Locale) TrError(err error) string {
	return l.TrTemplate(err.Error())
}
