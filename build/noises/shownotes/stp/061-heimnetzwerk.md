- Rückbezug: STP005 (Netzwerk-Schichtenmodell)
    - seither haben wir über die unteren Schichten des Netzwerks nicht allzu viel gesprochen, sondern eher über Anwendungen darauf (siehe STP018 zu DNS, STP022 zu dem Web, STP052/STP054 zu TLS, STP057 zu ActivityPub/Mastodon)
    - heute mal wieder mehr die Infrastruktursicht (Layer 1 bis 4), vor allem aus praktischer Endanwendersicht

- Layer 1/2 kabelgebunden: über achtadrige [Twisted-Pair-Kabel](https://de.wikipedia.org/w/index.php?title=Twisted-Pair-Kabel&oldid=243787992) mit üblicherweise vorkonfektionierten [RJ45-Steckern (8P8C)](https://de.wikipedia.org/w/index.php?title=RJ-Steckverbindung&oldid=245704057) nach [IEEE 802.3 "Ethernet"](https://de.wikipedia.org/w/index.php?title=Ethernet&oldid=244710581)
    - Twisted Pair: Kabel mit verdrillten Doppeladern und einer (je nach Kategorie) zunehmenden Menge an Abschirmungen, damit das Kabel nicht z.B. unbeabsichtigt zur Antenne wird :) -- Wikipedia hat eine [schöne Visualisierung](https://commons.wikimedia.org/w/index.php?title=File:Twisted_pair_based_ethernet.svg&oldid=507351021), wie besser geschirmte Kabel für mehr Bandbreite sorgen
    - vorkonfektionierte Stecker: man kann auch selber konfektionieren (sprich: Kabel auf die richtige Länge zuschneiden und Modularstecker anbringen)
    - im Rechenzentrum findet man auch Ethernet, aber oberhalb der 10 Gbit/s dann nur noch auf Basis von Glasfasern und [SFP-Modulen](https://de.wikipedia.org/w/index.php?title=Small_Form-factor_Pluggable&oldid=244881738)
    - Kabel alleine reichen nicht, man braucht mindestens noch einen [Switch](https://de.wikipedia.org/w/index.php?title=Switch_(Netzwerktechnik)&oldid=238943317), um tatsächlich ein Kabelnetzwerk zu erhalten
    - der Switch ist oft auch ein [Router](https://de.wikipedia.org/w/index.php?title=Router&oldid=242713208), der das lokale Kabelnetzwerk mit dem Internet verbindet
    - Wir empfehlen: Beim Hausbau oder -umbau am besten gleich Cat7 verlegen (oder zumindest Leerrohre), damit man sich nicht mit WLAN-Signalstärken rumärgern muss.
        - im Gespräch erwähnt: [Keystone-Modul](https://de.wikipedia.org/w/index.php?title=Keystone-Modul&oldid=234182002)

- Layer 1/2 kabellos: über Radiowellen im 2,4-GHz-Band oder im 5-GHz-Band nach [IEEE 802.11 "WLAN bzw. Wi-Fi"](https://de.wikipedia.org/w/index.php?title=IEEE_802.11&oldid=245761310)
    - was bei Ethernet der Switch, ist bei WLAN der [Access Point (AP)](https://de.wikipedia.org/w/index.php?title=Wireless_Access_Point&oldid=237904996)
    - übliches Setup zuhause: der zentrale Router macht sowohl LAN (Ethernet) als auch WLAN
    - Xyrills Setup: ein Server agiert als Router, aber macht nur LAN; am LAN hängt ein WLAN-AP (**kein** WLAN-Router, somit nur ein lokales Netz für alle Endgeräte)

- Layer 3 mit [IPv4](https://de.wikipedia.org/w/index.php?title=IPv4&oldid=243903744) (1981)
    - auf dem Router läuft mindestens ein [DHCP-Server](https://de.wikipedia.org/w/index.php?title=Dynamic_Host_Configuration_Protocol&oldid=244947062), um die lokalen Adressen zu vergeben; oft auch ein [DNS-Resolver](https://de.wikipedia.org/w/index.php?title=Domain_Name_System&oldid=245211339) (siehe STP018)
    - an und für sich sind Netzwerke ineinander verschachtelt, z.B. kann ein Netzwerk `10.11.0.0/16` ein Teilnetzwerk `10.11.12.0/24` enthalten (Analogie: Telefonanlage in einem Firmengebäude mit Verbindung ins weite Telefonnetz); damit fast beliebig feine Unterteilungen des Netzwerks möglich
    - Problem: insgesamt etwa 4,3 Mrd. unterschiedliche IP-Adressen, aber davon nur [etwa 3,7 Mrd. "weltweit gültige" Adressen](https://stackoverflow.com/a/2437185)
    - in der Praxis sehr viel [NAT](https://de.wikipedia.org/w/index.php?title=Netzwerkadress%C3%BCbersetzung&oldid=237026399): ein Heimnetzwerk kriegt nur eine IP-Adresse für den Router; innerhalb des lokalen Netzwerkes private Adressen (z.B. `192.168.x.y`), die nicht direkt aus dem Internet erreichbar sind

- Layer 3 mit [IPv6](https://de.wikipedia.org/w/index.php?title=IPv6&oldid=245056053) (1998)
    - IPv6-Adressen haben sehr viel mehr Ziffern und sind damit extrem viel zahlreicher, aber auch unhandlicher
    - noch sehr viele andere Unterschiede zu IPv4, die mit dazu führten, dass sich IPv6 nur sehr langsam verbreitet hat bzw. immer noch verbreitet
    - Endanwender kriegen bei IPv6 nicht nur einzelne Adressen, sondern wie ursprünglich vorgesehen ganze Netze, die ihr Router frei verwalten kann
    - statt DHCP meist nur SLAAC (Stateless Address Auto-Configuration; zustandslose Adressen-Autokonfiguration): Router teilt den Netzwerkbereich mit und Endgeräte würfeln selbsttätig gültige Adressen
    - Umfangreichere Infos im [RFC-Podcast zu IPv6](https://requestforcomments.de/archives/412)
    - Problem: Wenn jedes Gerät vom Computer bis zur smarten Glühbirne eine global eindeutige IPv6-Adresse hat, ist das nicht ein Traum für jeden Überwachungskapitalisten?

- Layer 3 mit [VPN (Virtual Private Network)](https://de.wikipedia.org/w/index.php?title=Virtual_Private_Network&oldid=245792132)
    - Idee: Netzwerkverkehr ist zumindest auf Layer 2 und darüber auch nur Bits und Bytes; die kann man als Nutzdaten behandeln und in eine andere Netzwerkverbindung zu einem VPN-Server einpacken (meist auch mit Verschlüsselung drumherum, z.B. TLS)
    - Xyrills Setup: ein Server im Rechenzentrum fungiert als VPN-Server, mit dem Xyrills Endgeräte alle verbunden sind; aus Sicht der Endgeräte sind alle im selben lokalen Netzwerk; damit auch z.B. vom Notebook von unterwegs "direkter" Zugriff auf den Heimserver möglich (trotz NAT im Plaste-Router)
    - das häufigere Setup beim Endanwender: aller Datenverkehr nach außen geht durch einen VPN-Tunnel und erst ab dem Server des VPN-Anbieters weiter ins weite Internet
    - tatsächlicher Nutzen: es sieht so aus, als ob man in einer anderen geografischen Region sitzt, was z.B. für Streaming-Dienste relevant sein kann (aber Vorsicht: das verstößt oft gegen die Geschäftsbedingungen des Streamers; das hier ist keine Rechtsberatung)
    - fragwürdiger Nutzen: man muss nicht mehr dem Internetanbieter vertrauen, dass der nicht den eigenen Datenverkehr mitliest (Problem 1: in Zeiten von allgegenwärtigem TLS sieht der ISP sowieso nur die Domain-Namen und IP-Adressen, nicht die Inhalte; Problem 2: man verschiebt die Vertrauensfrage nur vom Internetanbieter zum VPN-Anbieter)

- Layer 4 bis 7 dann beim nächsten Mal

