- Platz 9: [CWE-862](https://cwe.mitre.org/data/definitions/862.html); fehlende Autorisierung
    - siehe Besprechung von Authentifizierung und Autorisierung in STP039
    - Lösung: Just do it! :)
    - passend dazu Platz 14 und 18: [CWE-287](https://cwe.mitre.org/data/definitions/287.html) und [CWE-863](https://cwe.mitre.org/data/definitions/863.html); fehlerhafte Authentifizierung bzw. Autorisierung
    - außerdem Platz 15: [CWE-269](https://cwe.mitre.org/data/definitions/269.html); unzureichende Privilegienkontrolle
    - als Sonderfall Platz 25: [CWE-306](https://cwe.mitre.org/data/definitions/306.html); fehlende Authentifizierung an kritischer Stelle
    - daraus folgend Platz 17: [CWE-200](https://cwe.mitre.org/data/definitions/200.html); Preisgabe sensitiver Informationen an nicht Zugriffsberechtigte

- Platz 10: [CWE-434](https://cwe.mitre.org/data/definitions/434.html); Unbeschränktes Hochladen von Dateien des falschen Typs
    - und dadurch z.B. Einschleusung von Schadcode (Xyrill fühlt sich an den "OpenOffice kann dienstags nicht drucken"-Bug erinnert)
    - Lösung: nur bekannte Formate akzeptieren
    - passend dazu Platz 16: [CWE-502](https://cwe.mitre.org/data/definitions/502.html); Deserialisieren unvertrauenswürdiger Daten
    - Lösung: unbekannte Daten mit [sicheren Parsern](https://github.com/google/wuffs) validieren

- Platz 21: [CWE-476](https://cwe.mitre.org/data/definitions/476.html); Zugriff auf Nullzeiger
    - besprochen in STP045

- Platz 22: [CWE-798](https://cwe.mitre.org/data/definitions/798.html); Zugriff mit hartkodierten Anmeldedaten
    - In wieviele Plasterouter kommt man wohl immer noch mit Username `admin` und Passwort `letmein` rein? :)
    - tatsächlich auch bei Enterprise-Hardware oftmals Standardpasswörter voreingestellt, die dann beim Einrichten (versehentlich?) nicht geändert werden

- Platz 23: [CWE-190](https://cwe.mitre.org/data/definitions/190.html); Ganzzahlüberlauf ("Integer Overflow or Wraparound")
    - besprochen in STP003 (Überlauf in Logikgattern) und STP037 (berühmte Überlauffehler)

- Platz 24: [CWE-400](https://cwe.mitre.org/data/definitions/400.html); unkontrollierter Ressourcenverbrauch
    - Rückbezug zu STP079 (Komplexitätsattacke in Hashmaps)

- außerdem noch Xyrills Favorit: [CWE-655](https://cwe.mitre.org/data/definitions/655.html); unzureichende psychologische Akzeptabilität :)
