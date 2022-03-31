Letztes Mal hatten wir Komplikationen wegen der Zeit, diesmal wegen des Raumes. :)

* vergleiche: [Lokalisierung bei Filmen](https://de.wikipedia.org/wiki/Multimedialokalisierung )
    * Übersetzung als Untertitel, Voice-Over oder Synchronisation
    * viele Schwierigkeiten der Medienlokalisierung spiegeln sich im Softwarefall
        * Beispiel: Formulierung der Übersetzung muss sich an die verfügbare Zeit im Film anpassen
        * in Software müssen übersetzte Beschriftungen in das ursprüngliche Layout reinpassen

* Softwarelokalisierung: zwei getrennte Schritte
    * [Internationalisierung](https://de.wikipedia.org/wiki/Internationalisierung_(Softwareentwicklung) ) (**i18n**): Gestaltung des Programms derart, dass es einfach an verschiedene Sprachen und Kulturen angepasst werden kann -> Aufgabe für den Entwickler
    * [Lokalisierung](https://de.wikipedia.org/wiki/Lokalisierung_(Softwareentwicklung) ) (**l10n**): der Akt der Anpassung an eine Zielsprache -> Aufgabe für den Übersetzer

* Aspekte von Internationalisierung anhand der Locale-Einstellungen in Unix-Systemen (siehe [`man 7 locale`](https://de.wikipedia.org/wiki/Alphabetische_Sortierung#Sortierregeln_nach_Sprachen ))
    * **LC\_MESSAGES**: Sprache für übersetzte Texte
    * **LC\_TIME**, **LC\_DATE**: Darstellung von Datums- und Zeitangaben (siehe [XKCD 2562](https://xkcd.com/2562/ ) und [XKCD 1179](https://xkcd.com/1179/ ))
    * **LC\_COLLATE**: alphabetische Sortierreihenfolge ([siehe Wikipedia](https://de.wikipedia.org/wiki/Alphabetische_Sortierung#Sortierregeln_nach_Sprachen ), z.B. Deutsch sortiert "ä" wie "a", aber Dänisch sortiert "ä" hinter "z")
    * **LC\_NUMERIC**, **LC\_MONETARY**: Darstellung von Zahlen und Währungsangaben (z.B. Dezimalpunkt vs. Komma, Tausendertrennzeichen, Währungszeichen vor oder hinter dem Wert)
    * **LC\_NAME**: Anreden (z.B. "Herr" und "Frau" im Deutschen)
    * **LC\_MEASUREMENT**: Größenangaben in imperialen oder metrischen Einheiten
    * **LC\_PAPER**: Standard-Papierformat für Dokumente (z.B. DIN A4 in Deutschland, aber Letter in USA)
    * **LC\_CTYPE** (_character type_): Standard-Zeichensatz für Dateien
    * usw.

* mögliche Werte für eine Locale-Wahl in Unix
    * klassisch: `de_DE`, `de_AT`, `de_CH`, `en_US`, `en_GB` usw.
    * mit UTF-8-Präferenz: `de_DE.UTF-8`, `en_US.UTF-8` usw.
    * Spezialfall: `LC_MESSAGES=C` zeigt Texte, wie sie im originalen Quelltext stehen
    * gemischte Konfiguration möglich, z.B. [Xyrill hat fast überall `de_DE.UTF-8`, aber `LC_MESSAGES=C`](https://github.com/majewsky/system-configuration/blob/c42f3e38d6d84a10b69ada6eb59d7617065b3bfc/hologram-base.pkg.toml#L63-L68 )

* Aufgaben des Softwareentwicklers bei i18n
    * Verwendung von internationalisierbaren Betriebssystemfunktionen oder Programmbibliotheken, wann immer Ein- und Ausgaben lokalitätsabhängig sind
    * Testen von grafischen Oberflächen im Right-to-Left-Modus (in Sprachen mit RTL-Schrift läuft auch das GUI-Layout von rechts nach links, siehe z.B. [arabische Wikipedia](https://ar.wikipedia.org/ ))
    * Formulierung von übersetzbaren Texten derart, dass die Übersetzer möglichst gut arbeiten können
        * Angabe von Kontext (z.B. Beschriftungstext "Match" -> Ist das das Verb "Abgleichen" oder das Substantiv "Treffer"? Oder gar "Streichholz"?)
        * bei Textbaustein mit Zahlenparameter Berücksichtigung sprachabhängiger Pluralformen (im meist englischen Original wird ein Singulartext und ein Pluraltext bereitgestellt, z.B. `1 match` und `%d matches`; die Übersetzer können je nach Notwendigkeit beliebig viele Formen angeben)

* Arbeitsablauf für Übersetzung am Beispiel [GNU gettext](https://www.gnu.org/software/gettext/manual/gettext.html )
    * Programmierer fügt Textmeldungen ("Nachrichten") im englischen Original in den Quelltext ein
    * Maintainer erstellt mittels Hilfsprogramm aus dem Quelltext einen Nachrichtenkatalog für jede gewünschte Zielsprache (meist automatisiert, z.B. täglich oder wöchentlich)
    * Übersetzer erhält den Nachrichtenkatalog und füllt die Übersetzungen ein
    * zukünftige Änderungen der Originalnachrichten werden mit den bestehenden Nachrichtenkatalogen zusammengeführt
    * Anekdote: `x-test`

* Fundstück: [Amtssprachen der Europäischen Union](https://de.wikipedia.org/wiki/Amtssprachen_der_Europ%C3%A4ischen_Union )
