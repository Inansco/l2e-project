package main

import "testing"

func TestHex(t *testing.T) {
	got := Transform("1E (hex) files were added")
	want := "30 files were added"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBin(t *testing.T) {
	got := Transform("It has been 10 (bin) years")
	want := "It has been 2 years"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestUpper(t *testing.T) {
	got := Transform("go (up)")
	want := "GO"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestLower(t *testing.T) {
	got := Transform("STOP (low)")
	want := "stop"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCap(t *testing.T) {
	got := Transform("hello (cap)")
	want := "Hello"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestMultiCase(t *testing.T) {
	got := Transform("this is so exciting (up, 2)")
	want := "this is SO EXCITING"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestPunctuation(t *testing.T) {
	got := Transform("Hello , world !")
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestQuotes(t *testing.T) {
	got := Transform("I am ' awesome '")
	want := "I am 'awesome'"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAtoAn(t *testing.T) {
	got := Transform("a apple")
	want := "an apple"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
