package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/boutros/marc"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	marcFile, err := os.Open("../data/oclcfast/marc/FAST.mrc")
	if err != nil {
		log.Fatal(err)
	}

	records := marc.NewDecoder(marcFile, marc.MARC)
	for rec, err := records.Decode(); err != io.EOF; rec, err = records.Decode() {
		checkError(err)

		id, ok := rec.GetCField("001")
		if ok {
			recPath := fmt.Sprintf("../derived-data/fast-records/%v.mrc", id.Value)
			out, err := os.Create(recPath)
			checkError(err)

			defer out.Close()

			rec.DumpTo(out, false)
			checkError(err)

			out.Close()
		} else {
			panic("Error retrieving 001 for record.")
		}
	}
}
