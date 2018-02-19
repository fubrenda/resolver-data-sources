# Subjects Fuzzy Resolver Data Sources

## Guiding user story:

As a developer at a metadata aggregator, I would like to have an API to issue batch queries with a non-normalized subject heading / tag value string and then receive an identifier, primary label, & string-matching accuracy score for that string’s concept record in a controlled vocabulary.

## Data sources:

| Dataset | Source | Representations | Licenses / Usage Guidelines |
| --- | --- | --- | --- |
| LCSH | [Subscription Service (free version is from 2014)](https://www.loc.gov/cds/downloads/MDSConnect/Subjects.2014.utf8.gz) | MARC (ISO or XML) | covered by federal output / cannot license? There is subscription fee for full, up-to-date set. |
| LCSH | id.loc.gov downloads | SKOS, MADS/RDF, JSON | covered by federal output / cannot license? |
| LCSH | Classification Schema (partial) | CSV | covered by federal output / cannot license? |
| OCLC FAST (where relevant type) | ftp://anonftp.oclc.org/pub/researchdata/fast/FASTAll.marcxml.zip | MARC, RDF (using Schema.org) | Open Data Commons Attribution License (ODC-By): http://www.oclc.org/research/activities/fast/odcby.htm. |
| Getty AAT | http://vocab.getty.edu/dataset/aat/full.zip | JSON, RDF (local schema) | Open Data Commons Attribution License (ODC-By) 1.0 |
| Getty TGN | http://vocab.getty.edu/dataset/tgn/full.zip | JSON, RDF (local schema) | Open Data Commons Attribution License (ODC-By) 1.0 |
| Wikidata (where pre-matched & relevant type) | See access points at https://www.wikidata.org/wiki/Wikidata:Data_access or use [this SPARQL Endpoint Query](https://query.wikidata.org/sparql?query=SELECT%20*%20%0AWHERE%20%7B%0A%20%20%20%20VALUES%20%3Fp%20%7B%20wdt%3AP244%20wdt%3AP214%20wdt%3AP4801%20wdt%3AP1014%20wdt%3AP486%20%7D%20%0A%20%20%20%20%3Fs%20%3Fp%20%3Fo.%0A%7D) | local schema as JSON | CC-0 |


## Attributions

"This codebase and git repository contains information from FAST (Faceted Application of Subject Terminology) Data which is made available by OCLC Online Computer Library Center, Inc. under the [ODC Attribution License](https://www.oclc.org/research/themes/data-science/fast/odcby.html)."

"This codebase and git repository contains information from the Art & Architecture Thesaurus® (AAT), the Union List of Artists Names® (ULAN), the Getty Thesaurus of Geographic Names® (TGN), the Cultural Objects Name Authority® (CONA), and the Getty Iconography Authority (IA) are made available by the J. Paul Getty Trust under the [Open Data Commons Attribution License](https://www.oclc.org/research/themes/data-science/fast/odcby.html). 1.0. For questions, send an email to vocab@getty.edu."

"This codebase and git repository contains information [powered by Wikidata data](https://www.wikidata.org/wiki/Wikidata:Data_access)."
