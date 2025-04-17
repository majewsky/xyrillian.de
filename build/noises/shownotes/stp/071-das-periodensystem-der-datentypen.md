- [JSON](https://de.wikipedia.org/w/index.php?title=JavaScript_Object_Notation&oldid=249644714): ein populäres Format für die Serialisierung von Daten
    - [Serialisierung](https://de.wikipedia.org/w/index.php?title=Serialisierung&oldid=241468956): "Abbildung von strukturierten Daten auf eine sequenzielle Darstellungsform", insb. auf Bytefolgen wie in einer Datei oder einer Datenübertragung im Netzwerk
    - JSON ist bekannt dafür, nicht erfunden, sondern "entdeckt" worden zu sein
        - https://www.corecursive.com/json-vs-xml-douglas-crockford/
    - außerdem relativ minimal: nur sechs verschiedene Grundstrukturen
    - aber trotzdem relativ universell: damit unser Ausgangspunkt für die Kartografierung der Welt der Datentypen
    - außerdem Gegenüberstellung mit dem Typvorrat von Rust (einer zeitgenössischen Programmiersprache mit einem detaillierten Typsystem)

- atomare Datentypen: einzelne Werte (C++ nennt sie "Plain Old Data" (POD), Perl nennt sie "Skalare")
    - Boolean: ein Bit (wahr oder falsch)
        - Typvorrat in Rust: `bool`
        - in JSON: `true` oder `false`
    - Zahlen: in verschiedenen Bitgrößen, mit oder ohne Vorzeichen, Ganzzahl oder Fließkommazahl
        - Typvorrat in Rust: `u8/i8, u16/i16, u32/i32, u64/i64, u128/i128, usize/isize, f32, f64`
        - in JSON: wie gewohnt ausgeschrieben (z.B. `42`, `-23.5`) oder in wissenschaftlicher Notation (z.B. `7.297e-3` oder `-6.022E23`)
    - Zeichenfolgen: nicht wirklich unteilbar, aber wird für gewöhnlich als einzelner Wert verstanden
        - Typvorrat in Rust: `String` (UTF-8), `OsString` (Kodierung gemäß Zeichentabelle des Systems) und andere
        - in JSON: als Zeichenkette mit Anführungszeichen und mit Escape-Sequenzen für nicht druckbare Zeichen (z.B. `"Mit freundlichen Grüßen\nErika Mustermann"`)
    - fehlender Wert: in JSON `null` (heißt woanders vielleicht auch `nil`)
        - Typvorrat in Rust: siehe später

- Listen: ein abgeleiteter Datentyp
    - deswegen vorhin "atomar": der Typ einer Liste bestimmt sich aus dem Typ der Elemente (in statisch typisierten Programmiersprachen ist z.B. "Liste von 8-Bit-Zahlen" etwas anderes als "Liste von Strings" oder "Liste von Listen von Strings")
    - Typvorrat in Rust:
        - mit variabler Länge: `Vec<X>`, wobei `X` für den Typ der Elemente steht (Beispiele zu oben: `Vec<u8>`, `Vec<String>`, `Vec<Vec<String>>`); der Name "Vec" deutet auf die Struktur im Arbeitsspeicher hin
        - mit fester Länge: `[X; N]`, wobei `N` für die Anzahl von Elementen steht
    - in JSON: kommagetrennte Liste der Elemente in eckigen Klammern (z.B. `[42, 23]`)
    - Unterschiede zu Rust:
        - JSON ist nicht statisch typisiert und hat nur einen Listentyp, der alle möglichen Elemente enthalten kann und beliebig lang sein kann (z.B. `[42, "Hallo", []]`)

- assoziative Datenfelder: ein anderer abgeleiteter Datentyp
    - enthält nicht einzelne Elemente, sondern Paare von Schlüssel und Wert; ermöglicht schnellen Zugriff auf einen Wert anhand seines Schlüssels
    - heißt je nach Programmiersprache bzw. Datenformat "Objekt", "Map", "Dictionary" etc.
    - Typvorrat in Rust: unter anderem `HashMap<K, V>`, wobei `K` der Schlüsseltyp ("key") und `V` der Werttyp ("value") ist; der Name "HashMap" deutet auf die Struktur im Arbeitsspeicher hin
    - in JSON: kommagetrennte Liste von Wertpaaren in geschweiften Klammern, die Wertpaare jeweils mit Doppelpunkt getrennt; der Schlüssel muss ein String sein (z.B. `{"Karo": 9, "Herz": 10, "Pik": 11, "Kreuz": 12}`)
    - wiederum Unterschied zu Rust: Werte verschiedener Typen dürfen gemischt werden

- Damit ist alles beschrieben, was JSON kann, doch es gibt so viel mehr! Beispiele für abgeleitete Datentypen:
    - Mengen ("Sets"): wie Listen, aber ohne definierte Ordnung
        - in Rust: `HashSet<X>`; der Name "HashSet" deutet auf die Struktur im Arbeitsspeicher hin
    - Graphen: Netzwerke von Knoten mit Kanten dazwischen (z.B. Straßenkarte: Kreuzungen = Knoten, Straßen = Kanten)
        - in Rust: kein entsprechender Typ in der Standardbibliothek
        - ähnlich wie bei Listen und assoziativen Datenfeldern viele verschiedene Implementationsstrategien, aber bei Graphen gibt es keinen "Standardansatz"
    - Außerdem ist JSON relativ restriktiv in der Auswahl an PODs. Andere Serialisierungsformate unterstützen zum Beispiel Datums- oder Zeitangaben, Referenzen auf externe Dateien etc.

- im Gespräch erwähnt
    - Hörtipp: [Douglas Crockford bei CoRecursive zur Entstehung von JSON](https://www.corecursive.com/json-vs-xml-douglas-crockford/)
    - Lehrbuch: [Structure and Interpretation of Computer Programs](https://en.wikipedia.org/w/index.php?title=Structure_and_Interpretation_of_Computer_Programs&oldid=1244483390)
