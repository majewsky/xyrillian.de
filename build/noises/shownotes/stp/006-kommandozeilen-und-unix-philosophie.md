* Begriffsklärung
    * **Kommandozeile**: allgemein jede textbasierte Steueroberfläche
    * **Terminal** ("Endpunkt"): ursprünglich ein einfacher Computerarbeitsplatz, der an einen Großrechner angeschlossen ist und eine Kommandozeile zur Steuerung desselben bereitstellt; heutzutage synonym mit **Terminal-Emulator** (ein Programm, das denselben Zweck erfüllt)
    * **Shell** (auch "Befehlszeile" genannt): ein Programm, das die Grundstruktur der Befehlsinteraktion im Terminal bereitstellt und das Ausführen und Kombinieren konkreter Befehle ermöglicht
    * **Befehl**: ein Programm, das in der Shell ausgeführt wird und eine konkrete Aktion bewirkt

* Warum werden Terminals oft gegenüber grafischen Oberflächen (GUI) bevorzugt?
    * ~~Weil es viel cooler ist.~~
    * bessere Handhabbarkeit eines großen Funktionsumfangs, nach entsprechender Einarbeitung (vgl. GUI mit 1000 Knöpfen vs. Kommandozeile mit 1000 Befehlen)
    * Programmierbarkeit: Textmakros sind einfacher und robuster als GUI-Makros
    * Kombinierbarkeit: Textbausteine sind auf offensichtliche Art und Weise kombinierbar, grafische Bearbeitungsschritte nur begrenzt
    * Fernwartung: Übertragung von Text braucht weniger Bandbreite als Übertragung einer GUI
    * Zugriff bei Problemen oder bei Neuinstallation: wenn GUI nicht mehr oder noch nicht geht, geht oft noch das Terminal

* Nachteile von Terminals
    * schlechtere Auffindbarkeit von Funktionen ("Discoverability")
    * steile Lernkurve

* Befehlszeilen gibt es überall, aber besonders prägend ist die von Unix
    * **Unix**: Familie von Betriebssystemen
    * ursprünglich Anfang der 1970er bei [Bell Laboratories](https://de.wikipedia.org/wiki/Bell_Laboratories ) entwickelt von
        * [Dennis Ritchie](https://de.wikipedia.org/wiki/Dennis_Ritchie )
        * [Douglas McIlroy](https://de.wikipedia.org/wiki/Douglas_McIlroy)
        * [Ken Thompson](https://de.wikipedia.org/wiki/Ken_Thompson)
        * [Brian W. Kernighan](https://en.wikipedia.org/wiki/Brian_Kernighan ) wurde im Gespräch auch erwähnt. Er schrieb viele Programme für Unix und half, awk zu entwickeln.
    * wichtige aktuelle Vertreter: Linux (und damit Android), macOS, BSD-Familie
    * Gründe für den fortwährenden Einfluss des originalen Unix:
        * Quellcode war für universitäre Nutzung frei zugänglich
        * vergleichsweise einfacher und modularer Aufbau
        * dadurch einfach auf neue Systeme übertragbar

* [Unix-Philosophie](https://de.wikipedia.org/wiki/Unix-Philosophie )
    * Formulierung von Douglas McIlroy: "Mache nur eine Sache und mache sie gut."
    * Formulierung von Richard P. Gabriel: "Schlechter ist besser." (Einfachheit ist wichtiger als alle anderen Eigenschaften eines Programmes wie Fehlerfreiheit, Konsistenz oder Vollständigkeit, da die Weiterentwicklung vereinfacht wird.)
    * weitere Formulierungen siehe verlinkter Wikipedia-Artikel

* Unix-Philosophie im Kontext des Terminals
    * meist nicht ein Befehl, der die ganze Aufgabe erledigt; sondern viele kleine Befehle, die kombiniert werden können
    * gemeinsames Datenmodell: Ströme von Textdaten, die außerdem zumeist zeilenweise verarbeitet werden
    * Kombinieren der einzelnen Befehle erfolgt interaktiv (direkt in der Shell) und dadurch iterativ (schrittweise durch wiederholtes Probieren und Anpassen der Befehlszeile)

* Beispiel: Wir wollen aus einem gegebenen Text die häufigsten Wörter bestimmen.
    * Eingabedatei: [BGB als Markdown-Datei](https://github.com/bundestag/gesetze/raw/7c8a6a955bdefcf32632ff85ca671225e94513f9/b/bgb/index.md ) aus der inoffiziellen Quelle <https://github.com/bundestag/gesetze>
    * Schritt 1: Einlesen der Datei
        * `cat index.md`
    * Schritt 2: jedes Wort auf eine eigene Zeile
        * `cat index.md | tr ' ' '\n'`
        * passt nicht ganz, aber ist für unsere Zwecke gut genug ("Schlechter ist besser")
    * Schritt 3: Wörter sortieren, dann identische Wörter zusammengruppieren und zählen
        * `cat index.md | tr ' ' '\n' | sort | uniq -c`
    * Schritt 4: Wörter nach ihrer Häufigkeit sortieren
        * `cat index.md | tr ' ' '\n' | sort | uniq -c | sort -k1 -n`
    * Schritt 5: nur die letzten 20 Zeilen, die die häufigsten Wörter enthalten
        * `cat index.md | tr ' ' '\n' | sort | uniq -c | sort -k1 -n | tail -n 20`
    * siehe unten: Auszüge aus den Ausgaben jedes einzelnen Schrittes

* Abendgedanken: Eine Lernkurve ist nicht unbedingt ein Nachteil. Es ist durchaus wertvoll, wenn ein Programm seinem Nutzer richtig kommuniziert, wie viel Einarbeitung und Wissen erforderlich ist. Siehe LaTeX vs. Word.

* Im Gespräch erwähnt wurden auch:
     * [Noam Chomsky](https://de.wikipedia.org/wiki/Noam_Chomsky )
     * [Alan Turing](https://de.wikipedia.org/wiki/Alan_Turing )
     * [Steve Jobs](https://de.wikipedia.org/wiki/Steve_Jobs )

* Errata: Der Begriff "Mem" wurde nicht von [Douglas Hofstadter](https://de.wikipedia.org/wiki/Douglas_R._Hofstadter) eingeführt, sondern von [Richard Dawkins](https://de.wikipedia.org/wiki/Richard_Dawkins) (in seinem Buch ["Das egoistische Gen"](https://de.wikipedia.org/wiki/Das_egoistische_Gen)).

---

## Auszüge aus den einzelnen Schritten des Beispiels

Jeweils mit `| head`, um nur die ersten 10 Zeilen zu zeigen.

### Schritt 1

```
$ cat index.md | head
---
Title: Bürgerliches Gesetzbuch
jurabk: BGB
layout: default
origslug: bgb
slug: bgb

---

# Bürgerliches Gesetzbuch (BGB)
```

### Schritt 2

```
$ cat index.md | tr ' ' '\n' | head
---
Title:
Bürgerliches
Gesetzbuch
jurabk:
BGB
layout:
default
origslug:
bgb
```

### Schritt 3

Hier sieht man keine Wörter im engeren Sinne, sondern Folgen von Sonderzeichen, deren Zeichen aus technischen Gründen vor den eigentlichen Buchstaben einsortiert werden. In der ganz ersten Zeile taucht die Anzahl von Zeilen auf, die nach Schritt 2 komplett leer waren (zum Beispiel Leerzeilen, die im Ursprungsdokument Absätze voneinander trennen).

```
$ cat index.md | tr ' ' '\n' | sort | uniq -c | head
  18490
      1 #
      5 ##
     35 ###
    376 ####
   1216 #####
   1015 ######
     24 (§
      6 (§§
      1 *
```

### Schritt 4

Aufgrund der Sortierreihenfolge sehen wir hier besonders seltene Ergebnisse mit je einem Treffer.

```
$ cat index.md | tr ' ' '\n' | sort | uniq -c | sort -k1 -n | head
      1 #
      1 *
      1 ,
      1 0,75
      1 1).
      1 1\.
      1 1005
      1 1006
      1 1007
      1 1008
```

### Schritt 5

Das Endergebnis.

```
$ cat index.md | tr ' ' '\n' | sort | uniq -c | sort -k1 -n | tail -n 20
   1374 auf
   1383 nach
   1391 für
   1448 ist
   1523 von
   1568 so
   1574 das
   1630 wenn
   1692 in
   1850 zu
   1953 nicht
   2127 und
   2628 den
   2737 dem
   2909 oder
   3584 §
   4718 des
   5980 die
   8602 der
  18490
```
