* Vorbemerkung: Daten im Computer bzw. im Computernetzwerk sind immer nur Zahlen
    * Wenn wir Texte haben, haben wir tatsächlich Zahlenfolgen, von denen jede Zahl für einen Buchstaben steht.
    * Wenn wir Bilder haben, haben wir tatsächlich Zahlenfolgen, von denen jede Zahl für den Farbwert eines Bildpunktes steht.
    * etc.

* Rückerinnerung zu STP002: Protokolle
    * Wikipedia: "Ein Kommunikationsprotokoll ist eine Vereinbarung, nach der die Datenübertragung zwischen zwei oder mehreren Parteien abläuft. In seiner einfachsten Form kann ein Protokoll definiert werden als eine Menge von Regeln, die Syntax, Semantik und Synchronisation der Kommunikation bestimmen."
    * Beispiel aus dem normalen Leben: Interaktion zwischen Kunde und Verkäufer

* Computer-Netzwerke brauchen einen ganzen Stapel von Protokollen, die aufeinander aufbauen
    * Problemlösung auf verschiedenen Ebenen: physikalische Signalübertragung, Suche von Verbindungspfaden, Fehlerkorrektur, finale Anwendung
    * Nachrichten einer höheren Schicht werden in Nachrichten der nächsttieferen Schicht eingewickelt
        * metaphorisches Beispiel: verpacktes Geschenk in Postpaket in Postwertsack in Cargo-Flugzeug
    * Austauschbarkeit einzelner Teile des Stapels: z.B. höhere Protokollebenen hängen (fast) nicht davon ab, ob die Übertragung mittels Ethernet oder DSL oder WLAN oder LTE etc. läuft
        * innerhalb der Metapher: Paket kann sowohl über Luftpost als auch per Schiff zugestellt werden
    * Lehrmodell: [OSI-Modell](https://de.wikipedia.org/wiki/OSI-Modell ) (Open Systems Interconnection) mit 7 Schichten; [entspricht aber nicht ganz der Realität](https://computer.rip/2021-03-27-the-actual-osi-model.html )

* Schicht 1: Bitübertragungsschicht ("Physical Layer")
    * bei kabelgebundenen Wegen Spezifikation der Kabel (Aderanzahl, Materialien, Schirmung, etc.) und Stecker
    * bei kabellosen Wegen Spezifikation des Frequenzspektrums etc.
    * auf dieser Ebene im Allgemeinen nur Folgen von Bits (Nullen und Einsen)
    * Problem: echte Welt ist fehlerbehaftet (Störsignale)

* Schicht 2: Sicherungsschicht ("Data Layer")
    * Realisierung einer fehlerfreien Übertragung zwischen direkt miteinander verbundenen Geräten
    * Bitfolgen werden zu *Frames* zusammengefasst und mit Prüfsummen versehen
        * zu Prüfsummen siehe STP004
    * MTU: *Maximum Transfer Unit* (maximale Nutzlast in einem Frame)
    * manchmal enthalten: Protokolle zum Neustarten fehlgeschlagener Übertragungen (z.B. WLAN), ansonsten nur Verwerfen kaputter Frames (z.B. Ethernet)
    * hier meist hohe Komplexität aufgrund strenger Anforderungen an Verarbeitungsgeschwindigkeit

* Schicht 3: Vermittlungsschicht ("Network Layer")
    * Schicht 1 und 2 war nur für Geräte, die direkt miteinander verbunden sind (bzw. maximal durch Switches und Repeater verbunden)
    * Schicht 3 verschaltet mehrere solche lokale Netzwerke miteinander
    * *Router*: Computer, der mit mehreren Netzwerken verbunden ist und dazwischen vermittelt
    * Internet ist paketbasiert (wie die Post), nicht verbindungsbasiert (wie das Telefonnetz)
    * auf Schicht 3 große Einigkeit in den Protokollen, um Interoperabilität zu ermöglichen
    * IP (*Internet Protocol*)
        * Zuweisung global gültiger Adressen ("IP-Adressen"), z.B. `91.198.174.192` für IP-Version 4 oder `2620:0:862:ed1a::1` für IP-Version 6
        * IP-Pakete können über Netzwerkgrenzen hinweg an den Adressaten zugestellt werden; enthalten Rückadresse für Antwortpakete
    * IP-Adressen nicht immer eindeutig auf einen Computer zugewiesen
        * eine global erreichbare Adresse für mehrere Computer, z.B. Anycast (vgl. Großkundenadressen bei der Post)
        * private Adressen für die Verwendung in lokalen Netzwerken (analog zu internen Adressen in der Hauspost einer einzelnen Firma)
    * ICMP (*Internet Control Message Protocol*)
        * z.B. Ping, Pfad-MTU-Aushandlung, Anzeige von nicht verfügbaren Empfängern
    * NAT (*Network Address Translation*)
        * Anwendungsfeld: ein ganzes lokales Netzwerk hat nur eine öffentliche IP-Adresse
        * Router übersetzt IP-Adressen beim Übergang aus dem lokalen Netzwerk oder ins lokale Netzwerk
    * weitere Protokolle insb. für die dynamische Verschaltung von Netzwerken, z.B. BGP (*Border Gateway Protocol*)

* Schicht 4: Transportschicht ("Transport Layer")
    * Erweiterung von Schicht 3 um Sitzungsverwaltung und Fehlerkorrektur (analog zum Verhältnis von Schicht 2 zu Schicht 1)
    * Problem: IP-Adressen identifizieren nur Computer, nicht die Programme auf dem Computer
    * getrennte Sitzungen mittels *Ports*
        * Verbindung eindeutig identifizierbar durch die Verbindung aus Quelladresse, Quellport, Zieladresse und Zielport
    * TCP (*Transmission Control Protocol*)
        * garantierte fehlerfreie Zustellung der Datenpakete jeder Verbindung in der jeweils richtigen Reihenfolge
        * z.B. für große Downloads und Uploads (HTTP)
    * Fehlerkorrektur in TCP
        * verpflichtendes Bestätigen erfolgreicher Übertragungen
        * Sender hält Daten in einem Sendefenster vor, bis der Empfang von der Gegenseite bestätigt wurde; analog Empfangsfenster beim Empfänger
        * Überfüllungskontrolle: Sendefenstergröße wird angepasst, wenn die Fehlerrate zunimmt (weil jemand auf dem Weg Pakete verwerfen muss)
        * für den Benutzer sichtbar als "TCP Slow Start" (Downloads fangen meist erst mit geringer Geschwindigkeit an und werden dann schneller, sobald die optimale Sendefenstergröße gefunden wurde)
    * UDP (*User Datagram Protocol*)
        * keine Fehlerkorrektur; im Prinzip nur IP + Ports
        * z.B. für Echtzeit-Anwendungen (Voice over IP), wo geringe Verzögerung wichtiger ist und der Verlust einzelner Pakete hingenommen werden kann

* Schicht 5-7: Anwendungsschicht ("Application Layer")
    * feinere Aufteilung nicht wirklich hilfreich, da die verschiedenen Aspekte zunehmend schwer voneinander zu trennen sind
    * Protokolle auf dieser Ebene: z.B. HTTP/HTTPS (Webseiten), SMTP/IMAP (E-Mail), DNS (Namensauflösung)

* Grenzen des Schichtenmodells
    * tiefere Schichten können durch höhere Schichten abgebildet werden; Beispiel: VPN-Protokolle (*Virtual Private Network*; Protokolle der Schicht 7, das einen Datenkanal der Schicht 3 darstellt)
    * Schicht 8? Schicht 9?


* Zum Weiterhören
    * der großartige Podcast [Request for Comments](https://requestforcomments.de/ ) (dessen Produzentin natürlich Baecker und nicht Bauer heißt; ttimeless nimmt alle  Schuld auf sich)
    * hier der erwähnte [RFC-Podcast über TCP](https://requestforcomments.de/archives/254 )
    * außerdem die angesprochene Episode vom [Chaosradio Express](https://cre.fm/ ): [CRE197 IPv6](https://cre.fm/cre197-ipv6 )

* Lesestoff ;)
    * [ADSL over wet string](https://www.revk.uk/2017/12/its-official-adsl-works-over-wet-string.html )
    * [Twelve Networking Truths](https://tools.ietf.org/html/rfc1925 )
