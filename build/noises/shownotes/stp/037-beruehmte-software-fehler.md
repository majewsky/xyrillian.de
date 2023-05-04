Errata zu STP034 von Thomas aus Berlin:

> Am Ende der Folge wundert sich ttimeless darüber, warum `sin(pi)` auf dem Taschenrechner immer ein Ergebnis von 0,0548... produziert. Du versuchst, das dann mit möglichen Ungenauigkeiten der Abbildung von pi zu erklären.
>
> Ich glaube, die Antwort ist viel einfacher: In der normalen Einstellung DEG erwartet der Taschenrechner die Eingabe in Grad und `sin(pi°) = 0,0548...` ist damit auch korrekt. Stellt man den Taschenrechner auf RAD (Radiant), dann rechnet der Taschenrechner auch korrekt `sin(pi) = 0`.

- "Off by One" ("um eins daneben")
    - Diese Folge ist Nr. 37 und somit die 38. Folge dieses Podcasts!
    - häufige Konvention in der IT: Nummerierung fängt bei Null an
        - In einer Liste mit 10 Elementen das vorderste den Index 0 und das hinterste den Index 9.
        - formal plausibel, da Null die erste natürliche Zahl ist (ansonsten ist ℕ unter Addition keine Gruppe; siehe [EWD831](https://www.cs.utexas.edu/users/EWD/transcriptions/EWD08xx/EWD831.html))
        - technisch plausibel, da die Speicheradresse einer Liste für gewöhnlich die Speicheradresse des vordersten Elementes ist; dieses hat also einen Abstand von 0 zum Beginn der Liste
    - siehe auch [Zaunpfahlfehler](https://de.wikipedia.org/w/index.php?title=Zaunpfahlfehler&oldid=228434264)
    - Thema heute: Wie katastrophal können die Auswirkungen solcher kleinen Fehler sein?

- mechanische Fehler: ["First actual case of bug being found."](https://commons.wikimedia.org/w/index.php?title=File:First_Computer_Bug,_1945.jpg&oldid=682909364)

- Überlauf: [Ariane-Flug V88 (1996)](https://de.wikipedia.org/w/index.php?title=Ariane_V88&oldid=222991385)
    - erster Flug der Ariane 5
    - Steuerungssoftware war von der Ariane 4 übernommen und nicht auf derart große Horizontalgeschwindigkeiten ausgelegt
    - Überlauf bei Konvertierung von float64 in int16; aber bei dieser konkreten Variable keine Überlaufprüfung wegen Ressourcenknappheit
    - größeres Motiv laut Untersuchungsbericht: in der Gefahrenabschätzung wurde nicht hinreichend berücksichtigt, dass die Software (und nicht nur die Hardware) fehlerhaft sein könnte
    - Folge: erster großflächiger Einsatz von Fehlersuche mittels statischer Code-Analyse

- Überlauf: Y2K Bug -> siehe STP016

- Überlauf: [Heartbleed (2014)](https://de.wikipedia.org/w/index.php?title=Heartbleed&oldid=229430945)

- Drift: [Dhahran Missile Incident (1991)](https://www-users.cse.umn.edu/~arnold/disasters/Patriot-dharan-skeel-siam.pdf), [siehe auch](https://www.gao.gov/assets/imtec-92-26.pdf)
    - Hintergrund: 28 Amerikaner getötet durch Raketenbeschuss auf eine US-Basis in Saudi Arabien während des 1. Golfkriegs
    - Fehler im Raketenabwehrsystem: aufgrund Rundungsfehler driftet die interne Uhr um einen Faktor von 0,00000095
    - nach 100 Stunden Dauerbetrieb ist die Drift bei 0,34 Sekunden; entspricht bei Raketen mit Mach 4 einer Distanz von etwa 500 Metern

- Spezifikation falsch gelesen (oder Kommunikationsproblem): [Mars Climate Orbiter (1999)](https://de.wikipedia.org/w/index.php?title=Mars_Climate_Orbiter&oldid=230063104)
    - NASA hatte metrische Einheiten verwendet (Impuls in Newtonsekunden), die Navigationssoftware von Lockheed Martin hingegen imperiale Einheiten (Impuls in Pound-Force-Sekunden)

- Clusterfuck: [Therac-25 (1985-1987)](https://de.wikipedia.org/w/index.php?title=Therac-25&oldid=222407363), unter anderem:
    - ~~verteiltes~~ verstreutes System: fehlerhafte Synchronisation zwischen einem Messprozess und einem parallel arbeitenden Steuerprozess
    - Benutzerführung: in bestimmten Situationen wurde nach der Eingabe neuer Daten der alte Datensatz weiterverwendet, obwohl die Benutzeroberfläche den neuen Datensatz anzeigte
    - Überlauf: 8-Bit-Flag für die Anzeige der Notwendigkeit einer Sicherheitsprüfung konnte von 255 auf 0 überlaufen und dadurch die Prüfung überspringen

- überfordert von Komplexität: [Corrupted Blood (2005)](https://en.wikipedia.org/w/index.php?title=Corrupted_Blood_incident&oldid=1140307355)
    - vor COVID-19 eines der wichtigsten Studienobjekte für das Verbreitungsmuster einer Pandemie, inklusive Tiervektoren und asymptomatisch infektiösen Patienten
    - ein bestimmter Bosskampf in "World of Warcraft" versieht Spielercharaktere mit einem negativen Statuseffekt, der sie langsam tötet und der auf naheliegende Charaktere ansteckend wirkt
    - Spieler sollten mit diesem Statuseffekt den Bosskampf nicht verlassen können, aber die  Reittiere der Spieler konnten es!

- analog zu Copy-Paste-Fehler: [Mariner 1 (1962)](https://en.wikipedia.org/w/index.php?title=Mariner_1&oldid=1136631227)
    - Programmcode wurde auf Papier geschrieben, dort wurde fehlerhafterweise mit "R" statt mit dem korrekten "R̅" (R mit Überstrich) gerechnet

- Copy-Paste-Fehler: [Pentium-FDIV-Bug](https://de.wikipedia.org/w/index.php?title=Pentium-FDIV-Bug&oldid=226232131)
    - Division von Fließkommazahlen (siehe STP009) beim Pentium verwendet einen Algorithmus, der zur Beschleunigung eine Tabelle von vorberechneten Werten verwendet
    - fünf Einträge in dieser Tabelle wurden nicht richtig in das finale Chip-Design übernommen
    - Ergebnis: bestimmte Divisionen liegen um einen Bruchteil eines ‰ daneben
    - Rückrufaktion im Wert von über 400 Millionen $
    - "Wie viele Intel-Mitarbeiter braucht man, um eine Glühbirne zu wechseln? 1,9999983256"
    - im Gespräch erwähnt: [Fast Inverse Square Root](https://en.wikipedia.org/w/index.php?title=Fast_inverse_square_root&oldid=1142010709)
