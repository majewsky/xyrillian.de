ttimeless kocht eines seiner geTeXten Rezepte:

- [Präambel in einer Masterdatei](https://dl.xyrillian.de/noises/stp-088-shownotes/main.tex)
- Rezepte in einzelnen Dateien
    - Hier das [Hummusrezept](https://dl.xyrillian.de/noises/stp-088-shownotes/hummus.tex) wie von Andong auf [seinem Youtube-Kanal](https://www.youtube.com/mynameisandong) vorgestellt: [Can I Recreate the World's Smoothest Hummus?](https://www.youtube.com/watch?v=IRECUDMZteA), dazu gibt es [das passende Pita](https://dl.xyrillian.de/noises/stp-088-shownotes/pita.tex).
- [Gesamtergebnis als PDF](https://dl.xyrillian.de/noises/stp-088-shownotes/main.pdf)

Xyrill kocht ein C++\-Programm mit make(1):

- Grundproblem: im Kompilierungsmodell von C/C++ entsteht ein Programm in mehreren Schritten
    - jede einzelne Quellcode-Datei wird einzeln vom [Compiler](https://de.wikipedia.org/w/index.php?title=Compiler&oldid=258395815) in eine Objektdatei übersetzt
    - am Ende werden alle Objektdateien vom [Linker](https://de.wikipedia.org/w/index.php?title=Linker_(Computerprogramm)&oldid=239659784) zu einer ausführbaren Programmdatei verknüpft
    - [als Beispiel ein kleines, vollständiges C-Programm](https://github.com/majewsky/xmpp-bridge)
- Vorteile von Makefiles
    - Universalsprache zum Dokumentieren von Prozessen
- Nachteile/Beschränkungen
    - [tw. arkane Syntax](https://github.com/sapcc/keppel/blob/3107ae252039710e02ff06a25e0e55a2084f9487/Makefile#L114-L116)
    - [(fast?) turingvollständig](https://stackoverflow.com/q/3480950), vgl. [Ninja](https://en.wikipedia.org/w/index.php?title=Ninja_(build_system)&oldid=1327454812)
    - (unter anderem dadurch) vergleichsweise ineffizient, vgl. [Tup](https://gittup.org/tup/)
    - moderne Programmiersprachen verwenden andere Kompilierkonzepte, meist fokussiert um Module aus mehreren Quelldateien
