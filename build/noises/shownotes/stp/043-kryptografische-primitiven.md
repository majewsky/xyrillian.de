- Rückblick zu STP041
    - Kryptografische Systeme dienen dem Erreichen der vier Schutzziele (Vertraulichkeit, Integrität, Authentizität, Verbindlichkeit)
    - Implementierung vollständiger Kryptosysteme braucht **Primitiven** (Grundbausteine)
    - heute: Vorstellung grundlegender Klassen von Primitiven

- [Symmetrische Verschlüsselung](https://de.wikipedia.org/w/index.php?title=Symmetrisches_Kryptosystem&oldid=197227672)
    - Schritt 1: Klartext (Plaintext) + Schlüssel (Key) -> Geheimtext (Ciphertext)
    - Schritt 2: Geheimtext + selber Schlüssel -> Klartext
    - üblicherweise als **Blockchiffre** mit Beschränkung des Klartexts auf eine Blockgröße von z.B. 128 oder 256 Bits, übliche Schlüsselgrößen ähnlich groß
    - zwei Probleme:
       1.  Wie verschlüsselt man mehr als einen Block?
       2.  Wie einigt man sich auf einen gemeinsamen Schlüssel, ohne dass Beobachter den Schlüssel abgreifen können?

- zu Problem 1: Erweiterung der Blockchiffre zu einem **Stromchiffre** mittels eines geeigneten [Betriebsmodus](https://de.wikipedia.org/w/index.php?title=Betriebsmodus_(Kryptographie)&oldid=209978117)
    - Datenstrom wird in einzelne Blöcke der gewünschten Größe zerteilt; Betriebsmodus beschreibt dann die Anwendung des Blockchiffre auf die Folge von Blöcken
    - einfachster Betriebsmodus ist ECB (Electronic Codebook): der Blockchiffre wird auf jeden Block komplett separat angewendet -> [offensichtliche Probleme (siehe Bild im verlinkten Abschnitt)](https://en.wikipedia.org/w/index.php?title=Block_cipher_mode_of_operation&oldid=1154901199#Electronic_codebook_(ECB))
    - bei den meisten modernen Betriebsmodi werden meist Daten aus dem aktuellen Blockchiffre in den nächsten Block übertragen, um die Analyse des Geheimtextes zu erschweren -> Problem: schwer parallelisierbar, Lesezugriff nur sequentiell
    - CTR-Betriebsmodus (Counter): jeder Block wird beim Verschlüsseln mit einem Zähler kombiniert, damit sich keine wiederkehrenden Muster wie bei ECB ergeben -> Vorteil: wahlfreier Schreib- und Lesezugriff möglich (interessant für Festplattenverschlüsselung)
    - siehe unten: Beispiel für die Geschwindigkeit moderner Stromchiffern insb. mit Hardwareunterstützung

```
$ cryptsetup benchmark
...
#     Algorithm |       Key |      Encryption |      Decryption
...
        aes-xts        512b      3547,4 MiB/s      3540,4 MiB/s
    serpent-xts        512b       742,7 MiB/s       729,4 MiB/s
    twofish-xts        512b       398,9 MiB/s       404,8 MiB/s
```


- zu Problem 2 (Einigung auf einen gemeinsamen Schlüssel) zwei Ansätze:
    1. Schlüsselaustausch-Protokolle wie [Diffie-Hellman](https://media.ccc.de/v/ds21-122-schlsseltechnologie-live-das-diffie-hellman-protokoll)
    2. asymmetrische Verschlüsselung

- [Asymmetrische Verschlüsselung](https://de.wikipedia.org/w/index.php?title=Asymmetrisches_Kryptosystem&oldid=232100921)
    - statt einem einzelnen Schlüssel ein **Schlüsselpaar**
        - wir haben keine gute Bildmetapher dazu; Xyrills Vorschlag: zwei Ratschen; eine dreht nur nach rechts und eine nur nach links
    - Schritt 1: Klartext + **öffentlicher Schlüssel** -> Geheimtext
    - Schritt 2: Geheimtext + **privater Schlüssel** -> Klartext
    - Beispiel [RSA](https://www.youtube.com/playlist?list=PL6_AeYXBHF0OWS1GxV6lfivdReSyKhv3l):
        - Grundgeheimnis sind zwei sehr große Primzahlen (`p` und `q`, jeweils z.B. 500 Dezimalstellen)
        - aus dem Produkt dieser Primzahlen (`N = p * q`) werden zwei weitere Zahlen (`d` und `e`) mittels günstiger Rechenoperationen abgeleitet
        - dann ist der öffentliche Schlüssel das Zahlenpaar `d,N` und der private Schlüssel das Zahlenpaar `e,N`
        - Ableitung des privaten aus dem öffentlichen Schlüssel erfordert Primzahlfaktorisierung von `N`, was extrem teuer ist (siehe Beispiel unten)
    - Problem 1 teilweise gelöst: Empfänger muss nur seinen öffentlichen Schlüssel offenbaren, der kein Geheimnis enthält
    - Problem 2 ungelöst: auch asymmetrische Verschlüsselung arbeitet nur pro Block

```
$ time factor 1350918345091284620469824069824602946802496824562049682046982406982406980
1350918345091284620469824069824602946802496824562049682046982406982406980: 2 2 3 5 7 193 941 5077 3488395463852632669275489205712092360621005476304268749586169
factor   1,92s user 0,00s system 99% cpu 1,926 total
$ time factor 1350918345091284620469824069824602946802496824562049682046982406982406981
1350918345091284620469824069824602946802496824562049682046982406982406981: 11 3931 5242879872157046321 5958863859256362497309559371496021100634736129421
factor   462,19s user 0,01s system 99% cpu 7:43,24 total
```

- neues Problem: asymmetrische Verfahren sind sehr viel rechenaufwändiger als symmetrische Verfahren -> **hybrides Kryptosystem**
    - asymmetrische Verschlüsselung wird nur zum Austausch eines symmetrischen Schlüssels verwendet
    - Beispiel "verschlüsselte E-Mail": Alice hat den öffentlichen Schlüssel `O`, Bob hat den privaten Schlüssel `P` dazu
        - Alice erstellt die Nachricht `N` im Klartext
        - Alice erzeugt einen zufälligen symmetrischen Schlüssel `S`
        - Alice verschlüsselt `N` mit `S` und einem Stromchiffre zu `N'`
        - Alice verschlüsselt `S` mit `O` zu `S'`
        - Alice schickt `N'` und `S'` an Bob
        - Bob entschlüsselt `S'` mit `P` zu `S`
        - Bob entschlüsselt `N'` mit `S` zu `N`

- zweite Anwendung von asymmetrischer Kryptografie: [Signaturen](https://de.wikipedia.org/w/index.php?title=Digitale_Signatur&oldid=234047519)
    - Verschlüsselung: vorwärts mit dem öffentlichen Schlüssel, rückwärts mit dem privaten Schlüssel
    - bei Signierung umgekehrt: vorwärts mit dem privaten Schlüssel (Erzeugung), rückwärts mit dem öffentlichen Schlüssel (Überprüfung)
    - im Detail: erst Komprimierung der zu signierenden Daten auf eine feste Größe mittels einer Hashfunktion (siehe STP004), dann Verschlüsselung dieses Hash-Wertes mit dem privaten Schlüssel
    - Überprüfung: Berechnung des Hashes auf dieselbe Weise, Vergleich mit entschlüsselter Signatur

- noch eine Primitive zum Schluss: [Geheimnisteilung](https://de.wikipedia.org/w/index.php?title=Secret-Sharing&oldid=201132633)
    - Problem: Ein Geheimnis soll in N Teile so aufgeteilt werden, dass M beliebige Teile zum Entschlüsseln genutzt werden können.
    - Beispiel: Backup des Passwortspeichers in 5 Fragmente aufteilen und bei 5 Freunden so einlagern, dass man später nur 3 beliebige Fragmente zum Wiederherstellen braucht
    - Illustrationen: [Visuelle Kryptografie](https://de.wikipedia.org/w/index.php?title=Visuelle_Kryptographie&oldid=204692019), [Blakley-Schema](https://en.wikipedia.org/w/index.php?title=Secret_sharing&oldid=1152303869#Blakley's_scheme)

- im Gespräch erwähnt
    - [Anerzählt Podcast zu AES](https://anerzaehlt.net/aes-256/)
    - [Whitfield Diffie und Moxie Marlinspike und eine Flasche Whiskey auf der Bühne](https://www.youtube.com/watch?v=lt7uW6vDk00)
