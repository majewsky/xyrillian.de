Wir haben ein Errata zu STP019. Alex schreibt:

> DMA kenne ich anders, nicht das der Prozess direkt über die MMU auf ein Peripheriegerät durchgreift, sondern dass ein Peripheriegerät ohne Beteiligung der CPU Daten vom/zum Hauptspeicher übertragen kann.
> Man muss dem Kernel allerdings beim Reservieren von Speicher für DMA
sagen, dass das ein DMA Buffer sein soll, weil DMA selbst nicht über die MMU läuft sondern vom Peripheriegerät direkt von/zur physikalischen Adresse.
> Den eigentlichen Transfer zwischen Buffer und Peripherie muss
man dann aber per Software entsprechend aktivieren.

Xyrill antwortet: [Ja, stimmt.](https://de.wikipedia.org/w/index.php?title=Direct_Memory_Access&oldid=232059711) Nun zur eigentlichen Folge:

* Intro-Intro: in Referenz auf den klassischen Dilbert-Cartoon, den man [zum Beispiel hier](https://www.americanscientist.org/article/the-quest-for-randomness) sieht

* Warum sind hochwertige Zufallszahlen wichtig?
    * siehe STP043: Erzeugung eines RSA-Schlüsselpaars benötigt zwei große und **geheime** Primzahlen
    * geheim = niemand anders kann es raten = zufällig
    * oder zumindest von Zufall nicht zu unterscheiden (**Pseudozufall**)

* Was ist "guter" Zufall?
    * erstaunlich schwer zu definieren
    * vielleicht Periodenlänge (Dauer, bis sich die Zahlenfolge wiederholt)
    * vielleicht Gleichverteilung der Auftrittshäufigkeit von Teilfolgen (jede einstellige/zweistellige/dreistellige/etc. Zahl sollte gleich oft in der Folge vorkommen)
    * Problem: die Folge "2236067977499789805051478..." ist nach all diesen Maßstäben guter Zufall &ndash; bis man bemerkt, dass das die Dezimalstellen der Quadratwurzel von 5 sind
    * guter (aber unpraktikabler) Test: hohe [Kolmogorow-Komplexität](https://de.wikipedia.org/w/index.php?title=Kolmogorow-Komplexit%C3%A4t&oldid=207581838) (siehe STP025)
    * Errata: Beim Nachhören ist mir aufgefallen, dass die Kolmogorow-Komplexität eines Pseudozufallszahlengenerators nie besonders groß sein kann, da der Zufallszahlengenerator ja gerade ein Programm ist, was die entsprechende Zahlenfolge erzeugen kann. Das ist kein praktisches Problem, da das Programm im Kolmogorow-Sinne auch die Startwerte des Zufallszahlengenerators umfassen müsste, die aber gerade geheim sind.

* "echter" Zufall kommt aus physikalischen Zufallszahlengeneratoren (Hardware-RNG)
    * z.B. thermisches Rauschen eines elektrischen Widerstandes
    * z.B. radioaktive Zerfallsvorgänge
    * z.B. atmosphärisches Rauschen (Empfang auf einem Frequenzband, auf dem keiner sendet)
    * z.B. [Abfilmen eines Aquariums :)](https://www.semanticscholar.org/paper/True-Random-Number-Generator-using-Fish-Tank-Image-Katyal-Mishra/374875fdad184959a7e13eafe9c02bbd44e0aca7)
    * siehe auch: Würfeln, Lottozahlen
    * Probleme: nicht in jedem Computer verfügbar, pro Zeiteinheit begrenzte Menge von Zufall
    * Idee: Erzeugung beliebiger Mengen von Pseudozufall aus einem festen Algorithmus

* Beispiel: [Linearer Kongruenzgenerator](https://de.wikipedia.org/w/index.php?title=Kongruenzgenerator&oldid=214494984#Linearer_Kongruenzgenerator)
    * Parameter: geheime Ganzzahlen `a`, `b`, `m`; Startwert `y`
    * Erzeugung einer Folge von Pseudozufallszahlen durch die Formel `y' = (a * y + b) mod m`
    * Problem: bereits aus wenigen aufeinander folgenden `y`-Werten kann man auf die geheimen Zahlen zurückschließen (Ansatz vergleichbar mit Kryptoanalyse)
    * in besseren Pseudozufallszahlengeneratoren wird nach jedem Schritt nur ein kleiner Teil des Zustands herausgegeben, um Analyse zu erschweren
    * Anwendungsbereich für einfache Pseudozufallszahlengeneratoren: nicht sicherheitsrelevanter Pseudozufall, z.B. Ereignisse in Computerspielen oder Randomisierung von automatisierten Tests (hier kann Reproduzierbarkeit aus einem Startwert für spätere Nachvollziehung hilfreich sein)

* für sicherheitsrelevante Anwendungen: [Kryptografisch sichere Zufallszahlengeneratoren](https://de.wikipedia.org/w/index.php?title=Kryptographisch_sicherer_Zufallszahlengenerator&oldid=234755968)
    * Beispielidee: mit einem Hardware-RNG ein paar Bytes als Startwerte erzeugen, wiederholt mit AES-CTR (siehe STP043) verschlüsseln und die verschlüsselten Bits (oder Teile davon) als Ausgabe nehmen
    * wenn man keinen Hardware-RNG hat, kann man Zufallsbits auch aus Beobachtung des normalen Betriebs extrahieren (z.B. Mikrosekundenbruchteile der Zeitstempel von Ereignissen wie Tastaturdrücken oder dem Empfang von Datenpaketen übers Netzwerk)
    * unter Unix: `/dev/random` vs. `/dev/urandom` ; ist aber unter Linux mittlerweile so gut wie synonym ([siehe auch](https://media.ccc.de/v/32c3-7441-the_plain_simple_reality_of_entropy)); Probleme nur beim Systemstart, wenn noch nicht genug Entropie vorliegt (siehe auch: haveged, VirtIO RNG, systemd-random-seed.service)

* im Gespräch erwähnt:
    * [Video: "How we solved the worst minigame in Zelda's history"](https://www.youtube.com/watch?v=1hs451PfFzQ) (Errata: Xyrill hatte im Gespräch behauptet, dass das Video von Lowest Percent sei. Es ist aber von Linkus7.)
