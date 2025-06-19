Fortsetzung aus STP068. Wir lesen weiterhin die englische Erstausgabe von 1975, die [im Internet Archive](https://archive.org/details/MythicalManMonth) digitalisiert wurde.

### Kapitel 4: Aristokratie, Demokratie und System-Design

- Brooks beschreibt als Leitmotiv die Kathedrale von Reims, die von acht Generationen von Baumeistern errichtet wurde, aber trotzdem eine beachtliche gestalterische Kohärenz aufweist, da die nachfolgenden Generationen das ursprüngliche Design ungewöhnlich stark respektieren: "Das Ergebnis verkündet nicht nur die Herrlichkeit Gottes, sondern auch seine Macht, gefallene Menschen von ihrem Stolz zu befreien."
    - Kontrast nicht nur zu den meisten anderen Kirchenbauten, sondern auch zu Programmsystemen
    - dort aber nicht wegen Mehrgenerationalität, sondern wegen Vielstimmigkeit
    - These: konzeptionelle Integrität ist die wichtigste Eigenschaft im System-Design (vgl. STP006: Unix-Philosophie)
    - Vorteil: bessere Verwendbarkeit und Erlernbarkeit
    - Umsetzung: das Design sollte von einigen wenigen Architekten kontrolliert werden, die im Interesse des Benutzers handeln; doch hieraus drei Problemstellungen

- Problem 1: Folgt hieraus nicht eine aristokratische Klassengesellschaft von wenigen Vordenkern und bauendem Fußvolk?
    - nein, auch auf der Implementationsebene gibt es viele interessante Entscheidungen zu treffen
        - Brooks zitiert einen Aphorismus aus der Kunst: "Form befreit.", sprich: Kreativität erwächst aus dem Umgang mit Beschränkungen
    - in der Architektur geht es darum, was man dem Nutzer präsentiert (das "Was")
    - in der Bauphase geht es darum, wie man dieses Verhalten erreicht (das "Wie")
    - es ist absolut gewollt, dass die Erbauer auch Designideen vorschlagen und auf Beschränkungen hinweisen; nur entscheiden sie nicht final über das Design

- Problem 2: Wie verhindert man, dass die Architekten von den Zwängen der Realität abgekoppelt im Elfenbeinturm davonschweben?
    - Xyrill bringt eines seiner Lieblings-Schimpfwörter ins Spiel: ["Architektur-Astronaut"](https://en.wikipedia.org/w/index.php?title=Architecture_astronaut&oldid=1221960855)
    - Lösung dieselbe wie für das dritte Problem

- Problem 3: Wie stellt man sicher, dass die Erbauer mit allen wichtigen Details des Designs vertraut gemacht werden?
    - Lösung: die Erbauer frühzeitig in den Design-Prozess involvieren, Ideen aufnehmen, Designvorschläge auf praktische Umsetzbarkeit abklopfen etc.
    - Vergleich mit der Errichtung eines Gebäudes: man kann die Erbauer nicht erst einstellen, wenn der Plan fertig ist, weil in der Software-Entwicklung immer noch vieles explorativ und nicht nach Standard-Muster abläuft

### Kapitel 5: Das Problem des zweiten Systems ("Second-System Effect")

- in diesem Dialog zwischen Architekt und Erbauer stellen sich viele Ideen des Architekten als unpraktikabel heraus und werden verworfen... bis zum nächsten Mal
    - Szenario: das erste System war erfolgreich, der Architekt hat Prestige und Erfahrung gewonnen
    - "Dieses zweite ist das gefährlichste System, dass man jemals entwirft." -> [Second-System Effect](https://en.wikipedia.org/w/index.php?title=Second-system_effect&oldid=1247130968)
    - beim dritten und vierten System hat man dann mehr Erfahrung, was tatsächlich funktioniert und was nur auf dem Papier gut aussieht
    - Xyrill legt Wert darauf, dass seine jüngeren Kollegen früh ihre eigenen Komponenten alleine entwerfen, damit sie dann auch in die Probleme ihrer eigenen Design-Entscheidungen reinlaufen und daraus lernen

- der vorgeschlagene Lösungsansatz ist leider etwas antiquiert: "Eine Disziplin, die dem Architekten die Augen öffnen wird, ist es, jeder kleinen Funktion einen Wert zuzuweisen: Fähigkeit X ist nicht mehr wert als M Bytes Speicher und N Mikrosekunden pro Aufruf. Diese Werte dienen als Richtschnur für die ersten Entscheidungen und während der Implementierung als Leitfaden und Warnung an alle."
    - nicht realisitisch in Zeiten, wo die meisten Programmierer offenkundigerweise noch nie einen Profiler verwendet haben oder auch nur ein Gefühl für ihre eigene Ressourcenvergeudung haben

### Kapitel 6: Weitergabe von Informationen ("Passing the Word")

- Wie stellt eine Gruppe von 10 Architekten sicher, dass ihre Entwürfe von 1000 Mitarbeitenden berücksichtigt und von allen Benutzern verstanden werden?
    - gute Dokumentation ist erforderlich: Handbücher für die Benutzer, Architekturpläne für die Erbauer

- Qualitäten guter Dokumentation
    - vorzugsweise aus der Feder von nur ein oder zwei Personen -> Fokus auf konzeptionelle Integrität
    - Beschreibung sowohl des definierten Verhaltens als auch der Grenzen der Definiertheit
    - Definitionen in formaler Notation und/oder präziser Sprache, um Lücken und Widersprüche im Design möglichst früh aufzudecken
    - [Anti-Pattern](https://de.wikipedia.org/w/index.php?title=Anti-Pattern&oldid=248340518): Definition eines Systems durch sein eigenes Verhalten
        - führt zu Bug-für-Bug-Kompatibilität (siehe STP050) bzw. Ossifikation (siehe STP056): inhaltliche Widersprüche zwischen der ursprünglichen Implementation, späteren Revisionen sowie später abgeleiteten Spezifikationen
    - gutes Pattern: Spezifikation maschinenlesbar darstellen, damit die Erbauer die Konformität ihrer Implementation mit der Spezifikation automatisch prüfen können
        - aktuelles Beispiel: formale Spezifikation von Schnittstellen mittels [OpenAPI](https://de.wikipedia.org/w/index.php?title=OpenAPI&oldid=250760605)

- Wie erarbeitet man Designs? Vorschlag von Brooks:
    - Meetings zwischen den Architekten und involvierten Erbauern, in denen das Problem vorgestellt und mehrere Designideen gesammelt werden
    - dann Erarbeitung von detaillierten Designentwürfen mit Pro-/Kontra-Analyse
    - daraus entweder sofortiger Konsens oder Entscheidung eines Chefarchitekten
    - klingt offensichtlich, ist aber im [agilen Zeitalter](https://de.wikipedia.org/w/index.php?title=Agile_Softwareentwicklung&oldid=251327114) selten anzutreffen

- noch eine kleine Anmerkung am Rande: auch Kurzabsprachen am Telefon sollten dokumentiert werden :)

### Kapitel 7: Warum ist der Turm zu Babel gefallen?

- Woran können Projekte scheitern?
    - unklare Zielstellung
    - nicht genug Personal
    - nicht genug Material/Budget
    - nicht genug Zeit
    - unzureichende Technologien
    - *ineffektive Kommunikation*

- Wie sollen Teams kommunizieren? Neben direkten Gesprächen z.B. via Telefon und Gruppengesprächen z.B. in Meetings wird ein **Arbeitsbuch** vorgeschlagen.
    - nicht ein konkretes Dokument, sondern eine Struktur für alle Dokumente im Projekt
    - die Dokumentation am Ende geht aus den Memos am Anfang hervor, also gleich richtig machen
    - Memos sollen durchnummeriert werden, damit man später darauf Bezug nehmen kann (wir fühlen uns an RFCs erinnert) -> Arbeitsbuch ergibt sich als Baumstruktur aller Memos (wir fühlen uns an Hypertext erinnert, und in der Tat wird Douglas C. Engelbart am Rande erwähnt)
    - Brooks diskutiert dann praktische Unterschiede zwischen Papier und Microfiche, die heutzutage keine Relevanz mehr haben (aber ein Nachteil: die Verteilung aktualisierter physischer Seiten sorgte dafür, dass die Erbauer mitbekamen, wo sich wann was ändert)

> D. L. Parnas von der Carnegie-Mellon-Universität hat eine noch radikalere Lösung vorgeschlagen. Seine These lautet, dass ein Programmierer am effektivsten ist, wenn er von den Details der Konstruktion von Systemteilen, die nicht seine eigenen sind, abgeschirmt wird, anstatt ihnen ausgesetzt zu sein. Dies setzt voraus, dass alle Schnittstellen vollständig und genau definiert sind. Dies ist zwar ein solider Entwurf, aber in der Abhängigkeit von seiner perfekten Ausführung ist eine Katastrophe vorprogrammiert. Ein gutes Informationssystem deckt Schnittstellenfehler auf und regt zu deren Korrektur an.

- Xyrill stimmt voll und ganz zu: sorgfältiges Gegenlesen der Meeting-Minutes der benachbarten Teams ist zwar extrem zeitaufwändig, verhindert aber regelmäßig katastrophale Design-Entscheidungen

- Organisationsstruktur: Brooks hält eine baumförmige Organisationsstruktur für unausweichlich (mehrere Vorgesetzte für einen Angestellten = Chaos) und schaut darauf, was jedes Team haben sollte
    - Produzent ("Producer"): stellt das Team zusammen, verteilt Arbeit, definiert Ziele und Zeitpläne, kommuniziert **außerhalb** des Teams
    - Architekt bzw. technischer Leiter ("Technical Director"): entwirft das Design und entwickelt es fortlaufend weiter, kommuniziert **innerhalb** des Teams
    - drei Optionen, wie diese Rollen verteilt sein können

- Option 1: Produzent und Architekt sind dieselbe Person
    - nur für kleine Teams gangbar, bis etwa 3-6 Personen
    - "Denker sind rar; Macher sind rarer; und Denker-Macher sind am rarsten."

- Option 2: Produzent als Chef, Architekt als seine rechte Hand
    - funktioniert nur, wenn der Produzent die technische Autorität des Architekten akzeptiert
    - Xyrill hat lange in einem derartigen Zweierteam als Architekt gearbeitet und weiß, dass es funktionieren kann; kommt aber extrem auf die interpersonelle Dynamik des Duos an

- Option 3: Architekt als Chef, Produzent als seine rechte Hand
    - findet Brooks am besten; Xyrill ist geneigt, zuzustimmen

### Noch mehr

Es gibt noch acht weitere Kapitel, die wir irgendwann™ mal besprechen werden.
