- [blockbasierter Speicher ("Block Storage")](https://de.wikipedia.org/wiki/Datenblock )
    - Festplatten handeln immer nur in festen Einheiten (Blöcken)
    - traditionellerweise 512 Byte, heutzutage 4 KiB

- kurzer Überblick zur physikalischen Anbindungen blockbasierter Speicher
    - [USB (Universal Serial Bus)](https://de.wikipedia.org/w/index.php?title=Universal_Serial_Bus&oldid=226444113 ): für externe Festplatten, CD-Laufwerke, *\*räusper\** Mini-Raketenwerfer *\*räusper\**, etc.
    - [SATA (Serial AT Attachment)](https://de.wikipedia.org/w/index.php?title=Serial_ATA&oldid=220865700 ): für interne Festplatten, kabelgebundene Punkt-zu-Punkt-Verbindungen
    - [SAS (Serial Attached SCSI)](https://de.wikipedia.org/w/index.php?title=Serial_Attached_SCSI&oldid=218079361 ): für interne Festplatten, kabelgebundene Verbindungen auch mit Switching (mehrere Festplatten an einem Verteiler)
    - [NVMe (Non-Volatile Memory Express)](https://de.wikipedia.org/w/index.php?title=NVM_Express&oldid=222566309 ): für interne Festplatten, direkte Steckverbindung via PCI Express
    - [Warum alles immer seriell?](https://web.archive.org/web/20220929192234/https://old.reddit.com/r/explainlikeimfive/comments/5sgfzh/eli5_parallel_was_faster_than_serial_why_isnt/ ) Höhere Geschwindigkeiten = Synchronisierung immer schwieriger. Parallele Leitungen = Übersprechen/Interferenz = komplexe Kompensation notwendig bzw. bei flexiblen Systemen wie Kabeln nicht möglich.

- Rückbezug zu STP012: "Dateisysteme sind eine besondere Form von Key-Value-Datenbank"; keine Datenbank im engeren Sinne:
    - keine widerspruchsfreie Speicherung, da keine Beziehungen zwischen verschiedenen Dateien möglich
    - keine Unterstützung für bedarfsgerechte Darstellungsformen; jede Datei hat genau ein Format

- alternative Perspektive als Vergleich zum blockbasierten Speicher
    - jede Datei ist ihr eigener "kleiner Massenspeicher"
    - zusätzlich noch Metadaten (Dateiname, Erstelldatum, Zugriffsrechte, etc.)
    - Dateisystem kümmert sich um die Zuteilung des tatsächlichen blockbasierten Speichers an die einzelnen Dateien
    - Praxisbeispiel zur Zuteilung von Speicherplatz: Unix-Befehl `du -h` ("disk usage")

- Dateisystem-Struktur unter unixoiden Betriebssystemen: mittels [Inodes](https://de.wikipedia.org/w/index.php?title=Inode&oldid=226067530 )
    - ein Inode enthält alle Metadaten zu einem bestimmten Dateisystemeintrag (einer Datei, einem Verzeichnis, einer Verknüpfung, etc.) sowie (bei normalen Dateien) einen Verweis auf den Speicherort der tatsächlichen Daten
    - Größe eines Inodes: immer ein Vielfaches der Blockgröße (4 KiB, 8 KiB, etc.)
    - bei sehr kleinen Dateien mitunter direktes Ablegen des Dateiinhaltes im Inode
    - bei großen Dateien meist fragmentierter Inhalt: Ablage in mehreren separaten Stücken (wegen späterer Umschreibung oder wegen Platzmangel)
    - Verzeichnisse: Inode mit einer Liste von Einträgen (Name + Inode-Verweis)
    - Dateisystem = Inode-Liste + Dateiinhalte + Übersicht über noch freie Festplattenbereiche

- "aktuelle" Entwicklungen in Dateisystemen
    - [Journaling](https://de.wikipedia.org/w/index.php?title=Journaling-Dateisystem&oldid=214907336 ) (in NTFS seit 1993, in ext3 seit 2001): Transaktionen werden in einem separaten Festplattenbereich vorgemerkt, bevor sie tatsächlich ausgeführt werden; ermöglicht Abgeschlossenheit und Dauerhaftigkeit analog zum Write-Ahead-Log bei Datenbanken (siehe STP012), außerdem schnelleren Start nach Systemausfall (keine Komplettprüfung des Dateisystems nötig)
    - [Online-Defragmentierung](https://de.wikipedia.org/w/index.php?title=Fragmentierung_(Dateisystem)&oldid=201408972 ) bzw. optimierte Speicherzuteilung zur Vermeidung von Fragmentierung (aber: aktives Defragmentieren = mehr Schreibzyklen = mehr Abnutzung)

- antiintuitive Schlussfolgerungen aus dieser Struktur
    - "sicheres Löschen": einfaches Löschen einer Datei löscht nur den Inode, aber überschreibt nicht den Inhalt (selbst Überschreiben der Datei hilft nicht unbedingt, weil die neuen Daten vielleicht an einer anderen Stelle landen, oder weil die alten Daten trotzdem noch im Journal liegen)
    - Löschen von im Verbrauch befindlichen Dateien: kein Problem, eine "geöffnete Datei" ist nur ein Verweis auf einen Inode; "Löschen der Datei" löscht nur den Verzeichnis-Eintrag; der Inode bleibt noch so lang bestehen, wie ein Prozess einen Verweis auf den Inode hält
    - anders bei Windows: Verweis auf eine offene Datei blockiert das Löschen -> Konsequenzen für System-Updates

- Exkurs 1: virtuelles Dateisystem
    - bis hier: physikalische Perspektive (Welche Bits stehen tatsächlich auf der Platte?)
    - jetzt: logische Perspektive (Was zeigt das Betriebssystem den Programmen?)
    - Unix: ein einziger Verzeichnisbaum, in das die verschiedenen Dateisysteme [eingehängt](https://de.wikipedia.org/w/index.php?title=Mounten&oldid=199499166 ) werden können; anders als das System mit Laufwerksbuchstaben unter Windows
    - "Datei" und "Verzeichnis" als abstrakte Konzepte, die u.a. durch echte Dateisysteme abgebildet werden können, oder auch durch was anderes (vgl. ["Everything is a file."](https://de.wikipedia.org/w/index.php?title=Everything_is_a_file&oldid=223250415 ))
    - Beispiel: [procfs](https://de.wikipedia.org/w/index.php?title=Procfs&oldid=222577559 )
    - Beispiel: [sshfs](https://de.wikipedia.org/w/index.php?title=SSHFS&oldid=223439141 )
    - Beispiel: transparente Behandlung von Zip-Archiven im Windows Explorer
    - Beispiel: [/dev/random](https://de.wikipedia.org/w/index.php?title=/dev/random&oldid=216549123 ) etc.

- Exkurs 2: Festplattenverschlüsselung
    - Option 1: auf der Ebene des Dateisystems (einzelne Dateien oder Verzeichnisse werden verschlüsselt)
    - Option 2: auf der Ebene des blockbasierten Speichers -> Dateisystem sieht als Speichermedium einen _virtuellen_ blockbasierten Speicher, der die Verschlüsselung vornimmt und erst dann auf den tatsächlichen blockbasierten Speicher schreibt -> schöne Arbeitsteilung
    - andere Anwendungsfälle für virtuelle blockbasierte Speicher: Netzwerklaufwerke, Dateisysteme innerhalb von Dateien
