- Rückbezug zu STP048: Transportverschlüsselung vs. Ende-zu-Ende-Verschlüsselung
    - Ende-zu-Ende-Verschlüsselung: zwischen den Endpunkten einer Kommunikation (z.B. im Messenger zwischen Alice und Bob)
    - Transport-Verschlüsselung: entlang eines einzelnen Schrittes des Kommunikationsweges (z.B. im Messenger einmal zwischen Alice und dem Server, und dann einmal zwischen dem Server und Bob)
    - heute reden wir über Transport-Verschlüsselung

- [Transport Layer Security (Sicherheit auf der Transportschicht)](https://de.wikipedia.org/w/index.php?title=Transport_Layer_Security&oldid=237178060): "ein Verschlüsselungsprotokoll zur sicheren Datenübertragung im Internet"
    - nicht selbst ein vollständiges Protokoll, sondern ein Fundament, dass andere Protokolle als Nutzlast trägt (vgl. STP005: Netzwerkschichten-Modell)
    - Geschichte von TLS: 1994 erste stabile Version als Teil des Netscape-Browsers (unter dem Namen SSL); in den 90ern in schneller Abfolge neue Versionen, seitdem nach Bedarf in unregelmäßigen Abständen; aktuell TLS 1.3 (und 1.2 ist auch hinreichend okay)
    - wir beschreiben jetzt die Struktur des TLS-Protokolls; hierfür sind als notwendiges Vorwissen STP043 (Kryptografische Primitiven) und STP048 (Vertrauen) erforderlich

- TLS Record: zweite und einfachere Phase der Verbindung
    - einfach eine symmetrische Verschlüsselung, um die Nutzdaten einzupacken
    - offensichtliches Problem: Wo kommt der symmetrische Schlüssel her?
    - das klingt nach einem Job für asymmetrische Verschlüsselung

- TLS Handshake: initiale Aushandlungsphase zu Beginn der Verbindung
    - erste Aufgabe: Aushandlung eines symmetrischen Schlüssels -> Diffie-Hellman-Schlüsselaustausch (siehe [STP Live](https://media.ccc.de/v/ds21-122-schlsseltechnologie-live-das-diffie-hellman-protokoll))
    - zweite Aufgabe: Überprüfung der Identität der Gegenstelle -> digitale Zertifikate (siehe STP048)
    - beide Aufgaben müssen miteinander verbunden werden -> drei (teils überlappende) Phasen
    - Phase 1: Berechnung eines initialen symmetrischen Schlüssels mit Diffie-Hellman, damit Verschlüsselung von Phase 2 und 3
    - Phase 2: Validierung des Server-Zertifikates durch den Client und ggf. des Client-Zertifikates durch den Server
    - Phase 3: Client würfelt den finalen symmetrischen Schlüssel und überträgt diesen verschlüsselt mit dem öffentlichen Schlüssel des Server-Zertifikates

- Warum zwei Schlüssel?
    - die Art und Weise der Übertragung des finalen Schlüssels stellt sicher, dass der Server den privaten Schlüssel kontrolliert
    - der Handshake-Schlüssel aus Phase 1 stellt sicher, dass auch bei nachträglicher Kompromittierung des privaten Schlüssels des Servers eine mitgeschnittene Verbindung nicht geknackt werden kann ([Perfect Forward Secrecy/Perfekte Vorwärts-Geheimhaltung](https://de.wikipedia.org/w/index.php?title=Perfect_Forward_Secrecy&oldid=236329042))

- eine wesentliche Eigenschaft von TLS: Kryptoagilität
    - TLS ist erstmal nur eine Grundstruktur
    - konkrete kryptografische Primitiven sind bis zu einem gewissen Grad gegeneinander austauschbar
    - im Laufe der Zeit Hinzufügung neuentwickelter Primitiven und Abschaffung alter Primitiven
    - Handshake enthält eine Verhandlung, in der Client und Server sich über die zu verwendenen Primitiven handelseinig werden

- Im Gespräch erwähnt:
    - ARP(Address Resolution Protocol) erklärt im [RFC-Podcast](https://requestforcomments.de/archives/126?t=3%3A19%3A10)
