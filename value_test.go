package json

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	exampleJson = `
{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    }
}`
)

func TestNewJsonValue(t *testing.T) {
	root := NewJsonValue([]byte(exampleJson))
	get := root.Get("glossary", "title")
	glossaryTitle := get.String()
	require.Equal(t, "example glossary", glossaryTitle)

	get = root.Get("glossary", "GlossDiv", "GlossList", "GlossEntry", "GlossDef", "GlossSeeAlso")
	slice := get.AsSlice()
	require.Equal(t, "GML", slice[0].String())
	require.Equal(t, "XML", slice[1].String())
}
