* Rückbezug STP061: wir haben unser Heimnetzwerk von Schicht 1 bis 3 aufgebaut
    * Schicht 3 hatte uns einen Zustellweg durchs Internet gegeben
    * Schicht 4 ergänzt Sitzungsverwaltung (z.B. Flusskontrolle und Fehlerkorrektur; hierfür siehe STP005) und Ports
    * Warum braucht man Ports? Gedankenexperiment:
        * Annahme: Datenpakete werden nur anhand von Quelladresse und Zieladresse zugestellt
        * Situation: zwei Prozesse auf demselben PC wollen mit `wikipedia.de` reden
        * Abschicken der Anfragen klappt eventuell noch, da gleiches Ziel
        * Problem: Antwortpaket kommt -> Welcher Prozess kriegt das Paket?
        * Technisch formuliert: Wie kann man über diese Netzwerkleitung mehrere unabhängige Verbindungen [multiplexen](https://de.wikipedia.org/w/index.php?title=Multiplexverfahren&oldid=243098154)?

* Ports: Identifikationsnummer für offene Netzwerkverbindungen; 16 Bit groß (0-65535)
    * vgl. Briefkästen/Wohnungsnummern in Mehrparteienhäusern
    * Datenpakete auf Schicht 4 werden anhand der Gesamtheit aus Quelladresse, Quellport, Zieladresse, Zielport einer Verbindung zugeordnet
    * damit eindeutige Endzustellung an einen Prozess möglich
    * auf Server-Seite im Allgemeinen "wohlbekannte" Ports (z.B. 25 für SMTP, 80 für HTTP, 443 für HTTPS), damit Clients ein definiertes Ziel für neue Verbindungen haben
    * auf Client-Seite für gewöhnlich automatisch vergebene Ports (meist zwischen 32768-65535)
    * Praxisbeispiel: im Browser eine Webseite aufmachen, dann in der Konsole aktive Verbindungen auflisten mit `ss -utpn -o state connected`

* Wie kann ich zuhause einen Server-Prozess betreiben?
    * Außenwelt muss meine IP-Adresse finden können und den entsprechenden Port erreichen können
    * Problem 1: IP-Adresse wird durch den Internetanbieter bestimmt und ändert sich -> [Dynamisches DNS](https://de.wikipedia.org/w/index.php?title=Dynamisches_DNS&oldid=237965669) (nicht zu verwechseln mit der Firma "DynDNS", die nur ein möglicher Anbieter ist)
    * Problem 2: Router hat für gewöhnlich einen Paketfilter ([Firewall](https://de.wikipedia.org/w/index.php?title=Firewall&oldid=244493713)), der nur ausgehende sowie bekannte Verbindungen durchlässt -> Ausnahmeregel ("Port Opening") muss konfiguriert werden; je nach Betriebssystem auch dasselbe auf dem Server selbst erforderlich
    * Problem 3: IPv4 mit NAT, lokaler Rechner mit Server-Prozess hat nur eine lokale IP-Adresse -> Portweiterleitung muss konfiguriert werden (damit dann meist auch eine Ausnahmeregel in der Firewall verbunden)

* Was kann ich zuhause betreiben?
    * E-Mail definitiv nicht, MTAs aus Endkunden-IP-Bereichen werden von den großen Betreibern als Spam-Schleudern interpretiert und blockiert
    * Videotelefonie auch eher nicht, weil da NAT-Durchdringung noch problematischer ist
    * Faustregel: fast alles mit HTTP sollte gehen, sofern man nicht durch den oft dünnen Upstream von Endkunden-Internetanschlüssen begrenzt ist
    * Xyrill betreibt (teils zuhause, teils auf gemieteten Servern im Rechenzentrum): [Gitea](https://de.wikipedia.org/w/index.php?title=Gitea&oldid=241399660), [Jitsi Meet](https://de.wikipedia.org/w/index.php?title=Jitsi&oldid=245822087), [Matrix](https://de.wikipedia.org/w/index.php?title=Matrix_(Kommunikationsprotokoll)&oldid=243306410), [NextCloud](https://de.wikipedia.org/w/index.php?title=Nextcloud&oldid=246251388), [nginx](https://de.wikipedia.org/w/index.php?title=Nginx&oldid=246103192) (siehe STP022), [Prometheus](https://de.wikipedia.org/w/index.php?title=Prometheus_(Software)&oldid=236686905) (siehe STP032)
    * andere Vorschläge: [Bitwarden](https://en.wikipedia.org/w/index.php?title=Bitwarden&oldid=1230273157), [Kodi](https://de.wikipedia.org/w/index.php?title=Kodi_(Software)&oldid=243970444), [Mastodon](https://de.wikipedia.org/w/index.php?title=Mastodon_(Soziales_Netzwerk)&oldid=245943178) oder andere ActivityPub-Server (siehe STP057), [Paperless-NGX](https://docs.paperless-ngx.com/), [Pi-Hole](https://de.wikipedia.org/w/index.php?title=Pi-hole&oldid=243236097)

* im Gespräch erwähnt
    * Pentaradio zu Quellcode-Verwaltung mit Git: <https://c3d2.de/news/pentaradio24-20210727.html>
    * Audio-Streaming-Software: <https://www.navidrome.org/>
