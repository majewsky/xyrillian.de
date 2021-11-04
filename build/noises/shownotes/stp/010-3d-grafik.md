- Vergleich zu 2D-Desktop-Grafik (STP008)
    - Zeichenreihenfolge ist nicht so offensichtlich (Was ist vorne? Was ist hinten?)
    - höherer Anspruch an Realismus: dynamische Beleuchtung, Texturierung, etc.
    - Randbemerkung: auch 2D-Spiele werden heute meist in einer 3D-Szene aufgebaut
    - Grundeinheit für 3D-Modelle: nicht der Pixel, sondern das Polygon (deswegen kommen wir um Fließkommazahlen nicht herum)

- Vorbemerkungen zu den folgenden Schritten:
    - all dies passiert für jedes Einzelbild von vorne (also bei 60Hz 60-mal pro Sekunde)
    - zur Illustration verlinken wir auf Beispielbilder aus den [Übungsunterlagen für einen Kurs in 3D-Rendering-Grundlagen](https://github.com/ssloy/tinyrenderer/wiki/Lesson-0:-getting-started )

- Schritt 1: 3D-Modell mit Polygonen
    - Polygon heißt wörtlich "Vieleck", aber 3D-Modelle bestehen meist nur aus Dreiecken
    - Polygone mit vier und mehr Ecken können mittels [Triangulation](https://de.wikipedia.org/wiki/Triangulation_(Fl%C3%A4che) ) in Dreiecke zerlegt werden
    - [Darstellung eines 3D-Modells als Gitternetz](https://raw.githubusercontent.com/ssloy/tinyrenderer/gh-pages/img/01-bresenham/5da6818190.png )

- Schritt 2: Kamera
    - grundsätzlich ein Paar aus Blickpunkt und Blickrichtung, aber mit einigen weiteren Parametern
    - mathematische Grundlage: [Quaternionen](https://de.wikipedia.org/wiki/Quaternion )
    - [Sichtfeld](https://de.wikipedia.org/wiki/Sichtfeld ) (**Field of View**): beim Menschen 130°; bei Videospielen meist weniger, je nachdem, wieviel Platz der Bildschirm im Sichtfeld einnimmt (z.B. 90° am PC, 60° an Konsolen)

- Schritt 3: Ausmalen der Polygone (Rasterisierung)
    - **Backface Culling** (Rückseiten-Auslese): wir werfen alle Polygone weg, die von uns wegzeigen
    - [Beispiel mit ausgemalten Polygonen nach Backface Culling](https://raw.githubusercontent.com/ssloy/tinyrenderer/gh-pages/img/02-triangle/d5223f9b93.png )
    - **Z-Buffer**: wenn wir Polygone ausmalen, dann malen wir nur die Pixel aus, bei denen wir nicht schon Polygone gemalt haben, die näher an der Kamera dran sind
    - [Beispiel mit ausgemalten Polygonen unter Verwendung des Z-Buffers](https://raw.githubusercontent.com/ssloy/tinyrenderer/gh-pages/img/03-zbuffer/f93a1fc1cbaebb9c4670ae0003e62947.png )

- Schritt 4: Texturen
    - bis jetzt war jedes Polygon einfarbig, jetzt kleben wir Bilder auf (analog zu Modellbausätzen)
    - in der Praxis ein großes Bild für das ganze Modell
    - jede Ecke jedes Polygons wird zu einem Punkt in der Textur zugeordnet
    - [Beispiel mit Polygonen, die mit ihrer Textur ausgemalt sind](https://raw.githubusercontent.com/ssloy/tinyrenderer/gh-pages/img/03-zbuffer/73714966ad4a4377b8c4df60bef03777.png )

- Einwurf: Was unterscheidet einen Grafikprozessor (GPU) von einem Hauptprozessor (CPU)?
    - CPU-Kerne führen einzelne Befehle auf einzelnen Daten aus
    - GPU führt einzelne Befehle auf vielen Daten gleichzeitig aus, z.B. Rasterisierung von 10000 Polygonen gleichzeitig oder Z-Buffer-Prüfung für 10000 Pixel gleichzeitig
    - dadurch auch für andere massiv parallelisierbare Operationen interessant ("GPGPU", General Purpose GPU), z.B. wissenschaftliche Simulationen, maschinelles Lernen, Passwortknacken

- Schritt 5: Shader
    - bis jetzt haben wir die **Fixed Function Pipeline** beschrieben, eine im GPU-Treiber fest vorgegebene Kette von Rendering-Schritten
    - neuere GPU (seit den frühen 2000ern) erlauben das Einbetten von eigenen Programmen (**Shadern**) in die Pipeline
    - **Vertex Shader**: wird einmal pro Vertex (Ecke eines Polygons) aufgerufen, um die 2D-Position auf dem Bildschirm zu berechnen (ansonsten normale Berechnung anhand der konfigurierten Kameraposition)
    - **Fragment Shader**: wird einmal pro Pixel aufgerufen, um gegebenenfalls die Farbe des Pixels zu modifizieren
    - [Beispiel eines Fragment-Shaders, der eine Farbtextur auf 6 feste Farbstufen reduziert](https://raw.githubusercontent.com/ssloy/tinyrenderer/gh-pages/img/06-shaders/f2bf83c5994b9051aaba499cb05e65bf.png )

- Schritt 6: Beleuchtung
    - bei traditionallen 2D-Spielen mit Sprites war Beleuchtung meist global (z.B. jedes Objekt hat 2 Sprites, einmal für dunkle und einmal für helle Szenen)
    - dynamische Beleuchtung: heute meist als Teil des Fragment Shaders
    - Beispiel: [Phong's Approximation einer natürlichen Beleuchtung](https://raw.githubusercontent.com/ssloy/tinyrenderer/gh-pages/img/06-shaders/e3720a5dfedc49edb0bf70f8bc64204a.png)
        - 1. Anteil: ambiente Beleuchtung (überall gleich stark, Grundausleuchtung wie in einem Fernsehstudio)
        - 2. Anteil: diffuse Beleuchtung (abhängig davon, wie sehr die beleuchtete Fläche zur Lichtquelle hingewendet ist)
        - 3. Anteil: spekulare Beleuchtung (Glänzen; stark überbetont auf Flächen, die fast exakt zur Lichtquelle hingewendet sind)

- es gibt noch so viel mehr: Kantenglättung (Antialiasing), lineares und anisotropisches Filtern, Normal Mapping, Raytracing
