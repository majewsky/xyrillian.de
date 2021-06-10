* Rückblick STP001: Binärzahlen sind dargestellt als Folge von Bits
    * z.B. 16-Bit-Zahl entspricht 16 Stromleitungen, an denen die Spannung an oder aus ist
* Ziel: Berechnungen auf solchen Zahlen
    * z.B. Addition zweiter 16-Bit-Zahlen: 2x16 = 32 Bit gehen rein, 16 neue Bit kommen raus
    * Wie kombiniert man Bits sinnfällig?
    * Logikgatter: Bausteine für solche Schaltkreise
        * allgemein ein oder zwei Eingänge und ein Ausgang
        * ansonsten Logikschaltung bzw. Logikbaustein, der aus mehreren Gattern zusammengesetzt ist

* Logikgatter mit einem Eingang und einem Ausgang
    * Gatter unterscheiden sich in der Ausgabe bei Eingabe 0 (2 Optionen) und in der Ausgabe bei Eingabe 1 (2 Optionen)
    * insgesamt 2x2 = 4 mögliche Gatter
        * 0 -> 0, 1 -> 0 -- Ausgang ist konstant Null (langweilig)
        * 0 -> 0, 1 -> 1 -- Ausgang ist gleich Eingang (langweilig)
        * 0 -> 1, 1 -> 0 -- Ausgang ist Negation des Eingangs (NOT-Gatter)
        * 0 -> 1, 1 -> 1 -- Ausgang ist konstant Eins (langweilig)

* Logikgatter mit zwei Eingängen und einem Ausgang
    * vier mögliche Eingaben (00, 01, 10, 11) und je zwei mögliche Ausgaben (0, 1) -> 2^4=16 mögliche Gatter
    * darunter sechs langweilige Gatter
        * Ausgang ist konstant Null
        * Ausgang ist konstant Eins
        * Ausgang ist gleich Eingang 1
        * Ausgang ist gleich Eingang 2
        * Ausgang ist Negation von Eingang 1
        * Ausgang ist Negation von Eingang 2
    * sechs nützliche Gatter [Logikgatter](https://de.wikipedia.org/wiki/Logikgatter)
        * AND: Sind beide Eingänge 1?
        * OR: Ist wenigstens ein Eingang 1?
        * XOR: Ist genau ein Eingang 1? bzw. Sind beide Eingänge unterschiedlich?
        * NAND: Ist wenigstens ein Eingang 0?
        * NOR: Sind beide Eingänge 0?
        * NXOR, XNOR: Sind beide Eingänge gleich?
    * vier verbleibende Gatter sind nicht so intuitiv, können aber bei Bedarf aus den obigen Gattern zusammengesetzt werden

* Beispiel für eine Logikschaltung: Addiernetz
    * Idee analog zum Addieren auf Papier wie in der Grundschule mit Übertrag
        * Beispiel hier: 35 + 47 = 82
    * [Halbaddierer](https://de.wikipedia.org/wiki/Halbaddierer): zwei Eingänge, zwei Ausgänge
        * Zwei Ein-Bit-Zahlen werden zu einer Zwei-Bit-Zahl addiert.
        * Ausgang 1 = XOR (erste Stelle der addierten Zahl)
        * Ausgang 2 = AND (zweite Stelle der addierten Zahl)
        * z.B. bei dezimal: 5 + 7 = 2 + Übertrag 1
        * in binär: Ausgang 1 = XOR, Ausgang 2 = AND
    * [Volladdierer](https://de.wikipedia.org/wiki/Volladdierer): drei Eingänge, zwei Ausgänge
        * Wie der Halbaddierer, aber nimmt auch den Übertrag der vorherigen Stelle mit.
        * z.B. bei dezimal: 3 + 4 + Übertrag 1 = 8 + Übertrag 0
    * [Addiernetz](https://de.wikipedia.org/wiki/Addierwerk) für N-Bit-Zahlen: 2N Eingänge, N Ausgänge
        * ein Halbaddierer für die erste Stelle
        * je ein Volladdierer für jede weitere Stelle
        * Überlauf: Was ist, wenn der finale Übertrag nicht 0 ist?
            * Option 1: als Zustand im Prozessor speichern, der später abgefragt werden kann
            * Option 2: als Fehler behandeln, Prozess oder Prozessor anhalten

* Warum sind Logikgatter interessant?
    * starke Abstraktionsebene: fast das gesamte Verhalten des Computers als Rechenmaschine ist aus Logikgattern zusammengesetzt
    * konkrete Hardware (z.B. konkrete Art von Transistoren) ist oberhalb der Ebene der Logikgatter fast unerheblich
    * aber natürlich: keine Abstraktion ist perfekt...
    * Beispiel: Schaltverzögerung -> Takt
