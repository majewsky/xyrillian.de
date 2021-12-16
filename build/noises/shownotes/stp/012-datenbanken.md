[Laut Wikipedia:](https://de.wikipedia.org/wiki/Datenbank )

> Eine **Datenbank**, auch **Datenbanksystem** genannt, ist ein System zur elektronischen Datenverwaltung. Die wesentliche Aufgabe einer Datenbank ist es, große Datenmengen effizient, widerspruchsfrei und dauerhaft zu speichern und benötigte Teilmengen in unterschiedlichen, bedarfsgerechten Darstellungsformen für Benutzer und Anwendungsprogramme bereitzustellen.

* Beispiele für Datenbanken _im weitesten Sinne_: Passen diese mit der Definition oben zusammen?
    * Spreadsheets (Excel etc.)
        * was funktioniert: unterschiedliche, bedarfsgerechte Darstellungsformen
        * keine widerspruchsfreie Speicherung: [Beispiel Gen-Namen-Fehlformatierung](https://www.heise.de/-6165902 )
        * keine großen Datenmengen: [Beispiel verlorene Covid-Patienten aufgrund Beschränkung auf 1 Mio. Zeilen](https://www.theguardian.com/politics/2020/oct/05/how-excel-may-have-caused-loss-of-16000-covid-tests-in-england )
    * Dateisystem
        * was funktioniert: große Datenmengen effizient und dauerhaft speichern
        * keine widerspruchsfreie Speicherung, da keine Beziehungen zwischen Dateiinhalten möglich
        * keine Unterstützung für bedarfsgerechte Darstellungsformen: jede Datei hat genau ein Format

* Dateisysteme sind eine besondere Form von **Key-Value-Datenbank**
    * Datenbasis = ein großer Haufen von **Values** (Werte, im allgemeinen Fall beliebige Dateien), die alle einen Namen (**Key**) haben
    * keine Datenbank im engeren Sinne (keine Sicherstellung von Widerspruchsfreiheit, keine Unterstützung für bedarfsgerechte Darstellungsformen)
    * aber in der Praxis eine sinnvolle Ergänzung zu stark strukturierten Datenbanken, da die einfache Datenstruktur es ermöglicht, sehr groß zu skalieren (bis Milliarden/Billionen von individuellen Objekten)
    * Praxisbeispiel: Cache für vorgerenderte Wikipedia-Seiten

* Datenbanken im engeren Sinne: [relationale Datenbanken](https://de.wikipedia.org/wiki/Relationale_Datenbank )
    * Daten liegen in einer bestimmten Struktur vor: gleichartige Objekte sind jeweils als **Zeilen** in **Tabellen** abgelegt, deren **Spalten** im Voraus definiert sind (diese Grundstruktur nennt man **Schema**)
    * "relational": Kreuzreferenzen zwischen Tabellenspalten (Beispiel: wenn die Tabelle "Mitarbeiter" eine Spalte "Abteilungsnummer" hat, muss es für jeden Wert in dieser Spalte eine entsprechende Zeile in der Tabelle "Abteilungen" geben), unter anderem dadurch Sicherstellung der Widerspruchsfreiheit
    * strukturierte Datenhaltung ermöglicht komplexe Abfragen (z.B. "Wieviele Mitarbeiter in der Abteilung 172 haben ein Monatsgehalt über 3000€?"), typischerweise mittels [SQL (Structured Query Language)](https://de.wikipedia.org/wiki/SQL )

* Grundkriterien für das Verhalten von relationalen Datenbanken: [ACID](https://de.wikipedia.org/wiki/ACID ) am Beispiel "Überweisung zwischen zwei Bankkonten"
    * **Atomarität (Abgeschlossenheit)**
        * Änderungen werden entweder ganz oder gar nicht vorgenommen
        * am Beispiel: Überweisung muss immer beide Bankkonten anfassen (oder komplett fehlschlagen)
    * **Konsistenzerhaltung**
        * wie oben besprochen: Prüfung von Kreuzreferenzen oder Randbedingungen
        * am Beispiel: Überweisung darf nur durchgehen, wenn das Quellkonto genug Füllstand hat
    * **Isolation (Abgrenzung)**
        * parallel laufende Änderungen dürfen nicht dieselbe Zeile anfassen
        * im Zweifelsfall muss eine Änderung warten, bis die andere fertig ist
        * am Beispiel: Deckungsprüfung im Quellkonto muss alle vorherigen Überweisungen einbeziehen
    * **Dauerhaftigkeit**
        * wird eine Änderung als abgeschlossen gemeldet, darf auch ein Systemfehler (z.B. Systemabsturz oder Stromausfall) nicht zum Verlust der Änderung führen
        * Garantierung der Dauerhaftigkeit ist nichttrivial aufgrund der vielen Schichten der Speicherhierarchie (siehe STP007)
        * Sicherstellung meist in der Datenbanksoftware selbst mittels **Write-Ahead Log** und drumherum mittels Betriebsprozeduren (Backups, Replikation)

* klassisches Problem mit relationalen Datenbanken: Wie klassifiziert man die Welt? Welche Attribute kann man bei Objekten voraussetzen?
    * Welche Spalten kann man voraussetzen? z.B. Tabelle "Personen" mit Spalten "Name", "Adresse", "Geburtsdatum", "Geburtsort" -> für alle diese Attribute gibt es Menschen, bei denen dieses Attribut nicht existiert oder unbekannt ist
    * Welche Datentypen haben Spalten? z.B. wird die Postleitzahl oft als Nummer abgelegt, sodass dann "01069 Dresden" als "1069 Dresden" fehlverarbeitet wird
    * allgemein: Entwurf des Datenbankschemas ist einer dieser Punkte, wo man sich manchmal als Programmierer bewusst wird, wie wenig man das Problemfeld eigentlich verstanden hat

* abseits von relationalen Datenbanken und SQL: [NoSQL](https://de.wikipedia.org/wiki/NoSQL )
    * großer Hype in den 2010ern vor allem aufgrund der Hoffnung, die lästige Schemaentwurf-Phase los zu werden
    * aber: Speichern von unstruktierten oder wenig strukturierten Daten verlagert das Problem nur in die Zukunft

* Arten von NoSQL-Datenbanken, die tatsächlich nützlich sind:
    * Key-Value-Datenbanken (siehe oben)
    * Dokumentendatenbanken für Volltextsuche (z.B. Suchmaschinen)
    * Graphdatenbanken (z.B. soziale Graphen)
    * Datenströme (Darstellung der Welt als ein zeitlich geordneter Strom von Ereignissen)
