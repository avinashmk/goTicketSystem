package console

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func ExamplePrompt() {
	Prompt("Hello World!")
	// Output:
	// Hello World!
}

func TestGetString(t *testing.T) {
	want := "World!"
	close := SetStdin(want)
	defer close()
	if got := GetString("Hello"); got != want {
		t.Errorf("userInput failed! \nwant:%s \ngot:%s", want, got)
	}
}

func SetStdin(want string) func() {
	content := []byte(want + "\n")
	tmpfile, err := ioutil.TempFile("", "tempfile")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	os.Stdin = tmpfile

	return func() {
		defer func() { os.Stdin = oldStdin }()
		defer os.Remove(tmpfile.Name()) // clean up
		if err := tmpfile.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
