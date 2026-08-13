- Prolog: in der IT-Sicherheit meist der Angreifer im Vorteil, weil vielfache Wiederholung bzw. Fächerung von Angriffen vergleichsweise einfach ist und am Ende nur einer durchkommen muss
    - offensive Verteidigung ("Hackback") problematisch aufgrund schwieriger Attributierung ([wenn der Angreifer keine offensichtlichen Fehler macht](https://media.ccc.de/v/35c3-9716-du_kannst_alles_hacken_du_darfst_dich_nur_nicht_erwischen_lassen)), also defensive Strategien notwendig
    - bisher zur Angriffsseite: STP049 Schadcode, STP051 Ablauf eines IT-Angriffs
    - bisher zur Verteidigungsseite: STP023 Virtualisierung (und darin Sandboxing), STP038 Software-Tests, STP078/080 Common Weakness Enumeration
    - heute heben wir wieder den Schild anstelle des Schwertes

- [Code Signing](https://en.wikipedia.org/w/index.php?title=Code_signing&oldid=1345047115)
    - Idee: nur Code aus vertrauenswürdigen Quellen ausführen
    - vgl. STP048: Vertrauen klassischerweise durch eine PKI nach X.509 (analog zu Transportverschlüsselung TLS bei Webseiten)
    - als Wurzelzertifikatsautorität tritt der Plattformbetreiber auf, z.B. der App-Store-Betreiber oder bei Kerneltreibern der Betriebssystem-Hersteller

- Seitenleiste: die entsprechenden Signaturschlüssel sind sicherheitskritisch und werden heutzutage meist in [HSM (Hardware-Sicherheitsmodulen)](https://de.wikipedia.org/w/index.php?title=Hardware-Sicherheitsmodul&oldid=253958830) aufbewahrt

- Problem: Was bringt eine signierte Programmdatei, wenn die Ausführungsumgebung schon kompromittiert ist?
    - Idee: [Secure Boot](https://de.wikipedia.org/w/index.php?title=Unified_Extensible_Firmware_Interface&oldid=265805867#Secure_Boot) und vergleichbare Technologien
    - schon der Prozessor selber kann beim Start Code-Signaturen prüfen, ab dann Code Signing auf allen Ebenen über das Betriebssystem bis zum Anwendungsprogramm

- Vorteile von Code Signing und Secure Boot:
    - verhindert bestimmte Einschlesungspfade für Malware (z.B. wenn man "libreoffice download" sucht und nicht die offizielle LibreOffice-Webseite findet, sondern eine legitim wirkende Virenschleuder)
    - verhindert bestimmte Methoden für Lateral Movement und Persistierung (selbst wenn ein Prozess kompromittiert ist, kann er nicht zum Beispiel neue Programmdateien mit Malware auf die Platte schreiben und ausführen)

- Nachteile von Code Signing und Secure Boot:
    - kein perfekter Schutz (aber was ist das schon)
    - Ablauf des signierenden Zertifikats könnte ansonsten unveränderte Systeme funktionsunfähig machen -> deswegen bei Code Signing meist nur Prüfung auf "Signatur vor Ablauf", nicht auf "aktuelle Zeit vor Ablauf"
    - erlaubt Willkür seitens der Geräte- bzw. Betriebssystemhersteller (siehe z.B. Entfernung von ICEBlock aus dem iOS-App-Store)
    - zieht Eigentümerschaft am Gerät in Zweifel; vgl. [Jay Freemans](https://en.wikipedia.org/w/index.php?title=Jay_Freeman&oldid=1336519424) Diktum von "Felony Contempt of Business Model" ("verbrecherische Missachtung des Geschäftsmodells") zur Beschreibung des Effektes von DMCA Section 1201 und vergleichbaren Gesetzen (in Dtl. [§ 95a UrhG](https://www.gesetze-im-internet.de/urhg/__95a.html)), welches die Umgehung von "wirksamen technischen Maßnahmen" zum Schutz von Urheberrechten kriminalisiert

- ganz anderer Ansatz: Reduktion der [Trusted Computing Base (TCB)](https://en.wikipedia.org/w/index.php?title=Trusted_computing_base&oldid=1345500081) ("vertrauenswürdige Rechenbasis" bzw. eher "der man vertrauen muss")
    - TCB: "die Gesamtmenge an Hardware und Software, die für den sicheren Betrieb eines Systems essentiell sind"
    - nicht ganz dasselbe wie "weniger Programmcode = weniger Programmierfehler"
    - z.B. bei Sandboxing ist nur die Sandbox selber Teil der PCB; Code innerhalb der Sandbox sollte dann (hoffentlich!) auch im Fehler- oder Infektionsfall keinen Schaden anrichten können
    - praktisches Beispiel: in meinem LDAP-Server [Portunus](https://github.com/majewsky/portunus) gibt es einen Teil, der mit Root-Rechten ausgeführt werden muss; dort habe ich mal [reguläre Ausdrücke (vgl. STP021) durch explizite Parser ersetzt](https://github.com/majewsky/portunus/commit/411a8bb2fa0a1ecfe2325f2f543514667f5c5bda), weil das Entfernen der Regex-Bibliothek [das Root-Binary um 10-15% verkleinert hat](https://github.com/majewsky/portunus/commit/e783db98f5e290534c58365e51c108ce1aeaf457)
    - Randbemerkung: TCB-Reduktion ist auch, warum ich lieber 100 Zeilen Code selber schreibe, als eine Bibliothek zu importieren, die die eine Sache macht, die ich brauche... und dann noch zig andere Sachen mit unzähligen weiteren Abhängigkeiten ([Symbolbild](https://i.programmerhumor.io/2025/08/fe73cb25f8fb3c7790677b292f7f84a3ac75b728e693972784203f6e1bc6d34e.png))

- Videoempfehlung für diese Art, über Software nachzudenken: [Fefe über "Das Nützlich-Unbedenklich-Spektrum"](https://media.ccc.de/v/36c3-10608-das_nutzlich-unbedenklich_spektrum)
    - in den Lektionen einiger Überlapp mit "Vom Mythos des Mann-Monats" (vgl. STP068/STP074/082)
    - aber mit deprimierendem Einschlag: "Angenommen, jemand findet einen Weg, besser mit Komplexität umzugehen. Was passiert dann? Schreiben wir dann alle die alte Software neu, aber besser? Nein! Wir schreiben dann neue, noch größere Software. Wieder am Limit dessen, was mit den neuen Methoden machbar ist." (Der Talk ist von 2019. Jede Parallele zu Coding-Agenten ist rein zufällig.)
    - Xyrills Favorit: die Wortprägung "Bugwelle" für die immer größer werdende Flut von Bugreports, die ein langlaufendes Softwareprojekt vor sich herschiebt
