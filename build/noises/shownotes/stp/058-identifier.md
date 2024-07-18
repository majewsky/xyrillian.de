Wir haben einiges an Feedback bekommen, siehe Ende der Shownotes. Zuerst zum eigentlichen Thema:

- **Identifier** (ID): eindeutiges Kennzeichen zur Unterscheidung gleichartiger Objekte oder Individuen
    - während des Softwaredesign (z.B. beim Entwurf des Datenbankschemas): Entscheidung darüber, welche Identifier für welche Kategorien von Objekten verwendet werden

- Identifier für Menschen: [Falsehoods Programmers Believe About Names](https://web.archive.org/web/20240325164409/shinesolutions.com/2018/01/08/falsehoods-programmers-believe-about-names-with-examples/)
    - bei deutschen Behörden üblich: das Tupel aus Name und Geburtsdatum und evtl. Geburtsort -> [immer noch nicht ganz eindeutig](https://web.archive.org/web/20240404170829/https://www.tagesspiegel.de/berlin/stadtleben/die-doppelte-katja-4792738.html)
    - bei Benutzerkonten im wesentlichen drei Optionen:
        - fortlaufende Nummern (Vorteile: im Speicher kompakt darstellbar, robustes Pseudonym; Nachteile: leakt das Alter des Kontos, rare kurze Nummern vermitteln Status), siehe z.B. ICQ-Kontonummern (frühe 2000er)
        - frei wählbare Benutzernamen (Vorteile: eingängig, Pseudonymität möglich; Nachteile: belohnt frühe Registrierung, fördert Name-Squatting)
        - E-Mail-Adressen (Vorteile: enthält implizit eine Kontaktmöglichkeit, wird seltener vergessen als eine Nummer oder ein frei wählbarer Name; Nachteile: muss vertraulich behandelt werden, beim Wechsel der Mail-Adresse evtl. Neuanmeldung erforderlich)

- auch bei Objekten oder Datenstrukturen eine gute Option: **fortlaufende Nummern**
    - im deutschen Handelsrecht bei Rechnungen sogar in gewissem Rahmen rechtlich vorgeschrieben
    - Vorteil: Hochzählen ist trivial umzusetzen
    - Vorteil: kompakte Darstellung im Speicher (z.B. nur 4 Byte bei einer Gesamtzahl unter 4 Mrd. Objekte) -> insb. wichtig bei Kreuzreferenzen
    - Nachteil: ab und zu hat man doch mehr Objekte als erwartet und muss mühselig auf größere Zahlen umstellen (z.B. bei über 4 Mrd. Objekten von 32-Bit-Zahlen auf 64-Bit-Zahlen) -> [Beispiel bei der Schachwebseite chess.com](https://news.ycombinator.com/item?id=14539770)
    - Nachteil: in verteilten Systemen (siehe STP027) braucht man eine zentrale ID-Vergabestelle oder einen aufwändigen Konsensalgorithmus -> Geht das besser?

- [Universally Unique Identifier (UUID)](https://de.wikipedia.org/w/index.php?title=Universally_Unique_Identifier&oldid=243706271) gemäß [RFC 4122](https://datatracker.ietf.org/doc/html/rfc4122)
    - Idee: mit dem Schritt von 64 Bit auf 128 Bit so viele mögliche Werte, dass man z.B. bei zufälliger Wahl eine absurd geringe Kollisionswahrscheinlichkeit hat
    - Darstellung in Text meistens als Hexadezimalzahl in Gruppen von 8, 4, 4, 4 und 12 Stellen; Beispiel: `8b4b72a2-55dd-4ab7-b239-bb9caad9107a`
    - mehrere verschiedene Formate: heutzutage am häufigsten UUIDv4 (mit rein zufälligen Zahlen, außer dass die 13. Hexadezimalstelle die Formatzahl "4" enthält)
    - Vorteil: in verteilten Systemen keine Koordination bei der Vergabe erforderlich
    - Nachteil: etwas weniger kompakt, insb. wenn die UUID wie oft üblich als Text abgespeichert wird
    - Nachteil insb. bei UUIDv4: keine sinnvolle Sortierung -> Alternativvorschläge wie [UUIDv7](https://www.ietf.org/archive/id/draft-ietf-uuidrev-rfc4122bis-07.html#name-uuid-version-7) oder [ULID](https://github.com/ulid/spec) stellen eine sortierbar kodierte Form des Erstellungszeitstempels voran

- [Content-Addressed Storage (CAS)](https://de.wikipedia.org/wiki/Content-Addressed_Storage): Bezeichnung eines Datenstroms durch seine kryptografische Prüfsumme (siehe STP004)
    - z.B. Identifikation [dieses Bildes](https://dl.xyrillian.de/noises/stp-038-intro.png) nicht durch einen Dateinamen/Pfad/URL wie `https://dl.xyrillian.de/noises/stp-038-intro.png`, sondern durch die Prüfsumme `sha256:e68ae45acddfab6198c01ba552ca9f1b76e78bc12c4a4ca0476da27d688d8238`
    - Bildung von Datenstrukturen (z.B. Ordner oder Lesezeichenlisten) durch Benutzung dieser Prüfsummen-IDs analog zu Inode-Nummern (siehe STP031)
    - möglicherweise noch menschenlesbare Identifier, die auf CAS-Identifier abbilden (z.B. in [Git](https://c3d2.de/news/pentaradio24-20210727.html): Tags als menschenlesbare Namen für Commits)
    - Vorteil: bei sorgfältigem Design selbsttätige Deduplikation und Integritätsprüfung möglich
    - Nachteil: Prüfsummen sind nochmal deutlich sperriger als UUIDs

- zur Bezeichnung von Software-Versionen: [Versionsnummern](https://de.wikipedia.org/w/index.php?title=Versionsnummer&oldid=239964497)
    - konventionelles Format: Zahlen mit Punkten getrennt (z.B. `1.7.10`), sortiert leider nicht lexikografisch (z.B. `"1.7.10" > "1.10.7"`)
    - kein fester Formatstandard, aber es gibt Ansätze: z.B. [Semantic Versioning](https://semver.org/lang/de/)
    - Alternativschema: jahreszahlbasiert (z.B. Ubuntu 24.06 ist die im Juni 2024 veröffentlichte Version) -> Bemerkung von ttimeless: kann man auch mit Dateien machen (z.B. `2024-03-01-ticket-bahn.pdf`)
    - Alternativschema: das Marketing bestimmt (z.B. XBox -> XBox 360 -> XBox One -> XBox Series X/S)
    - Alternativschema: Konvergenz (die Versionsnummer von TeX nähert sich der Zahl π an) -> bräuchten wir vielleicht mehr davon :)

#### Feedback zu STP049

Till schreibt:

> Hallo Xyrill,
>
> dies bezieht sich auf Schlüsseltechnologie #49 bzw auf die darin erwähnten Farb-Begriffe.
>
> Rot vs blau kommt ursprünglich vom Militär, ich selbst kenn es von den Pfadfindern, damals bei Sommerlagern und/oder Geländespielen. Um ein bisschen Spannung in das Lagerleben zu kriegen (wovon es unter 200 pubertierenden Jünglingen offenbar nicht genug gab), wurden wir in zwei Unterlager aufgeteilt, die in Wettbewerb gingen. Am letzten Abend wurden die Sieger nach Punkten verkündet, hurra. Bei Geländespielen wurde oft eine Burg definiert, gehalten von Blau und attackiert von Rot. Hier die Ursprünge: $LINK

Xyrill ersetzt hiermit den bereitgestellten Link auf einen Personalagentur-Blog durch [den ausführlicheren Wikipedia-Artikel](https://en.wikipedia.org/wiki/Red_team).

> Die schwarzen und weißen Hüte kommen aus der Urform des Westernfilms, in der die guten Jungs immer weiße Hüte trugen und die Schurken schwarze. Zuletzt sehr pointiert in der ersten Staffel von "Westworld", die ich außerordentlich stark empfehlen möchte. (Steht bei euch popkulturell ignoranten Cyberspezls immer im Schatten von "Black Mirror".) <https://en.wikipedia.org/wiki/Black_and_white_hat_symbolism_in_film>
>
> re: Wurm - [Wiki sagt](https://en.wikipedia.org/wiki/Computer_worm#History), die Bezeichnung gehe zurück auf einen SF-Roman
>
> Schlapphüte (Farbe des Huts zweitrangig) sind geheimdienstlich aktive Personen. Vorkommen: ausschließlich in journalistisch schlampigen Kontexten. Das ist eine dieser Dummphrasen für alle Schreiber, die in Deutsch gelernt haben, dass man sich schöne Synonyme ausdenken soll, damit der Text eleganter wird. Wird er nicht, im Gegenteil fettet er zu, aber das ist wie nem Ochs ins Horn gepetzt. <https://de.wikipedia.org/wiki/Schlapphut#Sonstiges>
>
> Wo wir grade dabei sind, off-topic mein Lieblings-Hassobjekt: "soziale Medien". Nichts an diesen Plattformen ist sozial, Medien sind sie auch nicht, sondern deren Inhalt. Wie meistens bei Übersetzungen im Cyberbereich war die erstübersetzende Person des Englischen nicht mächtig und hat uns für immer mit einem False Friend gestraft. Das originale "social media" hat zwar ein ebenfalls dubioses Medium, aber der Begriff "social" bedeutet in diesem Fall "gesellig": Let's go socialising! (dt. "Gehn wir einen saufen").  Entsprechend sind die sozialen Medien in Wirklichkeit Geselligkeitsplattformen. Nörgel nörgel.
>
> Warum die doofen Trojaner sich ein Hottepferdchen voll waffenstarrender Griechen in die Stadt holten, schildert Stephen Fry in seiner sehr lesenswerten Ilias-Nacherzählung "Troy": Erstens war das Pferd ein exquisites Stück Handwerkskunst, zweitens war es überzogen mit Gold und Silber und Edelsteinen, drittens stand es unter dem Schutz einer mächtigen Göttin, Stichwort Laokoon, vor allem aber ergänzten diese verschlagenen Griechen das Pferd mit einem False-Flag-Helfershelfer, der alle Zweifel (und Cassandras Verzweiflungsschreie) beiseite wischte. Die Operation wäre einer Leningrader Trollfabrik würdig.
>
> Stuxnet: <https://darknetdiaries.com/episode/29/>

Wenn wir schon Podcasts zu Stuxnet verlinken, erhöht Xyrill um [ein Alternativlos](http://alternativlos.org/5/) und möchte sehen.

#### Feedback zu STP057

Klaus schlägt per Mail einen anderen Namen für Matrix/Mastodon-Server vor:

> Hallo Xyrill!
>
> Wie wär's mit 'Meldestelle' oder 'Anmeldedienst' oder etwas in der Art.
>
> Wenn man irgendwo neu hinzieht, muss man sich doch auch registrieren.
>
> Das ist hier nicht anders, schließlich ist man neu auf den entsprechenden Diensten.
>
> Gibt ja jede Menge Meldestellen in Deutschland, so auch im Fediverse, nur dass im Fediverse nicht vorgeschrieben wird, welche man zu nehmen hat.
>
> In der Regel interagiert man ja auch nicht notwendigerweise mit den Leuten, die sich bei der gleichen Meldestelle angemeldet habe.
>
> Grüße aus Belgien!

#### Mehr Feedback zu STP057

Jan schreibt:

> Hallo Xyrill und ttimeless,
>
> da ihr in der letzten Folge zu ActivityPub wiederholt nach Feedback gefragt habt, gebe ich doch auch gerne meinen Senf dazu!
>
> Ich denke das Wort, das Xyrill bei 00:14:53 sucht ist Galionsfigur. Wikipedia nennt im Deutschen Artikel von Figurehead auch "Frühstücksdirektor".
>
> Das war eigentlich schon alles was ich schreiben wollte. Ich danke euch vielmals für diesen tollen Podcast und freue mich alle
drei Wochen(!) über eine neue Folge.

Xyrill merkt hierzu an, dass er nicht mehr wirklich GPG verwendet und das Entschlüsseln dieser Mail ein Abenteuer war.

#### Cold Call

Xyrill hat eine Mail bekommen:

> Hallo Herr Majewsky,

Das sieht schon nicht gut aus.

> ich möchte Ihnen gerne einen spannenden Gast für Ihren Podcast vorschlagen

Haben wir ja grundsätzlich nichts dagegen. Ladet Euch bitte ein, wenn Ihr ein Thema mitbringen möchtet, zu dem Ihr was sagen könnt.

> Mit freundlichen Grüßen,
> $NAME

> Communications Executive
> $PR\_AGENTUR

Aber wenn Ihr Eure PR-Agentur vorschickt, dann gehe ich davon aus, dass es sich hier um Native Advertising handelt. Nicht nur mag ich das nicht, sondern vielleicht denkt dann das Finanzamt noch, dass hier Gewinnabsicht bestünde.
