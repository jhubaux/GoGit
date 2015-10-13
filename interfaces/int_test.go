package main

import (
	"fmt"
	"testing"
)

// Check that we have English and Francais available
func TestGreeter(t *testing.T) {
	greeters := []Greeter{English{}, Francais{}}

	for _, g := range greeters {
		if g == nil {
			t.Error("Language:", lang, "not implemented!")
		} else {
			if g.Language() != lang {
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
