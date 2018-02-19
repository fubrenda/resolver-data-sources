package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dutchcoders/goftp"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var ftp *goftp.FTP
	var inFile = "/pub/researchdata/fast/FASTAll.marc.zip"
	var outFile = "../data/oclcfast/fast.marc.zip"
	var outDir = "../data/oclcfast/marc/"

	ftp, err := goftp.Connect("anonftp.oclc.org:21")
	checkError(err)
	defer ftp.Close()

	if err = ftp.Login("anonymous", "anonymous"); err != nil {
		panic(err)
	}

	_, err = ftp.Retr(inFile, func(resp io.Reader) error {
		out, err := os.Create(outFile)
		checkError(err)
		defer out.Close()

		_, err = io.Copy(out, resp)
		checkError(err)
		out.Close()

		return err
	})

	ftp.Close()

	files, err := Unzip(outFile, outDir)
	checkError(err)

	fmt.Println("Unzipped: " + strings.Join(files, ", "))
}

// Unzip will un-compress a zip archive,
// moving all files and folders to an output directory
func Unzip(src, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			var fdir string
			if lastIndex := strings.LastIndex(fpath, string(os.PathSeparator)); lastIndex > -1 {
				fdir = fpath[:lastIndex]
			}

			err = os.MkdirAll(fdir, os.ModePerm)
			if err != nil {
				log.Fatal(err)
				return filenames, err
			}
			f, err := os.OpenFile(
				fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}
