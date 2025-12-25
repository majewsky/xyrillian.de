- siehe STP081: gemäß des neuen [Barrierefreiheitsstärkungsgesetzes](https://de.wikipedia.org/w/index.php?title=Barrierefreiheitsst%C3%A4rkungsgesetz&oldid=258239575) müssen Webseiten und Webapplikationen die AA-Kriterien der WCAG 2.1 erfüllen
    - WCAG = [Web Content Accessibility Guidelines](https://de.wikipedia.org/w/index.php?title=Web_Content_Accessibility_Guidelines&oldid=254032297) (Webinhalte-Barrierefreiheits-Richtlinien) der [Web Accessibility Initiative](https://www.w3.org/WAI/) des W3C
    - Kriterien sind gruppiert in 4 Grundprinzipien und darin 13 Richtlinien
    - Kriterien gehen jeweils von A (Pflicht) über AA (Kür) bis AAA (Sahnehäubchen)
    - in STP081: das Grundprinzip 1 (Wahrnehmbarkeit) mit insgesamt vier Richtlinien, heute der Rest
    - zum Mitlesen: [WCAG 2.2 (aktuelle Version, nur Englisch)](https://www.w3.org/TR/WCAG22/) oder [WCAG 2.0 (etwas ältere Version, in offizieller deutscher Übersetzung)](https://www.w3.org/Translations/WCAG20-de/)

- Grundprinzip 2: **Bedienbarkeit**
    - Richtlinie 5: *Zugänglichkeit per Tastatur*
        - A: alle Aktionen per Tastatur ansteuerbar, die nicht eine genaue geometrische Ansteuerung erfordern (z.B. Zeichenfläche in einem Zeichenprogramm); kein Steuerelement kann den Tastaturfokus so fangen, dass man den Fokus nur mit einer anderweitigen Aktion (z.B. Mausklick) brechen kann; etc.
        - AAA: ausnahmslos alle Aktionen per Tastatur ansteuerbar
    - Richtlinie 6: *Ausreichend Zeit*
        - A: wenn Interaktionen ein Zeitlimit haben, muss das Limit (wann immer möglich) in irgendeiner Form umgehbar oder zurücksetzbar sein (Aufgaben sollen abschließbar sein, ohne dass ein Zeitlimit plötzlich den Fortschritt zunichte macht); jegliches automatische Bewegen/Scrollen/Blinken/Aktualisieren muss abschaltbar sein
        - AAA: absolut keine Zeitlimits für Interaktionen; bei Relogin aufgrund abgelaufener Sitzung darf kein Fortschritt verloren gehen; Benutzer müssen jederzeit über laufende Zeitlimits informiert sein, die kürzer als 20 Stunden sind
    - Richtlinie 7: *Keine Anfälle*
        - A: Inhalte, die mehr als 3x pro Sekunde blinken, müssen eine Intensität unter den Grenzwerten für Epilepsierisiko liegen
        - AAA: gar kein Blinken schneller als 3x pro Sekunde; Animationen infolge von Benutzerinteraktionen müssen abschaltbar sein
    - Richtlinie 8: *Navigierbarkeit*
        - A: sinnvolle Titel auf jeder Seite; über mehrere Seiten wiederholte Informationen sind überspringbar; Steuerelemente haben eine sinnvolle Fokusreihenfolge; der Zweck von Links soll soweit wie möglich aus dem Linktext allein erkennbar sein (z.B. nicht "\[Hier klicken], um den Kauf abzuschließen")
        - AA: statische Seiten sollen auf mehrere Wege auffindbar sein; Seiten sind mittels Teilüberschriften und Labels sinnvoll strukturiert; Tastaturfokus auf Steuerelementen soll klar erkennbar sein; Steuerelemente mit Fokus sollen nie komplett verdeckt sein (z.B. durch ein danebenliegendes Ausklappmenü); etc.
        - AAA: aktueller Ort des Nutzers innerhalb einer Webseitenstruktur soll klar erkennbar sein (z.B. "Sie sind hier: \[Start] > \[Blog] > \[Archiv] > Name des aktuellen Artikels"); etc.
    - Richtlinie 9: *Alternative Eingabemethoden*
        - A: Webseiten/-applikationen, die Mehrpunkt-Interaktion erlauben ("Touch-Gesten"), sollen auch eine Alternative für zeigerbasierte Einzelpunkt-Interaktion erlauben; Aktionen infolge fehlerhaften Runterdrückens der Maustaste sollen abwendbar oder umkehrbar sein; Interaktion durch Bewegungssteuerung soll auch durch Steuerelemente ersetzbar sein sowie gegen Fehlaktivierung resistent sein; etc.
        - AA: Klickflächen müssen mindestens 24x24 CSS-Pixel groß sein (vgl. [Fitts' Gesetz](https://de.wikipedia.org/w/index.php?title=Fitts%E2%80%99_Gesetz&oldid=258729192))
        - AAA: Klickflächen müssen mindestens 44x44 CSS-Pixel groß sein; alternative Eingabemethoden sollen beliebig gleichzeitig nutzbar sein (ohne explizites Umschalten)
        - im Gespräch erwähnt: [Zenons Paradoxon von Achilles und der Schildkröte](https://de.wikipedia.org/w/index.php?title=Achilles_und_die_Schildkr%C3%B6te&oldid=258338766), [The Topological Problem with Voting (Video)](https://www.youtube.com/watch?v=v5ev-RAg7Xs)

In Teil 3 dieser Serie folgen noch Grundprinzip 3 (Verständlichkeit) und 4 (Robustheit).
