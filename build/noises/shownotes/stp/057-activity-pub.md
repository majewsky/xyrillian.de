Klaus schreibt per Mail:

> Hallo timeless und Xyrill,
>
> bin Anfang des Jahres irgendwo auf euren Podcast gestoßen und dort eingestiegen, wo ihr damals gerade wart, Folge 49 Schadcode und dann die weiteren.
>
> Diese haben mir so gut gefallen, dass ich alle Folgen mit Antennapod abgeholt habe und dabei bin euren Podcast von hinten aufzurollen.
>
> Großes Kompliment an Euch! Phantastische Serie! Freue mich auf jede neue und für mich neue-alte Folge.
>
> Bin zwar vom Fach, aber lerne immer was Neues oder andere Blickwinkel kennen. HERRLICH!
>
> ------------------------------------------------------------------------
>
> Eine kleine Bemerkung von mir zu den Folgen 13 und 14, Textkodierung und Textdarstellung.
>
> Ich bevorzuge so wie Xyrill auch die Verwendung von kombinierenden diakritischen Zeichen.
>
> Die meisten der für mich nützlichen davon (u.a. für DE, FR, NL, CN-Pinyin) [beginnen an der Hex-Adresse U+0300 ff.](https://en.wikipedia.org/w/index.php?title=Combining_character&oldid=1226621587)
>
> Leicht zu merken und praktisch, wenn man ständig mit anderen Tastaturbelegungen zu tun hat.
>
> Worauf ich aber hinaus möchte, ist folgendes:
>
> Diese diakritischen Zeichen werden sowohl am Schirm als auch im Druck gleich dargestellt, obwohl ihnen andere Byte-Folgen zugrunde liegen. Hat Xyrill alles beschrieben. Hab ja zugehört. Für Textprogramme, etc. ist das kein Problem, in Datenbanken aber schon.
>
> Es macht dann leider doch einen gravierenden Unterschied, ob diakritische Zeichen wie en "ä" aus U+00E4 oder aus der Kombination von U+0061 + U+0308 entstehen.
>
> * Tastaturen, die diese Zeichen direkt kodieren, verwenden nur die Version mit 1 Codepunkt.
> * Bei Längenbestimmungen von Zeichenketten ergeben sich Unterschiede. Die Zeichenketten sehen gleich aus, sind aber nicht gleich lang, obwohl sie auch gleich lang aussehen!
> * Eingabeformulare fangen das nicht immer gut ab. Gleich aussehender Begriff, aber doch nicht derselbe.
> * Idem bei Suchanfragen.
> * etc.
>
> Ich empfehle daher immer in Datenbanken Strings so zu standardisieren, dass nur die Version mit 1 Codepunkt verwendet wird, so wie eben Muttersprachler diesen Text auf ihrer sprachspezifischen Tastatur eingeben würden.
>
> Viele Grüße aus Belgien!
>
> Klaus

Nun zum eigentlichen Thema:

- Ausgangssituation
    - [Soziale Medien](https://de.wikipedia.org/w/index.php?title=Soziale_Medien&oldid=242703413): "digitale Medien bzw. Plattformen, die es Nutzern ermöglichen, sich im Internet zu vernetzen, \[...] untereinander auszutauschen und mediale Inhalte einzeln, in einer definierten Gemeinschaft oder offen in der Gesellschaft zu erstellen, zu diskutieren und weiterzugeben"
    - aber: alles zentralisiert (siehe STP056) und nicht interoperabel (siehe STP055)
    - [ActivityPub](https://de.wikipedia.org/w/index.php?title=ActivityPub&oldid=240310299): der zurzeit wahrscheinlich aussichtsreichste Versuch, soziale Medien dezentral zu modellieren
    - erste Anfänge in den frühen 2010ern mit [Activity Streams 1.0](https://activitystrea.ms/specs/atom/1.0/), dann in 2016-2018 Entwicklung von [Activity Streams 2.0](https://www.w3.org/TR/activitystreams-core/) und darauf aufbauend [ActivityPub](https://www.w3.org/TR/activitypub/)
    - Xyrill sieht zwei wesentliche Design-Einflüsse im ActivityPub-Protokoll

- Einfluss: [REST (Representational State Transfer)](https://de.wikipedia.org/w/index.php?title=Representational_State_Transfer&oldid=241086067)
    - HTTP kennt man allgemein als Protokoll zum Abruf von Webseiten (siehe STP002)
    - neben `GET` gibt es in HTTP-Anfragen aber noch viele andere Verben
    - `POST`: Ursprung in HTML für Abschicken ausgefüllter Formulare
    - `PUT`, `PATCH`, `DELETE` etc.: für Hochladen und Verwalten von Dateien
    - REST: Modellierung von APIs unter Einsatz dieser HTTP-Verben und unter Modellierung der Datensätze entlang von URLs (z.B. in einem Webshop `PATCH /products/23`, um Produkt Nr. 23 zu bearbeiten; oder in einem Blog `GET /articles/hello-world/comments`, um die Kommentare zu dem Post namens "Hello World" aufzurufen)

- Einfluss: [E-Mail](https://de.wikipedia.org/w/index.php?title=E-Mail&oldid=243190279)
    - entworfen in einer Zeit, als die meisten Verknüpfungen im Internet nicht durchgehend verfügbar waren
    - Mails werden nicht sofort an den Empfänger zugestellt, sondern landen erstmal auf dem Mail-Server des Absenders in einer Outbox (analog dem Postbriefkasten)
    - Zustellung an den jeweils nächsten Server in der Kette bei Verfügbarkeit der Leitung (heutzutage unverzüglich und immer direkt an den Zielserver), dann Ablage in der Inbox des Empfängers (analog dem Hausbriefkasten)
    - Problem bei ActivityPub: E-Mail ist Unicast (Nachrichten gehen immer an eine bestimmte Adresse bzw. eine bestimmte Liste von Adressen), aber soziale Medien können auch Broadcast sein (z.B. Tweets oder Blogposts sind global sichtbar)

- ActivityPub: Inbox und Outbox wie bei E-Mail, aber bidirektional
    - Direktnachricht von Alice an Bob: Alice postet in ihre eigene Outbox, Server leitet weiter in Bobs Inbox, Bob ruft aus seiner Inbox ab
    - öffentliche Nachricht von Alice: Alice postet in ihre eigene Outbox, Bob ruft aus Alices Outbox ab; bei Followern sendet Alices Server in deren Inbox
    - Nachrichten verbleiben bis zu ihrer eventuellen Löschung in der Outbox und bilden einen Feed aller Nachrichten dieser Person
    - REST-Einfluss: Inbox erlaubt `GET` durch den Inhaber und `POST` durch andere Server, Outbox erlaubt `POST` durch den Inhaber und `GET` durch alle
    - vgl. Bilder in [Kapitel 1 der ActivityPub-Spezifikation](https://www.w3.org/TR/activitypub/)
    - wesentliche Implementationen von ActivityPub: Kurznachrichtendienst [Mastodon](https://joinmastodon.org/), Linkaggregator [Lemmy](https://join-lemmy.org/), Bildplattform [Pixelfed](https://pixelfed.org/), Videoplattform [PeerTube](https://joinpeertube.org/)
    - Anekdote: Xyrill hatte überlegt, die Podcast-Seite mit ActivityPub auszustatten, und ist davon wieder abgekommen

- Beobachtung bzw. Rückbezug zu STP056: föderierte Netzwerke haben einen starken Drang zur (Re-)Zentralisierung
    - Anschauungsbeispiel: [Verteilung der Nutzerzahlen zwischen verschiedenen ActivityPub-Implementationen](https://en.wikipedia.org/w/index.php?title=ActivityPub&oldid=1212815451#Software_using_ActivityPub) -> Mastodon macht 99% der Nutzerbasis aus
    - innerhalb von Mastodon initial starke Konzentration bei `mastodon.social`; seit einiger Zeit ist beim Anlegen eines Nutzerkontos kein Server mehr voreingestellt
    - vgl. auch Matrix: schätzungsweise etwa 50% der Nutzer bei `matrix.org`
    - absolut subjektiver Eindruck: Mastodon-Entwickler dominieren die Weiterentwicklung von ActivityPub als ganzes

- Vergleich mit dem AT-Protokoll von BlueSky
    - [Motivation](https://web.archive.org/web/20200111021337/twitter.com/jack/status/1204766078468911106): Trennung des Zustellpfades von "dem Algorithmus" (Inhaltsvorschläge und Moderation), um letztere austauschbar zu gestalten
    - modulare/dezentrale Moderation (mit den daraus folgenden Konsequenzen, z.B. öffentlich einsehbare Blocklisten, also kein Shadowbanning möglich)
        - vgl. ActivityPub: Faustregel: 1 Moderator:in pro 1000 User
    - Konten bei AT sind nicht an bestimmte Server gebunden, sondern können unter Erhalt der gesamten Kommunikationshistorie umgezogen werden (ActivityPub-Adressen sind mehr wie Kontonummern oder E-Mail-Adressen; AT-Identifier sind mehr wie portierbare Telefonnummern)
        - vgl. Debatte über die Begrifflichkeit "Server" vs. "Instanz" vs. "Community" bei Mastodon
        - Update: [ActivityPub arbeitet an "nomadischer Identität"](https://wedistribute.org/2024/03/activitypub-nomadic-identity/) analog zu den Decentralized Identifiers in AT

- im Gespräch erwähnt: [Entwurf einer Änderung der Gemeinsamen Geschäftsordnung der Bundesministerien (GGO) zur Umsetzung der Vorhaben "Exekutiver Fußabdruck" und "Synopse"](https://www.bmi.bund.de/SharedDocs/downloads/DE/veroeffentlichungen/2024/ggo-fussabdruck-synopsenpflicht.pdf?__blob=publicationFile&v=1)

