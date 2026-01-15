Wir setzen unsere Betrachtung des [Cursed Computer Iceberg Meme](https://suricrasia.online/iceberg/) aus STP042 und STP064 fort. Dieses Mal hat Xyrill sich vorgenommen, einen Pick aus jeder Ebene des Eisbergs zu besprechen.

- "Above the Iceberg" - Oberhalb des Eisbergs: [Spectre/Meltdown](https://meltdownattack.com/)
    - Beispiel eines Seitenkanalangriffs (vgl. Besprechung von nicht-zeitkonstanten Operationen in STP041)
    - war damals (2018) auch [im Pentaradio besprochen worden](https://c3d2.de/news/pentaradio24-20180123.html)
    - und im [Chaosradio](https://chaosradio.de/cr242-einmaleinschmelzenbitte)
    - Grundlage: moderne Prozessoren erzielen ihre Geschwindigkeit zu guten Teilen durch spekulative Ausführung zukünftiger Instruktionen (siehe STP011)
    - Idee: Angreifer kann nicht auf geschützten Speicher zugreifen, aber die spekulative Ausführung weiß das nicht
    - spekulativ ausgeführter (und dann verworfener) Befehl ist so gebaut, dass je nach dem Wert eines geschützten Bytes eine unterschiedliche Speicherseite angefasst wird
    - dann Auslesen des Wert des Bytes durch Timing-Seitenkanal: zuletzt angefasste Speicherseite ist im Cache und schneller lesbar
- "On the Iceberg" - Auf dem Eisberg: [Schriftarten können Schadcode enthalten](https://superuser.com/a/1202556)
    - Kombo mit "Middle of the Iceberg" weiter unten
- "Below the Water" - Unter der Wasserlinie: [Magic the Gathering is Turing Complete](https://www.youtube.com/watch?v=pdmODVYPDLA)
    - vgl. unsere Besprechung von "aus Versehen turingvollständig" in STP042 sowie von Turingmaschinen in STP063
    - Kyle Hill hat dazu auch noch ein [neueres Video](https://www.youtube.com/watch?v=uDCj-QOp5gE) auf seinem eigenen Kanal
- "Middle of the Iceberg" - Mitte des Eisbergs: [Font-Hinting ist turingvollständig](https://en.wikipedia.org/w/index.php?title=TrueType&oldid=1305280909#Hinting_language)
    - noch ein Fall von "aus Versehen turingvollständig" (siehe STP042)
    - Hinting: Anpassung der Form einer Glyphe an die tatsächliche Schriftgröße in Pixeln etc., um ungewollte Rasterisierungs-Artefakte zu reduzieren ([grafisches Beispiel](https://commons.wikimedia.org/w/index.php?title=File:Font-hinting-example.png&oldid=808481633))
    - in TrueType: Hinting in einer VM, die durch Code in der Schriftartendatei gesteuert wird; damit Einschleusung von Schadcode möglich, wenn diese VM Schwachstellen hat
    - Problem: heute Schriftarten nicht nur im Betriebssystem vorinstalliert, sondern auch dynamisch geladen als Webfont
- "Bottom of the Iceberg" - Unterseite des Eisbergs: [Es gibt keine elegante Lösung für FizzBuzz](https://www.gayle.com/blog/2015/5/31/the-problem-with-the-fizzbuzz-problem)
- "Below the Iceberg" - Unterhalb des Eisbergs: [SAT-Solver werden immer schneller](https://ojs.aaai.org/index.php/aimagazine/article/view/2395)
    - SAT (von engl. "satisfiability"): [Erfüllbarkeitsproblem der Aussagenlogik](https://de.wikipedia.org/w/index.php?title=Erf%C3%BCllbarkeitsproblem_der_Aussagenlogik&oldid=256016000)
    - relevant deswegen, weil man viele Probleme als SAT-Problem umformulieren kann (z.B. Sudoku, Raumplanung, Paketmanager)
- "Deep Water" - Im tiefen Wasser: [FJCVTZS](https://developer.arm.com/documentation/dui0801/g/A64-Floating-point-Instructions/FJCVTZS)
    - FJCTVZS = "Floating-point Javascript Convert to Signed fixed-point, rounding toward Zero" (Konvertierung eines float64 in einen int32 mit Rundung zur Null hin und Modulo 2^32 bei Überlauf)
    - ARM ist zwar RISC (siehe STP075), aber eine spezielle Instruktion für Zahlenumwandlungen in JavaScript-Engines ist offenbar doch wertvoll
- "The Abyss" - Der Abgrund: [Die große Adressdatei der japanischen Post](https://www.dampfkraft.com/posuto.html)
    - eine Lektion darin, wie wichtig saubere Daten sind, und dass man auf bei der Erstellung von Dateien auf ihre Verwendung achten muss (Parallele zu Dokumentation)

Es gibt weiterhin mehr zu entdecken. Wir sehen uns wieder in STP105.
