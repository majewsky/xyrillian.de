* Einführungsbeispiel: simulierte schlechte Leitung ("xxxxxx heute xxxxxx Podcast xxxxxx Thema xxxxxx Kompression xxxxxx" -> "In unserer heutigen Folge unseres Podcasts ist das Thema die Datenkompression.")
    * naiver Impuls eines Erstsemester-Linguistikstudenten: "Warum sagt man nicht gleich 'heute Podcast Thema Kompression'?"
    * Sprache ist so gestaltet, dass auch teilweise fehlerhafte Informationen noch korrekt ankommen
    * beim Hören intuitiver Einsatz eines Vorhersagemodells, um kommende Silben/Wörter/Phrasen vorherzusagen

* theoretische Grundlagen
    * **Bit:** Basiseinheit für [Informationsgehalt](https://de.wikipedia.org/w/index.php?title=Informationsgehalt&oldid=219708601 ) (enstprechend der Auswahl aus zwei gleich wahrscheinlichen Möglichkeiten), allerdings nicht im SI-System verankert
    * Beispiel Münzwurf:
        * ideale Münze (50% Kopf + 50% Zahl) liefert 1 Bit pro Wurf
        * reelle Münze liefert mehr als 1 Bit pro Wurf (z.B. 49,5% Kopf + 49,5% Zahl + 1% Kante -> 1,07 Bit pro Wurf)
        * gezinkte Münze liefert weniger als 1 Bit pro Wurf (z.B. 10% Kopf + 90% Zahl -> 0,469 Bit pro Wurf)
    * gute Komprimierbarkeit = niedriger Informationsgehalt -> häufiger, als man denkt (Bsp. nebeneinanderliegende Pixel eines Bildes sind ähnlich)
    * [Zusammenhang mit physikalischer Entropie:](https://de.wikipedia.org/w/index.php?title=Entropie&oldid=223480319#Geschichtlicher_%C3%9Cberblick) Systeme mit geringer Entropie (z.B. Eisblock) sind einfach vorherzusagen, Beobachtungen haben geringen Informationsgehalt; Systeme mit hoher Entropie (z.B. Wasserdampf) sind schwer vorherzusagen, Beobachtungen haben hohen Informationsgehalt

* "naive" Kompressionsmethoden (die man auch als Mensch überblicken kann)
    * Nutzung von Allgemeinwissen
        * Beispiel aus der Vorbereitung: "Schachbrett" im Kontext von Exponentialfunktionen ruft die Assoziation einer ganzen Geschichte auf
    * Wörterbuchmethode für Text
        * [Beispiel aus Wikipedia](https://de.wikipedia.org/w/index.php?title=Datenkompression&oldid=222957700#Text ): "wenn Fliegen hinter Fliegen fliegen, fliegen Fliegen Fliegen nach" -> "wenn Fliegen hinter \\2 fliegen, \\5 \\2 \\2 nach"
    * RLE (Run-Length Encoding) z.B. in Bildern wie [dem Coverart dieses Podcasts](https://xyrillian.de/res/cover-stp.png )
    * [Huffman-Kodierung](https://de.wikipedia.org/w/index.php?title=Datenkompression&oldid=222957700#Entropiekodierung ) wie im [Morse-Code](https://de.wikipedia.org/w/index.php?title=Morsecode&oldid=222300739#Standard-Codetabelle ) (Verbindung zur Linguistik!)
    * Minifizierung von Skripten
        * Beispiel: jQuery 3.6 [ohne Minifizierung](https://code.jquery.com/jquery-3.6.0.js) und [mit Minifizierung](https://code.jquery.com/jquery-3.6.0.min.js )

* zwei fundamentale Abwägungen
    * Geschwindigkeit vs. Stärke (z.B. Debian-Pakete mit zwei getrennten Bereichen, die verschiedene Kompressionsverfahren nutzen)
    * verlustfrei vs. verlustbehaftet (Bsp. Pentaradio vom Mai 2022: Radiofassung als FLAC 271,7 MiB; Endprodukt als Opus 52,0 MiB)

* Kompressionsmethoden im Audio-Video-Bereich beispielhaft anhand ihres Verlustverhaltens
    * [Blockartefakte](https://de.wikipedia.org/w/index.php?title=Blockartefakt&oldid=121557554 ) bei JPEG
    * Psychoakustik bei [MP3](https://de.wikipedia.org/w/index.php?title=MP3&oldid=223021491#Verfahren )
    * Datamoshing bei Videos ([Beispiel](https://www.youtube.com/watch?v=rMSsw4CZvKg )) illustriert Bedeutung von [Intra-Frames](https://de.wikipedia.org/w/index.php?title=Intra-Frame&oldid=198265016 )
        * siehe auch [dasselbe Video, 1000-mal komprimiert und dekomprimiert](https://www.youtube.com/watch?v=icruGcSsPp0 )
    * Einfluss fester Bitrate auf Audio/Videos

* theoretische Grenze: [Kolmogorow-Komplexität](https://de.wikipedia.org/w/index.php?title=Kolmogorow-Komplexit%C3%A4t&oldid=207581838 )

* Fußnoten:
    * ttimeless fühlt sich verpflichtet, auf den [RFC-Podcast zu verlinken](https://requestforcomments.de/archives/320 )
    * David Kriesel: ["Traue keinem Scan, den du nicht selbst gefälscht hast"](https://media.ccc.de/v/31c3_-_6558_-_de_-_saal_g_-_201412282300_-_traue_keinem_scan_den_du_nicht_selbst_gefalscht_hast_-_david_kriesel )
