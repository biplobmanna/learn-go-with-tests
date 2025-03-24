package hello

import "testing"

// function name must be Test*
// must have (t *testing.T) are functional params
func TestHello(t *testing.T) {
	// anon-functions which also is a (t *testing.T) saying that it's a testing function
	t.Run("saying hello to people, in English", func(t *testing.T) {
		got := Hello("Marco", "English")
		want := "Hello, Marco"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		got := Hello("Marco", "Spanish")
		want := "Hola, Marco"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to people in French", func(t *testing.T) {
		got := Hello("Marco", "French")
		want := "Bonjour, Marco"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'hello world' in English if no people specified, or language specified", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

// refcatoring the test code repeats as well
// testing.TB is an interface over both testing.T, and testing.B\
// that way, no need to worry about tests or benchmarks
func assertCorrectMessage(t testing.TB, got, want string) {
	// this tells the test suite that this method is a helper
	// this way, when it fails, the line-number reported will be
	// the actual function, instead of this helper function
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
