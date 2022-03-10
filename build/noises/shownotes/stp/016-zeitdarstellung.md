- Zeiteinteilung ist ganz schön unhandlich
    - Kalender orientieren sich an astronomischen Fakten, die leider keine schönen ganzzahligen Einteilungen erlauben (1 Sonnenjahr = 365,242... Tage, 1 Mondmonat = 29.530... Tage, etc.)
    - [präzise Definition von kleinen Zeiteinheiten wie Sekunden](https://de.wikipedia.org/wiki/Atomuhr ) zeigt, dass die astronomischen Fakten schwanken -> [Schaltsekunden](https://de.wikipedia.org/wiki/Schaltsekunde ); später mehr
    - Zielkonflikt: möglichst gleichmäßige Zeitentwicklung oder möglichst intuitive Handhabung
    - [Internationale Atomzeit (TAI)](https://de.wikipedia.org/wiki/Internationale_Atomzeit ): exakt gleichmäßig (im Rahmen der Genauigkeit der Atomuhr), berücksichtigt nicht Schwankungen in der Tageslänge durch Änderungen der Erdrotationsgeschwindigkeit
    - [Koordinierte Weltzeit (UTC)](https://de.wikipedia.org/wiki/Koordinierte_Weltzeit ): TAI zuzüglich Schaltsekunden (zurzeit TAI = UTC + 37s), sodass der Tagesbeginn am nullten Breitengrad immer bei 00:00:00 Uhr UTC bleibt

Wie können Computer Zeit darstellen?

- Option 1: separate Zahlen für Jahr, Monat, Tag, Stunde, Minute, Sekunde
    - Verwendung vor allem an Stellen, wo der Nutzer es sehen kann (Zeitanzeigen, Zeitstempel in Dateinamen etc.)
    - Vorteil: sofort menschenlesbar
    - Nachteil: muss für Eindeutigkeit immer mit Zeitzoneninformation verknüpft sein; Gültigkeit von zukünftigen Zeitangaben evtl. durch nicht vorhergesehene Zeitzonenänderungen beeinträchtigt
    - Nachteil: textuelle Darstellung je nach Kulturkreis uneindeutig (z.B. "4/7/2021" kann entweder für den 4. Juli oder den 7. April stehen)

- Option 2: eine einzelne Zahl, die die vergangene Zeit in einer bestimmten Einheit (z.B. Sekunden oder Tage) seit einem Startzeitpunkt (der [Epoche](https://de.wikipedia.org/wiki/Epoche_(Chronologie) )) angibt
    - z.B. Unix: Sekunden seit 1. Januar 1970 00:00:00 Uhr UTC (bzw. 01:00:00 Uhr deutscher Zeit)
    - z.B. [julianisches Datum](https://de.wikipedia.org/wiki/Julianisches_Datum ): Tage seit 1. Januar 4713 v.u.Z.
    - z.B. Excel, LibreOffice Calc, etc.: Tage seit 1. Januar 1900, mit Nachkommastellen für Zeitpunkte innerhalb eines Tages

- Probleme durch Überlauf in Datumsfeldern
    - [Y2K-Problem](https://de.wikipedia.org/wiki/Jahr-2000-Problem ): Überlauf in der Jahreszahl, die damals zweistellig angegeben war (z.B. "76" für 1976)
    - [Y2038-Problem](https://de.wikipedia.org/wiki/Jahr-2038-Problem ): Überlauf in einem 32-Bit-Unix-Zeitstempel (am Dienstag, dem 19. Januar 2038 um 04:14:07 Uhr deutscher Zeit)
    - [GPS-Wochen-Rollover](https://de.wikipedia.org/wiki/GPS-Woche ): Überlauf in der Wochenzahl des Zeitstempel-Formats von GPS; erstmals 1999, zuletzt 2019, nächstes Mal 2038

- Probleme durch astronomische/politische Fakten
    - Berechnungen wie "10 Tage dazuaddieren" mühselig aufgrund Kalenderregeln (Monatswechsel, Schalttage etc.)
        - Beispiel: in einer Terminserie wie "täglich um 15 Uhr" hat wegen Sommerzeitwechsel nicht jeder Termin genau 24 Stunden Abstand zueinander
        - Beispiel: [Sommerzeitumstellung bei der Eisenbahn](https://de.wikipedia.org/wiki/Sommerzeit#%C3%96ffentliche_Verkehrsmittel ) (diesen Abschnitt können wir im Prinzip direkt vorlesen)
    - Schaltsekunden lassen UTC springen und erzeugen unerwartete Zeitzustände wie "23:59:60"
    - [Zeitzonendefinitionen](https://de.wikipedia.org/wiki/Zeitzonen-Datenbank ) folgen politischen Entscheidungen
        - z.B. Deutschland besteht in tzdata aus zwei Zeitzonen
        - z.B. 2015 [Chaos in der Türkei](https://codeofmatt.com/on-the-timing-of-time-zone-changes/ ) durch eine Änderung der Sommerzeitregel mit zu wenig Vorlauf

- Probleme durch technische Einflüsse
    - Hardware-Uhren sind nicht hinreichend genau und driften mit der Zeit: in verteilten Systemen mögliche Verwirrung über die Reihenfolge von Aktionen
    - Zeitsynchronisation mittels NTP (Network Time Protocol): im Moment der Richtigstellung der Uhr mögliche Verwirrung durch einen kleinen Sprung rückwärts in der Zeit
    - Suspend/Standby: im Moment des Wiederanschaltens mögliche Verwirrung durch einen drastischen Sprung vorwärts in der Zeit
    - monotonische Uhren: zählen die Zeit, die der Computer seit Systemstart gelaufen sind (z.B. mittels [TSC](https://en.wikipedia.org/wiki/Time_Stamp_Counter ); damit zuverlässige Messung von Laufzeitspannen etc.

- positiver Ausblick: Es könnte noch viel schlimmer sein!
    - gregorianischer Kalender hat sich weltweit für den Geschäftsverkehr durchgesetzt
    - Beispiel für einen alternativen Kalender: [Japanische Zeitrechnung](https://de.wikipedia.org/wiki/Japanische_Zeitrechnung )

- kryptografisch abgesicherte Zeitstempel: [RFC Podcast #15 Crypto for the Masses](https://requestforcomments.de/archives/559?t=07%3A55 )
