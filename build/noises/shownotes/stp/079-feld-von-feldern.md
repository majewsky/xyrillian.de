- Rückbezug STP077: Liste von Listen
    - verschiedene Datenstrukturen implementieren dasselbe Datenformat mit unterschiedlichen Laufzeitcharakteristiken
    - heute analog dazu Besprechung von assoziativen Datenfeldern

- Rückbezug STP072: assoziative Datenfelder (wir sagen hier "Maps"; heißt je nach Programmiersprache auch "Objekt", "Dictionary", etc.)
    - eine Sammlung von Wertepaaren mit jeweils einem eindeutigen Schlüssel ("Key") und einem zugeordneten Wert ("Value")
    - z.B. Wörterbuch: Schlüssel ist das Lemma, Wert ist die Definition
    - z.B. Verzeichnis im Dateisystem: Schlüssel ist der Dateiname, Wert ist die Datei (bzw. das Unterverzeichnis)
    - praktisches Beispiel: Wörter zählen
    - es geht hier nur um Wertepaare, die aktiv im Speicher vorliegen, nicht um berechnete Abbildungen (Gegenbeispiel: Suchmaschine bildet von Suchanfrage auf Ergebnisseite ab)
- Ansatz 1: "Sammlung von Wertepaaren" als Liste umsetzen
    - sofern unsortiert, ist die algorithmische Komplexität für die meisten Operationen sehr schlecht
        - z.B. Suchen nach einem bestimmten Schlüssel in linearer Zeit (`O(n)`), wenn man die Liste von vorne nach hinten durchgehen muss
        - z.B. Einfügen eines neuen Schlüssels auch nur in linearer Zeit, weil man einen evtl. existierenden Eintrag mit demselben Schlüssel finden und ersetzen
    - praktikable Umsetzung unter Verwendung einer sortierten Liste, um bestimmte Schlüssel schnell aufzufinden
        - meistens mittels eines balancierten Baums, z.B. in Rust [BTreeMap](https://doc.rust-lang.org/std/collections/struct.BTreeMap.html) oder in C++ [std::map](https://cppreference.com/w/cpp/container/map.html)
        - Problem: Schlüssel müssen sortierbar sein (Beispiel: Strings, Ganzzahlen; Gegenbeispiel: komplexe Zahlen, Fließkommazahlen, ungeordnete Mengen)
        - Problem: auch bei sortierbaren Schlüsseln können Vergleiche teuer sein
        - Wie können wir die Vergleiche günstiger machen?

- Ansatz 2: Vorsortierung anhand von Streuwertfunktionen (Hashes)
    - Idee: bei Suche nach einem bestimmten Schlüssel wird ein Hash des Schlüssels ermittelt (in Größenordnung einer kleinen Ganzzahl, z.B. 4 oder 8 Byte; nicht ein kryptografisch starker Hash wie SHA-2 mit 28-64 Bytes); damit schnelle Vergleiche möglich -> [Hashtabelle](https://de.wikipedia.org/w/index.php?title=Hashtabelle&oldid=255127849)
    - Strategie 1: Gruppierung der Wertepaare in "Buckets" anhand der ersten K Bits des Hash-Wertes des Schlüssels
        - K richtet sich nach der Gesamtanzahl an Wertepaaren (z.B. für 50 Elemente könnte man K = 5 oder 6 und damit 2^K = 32 oder 64 Buckets wählen)
        - innerhalb der Buckets einfache Listen von Wertepaaren, die in diesen Bucket fallen; idealerweise nur ein Eintrag pro Bucket
        - durch dynamisches Anpassen der Bucket-Anzahl Abwägung zwischen Speicheraufwand für evtl. leere Buckets einerseits und Zeitaufwand für Durchsuchen übervoller Buckets andererseits
    - Strategie 2: wie 1, aber jeder Bucket kann maximal ein Element aufnehmen
        - bei Kollisionen strategisches Ausweichen auf andere Buckets in einem systematischen Muster, bis ein freier Platz gefunden wird
        - große Variation von Implementationsstrategien für dieses "strategische Ausweichen"

- Vorwärtsbezug zu STP080: Komplexitätsattacke auf Hashmaps (eine Instanz von [CWE-400](https://cwe.mitre.org/data/definitions/400.html))
    - Problem, wenn Hashmap mit nutzerdefinierten Daten gefüllt wird: Angreifer könnte Daten so wählen, dass sie alle im selben Bucket landen wollen
    - damit extrem hohe Rechenlast durch ständige Kollisionen
    - Lösung: Durchmischen des Hash mittels prozessinterner Geheimzahl (siehe STP044 zu Zufallszahlen)

#### Feedback zu STP065

Mole schreibt:

> In STP065 vermutest du, dass der Ethernet-Port wegen Authentifizierung abgeschafft wurde. "Beim WLAN muss man sich authentifizieren, beim Ethernet kann man sich einfach anstecken und lossurfen." Das ist so natürlich nicht richtig. Im Enterprise-Umfeld gibt es für sowas 802.1x. Das wird übrigens auch beim WLAN zur Authentifizierung mit einem RADIUS-Server benutzt. Außerdem gibt es da SNMP-Traps. Da schaltet dann der Port ab und der Admin bekommt eine Nachricht, wenn die falsche MAC antwortet. Das gab es schon lange vor WLAN.
>
> Nächstes Problem: Du erklärst, das die NUCs vom VOC für das Streaming zuständig sind. Das ist so auch nicht richtig. Die machen nur das Encoding in die Subformate, wenn der Talk schon geschnitten ist. Für das Streaming sind die Cubes, ordentliche PCs mit extrem teuren Spezial-SDI-Encoding-Karten. Die Kisten brauchen auch echt Leistung, weil die das live machen.
