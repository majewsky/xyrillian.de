- Geschichte des [Internets](https://de.wikipedia.org/wiki/Internet )
    - 1989-1991 initiale Entwicklung des World Wide Web
    - Internet aber bereits seit 1969 als Darpanet, ab 1981 bis 1983 Umstellung auf TCP/IP
    - über 20 Jahre Internet ohne Web -> das Web ist nur eine von vielen Anwendungen für das Internet
    - Internet vs. Web: Netzwerk von Computern vs. Netzwerk von Dokumenten

- DNS: seit 1984, ausführlich siehe STP018

- [E-Mail](https://de.wikipedia.org/wiki/E-Mail ): seit 1971
    - Ablage von Mails in Mailboxen mittels [SMTP (Simple Mail Transfer Protocol)](https://de.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol )
    - Abruf von Mails aus Mailboxen zuerst durch direkten Login auf dem Zielsystem, ab 1984 mittels [POP3 (Post Office Protocol)](https://de.wikipedia.org/wiki/Post_Office_Protocol ), ab 1994 mittels [IMAP (Internet Message Access Protocol)](https://de.wikipedia.org/wiki/Internet_Message_Access_Protocol )
    - nicht dasselbe wie das Web: asynchrone Zustellung (siehe [Wikipedia zur ersten E-Mail nach Deutschland](https://de.wikipedia.org/wiki/E-Mail#Geschichte ))

- [NTP (Network Time Protocol)](https://de.wikipedia.org/wiki/Network_Time_Protocol ): seit 1985
    - Synchronisation der Computeruhr mit der Uhr eines Servers und damit indirekt mit einer Atomuhr
    - Zeitdifferenzen werden (bei Nachgehen) durch Vorstellen oder (bei Vorgehen) durch Verlangsamen der lokalen Uhr korrigiert
    - nicht dasselbe wie das Web: nicht dokumentenbasiert

- [Telnet](https://de.wikipedia.org/wiki/Telnet ), [RSH (Remote Shell)](https://de.wikipedia.org/wiki/Remote_Shell ) und [SSH (Secure Shell)](https://de.wikipedia.org/wiki/Secure_Shell ): seit 1974/1977/1995
    - Telnet: Klartextaustausch mit einem entfernten Rechner
    - RSH/SSH: Abwicklung einer Kommandozeilen-Sitzung auf dem entfernten Rechner
    - nicht dasselbe wie das Web: bidirektionaler Datenstrom

- [IRC (Internet Relay Chat)](https://en.wikipedia.org/wiki/Internet_Relay_Chat ), [XMPP (Extensible Messaging and Presence Protocol)](https://de.wikipedia.org/wiki/Extensible_Messaging_and_Presence_Protocol ): seit 1988/1998
    - IRC: textbasiertes Chatprotokoll zwischen vielen nutzerseitigen Clients und einigen untereinander vernetzten Servern
    - XMPP: ähnlich, statt Plain Text mittels Strömen von XML-Nachrichten
    - nicht dasselbe wie das Web: ebenfalls ein kontinuierlicher bidirektionaler Datenstrom

- [SIP (Session Initiation Protocol)](https://de.wikipedia.org/wiki/Session_Initiation_Protocol ): seit 1999
    - und darauf folgend ein ganzer Zoo von Telefonie-Protokollen... ITU-Standards sind unser blinder Fleck :)
    - nicht dasselbe wie das Web: ebenfalls ein kontinuierlicher bidirektionaler Datenstrom

- [FTP (File Transfer Protocol)](https://de.wikipedia.org/wiki/File_Transfer_Protocol ), [NFS (Network File System)](https://de.wikipedia.org/wiki/Network_File_System ), [SMB (Server Message Block)](https://de.wikipedia.org/wiki/Server_Message_Block ): seit 1985/1984/1996
    - Zugriff (FTP) bzw. Einbindung (NFS/SMB) eines entfernten Dateisystem auf dem lokalen Rechner
    - Protokollbefehle wie "Dateien in Ordner auflisten" oder "Dateiinhalt lesen" oder "Datei löschen"
    - nicht dasselbe wie das Web: keine Dokumente mit Verweise untereinander, "nur" eine Struktur aus Dateien ohne bestimmte Anforderungen an den Dateiinhalt

- [LDAP (Lightweight Directory Access Protocol)](https://de.wikipedia.org/wiki/Lightweight_Directory_Access_Protocol ): seit 1993
    - Wenn DNS das Telefonbuch des Internets ist, dann sollte LDAP das Benutzerverzeichnis des Internets sein.
    - nicht dasselbe wie das Web: nicht dokumentenbasiert, sondern datensatzbasiert

- und dann natürlich [HTTP (Hypertext Transfer Protocol)](https://de.wikipedia.org/wiki/Hypertext_Transfer_Protocol ): seit 1991
    - zunächst fokussiert auf Abruf von Hypertext-Dokumenten in HTML sowie der darin enthaltenen Ressourcen (zunächst vor allem Bilder)
    - dann erweitert um Hochladen von Daten zum Ausfüllen von Formularen
    - dann erweitert um Verwaltung von Sitzungsdaten mittels Cookies
    - dann erweitert um Methoden zum Durchlaufen von Dateisystemen (WebDAV, CalDAV, CardDAV, GroupDAV)
    - dann erweitert durch eine komplette Client-Server-Programmierumgebung (JavaScript, XHR)
    - dann erweitert durch serverseitige Benachrichtigungen (WebSockets, Server Sent Events)
    - dann erweitert... und erweitert... und erweitert

- Plot-Twist: Web ist nur ein Teil des Internets, aber [alle anderen Teile des Internets werden immer webbiger](https://twitter.com/fluepke/status/1504188393378430988 )
    - jedes einzelne Protokoll profitiert tendenziell davon, wenn es auf HTTP aufsetzt:
        - Webbrowser können direkt HTTP verstehen
        - die meisten Entwickler kennen HTTP gut (auf jeden Fall besser als die ganzen anderen Protokolle)
        - in HTTP sind häufige Probleme wie Verschlüsselung/Caching/etc. schon gelöst (und die Lösungen sind milliardenfach getestet)
        - HTTP kommt am besten durch Firewalls durch
    - quasi ein Netzwerkeffekt ("es ist populär, weil es alle benutzen, und alle benutzen es, weil es populär ist")

- Wie sieht es mit den vorgenannten Protokollen aus?
    - DNS -> DoH (DNS over HTTPS), um Verschlüsselung hinzuzufügen
    - IRC, XMPP -> meist HTTP-basierte app-spezifische Chatprotokolle (z.B. Matrix)
    - LDAP -> [SAML](https://en.wikipedia.org/wiki/Security_Assertion_Markup_Language ) (Authentifizierung mittels Umleitungskette im Webbrowser)
    - SMTP, SSH, SIP, NFS, SMB: gibt's so noch, aber als Webapplikation ist die letzte Meile über HTTPS oder das eigentliche Protokoll ist in HTTPS-Websockets eingewickelt
