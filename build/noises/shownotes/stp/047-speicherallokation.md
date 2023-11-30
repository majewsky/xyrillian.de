- Rückbezug auf STP045: **Heap** (Haufenspeicher)
    - keine Strukturvorgaben durch das Betriebssystem oder die Prozessorarchitektur
    - notwendig für Daten, deren Lebenszeit nicht der Struktur von Unterprogrammen folgt, oder deren Größe die 132 KiB des Stacks übersteigt
    - neue Frage: Wie kann man diesen unstrukturierten Haufen sinnvoll verwalten?
    - Analogie: Ein Team verwendet als Speicher ein großes Whiteboard, auf dem verschiedene Mitglieder des Teams rumschreiben wollen, ohne sich gegenseitig in die Quere zu kommen.

- man braucht einen [Allokator](https://de.wikipedia.org/w/index.php?title=Allokation_(Informatik)&oldid=199910001)
    - meist in Form einer Programmbibliothek, die auf die jeweilige Programmiersprache und/oder Laufzeitumgebung abgestimmt ist
    - klassischerweise zwei Funktionen: "malloc" ("memory allocate"; ich brauche ein Stück Speicher der Größe X) und "free" (das gegebene Stück Speicher brauche ich nicht mehr)
    - Allokationen mittels malloc sind meist sehr viel kleiner als die Speicherseiten, die das Betriebssystem rausgibt (min. 4 KB) -> Allokator holt sich ganze Speicherseiten und zerteilt die in mundgerechte Stücke
    - damit "free" funktioniert, muss der Allokator irgendwo über diese Zerstückelung Buch führen -> gewisser Mehraufwand an Speicher für die Buchführung
    - verschiedene Anwendungen machen ihre Speicherallokationen in unterschiedlichen Mustern (häufig z.B. ein gewisser Bodensatz an großen Allokationen, die für die ganze Lebensdauer des Programms erhalten bleiben, und dazu sehr viele kleine und kurzlebige Allokationen für bestimmte Einzeloperationen) -> interessantes Arbeits- und Forschungsfeld für Leute, die sich gern in Algorithmen eingraben

- Speicherverwaltung im Heap ist sicherheitskritisch -> mögliche Fehlerquellen
    - Speicherleck: Speicher wird allokiert, aber dann nicht wieder freigegeben; da der Allokator nicht wissen kann, ob der Speicher noch verwendet wird, muss er brachliegen
    - [Pufferüberlauf](https://de.wikipedia.org/w/index.php?title=Puffer%C3%BCberlauf&oldid=236800782): Allokation war für X Bytes, aber dann wird über das Ende (Basisadresse + X) hinausgeschrieben; dies zerstört dann Daten in anderen Allokationen oder Buchführungsdaten des Allokators
    - [Use after free](https://owasp.org/www-community/vulnerabilities/Using_freed_memory): Allokation wird mit "free" zurückgegeben, aber danach noch vom Programm verwendet, obwohl der Allokator dieses Stück Speicher vielleicht schon wieder anderweitig allokiert hat
    - Das klingt alles anstrengend. Kann man das nicht den Computer machen lassen? -> **automatische Speicherverwaltung**

- einfachste praktikable Idee: [Referenzen zählen](https://de.wikipedia.org/w/index.php?title=Referenzz%C3%A4hlung&oldid=221346853)
    - Ansatz: bei allokiertem Speicher nicht die lose Speicheradresse herumreichen, sondern noch einen Zähler mitführen
        - Zähler beginnt bei 1, da man am Anfang eine Referenz hat
        - bei jedem Kopieren der Referenz wird der Zähler um 1 erhöht, bei jedem Löschen einer Referenz um 1 verringert
        - die Allokation wird freigegeben, wenn der Zähler auf 0 fällt
    - Vorteil: sehr schnell, sehr zuverlässig, eliminiert bei konsequenter Anwendung das Risiko von "Use after free"
    - Nachteil: garantierte Speicherlecks bei zyklischen Referenzen (z.B. A verweist auf B verweist auf C verweist auf A, aber nichts anderes verweist auf A/B/C)
    - komplexere Formen von automatische Speicherverwaltung mit Zykluserkennung heißen meist [Garbage Collector](https://de.wikipedia.org/w/index.php?title=Garbage_Collection&oldid=236166816) (GC, "Müllsammler"); strenggenommen sind Referenzzähler "konservative GC"

- Beispiel für einen GC-Algorithmus: **Mark and Sweep**
    - aktives Aufräumen zu bestimmten Gelegenheiten, z.B. wenn bei malloc kein Speicher mehr frei ist und weiterer Speicher vom Betriebssystem angefordert werden müsste
    - Algorithmus:
        - Allokationen, deren Adressen im Stack liegen, werden als "erreichbar" markiert
        - alle anderen Allokationen werden initial als "nicht erreichbar" markiert
        - ausgehend von den erreichbaren Allokationen wird sukzessive nach Verweisen auf andere Allokationen gesucht und diese ebenfalls als "erreichbar" markiert
        - nachdem alles durchsucht wurde, werden alle Allokationen freigegeben, die immer noch als "nicht erreichbar" markiert sind
    - GC braucht meist eine enge Integration mit dem Allokator und mit der Programmiersprache oder der Laufzeitumgebung (z.B. insb. um die Frage "welche Allokationen sind im Stack referenziert" zu beantworten)
    - Nachteil: fast jeder GC muss zumindest für einen Teil der Aufräumphase "die Welt anhalten", was für den Nutzer wie Motorstottern aussieht

- Optimierung der Heap-Nutzung: Anzahl der Allokationen verringern (es gilt die Faustregel "der einzige Weg, zu optimieren, ist, weniger zu tun")
    - zusammenhängende Datenstrukturen möglichst in eine einzelne Allokation ablegen anstatt in mehrere Teilstücke ("Arena-Allokation", "Small Object Allocation")
    - kleine Daten bevorzugt im Stack halten und nicht auf den Heap legen (Xyrill hat eine Anekdote zu Minecraft)
    - noch eine Idee: GC zu opportunen Momenten explizit anstarten (z.B. bei Spielen während des VSync, d.h. wenn ein Einzelbild gerade fertig ist und es auf den Bildschirm kopiert wird)

- zum Schluss ein Blick auf ein Teilgebiet der Heap-Algorithmen: [Zwischenspeicher (Cache)](https://de.wikipedia.org/w/index.php?title=Cache&oldid=236047086)
    - Grundproblem: bestimmte Operationen sind teuer (z.B. Netzwerkabfragen, Einlesen von Dateien von der Festplatte, komplexe Berechnungen), aber reproduzierbar; wir würden die Ergebnisse gerne wiederverwenden, aber können nicht alle im Speicher halten
    - Speicherstück mit einer festen Größe wird als Cache designiert
    - idealerweise würde man genau die Ergebnisse dort ablegen, die in der Zukunft am häufigsten wieder nachgefragt werden
    - Problem: ["Vorhersagen sind schwierig, insbesondere, wenn sie die Zukunft betreffen"](https://en.wikiquote.org/w/index.php?title=Niels_Bohr&oldid=3344877) -> Cache-Replacement-Algorithmen sind ein Forschungsfeld
    - Beispiel **LRU** ("Least Recently Used"): wenn ein neues Element abgelegt wird, wird dasjenige existierende Element überschrieben, das am längsten ungenutzt ist

- Lesestoff: [www.memorymanagement.org](https://www.memorymanagement.org/) (englisch)
