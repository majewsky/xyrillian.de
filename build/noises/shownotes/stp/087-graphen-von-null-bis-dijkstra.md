- oft erwähnt, aber nie systematisch eingeführt: der Begriff **Graph**
    - [Graph](https://de.wikipedia.org/w/index.php?title=Graph_(Graphentheorie)&oldid=260114694): eine Menge von Knoten mit Kanten dazwischen
    - z.B. Nahverkehrsnetz: Knoten = Haltestellen, Kanten = Direktverbindung zwischen zwei Haltestellen
    - z.B. soziale Medien: Knoten = Personen, Kanten = "ist befreundet/folgt"
    - anders als bei Listen (STP077) oder Maps (STP079) große Variabilität darin, was ein Graph fundamental ist
        - z.B. gerichtete vs. ungerichtete Graphen: Haben die Kanten eine Richtung oder nicht? (siehe Beispiele oben)
    - häufig verwendete Unterarten von Graphen:
        - [Baum](https://de.wikipedia.org/w/index.php?title=Baum_(Graphentheorie)&oldid=260160360): ayzklischer ungerichteter zusammenhängender Graph (z.B. siehe STP077/STP079: zum effizienten Abspeichern von sortierten Listen bzw. Maps)
        - azyklischer gerichteter Graph (z.B. Versionsverwaltung, Familienstammbaum, Firmen-Organigramm)
        - [planarer Graph](https://de.wikipedia.org/w/index.php?title=Planarer_Graph&oldid=254694423): siehe ["Untangle"](https://www.chiark.greenend.org.uk/~sgtatham/puzzles/js/untangle.html)

- Wie durchschreitet man einen Graphen? -> Suchalgorithmen, im allgemeinsten Fall die zwei wichtigsten "uninformierten" Suchalgorithmen
    - [Breitensuche](https://de.wikipedia.org/w/index.php?title=Breitensuche&oldid=250572859): von einem Startpunkt pro Runde alle Nachbarn von bereits besuchten Knoten besuchen
    - [Tiefensuche](https://de.wikipedia.org/w/index.php?title=Tiefensuche&oldid=260875992): vom Startpunkt aus immer erst einen Nachbarn besuchen und dann dort wieder einen Nachbarn; sobald keine weiteren unbesuchten Nachbarn möglich sind, Backtracking zum vorherigen Knoten, bis ein neuer unbesuchter Nachbar gefunden wird
    - sowohl Breiten- als auch Tiefensuche erzeugen beim Durchschreiten des Graphen einen Suchbaum, der ein [Spannbaum](https://de.wikipedia.org/w/index.php?title=Spannbaum&oldid=258057046) ist (vgl. Spanning Tree bei Ethernet in STP028)
    - Abwägung: Tiefensuche ist einfacher zu implementieren und hat weniger Speicheraufwand (braucht nur eine "besucht-Markierung" und keine Kandidatenliste), [kann aber in der Praxis problematisch sein :)](https://xkcd.com/761/)

- Beobachtung: Breitensuche findet garantiert den kürzesten Pfad vom Startpunkt zu einem gewünschten Endpunkt
    - Abwandlung für Wegfindung: [Dijkstra-Algorithmus](https://de.wikipedia.org/w/index.php?title=Dijkstra-Algorithmus&oldid=259180765)
    - Voraussetzung: Kanten haben auch noch ein Gewicht (analog zu einer Streckenlänge)
    - Breitensuche merkt sich bei jedem Knoten den kürzesten Gesamtweg zu diesem Knoten
    - beim Auswählen des nächsten Besuchskandidaten Wahl des Knotens mit dem kürzesten Gesamtweg
    - Erweiterung zum [A\*-Algorithmus](https://de.wikipedia.org/w/index.php?title=A*-Algorithmus&oldid=258638060) mittels Ergänzung um eine Heuristik, die beschreibt, wie nahe gefundene Knoten am Ziel sind (damit je nach Situation enorme Effizienzsteigerung möglich, siehe z.B. [dieser Blogpost der Factorio-Entwickler](https://factorio.com/blog/post/fff-317))

- im Gespräch erwähnt: [#WissPodWeihnacht 3: Das Haus vom Niko…hä?!? mit Geschichten aus der Mathematik](https://wissenschaftspodcasts.de/wisspodcast/wisspodweihnacht-3-das-haus-vom-niko-hae-mit-geschichten-aus-der-mathematik/)

