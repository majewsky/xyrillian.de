- Rückblick: Logikgatter (STP003)
    - aus Gattern lassen sich Schaltkreise bilden, die konkrete Einzelberechnungen anstellen
        - hauptsächlich einfache Rechenoperationen wie "zwei 16-Bit-Zahlen addieren"
        - aber auch teils sehr komplexe Einzeloperationen, z.B. Teilschritte einer bestimmten Hashfunktion (siehe STP004) oder eines bestimmten Audio/Video-Codecs (heißt in Produktbeschreibungen meist **Hardware-Beschleunigung**)
    - Frage 1: Wie kann man mehrere Operationen nacheinander ausführen?
    - Frage 2: Wie kann man in einem einzelnen Taktschritt aus verschiedenen möglichen Operationen wählen?
    - Frage 3: Wie entsteht daraus ein ganzes Programm, also eine bestimmte Folge von verschiedenen Operationen?

- Antwort auf Frage 1: [Taktgenerator](https://de.wikipedia.org/wiki/Taktgenerator_(Computer) )
    - analog zum Taktgeber in einer Uhr (Unruh bei mechanischen Uhren, Quartzoszillator bei elektronischen Uhren), aber als mikroeletronische Schaltung realisiert
    - Geschwindigkeitsvergleich: bei Quartzuhren hat der Oszillator meist 32768 Hz
    - heutzutage ist die Taktfrequenz meistens regelbar
        - 1990er: Turbo-Schalter am Gehäuse zum Verdoppeln (bzw. in der anderen Richtung Halbieren) der Taktfrequenz
        - seit den 2000ern: [ACPI](https://en.wikipedia.org/wiki/Advanced_Configuration_and_Power_Interface ) und Co., Energiesparen durch bedarfsgerechtes Heruntertakten der CPU-Kerne
        - Heruntertakten als Schutz vor Überhitzung (nach Überlastung oder beim Versagen der Kühlung)
    - unter Linux: aktuelle Taktfrequenz auslesen mittels `grep '^cpu MHz' /proc/cpuinfo`

- Antwort auf Frage 2: [Arithmetisch-logische Einheit](https://de.wikipedia.org/wiki/Arithmetisch-logische_Einheit ) (ALU)
    - Idee: neben 2N Eingängen für zwei N-Bit-Eingabezahlen noch weitere Eingänge, die einen **Opcode** darstellen, also eine fortlaufende Nummer für die gewünschte Rechenoperation
    - Opcode bestimmt, von welchem Teilschaltkreis das Ergebnis nach außen durchgereicht wird
    - hierzu zusätzliche Logikschaltung am Ende: `Ausgabe = ((Opcode = 1) AND Ausgabe1) OR ((Opcode = 2) AND Ausgabe2) OR ...`
    - analog zur ALU gibt es meist noch eine FPU für Fließkommazahl-Operationen

- Antwort auf Frage 3: [Steuerwerk](https://de.wikipedia.org/wiki/Steuerwerk ) in der CPU (oder GPU)
    - Programmdatei besteht aus einer Folge von Befehlen, die im Arbeitsspeicher vorliegen
    - Steuerwerk hat einen **Befehlszeiger** (Instruction Pointer), der auf den jeweils nächsten Befehl zeigt

- Befehle sind in **Maschinensprache** kodiert
    - meist 16/32/64 Bit groß (je nach Prozessorarchitektur)
    - für Rechenbefehle wie "Addition" oder "Multiplikation"
        - den Opcode (welche Berechnung auszuführen ist)
        - aus welchen Prozessorregister die Eingabe kommt (evtl. ist ein Eingabewert eine Konstante wie "0" oder "1" und steckt direkt im Befehl drin)
        - in welches Prozessorregister die Ausgabe geht
    - andere Befehlsarten:
        - Load: aus dem Hauptspeicher einen Wert in ein Prozessorregister holen (Parameter: Quelladresse und Zielregister)
        - Store: genau andersherum (Parameter: Quellregister und Zieladresse)
        - Compare: schickt zwei Zahlen aus Registern an die ALU, aber statt einer Ergebniszahl gibt es einen Satz von **Flags** wie "A ist gleich B" oder "A ist kleiner als B" zurück
        - Jump: aktualisiert den Befehlszeiger, evtl. nur wenn ein bestimmtes Flag gesetzt ist (zum Beispiel Gleichheits-Flag aus einem vorherigen Compare oder Überlauf-Flag aus einer vorherigen Addition)
        - Call/Return: Call macht einen Sprung, aber speichert vorher die aktuelle Position in einen designierten Teil des Arbeitsspeichers; Return springt an den Ursprung des letzten Call zurück (damit verschachtelbare Unterprogramme realisierbar)

- Ausführung eines Befehls ist unterteilt in mehrere Schritte
    - grobe Skizze der Abarbeitung eines einzelnen Befehls (Schritt 2-5 jeweils nur, sofern für den entsprechenden Befehl zutreffend)
        - Schritt 1: Auslesen und Dekodieren des Befehls
        - Schritt 2: Laden der Eingabewerte aus den entsprechenden Registern
        - Schritt 3: Ausführen des Befehls (z.B. bei Rechenbefehlen mittels der ALU oder FPU)
        - Schritt 4: Speichern des Ausgabewertes im entsprechenden Zielregister
        - Schritt 5: Verschieben des Befehlszeigers auf den Folgebefehl
    - in der Praxis meist noch viel mehr Teilschritte (z.B. moderne x86-CPUs: über 40 Pipeline-Schritte)
    - **Pipelining**: wenn ein Befehl einen bestimmten Schritt durchlaufen hat, kann hier meist schon mit dem nächsten Befehl begonnen werden, auch wenn der vorherige Befehl noch nicht komplett abgearbeitet war
    - **spekulative Ausführung** (*Speculative Execution*): nachfolgende Befehle rechnen mit den Daten, von denen die CPU glaubt, dass sie wahrscheinlich aktuell sein werden, sobald derjenige Befehl tatsächlich an der Reihe ist
    - **Sprungvorhersage** (*Branch Prediction*): bei einem Sprungbefehl mit Vorbedingung muss die CPU möglichst gut abschätzen, ob die Vorbedingung eintritt und der Sprung vollzogen wird, um die nächsten Befehle vorhersagen zu können
    - Verbesserungen in spekulativer Ausführung und Sprungvorhersage sind einer der Hauptgründe, warum Desktop-CPUs auch ohne höhere Taktfrequenzen immer schneller geworden sind

- Überlegungen für schnelle Programme
    - Sprungbefehle sind für schnelle Programme ein Problem: wenn die CPU einen Sprung falsch vorhersagt, muss nach dem Ausführen des Sprungs die gesamte Pipeline verworfen und neu gefüllt werden
    - verschiedene Befehle dauern unterschiedlich lange in der Berechnung (z.B. Multiplikation ist langsamer als Addition ist langsamer als Bit-Shift); entsprechende Austauschungen finden in Compilern meist automatisch statt, z.B. `2 * x -> x + x` oder `8 * x -> x << 3`
