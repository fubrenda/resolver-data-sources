package main

import (
	"log"

	"github.com/knakk/sparql"
)

func main() {
	repo, err := sparql.NewRepo("https://query.wikidata.org/sparql")
	if err != nil {
		log.Fatal(err)
	}

	res, err := repo.Query("SELECT * WHERE { VALUES ?p { wdt:P244 wdt:P214 wdt:P4801 wdt:P1014 wdt:P486 } ?s ?p ?o. }")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res.Results)
}
