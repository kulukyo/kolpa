/*
kolpa is a fake data generator written in and for Go.

Usage:
	k := kolpa.C()
	k.Name()
	k.FirstName()
	k.Address()
*/

package kolpa

import (
	"strings"
)

// Generator struct to access various generator functions
type Generator struct {
	Locale string
}

// Creator function, initiates kolpa with or without locale
// setting. The default locale setting is "en_US".
// Returns a generator type that will be used to call generator methods.
func C(localeVar ...string) Generator {
	newGenerator := Generator{}
	if len(localeVar) > 0 {
		newGenerator.Locale = localeVar[0]
	} else {
		newGenerator.Locale = "en_US"
	}
	// newGenerator.populateFunctions()

	return newGenerator
}

// Language setter function. Language setting change be changed
// anytime by using this function.
func (g *Generator) SetLanguage(localeVar string) {
	g.Locale = localeVar
}


// Generic generator function.
// Recursively generates data for intended data type.
// intended variable should be slug version of a valid data type.
func (g *Generator) GenericGenerator(intended string) string {
	var result string
	var slice []string

	if len(g.fileToSlice(intended)) == 0 {
		slice = g.fileToSliceAll(intended)
	} else {
		slice = g.fileToSlice(intended)
	}
	
	if len(slice) == 0 {
		words := strings.Split(intended, "_")
 		intendedNew := strings.Join(words[:len(words)-1], "_")
 		slice = g.fileToSlice(intendedNew)
 		if len(slice) == 0 {
 			return ""
 		}
	} 

	if g.isParseable(slice) {
		randomFormat := getRandom(slice)
		tokens := g.formatToSlice(randomFormat)
	
		randomItems := make(map[string]string)

		for _, token := range tokens {
			randomItems[token] = g.GenericGenerator(token)
		}

		result = g.parser(randomFormat, randomItems)
	} else {
		result = getRandom(slice)
	}
	
	return result
}
