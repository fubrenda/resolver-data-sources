# OCLC FAST


## Mapped Fields for Resolver


## MARC Data Analysis (Subset)

tag       | field info | note
--------- | ---------- | ----
001       | OCLC FAST identifier |
003       | Do not map. |
005       | Do not map.
008       | Do not map. |
016 $a    | Keep if $a doesn't match 001 |
016 $z    | Keep if $z doesn't match 001 |
016 $2    | Keep $a if this is not OCoLC |
034 $d    | Westernmost Coordinate |
034 $e    | Easternmost Coordinate |
034 $f    | Northernmost Coordinate |
034 $g    | Southernmost Coordinate |
040       | Do not map. |
043 $a    | Geographic Area Code | Keep? Will need further parsing to be useable from MARC form.
046 $s    | Start period (of an event) |
046 $g    | Death date (of an event?) |
046 $t    | End period (of an event) |
046 $v    | Do not map. |
046 $u    | Do not map. | URI for a date?
046 $2    | Do not map. | Date scheme for field, possibly useful later.
053 $abc  | Classification number. | Useful sometimes for mapping?
111       | Meeting name. | Primary matching term.
148       | Chronological Term | Primary matching term.
150       | Topic Term | Primary matching term.
155       | Form / Genre Term | Primary matching term.
260       | Do not map. |
360       | Do not map. |
368 $a    | Type of Meeting / Event Body | Secondary matching term.
368 $c    | Other Meeting / Event Info | Secondary matching term.
368 $2    | Do not map. |
370 $c    | Associated Country with an Event. | Secondary matching term.
370 $e    | Headquarters for an Event. | Secondary matching term.
370 $f    | Associated Location for an Event. | Secondary matching term.
370 $g    | Place of Origin for an Event. | Secondary matching term.
370 rest  | Do not map. |
372 $a    | Field of Activity for an Event. | Secondary matching term.
372 $v    | Do not map. |
372 $2    | Do not map. |
373 $a    | Associated Group for an Event. | Secondary matching term.
373 $s    | Do not map. |
373 $v    | Do not map. |
373 $u    | Do not map. |
373 $2    | Do not map. |
377 $a    | Associated Language. | Do not use now?
400       | Alternate Form (technically, tracing title) of Personal Name | Secondary matching term.
400 $w    | Do not map.  | Ignore for now.
410       | Alternate Form (technically, tracing title) of Corporate Term | Secondary matching term.
410 $w    | Do not map.  | Ignore for now.
411       | Alternate Form (technically, tracing title) of Meeting Name | Secondary matching term.
411 $w    | Do not map.  | Ignore for now.
430       | Alternate Form (technically, tracing title) of Uniform Title | Secondary matching term.
430 $w    | Do not map.  | Ignore for now.
445       | Additional Form | Secondary matching term.
448       | Alternate Form (technically, tracing title) of Chronological Term | Secondary matching term.
448 $w    | Do not map.  | Ignore for now.
450       | Alternate Form (technically, tracing title) of Topic Term | Secondary matching term.
450 $w    | Do not map.  | Ignore for now.
451       | Alternate Form (technically, tracing title) of Geographic Term | Secondary matching term.
451 $w    | Do not map.  | Ignore for now.
455       | Alternate Form (technically, tracing title) of Form Genre Term | Secondary matching term.
455 $w    | Do not map.  | Ignore for now.
485 $v    | Alternate Form (technically, tracing title) of Form Subdivision Term | Secondary matching term.
485 $w$5  | Do not map.  | Ignore for now.
500       | Alternate Form (technically, See Also) of Personal Name | Secondary matching term.
510       | Alternate Form (technically, See Also) of Corporate Name | Secondary matching term.
510 $0    | Do not map. | Ignore for now.
510 $w    | Do not map. | Ignore for now.
511       | Alternate Form (technically, See Also) of Meeting Name | Secondary matching term.
511 $0    | Do not map. | Ignore for now.
511 $w    | Do not map. | Ignore for now.
530       | Alternate Form (technically, See Also) of Uniform Title | Secondary matching term.
530 $0    | Do not map. | Ignore for now.
530 $w    | Do not map. | Ignore for now.
548       | Alternate Form (technically, See Also) of Chronological Term | Secondary matching term.
548 $2    | Do not map. |
550       | Alternate Form (technically, See Also) of Topic Term | Secondary matching term.
550 $0    | Do not map. | Ignore for now.
550 $w    | Do not map. | Ignore for now.
551       | Alternate Form (technically, See Also) of Geographic Term | Secondary matching term.
551 $0    | Do not map. | Ignore for now.
551 $w    | Do not map. | Ignore for now.
555 $a$v  | Alternate Form (technically, See Also) of Form Genre Term | Secondary matching term.
555 $0    | Do not map. | Ignore for now.
670       | Do not map. |
680       | Public, general note. | Keep for general matching field?
688       | Do not map. | Usage stats.
700       | Linked Personal Name | Tertiary matching term dependent on type.
700 $0    | Do not map. | Reconciliation URI?
700 $w    | Do not map. | Ignore for now.
710       | Linked Corporate Name | Tertiary matching term dependent on type.
710 $0    | Do not map. | Reconciliation URI?
710 $w    | Do not map. | Ignore for now.
711       | Linked Meeting Name | Tertiary matching term dependent on type.
711 $0    | Do not map. | Reconciliation URI?
711 $w    | Do not map. | Ignore for now.
730       | Linked Uniform Title | Tertiary matching term dependent on type.
730 $0    | Do not map. | Reconciliation URI?
730 $w    | Do not map. | Ignore for now.
748       | Linked Chronological Term | Tertiary matching term dependent on type.
748 $0    | Do not map. | Reconciliation URI?
748 $w    | Do not map. | Ignore for now.
750       | Linked Topical Term | Tertiary matching term dependent on type.
750 $0    | Do not map. | Reconciliation URI?
750 $w    | Do not map. | Ignore for now.
751       | Linked Geographic Term | Tertiary matching term dependent on type.
751 $0    | Do not map. | Reconciliation URI?
751 $w    | Do not map. | Ignore for now.
755       | Linked GenreForm Term | Tertiary matching term dependent on type.
755 $0    | Do not map. | Reconciliation URI?
755 $w    | Do not map. | Ignore for now.
780       | Linked General Subdivision Term | Tertiary matching term. Map to Event or other term for context.
780 $0    | Do not map. | Ignore for now.
780 $w    | Do not map. | Ignore for now.
781       | Linked Geo Subdivision Term | Tertiary matching term. Map to Event or other term for context.
781 $0    | Do not map. | Ignore for now.
781 $w    | Do not map. | Ignore for now.
785       | Linked Form Subdivision Term | Tertiary matching term dependent on type.
785 $0    | Do not map. | Ignore for now.
785 $w    | Do not map. | Ignore for now.


## MARC Data Analysis (Raw)

### Chronological Fields

```
            001: |=========================|    676/676 | 100%
            003: |=========================|    676/676 | 100%
            005: |=========================|    676/676 | 100%
            008: |=========================|    676/676 | 100%
    016_7 _$a$2: |=========================|    676/676 | 100%
040_  _$a$b$c$f: |=========================|    676/676 | 100%
      148_  _$a: |=========================|    676/676 | 100%
      448_  _$a: |======================== |    675/676 |  99%
    448_  _$a$w: |                         |      3/676 |   0%
      688_  _$a: |=========================|    676/676 | 100%
```

### Event Fields

```
                    001: |=========================|  15046/15046 | 100%
                    003: |=========================|  15046/15046 | 100%
                    005: |=========================|  15046/15046 | 100%
                    008: |=========================|  15046/15046 | 100%
            016_7 _$a$2: |=======================  |  14203/15046 |  94%
          016_7 _$a$z$2: |=                        |    800/15046 |   5%
        016_7 _$a$z$z$2: |                         |     41/15046 |   0%
      016_7 _$a$z$z$z$2: |                         |      1/15046 |   0%
016_7 _$a$z$z$z$z$z$z$2: |                         |      1/15046 |   0%
      034_  _$d$e$f$g$2: |==                       |   1283/15046 |   8%
        040_  _$a$b$c$f: |=========================|  15046/15046 | 100%
              043_  _$a: |=====                    |   3518/15046 |  23%
              046_  _$s: |=======                  |   4314/15046 |  28%
            046_  _$s$2: |                         |      8/15046 |   0%
            046_  _$s$g: |                         |      1/15046 |   0%
            046_  _$s$t: |==                       |   1360/15046 |   9%
          046_  _$s$t$2: |                         |      2/15046 |   0%
          046_  _$s$v$u: |                         |      1/15046 |   0%
              046_  _$t: |                         |      5/15046 |   0%
            046_  _$t$2: |                         |      1/15046 |   0%
              053_  _$a: |                         |      9/15046 |   0%
            053_  _$a$b: |                         |      8/15046 |   0%
              053_ 0_$a: |==                       |   1281/15046 |   8%
            053_ 0_$a$b: |                         |    101/15046 |   0%
          053_ 0_$a$b$c: |                         |      4/15046 |   0%
            053_ 0_$a$c: |                         |     20/15046 |   0%
              111_  _$a: |                         |    579/15046 |   3%
            111_  _$a$x: |                         |     24/15046 |   0%
          111_  _$a$x$x: |                         |      1/15046 |   0%
            111_ 0_$a$x: |==                       |   1549/15046 |  10%
          111_ 0_$a$x$x: |                         |    106/15046 |   0%
        111_ 0_$a$x$x$x: |                         |      3/15046 |   0%
              111_0 _$a: |==                       |   1606/15046 |  10%
            111_0 _$a$x: |                         |      9/15046 |   0%
              111_2 _$a: |==================       |  11137/15046 |  74%
            111_2 _$a$d: |                         |      2/15046 |   0%
            111_2 _$a$x: |                         |     30/15046 |   0%
              368_  _$a: |                         |    104/15046 |   0%
            368_  _$a$2: |                         |    174/15046 |   1%
            368_  _$a$a: |                         |      3/15046 |   0%
          368_  _$a$a$2: |                         |      5/15046 |   0%
              368_  _$c: |                         |      8/15046 |   0%
            368_  _$c$2: |                         |      1/15046 |   0%
              370_  _$c: |                         |     11/15046 |   0%
            370_  _$c$2: |                         |      9/15046 |   0%
            370_  _$c$c: |                         |      2/15046 |   0%
          370_  _$c$c$2: |                         |      1/15046 |   0%
    370_  _$c$c$c$c$c$c: |                         |      1/15046 |   0%
        370_  _$c$c$e$2: |                         |      2/15046 |   0%
      370_  _$c$c$e$f$2: |                         |      1/15046 |   0%
            370_  _$c$e: |                         |     19/15046 |   0%
          370_  _$c$e$2: |                         |      9/15046 |   0%
          370_  _$c$e$e: |                         |      1/15046 |   0%
        370_  _$c$e$e$2: |                         |      1/15046 |   0%
            370_  _$c$f: |                         |      2/15046 |   0%
              370_  _$e: |                         |     78/15046 |   0%
            370_  _$e$2: |                         |     53/15046 |   0%
          370_  _$e$2$s: |                         |      1/15046 |   0%
          370_  _$e$c$2: |                         |      1/15046 |   0%
            370_  _$e$e: |                         |      2/15046 |   0%
        370_  _$e$e$e$e: |                         |      1/15046 |   0%
    370_  _$e$e$e$e$e$e: |                         |      1/15046 |   0%
          370_  _$e$s$t: |                         |      1/15046 |   0%
        370_  _$e$s$t$2: |                         |      1/15046 |   0%
          370_  _$e$v$u: |                         |      1/15046 |   0%
              370_  _$f: |                         |      9/15046 |   0%
            370_  _$f$2: |                         |      7/15046 |   0%
            370_  _$f$c: |                         |      1/15046 |   0%
        370_  _$f$f$f$2: |                         |      1/15046 |   0%
    370_  _$f$f$f$f$f$2: |                         |      1/15046 |   0%
          370_  _$f$s$t: |                         |      1/15046 |   0%
              370_  _$g: |                         |      1/15046 |   0%
              372_  _$a: |                         |     12/15046 |   0%
            372_  _$a$2: |                         |    120/15046 |   0%
          372_  _$a$2$v: |                         |      2/15046 |   0%
          372_  _$a$a$2: |                         |     41/15046 |   0%
          372_  _$a$a$a: |                         |      1/15046 |   0%
        372_  _$a$a$a$2: |                         |      7/15046 |   0%
    372_  _$a$a$a$a$a$2: |                         |      1/15046 |   0%
              373_  _$a: |                         |      4/15046 |   0%
            373_  _$a$2: |                         |     40/15046 |   0%
      373_  _$a$2$s$v$u: |                         |      1/15046 |   0%
          373_  _$a$a$2: |                         |      4/15046 |   0%
        373_  _$a$a$a$2: |                         |      1/15046 |   0%
            373_  _$a$s: |                         |      1/15046 |   0%
              377_  _$a: |                         |    112/15046 |   0%
            377_  _$a$a: |                         |     12/15046 |   0%
        377_  _$a$a$a$a: |                         |      1/15046 |   0%
  377_  _$a$a$a$a$a$a$a: |                         |      1/15046 |   0%
              410_  _$a: |                         |      1/15046 |   0%
            410_1 _$a$b: |                         |     19/15046 |   0%
          410_1 _$a$b$b: |                         |      4/15046 |   0%
        410_1 _$a$b$d$w: |                         |     31/15046 |   0%
          410_1 _$a$b$w: |                         |      1/15046 |   0%
              410_2 _$a: |                         |     23/15046 |   0%
            410_2 _$a$b: |                         |     37/15046 |   0%
          410_2 _$a$b$b: |                         |      1/15046 |   0%
        410_2 _$a$b$d$w: |                         |      1/15046 |   0%
          410_2 _$a$b$w: |                         |     15/15046 |   0%
            410_2 _$a$w: |                         |     17/15046 |   0%
          410_2 _$a$x$w: |                         |      1/15046 |   0%
              410_20_$a: |                         |      1/15046 |   0%
              411_  _$a: |==                       |   1211/15046 |   8%
            411_  _$a$5: |                         |      2/15046 |   0%
            411_  _$a$w: |                         |    565/15046 |   3%
          411_  _$a$x$w: |                         |     30/15046 |   0%
        411_  _$a$x$x$w: |                         |      1/15046 |   0%
              411_0 _$a: |=                        |    861/15046 |   5%
            411_0 _$a$w: |                         |     29/15046 |   0%
            411_0 _$a$x: |                         |      1/15046 |   0%
              411_1 _$a: |                         |     16/15046 |   0%
          411_1 _$a$b$b: |                         |      1/15046 |   0%
            411_1 _$a$w: |                         |      3/15046 |   0%
              411_2 _$a: |=====                    |   3299/15046 |  21%
          411_2 _$a$c$w: |                         |      2/15046 |   0%
            411_2 _$a$d: |                         |      4/15046 |   0%
          411_2 _$a$d$w: |                         |      2/15046 |   0%
            411_2 _$a$w: |                         |    165/15046 |   1%
          411_2 _$a$x$w: |                         |      1/15046 |   0%
              430_ 0_$a: |                         |      1/15046 |   0%
              450_  _$a: |                         |    307/15046 |   2%
            450_  _$a$5: |                         |      2/15046 |   0%
            450_  _$a$w: |                         |     86/15046 |   0%
            450_  _$a$x: |                         |      3/15046 |   0%
          450_  _$a$x$w: |                         |     46/15046 |   0%
              450_ 0_$a: |                         |      1/15046 |   0%
              451_  _$a: |                         |      1/15046 |   0%
          451_  _$a$x$w: |                         |      1/15046 |   0%
            451_  _$a$z: |                         |      2/15046 |   0%
          500_1 _$a$d$0: |                         |      2/15046 |   0%
            500_3 _$a$0: |                         |      3/15046 |   0%
            510_2 _$a$0: |                         |      2/15046 |   0%
          510_2 _$a$0$w: |                         |      1/15046 |   0%
              511_  _$a: |                         |      1/15046 |   0%
            511_  _$a$0: |=                        |    649/15046 |   4%
          511_  _$a$0$w: |=                        |    760/15046 |   5%
            511_  _$a$x: |                         |      5/15046 |   0%
          511_  _$a$x$0: |                         |      1/15046 |   0%
            511_0 _$a$0: |                         |    120/15046 |   0%
          511_0 _$a$0$w: |                         |    207/15046 |   1%
          511_0 _$a$x$0: |                         |      2/15046 |   0%
              511_2 _$a: |                         |      1/15046 |   0%
            511_2 _$a$0: |                         |    228/15046 |   1%
          511_2 _$a$0$w: |                         |     79/15046 |   0%
        511_2 _$a$2$0$5: |                         |      1/15046 |   0%
            548_  _$a$2: |                         |      1/15046 |   0%
            550_  _$a$0: |=                        |    947/15046 |   6%
          550_  _$a$0$w: |                         |    267/15046 |   1%
            550_  _$a$x: |                         |      1/15046 |   0%
          550_  _$a$x$0: |                         |    390/15046 |   2%
        550_  _$a$x$0$w: |                         |     67/15046 |   0%
        550_  _$a$x$x$0: |                         |      2/15046 |   0%
            551_  _$a$0: |                         |    267/15046 |   1%
          551_  _$a$0$w: |                         |     62/15046 |   0%
          551_  _$a$z$0: |==                       |   1527/15046 |  10%
        551_  _$a$z$0$w: |                         |     12/15046 |   0%
        551_  _$a$z$2$0: |                         |    349/15046 |   2%
        551_  _$a$z$z$0: |                         |      3/15046 |   0%
      551_  _$a$z$z$2$0: |                         |      2/15046 |   0%
          551_  _$w$a$0: |                         |      1/15046 |   0%
            555_  _$a$0: |                         |     27/15046 |   0%
            670_  _$a$b: |==                       |   1283/15046 |   8%
              688_  _$a: |====================     |  12614/15046 |  83%
    700_00_$a$b$c$d$x$0: |                         |      2/15046 |   0%
        700_10_$a$d$x$0: |                         |      1/15046 |   0%
      700_10_$a$q$d$x$0: |                         |      1/15046 |   0%
        710_10_$a$b$0$w: |                         |      1/15046 |   0%
      710_10_$a$b$b$x$0: |                         |      1/15046 |   0%
        710_10_$a$b$d$0: |                         |     31/15046 |   0%
    710_10_$a$b$d$0$0$9: |                         |     31/15046 |   0%
        710_10_$a$b$x$0: |                         |      3/15046 |   0%
      710_10_$a$b$x$y$0: |                         |      1/15046 |   0%
      710_17_$a$b$w$0$2: |                         |      1/15046 |   0%
        710_20_$a$0$0$9: |                         |      5/15046 |   0%
          710_20_$a$0$w: |                         |     51/15046 |   0%
      710_20_$a$b$0$0$9: |                         |      2/15046 |   0%
        710_20_$a$b$0$w: |                         |      2/15046 |   0%
        710_20_$a$b$d$0: |                         |      1/15046 |   0%
    710_20_$a$b$d$0$0$9: |                         |      1/15046 |   0%
          710_20_$a$b$x: |                         |      1/15046 |   0%
      710_20_$a$b$x$y$0: |                         |      1/15046 |   0%
            710_20_$a$x: |                         |     34/15046 |   0%
          710_20_$a$x$0: |                         |     11/15046 |   0%
        710_20_$a$x$y$0: |                         |      1/15046 |   0%
      710_27_$a$b$w$0$2: |                         |      3/15046 |   0%
    710_27_$a$l$2$4$0$9: |                         |      3/15046 |   0%
        710_27_$a$w$0$2: |                         |      4/15046 |   0%
          711_ 7_$a$0$2: |                         |      1/15046 |   0%
        711_ 7_$a$2$w$0: |                         |      1/15046 |   0%
        711_ 7_$a$w$0$2: |                         |    576/15046 |   3%
      711_ 7_$a$x$w$0$2: |                         |      4/15046 |   0%
            711_04_$a$0: |                         |     88/15046 |   0%
          711_07_$a$0$2: |                         |      2/15046 |   0%
          711_07_$a$2$0: |                         |     93/15046 |   0%
        711_07_$a$w$0$2: |                         |      3/15046 |   0%
    711_10_$a$d$c$0$0$9: |                         |      2/15046 |   0%
            711_20_$a$0: |========                 |   5173/15046 |  34%
        711_20_$a$0$0$9: |========                 |   4969/15046 |  33%
          711_20_$a$c$0: |                         |    109/15046 |   0%
      711_20_$a$c$0$0$9: |                         |     95/15046 |   0%
        711_20_$a$c$d$0: |                         |     19/15046 |   0%
    711_20_$a$c$d$0$0$9: |                         |     18/15046 |   0%
      711_20_$a$c$d$e$0: |                         |      1/15046 |   0%
  711_20_$a$c$d$e$0$0$9: |                         |      1/15046 |   0%
          711_20_$a$d$0: |                         |    546/15046 |   3%
      711_20_$a$d$0$0$9: |                         |    507/15046 |   3%
        711_20_$a$d$c$0: |==                       |   1298/15046 |   8%
    711_20_$a$d$c$0$0$9: |==                       |   1233/15046 |   8%
      711_20_$a$d$c$e$0: |                         |      1/15046 |   0%
  711_20_$a$d$c$e$0$0$9: |                         |      1/15046 |   0%
        711_20_$a$d$e$0: |                         |      1/15046 |   0%
    711_20_$a$d$e$0$0$9: |                         |      1/15046 |   0%
        711_20_$a$n$d$0: |                         |      7/15046 |   0%
    711_20_$a$n$d$0$0$9: |                         |      5/15046 |   0%
      711_20_$a$n$d$c$0: |                         |     37/15046 |   0%
  711_20_$a$n$d$c$0$0$9: |                         |     29/15046 |   0%
      711_20_$a$q$d$e$0: |                         |      1/15046 |   0%
            711_24_$a$0: |                         |      2/15046 |   0%
          711_27_$a$0$2: |                         |     23/15046 |   0%
          711_27_$a$2$0: |==                       |   1577/15046 |  10%
    711_27_$a$l$2$4$0$9: |                         |      3/15046 |   0%
        711_27_$a$w$0$2: |                         |     84/15046 |   0%
        711_27_$a$x$2$0: |                         |      2/15046 |   0%
        730_ 7_$a$w$0$2: |                         |      3/15046 |   0%
    730_07_$a$l$2$4$0$9: |                         |      1/15046 |   0%
            748_ 7_$a$2: |==                       |   1408/15046 |   9%
            750_ 0_$a$0: |=======                  |   4323/15046 |  28%
          750_ 0_$a$0$w: |                         |     21/15046 |   0%
          750_ 0_$a$x$0: |                         |    122/15046 |   0%
        750_ 0_$a$x$x$0: |                         |      2/15046 |   0%
        750_ 0_$a$x$y$0: |                         |      5/15046 |   0%
          750_ 0_$a$y$0: |                         |      9/15046 |   0%
        750_ 0_$a$z$x$0: |                         |      9/15046 |   0%
      750_ 0_$a$z$x$y$0: |                         |      1/15046 |   0%
          750_ 7_$a$0$2: |                         |      1/15046 |   0%
          750_ 7_$a$2$0: |==                       |   1322/15046 |   8%
    750_ 7_$a$l$2$4$0$9: |                         |    477/15046 |   3%
        750_ 7_$a$w$0$2: |                         |      6/15046 |   0%
        750_ 7_$a$x$2$0: |                         |    193/15046 |   1%
          751_ 0_$a$0$w: |                         |     53/15046 |   0%
            751_ 0_$a$x: |                         |     52/15046 |   0%
          751_ 0_$a$x$0: |                         |     42/15046 |   0%
        751_ 0_$a$x$x$0: |                         |      1/15046 |   0%
        751_ 0_$a$x$y$0: |==                       |   1490/15046 |   9%
          751_ 0_$a$y$0: |                         |      2/15046 |   0%
          751_ 7_$a$0$2: |                         |      1/15046 |   0%
          751_ 7_$a$2$0: |                         |     28/15046 |   0%
    751_ 7_$a$l$2$4$0$9: |                         |      2/15046 |   0%
        751_ 7_$a$w$0$2: |                         |      1/15046 |   0%
        751_ 7_$a$z$2$0: |                         |     12/15046 |   0%
          755_ 7_$a$2$0: |                         |    193/15046 |   1%
        755_ 7_$a$v$2$0: |                         |     12/15046 |   0%
          780_ 0_$x$0$w: |                         |     57/15046 |   0%
```


### FormGenre Fields

```
                        001: |=========================|   3104/3104 | 100%
                        003: |=========================|   3104/3104 | 100%
                        005: |=========================|   3104/3104 | 100%
                        008: |=========================|   3104/3104 | 100%
                016_7 _$a$2: |======================== |   3042/3104 |  98%
              016_7 _$a$z$2: |                         |     46/3104 |   1%
            016_7 _$a$z$z$2: |                         |      3/3104 |   0%
          016_7 _$a$z$z$z$2: |                         |      4/3104 |   0%
        016_7 _$a$z$z$z$z$2: |                         |      2/3104 |   0%
      016_7 _$a$z$z$z$z$z$2: |                         |      1/3104 |   0%
    016_7 _$a$z$z$z$z$z$z$2: |                         |      1/3104 |   0%
  016_7 _$a$z$z$z$z$z$z$z$2: |                         |      2/3104 |   0%
016_7 _$a$z$z$z$z$z$z$z$z$2: |                         |      1/3104 |   0%
            040_  _$a$b$c$f: |=========================|   3104/3104 | 100%
                  155_  _$a: |====================     |   2557/3104 |  82%
                155_  _$a$v: |===                      |    486/3104 |  15%
              155_  _$a$v$v: |                         |     61/3104 |   1%
                445_  _$a$5: |                         |      1/3104 |   0%
                  450_  _$a: |                         |      3/3104 |   0%
                  455_  _$a: |============             |   1540/3104 |  49%
                455_  _$a$v: |                         |      2/3104 |   0%
            455_  _$a$v$v$w: |                         |     50/3104 |   1%
              455_  _$a$v$w: |                         |     38/3104 |   1%
                455_  _$a$w: |                         |    110/3104 |   3%
                  485_  _$v: |                         |      1/3104 |   0%
                485_  _$v$5: |                         |      1/3104 |   0%
                555_  _$a$0: |=================        |   2125/3104 |  68%
                  680_  _$i: |======                   |    766/3104 |  24%
                680_  _$i$5: |                         |      1/3104 |   0%
                680_  _$i$a: |=                        |    157/3104 |   5%
              680_  _$i$a$i: |                         |      3/3104 |   0%
            680_  _$i$a$i$a: |                         |      7/3104 |   0%
          680_  _$i$a$i$a$i: |                         |      7/3104 |   0%
                  688_  _$a: |=======================  |   2922/3104 |  94%
          700_00_$a$v$v$0$w: |                         |      1/3104 |   0%
                750_ 0_$a$0: |                         |      4/3104 |   0%
              750_ 0_$a$0$w: |=                        |    212/3104 |   6%
            750_ 0_$a$v$0$w: |                         |     21/3104 |   0%
          750_ 0_$a$v$v$0$w: |                         |      4/3104 |   0%
                755_ 0_$a$0: |                         |     93/3104 |   2%
              755_ 7_$a$0$2: |                         |     86/3104 |   2%
              755_ 7_$a$2$0: |================         |   2065/3104 |  66%
            755_ 7_$a$v$2$0: |                         |      1/3104 |   0%
                785_  _$v$0: |                         |      1/3104 |   0%
              785_ 0_$a$v$0: |                         |      1/3104 |   0%
            785_ 0_$a$v$0$w: |                         |      1/3104 |   0%
            785_ 0_$a$v$v$0: |                         |      1/3104 |   0%
                785_ 0_$v$0: |==                       |    281/3104 |   9%
              785_ 0_$v$0$w: |                         |      6/3104 |   0%
              785_ 0_$v$v$0: |                         |     66/3104 |   2%
              785_ 0_$v$x$0: |                         |      1/3104 |   0%
```

### Topical Fields

```
                                          001: |=========================| 453891/453891 | 100%
                                          003: |=========================| 453891/453891 | 100%
                                          005: |=========================| 453891/453891 | 100%
                                          008: |=========================| 453891/453891 | 100%
                                  016_7 _$a$2: |======================== | 440368/453891 |  97%
                                016_7 _$a$z$2: |                         |  13280/453891 |   2%
                              016_7 _$a$z$z$2: |                         |    218/453891 |   0%
                            016_7 _$a$z$z$z$2: |                         |     22/453891 |   0%
                        016_7 _$a$z$z$z$z$z$2: |                         |      2/453891 |   0%
                              040_  _$a$b$c$f: |======================== | 453890/453891 |  99%
                                    043_  _$a: |                         |      1/453891 |   0%
                                    046_  _$s: |                         |      1/453891 |   0%
                                    053_  _$a: |                         |    137/453891 |   0%
                                  053_  _$a$b: |                         |      3/453891 |   0%
                                053_  _$a$b$c: |                         |      1/453891 |   0%
                                  053_  _$a$c: |                         |      2/453891 |   0%
                                    053_ 0_$a: |===                      |  55282/453891 |  12%
                                  053_ 0_$a$b: |                         |  12155/453891 |   2%
                                053_ 0_$a$b$c: |                         |   2771/453891 |   0%
                                  053_ 0_$a$c: |                         |  16218/453891 |   3%
                                150_  _$4$x$x: |                         |      1/453891 |   0%
                                    150_  _$a: |==========               | 192413/453891 |  42%
                                  150_  _$a$x: |==========               | 195840/453891 |  43%
                                150_  _$a$x$x: |=                        |  33453/453891 |   7%
                              150_  _$a$x$x$x: |                         |   1743/453891 |   0%
                            150_  _$a$x$x$x$x: |                         |     42/453891 |   0%
                          150_  _$a$x$x$x$x$x: |                         |      2/453891 |   0%
                                150_  _$a$x$y: |                         |      1/453891 |   0%
                                    150_ 0_$a: |                         |     97/453891 |   0%
                                  150_ 0_$a$v: |                         |      1/453891 |   0%
                                  150_ 0_$a$x: |=                        |  23390/453891 |   5%
                                150_ 0_$a$x$x: |                         |   6662/453891 |   1%
                              150_ 0_$a$x$x$x: |                         |    246/453891 |   0%
                                    260_  _$i: |                         |     20/453891 |   0%
                          360_  _$a$a$i$a$i$a: |                         |      1/453891 |   0%
                                  360_  _$a$i: |                         |      7/453891 |   0%
                              360_  _$a$i$a$i: |                         |      1/453891 |   0%
                                    360_  _$i: |                         |    321/453891 |   0%
                                  360_  _$i$a: |                         |    284/453891 |   0%
                                360_  _$i$a$a: |                         |      3/453891 |   0%
                              360_  _$i$a$a$a: |                         |      2/453891 |   0%
                            360_  _$i$a$a$i$a: |                         |      1/453891 |   0%
                                360_  _$i$a$i: |                         |    626/453891 |   0%
                              360_  _$i$a$i$a: |                         |   1909/453891 |   0%
                            360_  _$i$a$i$a$i: |                         |    129/453891 |   0%
                          360_  _$i$a$i$a$i$a: |                         |    137/453891 |   0%
                  360_  _$i$a$i$a$i$a$a$i$a$a: |                         |      1/453891 |   0%
                        360_  _$i$a$i$a$i$a$i: |                         |     29/453891 |   0%
                      360_  _$i$a$i$a$i$a$i$a: |                         |     98/453891 |   0%
                    360_  _$i$a$i$a$i$a$i$a$i: |                         |      8/453891 |   0%
                  360_  _$i$a$i$a$i$a$i$a$i$a: |                         |     18/453891 |   0%
                360_  _$i$a$i$a$i$a$i$a$i$a$i: |                         |      6/453891 |   0%
              360_  _$i$a$i$a$i$a$i$a$i$a$i$a: |                         |     16/453891 |   0%
          360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a: |                         |      3/453891 |   0%
        360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i: |                         |      2/453891 |   0%
      360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a: |                         |      1/453891 |   0%
  360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a: |                         |      1/453891 |   0%
360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i: |                         |      1/453891 |   0%
                            360_  _$i$a$i$i$a: |                         |      1/453891 |   0%
                          360_  _$i$i$a$i$a$i: |                         |      1/453891 |   0%
                          400_0 _$a$b$c$d$x$w: |                         |      1/453891 |   0%
                                  400_0 _$a$c: |                         |      1/453891 |   0%
                                  400_1 _$a$d: |                         |      1/453891 |   0%
                                  400_1 _$a$w: |                         |      1/453891 |   0%
                                400_3 _$a$x$w: |                         |      1/453891 |   0%
                                410_  _$a$b$w: |                         |     11/453891 |   0%
                                  410_1 _$a$b: |                         |      1/453891 |   0%
                                    410_2 _$a: |                         |      5/453891 |   0%
                                  410_2 _$a$b: |                         |      1/453891 |   0%
                                  410_2 _$a$w: |                         |     16/453891 |   0%
                                  411_  _$a$w: |                         |      7/453891 |   0%
                                    411_2 _$a: |                         |      1/453891 |   0%
                                    430_ 0_$a: |                         |     13/453891 |   0%
                                  430_ 0_$a$w: |                         |      6/453891 |   0%
                                    450_  _$a: |======                   | 109070/453891 |  24%
                                  450_  _$a$5: |                         |      1/453891 |   0%
                                  450_  _$a$w: |                         |   8497/453891 |   1%
                                  450_  _$a$x: |                         |  10831/453891 |   2%
                                450_  _$a$x$w: |                         |   4451/453891 |   0%
                                450_  _$a$x$x: |                         |    619/453891 |   0%
                              450_  _$a$x$x$w: |                         |   1095/453891 |   0%
                              450_  _$a$x$x$x: |                         |     21/453891 |   0%
                            450_  _$a$x$x$x$w: |                         |     80/453891 |   0%
                          450_  _$a$x$x$x$x$w: |                         |      2/453891 |   0%
                              450_  _$a$x$y$w: |                         |      1/453891 |   0%
                                451_  _$a$x$5: |                         |      1/453891 |   0%
                                451_  _$a$x$w: |                         |      1/453891 |   0%
                                451_  _$a$z$w: |                         |      3/453891 |   0%
                              451_  _$a$z$z$w: |                         |      1/453891 |   0%
                                  500_0 _$a$0: |                         |    160/453891 |   0%
                                500_0 _$a$c$0: |                         |    122/453891 |   0%
                              500_0 _$a$c$0$w: |                         |      2/453891 |   0%
                              500_0 _$a$c$d$0: |                         |      4/453891 |   0%
                                500_0 _$a$d$0: |                         |      1/453891 |   0%
                                  500_1 _$a$0: |                         |      1/453891 |   0%
                                500_1 _$a$d$0: |                         |      5/453891 |   0%
                                  500_3 _$a$0: |                         |      1/453891 |   0%
                                500_3 _$a$0$w: |                         |      1/453891 |   0%
                                510_1 _$a$b$0: |                         |    101/453891 |   0%
                                  510_2 _$a$0: |                         |    137/453891 |   0%
                                510_2 _$a$0$w: |                         |    126/453891 |   0%
                                  511_  _$a$0: |                         |     37/453891 |   0%
                                  511_  _$a$x: |                         |      2/453891 |   0%
                                  511_0 _$a$0: |                         |      6/453891 |   0%
                                511_0 _$a$0$w: |                         |      7/453891 |   0%
                                  511_2 _$a$0: |                         |     95/453891 |   0%
                                511_2 _$a$0$w: |                         |      8/453891 |   0%
                                  530_0 _$a$0: |                         |     28/453891 |   0%
                                530_0 _$a$0$w: |                         |     36/453891 |   0%
                                530_0 _$a$p$0: |                         |     16/453891 |   0%
                              530_0 _$a$p$0$w: |                         |     10/453891 |   0%
                                  548_  _$a$2: |                         |    771/453891 |   0%
                                    550_  _$a: |                         |     43/453891 |   0%
                                  550_  _$a$0: |=                        |  29456/453891 |   6%
                                550_  _$a$0$w: |=======                  | 130523/453891 |  28%
                                  550_  _$a$x: |                         |      4/453891 |   0%
                                550_  _$a$x$0: |                         |   2110/453891 |   0%
                              550_  _$a$x$0$w: |                         |  11478/453891 |   2%
                              550_  _$a$x$x$0: |                         |     71/453891 |   0%
                            550_  _$a$x$x$0$w: |                         |    601/453891 |   0%
                            550_  _$a$x$x$x$0: |                         |      3/453891 |   0%
                          550_  _$a$x$x$x$0$w: |                         |     11/453891 |   0%
                                  551_  _$a$0: |                         |  15634/453891 |   3%
                                551_  _$a$0$w: |                         |     54/453891 |   0%
                                551_  _$a$z$0: |                         |    160/453891 |   0%
                              551_  _$a$z$0$w: |                         |      6/453891 |   0%
                              551_  _$a$z$z$0: |                         |      1/453891 |   0%
                                  555_  _$a$0: |                         |   1284/453891 |   0%
                                555_  _$a$v$0: |                         |      1/453891 |   0%
                                    680_  _$i: |                         |     92/453891 |   0%
                                  680_  _$i$a: |                         |     27/453891 |   0%
                                680_  _$i$a$i: |                         |      6/453891 |   0%
                              680_  _$i$a$i$a: |                         |      2/453891 |   0%
                            680_  _$i$a$i$a$i: |                         |      1/453891 |   0%
                                    688_  _$a: |======================   | 407795/453891 |  89%
                                700_ 7_$a$0$2: |                         |      1/453891 |   0%
                              700_ 7_$a$c$0$2: |                         |      1/453891 |   0%
                              700_ 7_$a$c$2$0: |                         |      3/453891 |   0%
                          700_00_$a$b$c$d$x$0: |                         |      1/453891 |   0%
                            700_00_$a$c$d$x$0: |                         |      8/453891 |   0%
                              700_00_$a$c$x$0: |                         |     67/453891 |   0%
                              700_00_$a$d$x$0: |                         |      4/453891 |   0%
                                700_00_$a$v$0: |                         |      1/453891 |   0%
                                700_00_$a$x$0: |                         |     89/453891 |   0%
                              700_00_$a$x$0$w: |                         |      2/453891 |   0%
                              700_00_$a$x$x$0: |                         |      7/453891 |   0%
                                700_07_$a$0$2: |                         |      3/453891 |   0%
                                700_07_$a$2$0: |                         |      5/453891 |   0%
                              700_07_$a$c$0$2: |                         |     41/453891 |   0%
                              700_07_$a$c$2$0: |                         |    154/453891 |   0%
                          700_07_$a$c$d$w$0$2: |                         |      1/453891 |   0%
                            700_07_$a$c$w$0$2: |                         |      8/453891 |   0%
                              700_07_$a$d$2$0: |                         |      1/453891 |   0%
                            700_07_$a$d$w$0$2: |                         |      1/453891 |   0%
                              700_07_$a$w$0$2: |                         |      2/453891 |   0%
                              700_10_$a$c$0$w: |                         |      1/453891 |   0%
                              700_10_$a$c$x$0: |                         |      2/453891 |   0%
                          700_10_$a$d$x$x$0$w: |                         |      3/453891 |   0%
                                700_10_$a$x$0: |                         |      2/453891 |   0%
                              700_10_$a$x$0$w: |                         |      1/453891 |   0%
                                700_17_$a$0$2: |                         |      4/453891 |   0%
                              700_17_$a$c$2$0: |                         |     10/453891 |   0%
                              700_17_$a$d$0$2: |                         |      1/453891 |   0%
                                700_30_$a$x$0: |                         |      1/453891 |   0%
                            700_37_$a$d$w$0$2: |                         |      1/453891 |   0%
                              700_37_$a$w$0$2: |                         |      1/453891 |   0%
                                710_ 7_$a$0$2: |                         |     15/453891 |   0%
                        710_ 7_$a$b$b$b$w$0$2: |                         |      6/453891 |   0%
                          710_ 7_$a$b$b$w$0$2: |                         |      6/453891 |   0%
                            710_ 7_$a$b$w$0$2: |                         |      1/453891 |   0%
                              710_ 7_$a$w$0$2: |                         |      1/453891 |   0%
                              710_10_$a$b$x$0: |                         |    133/453891 |   0%
                            710_10_$a$b$x$x$0: |                         |     20/453891 |   0%
                          710_10_$a$b$x$x$x$0: |                         |      2/453891 |   0%
                            710_17_$a$b$b$2$0: |                         |      1/453891 |   0%
                          710_17_$a$b$b$w$0$2: |                         |      2/453891 |   0%
                              710_20_$a$0$0$9: |                         |      1/453891 |   0%
                                710_20_$a$x$0: |                         |     22/453891 |   0%
                              710_20_$a$x$0$w: |                         |     52/453891 |   0%
                              710_20_$a$x$x$0: |                         |      1/453891 |   0%
                                710_27_$a$0$2: |                         |      5/453891 |   0%
                                710_27_$a$2$0: |                         |     46/453891 |   0%
                          710_27_$a$l$2$4$0$9: |                         |    342/453891 |   0%
                              710_27_$a$w$0$2: |                         |     26/453891 |   0%
                                711_ 7_$a$0$2: |                         |      5/453891 |   0%
                                711_ 7_$a$2$0: |                         |     12/453891 |   0%
                              711_ 7_$a$2$w$0: |                         |      1/453891 |   0%
                              711_ 7_$a$w$0$2: |                         |     78/453891 |   0%
                            711_ 7_$a$x$w$0$2: |                         |      1/453891 |   0%
                                711_07_$a$2$0: |                         |     12/453891 |   0%
                              711_20_$a$c$x$0: |                         |      1/453891 |   0%
                                711_27_$a$0$2: |                         |      4/453891 |   0%
                                711_27_$a$2$0: |                         |     45/453891 |   0%
                          711_27_$a$l$2$4$0$9: |                         |    181/453891 |   0%
                              711_27_$a$w$0$2: |                         |     18/453891 |   0%
                              711_27_$a$x$2$0: |                         |      1/453891 |   0%
                              730_ 0_$a$0$0$9: |                         |      1/453891 |   0%
                              730_ 0_$a$x$0$w: |                         |      2/453891 |   0%
                            730_ 0_$a$x$v$0$w: |                         |      1/453891 |   0%
                                730_ 7_$a$0$2: |                         |      2/453891 |   0%
                                730_ 7_$a$2$0: |                         |      2/453891 |   0%
                              730_ 7_$a$w$0$2: |                         |     51/453891 |   0%
                                730_07_$a$2$0: |                         |      1/453891 |   0%
                          730_07_$a$l$2$4$0$9: |                         |    122/453891 |   0%
                              730_07_$a$w$0$2: |                         |      1/453891 |   0%
                                  748_ 7_$a$2: |                         |    631/453891 |   0%
                                    750_ 0_$a: |                         |  15751/453891 |   3%
                                  750_ 0_$a$0: |=========                | 175777/453891 |  38%
                                750_ 0_$a$0$w: |========                 | 159741/453891 |  35%
                                750_ 0_$a$4$w: |                         |      3/453891 |   0%
                                750_ 0_$a$v$0: |                         |   1035/453891 |   0%
                              750_ 0_$a$v$0$w: |                         |      2/453891 |   0%
                              750_ 0_$a$v$v$0: |                         |     47/453891 |   0%
                                  750_ 0_$a$x: |                         |     38/453891 |   0%
                                750_ 0_$a$x$0: |==                       |  41547/453891 |   9%
                              750_ 0_$a$x$0$w: |=                        |  19047/453891 |   4%
                              750_ 0_$a$x$v$0: |                         |     29/453891 |   0%
                              750_ 0_$a$x$x$0: |                         |   5444/453891 |   1%
                            750_ 0_$a$x$x$0$w: |                         |    680/453891 |   0%
                            750_ 0_$a$x$x$v$0: |                         |      6/453891 |   0%
                            750_ 0_$a$x$x$x$0: |                         |    271/453891 |   0%
                          750_ 0_$a$x$x$x$0$w: |                         |      4/453891 |   0%
                          750_ 0_$a$x$x$x$x$0: |                         |      9/453891 |   0%
                            750_ 0_$a$x$x$y$0: |                         |     10/453891 |   0%
                          750_ 0_$a$x$x$y$0$w: |                         |      1/453891 |   0%
                        750_ 0_$a$x$x$y$x$v$0: |                         |      1/453891 |   0%
                            750_ 0_$a$x$x$z$0: |                         |      1/453891 |   0%
                          750_ 0_$a$x$x$z$0$w: |                         |      1/453891 |   0%
                          750_ 0_$a$x$x$z$x$0: |                         |     10/453891 |   0%
                              750_ 0_$a$x$y$0: |                         |     94/453891 |   0%
                            750_ 0_$a$x$y$0$w: |                         |      2/453891 |   0%
                            750_ 0_$a$x$y$x$0: |                         |      4/453891 |   0%
                              750_ 0_$a$x$z$0: |                         |     14/453891 |   0%
                            750_ 0_$a$x$z$0$w: |                         |      1/453891 |   0%
                            750_ 0_$a$x$z$x$0: |                         |     63/453891 |   0%
                          750_ 0_$a$x$z$x$0$w: |                         |      1/453891 |   0%
                          750_ 0_$a$x$z$x$v$0: |                         |      1/453891 |   0%
                          750_ 0_$a$x$z$x$x$0: |                         |      1/453891 |   0%
                                750_ 0_$a$y$0: |                         |    324/453891 |   0%
                              750_ 0_$a$y$0$w: |                         |    348/453891 |   0%
                              750_ 0_$a$y$v$0: |                         |      2/453891 |   0%
                              750_ 0_$a$y$x$0: |                         |     22/453891 |   0%
                            750_ 0_$a$y$x$0$w: |                         |      5/453891 |   0%
                                750_ 0_$a$z$0: |                         |    112/453891 |   0%
                              750_ 0_$a$z$0$w: |                         |      1/453891 |   0%
                              750_ 0_$a$z$v$0: |                         |      1/453891 |   0%
                              750_ 0_$a$z$x$0: |                         |     53/453891 |   0%
                            750_ 0_$a$z$x$0$w: |                         |     11/453891 |   0%
                            750_ 0_$a$z$x$v$0: |                         |      2/453891 |   0%
                            750_ 0_$a$z$x$x$0: |                         |      1/453891 |   0%
                            750_ 0_$a$z$x$y$0: |                         |    152/453891 |   0%
                              750_ 0_$a$z$y$0: |                         |      2/453891 |   0%
                            750_ 0_$a$z$z$x$0: |                         |      6/453891 |   0%
                                750_ 4_$4$0$9: |                         |    627/453891 |   0%
                            750_ 4_$a$4$0$0$9: |===                      |  61232/453891 |  13%
                              750_ 4_$a$4$0$9: |                         |   1392/453891 |   0%
                        750_ 7_$226$l$2$4$0$9: |                         |      2/453891 |   0%
                        750_ 7_$254$l$2$4$0$9: |                         |     10/453891 |   0%
                                  750_ 7_$a$0: |                         |      1/453891 |   0%
                                750_ 7_$a$0$2: |                         |    750/453891 |   0%
                              750_ 7_$a$0$2$2: |                         |      1/453891 |   0%
                              750_ 7_$a$0$2$w: |                         |      1/453891 |   0%
                                  750_ 7_$a$2: |                         |      1/453891 |   0%
                                750_ 7_$a$2$0: |=                        |  24975/453891 |   5%
                            750_ 7_$a$2$4$0$9: |                         |   2597/453891 |   0%
                              750_ 7_$a$2$w$0: |                         |      8/453891 |   0%
                          750_ 7_$a$l$2$4$0$9: |===                      |  56131/453891 |  12%
                              750_ 7_$a$w$0$2: |                         |   7091/453891 |   1%
                              750_ 7_$a$w$2$0: |                         |     26/453891 |   0%
                              750_ 7_$a$x$0$2: |                         |    131/453891 |   0%
                            750_ 7_$a$x$0$2$w: |                         |      1/453891 |   0%
                              750_ 7_$a$x$2$0: |                         |   6775/453891 |   1%
                            750_ 7_$a$x$2$w$0: |                         |      2/453891 |   0%
                            750_ 7_$a$x$w$0$2: |                         |   3367/453891 |   0%
                            750_ 7_$a$x$w$2$0: |                         |      1/453891 |   0%
                            750_ 7_$a$x$x$0$2: |                         |      9/453891 |   0%
                            750_ 7_$a$x$x$2$0: |                         |    257/453891 |   0%
                          750_ 7_$a$x$x$w$0$2: |                         |    931/453891 |   0%
                          750_ 7_$a$x$x$x$2$0: |                         |      2/453891 |   0%
                        750_ 7_$a$x$x$x$w$0$2: |                         |     59/453891 |   0%
                                751_ 0_$a$x$0: |                         |     67/453891 |   0%
                              751_ 0_$a$x$0$w: |                         |     16/453891 |   0%
                            751_ 0_$a$x$v$0$w: |                         |      1/453891 |   0%
                              751_ 0_$a$x$x$0: |                         |     13/453891 |   0%
                            751_ 0_$a$x$x$0$w: |                         |     74/453891 |   0%
                            751_ 0_$a$x$x$x$0: |                         |      2/453891 |   0%
                              751_ 0_$a$x$y$0: |                         |      1/453891 |   0%
                            751_ 0_$a$x$y$0$w: |                         |     38/453891 |   0%
                                751_ 7_$a$0$2: |                         |      3/453891 |   0%
                                751_ 7_$a$2$0: |                         |     15/453891 |   0%
                          751_ 7_$a$l$2$4$0$9: |                         |    223/453891 |   0%
                              751_ 7_$a$w$0$2: |                         |      4/453891 |   0%
                              751_ 7_$a$z$0$2: |                         |      2/453891 |   0%
                            751_ 7_$a$z$0$2$w: |                         |      2/453891 |   0%
                              751_ 7_$a$z$2$0: |                         |      8/453891 |   0%
                            751_ 7_$a$z$w$0$2: |                         |     15/453891 |   0%
                          751_ 7_$a$z$z$w$0$2: |                         |      4/453891 |   0%
                                755_ 7_$a$0$2: |                         |    139/453891 |   0%
                              755_ 7_$a$0$2$2: |                         |      1/453891 |   0%
                                  755_ 7_$a$2: |                         |      1/453891 |   0%
                                755_ 7_$a$2$0: |=                        |  30402/453891 |   6%
                              755_ 7_$a$v$2$0: |                         |      5/453891 |   0%
                              755_ 7_$a$w$2$0: |                         |      1/453891 |   0%
                                780_  _$x$0$w: |                         |      1/453891 |   0%
                                  780_ 0_$a$0: |                         |      1/453891 |   0%
                                780_ 0_$a$0$w: |                         |      4/453891 |   0%
                                  780_ 0_$x$0: |                         |     38/453891 |   0%
                                780_ 0_$x$0$w: |=========                | 165644/453891 |  36%
                              780_ 0_$x$0$w$w: |                         |      1/453891 |   0%
                                780_ 0_$x$o$w: |                         |      4/453891 |   0%
                              780_ 0_$x$x$0$w: |                         |   1466/453891 |   0%
                            780_ 0_$x$x$x$0$w: |                         |      5/453891 |   0%
                                781_ 0_$z$0$w: |                         |      1/453891 |   0%
                                785_ 0_$v$0$w: |                         |     13/453891 |   0%
```
