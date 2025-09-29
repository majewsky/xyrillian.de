- Motivation: reguläre Ausdrücke verstehen
    - etwas Theorie ist notwendig, um die Eigenschaften und Grenzen von regulären Ausdrücken zu verstehen
    - quasi eine Erstsemestervorlesung -> Anekdote Xyrill: Büchlein "Einblick ins Studium"

[Siehe Wikipedia](https://de.wikipedia.org/wiki/Grammatik ):

> Grammatik bezeichnet in der Linguistik jede Form einer systematischen Sprachbeschreibung. Dabei steht der Begriff der Grammatik einmal für das Regelwerk selbst, auf der anderen Seite aber auch für die Theorie über eine bestimmte Sprache oder Sprachfamilie.

- im Kontext der Informatik: **formale Grammatiken** für **formale Sprachen**
    - formale Sprache: Menge von Symbolketten ("Wörter"), die aus Symbolen ("Alphabet") zusammengesetzt werden
    - formale Grammatik: endlich große Beschreibung einer (im Allgemeinen unendlich großen) formalen Sprache
    - Beispiele: Programmiersprachen, Tastenkombos in Videospielen, Münzeinwurf in Verkaufsautomat
    - Abwägung: komplexere Grammatiken sind flexibler, aber auch aufwändiger (z.B. Tastatureingaben vs. Sprachsteuerung)

- [Chomsky-Hierarchie](https://de.wikipedia.org/wiki/Chomsky-Hierarchie#%C3%9Cbersicht ): Klassifikation von formalen Grammatiken in aufsteigender Striktheit und absteigender Flexibilität
    - Beschreibung der entsprechenden formalen Grammatiken: hier nicht, zu technisch
    - zu jeder Ebene der Chomsky-Hierarchie gehört ein Maschinenmodell, dass entsprechende Grammatiken akzeptieren kann
    - Typ 3: reguläre Grammatiken mittels endlichem Zustandsautomat (Beispiel: Münzeinwurf in Verkaufsautomat)
    - Typ 2: kontextfreie Grammatiken mittels Kellerautomat (Beispiel: balancierte Klammerpaare)
    - Typ 1 und 0 mittels Turingmaschinen (erst mit beschränktem, dann mit unbeschränktem Speicher)

- Beispiel für grafische Repräsentation von Grammatiken: [siehe SQLite-Dokumentation](https://sqlite.org/lang_select.html )

- reguläre Grammatiken im Alltag: [reguläre Ausdrücke](https://de.wikipedia.org/wiki/Regul%C3%A4rer_Ausdruck ); formale Definition mit vier Grundbausteinen

| Grundbaustein  | Beispielausdruck | Findet z.B. |
| -------------- | ---------------- | ----------- |
| Zeichen(folge) | `Ausdruck` | `Ausdruck` |
| Alternative | `Ausdruck\|Ausdrücke` | `Ausdruck` und `Ausdrücke` |
| Gruppierung | `Ausdr(uck\|ücke)` | `Ausdruck` und `Ausdrücke` |
| Sternquantor | `to*ll` | `tll`, `toll`, `tooll`, `toooll`, etc. |

- in der Praxis weitere Bausteine, um häufige Anwendungen zu vereinfachen

| Zusatzbaustein | Definition | Beispielausdruck | Findet z.B. |
| --- | --- | --- | --- |
| Zeichenklasse | `[ABC...] = A\|B\|C\|...` | positive Ganzzahl: `0\|[1-9][0-9]*` | `0`, `42`, `900`, aber nicht `0815` |
| feste Zeichenklasse | `\s` für Whitespace, `\d = [0-9]` für Ziffern, etc. | |
| beliebiges Zeichen | `.` | Klammerausdruck: `(.*)` | `(2 + 4)`, `(siehe oben)`, aber auch `(siehe oben (oder nicht)` |
| Fragezeichenquantor | `A? = (\|A)` | Ganzzahl mit Vorzeichen: `[+-]?(0\|[1-9][0-9]*)` | `+0`, `-42`, `23` |
| Fixquantor | z.B. `A{3} = AAA`, `A{2,4} = AAA?A?` | Postleitzahl: `[0-9]{5}` | `01127`, `50616` |
| Plusquantor | `A+ = AA*` | Domainname: `([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+` | `www.example.com` |
| Anker | `^` für Zeilenanfang, `$` für Zeilenende, `\b` für Wortgrenze | `Schulz\b` | `Schulz`, aber nicht `Schulze` |

- Problem in der Praxis: **Escaping**
    - z.B. Klammern sind Teil der Syntax; um eine Klammer im Eingabetext darzustellen, schreibt man "\\("

- Regexen lernen
    - Online
        - [LearnRegex](https://regexlearn.com/learn/regex101 ) ist ein interaktives Tutorial zum Erlernen von regulären Ausdrücken, in englischer Sprache.
        - Mit [Regex Crossword](https://regexcrossword.com ) kann man spielerisch besonders das Lesen regulärer Ausdrücke üben.
    - Apps
        - [Regex-Kreuzworträtsel auf dem Mobiltelefon](https://f-droid.org/en/packages/de.chagemann.regexcrossword/ )
        - [Mehrere Ausdrücke mit regulären Ausdrücken erfassen bzw. ausschließen](https://f-droid.org/en/packages/com.phikal.regex/ )

