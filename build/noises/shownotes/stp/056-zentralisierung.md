* [RFC (Request for Comments)](https://de.wikipedia.org/w/index.php?title=Request_for_Comments&oldid=237305423): eine Publikationsreihe für zumeist technische Spezifikationen, aber auch Essays betreffend die Fortentwicklung der Internet-Infrastruktur
    * bereits besprochen in STP035 anhand der Tradition der Aprilscherz-RFCs
    * im Übrigen auch eines der anschaulichsten Beispiele für die ursprüngliche Intention von Hypertext: RFCs sind blanke Textdateien, die für die Darstellung als Webdokument automatisiert aufbereitet werden
    * heute etwas ernster, aber hoffentlich ähnlich zugänglich anhand [RFC 9518: "Centralization, Decentralization, and Internet Standards"](https://datatracker.ietf.org/doc/html/rfc9518)
    * wir folgen in unserer Besprechung des Themas der Struktur der Textgrundlage; deutschsprachige Zitate sind jeweils aus dem Englischen übersetzt

* Kernbegriff: "Zentralisierung \[im Sinne dieses Dokumentes] ist eine Sachlage, bei der eine einzelne Entität oder eine kleine Gruppe eine Funktion des Internets überwachen, abfangen, kontrollieren oder aus ihrem Betrieb oder ihrer Benutzung Pacht ziehen kann."
    * Die ersten drei Varianten (Überwachung, Abfangen, Kontrolle) sind die intuitiv klaren Formen von Zentralisierung (z.B. das Telefon-Festnetz ist zentralisiert, da es von einem einzigen Anbieter, der Telekom, betrieben und kontrolliert wird).
    * Der Teil mit "Pacht" ist im Original "rent extraction", was im Englischen ein durchaus geläufiger Begriff ist, während im Deutschen [der "Rentier"](https://de.wiktionary.org/w/index.php?title=Rentier&oldid=9846923#Substantiv,_m) wahrscheinlich eher für einen Rentner gehalten wird.
        * [en.wiktionary.org zu "rent" im Sinne "rent extraction"](https://en.wiktionary.org/w/index.php?title=rent&oldid=77955238): "A profit from possession of a valuable right, as a restricted license to engage in a trade or business." &ndash; "Profit aus dem Besitz eines wertvollen Rechtes, als limitierte Lizenz zum Betrieb eines Handwerks oder Geschäftes"
        * klassische Beispiele außerhalb der IT: Taximedaillon, Notarstellen, natürliche Monopole wie Netzbetreiber von Strom-/Wasser-/Gasleitungen
        * Beispiel im Internet: App-Store-Betreiber (wer Software für iPhones vertreiben will, muss 30% Pacht an Apple abführen), [IANA (Internet Assigned Numbers Authority)](https://de.wikipedia.org/w/index.php?title=Internet_Assigned_Numbers_Authority&oldid=235416916) als Betreiber der DNS-Rootserver (zu DNS siehe STP018)
        * im Gespräch erwähnt: [5 Gründe, warum der Markt in der Praxis nicht so funktioniert wie in der Theorie](https://news.ycombinator.com/item?id=33314970)

* mögliche negative Konsequenzen von Zentralisierung
    * Machtkonzentration (der RFC zitiert hier die Federalist Papers von 1788: "So wie \[John Madison in denselben\] Good Governance \[(verantwortungsbewusste Regierungsführung)] der US-Staaten beschreibt, so erfordert Good Governance des Internets, dass Kontrolle über seine einzelnen Funktionen nicht ohne angemessene Gewaltenteilung an einer zentralen Stelle konsolidiert wird.")
    * Begrenzung von Innovation: neue Entwicklungen brauchen Freiraum
        * im Gespräch erwähnt: [Freak Show #274 zum App-Store-Monopol von Apple](https://freakshow.fm/fs274-was-hat-apple-jemals-nicht-fuer-uns-getan)
    * Begrenzung von Wettbewerb: Wettbewerb mehrerer kompatibler Anbieter oder mehrerer Protokolle schafft die vorgenannte Gewaltenteilung
    * Verringerte Verfügbarkeit: Wenn alles durch einen zentralen Anbieter geht, führt ein Ausfall dort zu Folgeausfällen woanders. 
        * Greifbares Beispiel: Was passiert, wenn ein Stadtviertel über eine zentrale Glasfaserleitung angebunden ist und diese Glasfaser versehentlich aufgebaggert wird?
        * Der RFC zitiert einige Beispiele.
        * Xyrill hat eine Anekdote zu AWS und Minecraft.
    * Monokultur
    * Selbstverstärkung

* mögliche _positive_ Konsequenzen von Zentralisierung
    * Verständlichkeit, Zweckdienlichkeit
        * zentrale Steuerung von DNS verleiht Domain-Namen eine hilfreiche Universalität
        * z.B. bei IP-Adressen oder Telefonnummern ist derartig kontrollierte Universalität technisch notwendig
    * Effizienz durch Skaleneffekte: siehe Besprechung von CDNs in STP024 (Cloud-Computing)
    * Konzentration von Kontrolle kann Governance einfacher machen (was positiv sein kann, sofern man Governance an sich als grundsätzlich wünschenswert ansieht)
        * Beispiel: Bankenregulierung
    * Xyrill fügt noch hinzu: Zentralisierung kann die Weiterentwicklung von Protokollen vereinfachen, sofern Ossifikation entgegengewirkt werden muss
        * [Ossifikation](https://de.wikipedia.org/w/index.php?title=Ossifikation&oldid=239981566): im Wortsinne "Verknöcherung", im übertragenen Sinne die faktische Unänderbarkeit von Protokollen, wenn eine Vielzahl von verschiedenen unabhängigen Akteuren die Änderung gleichzeitig vollziehen muss

* Ansätze zur aktiven Dezentralisierung
    * Föderation: verschiedene Anbieter betreiben voneinander unabhängige Instanzen, die miteinander in Kontakt treten, um einen insgesamten Dienst zu bilden (klassisches Beispiel: E-Mail)
        * alleine unzureichend, da nicht-technische Faktoren Zentralisierungsdruck erzeugen
        * Beispiel E-Mail: wer einen E-Mail-Server betreiben will, muss faktisch nach der Pfeife von Microsoft/Google/United Internet/etc. tanzen, da Outlook und G-Mail den Markt dominieren
        * Beispiel XMPP: nachdem einige wenige Betreiber in der Föderation groß geworden sind, haben sie sich vom föderierten Netz getrennt und ihre Nutzer effektiv eingesperrt (z.B. WhatsApp war früher XMPP-kompatibel)
    * verteilter Konsens: Zusammenschluss aus vielen Teilnehmenden, bei denen die Ausübung von Macht innerhalb des Netzes zufällig und temporär zugewiesen wird (siehe Besprechung von Blockchain in STP048)
        * "Während diese Maßnahmen darin wirkungsvoll sein können, den Betrieb eines Dienstes zu dezentralisieren, können andere Aspekte seiner Vorhaltung weiterhin zentralisiert sein: zum Beispiel die Kontrolle über das Design, die gemeinsam genutzte Implementation, sowie die Dokumentation von Netzwerkprotokollen."
    * aktives Regierungshandeln ("Operational Governance"): Ausübung der operativen Kontrolle durch eine neutrale Treuhänderstelle, die die Interessen der verschiedenen Diensteanbietern und Nutzenden gegeneinander abwägt
        * Einrichtung mitunter auch im Nachhinein möglich, sofern das technische, ökonomische oder politische Umfeld die Diensteanbieter zur Kooperation zwingt
        * klassische Beispiele: [ICANN](https://de.wikipedia.org/w/index.php?title=Internet_Corporation_for_Assigned_Names_and_Numbers&oldid=238717925) für IP-Adressen, IANA für DNS, CA/Browser-Forum für Wurzel-Zertifikatsautoritäten (siehe STP048)

* Was können Autorinnen von Internet-Standards tun, um Zentralisierung vorzubeugen und Dezentralisierung zu fördern?
    * Legitimität erzeugen
        * durch Öffentlichkeitsarbeit gegenüber der Bevölkerung oder gesellschaftlichen und staatlichen Institutionen Vertrauen schaffen, um dann die öffentliche Diskussion in die gewünschte Richtung zu beeinflussen
        * Problem: mehr Geld kauft mehr PR
    * Diskussion fokussieren
        * nicht Zentralisierung als solches verteufeln (hat wie gesehen ja auch Vorteile), sondern explizit auf die Nachteile konzentrieren
        * konkreter Vorschlag: RFCs sollten neben den Standardabschnitten "Sicherheitsrelevante Überlegungen" und "IANA-relevante Überlegungen" immer auch "Zentralisierungsrelevante Überlegungen" explizieren
    * Abzielen auf proprietäre Dienste
        * zentralisierte Angebote und Protokolle sollten aktiv mit offenen Standards konfrontiert werden
    * Wechselkosten verringern
        * z.B. neue offene Standards so gestalten, dass sie mit bestehenden proprietären Diensten interoperabel sein können
        * selbst Open Source ist nicht immer die Lösung, wenn die Weiterentwicklung der Open-Source-Software die Expertise eines bestimmten Anbieters erfordert
    * Delegation von Kontrolle bändigen
        * Wenn ein Protokoll der Nutzerin ermöglicht, Aufgaben an einen Dritten zu delegieren, kann das eine Sollbruchstelle zur Machtkonzentration sein.
        * Beispiel (im RFC erwähnt): HTTP-Proxies (ähnliches war auch bei HTTP auf der Internetanbieter-Ebene zu beobachten)
    * starke Isolation durchsetzen
        * Beispiel: viele Protokolle hängen von der Verfügbarkeit und Zuverlässigkeit eines darunterliegenden Netzwerkes ab -> Verschlüsselung schützt gegen Machtkonzentration beim Betreiber dieser Basisschicht (schützt in diesem Beispiel nur vor Zentralisierung im Sinne von Überwachung, weniger im Sinne von Kontrolle)
    * sorgfältige Abwägung von Erweiterbarkeit
        * hilfreich, um Wettbewerb und Innovation zwischen verschiedenen Anbietern zu fördern
        * aber auch Risiko des "Vendor Lock-In" (Erweiterung durch anbieterspezifische Funktionen, die Nutzerinnen an den Anbieter binden)
    * wiederverwenden, was funktioniert
        * z.B. hat DNS trotz aller Probleme alles in allem ein relativ dezentrales Design und eine relativ gute Governance &ndash; tendenziell lieber darauf aufbauen, als eine eigene dezentrale Datenbank zu frickeln (oder `s/DNS/Ethereum/`, wenn man lieber Blockchain nehmen möchte)

* Xyrills Abendgedanken: (De-)Zentralisierung ist, so wie fast alle Fragen aus dem Feld der Netzpolitik, nicht eine rein technische Frage und erfordert neben technischen auch soziale und gesetzliche Lösungen
