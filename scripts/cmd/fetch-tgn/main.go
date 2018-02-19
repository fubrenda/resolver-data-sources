package main

import (
	"io"
	"log"
	"os"

	"github.com/boutros/marc"
)

func main() {
	marcFile, err := os.Open("../data/oclcfast/FAST.mrc")
	if err != nil {
		log.Fatal(err)
	}

	dec := marc.NewDecoder(marcFile, marc.MARC)
	log.Print(dec)
	for rec, err := dec.Decode(); err != io.EOF; rec, err = dec.Decode() {
		id, ok := rec.GetCField("001")
		if ok {
			log.Print(id)
		}
	}
}
