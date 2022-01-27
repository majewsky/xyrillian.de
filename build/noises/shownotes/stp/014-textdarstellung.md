- Nachbemerkungen zum letzten Mal: bidirektionaler Text
    - Mischung aus LTR-Schrift (_left to right_) und RTL-Schrift (_right to left_), z.B. hebräische oder arabische Zitate in lateinischer Schrift
    - in der Zeichenkette entsprechend der logischen Reihenfolge
    - in der Darstellung: Zeichen können (aber müssen nicht) eine präferierte Schreibrichtung haben, explizite Umschaltung der Schreibrichtung mit Steuerzeichen möglich

- Grundbegriffe
    - "Zeichen": wird im Bereich von Unicode kaum verwendet, da nicht klar definiert
    - [Graphem](https://de.wikipedia.org/wiki/Graphem ): die kleinste bedeutungsunterscheidende (aber nicht bedeutungstragende) Einheit eines Schriftsystems (z.B. lateinische Schrift: Buchstaben, Ziffern, Sonderzeichen, Leerzeichen); in Unicode ungefähr gleichsetzbar mit einem Codepunkt (abgesehen von Steuerzeichen)
    - [Glyphe](https://de.wikipedia.org/wiki/Glyphe ): die grafische Darstellung eines Graphems
    - [Schriftart/Font](https://de.wikipedia.org/wiki/Schriftart ): ein vollständiger Satz von Glyphen

- Unterschiede zwischen Schriftarten
    - [Serifenschriften](https://de.wikipedia.org/wiki/Serife ) vs. serifenlose, insb. [Groteskschriften](https://de.wikipedia.org/wiki/Grotesk_(Schrift) )
    - [Proportionalschriften](https://de.wikipedia.org/wiki/Proportionale_Schriftart ) vs. [Monospace-Schriften](https://de.wikipedia.org/wiki/Nichtproportionale_Schriftart )
    - Spezialanwendungen: [OCR-A](https://de.wikipedia.org/wiki/OCR-A ), [FE-Schrift](https://en.wikipedia.org/wiki/FE-Schrift ), Schriften zur Unterstützung von Analphabeten (unter den bekannten Schriften insbesondere [Comic Sans](https://de.wikipedia.org/wiki/Comic_Sans_MS ))
    - Beispiel aus Asien: [Ming](https://de.wikipedia.org/wiki/Ming_(Schrift) ) vs. [Gothic](https://en.wikipedia.org/wiki/East_Asian_Gothic_typeface ) vs. [Kalligrafie](https://de.wikipedia.org/wiki/Chinesische_Kalligrafie )

- Phase 1: Textlayout
    - Segmentierung in Wörter, Zeilen, Absätze, evtl. Seiten
    - abhängig von den Regeln der verwendeten Schriftsysteme, siehe z.B. [UAX 14: Unicode Line Breaking Algorithm](https://www.unicode.org/reports/tr14/ )
    - abhängig von der Schreibrichtung, interessant insbesondere bei bidirektionalem Text
        - Unicode beschreibt fast ausschließlich zeilenweises Layout mit LTR oder RTL
        - im Japanischen ist auch spaltenweises Layout mit TTB (_top to bottom_) gängig, siehe [UTR 50: Unicode Vertical Text Layout](https://www.unicode.org/reports/tr50/ )
    - [Ruby](https://de.wikipedia.org/wiki/Ruby_Annotation ): Ausspracheannotationen an CJK-Schriftzeichen, außerhalb der eigentlichen Textzeile (in LTR darüber, in TTB rechts daneben)

- Phase 2: Schriftwahl
    - kaum ein Font deckt alle Grapheme ab -> im Allgemeinen Auswahl verschiedener Fonts je nach Schriftsystem notwendig (**Font Stack**)
    - in europäischen Texten früher kaum relevant, heute wichtig wegen separater Emoji-Fonts
    - Schriftgröße früher wichtig, da manche Fonts nur als Bitmaps in festen Größen vorlagen; heute durch Vektorschriftarten kein Problem mehr
    - Schriftwahl beeinflusst Phase 1: verschiedene Schriften laufen unterschiedlich weit und ergeben somit andere Zeilenumbruchspunkte

- Phase 3: Glyphenwahl
    - einfacher Fall: eine Glyphe für ein Graphem
    - allgemeiner Fall: Bildung von **Graphem-Clustern**
        - z.B. Ligaturen: "f" + "l" = "ﬂ"
        - z.B. kombinierende diakritische Zeichen: "a" + "◌́" = "á"
        - z.B. Emoji-Hautfarben: "👨" + [Fitzpatrick](https://de.wikipedia.org/wiki/Hauttyp#Hauttypen_nach_Fitzpatrick ) Modifier 1-2 = "👨🏻"
        - z.B. Emoji-Flaggen: "🇩" + "🇪" = "🇩🇪"
        - z.B. Fonts mit kreativer Verwendung von Ligaturen: [Sans Bullshit Sans](https://www.sansbullshitsans.com/ ), [Fira Code](https://github.com/tonsky/FiraCode )
    - anspruchsvollster Fall: [Complex text layout](https://en.wikipedia.org/wiki/Complex_text_layout ), sprich: Anpassung von Glyphen an ihre Nachbarglyphen
        - v.a. im Arabischen und in indischen Schriften, aber auch z.B. in handschriftartigen Fonts
    - auch Glyphenwahl beeinflusst Phase 1, aus demselben Grund wie oben

- Phase 4: Darstellung
    - Glyphen aus [Vektorschriftarten](https://de.wikipedia.org/wiki/Vektorfont ) werden in den gewünschten Schriftgrößen gerendert und vorgehalten
    - dann "einfach" Zusammensetzen der Glyphen zu einem Gesamtbild unter Anwendung der gewünschten Textfarbe
    - Ausnahme: Emoji-Glyphen folgen meist nicht der Schriftfarbe

- Quelle: [Spezifikationen und Reports des Unicode Consortium](https://www.unicode.org/faq/specifications.html )

- Hörempfehlung: [CRE080 Geschichte der Typografie](https://cre.fm/cre080-geschichte-der-typographie )
