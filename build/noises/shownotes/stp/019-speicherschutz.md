- "Welchen Speicher müssen wir eigentlich schützen?"
    - hier nur Arbeitsspeicher und Gerätespeicher
    - andere Speicherarten (z.B. Festplatten) geschützt durch die Betriebssystem-Abstraktionen, siehe zukünftige Folge zu Privilegienkontrolle

- Aufgaben des Betriebssystems: Prozessisolation und Vermittlung des Hardwarezugriffs
    - Problem: das Betriebssystem ist auch nur "ein Programm"
    - Frage 1: Wie wird sichergestellt, dass das Betriebssystem die finale Kontrolle über Hardwarezugriffe hat?
    - Frage 2: Wie wird verhindert, dass ein Speicherbereich, den Prozess A verwendet, von Prozess B überschrieben (oder unberechtigterweise ausgelesen) wird?

- Antwort auf Frage 1: [Prozessor-Ringe](https://de.wikipedia.org/wiki/Ring_(CPU) )
    - Standardfunktion auf allen CPUs, die auf Mehrprozess-Betrieb ausgelegt sind (nicht immer im Embedded-Bereich)
    - zu jedem Zeitpunkt läuft der Prozess in einem von mehreren Ringen
    - sensible Operationen (v.a. Hardwarezugriff) sind nur in den höheren Ringen erlaubt und damit den Betriebssystemprozessen vorbehalten
    - Wechsel in einen niedrigeren Ring jederzeit möglich (z.B. Aktivierung eines Userspace-Prozesses durch das Betriebssystem)

- Wie kommen wir von einem niedrigeren Ring wieder in einen höheren Ring?
    - Interrupts (siehe STP015): Eintreten eines Hardware-Ereignisses (z.B. ankommendes Netzwerk-Paket oder abgelaufener Hardware-Timer), das durch das Betriebssystem behandelt werden muss
    - **Syscalls**: niedrigprivilegierter Prozess fragt eine definierte Schnittstelle des Betriebssystems an, um eine höherprivilegierte Operation auszuführen
    - z.B. Dateisystem-Zugriff, Hardware-Zugriff, Ändern der Betriebssystem-Konfiguration, Starten von neuen Prozessen, Nachrichtenübermittlung an andere Prozesse
    - dadurch Realisierung der Prozessisolation, weil das Betriebssystem die Kontrolle über alle privilegierten Operationen hat

- Antwort auf Frage 2: [Virtuelle Speicherverwaltung](https://de.wikipedia.org/wiki/Virtuelle_Speicherverwaltung )
    - wenn Prozesse auf Speicher zugreifen, verwenden sie nur "virtuelle" Speicheradressen: z.B. Adresse 0x42 für Prozess A ist nicht unbedingt derselbe Speicher wie Adresse 0x42 für Prozess B
    - Abbildung auf physische Speicheradressen innerhalb der CPU durch eine **Memory Management Unit** (MMU)
    - Konfiguration der MMU durch das Betriebssystem immer kurz vor der Übergabe der CPU an einen anderen Prozess
    - Aufteilung des physischen/virtuellen Speichers in Seiten (**Pages**), z.B. bei x86-64 wahlweise 4 KiB oder 2 MiB oder 1 GiB
    - pro Page verschiedene Zugriffsberechtigungen möglich, z.B. zur Verhinderung des Schreibens in ausführbare Programmteile, oder für **Shared Memory** als Kommunikationsweg zwischen Prozessen

- Paging der MMU ermöglicht **Swapping**
    - Pages müssen nicht unbedingt im tatsächlichen Arbeitsspeicher gehalten werden, sondern können bei Nichtverwendung in eine Swap-Datei auf der Festplatte ausgelagert werden
    - dadurch über alle Programme summiert meist deutlich mehr virtueller Speicher zugewiesen, als tatsächlich physischer Speicher vorhanden ist
    - wenn ein Programm auf eine im Arbeitsspeicher fehlende Seite zugreift, erzeugt die MMU einen Interrupt (**Page Fault**), sodass das Betriebssystem die fehlende Seite transparent nachladen kann
    - analog dazu: **Memory-Mapping** von Dateiinhalten direkt in den Arbeitsspeicher

- wenn wir schon mal über Speicheradressen reden: [Direct Memory Access](https://de.wikipedia.org/wiki/Direct_Memory_Access )
    - Abbildung eines Arbeitsspeicher-Addressbereiches auf einen Speicher eines angeschlossenen Gerätes, zum Beispiel den Arbeitsspeicher der Grafikkarte oder eines Netzwerkadapters
    - dadurch direkter Transfer von Daten zwischen niedrigprivilegierten Prozessen und Hardware möglich, z.B. schnelles Hochladen von Texturen in die Grafikkarte durch ein 3D-Spiel
    - DMA braucht nicht unbedingt eine MMU (siehe [VGA-Buffer bei `0xB8000`](https://wiki.osdev.org/Printing_to_Screen ) in DOS-Ära-Betriebssystemen)
