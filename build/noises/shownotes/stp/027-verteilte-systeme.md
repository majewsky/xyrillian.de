[Laut Wikipedia](https://de.wikipedia.org/w/index.php?title=Verteiltes_System&oldid=223080049 ) eine überraschend breite Definition:

> Ein verteiltes System ist \[...] ein Zusammenschluss unabhängiger Computer, die sich für den Benutzer als ein einziges System präsentieren.

- Motivationen für verteilte Systeme
    - Kommunikation zwischen Computern mit verschiedenen Eigentümern (z.B. Chatnetzwerk, Banküberweisung)
    - Teilen von Ressourcen (z.B. Druckerfreigabe)
    - Hochverfügbarkeit
    - Lastverteilung

- Grundlage des Problems: [CAP-Theorem](https://de.wikipedia.org/w/index.php?title=CAP-Theorem&oldid=222904512 )
    - Es gibt drei extrem wünschenswerte Eigenschaften. Wähle zwei.
    - **Konsistenz** (Consistency): Lesevorgänge liefern immer die Daten, die zuletzt geschrieben wurden.
    - **Verfügbarkeit** (Availability): Alle Anfragen an das System werden stets beantwortet.
    - **Partitionstoleranz**: Das System kann bei Störungen des Netzwerkes oder einzelner Computer weiterarbeiten.

- Beispiele bekannter Systeme gemäß CAP-Klassifikation
    - CP: Banküberweisung, Geldautomat
    - CA: Datenbanksysteme (siehe Folge 12)
    - AP: DNS (siehe Folge 18)

- Komplexbeispiel: Verteilte Datenbank mit [Raft-Konsensalgorithmus](https://en.wikipedia.org/w/index.php?title=Raft_(algorithm)&oldid=1095427922 )
    - Computernetzwerk aus N gleichartigen Teilnehmern (N meist ungerade, um Pattsituationen zu vermeiden)
    - pro **Wahlperiode** ein **Anführer**, alle anderen folgen dem Anführer
    - Schreibvorgänge gehen an den Anführer, werden von ihm an die Folger verteilt, und erst quittiert, wenn eine Mehrzahl der Folger Vollzug meldet
    - Anführer sagt regelmäßig (Größenordnung 150-300 ms) allen Folgern Bescheid, dass er noch anführt
    - bei Ausbleiben der Meldung vom Anführer startet jeder Folger eine neue Wahlperiode und stimmt für sich selbst als neuen Anführer
    - pro Wahlperiode:
        - wer jemand anderen abstimmen sieht, bevor er selbst abstimmt, folgt sofort der gesehenen Stimme
        - wer die Mehrheit der Stimmen auf sich vereinigt, ist der neue Anführer
    - bei Pattsituationen beginnt nach einer Wartefrist die nächste Wahlperiode (Wartefrist ist pro Teilnehmer randomisiert, um Kettenpatt zu vermeiden)
    - CAP-Bewertung: CA (nicht partitionstolerant, weil eine Mehrheit der Teilnehmer erforderlich ist)

- hier nicht betrachtet: [Byzantinische Fehler](https://de.wikipedia.org/w/index.php?title=Byzantinischer_Fehler&oldid=219637933 )
    - Paper: [The Saddest Moment](https://web.archive.org/web/20140124011140/research.microsoft.com/en-us/people/mickens/thesaddestmoment.pdf )
    - Analogie zu Wahlcomputern

- zum Nachlesen: [Irrtümer der verteilten Datenverarbeitung](https://de.wikipedia.org/w/index.php?title=Fallacies_of_Distributed_Computing&oldid=223882907 )

- Sidebar: [Volunteer Computing](https://de.wikipedia.org/w/index.php?title=Volunteer-Computing&oldid=204027523 ) (BOINC, SETI@Home, Folding@Home, etc.)
    - siehe [Pentacast 56](https://c3d2.de/news/pentacast-56-folding.html )

- Nachtrag: geführte Visualisierung des Raft-Algorithmus (englisch): <https://thesecretlivesofdata.com/raft/>
