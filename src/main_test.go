package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting("Felipe")
	expected := "<b>Felipe<b>"

	if got != expected {
		t.Errorf("Error on greeting:\ngot %s\nexpected %s", got, expected)
	}

	got = greeting("CodeEdu")
	expected = "<b>CodeEdu<b>"

	if got != expected {
		t.Errorf("Error on greeting:\ngot %s\nexpected %s", got, expected)
	}
}
