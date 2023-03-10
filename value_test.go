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

func TestJsonNew(t *testing.T) {
	root := NewJsonValue([]byte(exampleJson))

	require.NoError(t, root.Err())
}

func TestJsonGet(t *testing.T) {
	root := NewJsonValue([]byte(exampleJson))

	value := root.Get("glossary", "title")

	require.NoError(t, value.Err())
}

func TestString(t *testing.T) {
	root := NewJsonValue([]byte(exampleJson))
	value := root.Get("glossary", "GlossDiv", "GlossList", "GlossEntry", "GlossDef", "GlossSeeAlso").Slice()

	require.Equal(t, "GML", value[0].String())
	require.Equal(t, "XML", value[1].String())
}
