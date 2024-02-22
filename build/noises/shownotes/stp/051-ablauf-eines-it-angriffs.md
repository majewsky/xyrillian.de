> \[Unser Arbeitsfeld] heißt "Computer-Sicherheit", nicht "Computer-Optimismus". Wir machen uns der schlimmstmöglichen Resultate bewusst, denn auch wenn wir es nicht tun, unsere Gegner werden es auf jeden Fall tun.
>
> – [Matthew Green über IT-Sicherheit](https://threadreaderapp.com/thread/1433451378391883782.html) (übersetzt aus dem Englischen)

- Topologie eines IT-Angriffs auf ein verteiltes System (z.B. das interne Netz eines Unternehmens oder einer Organisation)
    - Erstangriff entweder durch reinen Exploit eines von außen erreichbaren Systemes (z.B. der Firmenwebseite) oder durch Einschleusung eines Exploits mittels [Social Engineering](https://de.wikipedia.org/w/index.php?title=Social_Engineering_(Sicherheit)&oldid=233291013) (Manipulation von Menschen mit Zugriff auf das System mit dem Ziel, einen Exploit einzuschleusen oder Informationen wie Passwörter auszuschleusen; z.B. [Phishing](https://de.wikipedia.org/w/index.php?title=Phishing&oldid=237076459))
    - basierend auf erstem Zugriff in das System dann **Privilege Escalation** (Verwendung weiterer Exploits, um schrittweise immer weitergehenden Zugriff zu erlangen) und **Lateral Movement** (Ausnutzung des Zugriffs auf ein Teilsystem, um andere verwundbare Systeme im Netzwerk aufzuspüren und anzugreifen)

- Welche Ziele kann ein IT-Angriff letztendlich verfolgen?
    - Datendiebstahl: z.B. Spionage, Wirtschaftsspionage, Erpressung mittels [Kompromat](https://de.wikipedia.org/w/index.php?title=Kompromat&oldid=237791818), [Doxing](https://de.wikipedia.org/w/index.php?title=Doxing&oldid=237524584)
    - Propaganda: z.B. Verbreitung von Desinformation und Fake-News durch gekaperte Accounts von Personen des öffentlichen Lebens
    - Datendiebstahl: z.B. Datenbanken von Nutzerdaten (Passwörter, Mail-Adressen, Kontodaten etc.) können für weitere Angriffe und Betrugsversuche vermarktet werden
    - Finanzdiebstahl: z.B. Ausnutzung illegitimen Zugriffs auf Bankkonten, Kreditkarten oder auch Konten in virtuellen Währungen
    - Sabotage: z.B. Ransomware (Infektion durch Virus oder Trojaner, dann evtl. Lateral Movement, dann Verschlüsselung aller erreichbaren Festplatten und Backups, dann Erpressung von Lösegeld)
    - Sabotage: z.B. [Stuxnet](https://de.wikipedia.org/w/index.php?title=Stuxnet&oldid=237385358)

- Was kann man auf der Verteidigerseite machen?
    - keine Sicherheitslücken haben :)
    - Verwendung von sichereren Programmiermethodiken (z.B. moderne Programmiersprachen, die Speicherverwaltungsfehler systematisch verhindern)
    - Einsatz von Sandboxing (siehe STP023) und anderen Techniken, um die Privilege Escalation zu erschweren ("Principle of Least Privilege": jeder Prozess sollte mit den geringsten möglichen Rechten laufen, und jedes Benutzerkonto sollte nur genau die Berechtigungen haben, die für die Erfüllung der jeweiligen legitimen Aufgaben notwendig sind)
    - aktive Verteidigungsmaßnahmen wie Virenscanner oder Intrusion-Detection-Software sind umstritten
    - allgemeiner: [Threat Modeling](https://en.wikipedia.org/w/index.php?title=Threat_model&oldid=1177652493); das systematische Auflisten von möglichen Angriffen und Schwachstellen, um dann entsprechende Gegenmaßnahmen zu koordinieren
    - nach dem Angriff: Attribuierung (die Zuschreibung eines Angriffs zu einem Angreifer) ist zwar oft gewünscht (man hat halt gerne einen Schuldigen, gerade in der politischen Arena), ist aber meist extrem schwer (es ist deutlich einfacher, als Angreifer falsche Fährten zu legen, als sie als Forensiker als solche zu erkennen)

- weiterführende Quellen
    - zu Social Engineering: [Hirne hacken](https://media.ccc.de/v/36c3-11175-hirne_hacken)
    - zu digitalen Waffen und Hackback: [Wie Hackback mit der Gesellschaft spielt](https://media.ccc.de/v/2019-218-wie-hackback-mit-der-gesellschaft-spielt)
