package main

import (
	"testing"
	"fmt"
)

// Check that we have English and Francais available
func TestGreeter(t *testing.T) {
	languages := []string{"English", "Francais"}

	for _, lang := range languages {
		l := NewGreeter(lang)
		if l == nil {
			t.Error("Language:", lang, "not implemented!")
		} else {
			if l.Language() != lang {
				t.Error("Language:", lang, "not correctly returned")
			}
		}
	}
}

// Uses the english Greeter to print some greetings
func ExampleEnglish() {
	english := NewGreeter("english")
	english.Hello()
	english.Greet("lcav")
	// Output:
	// 1 Hello world
	// 2 Hello lcav
}

// Uses the french Greeter to print some greetings
func ExampleFrancais() {
	francais := NewGreeter("francais")
	if francais == nil {
		fmt.Println("Francais not implemented")
		return
	}
	francais.Hello()
	francais.Greet("lcav")
	// Output:
	// 1 Bonjour le monde
	// 2 Bonjour monsieur/madame lcav
}