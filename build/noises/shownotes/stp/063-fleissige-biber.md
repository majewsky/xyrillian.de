Wir haben Feedback bekommen, siehe unten. Zuerst die Shownotes zum eigentlichen Thema:

- aktueller Anlass: [Der fünfte fleißige Biber wurde gefunden](https://www.quantamagazine.org/amateur-mathematicians-find-fifth-busy-beaver-turing-machine-20240702/)
    - Rückbezug zu STP000: theoretische Informatik ist ein Teilgebiet der Mathematik, dass Xyrill lieber "Berechnungstheorie" nennen würde
    - hierin insbesondere "Berechenbarkeitstheorie": Welche Probleme lassen sich mit Rechenmaschinen lösen?
    - [Church-Turing-These](https://de.wikipedia.org/w/index.php?title=Church-Turing-These&oldid=244116059): Berechenbarkeit ist eine intrinsische Eigenschaft eines Problems und nicht abhängig davon, welchen Computer ich zur Lösung einsetze
        - "Die Klasse der Turing-berechenbaren Funktionen stimmt mit der Klasse der intuitiv berechenbaren Funktionen überein."
    - Analyse der Berechenbarkeit eines Problems anhand von einfachen Modellmaschinen, z.B. der Turing-Maschine (daher "Turing-berechenbar")
    - [Turing-Vollständigkeit](https://de.wikipedia.org/w/index.php?title=Turing-Vollst%C3%A4ndigkeit&oldid=242871861): die Eigenschaft einer Rechenmaschine, alle Probleme lösen zu können, die auch eine Turing-Maschine lösen kann (mal abgesehen von weltlichen Einschränkungen wie Speicherplatz und Rechenzeit)
    - Rückbezug zu STP021 (Reguläre Ausdrücke): dort hatten wir stärker limitierte Maschinenmodelle erwähnt (endliche Zustandsautomaten, Kellerautomaten)

- Aufbau einer [Turing-Maschine](https://de.wikipedia.org/w/index.php?title=Turingmaschine&oldid=244930504) ([Beispielbild aus Wikipedia](https://commons.wikimedia.org/wiki/File:Model_of_a_Turing_machine.jpg))
    - unendlich langes Speicherband mit Feldern, auf denen entweder 0 oder 1 stehen kann
    - Lese-/Schreibkopf, der das aktuelle Feld lesen und schreiben kann sowie sich in Einzelschritten entlang des Bandes bewegen kann
    - Programm: endliche Menge von Zuständen, und eine Zustandsübergangsfunktion, die auf Basis der gelesenen Zeichen entscheidet und Zeichen überschreibt sowie den Lesekopf bewegt
    - Xyrill hat ein Beispiel mitgebracht: ein Programm, dass zwei positive Ganzzahlen subtrahiert; die Zahlen sind als Folgen von 1 gegeben, getrennt von einer 0; zum Beispiel "3-2" wird eingegeben als "111011" und der Lesekopf startet am Anfang der ersten Zahl
    - **Warnung:** Wir haben während der Aufzeichnung bemerkt, dass diese Maschine inkorrekt ist. Damit keine Verwirrung entsteht, haben wir die Notizen nicht angepasst.

- Zustandsvorrat (durchbuchstabiert in der Reihenfolge der Benutzung, Startzustand ist A):
    - A: suche nach rechts in der ersten Zahl
    - B: suche nach rechts in der zweiten Zahl
    - C: auf dem Ende der zweiten Zahl
    - D: suche nach links in der zweiten Zahl
    - E: suche nach links nach der ersten Zahl

- Übergangstabelle:
    - Zustand A, lese 0 -> schreibe 0, gehe nach rechts, neuer Zustand B
    - Zustand A, lese 1 -> schreibe 1, gehe nach rechts, neuer Zustand A
    - Zustand B, lese 0 -> schreibe 0, gehe nach links, neuer Zustand C
    - Zustand B, lese 1 -> schreibe 1, gehe nach rechts, neuer Zustand B
    - Zustand C, lese 0 -> schreibe 0, gehe nach links, neuer Zustand HALT
    - Zustand C, lese 1 -> schreibe 0, gehe nach links, neuer Zustand D
    - Zustand D, lese 0 -> schreibe 0, gehe nach links, neuer Zustand E
    - Zustand D, lese 1 -> schreibe 1, gehe nach links, neuer Zustand D
    - Zustand E, lese 0 -> schreibe 0, gehe nach links, neuer Zustand E
    - Zustand E, lese 1 -> schreibe 0, gehe nach rechts, neuer Zustand A

- Beispielausführung für die Aufgabe "3-2"
  ```
  ... 0 [1] 1  1  0  1  1  0 ... Zustand A
  ... 0  1 [1] 1  0  1  1  0 ... Zustand A
  ... 0  1  1 [1] 0  1  1  0 ... Zustand A
  ... 0  1  1  1 [0] 1  1  0 ... Zustand A
  ... 0  1  1  1  0 [1] 1  0 ... Zustand B
  ... 0  1  1  1  0  1 [1] 0 ... Zustand B
  ... 0  1  1  1  0  1  1 [0]... Zustand B
  ... 0  1  1  1  0  1 [1] 0 ... Zustand C
  ... 0  1  1  1  0 [1] 0  0 ... Zustand D
  ... 0  1  1  1 [0] 1  0  0 ... Zustand D
  ... 0  1  1 [1] 0  1  0  0 ... Zustand E
  ... 0  1  1  0 [0] 1  0  0 ... Zustand A
  ... 0  1  1  0  0 [1] 0  0 ... Zustand B
  ... 0  1  1  0  0  1 [0] 0 ... Zustand B
  ... 0  1  1  0  0 [1] 0  0 ... Zustand C
  ... 0  1  1  0 [0] 0  0  0 ... Zustand D
  ... 0  1  1 [0] 0  0  0  0 ... Zustand E
  ... 0  1 [1] 0  0  0  0  0 ... Zustand E
  ... 0  1  0 [0] 0  0  0  0 ... Zustand A
  ... 0  1  0  0 [0] 0  0  0 ... Zustand B
  ... 0  1  0 [0] 0  0  0  0 ... Zustand C
  ... 0  1 [0] 0  0  0  0  0 ... Zustand HALT
  ```

- Anmerkungen hierzu
    - man kann Turing-Maschinen auch auf viele andere Weisen modellieren, etwa mit größeren Zeichensatz (z.B. alle Dezimalziffern sowie mathematische Symbole, um eine Eingabe wie `127*(3+6/2)` zu ermöglichen) oder mehreren Bändern; an der Berechenbarkeit ändert das nichts
    - Xyrill belässt es bei obigem Beispiel, dass noch einigermaßen kompakt ist
    - Übungsaufgabe: Füge weitere Zustände ein, mit denen der Lesekopf am Ende in die Startposition zurückbewegt wird.
    - Aber was ist jetzt mit den Bibern?

- [N-ter fleißiger Biber](https://de.wikipedia.org/w/index.php?title=Flei%C3%9Figer_Biber&oldid=246581342): diejenige Turingmaschine mit N Zuständen, die auf ein initial leeres Band (alles voller Nullen) die meistmöglichen Einsen schreiben kann, bevor sie anhält
    - Die Haltebedingung ist wichtig! Sonst können wir mit nur einem Zustand unendlich viele Einsen schreiben ("Zustand A, lese irgendwas -> schreibe 1, gehe nach rechts, neuer Zustand A")
    - Problem: Man kann einer Turingmaschine nicht ohne weiteres ansehen, ob sie terminiert oder nicht. -> Das [Halteproblem](https://de.wikipedia.org/w/index.php?title=Halteproblem&oldid=239308188) ist eines der klassischen unberechenbaren Probleme.
    - Jetzt gefunden wurde `BB(5,2) = 47176870`; sprich: Der längstmögliche endliche Lauf einer Turingmaschine mit 5 Zuständen und 2 Zeichen auf initial leerem Band ist 47176870 Schritte. (Das ist die gleiche Maschinengröße wie unser Beispiel oben.)
    - Bei der Gelegenheit erwähnen wir auch mal die [OEIS](https://de.wikipedia.org/w/index.php?title=On-Line_Encyclopedia_of_Integer_Sequences&oldid=232132016). Die fleißigen Biber sind zum Beispiel [A060843](https://oeis.org/A060843).

#### Feedback zu STP038

laalsaas schreibt:

> In Minute 32 spricht xyrill von 1,6 Terabit und sagt, dass wären 1,6 Millionen Bit. Ich nehme an, dass es sich hierbei um einen Versprecher im Eifer des Gefechts handelt, sind doch 1 Terabit eine Billion Bit.

Anmerkung: Im Englischen sogar "a trillion bit". :)

> Etwas früher ging es um USB und darum, dass man den Stecker nie beim ersten Versuch richtig herum reinbekommt. Dazu möchte ich hier einen „Lifehack“ teilen. Ich hoffe, es gelingt mir, ihn in Textform verständlich zu beschreiben. Ein USB-Stecker ist meist folgendermaßen aufgebaut: Wenn man von vorne daraufguckt, ist der äußere Rand ein rechteckig gebogenes Metallblech. Innen ist eine Hälfte des Rechtecks eine Plastik-„Zunge“, auf der die Kontakte sitzen, in der anderen Hälfte ist Luft. Eine USB-Buchse wiederum besteht aus einem Rechteck, in dem die Hälfte ebenfalls aus einer „Zunge” besteht, die andere aus Luft. Damit nun ein Stecker in die Buchse passt, müssen sich die „Zungen“ gegenüberliegen, damit sie die Luft in der jeweils anderen Komponente ausfüllen. Nun ist es so, dass bei über 90% der horizontal fest verbauten USB-buchsen in Laptops, PCs, usw die „Zunge“ oben ist. Das heißt, man muss einfach die „Zunge“ des USB-steckers nach unten orientieren. Wer immer noch zu faul ist, jedes mal frontal auf seinen USB-Stecker draufzugucken, der sei darauf hingewiesen, dass man auch einfach herausfinden kann, auf welcher Seite die Zunge ist, indem man durch die seitlichen Löcher im Blech des USB-Steckers schaut und darauf achtet, ob da direkt Plastik ist (Zungenseite, die muss nach unten) oder erst etwas Luft und dann Kontakte, wenn man das sieht, hält man den Stecker richtig herum. Seitdem ich diesen Trick kenne, habe ich wesentlich weniger Probleme damit, USB-Stecker richtigherum reinzustecken. Das funktioniert natürlich nicht bei freifliegenden Hubs und auch bei seitlich oder „nach unten“ montierten Buchsen habe ich noch keine Konvention erkennen können, das kann aber auch daran liegen, dass ich zu selten mit denen zu tun habe. Ich hoffe trotzdem, dass ich manchem mit diesem Trick das Leben einfacher machen konnte :)
>
> Zu guter Letzt möchte ich noch anmerken, dass die in Minute 39 von timeless gewählte Formulierung „Der LaTeX-Bruder“ eine (vermutlich ungewollte) Doppelbödigkeit hat. Erfahrene STP-Hörer wissen natürlich, dass es sich hier um Xyrills Bruder handelt, der seine Abschlussarbeit in LaTeX geschrieben hat, aber ich finde, es klingt auch so, wie ein Mitglied des LaTeX-Klosters oder der LaTeX-Sekte, was mich sehr zum Schmunzeln gebracht hat.

#### Noch mehr Feedback

Weiterhin schreibt Matthias:

> Statt Postkarte: Liebe Hörergrüße aus dem Urlaub in Allgäu sendet Matthias.

Anbei ein Bild eines Gedenksteins für Konrad Zuse.
