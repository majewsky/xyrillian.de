- **Automatisierung:** nach DIN V 19233 definiert als "das Ausrüsten einer Einrichtung, sodass sie ganz oder teilweise ohne Mitwirkung des Menschen bestimmungsgemäß arbeitet" ([Quelle: Wikipedia](https://de.wikipedia.org/w/index.php?title=Automatisierung&oldid=222194579#Arten ))
    - Was heißt "bestimmungsgemäß"? (siehe 33C3-Motto "Works for me"; siehe ebenfalls "Overly-Clever-Syndrom")
    - heute feiern wir Autokonfiguration, die tatsächlich funktioniert
    - **Autokonfiguration:** automatisierte Führung eines neu entstehenden Systems in einen arbeitsfähigen Zustand

- Rückbezug zu STP020 (Bootstrap): Hardware-Erkennung während (und nach) dem Bootvorgang
    - früher: Betriebssysteme spezifisch für einzelne Computer-Modellreihen mit gleichartiger Hardware
    - heute: Betriebssysteme für ganze Kategorien von Computern, die oftmals vom Anwender selbst aus Komponenten zusammengestellt werden können
    - während des Bootvorgangs Abfrage der vorliegenden Hardware-Komponenten durch das Betriebssystem (siehe Xyrills Hardware-Upgrade)
    - nach dem Bootvorgang teilweise auch Neuerkennung möglich (Hotswap, USB, etc.)

- [Ethernet](https://de.wikipedia.org/w/index.php?title=Ethernet&oldid=223634192 )
    - ausführliche Beschreibung siehe später; Grundprinzip: Netzwerk aus Switches, an denen die Endgeräte hängen
    - Autokonfiguration 1: früher separate Kabel für Verbindungen zu Switches (Patch-Kabel) und Direktverbindungen (Cross-Kabel), heute automatische Erkennung mittels [Auto MDI-X](https://de.wikipedia.org/w/index.php?title=Medium_Dependent_Interface&oldid=202484104#Auto_MDI-X )
        - ähnlich dazu automatische Orientierungs-Erkennung in USB-C-Steckverbindungen
    - Autokonfiguration 2: früher Switches nicht in Schleifen verschaltbar (Broadcast-Stürme z.B. wegen ARP und DHCP)
    - heute automatische Bildung eines [Spannbaumes](https://de.wikipedia.org/w/index.php?title=Spanning_Tree_Protocol&oldid=224259716 ) für eindeutiges Routing von Frames

- apropos Ethernet: [ARP (Address Resolution Protocol)](https://de.wikipedia.org/w/index.php?title=Address_Resolution_Protocol&oldid=224525810 )
    - Identifikation im Internet (Layer 3-7) durch IP-Adressen, aber in lokalen Netzen (Layer 1-2) durch adapterspezifische MAC-Adressen
    - IP-Adressen können Computer sich im Prinzip auch einfach selber geben, grundsätzlich keine zentrale Vergabe
    - ARP: Welche MAC-Adresse gehört zur gewünschten IP-Adresse?
    - MAC-Adressen und ARP gibt's auch im WiFi...

- apropos Ethernet und WiFi: [DHCP (Dynamic Host Configuration Protocol)](https://de.wikipedia.org/w/index.php?title=Dynamic_Host_Configuration_Protocol&oldid=224991751 )
    - grundsätzlich Selbstwahl der IP möglich, aber man muss im richtigen Netzbereich sein und es darf keine Kollisionen geben
    - meist bessere Lösung: automatische IP-Vergabe durch eine zentrale Stelle (Clients können IP-Adressen für eine gewisse Zeit pachten)

- apropos DHCP: [TFTP (Trivial File Transfer Protocol)](https://de.wikipedia.org/w/index.php?title=Trivial_File_Transfer_Protocol&oldid=224254220 ) -> per DHCP automatisches Ankündigen von verfügbaren Betriebssystem-Images, Abholen der Images über TFTP für Booten mit leerer Festplatte
    - siehe auch: [Installieren von macOS auf Apple Macs mit leerer Festplatte via Internet Recovery](https://discussions.apple.com/thread/7645024?answerId=7645024021#7645024021 )

- IPv6 (Router Advertisement etc.) [siehe RFC-Podcast](https://requestforcomments.de/archives/412 )

- Autokonfiguration im Cloud-Computing-Umfeld
    - virtuelle Maschinen in der Cloud werden aus Abbildern (Images) erzeugt und müssen sich dann an die Laufzeitumgebung (IP-Adressen, verfügbare Netzwerk-Festplatten, etc.) anpassen
    - von Cloud-Seite: Metadata-Service meist auf der magischen IP-Adresse `169.254.169.254` (siehe z.B. [Dokumentation bei OpenStack](https://docs.openstack.org/nova/yoga/user/metadata.html#metadata-openstack-format ))
    - von VM-Seite: Initialisierung mittels dieser Metadaten mit z.B. [cloud-init](https://cloudinit.readthedocs.io/en/latest/index.html )
