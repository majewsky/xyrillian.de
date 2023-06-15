- grundsätzliche Begriffsklärung
    - [Authentifizierung](https://de.wikipedia.org/w/index.php?title=Authentifizierung&oldid=227594718): Nachweis einer behaupteten Eigenschaft einer Entität; hier: der Nachweis der Identität einer Person, eines Gerätes oder eines Prozesses ("Wer ist dieser Benutzer?")
    - [Autorisierung](https://de.wikipedia.org/w/index.php?title=Autorisierung&oldid=224714316): die Einräumung von Rechten gegenüber einem interessierten Subjekt; hier: das initiale Zuweisen und das wiederholt einleitende Überprüfen von Zugriffsrechten ("Darf dieser Benutzer auf dieses Dokument zugreifen?")
    - im englischen Sprachgebrauch "authentication" (AuthN) und "authorization" (AuthZ)

- AuthN durch einen oder mehrere von drei wesentlichen Arten von Identitäts**merkmalen**
    - Wissen: etwas, dass ich weiß (Passwörter, PINs, etc.)
    - Besitz: etwas, dass ich habe (Chipkarte, Lichtbildausweis, Mobiltelefon, PC, etc.)
    - Inhärenz: etwas, dass ich bin (Fingerabdruck, Iris-Abbild, etc.) -> Achtung: Problematisch, weil nicht änderbar!
    - [Multifaktor-Authentifizierung](https://de.wikipedia.org/w/index.php?title=Multi-Faktor-Authentisierung&oldid=228949120): Wer eines der drei Merkmale gut angreifen kann, kann selten eins der anderen Merkmale gut angreifen.

- schlechte Arten von Merkmalen
    - [Magnetstreifenkarte](https://de.wikipedia.org/w/index.php?title=Magnetstreifen&oldid=226316253) (Besitz-Merkmal): trivial kopierbar -> deswegen heute meist [Chipkarten](https://de.wikipedia.org/w/index.php?title=Chipkarte&oldid=232317097)
        - Chipkarten sind kleine Computer, die sich per [Challenge-Response](https://de.wikipedia.org/w/index.php?title=Challenge-Response-Authentifizierung&oldid=228328238) authentisieren können
        - siehe zum Beispiel [ChipTAN](https://de.wikipedia.org/w/index.php?title=Transaktionsnummer&oldid=230848111#TAN-Generator)
    - Code per SMS (Besitzfaktor): auf Telefonebene angreifbar durch Malware mit erschlichener SMS-Berechtigung, auf Netzwerkebene angreifbar durch Umleitung der SMS
        - als Besitzfaktor buchstäblich "besser als nichts"
        - besser ist eine Authentisierungs-App mit [TOTP](https://de.wikipedia.org/w/index.php?title=Time-based_One-time_Password_Algorithmus&oldid=222235161)
    - zu kreativen Angriffen auf Besitzmerkmale und Biometrie siehe [die Vorträge von Starbug](https://media.ccc.de/search/?q=starbug)

- Autorisierungs-Schemata
    - früher meist [Access Control Lists (ACL)](https://de.wikipedia.org/w/index.php?title=Access_Control_List&oldid=231251676): jedes Datenobjekt hat eine Liste mit zugriffsberechtigten Benutzerkonten und Gruppen
    - Beispiel ACL in Unix: Dateisystemeinträge haben rwx-Bits (Lesezugriff, Schreibzugriff, Ausführzugriff) jeweils für das besitzende Benutzerkonto, die besitzende Gruppe und alle Benutzerkonten
    - heute meist [Role-Based Access Control (RBAC)](https://de.wikipedia.org/w/index.php?title=Role_Based_Access_Control&oldid=229433588) auf der Ebene von semantischen Operationen
    - Beispiel RBAC in [OpenStack](https://docs.openstack.org/oslo.policy/latest/): Benutzerkonten oder Benutzergruppen können Rollenzuweisungen erhalten; die Menge der zugewiesenen Rollen bestimmt, welche Aktionen ein Benutzer ausführen kann

- Authentisierung im Webkontext: [Single Sign-On (SSO)](https://de.wikipedia.org/w/index.php?title=Single_Sign-on&oldid=232103662) via [SAML](https://de.wikipedia.org/w/index.php?title=Security_Assertion_Markup_Language&oldid=222991238), [OIDC](https://de.wikipedia.org/w/index.php?title=OpenID_Connect&oldid=230166769)
    - Idee: Authentisierung bei einem zentralen Identitätsanbieter, dann Verwendung dieser festgestellten Identität in einem Zielsystem
    - SSO könnte man auch mit LDAP haben, aber SAML/OIDC erlaubt eine Einschränkung, welche Identitätsdaten der autorisierende Dienst vom authentifizierenden Dienst erhält
