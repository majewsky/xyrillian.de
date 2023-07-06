- Programmcode ist ein immaterielles Wirtschaftsgut -> jeder Programmierer sollte ein Grundverständnis für **Immaterialgüterrecht** haben (neben Urheberrecht auch Patentrecht, Markenrecht, etc.)

- obligatorische Vorbemerkungen
    - **Wir sind keine Rechtsanwälte. Das hier ist keine Rechtsberatung.**
    - Fokus auf das deutsche Recht (insbesondere nicht auf das Copyright der englischen Rechtstradition)
    - insbesondere Fokus auf das [UrhG](https://www.gesetze-im-internet.de/urhg/), daneben noch zum Beispiel das [KunstUrhG](https://www.gesetze-im-internet.de/kunsturhg/) (Recht am eigenen Bild) oder das [ArbnErfG](https://www.gesetze-im-internet.de/arbnerfg/) (Arbeitnehmererfindungen)

- erste Facette: **Urheberpersönlichkeitsrecht**
    - [&sect; 12 UrhG](https://www.gesetze-im-internet.de/urhg/__12.html): "Der Urheber hat das Recht zu bestimmen, ob und wie sein Werk zu veröffentlichen ist."
    - [&sect; 13 UrhG](https://www.gesetze-im-internet.de/urhg/__13.html): "Der Urheber hat das Recht auf Anerkennung seiner Urheberschaft am Werk." (aber keine Pflicht)
    - [&sect; 14 UrhG](https://www.gesetze-im-internet.de/urhg/__14.html): "Der Urheber hat das Recht, eine Entstellung oder eine andere Beeinträchtigung seines Werkes zu verbieten, die geeignet ist, seine berechtigten geistigen oder persönlichen Interessen am Werk zu gefährden."
    - im Allgemeinen nicht übertragbar, aber insb. für Computerprogramme [&sect; 69b UrhG](https://www.gesetze-im-internet.de/urhg/__69b.html): "Wird ein Computerprogramm von einem Arbeitnehmer in Wahrnehmung seiner Aufgaben oder nach den Anweisungen seines Arbeitgebers geschaffen, so ist ausschließlich der Arbeitgeber zur Ausübung aller vermögensrechtlichen Befugnisse an dem Computerprogramm berechtigt, sofern nichts anderes vereinbart ist."

- zweite Facette: **Verwertungsrecht** in [&sect;&sect; 15 ff. UrhG](https://www.gesetze-im-internet.de/urhg/BJNR012730965.html#BJNR012730965BJNG000801377)
    - Recht auf Vervielfältigung, Verbreitung, Ausstellung, Vortrag/Aufführung/Vorführung, Sendung, Bearbeitung/Umgestaltung, Verfilmung, etc.
    - meistens Übertragung dieser Rechte, zum Beispiel bei Büchern auf Verlage oder bei Filmen auf Verleiher

- Schranken des Urheberrechts: [&sect;&sect; 44a ff. UrhG](https://www.gesetze-im-internet.de/urhg/BJNR012730965.html#BJNR012730965BJNG001302123)
    - Rechtspflege, Barrierefreiheit, Gebrauch bei religiösen Feierlichkeiten, Schulfunksendungen, Verbreitung öffentlicher Reden, Berichterstattung über Tagesereignisse, Zitate, Karikaturen/Parodien, [Privatkopien](https://www.gesetze-im-internet.de/urhg/__53.html)
    - insb. für Programme [&sect; 69e UrhG](https://www.gesetze-im-internet.de/urhg/__69e.html): "Die Zustimmung des Rechtsinhabers ist nicht erforderlich, wenn die Vervielfältigung des Codes oder die Übersetzung der Codeform \[...] unerläßlich \[sic] ist, um die erforderlichen Informationen zur Herstellung der Interoperabilität eines unabhängig geschaffenen Computerprogramms mit anderen Programmen zu erhalten, sofern folgende Bedingungen erfüllt sind: \[...]"

- [Freie Software](https://de.wikipedia.org/w/index.php?title=Freie_Software&oldid=230150686), [Open-Source-Bewegung](https://de.wikipedia.org/w/index.php?title=Open_Source&oldid=232516030)
    - Urkonzept im englischsprachigen Copyright: **Public Domain** (willentliche Abtretung der Urheberrechte, im deutschen Urheberrecht nicht möglich)
    - ab Beginn der Informatik: in der Forschung Behandlung von Programmcode wie wissenschaftliche Erkenntnisse, ohne Anwendung des Urheberrechts
    - ab 1980er Jahre: Freie-Software-Bewegung -> subversive Nutzung des Copyright nicht zur Beschränkung, sondern zur Maximierung von Nutzbarkeit
    - "Vier Grundfreiheiten": Use, Study, Share, Modify -> **Copyleft-Modell** (Urheber zwingt alle Nutzer, auch bei Folgenutzungen die vier Grundfreiheiten zu berücksichtigen)
    - Gegenbewegung: "Open Source" -> schwächere Definition, quelloffene Software ohne Copyleft kann auch in unfreier Software verwendet werden
    - Randbemerkung: diese Open-Source-Gemeinde hat bis auf den Namen nichts mit der [OSINT-Community](https://de.wikipedia.org/w/index.php?title=Open_Source_Intelligence&oldid=231133805) zu tun

- praktischer Hinweis für Programmierer: Schaut auf eure SBOM!
    - BOM = Bill of Material (zu deutsch: [Stückliste/Materialliste](https://de.wikipedia.org/w/index.php?title=St%C3%BCckliste&oldid=228719182))
    - SBOM = Stückliste für Software (Welche Programme oder Programmbibliotheken werden als Teil meines Programms mit ausgeliefert?), [Symbolbild](https://xkcd.com/2347/)
    - Teil der SBOM sind die Lizenzinformationen der verwendeten Programme und Bibliotheken
    - SBOM ist auch sonst eine gute Idee; kann man auch zum automatisierten Suchen nach Sicherheitslücken verwenden

- netzpolitische Aspekte
    - [&sect; 95a UrhG](https://www.gesetze-im-internet.de/urhg/__95a.html): der Kopierschutzparagraf, "Wirksame technische Maßnahmen zum Schutz eines \[...] Werkes \[...] dürfen ohne Zustimmung des Rechtsinhabers nicht umgangen werden, soweit \[...] die Umgehung erfolgt, um den Zugang zu einem solchen Werk \[oder dessen] Nutzung zu ermöglichen."
    - [&sect; 64 UrhG](https://www.gesetze-im-internet.de/urhg/__64.html): "Das Urheberrecht erlischt \[im Allgemeinen\] siebzig Jahre nach dem Tode des Urhebers."
        - [Studie: Optimale Dauer des Urheberrechtsschutzes wäre 14 Jahre](https://netzpolitik.org/2007/optimale-dauer-des-urheberrechtsschutzes-14-jahre/); [Link zur Originalstudie](https://rufuspollock.com/papers/optimal_copyright_term.pdf)
    - [&sect; 44b UrhG](https://www.gesetze-im-internet.de/urhg/__44b.html): der Data-Mining-Paragraf, "Zulässig sind Vervielfältigungen von rechtmäßig zugänglichen Werken für das Text und Data Mining. Die Vervielfältigungen sind zu löschen, wenn sie für das Text und Data Mining nicht mehr erforderlich sind. \[...] Ein Nutzungsvorbehalt bei online zugänglichen Werken ist nur dann wirksam, wenn er in maschinenlesbarer Form erfolgt."; siehe dazu auch [Pentaradio vom August 2022](https://c3d2.de/news/pentaradio24-20220823.html)

- im Gespräch erwähnt
    - [GNU-Projekt](https://de.wikipedia.org/w/index.php?title=GNU-Projekt&oldid=205119299)
    - [Postmortales Persönlichkeitsrecht](https://de.wikipedia.org/w/index.php?title=Postmortales_Pers%C3%B6nlichkeitsrecht&oldid=229281968)
