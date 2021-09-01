- Rückbezug zu STP003: Logikgatter
    - damit Rechnen möglich
    - für Eingabewerte und Ausgabewerte ist ein Speicher erforderlich

- Kurzabriss zum Thema [Harvard-Architektur](https://de.wikipedia.org/wiki/Harvard-Architektur) vs. [Von-Neumann-Architektur](https://de.wikipedia.org/wiki/Von-Neumann-Architektur)

- grundsätzliche Arten von Speichern
    - flüchtiger Speicher: braucht Strom zum Halten von Daten
    - nichtflüchtiger Speicher: behält Daten auch im ausgeschalteten Zustand
    - Der ideale Speicher wäre groß, schnell, nichtflüchtig, und beliebig oft wiederbeschreibbar.
    - da wir keinen idealen Speicher haben: **Speicherhierarchie**

In der folgenden Liste der verschiedenen Hierarchiestufen sind alle konkreten Zahlenwerte nur illustrative Beispielwerte für Hardware der aktuellen Generation.

- Stufe 1: [Festplatte](https://de.wikipedia.org/wiki/Festplattenlaufwerk )
    - nichtflüchtig, groß und günstig (ab etwa 25 Euro pro TiB)
    - Zugriffsgeschwindigkeit: 2.5 bis 10 ms (bei 3 GHz: 7,5 bis 30 Millionen Takte) plus Übertragungszeit
    - Zugriff immer mindestens in Blöcken zu 4 KiB, besonders effizient beim Zugriff auf Sequenzen aufeinanderfolgender Blöcke

- Stufe 2: [SSD](https://de.wikipedia.org/wiki/Solid-State-Drive )
    - etwas teurer als Festplatten (ab etwa 90 Euro pro TiB)
    - aber Zugriffsgeschwindigkeit: 0.05 ms (bei 3 GHz: 150.000 Takte)
    - anders als bei Festplatten schneller Wahlzugriff

- Stufe 3: [Arbeitsspeicher (RAM)](https://de.wikipedia.org/wiki/Random-Access_Memory )
    - ab hier flüchtig (Vorteil: öfter wiederbeschreibbar ohne Qualitätsverlust)
    - deutlich teurer (ab etwa 4 Euro pro **GiB**, entspr. ~4000 Euro pro TiB)
    - gerade deswegen meist kleiner als die nichtflüchtigen Speicher
        - übliche Größen: von 2-4 GiB bei Smartphones oder kleinen Notebooks bis mehrere TiB bei Servern
    - Zugriffsgeschwindigkeit: 15 ns Latenz plus Übertragungszeit, insgesamt z.B. 65 ns (bei 3 GHz: 200 Takte)

- Stufe 4: [Cache (Zwischenspeicher) des Prozessors](https://en.wikipedia.org/wiki/CPU_cache )
    - ab hier fest in den Prozessor eingebaut
    - Abbilder von einzelnen Zeilen (je 64 Byte) aus dem Arbeitsspeicher
    - Größe: 4-16 MiB
    - Zugriffsgeschwindigkeit: 1-5 ns (bei 3 GHz: 3-15 Takte)

- Stufe 5: [Registerbank des Prozessors](https://en.wikipedia.org/wiki/Register_file )
    - Größe: 64 bis einige 100 Byte
    - Zugriffsgeschwindigkeit: meist 1 Takt
    - nur dieser Speicher kann vom CPU direkt für Berechnungen verwendet werden

- andere mögliche Hierarchiestufen
    - Stufe 0: Speicher auf anderen Computern im Netzwerk
    - Stufe 2,5: nichtflüchtiger RAM (noch nicht weitverbreitet)

- Speicher aus Sicht des Programmierers
    - Abstraktion: lineare Kette von bytegroßen Speicherzellen, die fortlaufend nummeriert sind (**Speicheradressen**)
    - Probleme mit dieser Abstraktion:
        - kleinste sinnvolle Schreibmengen (4/8 Bytes bei Registern, 64 Bytes bei RAM, 4 KiB bei Festplatte/SSD)
        - Caching-Verhalten
    - Einfluss auf effizientes Rechnen: viel wichtiger als schnelles Rechnen ist lokales Rechnen
    - Beispiel: Repräsentation von 2D-Feldern (z.B. Grafiken) und effizientes Bearbeiten derselben

- Quelle: ["What every programmer should know about memory"](https://akkadia.org/drepper/cpumemory.pdf)
