![Bild](https://dl.xyrillian.de/noises/stp-038-intro.png)

- Warum testen? -> siehe STP037

- Anforderungen an Tests
    - reproduzierbar: nicht abhängig von unerwarteten äußeren Einflüssen oder menschlicher Intervention
    - fokussiert: im Fehlerfall soll möglichst klar sein, wo das Problem liegt
    - schnell: damit man beim Ändern des Systems schnelles Feedback bekommt (siehe auch [XKCD 303](https://xkcd.com/303/))
    - realitätsnah: überprüft eine Funktion, die der Benutzer so tatsächlich benötigt

- manuelle Tests
    - kostenintensiv: [Online-Shooter braucht(e) 200 Tester](https://www.destructoid.com/apex-legends-qa-baton-rouge-layoffs-electronic-arts-ea/) (das Spiel erlaubt Matches mit bis zu 60 Teilnehmern)
    - erfüllt von den obigen Anforderungen nur die Realitätsnähe
    - wenn's geht, lieber automatisieren
    - gängige Strategie: manuelle Tests zur Fehlersuche, im Zweifelsfall durch die Kunden ;) &ndash; dann Nachstellung des gefundenen Problems

- automatisierte Tests am Beispiel der [Testpyramide](https://commons.wikimedia.org/w/index.php?title=File:Testing_Pyramid.png&oldid=731870968#/media/File:Testing_Pyramid.png)
    - **Unit Tests**: Überprüfung einer einzelnen Komponente in Isolation
        - besonders wertvoll, wenn die Komponenten auf besonders reproduzierbare Weise arbeiten (zum Beispiel pure Berechnungen; Vorgänge, die *nicht* Netzwerkzugriff/Festplattenzugriff/etc. erfolgen)
        - Problem: es reicht nicht, wenn jede Komponente für sich genommen funktioniert, aber sie nicht zusammen passen
        - Bewertung: reproduzierbar, fokussiert, schnell, aber nicht realitätsnah
    - **Integration Tests**: Überprüfung des Zusammenspiels mehrerer Komponenten eines Gesamtsystems
        - nach Möglichkeit Eingrenzung äußerer Faktoren, z.B. statt Kommunikation mehrerer Komponenten über ein reales Netzwerk Simulation einer Netzwerkstruktur im Programm oder mittels Containern/VMs (siehe STP023)
        - Bewertung: tendenziell weniger reproduzierbar, etwas langsamer, weniger fokussiert, aber realitätsnäher
    - **End-to-End-Tests** (E2E): Überprüfung des Verhaltens eines Gesamtsystems
        - analog zu den manuellen Tests, aber nach festem Algorithmus
        - höchster Aufwand in der Umsetzung, siehe insbesondere UI-Tests mittels [Headless Browser](https://en.wikipedia.org/w/index.php?title=Headless_browser&oldid=1106195952) etc.
        - Bewertung: am wenigsten reproduzierbar (aufgrund der Vielzahl an Variablen), am unfokussiertesten, am langsamsten, aber am nächsten an der Realität
    - allgemeine Empfehlung: ganz viele kleine Unit Tests und einige mittelgroße Integrationstests für die schnelle Feedbackschleife, dazu einige wenige große E2E-Tests für den Abgleich mit der Realität

- alternative Teilung: [Blackbox-Tests vs. Whitebox-Tests](https://de.wikipedia.org/w/index.php?title=Black-Box-Test&oldid=226194172)
    - Blackbox: Umsetzung eines Tests ohne innere Kenntnis der Struktur des Systems, z.B. anhand einer Spezifikation
    - Vorteil: ein Blackbox-Test durch ein anderes Team kann Missverständnisse des ursprünglichen Teams aufdecken
    - Nachteil: kann nur E2E sein

- die reine Lehre: **Test Driven Development** ([Testgetriebene Entwicklung](https://de.wikipedia.org/w/index.php?title=Testgetriebene_Entwicklung&oldid=230163162))
    - neue Funktionalität darf man nur bauen, wenn es Testabdeckung dafür gibt
    - erst fehlschlagenden Test schreiben; dann Funktionalität einbauen, sodass der Test erfolgreich durchläuft; dann aufräumen
    - Xyrill macht lieber "Data Driven Development" und hat Bedenken

- Kann man die Qualität eines Tests bewerten? -> Idee: [Code Coverage](https://en.wikipedia.org/w/index.php?title=Code_coverage&oldid=1143329859)
    - Welcher Anteil des Programmcodes wird von meinem Test tatsächlich abgedeckt?
    - Was heißt eigentlich "Anteil"?
        - Beispiel: `if (a or b or c) { return "Kredit abgelehnt" }` (einfach bei Zeilenabdeckung, schwierig bei Bedingungsabdeckung)
        - Beispiel: "Kommt ein Software-Tester in eine Bar."
    - Problem: [Goodharts Gesetz](https://de.wikipedia.org/w/index.php?title=Goodharts_Gesetz&oldid=229408892), "Sobald eine Metrik zu einem Ziel wird, hört sie auf, eine gute Metrik zu sein."
    - Problem: für manche Zeilen sind Tests wichtiger als andere (wichtig für vertrackte Logik, weniger wichtig für Sofortabbruch bei unerwarteten Fehlern)

- Tests muss man trotzdem noch schreiben. Kann man das wegautomatisieren?
    - [Statische Analyse](https://de.wikipedia.org/w/index.php?title=Statische_Code-Analyse&oldid=221804280): automatische Suche nach bedenklichen Mustern (z.B. Datei wurde zum Schreiben geöffnet, wird aber am Ende nicht geschlossen -> Risiko des Datenverlusts; z.B. Verwendung veralteter kryptografischer Primitiven)
    - [Fuzzing](https://de.wikipedia.org/w/index.php?title=Fuzzing&oldid=225194301): anstatt eines Tests mit festen Werten ("bei Eingabe X wird Ausgabe Y erwartet") nur Vorgabe einer Struktur ("wenn hier ein Text reingesteckt wird, kommt da eine Zahl raus"), dann automatische Suche nach Absturzbedingungen

- Machen Tests die Sache immer besser?
    - [Heisenbug](https://de.wikipedia.org/w/index.php?title=Heisenbug&oldid=227133176): "ein Softwarefehler, der zu verschwinden oder sein Verhalten zu verändern scheint, wenn man versucht, ihn zu untersuchen"

- im Gespräch erwähnt:
    - [ISO9001](https://de.wikipedia.org/wiki/ISO_9001)
    - [FG034 zu Werner Heisenberg](https://forschergeist.de/podcast/fg034-werner-heisenberg/)
