package dotenv

import (
	"os"
	"testing"
)

func TestExtractQuottedString(t *testing.T) {
	ext := hasQuotePrefix(`"hacks"`)
	if !ext {
		t.Errorf(`Wrong "hacks"`)
	}
}

func TestLoad(t *testing.T) {

	Load() // use dotenv.Load()

	if _, exist := os.LookupEnv("HELLO_WORLD"); !exist {
		t.Errorf("Environment variable HELLO_WORLD is not exist!")
		return

	}
	if _, exist := os.LookupEnv("QUODET"); !exist {
		t.Errorf("Environment variable HELLO_WORLD is not exist!")
		return
	}

	t.Log(os.Getenv("QUODET"))
	t.Log(os.Getenv("HELLO_WORLD"))
}
