package extractzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/bodgit/sevenzip"
)

func ExtractFromZip(target, src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	destFile := filepath.Join(dest, target)
	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, f := range r.File {
		if filepath.Base(f.Name) == target {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			if _, err := io.Copy(out, rc); err != nil {
				return err
			}

			break
		}
	}

	return nil
}

func ExtractFrom7z(target, src, dest string) error {
	r, err := sevenzip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	destFile := filepath.Join(dest, target)
	out, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer out.Close()

	for _, f := range r.File {
		if filepath.Base(f.Name) == target {
			rc, err := f.Open()
			if err != nil {
				return err
			}

			if _, err := io.Copy(out, rc); err != nil {
				return err
			}
			rc.Close()

			break
		}
	}

	return nil
}
