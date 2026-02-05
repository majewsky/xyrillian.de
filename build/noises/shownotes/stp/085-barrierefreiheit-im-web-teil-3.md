- siehe STP081: gemäß des neuen [Barrierefreiheitsstärkungsgesetzes](https://de.wikipedia.org/w/index.php?title=Barrierefreiheitsst%C3%A4rkungsgesetz&oldid=258239575) müssen Webseiten und Webapplikationen die AA-Kriterien der WCAG 2.1 erfüllen
    - WCAG = [Web Content Accessibility Guidelines](https://de.wikipedia.org/w/index.php?title=Web_Content_Accessibility_Guidelines&oldid=254032297) (Webinhalte-Barrierefreiheits-Richtlinien) der [Web Accessibility Initiative](https://www.w3.org/WAI/) des W3C
    - Kriterien sind gruppiert in 4 Grundprinzipien und darin 13 Richtlinien
    - Kriterien gehen jeweils von A (Pflicht) über AA (Kür) bis AAA (Sahnehäubchen)
    - in STP081 und STP083: das Grundprinzip 1 (Wahrnehmbarkeit) und 2 (Bedienbarkeit) mit insgesamt neun Richtlinien, heute der Rest
    - zum Mitlesen: [WCAG 2.2 (aktuelle Version, nur Englisch)](https://www.w3.org/TR/WCAG22/) oder [WCAG 2.0 (etwas ältere Version, in offizieller deutscher Übersetzung)](https://www.w3.org/Translations/WCAG20-de/)

- Grundprinzip 3: **Verständlichkeit**
    - Richtlinie 10: *Lesbarkeit*
        - A: Sprache der Webseite muss programmatisch auslesbar sein
        - AA: Sprache jedes Teils der Webseite muss programmatisch auslesbar sein
        - AAA: für ungewöhnliche Ausdrücke, Fachwörter und Abkürzungen soll eine Definition zugänglich sein; für komplizierte Texte (alles über Hauptschulniveau) soll eine Alternativform in einfacher Sprache vorliegen; Aussprachehinweise sollen zugänglich sein, sofern die Bedeutung eines geschriebenen Wortes von der Aussprache abhängt (z.B. "umfahren" vs. "umfahren")
    - Richtlinie 11: *Vorhersehbarkeit*
        - A: Bedienung von Steuerelementen soll nicht unvorhersehbar eine Seitennavigation oder dergleichen auslösen; Hilfemechanismen (inkl. Kontaktinformationen) müssen (sofern vorhanden) auf jeder Unterseite gleich funktionieren
        - AA: Navigationsmechanismen und wiederholte Steuerelemente müssen auf jeder Unterseite gleich erscheinen und funktionieren
        - AAA: Navigation nur aufgrund von expliziter Anforderung durch die Nutzer:in
    - Richtlinie 12: *Hilfestellung bei der Eingabe*
        - A: fehlerhafte Eingaben müssen klar hervorgehoben werden; benötigte Eingaben müssen klar als solche gekennzeichnet sein; Daten sollen nicht mehrmals abgefragt werden (beim zweiten Mal entweder automatisches Ausfüllen oder eine Option zur Datenübernahme), außer wenn Sicherheit es erfordert oder die vorliegenden Daten veraltet sind
        - AA: bei fehlerhaften Eingaben sollen Korrekturvorschläge gegeben werden (soweit möglich; z.B. nicht, wenn Geheimniswahrung gefährdet wäre); rechtlich verbindliche Vorgänge sollen entweder umkehrbar sein oder einen separaten Bestätigungsschritt anbieten; Authentifizierung soll keine Tests für geistige Leistungsfähigkeit erfordern (z.B. für [Captcha](https://de.wikipedia.org/w/index.php?title=Captcha&oldid=257160068) muss eine alternative Methode existieren, außer wenn der kognitive Test nur Bilderkennung oder Wiedererkennung vergangener Eingaben umfasst; Passwort-Manager-Blocker sind verboten)
        - AAA: kontextsensitive Hilfe bei jeder Eingabe; jegliche formularbasierte Eingabe soll entweder umkehrbar sein oder einen separaten Bestätigungsschritt anbieten; Komplettverbot von Captchas ohne Alternativmethode, die keinen Test für geistige Leistungsfähigkeit umfasst

- Grundprinzip 4: **Robustheit**
    - Richtlinie 13: *Kompatibilität*
        - A: für alle Steuerelemente sind der Name und die Rolle programmatisch auslesbar (z.B. mittels [ARIA](https://de.wikipedia.org/w/index.php?title=Accessible_Rich_Internet_Applications&oldid=254032608) (*Accessible Rich Internet Applications*, barrierefreie reichhaltige Internet-Anwendungen))
        - AA: für alle Steuerelemente sind der aktuelle Status bzw. eventuelle Statusnachrichten programmatisch auslesbar (z.B. durch Screenreader)

## Feedback zu STP003

Dominic schreibt Themenvorschläge, und außerdem:

> Vor kurzem ist mir [ein Video untergekommen, das die unterschiedlichen Gatter sehr anschaulich zeigt](https://www.youtube.com/watch?v=_Pqfjer8-O4) und das ich hiermit empfehlen möchte. (Allerdings empfehle ich das Englische original denn die deutsche Synchronisation ist eine ✨KI✨-Stimme)
>
> macht weiter so und vielen Dank

## Feedback zu STP084

Daniel schreibt:

> Hallo Xyrill, hallo `stell es dir in einer unbekannten Fremdsprache vor` ttimeless,
>
> TL,DR: macht 'ne Folge über alternative Betriebssysteme für Smartphones
>
> Ich habe euren Call for Action aus der Episode 84 des STP gehört. Bei mir liegt aktuell das Thema `Betriebssystem fürs Smartphone` auf dem Tisch (dazu später mehr) und da habe ich mich gefragt, ob ihr dazu nicht eine Folge oder einen Teil einer Folge machen könnt.
>
> Xyrill sagte ja mal, er wechselt alle paar Jahre zwischen Android und iOS. ttimeless hat sich hierzu, glaube ich, noch nicht geäußert (Was mit Holz? *Zwinkersmiley*).
>
> Es gibt ja inzwischen einige Custom ROMs auf dem Markt und einige Projekte sind auch schon wieder eingestampft (dazu später mehr).
>
> Ich glaube, GrapheneOS wurde bei euch schon mal angesprochen, mindestens im Pentaradio. Das ist ja bekannt für seine Security Features  und Einstellmöglichkeiten, hat aber den klitzekleinen Nachteil, das es nur auf einer Hardware läuft (Google Pixel).
>
> Ubuntu Touch macht anscheinend auch Fortschritte. Außerdem schon lange verfügbar sind LineageOS und /e/os. Weitere: DivestOS, CalyxOS (dazu später mehr), Pixel Experience, usw, usf.
>
> Ich finde, auch im Zuge der Di.day Initiative lohnt sich ein Blick auf dieses Thema. Auch, wenn es hier ein gewisses Skill Level, die Fähigkeit zum Verzicht und Bock auf Bugs bzw Bastelfreude braucht, um mit einem alternativen Betriebssystem auf dem Smartphone unterwegs zu sein.
>
> Jetzt kommt das "später": Ich habe mir im letzten Jahr ein Fairphone 4 angeschafft und mir CalyxOS draufgeholfen. Nur um kurz nach der Installation und Einrichtung die Nachricht von den Leutis vom Projekt zu erhalten, dass erstmal keine Updates mehr over-the-air ausgeliefert werden. Wie es mit dem Projekt weitergeht, kann ich ihrer Kommunikation leider nicht entnehmen. `¯\_(^^)_/¯`
>
> Ich brauche also bald mal ein neues Betriebssystem und ich dachte, ihr habt bestimmt was Spannendes dazu beizutragen. Sei es aus eigener Erfahrung und/oder aus eurem technikbegeisterten und technikversierten Umfeld.
>
> Der Endgegner in diesem Game ist übrigens die einzige App, die ich nicht durch FOSS ersetzen kann: die Banking-App...
> Es liegt in der Natur der Sache, dass ich diese App nicht wie einen Messanger oder Musikstreamingdienst wechseln kann.
>
> Ich weiß, dass euer Wissensdampfer langsam den Eisberg umschifft (Lob der Sorgfalt) und ihr Zeit für die Produktion neuer Folgen braucht. Selbst wenn dieses Thema erst in einigen Monaten rankommen sollte, freue ich mich zu hören, was ihr dazu zu sagen habt.
>
> Ansonsten noch vielen Dank für viele unterhaltsame und lehrreiche Stunden und einen Podcast, bei dem ich schon einige Folgen  mehrfach gehört habe. Ganz stark.

Der im Gespräch erwähnte Datenspuren-Vortrag ist ["Na endlich: Linux auf dem Smartphone"](https://media.ccc.de/v/ds25-517-na-endlich-linux-auf-dem-smartphone).

## Feedback zu STP081/083

Toni schreibt:

> Lieber Xyrill, lieber ttimeless,
>
> ich wurde heute von einem Kollegen auf euren Podcast aufmerksam gemacht. Ausschlaggebend war eure kleine Episodenreihe zum Thema Barrierefreiheit im Web. Ich bin selbst geburtsblind und schon immer mit den Hürden im Netz konfrontiert, schlussendlich hat mich das dazu geführt, dass ich im Rahmen meiner Promotion derzeit zum Thema Zugänglichmachung von interaktiven visuellen Kollaborationswerkzeugen forsche. Studiert habe ich ursprünglich Spieleentwicklung.
>
> Ihr hattet in eurem Podcast zu Wortmeldungen von Betroffenen aufgefordert. Nun ist das ja schon eine Weile her - ihr meintet im Podcast, ihr nehmt die Folge Ende August auf, das sind ja immerhin schon 4 Monate. Ich finde es super, dass ihr in eurem Podcast über die Barrierefreiheit im Web aufklärt und würde mich noch viel mehr darüber freuen, wenn sie tatsächlich auch die Mehrheit umsetzen oder zumindest die Notwendigkeit dafür einsehen würde - in der Praxis wird man bei Nachfragen tatsächlich immer eher mit Skepsis und Unwillen konfrontiert, selbst aus der eher liberalen und auch der Open-Source-Community. Zudem ist das BFSG auch leider eher ein zahnloser Wolf - es gibt zwar Richtlinien, an die sich technisch jeder halten muss, insbesondere öffentliche Dienste und Dienstleister mit einer gewissen Mindestgröße, aber wo kein Kläger, da kein Richter. Menschen, die von Barrierefreiheit betroffen sind, befinden sich statistisch eher am unteren Ende der Existenzgrenze und machen sich mehr Sorgen um ihr täglich Brot als darüber, Geld in Klageverfahren zu stecken. Und selbst wenn damit Erfolg zu verzeichnen wäre, wäre gerade für große Firmen der Verlust immer noch sehr viel geringer, als das Investment, ihre Webpräsenzen von vornherein umzurüsten.
>
> Mein Ansatz hier ist daher, Barrierefreiheit direkt schon in der Lehre mitzudenken. Darum bin ich unter Anderem an der Hochschule Anhalt tätig und versuche, meine Studierenden von vornherein darauf einzuschwören.
>
> Ich könnte vermutlich einen halben Roman dazu verfassen, und auch auf den einen oder anderen Punkt aus eurem Episoden eingehen, aber wie zielführend das ist kann ich nicht sagen. Wenn euch so etwas interessiert, könnte man dazu aber ja gern mal quatschen, oder ihr fragt nochmal nach konkreten Kommentaren, dann setze ich mich hin und mache euch eine Liste fertig ;).
>
> Bis dahin habt ihr vermutlich einen neuen Hörer dazugewonnen, je nachdem, ob ich noch einen Podcast mehr in meinen täglichen Informationswust reinquätschen kann ;). Macht weiter so, haltet die Ohren steif, und viel Erfolg und Glück im Jahr 2026.
