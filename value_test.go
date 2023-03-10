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
	_, err := NewJsonValue([]byte(exampleJson))

	require.NoError(t, err)
}

func TestJsonGet(t *testing.T) {
	root, _ := NewJsonValue([]byte(exampleJson))

	_, err := root.Get("glossary", "title")

	require.NoError(t, err)
}

func TestJsonAsSlice(t *testing.T) {
	root, _ := NewJsonValue([]byte(exampleJson))
	value, _ := root.Get("glossary", "GlossDiv", "GlossList", "GlossEntry", "GlossDef", "GlossSeeAlso")

	_, err := value.AsSlice()

	require.NoError(t, err)
}

func TestString(t *testing.T) {
	root, _ := NewJsonValue([]byte(exampleJson))
	value, _ := root.Get("glossary", "GlossDiv", "GlossList", "GlossEntry", "GlossDef", "GlossSeeAlso")
	slice, _ := value.AsSlice()

	require.Equal(t, "GML", slice[0].String())
	require.Equal(t, "XML", slice[1].String())
}
