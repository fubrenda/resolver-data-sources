# Subjects Fuzzy Resolver Data Sources

## Guiding user story:

As a developer at a metadata aggregator, I would like to have an API to issue batch queries with a non-normalized subject heading / tag value string and then receive an identifier, primary label, & string-matching accuracy score for that string’s concept record in a controlled vocabulary.

## Data sources:

| Dataset | Source | Representations | Licenses / Usage Guidelines |
| --- | --- | --- | --- |
| LCSH | Subscription Service with updates feed | MARC (ISO or XML) | covered by federal output / cannot license? |
| LCSH | id.loc.gov downloads | SKOS, MADS/RDF, JSON | covered by federal output / cannot license? |
| LCSH | Classification Schema (partial) | CSV | covered by federal output / cannot license? |
| OCLC FAST (where relevant type) | https://www.oclc.org/research/themes/data-science/fast/download.html | MARC, RDF (using Schema.org) | Open Data Commons Attribution License (ODC-By): http://www.oclc.org/research/activities/fast/odcby.htm. |
| Getty AAT | http://www.getty.edu/research/tools/vocabularies/obtain/download.html | JSON, RDF (local schema) | Open Data Commons Attribution License (ODC-By) 1.0 |
| Getty TGN | http://www.getty.edu/research/tools/vocabularies/obtain/download.html | JSON, RDF (local schema) | Open Data Commons Attribution License (ODC-By) 1.0 |
| Wikidata (where pre-matched & relevant type) | See access points at https://www.wikidata.org/wiki/Wikidata:Data_access | local schema RDF & JSON | CC-0 |


## Attributions

"This codebase and git repository contains information from FAST (Faceted Application of Subject Terminology) Data which is made available by OCLC Online Computer Library Center, Inc. under the [ODC Attribution License](https://www.oclc.org/research/themes/data-science/fast/odcby.html)."

"This codebase and git repository contains information from the Art & Architecture Thesaurus® (AAT), the Union List of Artists Names® (ULAN), the Getty Thesaurus of Geographic Names® (TGN), the Cultural Objects Name Authority® (CONA), and the Getty Iconography Authority (IA) are made available by the J. Paul Getty Trust under the [Open Data Commons Attribution License](https://www.oclc.org/research/themes/data-science/fast/odcby.html). 1.0. For questions, send an email to vocab@getty.edu."

"This codebase and git repository contains information [powered by Wikidata data](https://www.wikidata.org/wiki/Wikidata:Data_access)."
