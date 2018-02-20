package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/boutros/marc"
)

type ResoRecord struct {
	Type            string   `json:"type"`
	Identifier      string   `json:"identifier"`
	AltIdentifier   []string `json:"alt-identifier"`
	OldIdentifier   []string `json:"old-identifier"`
	Heading         []string `json:"heading"`
	AltHeading      []string `json:"alt-heading"`
	WestCoordinate  []string `json:"west-coordinate"`
	EastCoordinate  []string `json:"east-coordinate"`
	NorthCoordinate []string `json:"north-coordinate"`
	SouthCoordinate []string `json:"south-coordinate"`
	MARCGeoCode     []string `json:"marc-geo-code"`
	Classification  []string `json:"classification"`
	GeneralNote     []string `json:"general-note"`
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
		output = flag.String("output", "", "output directory for JSON records")
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

	marcFile, err := os.Open(flag.Args()[0])
	checkError(err)
	records := marc.NewDecoder(marcFile, marc.MARC)
	for rec, err := records.Decode(); err != io.EOF; rec, err = records.Decode() {
		checkError(err)

		id, ok := rec.GetCField("001")
		if ok {
			recPath := fmt.Sprintf("%v/%v.json", *output, strings.Replace(id.Value, " ", "", -1))
			out, err := os.Create(recPath)
			checkError(err)
			defer out.Close()

			var NewRecord ResoRecord

			NewRecord.Identifier = strings.Replace(id.Value, " ", "", -1)
			NewRecord.AltIdentifier = getFields(rec, []string{"010"}, "a", []string{})
			NewRecord.OldIdentifier = getFields(rec, []string{"010"}, "z", []string{})
			NewRecord.Heading = getFields(rec, []string{"100", "110", "111", "130", "150", "151", "155", "180", "181", "182", "185"}, "", []string{})
			NewRecord.AltHeading = getFields(rec, []string{"400", "500", "410", "510", "411", "430", "530", "450", "550", "451", "551", "455", "555", "480", "580", "581", "781", "482", "485", "585"}, "", []string{"w", "5"})
			NewRecord.WestCoordinate = getFields(rec, []string{"034"}, "d", []string{})
			NewRecord.EastCoordinate = getFields(rec, []string{"034"}, "e", []string{})
			NewRecord.NorthCoordinate = getFields(rec, []string{"034"}, "f", []string{})
			NewRecord.SouthCoordinate = getFields(rec, []string{"034"}, "g", []string{})
			NewRecord.MARCGeoCode = getFields(rec, []string{"043"}, "a", []string{})
			NewRecord.Classification = getFields(rec, []string{"050", "053", "072", "073"}, "", []string{})
			NewRecord.GeneralNote = getFields(rec, []string{"680"}, "", []string{})

			b, err := json.Marshal(NewRecord)
			checkError(err)
			_, err = out.Write(b)
			checkError(err)
			out.Close()
		} else {
			panic("Error retrieving identifier for record.")
		}
	}
}

func getFields(rec *marc.Record, tags []string, subfield string, ignores []string) []string {
	var Values []string

	for _, tag := range tags {

		Field := rec.GetDFields(tag)

		if len(subfield) > 0 {
			if len(Field) > 0 {
				for i := 0; i < len(Field); i++ {
					s := string(Field[i].SubField(subfield))
					Values = append(Values, s)
				}
			}
		} else {
			if len(Field) > 0 && len(ignores) == 0 {
				for i := 0; i < len(Field); i++ {
					s := ""
					for _, subf := range Field[i].SubFields {
						s += subf.Value + ". "
					}
					Values = append(Values, s)
				}
			} else {
				for i := 0; i < len(Field); i++ {
					s := ""
					for _, subf := range Field[i].SubFields {
						if !stringInSlice(subf.Code, ignores) {
							s += subf.Value + ". "
						}
					}
					Values = append(Values, s)
				}
			}
		}

	}

	return Values
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
