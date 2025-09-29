> Through twenty years of effort, we've successfully trained everyone to use passwords that are hard for humans to remember, but easy for computers to guess.
>
> Wir haben zwanzig \[mittlerweile dreißig &ndash; d. Red.\] Jahre lang allen eingeredet, Passwörter zu verwenden, die schwer zu merken sind, aber leicht von Computern geraten werden können.

Quelle: Zitat aus [XKCD Comic #936](https://xkcd.com/936/).

* Motivation: Authentifizierung (Identitätsbeweis)
    * mögliche Auth.faktoren:
        * etwas, das ich weiß (z.B. Passwort, PIN, ~~Sicherheitsfrage~~)
        * etwas, das ich habe (z.B. Chipkarte, Hardware-Token, TOTP/HOTP)
        * etwas, das ich bin (z.B. Fingerabdruck, Iris-Abbild)
    * Ein-Faktor-Auth. vs. Zwei-Faktor-Auth.

* Theorie: Hash-Funktionen
    * [Hash-Funktion](https://de.wikipedia.org/wiki/Hashfunktion ):
        * Wikipedia sagt als Begriff "Streuwertfunktion", das erklärt es ganz gut
        * Anwendungsfelder: Hashtabellen, Prüfsummen
    * [kryptografische Hash-Funktion](https://de.wikipedia.org/wiki/Kryptographische_Hashfunktion ):
        * kollisionsresistent und dadurch zur "Ein-Weg-Verschlüsselung" geeignet
        * typische Hash-Länge: 256 Bit = 32 Byte = 64 Hexadezimalziffern
        * Größenvergleich: 2^256 Nanosekunden = 3 Mio. Trio. Trio. Trio. Jahre
    * Ziel: Passwort nicht direkt auf dem Webserver speichern
        * Server speichert nur Hash vom Passwort
        * zur Überprüfung präsentiertes Passwort hashen und mit gespeichertem Hash vergleichen
        * Salting: zu jedem Hash noch ein Extra-Zufallstext, damit dasselbe Passwort nicht immer denselben Hash ergibt

* Was können Nutzer machen, um sich zu schützen?
    * schwer ratbare Passwörter verwenden
        * keine persönlichen Informationen hineinkodieren (Namen von Angehörigen oder nahestehenden Personen, Geburtsdaten, etc.)
        * NICHT "passwort" oder "123456" oder "123qwe"
        * z.B. XKCD-Schema: Wörterbuch aufschlagen, vier zufällige Wörter wählen
        * NICHT bekannte Phrasen (z.B. Bibelverse)
    * Passwörter nicht wiederverwenden
        * Wenn überall das gleiche Passwort verwendet wird, sind alle Konten gleichzeitig offen.
        * Mit dem Zugang zum E-Mail Konto kann die "Passwort zurücksetzen"-Funktion ausgenutzt werden.
        * Wenn du nur ein starkes Passwort benutzt, dann das für das E-Mail-Konto.
    * Passwort-Manager verwenden
        * auf dem lokalen PC: [KeePassXC](https://keepassxc.org/ )
        * oder cloud-basiert: [Bitwarden](https://bitwarden.com/ ) (Server kann man auch selbst betreiben)
        * oder für Nerds: [pass](https://www.passwordstore.org/ )
    * Passwörter nicht mit anderen teilen, auch nicht "nur mal schnell"
        * siehe IT-Support-Scam
        * bei Mehrbenutzersystemen auch mehrere Benutzerkonten verwenden
    * Passwort nur auf vertrauenswürdigen Seiten eingeben
        * nicht Links in irgendwelchen Mails folgen (siehe Phishing-Scam), sondern direkt zur bekannten Webseite navigieren
    * [Passwort für Webseiten ändern, bei denen die Datenbanken gestohlen wurden](https://haveibeenpwned.com/)
        * in vielen Passwort-Managern als automatische Prüfung integriert

* Rant: Sicherheitsmaßnahmen, die nicht wirklich bessere Sicherheit bringen
    * Sicherheitsfragen
    * Passwortbeschränkungen: maximal N Zeichen, mindestens ein Sonderzeichen etc.
    * [Zwangswechsel nach N Monaten](https://arstechnica.com/information-technology/2019/06/microsoft-says-mandatory-password-changing-is-ancient-and-obsolete/)

* Was können Webseitenbetreiber machen, um Auth. sicher zu machen?
    * Passwörter nur mit starken Hashes ablegen (Argon2, SCrypt, BCrypt, notfalls auch PBKDF2)
    * 2FA unterstützen
    * [HIBP-API](https://haveibeenpwned.com/API/v3) verwenden, um Passwort-Wiederverwendung zu erschweren

* Gibt es eine Alternative zu Passwörtern?
    * anderer Faktor (z.B. Hardware-Token oder Biometrie): nicht unbedingt besser, nur anders (Abwägungsfrage)
    * Zwei-Faktor-Auth.
    * Anmeldungslink per E-Mail
