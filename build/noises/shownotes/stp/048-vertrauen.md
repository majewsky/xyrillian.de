- Rückbezug zu STP039: Authentifizierung ("Wer ist diese Person?" bzw "Was darf diese Person?")
    - Problem: Woher weiß man, ob ein Identitätsnachweis verlässlich ist? -> **Vertrauen**
    - Vorstellung verschiedener Arten von Vertrauensvermittlung -> Ziel: Vokabular aufbauen, Methoden gegeneinander abwägen
    - Beispiele im realen Leben: gemeinsame Vorgeschichte und Inside-Jokes, Geheimlosungen, Personalausweise

- Vererbung von Vertrauen: [Digitale Zertifikate](https://de.wikipedia.org/w/index.php?title=Digitales_Zertifikat&oldid=236185751) nach dem [X.509-Standard](https://de.wikipedia.org/w/index.php?title=X.509&oldid=237108927)
    - basierend auf asymmetrischer Verschlüsselung (siehe STP043)
        - Zertifikat = öffentlicher Schlüssel + Identitätsinformationen + Signatur durch ein übergeordnetes Zertifikat
        - durch die Signaturkette geht Vertrauen vom Anfang der Kette (**Zertifikatsautorität**, CA) auf die abgeleiteten Zertifikate über
    - authentifizierende Systeme müssen nur das allererste Zertifikat in der Kette kennen (das Wurzel-Zertifikat bzw. Root-CA)
        - Beispiel: Wurzelzertifikat signiert Zertifikat einer unterliegenden CA, und dieses signiert ein Zertifikat für einen Nutzer
    - Einsatzfälle
        - Transportverschlüsselung im Web (HTTPS, Zertifikat enthält Domain-Namen)
        - Impfzertifikate in der Corona-Warn-App
        - Benutzerseitige Zertifikate im Firmen-Intranet oder bei Elster
        - Äquivalente in der realen Welt: Personalausweis oder Reisepass (ausgestellt durch eine Regierungsstelle), Zugangskarte in der Firma (ausgestellt durch die Liegenschaftsverwaltung der Firma), Mitarbeiteruniform (eine Person in DB-Uniform auf einem Bahnhof wird wahrscheinlich ein DB-Mitarbeiter sein)
    - Vorteile: skaliert sehr gut auf viele Akteure, nur relativ wenig Speicher notwendig (für die Root-CAs)
    - Nachteil: erfordert Vertrauen in eine zentrale Stelle

- Vertrauen durch Beständigkeit: [Trust on first use (TOFU)](https://en.wikipedia.org/w/index.php?title=Trust_on_first_use&oldid=1157774410)
    - beim ersten Kontakt mit einer Person wird deren öffentlichen Schlüssel blind vertraut
    - bei weiteren Kontakten wird demselben Schlüssel weiterhin vertraut, aber die Präsentation eines neuen Schlüssels wird als Angriffsversuch bzw. Betrugsversuch gewertet
    - Einsatzfälle
        - die meisten Messenger mit Ende-zu-Ende-Verschlüsselung (z.B. Signal, WhatsApp, verschlüsselte Chats in Telegram)
        - Fernzugriff auf Server mittels SSH (für den Teil, wo die Identität des Servers überprüft wird)
        - Äquivalente in der realen Welt: Bekanntschaften unter Nachbarn oder Vereinsmitgliedern (alles, was auf der Basis von "den kenn ich" läuft)
    - Vorteil: braucht keinen zentralen Vertrauensanker
    - Nachteil: anfällig gegen Attacken beim initialen Verbindungsaufbau

- Vertrauen durch Vitamin B: [Web of Trust (WoT)](https://de.wikipedia.org/w/index.php?title=Web_of_Trust&oldid=235927300)
    - Zertifikate ähnlich wie bei X.509: öffentlicher Schlüssel + Identitätsinformatione + **Liste von** Signaturen
    - Signaturen nicht vertikal (in einer Hierarchie), sondern horizontal (zwischen gleichrangigen Zertifikaten verschiedener Nutzer)
    - Vertrauen vererbt sich zwischen Nutzern:
        - unendliches Vertrauen in das eigene Zertifikat
        - hohes Vertrauen in Zertifikate, die man selbst signiert hat
        - mittelhohes Vertrauen in Zertifikate, die von diesen Zertifikaten signiert wurden
        - und so weiter
    - Einsatzfälle
        - E-Mail-Verschlüsselung mit PGP/GPG, siehe auch [Keysigning-Parties](https://de.wikipedia.org/w/index.php?title=Keysigning-Party&oldid=237529163)
        - Verifikation eines Chat-Kontakts über einen Seitenkanal
        - Äquivalent in der realen Welt: Bekanntschaften über Bande ("der Freund der Freundin des Schwippschwagers")
    - Vorteil: braucht keinen zentralen Vertrauensanker
    - Nachteil: Schlüssel mit vielen Signaturen sind eigentlich wertvoll, aber auch unhandlich; Angriff möglich durch bösartige Müll-Signaturen

- Seitenleiste: Wie könnte man solchen Signatur-Spam verhindern?
    - Problem: Spam (sowohl bei E-Mails als auch GPG-Signaturen) ist für den Versender sehr günstig und für den Empfänger extrem lästig -> Kann man den Spieß umdrehen?
    - Idee (1990er): wir nutzen Einwegfunktionen (Streuwertfunktionen siehe STP004, STP043), aber drehen die Sache um
        - Sender muss einen Datensatz bilden, für den eine Streuwertfunktion ein Ergebnis mit X führenden Nullen hat (X kann variiert werden je nach Leistungsfähigkeit aktueller Prozessoren) -> erfordert sehr viel Durchprobieren von Eingabewerten
        - Empfänger muss nur schauen, dass die Lösung stimmt -> erfordert nur eine Berechnung mit dem präsentierten Eingabewert
    - Implementation (2002): [Hashcash](https://en.wikipedia.org/w/index.php?title=Hashcash&oldid=1166511502) als eines der ersten [Proof-of-Work](https://de.wikipedia.org/w/index.php?title=Proof_of_Work&oldid=235827513)-Verfahren

- Vertrauen durch Proof of Work: [Blockchain](https://de.wikipedia.org/w/index.php?title=Blockchain&oldid=235793427)
    - Kombination von Proof of Work mit Merkle-Bäumen
    - Merkle-Baum = Baumstruktur aus Datenpaketen, die jeweils die Hash-Werte ihrer untergeordneten Knoten enthalten -> Hash-Wert des Wurzelknotens ist die einzige benötigte Information, um die Integrität des gesamten Baumes zu überprüfen
    - Blockchain: System zur kontinuierlichen Erweiterung eines Merkle-Baums (genauer: einer Merkle-Kette) um neue Blöcke, aber die Hashes jedes Blocks müssen eine Proof-of-Work-Bedingung erfüllen
    - kein Vertrauen in einzelne Akteure notwendig, nur in die Sicherheit des Algorithmus

- Abendgedanken: Vertrauen lässt sich nicht eliminieren, nur verschieben
    - Vertraue ich einem Menschen? Einer Maschine? Einer Organisation? Einem Algorithmus?
    - Niemandem vertrauen ist auch keine Lösung.

- im Gespräch erwähnt
    - Versionsverwaltungssystem "git", wie von uns besprochen im [Pentaradio vom Juli 2021](https://c3d2.de/news/pentaradio24-20210727.html)
