* Schritt 1: Webseitenadresse, z.B. <https://de.wikipedia.org/wiki/Computer>
    * `https`: Protokoll (wie wir mit der anderen Seite reden)
    * `de.wikipedia.org`: Domain-Name (wo die andere Seite zu finden ist)
    * `/wiki/Computer`: Pfad (welches Dokument wir von der anderen Seite haben wollen)

* Schritt 2: Namensauflösung
    * Datenverbindungen im Internet gehen nur mit einer IP-Adresse (sieht aus wie `192.168.12.169` oder `fe80::42e2:30ff:fe12:17b`)
    * Webseitenadressen enthalten Domain-Namen, weil:
        * menschenlesbar
        * kann auf mehrere verschiedene Server verweisen (Lastverteilung, Redundanz, Geo-Routing)
    * Frage: Welche IP-Adresse gehört zu dem Domain-Namen? -> Namensauflösung mittels Domain Name System (DNS)
    * Computer der Benutzerin kennt einen DNS-Server (meist vom Internetanbieter)
    * DNS-Server kann zu jedem Domain-Namen die richtige IP-Adresse bestimmen
    * Privatsphäre-Relevanz: dein DNS-Server kennt alle Domains, die du besuchst

* Schritt 3: Verbindung zum Webserver
    * Broadcast vs. Unicast
        * Fernsehen/Radio/etc. ist Broadcast (eine sendet einmal, viele empfangen)
        * Internet ist im Allgemeinen Unicast (eine sendet, eine empfängt) wie das Telefonnetz
    * grobe Analogie: IP-Adressen zu Telefonnummern (Ländervorwahl, Ortsvorwahl, Durchwahl, etc.)
    * Datenleitung steht, jetzt können wir unser Protokoll (HTTPS) darüber sprechen
        * Wikipedia: "Ein Kommunikationsprotokoll ist eine Vereinbarung, nach der die Datenübertragung zwischen zwei oder mehreren Parteien abläuft. In seiner einfachsten Form kann ein Protokoll definiert werden als eine Menge von Regeln, die Syntax, Semantik und Synchronisation der Kommunikation bestimmen."
    * HTTPS verwendet Verschlüsselung, da sonst jeder entlang des Weges mitlesen könnte
        * TLS: Aushandlung einer Transportverschlüsselung, dann Überprüfung der Identität des Servers
        * Server muss dazu ein Zertifikat für seinen Domain-Namen haben
    * Datenleitung mit Verschlüsselung steht, jetzt kann HTTP gesprochen werden (HTTPS = HTTP mit TLS)
        * HTTP: allgemein Protokoll zum Herunterladen und Hochladen von Dateien (Hochladen z.B. auch beim Abschicken von Webseiten-Formularen)

* Schritt 4: Webseite
    * Computer der Benutzerin fragt über die bestehende HTTPS-Verbindung nach dem Dokument unter dem Pfad in der Webseitenadresse
    * Server liefert eine Datei, üblicherweise im HTML-Format
    * Hausaufgabe: im Browser Webseite öffnen und im Rechtsklick-Menü "Element untersuchen" wählen
    * HTML: Dokumentenformat
        * dieselbe Grundkategorie wie PDF oder DOC
        * aber mehr für die Bildschirmanzeige als für den Druck optimiert
        * seit 2000 zunehmend auch für dynamische Inhalte ("Webanwendungen")

* Schritt 5: Ressourcen
    * HTML alleine enthält nur die Dokumentenstruktur und den Text
    * weitere Teile der Webseite liegen in anderen Dateien, auf die das HTML verweist und die der Browser separat herunterlädt
        * CSS: Darstellungsanweisungen (z.B. "alle Absätze mit 1,5 Zeilen Zeilenhöhe")
        * Bilder, Audios, Videos
        * Skripte für dynamische Inhalte
    * siehe auch: Webseite wird schon angezeigt, aber Bilder erscheinen erst nach und nach
    * Verweise auch auf Dokumente auf anderen Domains möglich
        * Privatsphäre-Relevanz: Tracking-Pixel und Tracking-Skripte

* Schritt 6: Cookies
    * HTTP an sich ist zustandslos: Anfrage nach einem Dokument ist nicht abhängig davon, was ich vorher angefragt hatte
    * Problem: Wie ordnet man die Anfrage nach einem Dokument zu einer Benutzerin zu? (z.B. um Anmeldestatus oder Warenkorb zu verfolgen)
    * Cookie: kleine Datei, die der Server bei einem Dokument mitschickt und die vom Browser aufbewahrt wird
    * beim nächsten Request werden alle Cookies von diesem Server wieder mitgeschickt
    * Privatsphäre-Relevanz: Tracking durch Facebook-Cookie etc.

* Exkurs: "Was passiert, wenn ich einen Podcast abrufe?"
    * prinzipiell genauso, aber statt HTML anderes Dokumentenformat
    * RSS: "Really Simple Syndication" ("wirklich einfache Verbreitung")
    * RSS-Dokument ist ein Feed
        * Feed = Liste von Nachrichten
        * vergleichbar mit dem Feed auf Twitter oder Facebook
    * RSS ursprünglich für Blogs und Nachrichtenwebseiten entwickelt (Titel, Kurzbeschreibung, Link zur Webseite)
    * für Podcasts statt Link zur Webseite Links zu den Audio-Dateien
