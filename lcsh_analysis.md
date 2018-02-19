# LCSH


## Mapped Fields for Resolver


## MARC Data Analysis (Subset)

tag       | field info | note
--------- | ---------- | ----
001       | LCSH identifier. |
003       | Do not map. |
005       | Do not map.
008       | Do not map. |
010       | 
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

```


                                                                                                                                          001: |=========================| 414695/414695 | 100%
                                                                                                                                          003: |=========================| 414695/414695 | 100%
                                                                                                                                          005: |=========================| 414695/414695 | 100%
                                                                                                                                          008: |=========================| 414695/414695 | 100%
                                                                                                                                    010_  _$a: |======================== | 411019/414695 |  99%
                                                                                                                                  010_  _$a$z: |                         |   3464/414695 |   0%
                                                                                                                                010_  _$a$z$z: |                         |    160/414695 |   0%
                                                                                                                              010_  _$a$z$z$z: |                         |     28/414695 |   0%
                                                                                                                            010_  _$a$z$z$z$z: |                         |     11/414695 |   0%
                                                                                                                          010_  _$a$z$z$z$z$z: |                         |      6/414695 |   0%
                                                                                                                        010_  _$a$z$z$z$z$z$z: |                         |      3/414695 |   0%
                                                                                            010_  _$a$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z: |                         |      1/414695 |   0%
                                                                                010_  _$a$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z: |                         |      1/414695 |   0%
                                                                  010_  _$a$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z: |                         |      1/414695 |   0%
010_  _$a$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z$z: |                         |      1/414695 |   0%
                                                                                                                              034_  _$d$e$f$g: |                         |      1/414695 |   0%
                                                                                                                            034_  _$d$e$f$g$2: |                         |  16546/414695 |   3%
                                                                                                                                    035_  _$a: |                         |      7/414695 |   0%
                                                                                                                                040_  _$a$b$c: |=========                | 154382/414695 |  37%
                                                                                                                              040_  _$a$b$c$d: |                         |  14958/414695 |   3%
                                                                                                                            040_  _$a$b$c$d$d: |                         |    360/414695 |   0%
                                                                                                                          040_  _$a$b$c$d$d$d: |                         |     59/414695 |   0%
                                                                                                                        040_  _$a$b$c$d$d$d$d: |                         |      8/414695 |   0%
                                                                                                                      040_  _$a$b$c$d$d$d$d$d: |                         |      1/414695 |   0%
                                                                                                                            040_  _$a$b$c$d$f: |                         |      9/414695 |   0%
                                                                                                                              040_  _$a$b$c$f: |                         |    554/414695 |   0%
                                                                                                                            040_  _$a$b$c$f$d: |                         |     92/414695 |   0%
                                                                                                                          040_  _$a$b$c$f$d$d: |                         |     12/414695 |   0%
                                                                                                                                  040_  _$a$c: |=====                    |  89454/414695 |  21%
                                                                                                                                040_  _$a$c$ : |                         |      1/414695 |   0%
                                                                                                                              040_  _$a$c$b$d: |                         |      2/414695 |   0%
                                                                                                                                040_  _$a$c$d: |========                 | 148685/414695 |  35%
                                                                                                                              040_  _$a$c$d$d: |                         |   4460/414695 |   1%
                                                                                                                            040_  _$a$c$d$d$d: |                         |   1506/414695 |   0%
                                                                                                                          040_  _$a$c$d$d$d$d: |                         |    111/414695 |   0%
                                                                                                                        040_  _$a$c$d$d$d$d$d: |                         |     37/414695 |   0%
                                                                                                                      040_  _$a$c$d$d$d$d$d$d: |                         |      3/414695 |   0%
                                                                                                                040_  _$a$c$d$d$d$d$d$d$d$d$d: |                         |      1/414695 |   0%
                                                                                                                                    043_  _$a: |                         |      1/414695 |   0%
                                                                                                                                    050_ 0_$a: |                         |      1/414695 |   0%
                                                                                                                                    052_  _$a: |                         |      2/414695 |   0%
                                                                                                                                    053_  _$a: |==                       |  39454/414695 |   9%
                                                                                                                                  053_  _$a$b: |                         |   7337/414695 |   1%
                                                                                                                                053_  _$a$b$c: |                         |   1464/414695 |   0%
                                                                                                                                  053_  _$a$c: |                         |   7320/414695 |   1%
                                                                                                                                    053_ 0_$a: |=                        |  17468/414695 |   4%
                                                                                                                                  053_ 0_$a$b: |                         |   5575/414695 |   1%
                                                                                                                                053_ 0_$a$b$c: |                         |   1473/414695 |   0%
                                                                                                                                  053_ 0_$a$c: |                         |  10215/414695 |   2%
                                                                                                                                  072_ 7_$a$2: |                         |   1022/414695 |   0%
                                                                                                                                  072_ 7_$a$a: |                         |      1/414695 |   0%
                                                                                                                073_  _$a$a$a$a$a$a$a$a$a$a$z: |                         |      3/414695 |   0%
                                                                                                                  073_  _$a$a$a$a$a$a$a$a$a$z: |                         |     10/414695 |   0%
                                                                                                                    073_  _$a$a$a$a$a$a$a$a$z: |                         |      7/414695 |   0%
                                                                                                                      073_  _$a$a$a$a$a$a$a$z: |                         |     31/414695 |   0%
                                                                                                                        073_  _$a$a$a$a$a$a$z: |                         |     32/414695 |   0%
                                                                                                                          073_  _$a$a$a$a$a$z: |                         |     51/414695 |   0%
                                                                                                                            073_  _$a$a$a$a$z: |                         |    106/414695 |   0%
                                                                                                                              073_  _$a$a$a$z: |                         |    188/414695 |   0%
                                                                                                                                073_  _$a$a$z: |                         |    432/414695 |   0%
                                                                                                                                  073_  _$a$z: |                         |   2744/414695 |   0%
                                                                                                                                    100_0 _$a: |                         |      2/414695 |   0%
                                                                                                                            100_0 _$a$b$c$d$v: |                         |      3/414695 |   0%
                                                                                                                            100_0 _$a$b$c$d$x: |                         |      3/414695 |   0%
                                                                                                                                  100_0 _$a$c: |                         |      1/414695 |   0%
                                                                                                                                100_0 _$a$c$d: |                         |      6/414695 |   0%
                                                                                                                              100_0 _$a$c$d$v: |                         |      9/414695 |   0%
                                                                                                                              100_0 _$a$c$d$x: |                         |     18/414695 |   0%
                                                                                                                            100_0 _$a$c$d$x$x: |                         |      1/414695 |   0%
                                                                                                                                100_0 _$a$c$v: |                         |     17/414695 |   0%
                                                                                                                                100_0 _$a$c$x: |                         |     64/414695 |   0%
                                                                                                                              100_0 _$a$c$x$x: |                         |      1/414695 |   0%
                                                                                                                            100_0 _$a$c$x$x$v: |                         |      1/414695 |   0%
                                                                                                                              100_0 _$a$c$x$z: |                         |     25/414695 |   0%
                                                                                                                                  100_0 _$a$d: |                         |      6/414695 |   0%
                                                                                                                                100_0 _$a$d$t: |                         |      1/414695 |   0%
                                                                                                                              100_0 _$a$d$t$v: |                         |      1/414695 |   0%
                                                                                                                                100_0 _$a$d$x: |                         |      5/414695 |   0%
                                                                                                                              100_0 _$a$d$x$x: |                         |      1/414695 |   0%
                                                                                                                                  100_0 _$a$t: |                         |      1/414695 |   0%
                                                                                                                                  100_0 _$a$v: |                         |     14/414695 |   0%
                                                                                                                                100_0 _$a$v$v: |                         |      4/414695 |   0%
                                                                                                                              100_0 _$a$v$v$v: |                         |      1/414695 |   0%
                                                                                                                                  100_0 _$a$x: |                         |     38/414695 |   0%
                                                                                                                                100_0 _$a$x$v: |                         |     23/414695 |   0%
                                                                                                                                100_0 _$a$x$x: |                         |      1/414695 |   0%
                                                                                                                              100_0 _$a$x$x$v: |                         |      2/414695 |   0%
                                                                                                                                100_0 _$a$x$z: |                         |      2/414695 |   0%
                                                                                                                            100_00_$a$b$c$d$v: |                         |      3/414695 |   0%
                                                                                                                            100_00_$a$b$c$d$x: |                         |     14/414695 |   0%
                                                                                                                          100_00_$a$b$c$d$x$z: |                         |      1/414695 |   0%
                                                                                                                              100_00_$a$c$d$v: |                         |      2/414695 |   0%
                                                                                                                              100_00_$a$c$d$x: |                         |      2/414695 |   0%
                                                                                                                            100_00_$a$c$d$x$z: |                         |      1/414695 |   0%
                                                                                                                                100_00_$a$c$x: |                         |     18/414695 |   0%
                                                                                                                              100_00_$a$c$x$z: |                         |      3/414695 |   0%
                                                                                                                                  100_00_$a$v: |                         |      1/414695 |   0%
                                                                                                                                  100_00_$a$x: |                         |    100/414695 |   0%
                                                                                                                                100_00_$a$x$v: |                         |      6/414695 |   0%
                                                                                                                                100_00_$a$x$x: |                         |     18/414695 |   0%
                                                                                                                                100_00_$a$x$y: |                         |      8/414695 |   0%
                                                                                                                                100_00_$a$x$z: |                         |      2/414695 |   0%
                                                                                                                                    100_1 _$a: |                         |      1/414695 |   0%
                                                                                                                                100_1 _$a$c$d: |                         |      1/414695 |   0%
                                                                                                                              100_1 _$a$c$d$t: |                         |      1/414695 |   0%
                                                                                                                              100_1 _$a$c$d$x: |                         |      2/414695 |   0%
                                                                                                                              100_1 _$a$c$x$z: |                         |      1/414695 |   0%
                                                                                                                                  100_1 _$a$d: |                         |     15/414695 |   0%
                                                                                                                                100_1 _$a$d$t: |                         |     24/414695 |   0%
                                                                                                                              100_1 _$a$d$t$v: |                         |     11/414695 |   0%
                                                                                                                              100_1 _$a$d$t$x: |                         |      1/414695 |   0%
                                                                                                                                100_1 _$a$d$v: |                         |     75/414695 |   0%
                                                                                                                              100_1 _$a$d$v$v: |                         |      6/414695 |   0%
                                                                                                                                100_1 _$a$d$x: |                         |    173/414695 |   0%
                                                                                                                              100_1 _$a$d$x$v: |                         |     14/414695 |   0%
                                                                                                                              100_1 _$a$d$x$x: |                         |    199/414695 |   0%
                                                                                                                            100_1 _$a$d$x$x$v: |                         |      1/414695 |   0%
                                                                                                                            100_1 _$a$d$x$x$y: |                         |      4/414695 |   0%
                                                                                                                              100_1 _$a$d$x$y: |                         |      4/414695 |   0%
                                                                                                                              100_1 _$a$d$x$z: |                         |     14/414695 |   0%
                                                                                                                              100_1 _$a$q$d$x: |                         |     13/414695 |   0%
                                                                                                                            100_1 _$a$q$d$x$x: |                         |      9/414695 |   0%
                                                                                                                                100_1 _$a$q$x: |                         |      1/414695 |   0%
                                                                                                                                  100_1 _$a$v: |                         |      5/414695 |   0%
                                                                                                                                  100_1 _$a$x: |                         |     26/414695 |   0%
                                                                                                                                100_1 _$a$x$x: |                         |     13/414695 |   0%
                                                                                                                                100_1 _$a$x$z: |                         |      1/414695 |   0%
                                                                                                                              100_10_$a$c$d$x: |                         |      2/414695 |   0%
                                                                                                                            100_10_$a$c$d$x$z: |                         |      2/414695 |   0%
                                                                                                                                100_10_$a$d$v: |                         |     10/414695 |   0%
                                                                                                                                100_10_$a$d$x: |                         |     70/414695 |   0%
                                                                                                                              100_10_$a$d$x$x: |                         |      5/414695 |   0%
                                                                                                                              100_10_$a$d$x$z: |                         |     10/414695 |   0%
                                                                                                                              100_10_$a$q$d$x: |                         |      1/414695 |   0%
                                                                                                                            100_10_$a$q$d$x$x: |                         |      1/414695 |   0%
                                                                                                                                  100_10_$a$v: |                         |      1/414695 |   0%
                                                                                                                                  100_10_$a$x: |                         |      3/414695 |   0%
                                                                                                                                100_10_$a$x$x: |                         |      2/414695 |   0%
                                                                                                                                100_10_$a$x$z: |                         |      1/414695 |   0%
                                                                                                                                    100_3 _$a: |                         |   6700/414695 |   1%
                                                                                                                                  100_3 _$a$d: |                         |     17/414695 |   0%
                                                                                                                                  100_3 _$a$v: |                         |      1/414695 |   0%
                                                                                                                                  100_3 _$a$x: |                         |      2/414695 |   0%
                                                                                                                                    100_30_$a: |                         |  15863/414695 |   3%
                                                                                                                                  100_30_$a$d: |                         |     34/414695 |   0%
                                                                                                                                  100_30_$a$x: |                         |      2/414695 |   0%
                                                                                                                                    110_1 _$a: |                         |      1/414695 |   0%
                                                                                                                                  110_1 _$a$b: |                         |      3/414695 |   0%
                                                                                                                                110_1 _$a$b$b: |                         |      1/414695 |   0%
                                                                                                                              110_1 _$a$b$b$v: |                         |      1/414695 |   0%
                                                                                                                              110_1 _$a$b$b$x: |                         |      8/414695 |   0%
                                                                                                                                110_1 _$a$b$v: |                         |     46/414695 |   0%
                                                                                                                              110_1 _$a$b$v$v: |                         |      2/414695 |   0%
                                                                                                                                110_1 _$a$b$x: |                         |     44/414695 |   0%
                                                                                                                              110_1 _$a$b$x$v: |                         |     14/414695 |   0%
                                                                                                                              110_1 _$a$b$x$x: |                         |      7/414695 |   0%
                                                                                                                              110_1 _$a$b$x$y: |                         |      2/414695 |   0%
                                                                                                                                110_1 _$a$t$x: |                         |      3/414695 |   0%
                                                                                                                                    110_10_$a: |                         |      2/414695 |   0%
                                                                                                                                  110_10_$a$b: |                         |      5/414695 |   0%
                                                                                                                                110_10_$a$b$b: |                         |      3/414695 |   0%
                                                                                                                              110_10_$a$b$b$v: |                         |      3/414695 |   0%
                                                                                                                              110_10_$a$b$b$x: |                         |     14/414695 |   0%
                                                                                                                            110_10_$a$b$b$x$y: |                         |      1/414695 |   0%
                                                                                                                                110_10_$a$b$v: |                         |     12/414695 |   0%
                                                                                                                              110_10_$a$b$v$v: |                         |      1/414695 |   0%
                                                                                                                                110_10_$a$b$x: |                         |    358/414695 |   0%
                                                                                                                              110_10_$a$b$x$v: |                         |      2/414695 |   0%
                                                                                                                              110_10_$a$b$x$x: |                         |     29/414695 |   0%
                                                                                                                            110_10_$a$b$x$x$x: |                         |      2/414695 |   0%
                                                                                                                              110_10_$a$b$x$y: |                         |     19/414695 |   0%
                                                                                                                                110_10_$a$t$x: |                         |      2/414695 |   0%
                                                                                                                                    110_2 _$a: |                         |   5489/414695 |   1%
                                                                                                                                110_2 _$a$b$v: |                         |      2/414695 |   0%
                                                                                                                                  110_2 _$a$v: |                         |     70/414695 |   0%
                                                                                                                                110_2 _$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                  110_2 _$a$x: |                         |     78/414695 |   0%
                                                                                                                                110_2 _$a$x$v: |                         |     33/414695 |   0%
                                                                                                                              110_2 _$a$x$v$v: |                         |      1/414695 |   0%
                                                                                                                                110_2 _$a$x$x: |                         |     11/414695 |   0%
                                                                                                                              110_2 _$a$x$x$v: |                         |      3/414695 |   0%
                                                                                                                                110_2 _$a$x$y: |                         |      3/414695 |   0%
                                                                                                                                110_2 _$a$x$z: |                         |      3/414695 |   0%
                                                                                                                              110_2 _$a$x$z$x: |                         |      2/414695 |   0%
                                                                                                                            110_2 _$a$x$z$x$y: |                         |      1/414695 |   0%
                                                                                                                                  110_2 _$a$z: |                         |      3/414695 |   0%
                                                                                                                                110_2 _$a$z$x: |                         |      2/414695 |   0%
                                                                                                                                    110_20_$a: |                         |   2597/414695 |   0%
                                                                                                                                  110_20_$a$b: |                         |      1/414695 |   0%
                                                                                                                                110_20_$a$b$x: |                         |      3/414695 |   0%
                                                                                                                                  110_20_$a$t: |                         |      1/414695 |   0%
                                                                                                                                  110_20_$a$v: |                         |     26/414695 |   0%
                                                                                                                                110_20_$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                110_20_$a$v$x: |                         |      1/414695 |   0%
                                                                                                                                  110_20_$a$x: |                         |    252/414695 |   0%
                                                                                                                                110_20_$a$x$v: |                         |      4/414695 |   0%
                                                                                                                                110_20_$a$x$x: |                         |     55/414695 |   0%
                                                                                                                              110_20_$a$x$x$x: |                         |      1/414695 |   0%
                                                                                                                              110_20_$a$x$x$y: |                         |      1/414695 |   0%
                                                                                                                                110_20_$a$x$y: |                         |     22/414695 |   0%
                                                                                                                                110_20_$a$x$z: |                         |      3/414695 |   0%
                                                                                                                                  110_20_$a$z: |                         |     17/414695 |   0%
                                                                                                                                110_20_$a$z$x: |                         |      5/414695 |   0%
                                                                                                                              110_20_$a$z$x$y: |                         |      6/414695 |   0%
                                                                                                                                    111_2 _$a: |                         |      1/414695 |   0%
                                                                                                                                111_2 _$a$c$x: |                         |      1/414695 |   0%
                                                                                                                                  111_2 _$a$d: |                         |      2/414695 |   0%
                                                                                                                              111_2 _$a$d$c$x: |                         |      1/414695 |   0%
                                                                                                                                111_2 _$a$n$d: |                         |      1/414695 |   0%
                                                                                                                                  111_2 _$a$v: |                         |      1/414695 |   0%
                                                                                                                                  111_2 _$a$x: |                         |      2/414695 |   0%
                                                                                                                                    130_ 0_$a: |                         |     85/414695 |   0%
                                                                                                                                  130_ 0_$a$f: |                         |      2/414695 |   0%
                                                                                                                                  130_ 0_$a$l: |                         |      3/414695 |   0%
                                                                                                                                130_ 0_$a$l$v: |                         |      3/414695 |   0%
                                                                                                                                130_ 0_$a$l$x: |                         |      6/414695 |   0%
                                                                                                                              130_ 0_$a$l$x$x: |                         |      4/414695 |   0%
                                                                                                                                  130_ 0_$a$p: |                         |      3/414695 |   0%
                                                                                                                              130_ 0_$a$p$l$v: |                         |      1/414695 |   0%
                                                                                                                              130_ 0_$a$p$l$x: |                         |      3/414695 |   0%
                                                                                                                            130_ 0_$a$p$l$x$x: |                         |      1/414695 |   0%
                                                                                                                                130_ 0_$a$p$p: |                         |      1/414695 |   0%
                                                                                                                              130_ 0_$a$p$p$v: |                         |      1/414695 |   0%
                                                                                                                                130_ 0_$a$p$v: |                         |      6/414695 |   0%
                                                                                                                                130_ 0_$a$p$x: |                         |     35/414695 |   0%
                                                                                                                              130_ 0_$a$p$x$x: |                         |      2/414695 |   0%
                                                                                                                                  130_ 0_$a$v: |                         |    101/414695 |   0%
                                                                                                                                130_ 0_$a$v$v: |                         |      6/414695 |   0%
                                                                                                                                130_ 0_$a$v$x: |                         |     15/414695 |   0%
                                                                                                                                  130_ 0_$a$x: |                         |    160/414695 |   0%
                                                                                                                                130_ 0_$a$x$v: |                         |     13/414695 |   0%
                                                                                                                              130_ 0_$a$x$v$v: |                         |      1/414695 |   0%
                                                                                                                                130_ 0_$a$x$x: |                         |     28/414695 |   0%
                                                                                                                              130_ 0_$a$x$x$y: |                         |      9/414695 |   0%
                                                                                                                                    150_  _$a: |====                     |  76845/414695 |  18%
                                                                                                                                  150_  _$a$v: |                         |  16110/414695 |   3%
                                                                                                                                150_  _$a$v$v: |                         |    440/414695 |   0%
                                                                                                                                150_  _$a$v$x: |                         |    563/414695 |   0%
                                                                                                                              150_  _$a$v$x$v: |                         |      3/414695 |   0%
                                                                                                                                150_  _$a$v$y: |                         |      1/414695 |   0%
                                                                                                                                150_  _$a$v$z: |                         |     10/414695 |   0%
                                                                                                                                  150_  _$a$x: |=                        |  20366/414695 |   4%
                                                                                                                                150_  _$a$x$v: |                         |   3767/414695 |   0%
                                                                                                                              150_  _$a$x$v$v: |                         |     63/414695 |   0%
                                                                                                                            150_  _$a$x$v$v$v: |                         |      1/414695 |   0%
                                                                                                                              150_  _$a$x$v$x: |                         |     17/414695 |   0%
                                                                                                                                150_  _$a$x$x: |                         |   3415/414695 |   0%
                                                                                                                              150_  _$a$x$x$v: |                         |    491/414695 |   0%
                                                                                                                            150_  _$a$x$x$v$v: |                         |      2/414695 |   0%
                                                                                                                              150_  _$a$x$x$x: |                         |    216/414695 |   0%
                                                                                                                            150_  _$a$x$x$x$v: |                         |     32/414695 |   0%
                                                                                                                            150_  _$a$x$x$x$x: |                         |      7/414695 |   0%
                                                                                                                          150_  _$a$x$x$x$x$v: |                         |      3/414695 |   0%
                                                                                                                            150_  _$a$x$x$x$y: |                         |      2/414695 |   0%
                                                                                                                            150_  _$a$x$x$x$z: |                         |     29/414695 |   0%
                                                                                                                          150_  _$a$x$x$x$z$v: |                         |      2/414695 |   0%
                                                                                                                              150_  _$a$x$x$y: |                         |     64/414695 |   0%
                                                                                                                            150_  _$a$x$x$y$v: |                         |      4/414695 |   0%
                                                                                                                          150_  _$a$x$x$y$x$v: |                         |      1/414695 |   0%
                                                                                                                              150_  _$a$x$x$z: |                         |    591/414695 |   0%
                                                                                                                            150_  _$a$x$x$z$v: |                         |    151/414695 |   0%
                                                                                                                          150_  _$a$x$x$z$v$v: |                         |      3/414695 |   0%
                                                                                                                            150_  _$a$x$x$z$x: |                         |     17/414695 |   0%
                                                                                                                          150_  _$a$x$x$z$x$y: |                         |      2/414695 |   0%
                                                                                                                                150_  _$a$x$y: |                         |    295/414695 |   0%
                                                                                                                              150_  _$a$x$y$v: |                         |     30/414695 |   0%
                                                                                                                              150_  _$a$x$y$x: |                         |      4/414695 |   0%
                                                                                                                            150_  _$a$x$y$x$v: |                         |      1/414695 |   0%
                                                                                                                                150_  _$a$x$z: |                         |   9394/414695 |   2%
                                                                                                                              150_  _$a$x$z$v: |                         |   1777/414695 |   0%
                                                                                                                            150_  _$a$x$z$v$v: |                         |    113/414695 |   0%
                                                                                                                            150_  _$a$x$z$v$x: |                         |      1/414695 |   0%
                                                                                                                              150_  _$a$x$z$x: |                         |    512/414695 |   0%
                                                                                                                            150_  _$a$x$z$x$v: |                         |     55/414695 |   0%
                                                                                                                          150_  _$a$x$z$x$v$v: |                         |      4/414695 |   0%
                                                                                                                            150_  _$a$x$z$x$x: |                         |     10/414695 |   0%
                                                                                                                          150_  _$a$x$z$x$x$v: |                         |      6/414695 |   0%
                                                                                                                        150_  _$a$x$z$x$x$v$v: |                         |      4/414695 |   0%
                                                                                                                            150_  _$a$x$z$x$y: |                         |    207/414695 |   0%
                                                                                                                          150_  _$a$x$z$x$y$v: |                         |     27/414695 |   0%
                                                                                                                        150_  _$a$x$z$x$y$x$v: |                         |      1/414695 |   0%
                                                                                                                            150_  _$a$x$z$x$z: |                         |      1/414695 |   0%
                                                                                                                              150_  _$a$x$z$z: |                         |    124/414695 |   0%
                                                                                                                            150_  _$a$x$z$z$v: |                         |     12/414695 |   0%
                                                                                                                          150_  _$a$x$z$z$v$v: |                         |      2/414695 |   0%
                                                                                                                            150_  _$a$x$z$z$x: |                         |      4/414695 |   0%
                                                                                                                          150_  _$a$x$z$z$x$y: |                         |      3/414695 |   0%
                                                                                                                        150_  _$a$x$z$z$x$y$v: |                         |      1/414695 |   0%
                                                                                                                                  150_  _$a$y: |                         |   1065/414695 |   0%
                                                                                                                                150_  _$a$y$v: |                         |    888/414695 |   0%
                                                                                                                              150_  _$a$y$v$v: |                         |     31/414695 |   0%
                                                                                                                                150_  _$a$y$x: |                         |    631/414695 |   0%
                                                                                                                              150_  _$a$y$x$v: |                         |     61/414695 |   0%
                                                                                                                              150_  _$a$y$x$x: |                         |     15/414695 |   0%
                                                                                                                              150_  _$a$y$x$z: |                         |      1/414695 |   0%
                                                                                                                                150_  _$a$y$z: |                         |      2/414695 |   0%
                                                                                                                              150_  _$a$y$z$v: |                         |      1/414695 |   0%
                                                                                                                                  150_  _$a$z: |=                        |  24696/414695 |   5%
                                                                                                                                150_  _$a$z$v: |                         |   5758/414695 |   1%
                                                                                                                              150_  _$a$z$v$v: |                         |    294/414695 |   0%
                                                                                                                            150_  _$a$z$v$v$v: |                         |      5/414695 |   0%
                                                                                                                              150_  _$a$z$v$x: |                         |      7/414695 |   0%
                                                                                                                                150_  _$a$z$x: |                         |   2583/414695 |   0%
                                                                                                                              150_  _$a$z$x$v: |                         |    132/414695 |   0%
                                                                                                                            150_  _$a$z$x$v$v: |                         |      8/414695 |   0%
                                                                                                                              150_  _$a$z$x$x: |                         |     18/414695 |   0%
                                                                                                                            150_  _$a$z$x$x$v: |                         |      2/414695 |   0%
                                                                                                                            150_  _$a$z$x$x$y: |                         |      1/414695 |   0%
                                                                                                                              150_  _$a$z$x$y: |                         |   1015/414695 |   0%
                                                                                                                            150_  _$a$z$x$y$v: |                         |     44/414695 |   0%
                                                                                                                                150_  _$a$z$y: |                         |    206/414695 |   0%
                                                                                                                              150_  _$a$z$y$v: |                         |     13/414695 |   0%
                                                                                                                              150_  _$a$z$y$x: |                         |     24/414695 |   0%
                                                                                                                                150_  _$a$z$z: |                         |    980/414695 |   0%
                                                                                                                              150_  _$a$z$z$v: |                         |    655/414695 |   0%
                                                                                                                            150_  _$a$z$z$v$v: |                         |     15/414695 |   0%
                                                                                                                              150_  _$a$z$z$x: |                         |     97/414695 |   0%
                                                                                                                            150_  _$a$z$z$x$v: |                         |      4/414695 |   0%
                                                                                                                            150_  _$a$z$z$x$x: |                         |      3/414695 |   0%
                                                                                                                            150_  _$a$z$z$x$y: |                         |     45/414695 |   0%
                                                                                                                          150_  _$a$z$z$x$y$v: |                         |      3/414695 |   0%
                                                                                                                              150_  _$a$z$z$y: |                         |      7/414695 |   0%
                                                                                                                            150_  _$a$z$z$y$v: |                         |     11/414695 |   0%
                                                                                                                            150_  _$a$z$z$y$x: |                         |      2/414695 |   0%
                                                                                                                                    150_ 0_$a: |=====                    |  95085/414695 |  22%
                                                                                                                                  150_ 0_$a$v: |                         |    964/414695 |   0%
                                                                                                                                150_ 0_$a$v$v: |                         |     27/414695 |   0%
                                                                                                                                150_ 0_$a$v$x: |                         |      4/414695 |   0%
                                                                                                                                  150_ 0_$a$x: |=                        |  22723/414695 |   5%
                                                                                                                                150_ 0_$a$x$v: |                         |     15/414695 |   0%
                                                                                                                                150_ 0_$a$x$x: |                         |   3587/414695 |   0%
                                                                                                                              150_ 0_$a$x$x$x: |                         |     92/414695 |   0%
                                                                                                                            150_ 0_$a$x$x$x$x: |                         |      2/414695 |   0%
                                                                                                                              150_ 0_$a$x$x$y: |                         |      5/414695 |   0%
                                                                                                                              150_ 0_$a$x$x$z: |                         |     14/414695 |   0%
                                                                                                                                150_ 0_$a$x$y: |                         |    202/414695 |   0%
                                                                                                                              150_ 0_$a$x$y$v: |                         |      2/414695 |   0%
                                                                                                                              150_ 0_$a$x$y$x: |                         |      2/414695 |   0%
                                                                                                                                150_ 0_$a$x$z: |                         |    803/414695 |   0%
                                                                                                                              150_ 0_$a$x$z$x: |                         |      4/414695 |   0%
                                                                                                                              150_ 0_$a$x$z$z: |                         |      3/414695 |   0%
                                                                                                                                  150_ 0_$a$y: |                         |   1348/414695 |   0%
                                                                                                                                150_ 0_$a$y$v: |                         |      9/414695 |   0%
                                                                                                                                150_ 0_$a$y$x: |                         |     42/414695 |   0%
                                                                                                                              150_ 0_$a$y$x$x: |                         |      1/414695 |   0%
                                                                                                                                150_ 0_$a$y$z: |                         |      2/414695 |   0%
                                                                                                                              150_ 0_$a$y$z$z: |                         |      1/414695 |   0%
                                                                                                                                  150_ 0_$a$z: |                         |  11709/414695 |   2%
                                                                                                                                150_ 0_$a$z$v: |                         |     14/414695 |   0%
                                                                                                                                150_ 0_$a$z$x: |                         |    426/414695 |   0%
                                                                                                                              150_ 0_$a$z$x$v: |                         |      2/414695 |   0%
                                                                                                                              150_ 0_$a$z$x$x: |                         |      4/414695 |   0%
                                                                                                                            150_ 0_$a$z$x$x$x: |                         |      1/414695 |   0%
                                                                                                                              150_ 0_$a$z$x$y: |                         |    124/414695 |   0%
                                                                                                                                150_ 0_$a$z$y: |                         |      6/414695 |   0%
                                                                                                                              150_ 0_$a$z$y$x: |                         |      1/414695 |   0%
                                                                                                                                150_ 0_$a$z$z: |                         |    105/414695 |   0%
                                                                                                                              150_ 0_$a$z$z$x: |                         |     25/414695 |   0%
                                                                                                                              150_ 0_$a$z$z$y: |                         |      1/414695 |   0%
                                                                                                                                    151_  _$a: |==                       |  45831/414695 |  11%
                                                                                                                                  151_  _$a$v: |                         |   1884/414695 |   0%
                                                                                                                                151_  _$a$v$v: |                         |     41/414695 |   0%
                                                                                                                              151_  _$a$v$v$v: |                         |      1/414695 |   0%
                                                                                                                              151_  _$a$v$v$x: |                         |      1/414695 |   0%
                                                                                                                                  151_  _$a$x: |                         |   2489/414695 |   0%
                                                                                                                                151_  _$a$x$v: |                         |    545/414695 |   0%
                                                                                                                              151_  _$a$x$v$v: |                         |     10/414695 |   0%
                                                                                                                                151_  _$a$x$x: |                         |    646/414695 |   0%
                                                                                                                              151_  _$a$x$x$v: |                         |     12/414695 |   0%
                                                                                                                            151_  _$a$x$x$v$v: |                         |      1/414695 |   0%
                                                                                                                              151_  _$a$x$x$x: |                         |      6/414695 |   0%
                                                                                                                            151_  _$a$x$x$x$y: |                         |      3/414695 |   0%
                                                                                                                              151_  _$a$x$x$y: |                         |     24/414695 |   0%
                                                                                                                              151_  _$a$x$x$z: |                         |      2/414695 |   0%
                                                                                                                            151_  _$a$x$x$z$v: |                         |      1/414695 |   0%
                                                                                                                                151_  _$a$x$y: |                         |   2051/414695 |   0%
                                                                                                                              151_  _$a$x$y$v: |                         |    431/414695 |   0%
                                                                                                                            151_  _$a$x$y$v$v: |                         |      2/414695 |   0%
                                                                                                                              151_  _$a$x$y$x: |                         |    130/414695 |   0%
                                                                                                                            151_  _$a$x$y$x$v: |                         |      6/414695 |   0%
                                                                                                                            151_  _$a$x$y$x$x: |                         |      3/414695 |   0%
                                                                                                                          151_  _$a$x$y$x$x$v: |                         |      1/414695 |   0%
                                                                                                                            151_  _$a$x$y$x$z: |                         |      1/414695 |   0%
                                                                                                                                151_  _$a$x$z: |                         |    469/414695 |   0%
                                                                                                                              151_  _$a$x$z$v: |                         |     11/414695 |   0%
                                                                                                                              151_  _$a$x$z$x: |                         |      2/414695 |   0%
                                                                                                                            151_  _$a$x$z$x$y: |                         |      2/414695 |   0%
                                                                                                                                  151_  _$a$y: |                         |      2/414695 |   0%
                                                                                                                                  151_  _$a$z: |                         |      1/414695 |   0%
                                                                                                                                    151_ 0_$a: |                         |    161/414695 |   0%
                                                                                                                                  151_ 0_$a$v: |                         |     33/414695 |   0%
                                                                                                                                151_ 0_$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                  151_ 0_$a$x: |                         |   2913/414695 |   0%
                                                                                                                                151_ 0_$a$x$v: |                         |     41/414695 |   0%
                                                                                                                              151_ 0_$a$x$v$v: |                         |      3/414695 |   0%
                                                                                                                                151_ 0_$a$x$x: |                         |    488/414695 |   0%
                                                                                                                              151_ 0_$a$x$x$v: |                         |      7/414695 |   0%
                                                                                                                              151_ 0_$a$x$x$x: |                         |     23/414695 |   0%
                                                                                                                              151_ 0_$a$x$x$y: |                         |      7/414695 |   0%
                                                                                                                                151_ 0_$a$x$y: |                         |   5646/414695 |   1%
                                                                                                                              151_ 0_$a$x$y$v: |                         |    145/414695 |   0%
                                                                                                                              151_ 0_$a$x$y$x: |                         |    235/414695 |   0%
                                                                                                                            151_ 0_$a$x$y$x$v: |                         |      5/414695 |   0%
                                                                                                                            151_ 0_$a$x$y$x$x: |                         |      1/414695 |   0%
                                                                                                                                151_ 0_$a$x$z: |                         |     69/414695 |   0%
                                                                                                                                    155_  _$a: |                         |    667/414695 |   0%
                                                                                                                                    180_  _$x: |                         |   2272/414695 |   0%
                                                                                                                                  180_  _$x$v: |                         |     71/414695 |   0%
                                                                                                                                180_  _$x$v$v: |                         |      3/414695 |   0%
                                                                                                                                  180_  _$x$x: |                         |    581/414695 |   0%
                                                                                                                                180_  _$x$x$v: |                         |      5/414695 |   0%
                                                                                                                                180_  _$x$x$x: |                         |     22/414695 |   0%
                                                                                                                                180_  _$x$x$y: |                         |     17/414695 |   0%
                                                                                                                                  180_  _$x$y: |                         |     79/414695 |   0%
                                                                                                                                180_  _$x$y$v: |                         |      3/414695 |   0%
                                                                                                                              180_  _$x$y$v$v: |                         |      2/414695 |   0%
                                                                                                                                180_  _$x$y$x: |                         |      5/414695 |   0%
                                                                                                                              180_  _$x$y$x$v: |                         |      1/414695 |   0%
                                                                                                                              180_  _$x$y$x$x: |                         |      1/414695 |   0%
                                                                                                                                  180_  _$x$z: |                         |      6/414695 |   0%
                                                                                                                                180_  _$x$z$x: |                         |      1/414695 |   0%
                                                                                                                                    181_  _$z: |                         |      1/414695 |   0%
                                                                                                                                  181_  _$z$x: |                         |      1/414695 |   0%
                                                                                                                                    182_  _$y: |                         |     34/414695 |   0%
                                                                                                                                    185_  _$v: |                         |    406/414695 |   0%
                                                                                                                                  185_  _$v$v: |                         |     80/414695 |   0%
                                                                                                                                  185_  _$v$x: |                         |     13/414695 |   0%
                                                                                                                                    260_  _$a: |                         |      1/414695 |   0%
                                                                                                                              260_  _$a$i$a$i: |                         |      1/414695 |   0%
                                                                                                                          260_  _$a$i$a$i$a$i: |                         |      1/414695 |   0%
                                                                                                                                    260_  _$i: |                         |      4/414695 |   0%
                                                                                                                                  260_  _$i$a: |                         |    165/414695 |   0%
                                                                                                                                260_  _$i$a$i: |                         |    144/414695 |   0%
                                                                                                                              260_  _$i$a$i$a: |                         |    360/414695 |   0%
                                                                                                                            260_  _$i$a$i$a$i: |                         |     26/414695 |   0%
                                                                                                                          260_  _$i$a$i$a$i$a: |                         |     31/414695 |   0%
                                                                                                                        260_  _$i$a$i$a$i$a$i: |                         |      7/414695 |   0%
                                                                                                                      260_  _$i$a$i$a$i$a$i$a: |                         |     16/414695 |   0%
                                                                                                                    260_  _$i$a$i$a$i$a$i$a$i: |                         |      4/414695 |   0%
                                                                                                                  260_  _$i$a$i$a$i$a$i$a$i$a: |                         |      6/414695 |   0%
                                                                                                              260_  _$i$a$i$a$i$a$i$a$i$a$i$a: |                         |      1/414695 |   0%
                                                                                                                      260_  _$i$a$i$a$i$i$a$i: |                         |      1/414695 |   0%
                                                                                                                    260_  _$i$i$a$i$a$i$a$i$a: |                         |      2/414695 |   0%
                                                                                                                                    360_  _$a: |                         |      1/414695 |   0%
                                                                                                                          360_  _$a$a$i$a$i$a: |                         |      1/414695 |   0%
                                                                                                                                  360_  _$a$i: |                         |      7/414695 |   0%
                                                                                                                              360_  _$a$i$a$i: |                         |      1/414695 |   0%
                                                                                                                                    360_  _$i: |                         |    630/414695 |   0%
                                                                                                                                  360_  _$i$a: |                         |    357/414695 |   0%
                                                                                                                                360_  _$i$a$a: |                         |      5/414695 |   0%
                                                                                                                              360_  _$i$a$a$a: |                         |      4/414695 |   0%
                                                                                                                            360_  _$i$a$a$i$a: |                         |      1/414695 |   0%
                                                                                                                                360_  _$i$a$i: |                         |    779/414695 |   0%
                                                                                                                              360_  _$i$a$i$a: |                         |   2044/414695 |   0%
                                                                                                                            360_  _$i$a$i$a$i: |                         |    142/414695 |   0%
                                                                                                                          360_  _$i$a$i$a$i$a: |                         |    162/414695 |   0%
                                                                                                                        360_  _$i$a$i$a$i$a$i: |                         |     31/414695 |   0%
                                                                                                                      360_  _$i$a$i$a$i$a$i$a: |                         |    115/414695 |   0%
                                                                                                                    360_  _$i$a$i$a$i$a$i$a$i: |                         |      9/414695 |   0%
                                                                                                                  360_  _$i$a$i$a$i$a$i$a$i$a: |                         |     18/414695 |   0%
                                                                                                                360_  _$i$a$i$a$i$a$i$a$i$a$a: |                         |      1/414695 |   0%
                                                                                                                360_  _$i$a$i$a$i$a$i$a$i$a$i: |                         |      5/414695 |   0%
                                                                                                              360_  _$i$a$i$a$i$a$i$a$i$a$i$a: |                         |     15/414695 |   0%
                                                                                                          360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a: |                         |      4/414695 |   0%
                                                                                                        360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i: |                         |      2/414695 |   0%
                                                                                                      360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a: |                         |      1/414695 |   0%
                                                                                                  360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a: |                         |      1/414695 |   0%
                                                                                                360_  _$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i$a$i: |                         |      1/414695 |   0%
                                                                                                                            360_  _$i$a$i$i$a: |                         |      1/414695 |   0%
                                                                                                                            360_  _$i$a$i$i$i: |                         |      1/414695 |   0%
                                                                                                                                360_  _$i$i$a: |                         |      1/414695 |   0%
                                                                                                                          360_  _$i$i$a$i$a$i: |                         |      1/414695 |   0%
                                                                                                                                    400_0 _$a: |                         |      2/414695 |   0%
                                                                                                                            400_0 _$a$b$c$d$x: |                         |      1/414695 |   0%
                                                                                                                                400_0 _$a$c$d: |                         |      7/414695 |   0%
                                                                                                                              400_0 _$a$c$d$x: |                         |     10/414695 |   0%
                                                                                                                                400_0 _$a$c$v: |                         |      1/414695 |   0%
                                                                                                                                400_0 _$a$c$x: |                         |      8/414695 |   0%
                                                                                                                              400_0 _$a$c$x$x: |                         |      1/414695 |   0%
                                                                                                                                  400_0 _$a$d: |                         |      4/414695 |   0%
                                                                                                                                400_0 _$a$d$t: |                         |      1/414695 |   0%
                                                                                                                                400_0 _$a$d$x: |                         |      1/414695 |   0%
                                                                                                                                  400_0 _$a$t: |                         |      1/414695 |   0%
                                                                                                                                  400_0 _$a$v: |                         |      2/414695 |   0%
                                                                                                                                400_0 _$a$v$v: |                         |      2/414695 |   0%
                                                                                                                                  400_0 _$a$x: |                         |     18/414695 |   0%
                                                                                                                                400_0 _$a$x$v: |                         |      2/414695 |   0%
                                                                                                                          400_0 _$w$a$b$c$d$k: |                         |      2/414695 |   0%
                                                                                                                          400_0 _$w$a$b$c$d$v: |                         |      1/414695 |   0%
                                                                                                                            400_0 _$w$a$c$d$x: |                         |      4/414695 |   0%
                                                                                                                              400_0 _$w$a$c$v: |                         |      1/414695 |   0%
                                                                                                                              400_0 _$w$a$c$x: |                         |     13/414695 |   0%
                                                                                                                                400_0 _$w$a$d: |                         |      1/414695 |   0%
                                                                                                                              400_0 _$w$a$d$x: |                         |      1/414695 |   0%
                                                                                                                                400_0 _$w$a$k: |                         |      1/414695 |   0%
                                                                                                                                400_0 _$w$a$v: |                         |      1/414695 |   0%
                                                                                                                                400_0 _$w$a$x: |                         |      4/414695 |   0%
                                                                                                                            400_0 _$w$a$x$x$v: |                         |      1/414695 |   0%
                                                                                                                            400_00_$a$b$c$d$x: |                         |      2/414695 |   0%
                                                                                                                                400_00_$a$c$x: |                         |      7/414695 |   0%
                                                                                                                                400_00_$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                  400_00_$a$x: |                         |     34/414695 |   0%
                                                                                                                                400_00_$a$x$x: |                         |      4/414695 |   0%
                                                                                                                          400_00_$w$a$b$c$d$v: |                         |      1/414695 |   0%
                                                                                                                          400_00_$w$a$b$c$d$x: |                         |      1/414695 |   0%
                                                                                                                            400_00_$w$a$c$d$v: |                         |      2/414695 |   0%
                                                                                                                                400_00_$w$a$x: |                         |      4/414695 |   0%
                                                                                                                              400_00_$w$a$x$x: |                         |      1/414695 |   0%
                                                                                                                                    400_1 _$a: |                         |      2/414695 |   0%
                                                                                                                                400_1 _$a$c$d: |                         |      1/414695 |   0%
                                                                                                                              400_1 _$a$c$d$t: |                         |      1/414695 |   0%
                                                                                                                              400_1 _$a$c$d$x: |                         |      1/414695 |   0%
                                                                                                                                  400_1 _$a$d: |                         |     16/414695 |   0%
                                                                                                                                400_1 _$a$d$t: |                         |     24/414695 |   0%
                                                                                                                                400_1 _$a$d$v: |                         |     10/414695 |   0%
                                                                                                                              400_1 _$a$d$v$v: |                         |      1/414695 |   0%
                                                                                                                                400_1 _$a$d$x: |                         |     29/414695 |   0%
                                                                                                                              400_1 _$a$d$x$x: |                         |     12/414695 |   0%
                                                                                                                              400_1 _$a$q$d$x: |                         |      1/414695 |   0%
                                                                                                                            400_1 _$w$a$c$d$x: |                         |      1/414695 |   0%
                                                                                                                              400_1 _$w$a$d$k: |                         |      9/414695 |   0%
                                                                                                                              400_1 _$w$a$d$v: |                         |     20/414695 |   0%
                                                                                                                            400_1 _$w$a$d$v$x: |                         |      2/414695 |   0%
                                                                                                                              400_1 _$w$a$d$x: |                         |     16/414695 |   0%
                                                                                                                            400_1 _$w$a$d$x$v: |                         |      3/414695 |   0%
                                                                                                                            400_1 _$w$a$d$x$x: |                         |     25/414695 |   0%
                                                                                                                                    400_10_$a: |                         |      2/414695 |   0%
                                                                                                                                400_10_$a$d$t: |                         |      8/414695 |   0%
                                                                                                                                400_10_$a$d$x: |                         |      7/414695 |   0%
                                                                                                                                  400_10_$a$t: |                         |      1/414695 |   0%
                                                                                                                                  400_10_$a$x: |                         |      1/414695 |   0%
                                                                                                                          400_10_$w$a$b$c$d$x: |                         |      1/414695 |   0%
                                                                                                                              400_10_$w$a$d$k: |                         |      1/414695 |   0%
                                                                                                                              400_10_$w$a$d$v: |                         |      2/414695 |   0%
                                                                                                                              400_10_$w$a$d$x: |                         |      7/414695 |   0%
                                                                                                                                400_10_$w$a$v: |                         |      1/414695 |   0%
                                                                                                                                    400_3 _$a: |                         |   2998/414695 |   0%
                                                                                                                                400_3 _$a$c$d: |                         |      2/414695 |   0%
                                                                                                                                  400_3 _$a$d: |                         |     27/414695 |   0%
                                                                                                                                  400_3 _$w$a: |                         |    140/414695 |   0%
                                                                                                                                    400_30_$a: |                         |   7558/414695 |   1%
                                                                                                                                  400_30_$a$c: |                         |      1/414695 |   0%
                                                                                                                                400_30_$a$c$d: |                         |     10/414695 |   0%
                                                                                                                                  400_30_$a$d: |                         |     24/414695 |   0%
                                                                                                                                  400_30_$w$a: |                         |    275/414695 |   0%
                                                                                                                                400_30_$w$a$c: |                         |      5/414695 |   0%
                                                                                                                                    410_0 _$a: |                         |      2/414695 |   0%
                                                                                                                                    410_1 _$a: |                         |      1/414695 |   0%
                                                                                                                                  410_1 _$a$b: |                         |     38/414695 |   0%
                                                                                                                                410_1 _$a$b$v: |                         |      5/414695 |   0%
                                                                                                                                410_1 _$a$b$x: |                         |     18/414695 |   0%
                                                                                                                              410_1 _$a$b$x$v: |                         |      4/414695 |   0%
                                                                                                                                410_1 _$w$a$b: |                         |      1/414695 |   0%
                                                                                                                            410_1 _$w$a$b$b$x: |                         |      7/414695 |   0%
                                                                                                                              410_1 _$w$a$b$v: |                         |      7/414695 |   0%
                                                                                                                            410_1 _$w$a$b$v$v: |                         |      1/414695 |   0%
                                                                                                                              410_1 _$w$a$b$x: |                         |      9/414695 |   0%
                                                                                                                            410_1 _$w$a$b$x$v: |                         |      2/414695 |   0%
                                                                                                                            410_1 _$w$a$b$x$x: |                         |      1/414695 |   0%
                                                                                                                            410_1 _$w$a$b$x$y: |                         |      1/414695 |   0%
                                                                                                                                    410_10_$a: |                         |      4/414695 |   0%
                                                                                                                                  410_10_$a$b: |                         |      6/414695 |   0%
                                                                                                                                410_10_$a$b$b: |                         |      2/414695 |   0%
                                                                                                                              410_10_$a$b$b$x: |                         |      2/414695 |   0%
                                                                                                                                410_10_$a$b$t: |                         |      1/414695 |   0%
                                                                                                                                410_10_$a$b$v: |                         |      1/414695 |   0%
                                                                                                                                410_10_$a$b$x: |                         |     29/414695 |   0%
                                                                                                                              410_10_$a$b$x$y: |                         |      2/414695 |   0%
                                                                                                                                410_10_$a$t$x: |                         |      1/414695 |   0%
                                                                                                                                410_10_$w$a$b: |                         |      3/414695 |   0%
                                                                                                                              410_10_$w$a$b$v: |                         |      3/414695 |   0%
                                                                                                                              410_10_$w$a$b$x: |                         |      5/414695 |   0%
                                                                                                                                    410_2 _$a: |                         |   3792/414695 |   0%
                                                                                                                                  410_2 _$a$b: |                         |      1/414695 |   0%
                                                                                                                                  410_2 _$a$v: |                         |     13/414695 |   0%
                                                                                                                                  410_2 _$a$x: |                         |     15/414695 |   0%
                                                                                                                                410_2 _$a$x$v: |                         |      1/414695 |   0%
                                                                                                                                410_2 _$a$x$y: |                         |      1/414695 |   0%
                                                                                                                                  410_2 _$a$y: |                         |      1/414695 |   0%
                                                                                                                                  410_2 _$a$z: |                         |      1/414695 |   0%
                                                                                                                                  410_2 _$w$a: |                         |    210/414695 |   0%
                                                                                                                                410_2 _$w$a$b: |                         |      1/414695 |   0%
                                                                                                                                410_2 _$w$a$v: |                         |      9/414695 |   0%
                                                                                                                              410_2 _$w$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                410_2 _$w$a$x: |                         |      9/414695 |   0%
                                                                                                                              410_2 _$w$a$x$v: |                         |      3/414695 |   0%
                                                                                                                              410_2 _$w$a$x$x: |                         |      1/414695 |   0%
                                                                                                                                    410_20_$a: |                         |   1841/414695 |   0%
                                                                                                                                  410_20_$a$b: |                         |      2/414695 |   0%
                                                                                                                                  410_20_$a$k: |                         |      1/414695 |   0%
                                                                                                                                  410_20_$a$v: |                         |      1/414695 |   0%
                                                                                                                                  410_20_$a$x: |                         |     35/414695 |   0%
                                                                                                                                410_20_$a$x$x: |                         |      3/414695 |   0%
                                                                                                                                410_20_$a$x$y: |                         |      1/414695 |   0%
                                                                                                                                  410_20_$a$z: |                         |      1/414695 |   0%
                                                                                                                                  410_20_$w$a: |                         |    107/414695 |   0%
                                                                                                                                410_20_$w$a$b: |                         |      1/414695 |   0%
                                                                                                                                410_20_$w$a$k: |                         |      3/414695 |   0%
                                                                                                                                410_20_$w$a$v: |                         |      6/414695 |   0%
                                                                                                                                410_20_$w$a$x: |                         |     27/414695 |   0%
                                                                                                                              410_20_$w$a$x$v: |                         |      1/414695 |   0%
                                                                                                                              410_20_$w$a$x$x: |                         |     10/414695 |   0%
                                                                                                                              410_20_$w$a$x$y: |                         |      2/414695 |   0%
                                                                                                                              410_20_$w$a$z$x: |                         |      2/414695 |   0%
                                                                                                                                    411_2 _$a: |                         |      1/414695 |   0%
                                                                                                                                  411_2 _$a$c: |                         |      1/414695 |   0%
                                                                                                                                  411_2 _$a$d: |                         |      1/414695 |   0%
                                                                                                                                411_2 _$a$d$c: |                         |      1/414695 |   0%
                                                                                                                                411_2 _$a$n$d: |                         |      1/414695 |   0%
                                                                                                                                    411_20_$a: |                         |      1/414695 |   0%
                                                                                                                                411_20_$w$a$x: |                         |      1/414695 |   0%
                                                                                                                                    430_ 0_$a: |                         |     32/414695 |   0%
                                                                                                                                  430_ 0_$a$f: |                         |      1/414695 |   0%
                                                                                                                                430_ 0_$a$l$v: |                         |      4/414695 |   0%
                                                                                                                                430_ 0_$a$l$x: |                         |      2/414695 |   0%
                                                                                                                              430_ 0_$a$l$x$x: |                         |      1/414695 |   0%
                                                                                                                                430_ 0_$a$p$v: |                         |      3/414695 |   0%
                                                                                                                                430_ 0_$a$p$x: |                         |     25/414695 |   0%
                                                                                                                                  430_ 0_$a$v: |                         |     25/414695 |   0%
                                                                                                                                  430_ 0_$a$x: |                         |     79/414695 |   0%
                                                                                                                                430_ 0_$a$x$v: |                         |      8/414695 |   0%
                                                                                                                                430_ 0_$a$x$x: |                         |      9/414695 |   0%
                                                                                                                                430_ 0_$a$x$y: |                         |      1/414695 |   0%
                                                                                                                                  430_ 0_$w$a: |                         |      1/414695 |   0%
                                                                                                                                430_ 0_$w$a$g: |                         |      1/414695 |   0%
                                                                                                                              430_ 0_$w$a$p$v: |                         |      1/414695 |   0%
                                                                                                                              430_ 0_$w$a$p$x: |                         |      1/414695 |   0%
                                                                                                                                430_ 0_$w$a$v: |                         |      9/414695 |   0%
                                                                                                                              430_ 0_$w$a$v$v: |                         |      1/414695 |   0%
                                                                                                                              430_ 0_$w$a$v$x: |                         |      1/414695 |   0%
                                                                                                                                430_ 0_$w$a$x: |                         |     15/414695 |   0%
                                                                                                                              430_ 0_$w$a$x$v: |                         |      1/414695 |   0%
                                                                                                                            430_ 0_$w$a$x$v$y: |                         |      1/414695 |   0%
                                                                                                                              430_ 0_$w$a$x$x: |                         |      3/414695 |   0%
                                                                                                                                    450_  _$a: |===                      |  49979/414695 |  12%
                                                                                                                                  450_  _$a$v: |                         |    404/414695 |   0%
                                                                                                                                450_  _$a$v$v: |                         |     22/414695 |   0%
                                                                                                                                450_  _$a$v$x: |                         |      2/414695 |   0%
                                                                                                                                  450_  _$a$x: |                         |   4088/414695 |   0%
                                                                                                                                450_  _$a$x$v: |                         |     19/414695 |   0%
                                                                                                                              450_  _$a$x$v$v: |                         |      4/414695 |   0%
                                                                                                                                450_  _$a$x$x: |                         |    106/414695 |   0%
                                                                                                                              450_  _$a$x$x$x: |                         |      4/414695 |   0%
                                                                                                                                450_  _$a$x$y: |                         |     17/414695 |   0%
                                                                                                                                450_  _$a$x$z: |                         |     13/414695 |   0%
                                                                                                                              450_  _$a$x$z$z: |                         |      1/414695 |   0%
                                                                                                                                  450_  _$a$y: |                         |     48/414695 |   0%
                                                                                                                                450_  _$a$y$v: |                         |      6/414695 |   0%
                                                                                                                                450_  _$a$y$x: |                         |     10/414695 |   0%
                                                                                                                                  450_  _$a$z: |                         |    771/414695 |   0%
                                                                                                                                450_  _$a$z$v: |                         |      1/414695 |   0%
                                                                                                                                450_  _$a$z$x: |                         |      8/414695 |   0%
                                                                                                                              450_  _$a$z$x$x: |                         |      1/414695 |   0%
                                                                                                                              450_  _$a$z$x$y: |                         |      4/414695 |   0%
                                                                                                                                450_  _$a$z$z: |                         |      3/414695 |   0%
                                                                                                                                  450_  _$w$a: |                         |  10919/414695 |   2%
                                                                                                                                450_  _$w$a$v: |                         |    372/414695 |   0%
                                                                                                                              450_  _$w$a$v$v: |                         |      2/414695 |   0%
                                                                                                                              450_  _$w$a$v$x: |                         |      8/414695 |   0%
                                                                                                                                450_  _$w$a$x: |                         |    879/414695 |   0%
                                                                                                                              450_  _$w$a$x$v: |                         |     17/414695 |   0%
                                                                                                                              450_  _$w$a$x$x: |                         |     40/414695 |   0%
                                                                                                                            450_  _$w$a$x$x$x: |                         |      9/414695 |   0%
                                                                                                                              450_  _$w$a$x$y: |                         |      9/414695 |   0%
                                                                                                                              450_  _$w$a$x$z: |                         |     14/414695 |   0%
                                                                                                                                450_  _$w$a$y: |                         |    121/414695 |   0%
                                                                                                                              450_  _$w$a$y$x: |                         |      5/414695 |   0%
                                                                                                                              450_  _$w$a$y$z: |                         |     65/414695 |   0%
                                                                                                                                450_  _$w$a$z: |                         |    150/414695 |   0%
                                                                                                                              450_  _$w$a$z$v: |                         |      2/414695 |   0%
                                                                                                                              450_  _$w$a$z$x: |                         |     28/414695 |   0%
                                                                                                                            450_  _$w$a$z$x$v: |                         |      1/414695 |   0%
                                                                                                                            450_  _$w$a$z$x$y: |                         |     43/414695 |   0%
                                                                                                                              450_  _$w$a$z$y: |                         |     51/414695 |   0%
                                                                                                                              450_  _$w$a$z$z: |                         |      4/414695 |   0%
                                                                                                                            450_  _$w$a$z$z$v: |                         |      3/414695 |   0%
                                                                                                                            450_  _$w$a$z$z$x: |                         |      1/414695 |   0%
                                                                                                                                    450_ 0_$a: |===                      |  53090/414695 |  12%
                                                                                                                                  450_ 0_$a$v: |                         |     35/414695 |   0%
                                                                                                                                450_ 0_$a$v$v: |                         |      5/414695 |   0%
                                                                                                                                  450_ 0_$a$x: |                         |   4267/414695 |   1%
                                                                                                                                450_ 0_$a$x$x: |                         |    201/414695 |   0%
                                                                                                                              450_ 0_$a$x$x$x: |                         |      2/414695 |   0%
                                                                                                                              450_ 0_$a$x$x$z: |                         |      1/414695 |   0%
                                                                                                                                450_ 0_$a$x$y: |                         |     16/414695 |   0%
                                                                                                                              450_ 0_$a$x$y$x: |                         |      1/414695 |   0%
                                                                                                                                450_ 0_$a$x$z: |                         |     30/414695 |   0%
                                                                                                                                  450_ 0_$a$y: |                         |     39/414695 |   0%
                                                                                                                                450_ 0_$a$y$x: |                         |      2/414695 |   0%
                                                                                                                              450_ 0_$a$y$x$x: |                         |      1/414695 |   0%
                                                                                                                                450_ 0_$a$y$z: |                         |      1/414695 |   0%
                                                                                                                                  450_ 0_$a$z: |                         |    595/414695 |   0%
                                                                                                                                450_ 0_$a$z$x: |                         |      4/414695 |   0%
                                                                                                                                450_ 0_$a$z$z: |                         |      1/414695 |   0%
                                                                                                                                  450_ 0_$w$a: |                         |   7356/414695 |   1%
                                                                                                                                450_ 0_$w$a$v: |                         |     74/414695 |   0%
                                                                                                                              450_ 0_$w$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                450_ 0_$w$a$x: |                         |   2196/414695 |   0%
                                                                                                                              450_ 0_$w$a$x$v: |                         |      2/414695 |   0%
                                                                                                                              450_ 0_$w$a$x$x: |                         |    109/414695 |   0%
                                                                                                                              450_ 0_$w$a$x$y: |                         |      6/414695 |   0%
                                                                                                                              450_ 0_$w$a$x$z: |                         |     13/414695 |   0%
                                                                                                                            450_ 0_$w$a$x$z$x: |                         |      1/414695 |   0%
                                                                                                                                450_ 0_$w$a$y: |                         |     83/414695 |   0%
                                                                                                                              450_ 0_$w$a$y$x: |                         |      8/414695 |   0%
                                                                                                                                450_ 0_$w$a$z: |                         |    262/414695 |   0%
                                                                                                                              450_ 0_$w$a$z$v: |                         |      1/414695 |   0%
                                                                                                                              450_ 0_$w$a$z$x: |                         |    206/414695 |   0%
                                                                                                                            450_ 0_$w$a$z$x$x: |                         |      1/414695 |   0%
                                                                                                                            450_ 0_$w$a$z$x$y: |                         |      2/414695 |   0%
                                                                                                                              450_ 0_$w$a$z$y: |                         |      2/414695 |   0%
                                                                                                                              450_ 0_$w$a$z$z: |                         |      6/414695 |   0%
                                                                                                                            450_ 0_$w$a$z$z$x: |                         |      1/414695 |   0%
                                                                                                                                    451_  _$a: |=                        |  24960/414695 |   6%
                                                                                                                                  451_  _$a$v: |                         |     75/414695 |   0%
                                                                                                                                  451_  _$a$x: |                         |    105/414695 |   0%
                                                                                                                                451_  _$a$x$v: |                         |      7/414695 |   0%
                                                                                                                                451_  _$a$x$x: |                         |     10/414695 |   0%
                                                                                                                              451_  _$a$x$x$v: |                         |      2/414695 |   0%
                                                                                                                                451_  _$a$x$y: |                         |    155/414695 |   0%
                                                                                                                              451_  _$a$x$y$v: |                         |      6/414695 |   0%
                                                                                                                              451_  _$a$x$y$x: |                         |      8/414695 |   0%
                                                                                                                                451_  _$a$x$z: |                         |      7/414695 |   0%
                                                                                                                                  451_  _$a$y: |                         |      1/414695 |   0%
                                                                                                                                  451_  _$a$z: |                         |      1/414695 |   0%
                                                                                                                                  451_  _$w$a: |                         |   2361/414695 |   0%
                                                                                                                                451_  _$w$a$v: |                         |     31/414695 |   0%
                                                                                                                              451_  _$w$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                451_  _$w$a$x: |                         |     76/414695 |   0%
                                                                                                                              451_  _$w$a$x$v: |                         |      9/414695 |   0%
                                                                                                                              451_  _$w$a$x$x: |                         |     13/414695 |   0%
                                                                                                                              451_  _$w$a$x$y: |                         |    221/414695 |   0%
                                                                                                                            451_  _$w$a$x$y$v: |                         |     22/414695 |   0%
                                                                                                                            451_  _$w$a$x$y$x: |                         |      3/414695 |   0%
                                                                                                                                451_  _$w$a$y: |                         |      4/414695 |   0%
                                                                                                                                    451_ 0_$a: |                         |     96/414695 |   0%
                                                                                                                                  451_ 0_$a$x: |                         |     83/414695 |   0%
                                                                                                                                451_ 0_$a$x$x: |                         |      7/414695 |   0%
                                                                                                                                451_ 0_$a$x$y: |                         |     89/414695 |   0%
                                                                                                                              451_ 0_$a$x$y$x: |                         |     16/414695 |   0%
                                                                                                                                451_ 0_$a$x$z: |                         |      1/414695 |   0%
                                                                                                                                  451_ 0_$w$a: |                         |     11/414695 |   0%
                                                                                                                                451_ 0_$w$a$v: |                         |      2/414695 |   0%
                                                                                                                                451_ 0_$w$a$x: |                         |    361/414695 |   0%
                                                                                                                              451_ 0_$w$a$x$v: |                         |      6/414695 |   0%
                                                                                                                              451_ 0_$w$a$x$x: |                         |     50/414695 |   0%
                                                                                                                              451_ 0_$w$a$x$y: |                         |    521/414695 |   0%
                                                                                                                            451_ 0_$w$a$x$y$v: |                         |     19/414695 |   0%
                                                                                                                            451_ 0_$w$a$x$y$x: |                         |     12/414695 |   0%
                                                                                                                          451_ 0_$w$a$x$y$x$x: |                         |      1/414695 |   0%
                                                                                                                              451_ 0_$w$a$x$z: |                         |      2/414695 |   0%
                                                                                                                                451_ 0_$w$a$y: |                         |      6/414695 |   0%
                                                                                                                                    455_  _$a: |                         |    506/414695 |   0%
                                                                                                                                  455_  _$w$a: |                         |     54/414695 |   0%
                                                                                                                                  480_  _$w$x: |                         |    151/414695 |   0%
                                                                                                                                480_  _$w$x$v: |                         |     11/414695 |   0%
                                                                                                                                480_  _$w$x$x: |                         |     35/414695 |   0%
                                                                                                                                480_  _$w$x$z: |                         |      2/414695 |   0%
                                                                                                                                    480_  _$x: |                         |    510/414695 |   0%
                                                                                                                                  480_  _$x$v: |                         |      8/414695 |   0%
                                                                                                                                  480_  _$x$x: |                         |     33/414695 |   0%
                                                                                                                                  482_  _$w$y: |                         |      5/414695 |   0%
                                                                                                                                    485_  _$v: |                         |    122/414695 |   0%
                                                                                                                                  485_  _$v$v: |                         |      7/414695 |   0%
                                                                                                                                  485_  _$w$v: |                         |     61/414695 |   0%
                                                                                                                                485_  _$w$v$v: |                         |      4/414695 |   0%
                                                                                                                                485_  _$w$v$x: |                         |      3/414695 |   0%
                                                                                                                                500_0 _$a$c$x: |                         |      2/414695 |   0%
                                                                                                                                500_0 _$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                  500_0 _$a$x: |                         |     14/414695 |   0%
                                                                                                                                500_0 _$a$x$v: |                         |      1/414695 |   0%
                                                                                                                                500_0 _$a$x$x: |                         |      1/414695 |   0%
                                                                                                                            500_0 _$w$a$c$d$v: |                         |      2/414695 |   0%
                                                                                                                            500_0 _$w$a$c$d$x: |                         |      3/414695 |   0%
                                                                                                                              500_0 _$w$a$c$v: |                         |      3/414695 |   0%
                                                                                                                              500_0 _$w$a$c$x: |                         |     21/414695 |   0%
                                                                                                                            500_0 _$w$a$c$x$z: |                         |     68/414695 |   0%
                                                                                                                              500_0 _$w$a$d$x: |                         |      2/414695 |   0%
                                                                                                                                500_0 _$w$a$v: |                         |      4/414695 |   0%
                                                                                                                              500_0 _$w$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                500_0 _$w$a$x: |                         |     31/414695 |   0%
                                                                                                                              500_0 _$w$a$x$v: |                         |      1/414695 |   0%
                                                                                                                              500_0 _$w$a$x$x: |                         |      2/414695 |   0%
                                                                                                                              500_0 _$w$a$x$y: |                         |      3/414695 |   0%
                                                                                                                              500_0 _$w$a$x$z: |                         |      2/414695 |   0%
                                                                                                                                500_00_$a$c$x: |                         |      4/414695 |   0%
                                                                                                                                  500_00_$a$x: |                         |     19/414695 |   0%
                                                                                                                              500_00_$w$a$c$v: |                         |      4/414695 |   0%
                                                                                                                              500_00_$w$a$c$x: |                         |     13/414695 |   0%
                                                                                                                              500_00_$w$a$d$t: |                         |      1/414695 |   0%
                                                                                                                                500_00_$w$a$x: |                         |     82/414695 |   0%
                                                                                                                              500_00_$w$a$x$x: |                         |      2/414695 |   0%
                                                                                                                              500_00_$w$a$x$y: |                         |      1/414695 |   0%
                                                                                                                              500_00_$w$a$x$z: |                         |      5/414695 |   0%
                                                                                                                                    500_1 _$a: |                         |      1/414695 |   0%
                                                                                                                                500_1 _$a$d$v: |                         |      1/414695 |   0%
                                                                                                                                500_1 _$a$d$x: |                         |      8/414695 |   0%
                                                                                                                            500_1 _$w$a$c$x$z: |                         |      1/414695 |   0%
                                                                                                                              500_1 _$w$a$d$v: |                         |      3/414695 |   0%
                                                                                                                              500_1 _$w$a$d$x: |                         |     14/414695 |   0%
                                                                                                                            500_1 _$w$a$d$x$v: |                         |      1/414695 |   0%
                                                                                                                            500_1 _$w$a$d$x$x: |                         |      1/414695 |   0%
                                                                                                                            500_1 _$w$a$d$x$z: |                         |      2/414695 |   0%
                                                                                                                                500_1 _$w$a$v: |                         |      1/414695 |   0%
                                                                                                                              500_10_$w$a$d$x: |                         |      4/414695 |   0%
                                                                                                                                    500_3 _$a: |                         |   1357/414695 |   0%
                                                                                                                                  500_3 _$w$a: |                         |      4/414695 |   0%
                                                                                                                                    500_30_$a: |                         |   2333/414695 |   0%
                                                                                                                                  500_30_$w$a: |                         |     20/414695 |   0%
                                                                                                                                510_1 _$a$b$x: |                         |      2/414695 |   0%
                                                                                                                              510_1 _$w$a$b$x: |                         |     25/414695 |   0%
                                                                                                                                510_10_$a$b$x: |                         |      1/414695 |   0%
                                                                                                                            510_10_$w$a$b$b$x: |                         |      1/414695 |   0%
                                                                                                                              510_10_$w$a$b$x: |                         |    101/414695 |   0%
                                                                                                                            510_10_$w$a$b$x$x: |                         |      4/414695 |   0%
                                                                                                                          510_10_$w$a$b$x$x$x: |                         |      1/414695 |   0%
                                                                                                                                    510_2 _$a: |                         |     18/414695 |   0%
                                                                                                                                  510_2 _$a$v: |                         |      1/414695 |   0%
                                                                                                                                  510_2 _$a$x: |                         |      4/414695 |   0%
                                                                                                                                  510_2 _$w$a: |                         |    130/414695 |   0%
                                                                                                                                510_2 _$w$a$v: |                         |      5/414695 |   0%
                                                                                                                                510_2 _$w$a$x: |                         |     30/414695 |   0%
                                                                                                                              510_2 _$w$a$x$v: |                         |      1/414695 |   0%
                                                                                                                              510_2 _$w$a$x$x: |                         |      4/414695 |   0%
                                                                                                                                510_2 _$w$a$z: |                         |      3/414695 |   0%
                                                                                                                                    510_20_$a: |                         |      3/414695 |   0%
                                                                                                                                  510_20_$a$x: |                         |      3/414695 |   0%
                                                                                                                                510_20_$a$x$x: |                         |      2/414695 |   0%
                                                                                                                                  510_20_$w$a: |                         |     55/414695 |   0%
                                                                                                                                510_20_$w$a$x: |                         |     91/414695 |   0%
                                                                                                                              510_20_$w$a$x$x: |                         |      5/414695 |   0%
                                                                                                                              510_20_$w$a$x$y: |                         |      2/414695 |   0%
                                                                                                                              510_20_$w$a$x$z: |                         |      1/414695 |   0%
                                                                                                                                510_20_$w$a$z: |                         |     12/414695 |   0%
                                                                                                                              510_20_$w$a$z$x: |                         |      2/414695 |   0%
                                                                                                                            510_20_$w$a$z$x$y: |                         |      1/414695 |   0%
                                                                                                                                    530_ 0_$a: |                         |      3/414695 |   0%
                                                                                                                                530_ 0_$a$p$v: |                         |      1/414695 |   0%
                                                                                                                                530_ 0_$a$p$x: |                         |      4/414695 |   0%
                                                                                                                                  530_ 0_$a$v: |                         |      1/414695 |   0%
                                                                                                                                  530_ 0_$a$x: |                         |     18/414695 |   0%
                                                                                                                                  530_ 0_$w$a: |                         |     15/414695 |   0%
                                                                                                                            530_ 0_$w$a$l$x$x: |                         |      1/414695 |   0%
                                                                                                                                530_ 0_$w$a$p: |                         |      2/414695 |   0%
                                                                                                                            530_ 0_$w$a$p$l$v: |                         |      1/414695 |   0%
                                                                                                                              530_ 0_$w$a$p$v: |                         |      1/414695 |   0%
                                                                                                                              530_ 0_$w$a$p$x: |                         |     21/414695 |   0%
                                                                                                                                530_ 0_$w$a$v: |                         |      9/414695 |   0%
                                                                                                                                530_ 0_$w$a$x: |                         |     41/414695 |   0%
                                                                                                                              530_ 0_$w$a$x$x: |                         |      3/414695 |   0%
                                                                                                                                    550_  _$a: |                         |   5485/414695 |   1%
                                                                                                                                  550_  _$a$v: |                         |    133/414695 |   0%
                                                                                                                                  550_  _$a$x: |                         |    529/414695 |   0%
                                                                                                                                550_  _$a$x$v: |                         |      3/414695 |   0%
                                                                                                                                550_  _$a$x$x: |                         |     16/414695 |   0%
                                                                                                                              550_  _$a$x$x$x: |                         |      1/414695 |   0%
                                                                                                                                550_  _$a$x$z: |                         |      6/414695 |   0%
                                                                                                                                  550_  _$a$y: |                         |    102/414695 |   0%
                                                                                                                                  550_  _$a$z: |                         |     52/414695 |   0%
                                                                                                                                550_  _$a$z$y: |                         |      7/414695 |   0%
                                                                                                                                  550_  _$w$a: |===                      |  56578/414695 |  13%
                                                                                                                                550_  _$w$a$v: |                         |    314/414695 |   0%
                                                                                                                              550_  _$w$a$v$v: |                         |      1/414695 |   0%
                                                                                                                                550_  _$w$a$x: |                         |   4901/414695 |   1%
                                                                                                                              550_  _$w$a$x$v: |                         |      3/414695 |   0%
                                                                                                                              550_  _$w$a$x$x: |                         |    235/414695 |   0%
                                                                                                                            550_  _$w$a$x$x$x: |                         |      6/414695 |   0%
                                                                                                                          550_  _$w$a$x$x$x$z: |                         |      1/414695 |   0%
                                                                                                                            550_  _$w$a$x$x$z: |                         |      7/414695 |   0%
                                                                                                                              550_  _$w$a$x$y: |                         |     74/414695 |   0%
                                                                                                                              550_  _$w$a$x$z: |                         |   1415/414695 |   0%
                                                                                                                            550_  _$w$a$x$z$x: |                         |      5/414695 |   0%
                                                                                                                            550_  _$w$a$x$z$z: |                         |      9/414695 |   0%
                                                                                                                                550_  _$w$a$y: |                         |    651/414695 |   0%
                                                                                                                              550_  _$w$a$y$x: |                         |     11/414695 |   0%
                                                                                                                                550_  _$w$a$z: |===                      |  51153/414695 |  12%
                                                                                                                              550_  _$w$a$z$v: |                         |      9/414695 |   0%
                                                                                                                              550_  _$w$a$z$x: |                         |     86/414695 |   0%
                                                                                                                            550_  _$w$a$z$x$y: |                         |     10/414695 |   0%
                                                                                                                            550_  _$w$a$z$y$x: |                         |      1/414695 |   0%
                                                                                                                              550_  _$w$a$z$z: |                         |      4/414695 |   0%
                                                                                                                            550_  _$w$a$z$z$x: |                         |      1/414695 |   0%
                                                                                                                                    550_ 0_$a: |                         |   5962/414695 |   1%
                                                                                                                                  550_ 0_$a$v: |                         |      3/414695 |   0%
                                                                                                                                  550_ 0_$a$x: |                         |    583/414695 |   0%
                                                                                                                                550_ 0_$a$x$x: |                         |     18/414695 |   0%
                                                                                                                              550_ 0_$a$x$x$x: |                         |      1/414695 |   0%
                                                                                                                                550_ 0_$a$x$y: |                         |      1/414695 |   0%
                                                                                                                                  550_ 0_$a$y: |                         |      7/414695 |   0%
                                                                                                                                550_ 0_$a$y$x: |                         |      1/414695 |   0%
                                                                                                                                  550_ 0_$a$z: |                         |     17/414695 |   0%
                                                                                                                                550_ 0_$a$z$x: |                         |      1/414695 |   0%
                                                                                                                                  550_ 0_$w$a: |====                     |  72219/414695 |  17%
                                                                                                                                550_ 0_$w$a$v: |                         |     20/414695 |   0%
                                                                                                                                550_ 0_$w$a$x: |                         |   7217/414695 |   1%
                                                                                                                              550_ 0_$w$a$x$x: |                         |    388/414695 |   0%
                                                                                                                            550_ 0_$w$a$x$x$x: |                         |      6/414695 |   0%
                                                                                                                            550_ 0_$w$a$x$x$y: |                         |      1/414695 |   0%
                                                                                                                            550_ 0_$w$a$x$x$z: |                         |     10/414695 |   0%
                                                                                                                              550_ 0_$w$a$x$y: |                         |     95/414695 |   0%
                                                                                                                            550_ 0_$w$a$x$y$x: |                         |      1/414695 |   0%
                                                                                                                              550_ 0_$w$a$x$z: |                         |    510/414695 |   0%
                                                                                                                            550_ 0_$w$a$x$z$x: |                         |      1/414695 |   0%
                                                                                                                            550_ 0_$w$a$x$z$z: |                         |      2/414695 |   0%
                                                                                                                                550_ 0_$w$a$y: |                         |    314/414695 |   0%
                                                                                                                              550_ 0_$w$a$y$x: |                         |     25/414695 |   0%
                                                                                                                                550_ 0_$w$a$z: |                         |   7300/414695 |   1%
                                                                                                                              550_ 0_$w$a$z$x: |                         |    108/414695 |   0%
                                                                                                                            550_ 0_$w$a$z$x$y: |                         |     13/414695 |   0%
                                                                                                                              550_ 0_$w$a$z$y: |                         |      1/414695 |   0%
                                                                                                                              550_ 0_$w$a$z$z: |                         |      2/414695 |   0%
                                                                                                                                    551_  _$a: |                         |    213/414695 |   0%
                                                                                                                                  551_  _$a$x: |                         |     21/414695 |   0%
                                                                                                                                551_  _$a$x$x: |                         |      2/414695 |   0%
                                                                                                                                551_  _$a$x$y: |                         |     14/414695 |   0%
                                                                                                                                551_  _$a$z$y: |                         |      1/414695 |   0%
                                                                                                                                  551_  _$w$a: |                         |   2035/414695 |   0%
                                                                                                                                551_  _$w$a$v: |                         |    451/414695 |   0%
                                                                                                                                551_  _$w$a$x: |                         |   9305/414695 |   2%
                                                                                                                              551_  _$w$a$x$v: |                         |      2/414695 |   0%
                                                                                                                              551_  _$w$a$x$x: |                         |     36/414695 |   0%
                                                                                                                          551_  _$w$a$x$x$x$y: |                         |      1/414695 |   0%
                                                                                                                            551_  _$w$a$x$x$z: |                         |      1/414695 |   0%
                                                                                                                              551_  _$w$a$x$y: |                         |    982/414695 |   0%
                                                                                                                            551_  _$w$a$x$y$x: |                         |    600/414695 |   0%
                                                                                                                              551_  _$w$a$x$z: |                         |     12/414695 |   0%
                                                                                                                                551_  _$w$a$z: |                         |     18/414695 |   0%
                                                                                                                            551_  _$w$a$z$z$z: |                         |      1/414695 |   0%
                                                                                                                                    551_ 0_$a: |                         |     19/414695 |   0%
                                                                                                                                  551_ 0_$a$x: |                         |      7/414695 |   0%
                                                                                                                                551_ 0_$a$x$y: |                         |     29/414695 |   0%
                                                                                                                              551_ 0_$a$x$y$x: |                         |      8/414695 |   0%
                                                                                                                                551_ 0_$a$x$z: |                         |      1/414695 |   0%
                                                                                                                                  551_ 0_$w$a: |                         |     40/414695 |   0%
                                                                                                                                551_ 0_$w$a$x: |                         |   2525/414695 |   0%
                                                                                                                              551_ 0_$w$a$x$x: |                         |     29/414695 |   0%
                                                                                                                              551_ 0_$w$a$x$y: |                         |    483/414695 |   0%
                                                                                                                            551_ 0_$w$a$x$y$x: |                         |     86/414695 |   0%
                                                                                                                              551_ 0_$w$a$x$z: |                         |     12/414695 |   0%
                                                                                                                                    555_  _$a: |                         |    118/414695 |   0%
                                                                                                                                  555_  _$w$a: |                         |    660/414695 |   0%
                                                                                                                                  580_  _$w$x: |                         |    880/414695 |   0%
                                                                                                                                580_  _$w$x$x: |                         |      4/414695 |   0%
                                                                                                                                  581_  _$w$z: |                         |      1/414695 |   0%
                                                                                                                                  585_  _$w$v: |                         |    169/414695 |   0%
                                                                                                                                585_  _$w$v$v: |                         |      6/414695 |   0%
                                                                                                                                    667_  _$a: |=====                    |  84396/414695 |  20%
                                                                                                                                    670_  _$a: |===========              | 186217/414695 |  44%
                                                                                                                                  670_  _$a$b: |======                   | 114421/414695 |  27%
                                                                                                                                670_  _$a$b$b: |                         |      1/414695 |   0%
                                                                                                                                670_  _$a$b$u: |                         |     38/414695 |   0%
                                                                                                                                  670_  _$a$u: |                         |      1/414695 |   0%
                                                                                                                                    675_  _$a: |                         |  15114/414695 |   3%
                                                                                                                                  675_  _$a$a: |=                        |  16676/414695 |   4%
                                                                                                                                675_  _$a$a$a: |                         |   8082/414695 |   1%
                                                                                                                              675_  _$a$a$a$a: |                         |   3351/414695 |   0%
                                                                                                                            675_  _$a$a$a$a$a: |                         |   1468/414695 |   0%
                                                                                                                          675_  _$a$a$a$a$a$a: |                         |    671/414695 |   0%
                                                                                                                        675_  _$a$a$a$a$a$a$a: |                         |    306/414695 |   0%
                                                                                                                      675_  _$a$a$a$a$a$a$a$a: |                         |    149/414695 |   0%
                                                                                                                    675_  _$a$a$a$a$a$a$a$a$a: |                         |     62/414695 |   0%
                                                                                                                  675_  _$a$a$a$a$a$a$a$a$a$a: |                         |     20/414695 |   0%
                                                                                                                675_  _$a$a$a$a$a$a$a$a$a$a$a: |                         |     19/414695 |   0%
                                                                                                              675_  _$a$a$a$a$a$a$a$a$a$a$a$a: |                         |      5/414695 |   0%
                                                                                                            675_  _$a$a$a$a$a$a$a$a$a$a$a$a$a: |                         |      6/414695 |   0%
                                                                                                          675_  _$a$a$a$a$a$a$a$a$a$a$a$a$a$a: |                         |      2/414695 |   0%
                                                                                                        675_  _$a$a$a$a$a$a$a$a$a$a$a$a$a$a$a: |                         |      2/414695 |   0%
                                                                                                  675_  _$a$a$a$a$a$a$a$a$a$a$a$a$a$a$a$a$a$a: |                         |      1/414695 |   0%
                                                                                                                                    678_  _$a: |                         |      1/414695 |   0%
                                                                                                                                    680_  _$a: |                         |      1/414695 |   0%
                                                                                                                                  680_  _$a$a: |                         |      1/414695 |   0%
                                                                                                                              680_  _$a$a$i$a: |                         |      1/414695 |   0%
                                                                                                                                    680_  _$i: |                         |   8305/414695 |   2%
                                                                                                                                  680_  _$i$a: |                         |   2575/414695 |   0%
                                                                                                                                680_  _$i$a$a: |                         |      5/414695 |   0%
                                                                                                                              680_  _$i$a$a$a: |                         |      1/414695 |   0%
                                                                                                                            680_  _$i$a$a$i$a: |                         |      2/414695 |   0%
                                                                                                                                680_  _$i$a$i: |                         |    207/414695 |   0%
                                                                                                                              680_  _$i$a$i$a: |                         |    576/414695 |   0%
                                                                                                                            680_  _$i$a$i$a$a: |                         |      1/414695 |   0%
                                                                                                                            680_  _$i$a$i$a$i: |                         |     37/414695 |   0%
                                                                                                                          680_  _$i$a$i$a$i$a: |                         |    127/414695 |   0%
                                                                                                                        680_  _$i$a$i$a$i$a$i: |                         |      9/414695 |   0%
                                                                                                                      680_  _$i$a$i$a$i$a$i$a: |                         |     43/414695 |   0%
                                                                                                                    680_  _$i$a$i$a$i$a$i$a$i: |                         |      5/414695 |   0%
                                                                                                                  680_  _$i$a$i$a$i$a$i$a$i$a: |                         |     22/414695 |   0%
                                                                                                                680_  _$i$a$i$a$i$a$i$a$i$a$i: |                         |      3/414695 |   0%
                                                                                                              680_  _$i$a$i$a$i$a$i$a$i$a$i$a: |                         |      6/414695 |   0%
                                                                                                                              680_  _$i$a$i$i: |                         |      2/414695 |   0%
                                                                                                                        680_  _$i$i$a$i$a$i$a: |                         |      2/414695 |   0%
                                                                                                                                  681_  _$i$a: |                         |   9633/414695 |   2%
                                                                                                                                681_  _$i$a$a: |                         |      3/414695 |   0%
                                                                                                                              681_  _$i$a$i$a: |                         |    248/414695 |   0%
                                                                                                                          681_  _$i$a$i$a$i$a: |                         |      2/414695 |   0%
                                                                                                                                  781_ 0_$w$z: |                         |     10/414695 |   0%
                                                                                                                                781_ 0_$w$z$z: |                         |     16/414695 |   0%
                                                                                                                                    781_ 0_$z: |                         |   3725/414695 |   0%
                                                                                                                                  781_ 0_$z$z: |==                       |  36872/414695 |   8%

```
