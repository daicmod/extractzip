package extractzip

import (
	"os"
	"testing"
)

func TestExtractFromZip(t *testing.T) {
	err := ExtractFromZip("sample.txt", "./sample.zip", "./")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	_, err = os.Stat("sample.txt")
	if err != nil {
		t.Errorf("Not found sample.txt: %v", err)
	}

}
