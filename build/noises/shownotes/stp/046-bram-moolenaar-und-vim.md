Aus aktuellem Anlass: Anfang August 2023 ist Bram Moolenaar gestorben.

- [Bram Moolenaar](https://de.wikipedia.org/w/index.php?title=Bram_Moolenaar&oldid=236301189)
    - Jahrgang 1961, bis 1985 Studium der Elektrotechnik in Delft mit Abschlussarbeit zu Multiprozessor-Unix-Architektur
    - seit 1988 Entwicklung eines Klons des Texteditors vi mit dem Namen [vim](https://de.wikipedia.org/w/index.php?title=Vim&oldid=236190917) (erst "vi imitation", später "vi improved"; "vim" kann aber auch "Elan" bedeuten)
    - bis kurz vor seinem Tod (d.h. 35 Jahre) Leitung des vim-Projektes in seiner Freizeit (neben der Arbeit u.a. bei Google)
    - außerdem Engagement für die NGO [ICCF Holland](https://iccf.nl/index.html), die Waisenkinder in Uganda unterstützt

- in diesem Zusammenhang: Was ist eigentlich vim? Und warum verbringt Xyrill sein halbes Leben darin? (Spoiler: man kann alle intellektuelle Arbeit als Texteditieren darstellen; siehe auch STP006)
    - frühe 1970er-Jahre: [`ed`](https://de.wikipedia.org/w/index.php?title=Ed_(Texteditor)&oldid=207200962) &ndash; der erste Unix-Texteditor; damals hatten Unix-Computer noch keinen Bildschirm, sondern nur Tastatur und Drucker
    - 1976: [`ex`](https://de.wikipedia.org/w/index.php?title=Ex_(Texteditor)&oldid=217207358) &ndash; eine von vielen Weiterentwicklungen von `ed` mit zusätzlichen Funktionen
    - darin enthalten: ein neuartiger "visual mode", der 1979 der Standard wurde &ndash; [`vi`](https://de.wikipedia.org/w/index.php?title=Vi&oldid=234166542) ist geboren
    - 1989 dann `vim` als Weiterentwicklung von Brams Portierung von `vi` auf Atari-Computer
    - sprich: im Großen und Ganzen eine 50 Jahre lang ungebrochene Evolutionslinie von Texteditoren, analog und parallel zu Unix

- Kernidee von vi-artigen Texteditoren: **modales Editieren**
    - in herkömmlichen Texteditoren: jeder Tastendruck hat immer die gleiche Bedeutung (z.B. "u" fügt ein kleines U ein; z.B. "Strg+Z" macht die letzte Operation rückgängig)
    - Nachteil: komplexe Operationen sind begrenzt darin, wieviele Tastenkombinationen man sich merken kann
    - in vi-artigen Editoren: Bedeutung eines Tastendrucks hängt vom aktuellen Modus ab (z.B. "u" im Eingabemodus fügt ein kleines U ein, jedoch "u" im Befehlsmodus macht die letzte Operation rückgängig)
    - erster Vorteil: mehr Befehle sind durch einfache Tastendrücke erreichbar, die sich einfacher tippen lassen als Akkorde wie "Strg+Umschalt+R"

- zweite Besonderheit: **Aktionsgrammatik**
    - in herkömmlichen Texteditoren: viele Aktionen, die alle ziemlich ähnlich sind (z.B. ein Zeichen nach links löschen, ein Wort nach links löschen, ein Zeichen nach links auswählen, ein Wort nach links auswählen, etc.)
    - Befehlsmodus in vi: Befehle werden wie mit einer Grammatik zusammengesetzt
        - Substantive = Bewegungen (z.B. "w" = ein Wort nach rechts, "b" = ein Wort nach links, "j" = eine Zeile nach unten)
        - Zahlen = Vervielfältigungen (z.b. "3w" = drei Worte nach rechts, "5j" = fünf Zeilen nach unten)
        - Verben = Aktionen (z.B. "d" = löschen, "c" = ändern)
        - ganzer Befehl: nur eine Bewegung (bewegt den Cursor, z.B. "3w") oder eine Aktion und eine Bewegung (z.B. "d5j")
    - damit zweiter Vorteil: eine große Menge an komplexen Aktionen ist mit sehr kompakten Befehlen auszudrücken (z.B. "Ich möchte den Inhalt des Klammerpaars löschen, in dem mein Cursor steht" = "di(", "ich möchte diesen Absatz und die nächsten drei Absätze automatisch formattieren" = "ga4p")

- weitere coole Funktionen in vim
    - Makrofunktion: Folgen von Tastendrücken können aufgezeichnet und dann automatisiert abgespielt werden (gibt's auch in anderen Editoren, aber paart sich mit vi's Befehlssprache besser als mit denen der meisten anderen Texteditoren)
    - man kann damit [Hacker Jeopardy](https://media.ccc.de/v/34c3-9007-hacker_jeopardy#t=1093) gewinnen :3

- Kann man das lernen?
    - Ja, auf jeden Fall. Was man da an Zeit aufwendet, zahlt sich dann über Jahrzehnte aus.
    - Xyrill empfiehlt: Erstmal nur mit den Basics anfangen (Dateien öffnen und schließen, zwischen Modi wechseln). Dann schrittweise Befehle lernen, wann immer man das Gefühl hat, dass man was von Hand macht, was auch einfacher gehen sollte.
    - vim hat ein eingebautes Tutorial: `vimtutor` auf der Kommandozeile ausführen
    - vim-Cheatsheet auf dem Schreibtisch bereitlegen ([Beispiel (deutsch)](https://cheatography.com/mipapo/cheat-sheets/vim-cheatsheet/), [Beispiel (englisch)](https://i.redd.it/ve1jv3m3qqj21.png))

- im Gespräch erwähnt
    - [Perl-Programm mit Spracherkennung schreiben](https://www.youtube.com/watch?v=KyLqUf4cdwc)
    - [Pentaradio zu Git (Versionsverwaltungssystem)](https://c3d2.de/news/pentaradio24-20210727.html)
