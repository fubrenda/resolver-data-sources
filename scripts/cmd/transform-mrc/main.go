package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/boutros/marc"
)

type Mapping struct {
	Type            string  `json:"type"`
	Identifier      string  `json:"identifier"`
	AltIdentifier   []Field `json:"alt-identifier"`
	OldIdentifier   []Field `json:"old-identifier"`
	Heading         []Field `json:"heading"`
	AltHeading      []Field `json:"alt-heading"`
	WestCoordinate  []Field `json:"west-coordinate"`
	EastCoordinate  []Field `json:"east-coordinate"`
	NorthCoordinate []Field `json:"north-coordinate"`
	SouthCoordinate []Field `json:"south-coordinate"`
	MARCGeoCode     []Field `json:"marc-geo-code"`
	Classification  []Field `json:"classification"`
	GeneralNote     []Field `json:"general-note"`
}

type Field struct {
	Tag      string   `json:"tag"`
	Subfield []string `json:"subfield"`
	Type     string   `json:"type"`
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("transform-mrc: ")
	var (
		output       = flag.String("output", "", "output directory for JSON records")
		mappingInput = flag.String("mapping", "", "JSON mapping file for transform")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: transform-mrc [options...] file\n\nOptions:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	mappingFile, err := os.Open(*mappingInput)
	checkError(err)
	byteValue, _ := ioutil.ReadAll(mappingFile)
	json.Unmarshal(byteValue, &mapping)

	mapping.Identifier

	marcFile, err := os.Open(flag.Args()[0])
	checkError(err)

	records := marc.NewDecoder(marcFile, marc.MARC)
	log.Print(records)
	log.Print(*output)
	// for rec, err := records.Decode(); err != io.EOF; rec, err = records.Decode() {
	// 	checkError(err)
	//
	// 	id, ok := rec.GetCField("001")
	// 	if ok {
	//
	// 		recPath := fmt.Sprintf("%v/%v.json", *output, strings.Replace(id.Value, " ", "", -1))
	// 		log.Print(recPath)
	// 		out, err := os.Create(recPath)
	// 		checkError(err)
	//
	// 		defer out.Close()
	//
	// 		rec.DumpTo(out, false)
	// 		checkError(err)
	//
	// 		out.Close()
	// 	} else {
	// 		panic("Error retrieving identifier for record.")
	// 	}
	// }
}

func toJson() {

}
