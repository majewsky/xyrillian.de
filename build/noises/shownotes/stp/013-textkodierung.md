Laut Wiktionary:

> \[**Text** ist eine\] mündliche oder schriftliche Folge von Sätzen, die miteinander syntaktisch und semantisch verbunden sind (Kohärenz, Kohäsion), eine abgeschlossene Einheit bilden (Kompletion) und eine bestimmte kommunikative Funktion (Textfunktion) erfüllen.
> 
> \[Von\] lateinisch _textus_ für "Inhalt" oder "Gewebe der Rede", "Text".

- zur Definition
    - Wiktionary geht vom linguistischen Blickwinkel aus
    - in der Informatik: Zeichenketten (oder synonym "Zeichenfolgen"), in Programmiersprachen meist als [String](https://de.wikipedia.org/wiki/Zeichenkette ) bezeichnet

- Problem: Computer können nur Zahlen! Wie bilden wir Zeichenketten als Zahlenfolgen ab?
    - Idee 1: jedes Zeichen als eine Zahl einer bestimmten Größe, z.B. 1 Byte oder 2 Byte (_fixed-length code_)
        - primitives Beispiel: alle Buchstaben des Alphabets durch ihre Position darstellen, z.B. "hallo" -> "08 01 12 12 15"
    - Idee 2: häufige Zeichen mit kurzen Bytefolgen darstellen, seltene Zeichen mit längeren Bytefolgen (_variable-length code_, siehe [Entropiekodierung](https://de.wikipedia.org/wiki/Entropiekodierung))

- Beispiel: [Morse-Code](https://de.wikipedia.org/wiki/Morsezeichen)
    - variable Länge: z.B. "e" ist am häufigsten und deswegen nur 1 Ton kurz, "q" ist sehr selten und deswegen 4 Töne lang
    - Kodiertabelle enthält nicht nur Buchstaben und Ziffern, sondern auch Sonderzeichen (Punkt, Komma, etc.) und Spezialzeichen wie "Spruchende" oder "Fehler, Wiederholung ab letztem vollständigem Wort"

- in der frühen Informatik (bis in die 1990er Jahre): **Codepages** (Zeichentabellen)
    - Codes mit fester Länge von meist 1 Byte (7 oder 8 Bit) pro Zeichen
    - Beispiel mit 7 Bit: [ASCII](https://de.wikipedia.org/wiki/ASCII ), unter Linux siehe `man ascii` ([RFC 20](https://datatracker.ietf.org/doc/html/rfc20 ))
    - Beispiel mit 8 Bit: [Codepage 437](https://de.wikipedia.org/wiki/Codepage_437 )
    - Standardisierung als [ISO/IEC 8859](https://en.wikipedia.org/wiki/ISO/IEC_8859 )
    - Problem: 8 Bit (256 Zeichen) sind zu wenig für alle in europäischen Sprachen verwendeten Zeichen, insb. wegen [diakritischer Zeichen](https://de.wikipedia.org/wiki/Diakritisches_Zeichen ) -> verschiedene Codepages für verschiedene Sprachen -> [Mojibake](https://de.wikipedia.org/wiki/Zeichensalat )
    - spätestens mit zunehmend globaler Kommunikation im Internetzeitalter musste eine bessere Lösung her

- [Unicode](https://de.wikipedia.org/wiki/Unicode )
    - keine Textkodierung im engeren Sinne, sondern erstmal nur eine durchnummerierte Liste aller möglichen Zeichen (die Nummer zu einem Zeichen heißt **Codepunkt**)
    - Anspruch: wenn ein Zeichen in tatsächlichen Textdokumenten verwendet wird oder wurde, gehört es in Unicode (von Buchstaben über Redigierungssymbole mittelalterlicher Typografen bis zu Hieroglyphen und Keilschrift)
    - Größe: 17 Ebenen (_Planes_) zu je 65536 Codepunkten = 1.114.112 Codepunkte; davon sind 2048 Codepunkte aus technischen Gründen (für die Kodierung von Surrogatpaaren) unbelegbar, somit maximal möglich 1.112.064 Codepunkte

- Ebenen in Unicode
    - Ebene 0: **Basic Multilingual Plane** mit allen zeitgenössischen und weitverbreiteten Schriftsystemen ([Übersichtskarte](https://commons.wikimedia.org/wiki/File:Roadmap_to_Unicode_BMP_multilingual.svg ))
    - Ebene 1: **Supplementary Multilingual Plane** mit historischen Schriftsystemen und Obskuritäten wie Domino/Mahjongg-Steinen und 😍 Emoji 😍
    - Ebene 2: **Supplementary Ideographic Plane** mit selten benutzten CJK-Schriftzeichen
    - Ebene 3-13: noch frei
    - Ebene 14: **Supplementary Special-purpose Plane** mit Steuerzeichen zur Selektion alternativer Schriftzeichenformen
    - Ebene 15/16: für private Verwendung reserviert

- Kodierungen für Unicode mit fixer Länge
    - UCS-2: eine 16-Bit-Zahl pro Zeichen -> kann nur die BMP kodieren (UCS-2 stammt aus der Zeit, bevor Unicode mehrere Planes hatte)
    - UCS-4/UTF-32: eine 32-Bit-Zahl pro Zeichen -> extrem verschwenderisch

- Kodierungen für Unicode mit variabler Länge
    - UTF-16: BMP-Zeichen als eine 16-Bit-Zahl, andere Zeichen als zwei 16-Bit-Zahlen -> Probleme mit Endianness
    - [UTF-8](https://de.wikipedia.org/wiki/UTF-8 ): jedes ASCII-Zeichen als ein Byte, alle anderen Zeichen als eine Folge von 2-5 Bytes (unter Linux siehe `man utf-8`)
    - UTF-8 ist spektakulär: sehr kompakt, selbstsynchronisierend, einfach zu implementieren (auch in Hardware), sortierungserhaltend, ASCII-abwärtskompatibel, Nicht-ASCII-Zeichen enthalten niemals ASCII-Bytes; und wurde buchstäblich in einem Abend von Ken Thompson und Rob Pike auf einer Serviette entworfen

- Probleme von Unicode
    - Anspruch: jeder bestehende Kodierungsstandard soll Zeichen für Zeichen in Unicode abbildbar sein -> dadurch unnötig viele Codepunkte insb. für diakritische Zeichen und Probleme mit Normalisierung
    - allgemeiner: Was ist eigentlich ein Zeichen im Sinne von "ein Codepunkt"? (mehr dazu in der nächsten Folge) -> Unicode übernimmt hier oftmals teils kontroverse Klassifikationen aus nationalen Standards z.B. für indische und thailändische Schriften
    - am signifikantesten hier: [Han-Vereinheitlichung](https://de.wikipedia.org/wiki/Han-Vereinheitlichung ) ([Beispiel](https://commons.wikimedia.org/wiki/File:Source_Han_Sans_Version_Difference.svg ))

Für die historische Herleitung von Schriftkodierung empfehlen wir wieder den großartigen Podcast [Request for Comments](https://requestforcomments.de/ ). Diesmal [die Folge über RFC 20 ASCII](https://requestforcomments.de/archives/400 ).

Das nächste Mal geht es dann um T̪͙̪̻̹̩̠͔͑͋ͥ͘ę̥͔̝̘ͦͬx̨̜̩̜͙̦̻̗̳̒̊͆t̯̳̓ͥ͐ͣ́d͔̱̗̯̪̭́͛͜ͅa̵͕̯͕̻̜ͮ͒ͭȓ̜͚͚̯͎͓̋̎̈́͘s̥̯̠̬̆̿͡t̫̱̺͉̱̓͒͑͐͢e̸͖̮̙̺̠̗̬̅ͅl̝͉͚ͪ͋ͤ͜l̵̦̪͔͈̉͑u̴̝̲̠̳̖̘͚̔n͙̹̍̆͊̈͢g̨͍̣̝̣͇̯̮̠̏.
