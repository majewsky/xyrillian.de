- im Intro erwähnt:
    - [Eternal September](https://de.wikipedia.org/wiki/Eternal_September)
    - [Im Dunkeln ist alles viel aufregender](https://media.ccc.de/v/subscribe11-58195-im-dunkeln-ist-alles-viel-aufregender)

**Talk Nr. 1:** [Geek Code](https://en.wikipedia.org/wiki/Geek_Code)

- Primärquelle (Version 3.12, Stand vom 5. März 1996): <https://web.archive.org/web/20090220181018/http://geekcode.com/geek.html>
- vordergründig: eine kompakte Notation zur Selbstbeschreibung von Geeks durch eine Menge von Attributen und Interessen
    - Beispiel: das Attribut `d` beschreibt den eigenen Dresscode (`d++` für den Anzugträger, `d-` für Jeans und T-Shirt, `!d` für nackt etc.)
    - diverse Modifier, z.B. `$` für "damit verdiene ich Geld" oder `>` für ein persönliches Entwicklungsziel
    - Beispiel: `C++` alleine ist "ich kenne mich gut mit Computern aus", `C++>$` ist "...und möchte damit Geld verdienen"
    - Xyrill fühlt sich an die Grammatik [agglutinierender Sprachen](https://de.wikipedia.org/w/index.php?title=Agglutinierende_Sprache&oldid=232734095) erinnert
- Parallele zum Jargon File (siehe STP050): Geek Code ist ein Produkt seiner Zeit und Subkultur mit veralteten Debatten (z.B. interessiert sich heute noch jemand für Babylon 5 oder Akte X?) tendenziösen Einschätzungen (z.B. vgl. `h---` mit `r+++`)
- Kann es sowas heute nochmal geben?
    - Wahrscheinlich schwierig, man bräuchte zu viele Subkategorien.
    - Passt wahrscheinlich nicht für die heutigen visuell dominierten Medien (YouTube, TikTok, Instagram etc.).
    - Vielleicht für Dating?

**Talk Nr. 2:** Ist was dran an "try turning it off and on again"? Ja! Hierzu eine Kurzeinführung in Kombinatorik:

- ein Speicher mit einem Bit kann zwei Zustände annehmen (klar)
- ein Speicher mit zwei Bit kann `2 * 2 = 4` Zustände annehmen
    - Beispiel/Analogie: UI mit zwei Checkboxen
- aber nur, wenn die Bits voneinander unabhängig sind
    - Beispiel: UI für Ticketkauf, Checkbox 1 = "ist Kind", Checkbox 2 = "ist Rentner"
    - von vier Zuständen ist einer ungültig: Kinder können nicht Rentner sein
- kombinatorische Explosion: bei mehr Bits exponentiell mehr mögliche Zustände, die meisten davon ungültig
    - Programmfehler oder "kosmische Strahlung" können dazu führen, dass ungültige Zustände erreicht werden
    - "Aus- und Wiederanschalten" ist die ultimative Form des Zurücksetzens in einen gültigen Zustand

**Talk Nr. 3:** Umgekehrte polnische Notation

- Wie schreibt man mathematische Ausdrücke?
    - **Infix-Notation:** die klassische Schreibweise für mathematische Ausdrücke, bei der Operatoren zwischen ihren Operanden stehen
    - Problem: Operationsreihenfolge muss entweder per Konvention ("Punkt vor Strich") oder durch Klammern klargestellt werden (Bsp. `2 * 3 + 5 = (2 * 3) + 5 ≠ 2 * (3 + 5)`)
    - [Polnische Notation](https://de.wikipedia.org/w/index.php?title=Polnische_Notation&oldid=249108401): Operatoren stehen vor ihren Operanden (Bsp. `+ * 2 3 5` bzw. `* 2 + 3 5`)
    - hier technisch gesehen keine Klammern notwendig (außer um Menschen das Lesen zu erleichtern, Bsp. `+ (* 2 3) 5` bzw. `* 2 (+ 3 5)`)
    - [Umgekehrte polnische Notation (RPN)](https://de.wikipedia.org/w/index.php?title=Umgekehrte_polnische_Notation&oldid=249951216): Operatoren stehen hinter ihren Operanden (Bsp. `2 3 * 5 +` bzw. `2 3 5 + *`)
- Warum würde man das machen wollen?
    - RPN entspricht dem Betriebsmodus eines Kellerautomaten (siehe STP021)
    - Kellerautomaten kann man extrem günstig in Hardware implementieren (d.h. mit sehr wenig Code, mit sehr wenig Prozessorregistern, mit wenig Speicher)
    - Praxisbeispiel: [Forth](https://de.wikipedia.org/w/index.php?title=Forth_(Programmiersprache)&oldid=249759593)
    - Praxisbeispiel: [HP-12C](https://de.wikipedia.org/w/index.php?title=HP-12C&oldid=245054308)
