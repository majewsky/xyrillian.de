- Rückbezug zu STP008 (2D-Grafik) und STP010 (3D-Grafik): dort besprochen die grundsätzliche Methodik zum Zeichnen mit Pixeln, aber wer zeichnet hier eigentlich?
    - [Infografik](https://commons.wikimedia.org/w/index.php?title=File:Schema_der_Schichten_der_grafischen_Benutzeroberfl%C3%A4che.svg&oldid=796566713) zum Rendering-Ablauf unter Linux, mit Display-Manager (bzw. heutzutage Wayland-Compositor) und Desktop-Shell
    - unter Windows war früher alles in explorer.exe, heute gibt es den [Desktop Window Manager](https://learn.microsoft.com/en-us/windows/win32/learnwin32/the-desktop-window-manager) dafür

- Thema heute: Wie baut man als Anwendungsprogramm seine GUI zusammen?
    - klassischer Ansatz: [GUI-Toolkit](https://de.wikipedia.org/w/index.php?title=GUI-Toolkit&oldid=246182813); Programmbibliothek mit vorgefertigten Steuerelementen (Texteingabefeld, Knopf, Checkbox, Listenansicht), siehe Beispielbilder im verlinkten Artikel
    - neumodischer Ansatz: Darstellung einer Programmoberfläche mit einer Webbrowser-Engine (z.B. [Electron](https://de.wikipedia.org/w/index.php?title=Electron_(Framework)&oldid=247717134)); somit Verwendung von Webtechnologien (HTML, CSS und JavaScript) an Stelle (im Wortsinne) eines GUI-Toolkits
        - bis dahin, dass selbst die Interface-Elemente von Computerspielen mittlerweile quasi mit einem Webbrowser gerendert werden ([Beispielprodukt](https://coherent-labs.com/products/coherent-gameface/))
        - eine Variation von STP022: Webentwickler gibt es wie Sand am Meer

- Wie baut man so ein GUI-Toolkit? [Retained Mode vs. Immediate Mode](https://de.wikipedia.org/w/index.php?title=Retained_Mode&oldid=245388723)
    - Retained Mode: die Grafikbibliothek behandelt die Steuerelemente als langlebige Objekte, die vom Programm über die Zeit manipuliert werden
    - Immediate Mode: beim Zeichnen jedes Einzelbildes muss das Programm die gewünschten Steuerelemente neu deklarieren; Grafikbibliothek speichert nur wenig Zustand (z.B. aktuelle Cursorposition oder noch nicht verarbeitete Eingabeereignisse)
    - Unterschiede:
        - Retained Mode ist intuitiver (aus Entwicklersicht)
        - Retained Mode führt oft zu einer Duplikation von Zustandsverwaltung (z.B. speichert das Checkbox-Steuerelement ein Bit, ob es angehakt ist; und der Applikationscode selber speichert ein gleichlautendes Bit, ob die entsprechende Funktion aktiviert ist)
        - Immediate Mode ist flexibler (z.B. gibt eine direkte Möglichkeit, die Standard-Steuerelemente noch mit einzelnen Zeichenbefehlen zu erweitern)
        - Immediate Mode kann resourcenschonender sein (z.B. geringerer Speicherverbrauch, weil nicht sichtbare Steuerelemente weniger wahrscheinlich im Speicher gehalten werden)
        - Immediate Mode kann auch mehr Ressourcen verbrauchen (z.B. weil jedes Einzelbild neu gerendert werden muss, auch wenn sich meist in weiten Teilen gar nichts ändert)
    - Mischform: [React](https://de.wikipedia.org/w/index.php?title=React&oldid=248452835)-artige GUI-Toolkits arbeiten im Retained Mode, aber alle Aktualisierungen führen zu einer kompletten Neudeklaration der betroffenen Steuerelemente; diese Deklaration wird dann mit den vorhandenen Steuerelementen abgeglichen

- Layoutstrategien: Wie werden Steuerelemente gegeneinander angeordnet?
    - Ziel: ein Arrangement von Steuerelementen soll z.B. auf Änderungen der Fenstergröße sinnfällig reagieren
    - [Beispielbild](https://doc.qt.io/qt-6/images/qgridlayout.png) aus der [Dokumentation zu QGridLayout](https://doc.qt.io/qt-6/qgridlayout.html) -> bei Änderungen der Höhe des Fensters sollen die beiden oberen Zeilen gleichbleiben und nur die unterste Zeile wachsen
    - explizite Deklaration durch den Programmentwickler mühselig: stattdessen berechnen Steuerelemente selbst bestimmte Eigenschaften wie "minimale Größe", "maximale Größe", "sinnvolle Größe" oder "wieviel profitiere ich von Wachstum"

- Layoutstrategien bei Webseiten und Webapplikationen
    - anhand des Textflusses, z.B. Eingabefelder und Links/Schaltflächen innerhalb von Mengentext (in anderen GUI-Toolkits unüblich; das Web ist durch seine Dokumentenorientierung hier besonders)
    - für alles andere bis ca. 2013: tabellenbasierte Layouts mit manuell (z.B. per JavaScript) berechneten Größen; in der Praxis frustrierend und fehleranfällig
    - seit 2013: [Flexbox](https://css-tricks.com/snippets/css/a-guide-to-flexbox/) für Layout von Elementen mit dynamischer Größe nebeneinander bzw. in Form eines Textflusses
    - seit 2017: [Grid](https://css-tricks.com/snippets/css/complete-guide-grid/) für Layout von Elementen, die sich in ein dynamisches Raster einfügen
    - außerdem seit 2009/2010: [CSS Media Queries](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_media_queries/Using_media_queries) zum Umschalten zwischen Layouts auf Basis von Eigenschaften des Endgerätes (Bildschirmgröße und -ausrichtung, Maus- oder Touch-Bedienung, Dark Mode oder Light Mode, etc.)

- zum Weiterhören: [Pentaradio vom März 2024: "UI: Hui oder pfui?"](https://hedgedoc.c3d2.de/pentaradio-2024-03#)
