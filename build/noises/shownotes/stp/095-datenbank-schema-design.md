- Rückbezug STP012: [relationale Datenbanken (RDBMS)](https://de.wikipedia.org/w/index.php?title=Relationale_Datenbank&oldid=262381831)
    - RDBMS sind **tabellenbasiert**: gleichartige Objekte als **Zeilen** in **Tabellen** abgelegt, deren **Spalten** im Voraus definiert sind -> Zeilen werden auch **Records** (Datensätze) genannt
    - Rückbezug STP071/072: Tabelle ist eine Liste von Datensätzen, Datensatz ist ein Produkttyp (siehe STP072) von Einzelwerten pro Spalte
    - RDBMS sind **relational**: Kreuzreferenzen zwischen Tabellen (z.B. Datensatz für Mitarbeiter enthält eine Abteilungsnummer, damit eine Referenz in den entsprechenden Datensatz der Abteilung)
    - wie in STP012 schon erwähnt: der Entwurf eines guten Datenbankschemas ist eine Kunst (vgl. den Abendgedanken in STP072: Datenstrukturen und Datenfluss als Fundament des Programmentwurfs)

- [Normalisierung](https://de.wikipedia.org/w/index.php?title=Normalisierung_(Datenbank)&oldid=264921953): Strukturierung von Datenbankschemata in einer Art und Weise, die Redundanzen reduziert
    - Ziel: keine redundante Speicherung von Informationen (um Platz zu sparen, und um Inkonsistenzen zu vermeiden)
    - der theoretische Kanon kennt eine Reihe zunehmend strikterer Normalformen: 1NF -> 2NF -> 3NF -> BCNF -> 4NF -> 5NF
    - Diskussion anhand der Beispiele aus dem verlinkten Wikipedia-Artikel
    - übergreifendes Motiv: sollte man sich mal klar machen, aber Xyrill wird (anders als der Wikipedia-Artikel) an einigen Stellen Gegenrede leisten

- erste Normalform (1NF): einzelne Tabellenspalten dürfen nur atomare Werte enthalten, keine strukturierten Datensätze
    - siehe Beispiel: Tabelle "CD-Lieder" mit Spalte "Album", die sowohl Künstler als auch Albumtitel enthält -> in zwei Spalten aufteilen
    - außerdem: Spalte "Titelliste" enthält eine Liste von Titelnamen -> aufteilen in Spalten "Track" und "Titel", ein Datensatz pro Titel
    - Begründung in der reinen Lehre: Abfragen nach Teilwerten eines nichtatomaren Attributes werden erschwert
    - Xyrills Gegenreden:
        - moderne Datenbanken ermöglichen auch ergonomischen Zugriff in strukturierte Spalten (vgl. z.B. [der JSON-Spaltentyp in PostgreSQL](https://www.postgresql.org/docs/18/functions-json.html))
        - Denormalisierung kann das Leben einfacher machen, wenn man keine strukturierte Abfrage in die Spalte hinein benötigt (Beispiel: GC-Policies in Keppel)
    - Xyrills Anekdote: unerwartete Probleme mit numerischen IDs beim Umstieg von MaxDB auf PostgreSQL

- zweite Normalform (2NF): 1NF + keine Spalte außerhalb eines Schlüsselkandidaten hängt funktional von einer echten Teilmenge des Schlüsselkandidaten ab
    - hierbei **Schlüsselkandidat**: ein Satz von Spalten, deren Werte Datensätze eindeutig identifizieren (z.B. bei der Tabelle "CD-Lieder" die Spalten "CD-ID" und "Track") -> beim Schemadesign wird hieraus (meist) ein Kandidat zum **Primärschlüssel** gewählt (Beispiel: in Dateisystemen ist der vollständige Pfad ein Schlüsselkandidat, aber in der Praxis ist meist die Inode-Nummer der bevorzugte Schlüssel)
    - siehe Beispiel: Tabelle "CD-Lieder" mit Spalte "CD-ID" und "Track", aber Spalten mit Informationen zum Album (z.B. "Interpret" und "Albumtitel") hängen nur von der CD-ID ab, nicht vom Track
    - damit unnötige Duplikation in diesen Spalten zwischen Datensätzen betreffend dieselbe CD, sowie Risiko von Inkonsistenz zwischen diesen Datensätzen
    - Lösung: Aufspaltung in zwei Tabellen "CDs" und "Lieder", letzteres mit Bezug auf die CD-ID

- dritte Normalform (3NF): 2NF + keine Nichtschlüsselspalten hängen untereinander funktional ab
    - siehe Beispiel: Tabelle "CD" mit Spalten "Interpret" und "Gründungsjahr", aber das Gründungsjahr hängt vom Interpreten ab -> wiederum unnötige Duplikation und Risiko von Inkonsistenzen
    - Lösung: Abspaltung in eine separate Tabelle "Künstler", möglicherweise unter Einführung eines neuen synthetischen Schlüssels

- weitere Normalformen: Boyce-Codd-Normalform (BCNF), 4. Normalform (4NF), 5. Normalform (5NF)
    - Xyrills Gegenrede: die Beispiel werden zunehmend esoterisch und es sieht danach aus, dass hier vor allem akademische Probleme gelöst werden
    - ttimeless's Gegenrede: nicht die Kreativität der Nutzer:innen unterschätzen
    - vielleicht als allgemeines Prinzip: beim Design immer erst einen Primärschlüssel wählen und sich dann bei jedem neuen Attribut fragen, zu welchem Schlüssel dieses gehört -> wenn keine eindeutige Antwort möglich, ist wohl eine neue Tabelle erforderlich (z.B. häufig für M:N-Relationen)
