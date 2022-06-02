- Zustand bei frühen Rechenmaschinen und Computern
    - kein Bootstrap: Strom an -> Programm starten, Programm beenden -> Strom aus (siehe [Konrad-Zuse-Clip](https://www.youtube.com/watch?v=n8Yo-wD-QTo ))
    - Programme waren spezifisch für die Hardware, auf der sie laufen
    - siehe auch: [damals (TM) Podcast](https://damals-tm-podcast.de/)
    - und [ChaosRadioExpress193 Old School Computing](https://cre.fm/cre193-old-school-computing)

- Zustand bei heutigen Computern
    - Programme benötigen Funktionen des Betriebssystems, Betriebssystem muss sich mit der Hardware vertraut machen -> separate Startphase notwendig
    - **Bootstrap** (oder kurz **Boot**): sich selbst an den Schnürsenkeln aus dem Sumpf ziehen; aus einem einfachen System heraus ein komplexeres System aktivieren
    - Konzept taucht mehrmals in verschiedenen Kontexten auf ("Wie stellt man einen Hobel her, wenn man keinen Hobel hat?"), siehe z.B. Bootstrapping von Programmiersprachen
    - Hardware-Bootvorgang meist in mehreren aufeinanderfolgenden Phasen

- Phase 1: Firmware
    - minimales Startprogramm
    - in einen separaten [Flash-Speicher](https://de.wikipedia.org/wiki/Flash-Speicher ) im Chip fest verbaut
    - Aufgabe: Hardware in einen definierten Zustand bringen (z.B. interne Speicher initialisieren), nächste Stufe finden und laden
    - heute mitunter alles andere als minimal: siehe [UEFI](https://de.wikipedia.org/wiki/Unified_Extensible_Firmware_Interface ) und [Intel ME](https://de.wikipedia.org/wiki/Intel#Intel_Management_Engine_seit_2008 )/[AMD PSP](https://de.wikipedia.org/wiki/AMD_Security_Processor )
    - bei x86 früher [BIOS](https://de.wikipedia.org/wiki/BIOS ); heute ersetzt durch UEFI, dessen definierter Endzustand modernen Konventionen folgt

- Phase 2: Bootloader
    - immer noch ziemlich minimal
    - liegt auf einem Massenspeicher (Festplatte, USB-Stick)
    - Aufgabe: Massenspeicher nach Betriebssystemen durchsuchen, Auswahldialog anbieten, gewähltes Betriebssystem starten
    - unter Linux meistens [GRUB](https://de.wikipedia.org/wiki/Grand_Unified_Bootloader ) oder [systemd-boot](https://en.wikipedia.org/wiki/Systemd-boot )

- Alternativen zum klassischen Bootloader
    - UEFI-Firmware kann direkt das Betriebssystem starten, sofern keine interaktive Auswahl erforderlich ist
    - [Netzwerk-Boot](https://de.wikipedia.org/wiki/Netzboot ) (entweder durch die Firmware oder als Option in Bootloadern wie GRUB)

- Phase 3: Betriebssystem
    - Erinnerung (siehe STP019): Kernel = Betriebssystemteile in höchster Privilegenstufe, Userspace = alle Programme in niedrigerer Privilegienstufe
    - Bootloader lädt und startet den Basisteil des Kernels
    - Kernel erkennt die verfügbare Hardware, lädt von der Festplatte die benötigten Kernel-Module nach und startet den System-Manager

- Phase 4: System-Manager
    - unter Unix meist als "PID 1" benannt, weil dieser erste Prozess mit der Prozess-ID (PID) 1 läuft
    - klassischerweise [sysvinit](https://de.wikipedia.org/wiki/SysVinit ), heute unter Linux meist [systemd](https://de.wikipedia.org/wiki/Systemd )
    - startet alle Userspace-Programme, die als Teil des Betriebssystems aufgefasst werden können (im Unix-Sprech "Daemons": Disk And Execution MONitors)
    - auf Systemebene z.B. Druckwarteschlange, Bluetooth-Dienst, Netzwerk-Konfigurationsdienst, Zeitsynchronisation
    - auf der Ebene einer grafischen Sitzung z.B. Dienste zum Auswählen von Tastaturlayout/Netzwerk, Einstellen der Lautstärke, Screen-Reader

    - [ChaosRadioExpress209 Das Linux System](https://cre.fm/cre209-das-linux-system)

- Phase 5: Display-Manager
    - bietet auf grafischen Systemen den Anmeldedialog
    - im Prinzip selbst eine komplette grafische Sitzung (Anekdote: GDM vs. LightDM)
    - startet nach erfolgreicher Anmeldung die entsprechende grafische Sitzung

- Phase 6: grafische Sitzung
    - Strukturierung je nach Desktopoberfläche
    - immer mit dabei: ein Window-Manager, der die einzelnen Fenster in ein Gesamtbild zusammensetzt (heutzutage auf der GPU, dann heißt der Window-Manager "Compositor")
    - unter Windows ist der Window-Manager aus historischen Gründen Teil von `explorer.exe`
    - wie besprochen: unter Linux Bootvorgang visualisieren mit `systemd-analyze plot > output.svg`
