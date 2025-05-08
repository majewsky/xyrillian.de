- Rückbezug STP071: anhand von JSON haben wir die sichtbaren Datentypen besprochen
    - heute tauchen wir unter die Wasseroberfläche ab und schauen auf Typen, die man eher im Speicher als in einer Datei sieht

- Einwurf Kombinatorik: Wieviele mögliche Werte haben die vorgenannten Typen (sofern nicht unendlich)? (wir schreiben `K(T)` für die Kardinalität des Typs `T`)
    - Boolean: `K(bool) = 2` (`true` oder `false`)
    - Zahlen: `K(T_N) = 2^N` Werte für Zahltypen mit `N` Bits (z.B. 256 für `i8` oder 65536 für `u16`), weil jedes Bit zwei Werte hat, die unabhängig voneinander gewählt werden können
    - Listen fester Länge: `K([T; N]) = K(T)^N` für `K(T)` mögliche Werte des Elementtyps `T` und Länge `N`
    - Listen variabler Länge: `K(Vec<T>) = ∞` bis auf praktische Platzbeschränkungen

- Listen wirken sich auf die Kardinalität also wie eine Potenzierung aus
    - Gibt es auch andere Rechenoperationen? -> [algebraische Datentypen](https://de.wikipedia.org/w/index.php?title=Algebraischer_Datentyp&oldid=248147888)

- Produkttypen: Gesamttyp `P` enthält je einen Wert von Typ `T_1, T_2, ..., T_N`, somit `K(P) = K(T_1) * K(T_2) * ... * K(T_N)`
    - in Rust: entweder als Tupel (geordnete Folge von Werten, die mit Zahlenindizes bezeichnet sind), z.B. `type Vector3D = (f64, f64, f64)`, oder als Strukturtypen (ungeordnete Menge von Werten, die mit Namen bezeichnet sind, z.B. `struct Person { first_name: String, last_name: String }`)
    - in JSON: Tupel als Listen, Strukturtypen als Objekte (beim Einlesen muss die Elementzahl bzw. die Namen der Elemente geprüft werden)

- Summentypen: Gesamttyp `S` enthält entweder einen Wert vom Typ `T_1` oder `T_2` oder ... oder `T_N`, somit `K(P) = K(T_1) + K(T_2) + ... + K(T_N)`
    - in Rust: Enumerationstyp, z.B. `enum ParseResult { Ok(DataStructure), Err(ParseError) }`
    - in JSON: keine klare Konvention, Serialisierung in Rust erzeugt ein Objekt mit einem Eintrag, die nach der jeweiligen Variante benannt ist (z.B. `{"Ok": ...}`)

- sehr kleine Spezialfälle: Gibt es Datentypen, die kleiner als `bool` sind?
    - Produkttyp mit null Feldern (in Rust `struct Empty {}`): nur ein möglicher Wert (1 ist das neutrale Element der Multiplikation)
        - in JSON kann man das als `null` serialisieren
        - in Rust kommt dies im Summentyp `enum Option<T> { Some(T), None }` vor; `None` ist hier im Prinzip eine Kurzschreibweise für `None(Empty)`
    - Summentyp mit null Varianten (in Rust `enum Impossible {}`): keine möglichen Werte (0 ist das neutrale Element der Addition)
        - in JSON nicht darstellbar (keine Werte haben keine Serialisierung)
        - in Rust kommt so ein Typ in Funktionen vor, die kein Ergebnis liefern (z.B. `quit()`)
        - nützlicher, als es zunächst erscheint, sobald [Monomorphisierung](https://en.wikipedia.org/w/index.php?title=Monomorphization&oldid=1246708145) ins Spiel kommt

- Abendgedanken: Warum das ganze Gerede über Datentypen?
    - angewandte Programmierung besteht aus zwei (verwandten) Teildisziplinen: Algorithmen (das Wie, der Rechenteil) und Datenstrukturen (das Was, der Speicherteil)
    - Xyrill hält Datenstrukturen für das Fundament; die Wahl der Algorithmen muss an zweiter Stelle kommen
    - vergleichbarer Ansatz in klassischen Software-Design-Lehrbüchern wie [SICP](https://en.wikipedia.org/w/index.php?title=Structure_and_Interpretation_of_Computer_Programs&oldid=1244483390) und [HtDP](https://en.wikipedia.org/w/index.php?title=How_to_Design_Programs&oldid=1251901750)
