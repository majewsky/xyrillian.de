- Rückbezug und Abgrenzung zu STP015 (Multitasking in Betriebssystemen)
    - Definition von [Nebenläufigkeit](https://de.wikipedia.org/w/index.php?title=Nebenl%C3%A4ufigkeit&oldid=242751497): "in der Informatik die Eigenschaft eines Systems, mehrere Aufgaben, Berechnungen, Anweisungen oder Befehle gleichzeitig ausführen zu können"
    - Definition von [Multitasking](https://de.wikipedia.org/w/index.php?title=Multitasking&oldid=241562607): "die Fähigkeit eines Betriebssystems, mehrere Aufgaben \[...] (quasi-)nebenläufig auszuführen"
    - eins definiert das andere \o/ -> wir schauen auf den Begriffsgebrauch in der Praxis
    - Multitasking: die funktionale Umsetzung einer Multiprozess-Architektur in Hardware und Software (auf Betriebssystem-Ebene)
    - Nebenläufigkeit: die Ertüchtigung von Userspace-Programmen zur Ausnutzung dieser Möglichkeiten unter Wahrung des korrekten Verhaltens

- Grundproblem: Wie vermeidet man Konflikte und Verwirrung beim Umgang mit geteilten Ressourcen?
    - "Ressource" bedeutet vor allem: Speicherstellen, Dateisystem-Einträge (Dateien und Verzeichnisse), Geräte, (Aufmerksamkeit der Benutzerin)
    - explizit nicht Zeit; darum kümmert sich bereits die Multitasking-Unterstützung des Betriebssystems

- **Race**: eine Situation, bei der das Ergebnis (und insbesondere die Korrektheit) mehrerer nebenläufiger Prozesse davon abhängt, in welcher Reihenfolge die einzelnen Rechenschritte verschiedener Prozesse zufälligerweise ausgeführt werden
    - allgemein bekannt als **Race Condition** ([Wettlaufsituation](https://de.wikipedia.org/w/index.php?title=Wettlaufsituation&oldid=203715002)) oder beim Speicherzugriff insbesondere **Data Race**
    - Beispielsituation: im Arbeitsspeicher liegt ein Zähler mit aktuellem Wert 40; zwei Prozesse A und B wollen diesen Zähler gleichzeitig um 1 erhöhen -> erwarteter Endwert 42
    - Problem: "Zahl im Arbeitsspeicher verändern" ist nicht, wie Speicherzugriff in CPUs funktioniert (siehe STP007); tatsächlich sind jeweils drei Schritte erforderlich (Einlesen in CPU-Register, Erhöhen um 1, Zurückschreiben in den RAM)
    - möglicher Ausgang: beide Prozesse laufen auf verschiedenen CPUs, lesen gleichzeitig den Wert 40 in ihre CPU-Register, erhöhen gleichzeitig auf 41, schreiben dies zurück -> Ergebnis 41 statt 42
    - "auf verschiedenen CPUs" ist hier nicht erforderlich: z.B. A liest ein und erhöht, wird unterbrochen, B liest ein und erhöht, B schreibt zurück, wird unterbrochen, A schreibt zurück
    - "zwei Prozesse" ist auch nicht erforderlich: Prozesse können auch in **Threads** (parallele Ausführungsstränge) unterteilt sein, die nebenläufig Code ausführen, aber ansonsten fast alle Ressourcen (Speicherseiten, offene Dateien, etc.) teilen

- wir brauchen ein [Mutex](https://de.wikipedia.org/w/index.php?title=Mutex&oldid=217939607): einen Mechanismus zum wechselseitigen Ausschluss ("**Mut**ual **Ex**clusion")
    - Problem: Wie implementiert man sowas?

- Idee: bevor wir den Zähler anfassen, fragen wir bei einem zentralen Prozess nach einer Sperre für diesen Zähler an; dieser Prozess vermerkt Sperr- und Entsperrvorgänge in seinem internen Speicher
    - dieser Kontrollprozess könnte auch einfach ein Teil des Betriebssystems sein und der Sperr-/Entsperrvorgang ein Syscall (siehe STP019)
    - Vorteil: innerhalb dieses Kontrollprozesses keine Nebenläufigkeit und damit keine Gefahr eines Data Race
    - Nachteil: Interprozess-Kommunikation ist vergleichsweise grauenhaft langsam (Millisekunden vs. Mikrosekunden)

- Idee: in der **kritischen Region** (von Auslesen des Zählers bis Zurückschreiben) verbieten wir dem Betriebssystem, unseren Prozess zu unterbrechen
    - Problem 1: hilft nur bei nebenläufigen Prozessen auf demselben CPU-Kern
    - Problem 2: immer noch ein teurer Syscall
    - Problem 3: böswillige Prozesse könnten einfach ihre gesamte Laufzeit als kritische Region markieren und die Rechenzeit blockieren

- praktische Umsetzung von Mutexen mittels [Atomics](https://de.wikipedia.org/w/index.php?title=Atomare_Operation&oldid=232565909): spezielle CPU-Instruktionen, die nicht unterbrochen werden können
    - Beispiel für Mutex: "Fetch and Add" liest einen Wert aus dem Speicher aus, erhöht ihn um das Argument, und schreibt den erhöhten Wert zurück
    - schneller als ein Kontextwechsel zu einem Kontrollprozess oder ein Syscall
    - langsamer als ein normaler Speicherzugriff, da eventuell Caches ignoriert oder aktiv geleert werden müssen
    - in der Praxis evtl. Kombination mit Syscalls, um bei blockiertem Mutex den Prozess zu unterbrechen (z.B. unter Linux das ["Fast Userspace Mutex" bzw. Futex](https://de.wikipedia.org/w/index.php?title=Futex&oldid=242228263))

- andere Perspektive, hier zitiert aus der Programmiersprache Go: ["Do not communicate by sharing memory; instead, share memory by communicating."](https://go.dev/doc/effective_go)
    - statt Zugriffssicherungen für geteilten Speicher dort eher Nutzung von "Kanälen" (**Channels**) zur Nachrichtenübermittlung zwischen Threads
    - Beispiel "Worker Pool": mehrere gleichartige und voneinander unabhängige Teilaufgaben sind abzuarbeiten (z.B. 100 Bilder in ein anderes Dateiformat umwandeln)
    - Idee: ein Worker (Arbeits-Prozess oder Arbeits-Thread) pro CPU-Kern; außerdem ein zentraler Prozess, der die Aufgaben verteilt; Zentrale stellt alle Dateinamen in einen Kanal, Arbeiter greifen nacheinander aus dem Kanal die Dateinamen heraus
    - unter der Haube nutzt der Kanal Atomics, um sich vor Data Races zu schützen
    - Rückbezug zu STP027: sobald man mehrere Threads untereinander koordinieren muss, hat man das ganze Problemfeld "Verteilte Systeme", was nach Xyrills Erfahrung nochmal wesentlich nerviger ist als Data Races

- Abendgedanken: [Amdahl'sches Gesetz](https://de.wikipedia.org/w/index.php?title=Amdahlsches_Gesetz&oldid=233711717)
    - mehr CPU-Kerne machen nur Dinge schneller, die wahrhaft nebenläufig sind
