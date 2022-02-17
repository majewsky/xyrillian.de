* Vorbemerkung: Begriffsklärung Programm vs. Prozess
    * Programm: ein bestimmtes Stück Maschinencode, das zumeist auf der Festplatte vorgehalten wird
    * Prozess: eine bestimmte Laufzeitinstanz eines Programms
    * Prozess = Programm + zugehörige Daten im Speicher + zugehörige Berechtigungen etc.

* [Multitasking](https://de.wikipedia.org/wiki/Multitasking ): Ausführung mehrerer Aufgaben zur selben Zeit oder abwechselnd in kurzen Zeitabschnitten
    * konkurrente Ausführung: abwechselnd in kurzen Zeitabschnitten
    * parallelle Ausführung: zur selben Zeit (z.B. Parallelbearbeitung mehrerer Daten in der GPU, Parallelität in mehreren CPU-Kernen); heute nicht so sehr im Fokus

* Analogie: Multitasking [beim Menschen](https://de.wikipedia.org/wiki/Multitasking_(Psychologie) )
    * Idee 1: tatsächlich parallele Ausführung, z.B. gehen und gleichzeitig sprechen
    * Idee 2: Teilaufgaben an einer jeweils sinnvollen Stelle unterbrechen und während der Wartezeiten etwas anderes machen, z.B. Kochen/Backen (entspricht dem kooperativen Multitasking)
    * Idee 3: Teilaufgaben nach einer bestimmten Zeit oder bei äußeren Ereignissen unterbrechen, z.B. Büroarbeit (entspricht dem präemptiven Multitasking)

* kooperatives Multitasking: Prozesse unterbrechen sich selbst, wenn sich abzeichnet, dass sie warten müssen
    * z.B. auf das Lesen von Daten von der Festplatte oder eine Antwort aus dem Netzwerk
    * Problem 1: jedes Programm muss Unterbrechungspunkte enthalten
    * Problem 2: unkooperative Programme (z.B. fehlerhafte oder Schadsoftware) können das gesamte System blockieren
    * weit verbreitet in der Ära der frühen Heimcomputer (C64, DOS-PCs, Macs bis Mac OS 9)
    * heute nicht mehr im Einsatz, da moderne Prozessoren interruptfähig sind und somit präemptives Multitasking möglich ist

* präemptives Multitasking: Prozesse werden nach dem Ablauf ihrer zugeteilten Zeitscheibe unterbrochen, damit jemand anderes drankommen kann
    * bevor der Prozess drankommt, stellt das Betriebssystem einen Hardware-Timer in der CPU
    * nach Ablauf des Timers löst die CPU einen **Interrupt** (eine Unterbrechung) aus
    * CPU unterbricht die Arbeit am Prozess und führt eine vom Betriebssystem definierte ISR (**Interrupt Service Routine**) aus
    * ISR speichert den Zustand des Prozesses (insb. die Inhalte aller Prozessor-Register) und gibt die Ausführung zurück an den Scheduler
    * um Prozess später wiederherzustellen, werden im Prinzip einfach alle Prozessor-Register auf die gesicherten Werte zurückgesetzt

* **Scheduler** (Zeitplaner): Komponente des Betriebssystems, die den jeweils nächsten Prozess wählt, der laufen darf
    * interne Buchführung des Schedulers unterscheidet zwischen lauffähigen und nicht lauffähigen Prozessen
    * ist ein laufender Prozess nicht mehr lauffähig (weil auf I/O gewartet werden muss), kann er auch selbst einen Interrupt auslösen und wird dann im Scheduler als nicht lauffähig markiert
    * nächster Prozess wird immer aus der Liste der lauffähigen Prozesse gewählt, nach Kriterien wie Priorität (z.B. Audio-Prozesse)
    * gibt es keinen lauffähigen Prozess, kann die CPU bis zum nächsten Interrupt angehalten werden

* Wie werden Prozesse wieder lauffähig?
    * Eingabe von der Hardwareseite (z.B. Netzwerk, Festplatte) erzeugt ebenfalls Interrupts, auf die das Betriebssystem reagiert
    * Eingabe von anderen Prozessen (z.B. via Pipes) landet direkt beim Betriebssystem
    * Betriebssystem stellt die Daten dem entsprechenden Prozess bereit und markiert diesen wieder als lauffähig
    * analog bei vollgelaufenen Ausgabepuffern
