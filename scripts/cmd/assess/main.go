package main

import (
	"flag"
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
	log.SetFlags(0)
	log.SetPrefix("assess: ")
	var (
		cfield = flag.String("cfield", "", "control field to analyze")
		tag    = flag.String("tag", "", "data field to analyze")
		subtag = flag.String("subtag", "", "data field to analyze")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: assess [options...] file\n\nOptions:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	marcFile, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	records := marc.NewDecoder(marcFile, marc.MARC)
	for rec, err := records.Decode(); err != io.EOF; rec, err = records.Decode() {
		checkError(err)
		if len(*cfield) > 0 {
			fieldVal, ok := rec.GetCField(*cfield)
			if ok {
				fmt.Print(fieldVal.Value + "\n")
			}
		} else if len(*tag) > 0 {
			if len(*subtag) > 0 {
				fieldVal := rec.GetDFields(*tag)
				if len(fieldVal) > 0 {
					for i := 0; i < len(fieldVal); i++ {
						fmt.Print(fieldVal[i].SubField(*subtag) + "\n")
					}
				}
			} else {
				fieldVal := rec.GetDFields(*tag)
				if len(fieldVal) > 0 {
					for i := 0; i < len(fieldVal); i++ {
						for j := 0; j < len(fieldVal[i].SubFields); j++ {
							fmt.Print(fieldVal[i].SubFields[j].Code + ":\t" + fieldVal[i].SubFields[j].Value + "\n")
						}
					}
				}
			}
		}
	}
}
