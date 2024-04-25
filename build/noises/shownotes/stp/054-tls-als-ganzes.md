Feedback per Mail:

> Bei ca. 8 Minuten sagst du, das ein 64-Bit-Prozessor eine Adressbreite von 64 Bit hat, 32 Bit etc. analog. Das stimmt so nicht. Die Zahl gibt die Datenbreite an, nicht die Adressbreite (und damit auch die Registerbreite). Ein aktueller 64 Bit amd64 Prozessor kann max. 48 Bit adressieren (siehe /proc/cpuinfo, bzw. die Länge deiner Adressen in der maps Ausgabe). Der 286er konnte 20 bit Adressbreite, obwohl es ein 16-Bit-Prozessor war &ndash; Ich weiss nicht, ob du die Rumfummelei mit den Segmentregistern noch machen musstest. ;) Bei den 32-Bit-Prozessoren hat es übereingestimmt.
>
> Ca 14 Minuten: Das Programm muss nicht komplett im Speicher liegen, sondern nur Teil, der abgearbeitet wird. Heute eventuell unüblich, aber früher war es durchaus so, das Programmteile während der Ausführung nachgeladen wurden und dann andere, gerade nicht benutzte Programmteile ersetzt haben.
>
> Bei den 8-Bit-Spielekonsolen teilweise mit per Software umschaltbaren Speicherseiten des Game-ROMs, als z.B. das von einem 64k-ROM immer nur 16k-Blöcke für den Prozessor sichtbar waren (und die gleichen Adressen benutzt haben) und per IO-Befehl die Speicherseiten vom Programm gewechselt werden konnten, da ansonsten der Adressraum nicht ausgereicht hätte.
>
> Stack: Soweit ich weiss, ist der Stack nicht wirklich hart beschränkt, sondern nur über das eingestellte ulimit. Bis zu diesem (änderbarem) Limit wird der dynamisch zur Laufzeit verwaltet und ggf. vergrößert.

Xyrill nimmt die Ausführungen wohlwollend zur Kenntnis, sieht sich jedoch nicht gezwungen, seine Position zu revidieren. Nun zum eigentlichen Thema:

- Wo kriegt man Zertifikate her?
    - Rückblick auf STP048: Zertifikate sind signiert von übergeordneten Zertifikaten (**Zertifikatsautoritäten**, CA), wodurch sich eine Kette bis hoch zu einem **Wurzelzertifikat** (Root-CA) aufbaut
    - Clientzertifikate werden durch eine Stelle ausgegeben, der der Server vertraut (z.B. in Unternehmensnetzwerken eine CA unter der Obhut der Unternehmens-IT)
    - für Serverzertifikate enthält der Browser bzw. das Betriebssystem eine Liste von vertrauenswürdigen Root-CAs -> Wie kann man eine Root-CA werden?

- [CA/Browser Forum](https://de.wikipedia.org/w/index.php?title=CA/Browser_Forum&oldid=235926243)
    - Verein von Browser-Herstellern und Root-CA-Betreibern zur Selbstregulierung der Anforderungen an Browser und Root-CAs
    - Root-CAs, die gegen die entsprechenden Sicherheits- und Transparenzregeln verstoßen, wird von den Browsern das Vertrauen entzogen, was deren Geschäftsgrundlage entzieht (siehe z.B. [diese tolle lange Liste von Vorfällen mit Symantec-CAs, die Mozilla dokumentiert hat](https://wiki.mozilla.org/CA/Symantec_Issues))

- eine sehr bekannte Root-CA: [ISRG (Internet Security Research Group)](https://www.abetterinternet.org/) unter dem Markennamen [Let's Encrypt](https://letsencrypt.org/)
    - Situation vor 2013: TLS-Zertifikate kosten Geld und werden im Prinzip nur von Firmenwebseiten gekauft (Teufelskreis: teuer -> kleine Kundschaft -> keine Skaleneffekte -> teuer)
    - 2013: Snowden-Enthüllungen ([Pentaradio berichtete neulich über das 10-jährige Jubiläum](https://c3d2.de/news/pentaradio24-20230725.html)) -> Erkenntnis: wir brauchen mehr Verschlüsselung im Internet
    - 2015: ISRG bietet mit Let's Encrypt automatisierbare und kostenlose TLS-Zertifikate für Jedermann

- Was tun nach einem IT-Angriff, wenn ein Server-Zertifikat kompromittiert wurde?
    - Zertifikate können der CA als gestohlen gemeldet werden
    - [Online Certificate Status Protocol (OCSP)](https://de.wikipedia.org/w/index.php?title=Online_Certificate_Status_Protocol&oldid=235873169): Protokoll für die Abfrage von Rückzugsanzeigen (dem digitalen Äquivalent von Diebstahlsmeldungen)
    - Probleme: wenn das die Clients machen, überlastet es die kleinen CAs und die großen CAs könnten damit wunderbar #Datenspuren sammeln
    - OCSP-Stapling ("OCSP-Tackern"): der Server lässt sich alle paar Minuten per OCSP bescheinigen, dass das eigene Zertifikat noch nicht zurückgezogen wurde, und tackert das an alle TLS-Handshakes ran

- Angriffe auf TLS
    - grundsätzlich alle Angriffe auf einzelne Primitiven sowie die Interaktion bestimmter Primitiven miteinander im Rahmen von TLS -> das geht uns zu sehr ins Detail, ist aber einer der Hauptgründe für die fortwährende Weiterentwicklung des Protokolls
    - Implementierungsfehler (z.B. Heartbleed, siehe STP037)
    - Downgrade-Angriffe: Kryptoagilität ausnutzen, um die anderen Teilnehmer zu einer weniger starken Verschlüsselung zu überreden; z.B. auch durch Machine-in-the-Middle (MITM), weswegen die Verschlüsselung schon der allerersten Handshake-Schritte so wichtig ist
    - Export-Downgrade: Spezialform des Downgrade-Angriffs, der die anderen Teilnehmer auf eine extrem schwache Exportverschlüsselung herunterzieht (bewusst schwache Verschlüsselungen aus der Zeit vor dem [Wassenaar-Abkommen](https://de.wikipedia.org/wiki/Wassenaar-Abkommen), als starke Verschlüsselung noch als exportlimitierte Militärtechnik eingestuft war)

