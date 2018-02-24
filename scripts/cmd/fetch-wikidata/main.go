package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knakk/sparql"
)

type ResoRecord struct {
	Identifier       string   `json:"identifier"`
	Heading          []string `json:"heading"`
	LCSHIdentifier   []string `json:"lcshIdentifier"`
	VIAFIdentifier   []string `json:"viafIdentifier"`
	LCMARCIdentifier []string `json:"lcmarcIdentifier"`
	AATIdentifier    []string `json:"aatIdentifier"`
	MESHIdentifier   []string `json:"meshIdentifier"`
}

type ObjSet struct {
	set map[string]bool
}

func NewObjSet() *ObjSet {
	return &ObjSet{make(map[string]bool)}
}

func (set *ObjSet) Add(i string) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found
}

func (set *ObjSet) Get(i string) bool {
	_, found := set.set[i]
	return found
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("fetch-wikidata: ")
	var (
		output = flag.String("output", "", "output directory for JSON records")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: fetch-wikidata [options...] Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	lcsh := "http://www.wikidata.org/prop/direct/P244"
	viaf := "http://www.wikidata.org/prop/direct/P214"
	lcmarc := "http://www.wikidata.org/prop/direct/P4801"
	aat := "http://www.wikidata.org/prop/direct/P1014"
	mesh := "http://www.wikidata.org/prop/direct/P486"
	labelQ := "SELECT ?heading WHERE { <%s> rdfs:label ?heading . }"

	repo, err := sparql.NewRepo("https://query.wikidata.org/sparql")
	checkError(err)

	res, err := repo.Query("SELECT ?item ?p ?o WHERE { VALUES (?p) { (wdt:P244) (wdt:P214) (wdt:P4801) (wdt:P1014) (wdt:P486) } ?item ?p ?o . }")
	checkError(err)

	URIsSet := NewObjSet()
	URIs := []string{}

	for k := range res.Results.Bindings {
		item := res.Results.Bindings[k]["item"].Value
		if len(item) != 0 {
			if URIsSet.Get(item) == false {
				URIsSet.Add(item)
				URIs = append(URIs, item)
			}
		}
	}

	for _, URI := range URIs {
		var WikiRecord ResoRecord
		WikiRecord.Identifier = URI
		Headings := []string{}
		LCSHes := []string{}
		VIAFes := []string{}
		LCMARCes := []string{}
		AATes := []string{}
		MESHes := []string{}

		for k := range res.Results.Bindings {
			item := res.Results.Bindings[k]["item"].Value
			if item == URI {
				pred := res.Results.Bindings[k]["p"].Value
				switch pred {
				case lcsh:
					LCSHes = append(LCSHes, res.Results.Bindings[k]["o"].Value)
				case viaf:
					VIAFes = append(VIAFes, res.Results.Bindings[k]["o"].Value)
				case lcmarc:
					LCMARCes = append(LCMARCes, res.Results.Bindings[k]["o"].Value)
				case aat:
					AATes = append(AATes, res.Results.Bindings[k]["o"].Value)
				case mesh:
					MESHes = append(MESHes, res.Results.Bindings[k]["o"].Value)
				}
				if len(Headings) == 0 {
					labelRes, err := repo.Query(fmt.Sprintf(labelQ, item))
					checkError(err)
					for k := range labelRes.Results.Bindings {
						label := labelRes.Results.Bindings[k]["heading"].Value
						Headings = append(Headings, label)
					}
				}
			}
		}
		WikiRecord.Heading = Headings
		WikiRecord.AATIdentifier = AATes
		WikiRecord.LCMARCIdentifier = LCMARCes
		WikiRecord.LCSHIdentifier = LCSHes
		WikiRecord.VIAFIdentifier = VIAFes
		WikiRecord.MESHIdentifier = MESHes

		recPath := fmt.Sprintf("%v/%v.json", *output, strings.Replace(URI, "http://www.wikidata.org/entity/", "", -1))
		out, err := os.Create(recPath)
		checkError(err)
		defer out.Close()
		b, err := json.Marshal(WikiRecord)
		checkError(err)
		_, err = out.Write(b)
		checkError(err)
		out.Close()
	}
}
