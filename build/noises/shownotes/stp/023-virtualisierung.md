Laut Wikipedia:

> [Virtualisierung](https://de.wikipedia.org/wiki/Virtualisierung_(Informatik) ) bezeichnet in der Informatik die Nachbildung eines Hardware- oder Software-Objektes durch ein ähnliches Objekt vom selben Typ mithilfe einer Abstraktionsschicht.

* verschiedene Arten von Virtualisierung je nachdem, wo die Abstraktionsschicht liegt
    * Virtuelle Maschinen (VM): Hardware gleich, Betriebssysteme (OS) und Prozesse getrennt
    * Container: OS gleich, Prozesse getrennt
    * Sandboxing: getrennte Teile innerhalb eines Prozesses

* Anwendungsfälle für VM (OS als Anwendung innerhalb eines anderen OS)
    * gleichzeitiger Betrieb von mehreren OS auf einem Computer (z.B. Linux als Hauptsystem, Windows-VM für bestimmte Anwendungen)
    * saubere Abgrenzung von verschiedenen Programmen durch das Einsperren in VMs mit festen Ressourcen-Zuteilungen und Zugriffsschranken
    * daraus folgend: Aufteilen einer einzelnen Server-Maschine in mehrere VMs für verschiedene Kunden
    * Ermöglichung der Verschiebung von Programmen zwischen Maschinen im laufenden Betrieb ("Live-Migration", z.B. bei Hardwarefehlern)

* Ansätze für VMs (Wie kann der Gast-Kernel seinen administrativen Aufgaben nachgehen, ohne selber vollen Zugriff auf das System zu erhalten?) anhand von Meilensteinen auf x86-CPUs
    * User Mode Linux (ab ca. 2000), Xen (ab 2003): "Paravirtualisierung", Gast-Kernel ist von Anfang an so aufgebaut, dass er auch als Prozess unter einem anderen Kernel laufen kann (Hardware-Zugriff mittels Emulation, siehe unten)
    * VMware Workstation (ab 1999), VirtualBox (ab 2007), etc.: "Software-Virtualisierung", Gast-Kernel läuft in einem teilprivilegierten Modus (z.B. Ring 1 bei x86), sodass Hardware-Zugriff einen Fehler auslöst; der Host-Kernel erhält die Kontrolle zurück und macht den Hardwarezugriff im Namen des Gast-Kernels (Problem: Performance-Verlust gegenüber Direktzugriff auf die Hardware) ([detaillierte Erklärung in der VirtualBox-Dokumentation](https://docs.oracle.com/en/virtualization/virtualbox/6.0/admin/swvirt-details.html ))1
    * Intel VT-x (ab 2005), AMD-V (ab 2006): "Hardware-Virtualisierung", Host-Kernel richtet durch spezialisierte CPU-Instruktionen einen abgetrennten Bereich für den Gast-Kernel ein (beschränkter Speicherbereich, beschränkte CPU-Zeit, beschränkter Gerätezugriff), innerhalb dessen das System für den Gast-Kernel wie unvirtualisiert erscheint (somit so gut wie kein Performance-Unterschied zwischen nativem und virtualisiertem Kernel mehr)

* [Container](https://de.wikipedia.org/wiki/Containervirtualisierung ): isolierte Umgebungen innerhalb eines OS, aus denen Programme im Allgemeinen nicht ohne weiteres ausbrechen können
    * Vorteil gegenüber VM: ressourcenschonender (nur ein Kernel für alle), Container starten schneller als VM (kein Boot eines Gast-Kernels notwendig)
    * Urform des Containers: [chroot in Unix-Systemen](https://de.wikipedia.org/wiki/Chroot ) (ab 1979) modifiziert für einen Prozess das sichtbare Dateisystem so, dass nur Dateien unterhalb eines bestimmten Verzeichnisses erreicht werden können (Intention war die Bereitstellung von angepassten Dateisystemstrukturen für bestimmte Programme; chroot ist kein wirklicher Container, da die allermeisten Ressourcen nicht eingeschränkt werden und ein privilegierter Prozess das chroot jederzeit durch Neukonfiguration umgehen kann)
    * echte Container ab ca. 2000: z.B. FreeBSD Jails ab 1999, Virtuozzo/OpenVZ ab 2000, Solaris Zones ab 2005, LXC (Linux Containers) ab 2008, Docker ab 2013

* ausführliche Besprechung von Containern unter Linux (Docker/Kubernetes) siehe [Pentaradio von 2022-05-24](https://c3d2.de/news/pentaradio24-20220524.html )

* Sandboxing: Container innerhalb einer Applikation
    * z.B. Webbrowser führen JavaScript-Code aus dem Internet in separaten Prozessen mit geringen Privilegien aus, um Malware das Leben schwer zu machen
    * Begriffstrennung nicht ganz sauber; mitunter werden auch Container als eine Art von Sandboxing eingeordnet

* Simulation und Emulation
    * [Simulation](https://de.wikipedia.org/wiki/Simulation ): Nachbildung eines realen Systems oder Szenarios mittels andersgearteter Technik (z.B. Fahrsimulation ohne echtes Auto, Militärsimulation ohne echtes Schlachtfeld, Simulation von quantenmechanischen Wellen mittels Ultraschallwellen... oder gleich mit Zahlen im Computer)
    * [Emulation](https://de.wikipedia.org/wiki/Emulator ): Nachbildung eines Systems durch ein ähnliches, aber anderes System (z.B. Nachbildung des Verhaltens eines CPU-Typs als Programm auf einem anderen CPU-Typ)
    * Einsatz von Emulation hauptsächlich zum Nachbilden alter Systeme (siehe Emulation von Videospielen, wie besprochen im [Pentaradio vom Januar 2022](https://c3d2.de/news/pentaradio24-20220125.html )) und zum Testen von Software auf besser zugänglichen Systemen (z.B. Testumgebungen für Smartphone-Apps auf PC-Betriebssystemen)
