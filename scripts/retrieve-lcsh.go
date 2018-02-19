package main

import (
	"compress/gzip"
	"io"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var outFile = "../data/lcsh/lcsh2014.mrc.gz"
	var outExpandedFile = "../data/lcsh/lcsh2014.mrc"

	resp, err := http.Get("https://www.loc.gov/cds/downloads/MDSConnect/Subjects.2014.utf8.gz")
	checkError(err)
	defer resp.Body.Close()

	// Write out gzipped file
	out, err := os.Create(outFile)
	checkError(err)
	defer out.Close()
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	checkError(err)
	out.Close()
	defer resp.Body.Close()

	// Expand gzipped file in given directory
	ungzipout, err := os.Create(outExpandedFile)
	checkError(err)
	defer ungzipout.Close()
	defer resp.Body.Close()

	r, err := gzip.NewReader(resp.Body)
	checkError(err)

	io.Copy(ungzipout, r)
	r.Close()
	ungzipout.Close()
	resp.Body.Close()
}
