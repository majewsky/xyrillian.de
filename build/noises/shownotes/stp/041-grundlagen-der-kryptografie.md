Definition laut [Wiktionary](https://de.wiktionary.org/w/index.php?title=Kryptografie&oldid=9657177):

> "Kryptografie", aus griechisch κρυπτός (kryptos) „verborgen“ und -γραφία (graphia), zum Verb γράφειν (graphein) „schreiben“ gebildet
>
> 1. _Informatik, Mathematik:_ Wissenschaft der Verschlüsselung von Informationen
> 2. _Psychologie:_ absichtslos entstandene Kritzelzeichnung eines Erwachsenen

- Warum Verschlüsselung? Vier **Schutzziele**
    - **Vertraulichkeit:** Eine Nachricht soll nur von berechtigten Personen oder Maschinen gelesen werden dürfen.
    - **Integrität:** Nachrichten sollen nachweislich vollständig und unverändert sein.
    - **Authentizität:** Die Nachricht soll ihrem tatsächlichen Absender zugeordnet werden können.
    - **Verbindlichkeit:** Der Absender einer Nachricht soll seine Urheberschaft nicht abstreiten können.
    - Beispiele aus der analogen Welt: Verträge, Geheimnisträger:innen (Anwält:innen, Geistliche, Mediziner:innen, etc.)
    - ["Ich hab doch nichts zu verbergen"](https://digitalcourage.de/nichts-zu-verbergen) -> Unterschied zwischen "privat" und "geheim"

- Verschlüsselung im engeren Sinne: Übertragung eines **Klartext** ([Plaintext](https://en.wikipedia.org/w/index.php?title=Plaintext&oldid=1147935435)) in einen **Geheimtext** ([Ciphertext](https://en.wikipedia.org/w/index.php?title=Plaintext&oldid=1147935435)) und später wieder zurück
    - Beispiel für (historisch signifikanten) Ciphertext: [Zimmermann-Depesche](https://de.wikipedia.org/w/index.php?title=Zimmermann-Depesche&oldid=232609215)
    - die Relation zwischen Klartext und Ciphertext wird vermittelt durch einen **Schlüssel**, also eine Folge von Bits fester Länge (je nach Verfahren zwischen 128 Bits bis zu 4096 Bits)

- Grundbausteine eines Kryptosystems: **Primitiven**
    - Beispiel für eine Klasse von Primitiven: symmetrische Verschlüsselung (Funktionsprinzip: Klartext + Schlüssel -> Geheimtext, Geheimtext + selber Schlüssel -> Klartext)
    - Qualitätsmaßstab: Angreifer mit Zugriff auf den Geheimtext soll nicht auf den Klartext schließen können
    - immer möglich: Angriff nach [Brute-Force-Methode](https://de.wikipedia.org/w/index.php?title=Brute-Force-Methode&oldid=230198355), also Durchprobieren aller möglichen Schlüssel (Beispiel: Zahlenschloss mit vierstelliger PIN)
    - praktischer Maßstab: Brechen einer Verschlüsselung soll nicht weniger als `2^100` Rechenoperationen erfordern (`2^100 = 1.267... * 10^30`)
    - Rechenbeispiel: Xyrills Grafikkarte kann im Idealfall etwa 64 TFLOPS liefern, also 64 Trillionen Rechenoperationen pro Sekunde; bei dieser Geschwindigkeit benötigen `2^100` Rechenoperationen etwa 627 Millionen Jahre
    - Seitenleiste: Quantencomputer sind nicht automatisch Game Over, aber es gibt aktive Forschung an [Post-Quanten-Kryptografie](https://de.wikipedia.org/w/index.php?title=Post-Quanten-Kryptographie&oldid=228540985)

- Primitiven alleine machen noch kein vollständiges Kryptosystem
    - die tollste Festplattenverschlüsselung hilft nicht, wenn ich das Passwort am Monitor kleben habe :)
    - besonders spannend: [Seitenkanal-Angriffe](https://de.wikipedia.org/w/index.php?title=Seitenkanalattacke&oldid=218845820)
    - Beispiel: je nach Bitfolge im Schlüssel dauert das Entschlüsseln unterschiedlich lang oder braucht unterschiedlich viel Energie oder erzeugt unterschiedlich viel Strahlung
    - Beispiel: Login mit Benutzername und Passwort sollte sich bei falschem Passwort genauso verhalten wie bei nicht vorhandenem Benutzerkonto

- Vorschau
    - nächstes Mal (STP042) was ganz anderes... aber dann!
    - STP043: Überblick über Klassen von Primitiven (Fokus auf das allgemeine Verhalten anstatt auf konkrete mathematische Konstruktion)
    - STP044: Starke Kryptografie ist kein Zufall... oder doch?

- im Gespräch erwähnt
    - [Schlüsseltechnologie Live: Das Diffie-Hellman-Protokoll](https://media.ccc.de/v/ds21-122-schlsseltechnologie-live-das-diffie-hellman-protokoll)
    - [Kerckhoffs' Prinzip](https://de.wikipedia.org/w/index.php?title=Kerckhoffs%E2%80%99_Prinzip&oldid=213145792)
    - [XKCD 538: Security](https://xkcd.com/538/)
    - [crypto/subtle](https://pkg.go.dev/crypto/subtle)

