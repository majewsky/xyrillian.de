- Rückblick: STP006 "Kommandozeilen und Unixphilosophie"
    - ttimeless hat gerade eine Gelegenheit, Daten zu analysieren, an der man das praktische Vorgehen mit Kommandozeilen-Werkzeugen gut demonstrieren kann
    - Ausgangssituation: [Klimasensor-Datenlogger](https://www.conrad.de/de/p/voltcraft-dl-210th-dl-210th-temperatur-datenlogger-luftfeuchte-datenlogger-messgroesse-temperatur-luftfeuchtigkeit-30-1435091.html) in verschiedenen Räumen messen einmal pro Minute Temperatur und Luftfeuchtigkeit
    - mitgelieferte Software bietet eine rudimentäre Visualisierung sowie einen Datenexport als [CSV](https://de.wikipedia.org/w/index.php?title=CSV_(Dateiformat)&oldid=244981555)
    - die [hier besprochenen CSV-Dateien](http://dl.xyrillian.de/noises/stp-060-shownotes/) können zum Nachvollziehen heruntergeladen werden
    - Ziel: eigene Datenvisualisierung mit [gnuplot](https://de.wikipedia.org/w/index.php?title=Gnuplot&oldid=226771326)

- erste Sichtprüfung: `<Balkon.csv head`
    - acht Kopfzeilen mit Metadaten und Auswertungen (Berichtszeitraum, Maximum/Minimum/Durchschnitt)
    - diesen Header brauchen wir nicht, wir wollen die Rohdaten
    - Option 1: sofern immer acht Zeilen, z.B. `tail -n +9` oder `sed 1,8d`
    - Option 2: Suchmuster auf den Zeileninhalt der letzten Kopfzeile, z.B. `sed 1,/^Aufnahmezeit/d`
    - Option 3: nur Zeilen auswählen, die nach Daten aussehen, z.B. `grep '^"[0-9]'`

- Kopfzeilen entfernt: `<Balkon.csv sed 1,8d`
    - nächstes Problem: die Anführungszeichen um die einzelnen Werte werden gnuplot verwirren -> würden wir mit `tr -d \"` wegbekommen
    - Problem: dann wären die Dezimalkommas nicht von den wertetrennenden Kommas zu unterscheiden
    - unsere Lösung: während wir Anführungszeichen entfernen, werden dazwischenliegende Kommas gleichzeitig durch ein anderes Trennzeichen (hier einen Tabulator) ersetzt -> `sed 's/^"//; s/","/\t/g; s/"\s*$//'`
    - hierin ist Trickserei erforderlich, weil die erzeugten Dateien Windows-Zeilenenden haben (CR+LF), nicht Unix-Zeilenenden (nur LF) -> könnte man auch vorher mit `dos2unix` beheben, dann ist statt `s/"\s*$//` nur `s/"$//` erforderlich
    - Alternative (erfordert [csvkit](https://csvkit.readthedocs.org)): `csvformat -d , -q \" -U 3 -D $'\t'`

- Anführungszeichen entfernt: `<Balkon.csv sed 1,8d | sed 's/^"//; s/","/\t/g; s/"\s*$//' >Balkon.dat`
    - zu Xyrills Überraschung ist das schon ausreichend, den Rest können wir in gnuplot machen

- Anzeige in gnuplot: das untenstehende Programm als `Balkon.gp` abspeichern, `gnuplot` aufrufen und `load 'Balkon.gp'` eingeben)
  ```
  set xdata time
  set timefmt '%d-%m-%Y %H:%M:%S'
  set format x '%d.%m.%Y %H:%M:%S'
  set decimalsign ','
  set datafile separator "\t"
  set ylabel "Temperatur °C"
  set y2label "Luftfeuchte %"
  set y2tics
  plot 'Balkon.dat' using 1:2 title 'Temperatur' with lines axes x1y1, 'Balkon.dat' using 1:3 title 'Luftfeuchtigkeit' with lines axes x1y2
  ```

- Für die Ausgabe in eine Datei können noch diese beiden Zeilen vorangestellt werden:
  ```
  set terminal svg background "white" dashed size 1920,1080
  set output "Balkon.svg"
  ```

- Was, wenn gnuplot nicht so clever wäre?
    - Dezimaltrennzeichen von Komma in Punkt ändern: `<Balkon.dat tr , .`
    - Datumsformat in ISO ändern: `<Balkon.dat awk -F'\t' '{ split($1, DATE, /[^0-9]*/); $1 = DATE[3] "-" DATE[2] "-" DATE[1] "T" DATE[4] ":" DATE[5] ":" DATE[6] "Z"; print }'`
    - Xyrill hatte auch geschaut, wie man die Zeitangaben in `date -d` bekommen könnte, um sie damit umzuwandeln, und wurde von [einem wunderschönen Zitat](https://www.gnu.org/software/coreutils/manual/html_node/Date-input-formats.html) begrüßt (nachfolgende deutsche Übersetzung mittels [DeepL](https://www.deepl.com/translator)):
      > Unsere zeitlichen Maßeinheiten, von Sekunden bis hin zu Monaten, sind so kompliziert, asymmetrisch und disjunkt, dass ein kohärentes mentales Rechnen in der Zeit nahezu unmöglich ist. In der Tat, hätte sich ein tyrannischer Gott ausgedacht, unseren Geist an die Zeit zu versklaven, um es uns nahezu unmöglich zu machen, der Unterwerfung unter durchweichte Routinen und unangenehme Überraschungen zu entkommen, hätte er kaum etwas Besseres tun können, als uns unser gegenwärtiges System zu überliefern. Es ist wie ein Satz trapezförmiger Bauklötze, ohne vertikale oder horizontale Flächen, wie eine Sprache, in der der einfachste Gedanke verschnörkelte Konstruktionen, nutzlose Partikel und langatmige Umschreibungen erfordert. Im Gegensatz zu den erfolgreicheren Mustern der Sprache und der Wissenschaft, die es uns ermöglichen, Eindrücken kühn oder zumindest nüchtern zu begegnen, fördert unser System der zeitlichen Berechnung stillschweigend und beharrlich unseren Terror vor der Zeit.
      >
      > ... Es ist, als ob Architekten die Länge in Fuß, die Breite in Metern und die Höhe in Ellen messen müssten; als ob grundlegende Bedienungsanleitungen die Kenntnis von fünf verschiedenen Sprachen erfordern würden. Kein Wunder also, dass wir in unsere eigene unmittelbare Vergangenheit oder Zukunft, auf den letzten Dienstag oder den Sonntag in einer Woche, oft mit einem Gefühl hilfloser Verwirrung blicken.  ...

- im Gespräch erwähnt: [ASCII-Folge von Request for Comments](https://requestforcomments.de/archives/400)
