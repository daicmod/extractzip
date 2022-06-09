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

func TestExtractFromu7z(t *testing.T) {
	err := ExtractFrom7z("wkhtmltopdf.exe", "./wkhtmltox-0.12.6-1.mxe-cross-win64.7z", "./")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	_, err = os.Stat("wkhtmltopdf.exe")
	if err != nil {
		t.Errorf("Not found wkhtmltopdf.exe: %v", err)
	}

}
