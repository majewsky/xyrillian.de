- Wortklärung
    - "effektiv": die Fähigkeit besitzend, eine Aufgabe erfolgreich zu erledigen
    - "effizient": fähig, viel Leistung in Relation zum Aufwand zu erbringen
    - Leitsatz: "Der einzige Weg, ein Programm effizienter zu machen, ist, weniger zu tun." (Quelle unklar)
    - Welche Wege gibt es, weniger zu tun?
    - hängt vom Optimierungsziel ab: Speicherverbrauch vs. Zeitverbrauch vs. Energieverbrauch vs. Kosten bei Zugriff auf externe Systeme -> Analogie zum Navigationssystem im Auto

- [Caching](https://de.wikipedia.org/wiki/Cache ): Ergebnisse wiederverwenden
    - Abruf aus dem Arbeitsspeicher ist effizienter (in Zeit und Energie) als Abruf aus dem Netzwerk

- [Miniaturisierung](https://de.wikipedia.org/wiki/Miniaturisierung ): weniger Elektronen durch die Gegend schaufeln
    - "viele Zahlen für wenig Wärme"
    - Transistoren sind ähnlich wie Relais-Schalter, die durch Auffüllen und Leeren einer Elektronenfalle umschalten
    - kleinere Transistoren = weniger Elektronen schaufeln = weniger Stromverbrauch (und weniger Abwärme)
    - zeitlicher Verlauf der [Strukturgrößen](https://de.wikipedia.org/wiki/Strukturgr%C3%B6%C3%9Fe ) anhand von x86-Prozessoren (mit bildlichem Vergleich zu mechanischen Schaltern im Maßstab 1.000.000:1)
        - 1971 (Intel 4004): 10µm (vgl. 10m = nur im Cartoon, etwa 3-4 Etagen hoch)
        - 1993 (Intel Pentium): 800nm (vgl. 80cm = Weichenhebel in manuellem Stellwerk, kleiner Scherenstromabnehmer)
        - 2007 (ohne Produktbezeichnung): 65nm (vgl. 65mm = Lichtschalter)
        - 2019 (AMD Ryzen): 7nm (vgl. 7mm = Kippschalter an der Schreibtischlampe)

- [heterogene CPU-Architektur](https://en.wikipedia.org/wiki/Heterogeneous_computing ), z.B. "big.LITTLE" bei ARM
    - kleine CPU-Kerne (mit geringem Stromverbrauch) für Routinearbeit und Hintergrundprozesse
    - große CPU-Kerne (die meistens schlafen) für energieintensive Anwendungen (3D-Grafik bei Spielen, Enkodierung/Dekodierung bei Videoanrufen, etc.)
    - außerdem Betrieb der kleinen Kerne sehr stoßartig: lieber kurz mit hoher Taktfrequenz und dann gleich wieder schlafen als permanent langsam laufen lassen (Randbemerkung: anders als z.B. bei Verbrennungsmotoren, aber Parallelen zu Elektromotoren)

- im Zusammenhang damit: "Aufwachverschmelzung" (z.B. systemd.timer)
    - Idee: für geplante Hintergrundaktivitäten die Takte so verschieben, dass möglichst viele Aktivitäten gleichzeitig ausgeführt werden und die CPU-Kerne möglichst selten aufwachen müssen

- ebenfalls im Zusammenhang damit: Push Gateways auf iOS/Android
    - Netzwerkverbindungen = Energieverbrauch (für das Funken und das Verarbeiten in der CPU)
    - Netzwerkverbindung aufhalten = Energieverbrauch (für Keepalive-Pakete)
    - Idee: statt einer Netzwerkverbindung pro Messenger-App eine zentrale Netzwerkverbindung zu einem zentralen Relay-Server vom Betriebssystem-Anbieter
    - aber: aus Datenschutzsicht problematisch

- Xyrills Abendgedanken: Wie kann es sein, dass alles soviel effizienter ist, aber wir machen nicht mehr damit? (siehe klassische Floskel: "modernes Smartphone ist mächtiger als die Computer in der Apollo-Kapsel")
    - wir erwarten gleichzeitig immer mehr von den Geräten (z.B. Grafikqualität: gleich komplizierte App, aber früher auf 80x25-Terminal und heute auf Full-HD-Bildschirm)
    - vgl. [Jevons-Paradoxon](https://de.wikipedia.org/wiki/Jevons-Paradoxon ): mehr Effizienz ermöglicht mehr Anwendungen, sodass letztlich der Energieverbrauch gleich bleibt oder sogar steigt
    - waghalsige Prognose: viel mehr E-Ink-Displays, wenn Energieeffizienz wichtiger wird
