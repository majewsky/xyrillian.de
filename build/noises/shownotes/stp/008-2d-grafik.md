* computergenerierte Bilder sind **rasterisiert** (in Bildpunkte unterteilt)
    * jeder Bildpunkt hat genau eine Farbe
    * generell sind verschiedene Rasterisierungen möglich (z.B. bei Facettenaugen von Insekten meist hexagonale Rasterung)
    * bei Computern meist quadratische Pixelraster: vereinfacht Speicherabbildung und bestimmte Berechnungen, die Nachbarpixel einbeziehen
* Farbe ist nicht eine einzige Zahl, sondern ein Punkt in einem mehrdimensionalen **Farbraum**
    * verschiedene Farbräume/Farbmodelle je nach Anwendung
        * [RGB](https://de.wikipedia.org/wiki/RGB-Farbraum ) (Rot-Grün-Blau): für selbstleuchtende Monitore mit drei verschiedenen Grundfarben ([Beispiel aus einem CRT-Monitor](https://commons.wikimedia.org/wiki/File:Lochmaske_makro.jpg ))
        * [YUV](https://de.wikipedia.org/wiki/YUV-Farbmodell ) (Y-Luminanz, UV-Chrominanz) etc.: für Fernsehbildschirme (abwärtskompatibel zu Schwarzweißfernsehern, die UV ignorieren und nur Y anzeigen)
        * [HSV](https://de.wikipedia.org/wiki/HSV-Farbraum ) und [CIELAB](https://de.wikipedia.org/wiki/Lab-Farbraum ): modellieren die natürliche Farbwahrnehmung des Menschen
        * [CMYK](https://de.wikipedia.org/wiki/CMYK-Farbmodell ) (Cyan-Magenta-Yellow-Black): für Druckprozesse mit vier Standardtinten
    * Farbräume sind meistens dreidimensional, weil wir drei Arten von Farbrezeptoren im Auge haben (zum Vergleich: Fangschreckenkrebse haben bis zu 12 Rezeptorarten und bräuchten also einen zwölfdimensionalen Farbraum)
* Beispielrechnung: Wie groß ist der Datenstrom zu einem normalen Computermonitor?
    * angenommen 4K-Auflösung: 3840x2160 Pixel
    * angenommen 8 Bit Farbtiefe: 8 Bit pro Pixel und Farbe, also 3 Byte pro Pixel (es gibt auch schon Monitore mit 10 Bit oder 12 Bit Farbtiefe)
    * angenommen 60 Hz, also 60 Bilder pro Sekunde
    * Ergebnis (ohne Datenkompression): 3840 \* 2160 Pixel / Bild \* 3 Byte / Pixel \* 60 Bilder / Sekunde = 1.39 GiB/s bzw. 11.1 Gib/s
    * zum Vergleich: komprimiertes 4K-Video typischerweise 2 MiB/s (16 Mib/s), Kompressionsfaktor ~700

* Bildaufbau in traditionellen Systemen mit 2D-Grafik (z.B. Spielkonsolen der 80er und frühen 90er) erfolgt meist mittels **Sprites**
    * Sprites: einzelne vorgefertigte Bilder (z.B. Glyphen einer Terminal-Schriftart oder einzelne Objekte in einem Videospiel)
        * [Pixel-Art](https://en.wikipedia.org/wiki/Pixel_art )
    * Grafikprozessor erhält eine Liste von Sprites mit den entsprechenden Positionen, an denen diese eingefügt werden sollen
    * Bild entsteht als Überlagerung aller Sprites in ein Gesamtbild
    * Gesamtbild liegt somit im Videospeicher vor (einem separaten Speicher des Grafikprozessors) und kann dann zum Monitor geschickt werden (siehe auch [Double Buffering](https://de.wikipedia.org/wiki/Doppelpufferung ) und [V-Sync](https://de.wikipedia.org/wiki/Vertikale_Synchronisation ))

* in Abgrenzung dazu: **Vektorgrafik**
    * Beschreibung eines Bildes mittels grafischer Primitiven (Rechtecke, Kreise, Linien, Kurven)
    * Bildaufbau durch **Rasterisierung** dieser Primitiven, z.B. mittels des [Bresenham-Algorithmus](https://de.wikipedia.org/wiki/Bresenham-Algorithmus )
    * bis Ende der 90er oft mittels separater Koprozessoren (z.B. [Super FX](https://de.wikipedia.org/wiki/Super_FX )), da die normalen Prozessoren nicht schnell genug waren

* Bildaufbau in traditionellen Desktopsystemen (z.B. Windows 95 bis XP) ist eine Mischung aus beiden
    * Fensterrahmen, Knöpfe etc. als Vektorgrafik (teilweise in Sprites vorgerendert)
    * Textsatz ebenfalls auf Basis vorgerenderter Sprites für die einzelnen Zeichen, die wiederum als Vektorgrafik definiert sind
* seit den 2000er-Jahren werden auch 2D-Bilder zunehmend auf der Grafikkarte berechnet, denn:
    * Grafikkarten bzw. Grafikprozessoren (GPUs) sind seit diesem Punkt hinreichend allgegenwärtig
    * GPUs sind für diese Aufgabe deutlich effizienter (sowohl in Geschwindigkeit als auch Energieverbrauch)
    * GPUs können auch fortgeschrittene Effekte wie Schattierungen oder Animationen effizient berechnen
