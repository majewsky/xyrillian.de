Wir lesen die englische Erstausgabe von 1975, die [im Internet Archive](https://archive.org/details/MythicalManMonth) digitalisiert wurde.

- zum Autor: [Frederick P. Brooks, Jr.](https://de.wikipedia.org/w/index.php?title=Frederick_P._Brooks&oldid=229001327) (US-amerikanischer Informatiker, 1931-2022)
    - in den 60ern Entwicklungsleiter bei IBM für das Betriebssystem OS/360 (eines der ersten Systeme mit Unterstützung für persistente Massenspeicher mit Wahlzugriff, sprich: Festplatten)
    - gilt als Erfinder des Bytes mit 8 Bit; Wikipedia zitiert Brooks aus einem Interview von 2010:

> „Die wichtigste Entscheidung, die ich je getroffen habe, war, die IBM-360-Serie von einem 6-Bit-Byte auf ein 8-Bit-Byte umzustellen und damit die Verwendung von Kleinbuchstaben zu ermöglichen. Diese Veränderung verbreitete sich überall.“

- zum Buch: "Vom Mythos des Mann-Monats" ("The Mythical Man-Month") dokumentiert Brooks' Erfahrungen als Projekt-Manager in einem der frühen großen Software-Projekte
    - nicht als Lehrbuch angelegt, weil es dafür "noch zu früh" gewesen sei, sondern als Serie persönlicher Eindrücke und Lektionen
    - fast jedes Kapitel ist ein alleinstehendes Essay
    - wir wollen schauen, ob diese Lektionen auch heute noch Bestand haben
    - in der ersten Hälfte des Buches (Kapitel 2-7) ist das übergreifende Motiv, wie sich die Arbeit in großen Teams von kleinen Teams unterscheidet

- Hörempfehlung: [Pentaradio vom Dezember 2022](https://c3d2.de/news/pentaradio24-20221227.html)
    - anlässlich Brooks' Tod hatten wir über seinen anderen wichtigen Text "No Silver Bullet" und das Thema Komplexität gesprochen

### Kapitel 1: Die Teergrube

- Teergruben-Analogie: steckt ein Tier mit einem Bein in der Teergrube, kann es das Bein meist aus eigener Kraft herausziehen; alle Beine in der Teergrube sind hingegen meist ein Todesurteil
    - analog hierzu: Zähheit/Trägheit in der Software-Entwicklung nicht aus einem bestimmten Grund, sondern aufgrund mehrerer parallel wirkender Gründe

- Programme können von Einzelpersonen oder Zweierteams erschaffen werden
    - These: aber nicht "Programmsystemprodukte"

- vom Programm zum Programmprodukt: Generalisierung (Nützlichkeit für eine ganze Klasse von Situationen), Testabdeckung (siehe STP038), Dokumentation, fortlaufende Wartung -> heute sagt man "Produktisierung" dazu
    - Brooks schätzt hier dreimal mehr Aufwand als für das reine Programm
    - Xyrill würde heute sogar einen größeren Faktor schätzen, vllt. 4 oder 5: im Bereich der Programmierproduktivität gab es mehr Innovationen beim initialen Schreiben als bei Maintenance (siehe Xyrills Kritik an Ruby als einzeilerorientiert; siehe auch LLMs, aber das verschiebt sich eventuell gerade)
    - Parallele zu "warum setzt sich Open Source nicht durch", siehe Besprechnung von ["Open Source on its own is no alternative to Big Tech"](https://berthub.eu/articles/posts/open-source-by-itself-is-no-alternative-for-big-tech/) im [Pentaradio vom Dezember 2024](https://c3d2.de/news/pentaradio24-20241224.html)

- vom Programm zum Programmsystem: Einbettung in ein Netzwerk anderer Programme mit konsistentem Stil und definierten Schnittstellen untereinander
    - Brooks legt hier auch Wert auf ein vordefiniertes Ressourcenbudget (CPU-Zeit, Speicher, Datenverkehr) für jedes Teilprogramm -> heute überholt; Effizienz der Belegschaft geht über Effizienz des Produktes
    - Brooks schätzt hier *mindestens* dreimal mehr Aufwand als für die reinen Programme
    - Xyrill stimmt im Großen und Ganzen zu: heute zwar öfter von Anfang an wohlgeformte Schnittstellen (vgl. Unix-Philosophie in STP006, JSON in STP071), aber andererseits gestiegene Gesamtkomplexität

- Programmsystemprodukt: ein produktisiertes Programmsystem
    - somit laut Brooks' Schätzung `3 x 3 = 9` Mal mehr Aufwand als für die einzelnen Programme

### Kapitel 2: Vom Mythos des Mann-Monats

- Problem 1: Programmierer sind schlecht im Schätzen von Zeitaufwand
    - inhärenter Optimismus: für jede Teilaufgabe schätzt man den Aufwand im Fall, dass alles gut geht
    - These: das Programm entsteht erst im Kopf und dann in der Realität; letzteres dauert länger, aber ein Programm, dass im Kopf fertig ist, fühlt sich schon fertig an
    - Programme sind kristallisierte Gedanken auf eine sehr viel direktere Art und Weise als z.B. Handwerksstücke, die kategorisch in ein physisches Artefakt umgesetzt werden müssen
    - Testphasen sind von diesem Optimismus am stärksten betroffen
        - Xyrill stimmt im Prinzip zu; allerdings wird heute versucht, Tests bei der Entwicklung einzugliedern (siehe Besprechung von TDD in STP038)

- Problem 2: konventielle Planungsmethoden funktionieren nur, wenn "Personen und Monate austauschbar sind"
    - Basiseinheit: Mann-Monat (heute sagt man Personenmonat); die Menge Arbeit, die eine Person in einem Monat verrichten kann
    - Messung von Aufwand in Personenmonaten impliziert, dass eine Aufgabe unter mehreren Personen aufgeteilt werden kann und keine Kommunikation zwischen diesen Personen erforderlich ist
        - kommt hin für z.B. Spargelstechen oder Zimmerreinigung im Hotel, aber: "Das Austragen eines Kindes dauert neun Monate, egal, wieviele Frauen der Aufgabe zugewiesen werden."
    - zwei Komplikationen: Einarbeitungszeit (linear in der Größe des Personals) und fortwährender Kommunikationsaufwand (quadratisch in der Größe des Personals)
        - Xyrill kann das 100-%-ig bestätigen; in großen Teams verbringen die hochrangigen Techniker signifikant mehr Zeit mit Meetings/Mails/Chat als mit ihrer nominalen Tätigkeit
        - quadratisches Wachstum kann teilweise begrenzt werden durch Einführung von Unterteams mit strikten Rollenverteilungen; dann tritt das [Gesetz von Conway](https://de.wikipedia.org/w/index.php?title=Gesetz_von_Conway&oldid=238711792) in Kraft

- Problem 3: der Zeitplan wird durch die Bedürfnisse des Kunden bestimmt, nicht durch die Bedürfnisse der Programmierer
    - Restaurant-Analogie: wenn man das langsam gegarte Steak in nur 5 Minuten haben will, wird der Koch sich dem entgegenstellen... oder er dreht halt die Hitze hoch und man kriegt ein schlechteres Produkt
    - Programmierprojekte richten sich aber oft dem vom Kunden und/oder Management vorgegebenen Liefertermin: "Es ist sehr schwierig, eine Schätzung energisch, plausibel und arbeitsplatzgefährdend zu verteidigen, die auf keiner quantitativen Methode beruht, durch wenig Daten gestützt und hauptsächlich durch das Bauchgefühl der Manager getragen wird."
    - wenn dann beim Testen am Ende eine Verspätung eintritt, ist die Versuchung groß, halb fertige Produkte zu liefern (vgl. "Bananaware")

- Problem 4: Was macht man, wenn dann das Kind in den Brunnen gefallen ist und man zu spät dran ist?
    - Option 1: Lieferdatum nach hinten verschieben
    - Option 2: Umfang reduzieren
        - passiert eventuell auch unter der Hand, zum Beispiel indem beim Design und beim Testen Arbeitstempo gegen Sorgfalt abgewogen wird
        - Xyrill schielt auf sein ständig wachsendes Backlog an "müsste man mal machen"-Aufgaben
    - Option 3: weitere Leute hinzufügen
        - [Brooks'sches Gesetz](https://en.wikipedia.org/w/index.php?title=Brooks%27s_law&oldid=1233664913): "Das Hinzufügen weiteren Personals zu einem verspäteten Softwareprojekt verspätet es weiter."
        - siehe Problem 2: neue Mitarbeiter müssen von den bestehenden Teammitgliedern eingearbeitet werden und erhöhen dann permanent den Kommunikationsaufwand
    - Vorschlag von Brooks: Planung basierend auf sequentiellen Teilschritten mit jeweils voneinander unabhängigen Teilaufgaben
        - Personalgröße ergibt sich aus der Anzahl jeweils unabhängiger Teilaufgaben
        - Zeitaufwand ergibt sich aus der längsten Kette von sequentiellen Teilschritten
        - Parallele zum [Gantt-Diagramm](https://de.wikipedia.org/w/index.php?title=Gantt-Diagramm&oldid=248831153), dass zu Brooks' Zeit bekannt gewesen sein sollte

### Kapitel 3: Das OP-Team ("The Surgical Team")

- Beobachtung: sehr weite Spannbreite im Produktivitätsspektrum der Software-Entwickler; die besten sind 10x so produktiv wie die schlechtesten
    - Xyrill war überrascht, dass diese Beobachtung schon so alt ist; in den 2010ern gab es einen Hype darum und man hat die produktivsten Entwickler als "Rock Stars" gehandelt
    - naive Schlussfolgerung: statt einem großen Team lieber ein kleines Team der absolut Besten
        - kleine Teams kommunizieren ja auch besser, also weniger Overhead
    - Problem: skaliert irgendwann nicht mehr
        - 10 Leute können nicht z.B. ein modernes Betriebssystem bauen (vielleicht als Programm, aber definitiv nicht als fortlaufend gewartetes Programmsystemprodukt)

- Alternativvorschlag: Teamstruktur analog zu Operations-Teams in der Chirurgie
    - ein Spezialist in der Mitte, der die Schnitte macht
    - viele Unterstützer drumherum, die die Werkzeuge anreichen und Instrumente überwachen

- genaue Rollenverteilung in diesem Vorschlag
    - Chirurg/Chefprogrammierer: erarbeitet die Design-Spezifikation und schreibt Programm, Dokumentation und Tests
    - Copilot: berät den Chirurgen in allen Design- und Implementations-Entscheidungen; fungiert als Stellvertreter des Chirurgen
    - Administrator: kümmert sich für den Chirugen um alle administrativen Aufgaben (Geldmittel, Personal, Räumlichkeiten, Inventar); bei kleineren Teams teilen sich evtl. mehrere Teams einen Administrator
    - Editor: bringt die vom Chirurgen ins Unreine geschriebene oder diktierte Dokumentation in Reinform
    - außerdem werden Sekretariatsrollen beschrieben; z.B. ein Schreiber, der Code auf Lochkarten bringt und Programmausgaben sammelt und archiviert
    - weitere Randrollen, die von mehreren Teams geteilt werden können: Werkzeugmacher, Tester, Language Lawyer (jemand, der die jeweilige Programmiersprache in- und auswendig kennt und hochspezialisierte Optimierungen oder besonders leichtgängige Strukturen finden kann)

- großer Fokus auf den unterstützenden Sekretärs- und Schreiberrollen, um den Chirurgen und Copilot von Schreibarbeiten zu entlasten
    - Xyrill stellt sich das heute schwierig vor; Digitalisierung hat zwar viele Schreibarbeiten vereinfacht, aber auf eine Art und Weise, die das Delegieren schwierig macht
    - hier könnte der Kern liegen, warum verschiedene Leute so unterschiedlich starken Nutzen aus LLMs ziehen

- Skalierungsfrage ist hierdurch vereinfacht: ein Projekt mit 2000 Mitarbeitern kann damit in 200 Zehnerteams geteilt werden, sodass nur die 200 Chirurgen miteinander reden müssen
    - aber nicht gelöst: 200 Chirurgen können immer noch ganz schön viel Kakophonie erzeugen
