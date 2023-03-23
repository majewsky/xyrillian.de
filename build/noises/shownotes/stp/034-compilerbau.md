- [Compiler](https://de.wikipedia.org/w/index.php?title=Compiler&oldid=225542576 ): wandelt eine Programmiersprache (siehe STP033) in eine andere Programmiersprache oder Maschinensprache (siehe STP011)
    - insb. Erkennung und Abweisung fehlerhafter Konstrukte (z.B. `42 + 5` ist okay, aber `42 + "Test"` nicht)

- Unterbegriffe zu "Compiler"
    - **Interpreter**: anstatt Erzeugung eines Zielcodes direkte Ausführung
    - **REPL**: Interpreter für die interaktive Verwendung ("Read Evaluate Print Loop")
    - **JIT-Compiler**: direkte Ausführung wie ein Interpreter, aber es wird echter Maschinencode erzeugt ("Just In Time")
    - z.B. portable Programme: Compiler erzeugt Maschinencode für ein virtuelles Maschinenmodell, Laufzeitumgebung ist ein Interpreter oder JIT-Compiler für diesen Maschinencode -> analog dazu Frontend/Backend-Aufteilung innerhalb eines Compilers

- Frontend: Umformen der Texteingabe in eine Datenstruktur, Prüfung dieser Struktur auf inhaltliche Korrektheit (maschinenunabhängig, aber sprachabhängig)
    - lexikalische Analyse: Zerteilen des Eingabetextes in eine Folge von Tokens mittels regulärer Ausdrücke (siehe STP021; diese Vorstufe ist eine Optimierung, weil reguläre Ausdrücke schneller ausgewertet werden können)
    - syntaktische Analyse: Umformen der Tokenliste in einen **Syntaxbaum** mittels eines kontextfreien Parsers (siehe [Beispiel bei Wikipedia](https://de.wikipedia.org/w/index.php?title=Parser&oldid=225991691#Beispiel ))
    - semantische Analyse: Prüfen des Syntaxbaums auf inhaltliche Korrektheit, meist mittels Markierungen an jedem Knoten des Baums (z.B. Welchen Typ hat der hier berechnete Wert? Passen die Werte des Elternknotens bzw. des Kindknotens dazu? etc.)
    - IR-Erzeugung: Umformung des Syntaxbaums in eine **Intermediate Representation** (IR), die sich für maschinenunabhängige Optimierung eignet

- Fortsetzung des Beispiels aus Wikipedia (aber mit einer Eingabevariable `y` und Ausgabevariable `x`): mögliche IR für den Ausdruck `x = y + (2 + 2) - sin(pi)`

```
$1 = 2
$2 = $1 + $1
$3 = load %y
$4 = $3 + $2
$5 = pi
$6 = call sin($5)
$7 = $4 - $6
store %x, $7
```

- Middleend: Optimierungsläufe auf IR (Auswahl)
    - Statische Formelauswertung zur Programmzeit: z.B. oben statt `$2 = $1 + $1` direkt `$2 = 4`
    - Löschung von unerreichbarem Code: sehr einfach bei IR mit **Static Single Assignment** (SSA) wie im Beispiel
    - Invariantenextraktion aus Schleifen: wenn z.B. für verschiedene Werte von `x` der Ausdruck `sin(x) + cos(y)` berechnet wird, so kann man `cos(y)` vor der Schleife ein einziges Mal berechnen und das Zwischenergebnis wiederverwenden
    - Inlining: ist eine Funktion sehr kurz, kann es effizienter sein, einen Funktionsaufruf durch eine Kopie der Funktion zu ersetzen (insb. wenn dann weitere Optimierungen möglich werden)

- Beispiel: optimierte Variante der IR von oben

```
$3 = load %y
$7 = $3 + 4
store %x, $7
```

- Backend: Umwandlung der maschinenunabhängigen IR in tatsächlichen Maschinencode
    - Codeauswahl: für die einzelnen Operationen Auswahl der besten Maschinenbefehle
        - je nach Optimierungsziel: Geschwindigkeit, Speicherverbrauch, Energieverbrauch, etc. (z.B. `x * 3` kann je nach Maschinenarchitektur gewinnbringend durch `x + x + x` ersetzt werden)
        - schwieriger als gedacht wegen komplexen Befehlen (z.B. SIMD: Single Instruction Multiple Data)
    - Registerzuweisung: nummerierte Register aus der IR ersetzen durch die tatsächlich in der Maschine verfügbaren Register
        - reduzierbar auf ein [Graphenfärbungsproblem](https://de.wikipedia.org/w/index.php?title=F%C3%A4rbung_(Graphentheorie)&oldid=228052439 ): Knoten = IR-Register, Kanten = gemeinsame Verwendung in Berechnungen, Farben = echte Register
    - Scheduling: optimale Sortierung der einzelnen Maschinenbefehle (soweit umordenbar)
        - z.B. parallele Ausnutzung verschiedener Recheneinheiten in der CPU (für Ganzzahl- und Fließkommazahl-Operationen)
    - maschinenabhängige Optimierungen (aus Zeitgründen hier nicht)
