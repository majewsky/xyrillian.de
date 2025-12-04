Fortsetzung aus STP068 und STP074. Wir lesen weiterhin die englische Erstausgabe von 1975, die [im Internet Archive](https://archive.org/details/MythicalManMonth) digitalisiert wurde.
Erneut als Vorbemerkung: Aus Quellentreue sagen wir "Mann-Monat" statt "Person-Monat".

### Kurzzusammenfassung der bisherigen Folgen

Für die Zuschauer im Saal, die die vergangenen Folgen nicht gehört haben. Das übergreifende Motiv in dieser ersten Hälfte des Buches ist, wie sich die Arbeit in großen Teams von kleinen Teams unterscheidet.

1. Die Teergrube
   > So wie man in der Teergrube erst gefangen ist, wenn alle Beine drinstecken, so resultiert die Trägheit von Software-Entwicklung aus einem Zusammenspiel mehrerer Faktoren. Ein Programm zu schreiben, ist einfach. Aber die Produktisierung erfordert viel mehr Aufwand: für Generalisierung, Testabdeckung, Dokumentation und fortlaufende Wartung. Außerdem muss sich das Programm in ein größeres System einfügen. Vom Programm zum Programmsystemprodukt erhöht sich der Aufwand um ungefähr eine ganze Größenordnung.
2. Vom Mythos des Mann-Monats
   > Konventionelle Planungsmethoden funktionieren nur, wenn Personen und Monate austauschbar sind. Das beachtet jedoch weder die unterschiedlichen Kompetenzniveaus verschiedener Teammitglieder noch den zusätzlichen Zeitaufwand durch Kommunikation innerhalb des Teams sowie zwischen Teams. Da Programmierer aber schlecht im Schätzen von Zeitaufwand sind, wird in der Praxis der Zeitplan oft an den Bedürfnissen der Kunden ausgerichtet, nicht an den technischen Randbedingungen. Wenn dann das Kind in den Brunnen gefallen ist, tritt das Brooks'sche Gesetz in Kraft: "Das Hinzufügen weiteren Personals zu einem verspäteten Softwareprojekt verspätet es weiter."
3. Das OP-Team
   > Um die Kommunikationskosten zu verringern, stellt Brooks ein Modell für kleinere Teams vor, die in Analogie zu einem OP-Team organisiert sind, in dem immer nur einer auf einmal die eigentliche Programmierarbeit leistet und alle anderen nur zuarbeiten. Die genaue Beschreibung ist sicherlich überholt: Heutzutage braucht man zum Beispiel keinen Schreiber mehr, der Lochkarten sortiert. Aber vielleicht haben wir auch etwas verloren, als wir Sekretärsrollen ersetzt haben dadurch, dass jetzt die rare Expertin Termine im Outlook schubsen und Präsentationsfolien zusammenklicken muss.
4. Aristokratie, Demokratie und System-Design
   > Die Kathedrale von Reims wurde von acht Generationen von Baumeistern errichtet, weist aber trotzdem eine beachtliche gestalterische Kohärenz auf, weil die nachfolgenden Generationen das ursprüngliche Design respektiert haben. Daraus muss aber nicht folgen, dass ein einzelner Architekt alle Entscheidungen trifft und die Erbauer nur stupide die Pläne umsetzen. Auch auf der Umsetzungsebene gibt es viele interessante Designentscheidungen zu treffen; und ein guter Architekt hört auch auf seine Erbauer, wenn es um Fragen von Praktikabilität geht.
5. Das Problem des zweiten Systems
   > In diesem Dialog zwischen Architekt und Erbauer stellen sich also oftmals Ideen des Architekten als unpraktikabel heraus und werden verworfen. Doch wenn das erste System erfolgreich war, gewinnt der Architekt an Prestige und Erfahrung, die dazu verleiten kann, beim nächsten Mal alle Sachzwänge zu ignorieren. Ein weiteres klassisches Diktum von Brooks ist deswegen: "Dieses zweite ist das gefährlichste System, dass man jemals entwirft."
6. Weitergabe von Informationen
   > Damit ein Entwurf von 10 Architekt:innen von 1000 Mitarbeiter:innen und nochmal deutlich mehr Nutzer:innen verstanden wird, braucht es gute Dokumentation: Handbücher für Benutzer, und Architekturpläne für Erbauer. Brooks schlägt hier einiges vor, was heute in der Industrie gelebte Praxis ist (oder zumindest sein sollte), zum Beispiel maschinenlesbare Spezifikation von Schnittstellen und automatisierte Tests zur Überprüfung der Kompatibilität. Sein Ruf nach sorgfältiger Pro/Kontra-Abwägung zwischen mehreren möglichen Design-Optionen verhallt im Zeitalter agiler Softwareentwicklung jedoch oftmals ungehört.
7. Warum ist der Turm zu Babel gefallen?
   > Offensichtlich ist ineffektive Kommunikation einer der wesentlichen Gründe dafür, dass Software-Großprojekte scheitern oder erst mit deutlicher Verspätung liefern. Somit stellt sich die Frage nach der optimalen Teamstruktur in einer Software-Entwicklungsabteilung. Brooks hält eine baumförmige Struktur für unausweichlich, damit Entscheidungskompetenzen klar geregelt sind. In jedem Entwicklungsteam sollte es sollte es neben einem Architekten (dem Chefentwickler des Teams) auch einen Produzenten geben, der sich um die Kommunikation mit anderen Teams kümmert. Da eine Personalunion zwischen beiden Rollen nicht skaliert, muss entweder der Architekt der Chef des Produzenten sein: Dies passt aber nicht zu den separaten Management-Karrieren in den großen Firmen. Oder der Produzent ist der Chef des Architekten: Dies kann aber nur funktionieren, wenn der Produzent die technische Autorität des Architekten akzeptiert.

### Kapitel 8: "Calling the Shot"

- "Calling the Shot" hier im Baseball-Sinne: vorher ansagen, wo der Ball nach dem Schlag landet
    - Wie lange wird eine System-Programmierungsaufgabe dauern? Wieviel Aufwand ist nötig? Wie schätzt man das?

- Problem: Code schreiben alleine ist noch einigermaßen abschätzbar, macht aber nur einen kleinen Teil (vllt. ein Sechstel) des Gesamtaufwandes aus
    - Aufwand steigt exponentiell mit der Komplexität des Programmiersystems, weil Interaktionen zwischen verschiedenen Teilen des Systems beachtet werden müssen und Kompromisse im Design erfordern (vgl. Kapitel 1)
    - und dann zusätzlicher Overhead, sobald Kommunikation zwischen mehreren Teammitgliedern erforderlich ist (vgl. Kapitel 2)

- Brooks fasst dann Erkenntnisse aus mehreren quantitativen Studien zusammen; Highlights:
    - Geschwindigkeitsschätzungen waren zu Beginn immer ziemlich akkurat, aber im Verlauf des Projektes wird mit zunehmender Komplexität alles immer träger
    - Aufwand scheint am stärksten von der Gesamtmenge an Code abzuhängen (sprich: 1000 Zeilen Code sind ähnlicher Aufwand, unabhängig von der Sprache) -> Überleitung in das nächste Kapitel

### Kapitel 9: Zehn Pfund in einem Fünf-Pfund-Sack

- wenn Code kein Wert an sich, sondern eine Verbindlichkeit ist, wie kann man Code reduzieren?
    - Hochsprachen -> vgl. "No Silver Bullet" und intrinsiche vs. versehentliche Komplexität ("essential vs. accidental complexity"; besprochen im [Pentaradio vom Dezember 2022](https://c3d2.de/news/pentaradio24-20221227.html))
    - Xyrill sieht einen Bezug zum Hype um Coding-KI und die damit verbundenen Werbeversprechen

- Brooks zieht eine Parallele zu Hardware-Entwurf: "Weil die Größe ein wichtiger Teil der Benutzerkosten eines Programmiersystem-Produktes ist, muss der Erbauer Größenziele setzen, die Größe im Auge behalten, und Techniken zur Größenreduktion ersinnen; genau so, wie ein Hardware-Erbauer Komponentenzahlziele setzt, die Komponentenzahl im Auge behält, und Techniken zur Reduktion der Komponentenzahl ersinnt. Wie jede Art von Kosten ist Größe an sich nicht schlecht, aber überflüssige Größe ist es."
    - vgl. Kapitel 1 (siehe STP068), wo bereits ein vordefiniertes Ressourcenbudget für jedes Teilprogramm gefordert wurde
    - heute (leider?) nicht mehr en vogue; Speicher ist zu günstig geworden

- "Repräsentation ist die Essenz des Programmierens": Brooks hält den Entwurf passender Datenstrukturen für sehr viel relevanter als die Auswahl passender Algorithmen
    - Xyrill stimmt offensichtlicherweise zu, siehe STP071/072
    - "Zeig mir deine Flussdiagramme ohne die [Datenbank-]Tabellen, und ich werde verwirrt bleiben. Zeig mir deine Tabellen, und ich werde die Flussdiagramme wahrscheinlich nicht brauchen; sie werden offensichtlich sein."
    - Xyrill vermutet, dass hierher seine tiefe Abneigung gegen Prozess-Designer rührt

### Kapitel 10: Die dokumentarische Hypothese

> Die Hypothese: Verborgen in einem riesigen Papierhaufen verborgen, werden eine kleine Zahl von Dokumenten zu kritischen Angelpunkten, um die sich das Management eines jeden Projektes dreht. Dies sind die zentralen Werkzeuge eines Managers.

- "Auf einen neuen Manager, der gerade aus dem Handwerk gewechselt ist, erscheint der meiste Papierkram als vollständiges Ärgernis, als unnütze Ablenkung, und eine weiße Flut, die droht, ihn zu umspülen. Und in der Tat sind die meisten dieser Dokumente genau das."
    - doch einige Dokumente sind essentiell und das wesentliche Arbeitsergebnis des Managers
    - Mindestsatz laut Brooks: Zielsetzungen/Spezifikationen, Zeitplan, Budget, Ressourcenzuteilung, Personalzuteilung (letzteres vgl. Conway's Gesetz: "Organisationen, die Systeme entwerfen, sind gezwungen, Systeme hervorzubringen, die die Kommunikationsstruktur dieser Organisation kopieren.")

- wesentliche Funktionen von Dokumentation für den Manager:
    - Festhalten von Entscheidungen: Aufschreiben legt Widersprüche, Denkfehler und -lücken offen
    - Kommunizieren von Entscheidungen: Manager sind die meiste Zeit Kommunikatoren und nur relativ selten tatsächlich Entscheider
    - Erinnern von Entscheidungen: klare Dokumentation einer Entscheidung mit Begründung ermöglicht spätere Rezension und Entscheidung, wann eine Richtungsänderung angebracht ist

### Kapitel 11: Plane damit, eines wegzuwerfen ("Plan to Throw One Away")

- Analogie zu **Pilotanlagen** in der Chemietechnik: beim Übergang von Laborprozessen zu Fabrikproduktion im großen Maßstab sind Skalierungsprobleme und methodische Fragen zu lösen, sodass man nicht sofort die Fabrikanlage in der finalen Größe bauen kann
    - "Plane damit, eines wegzuwerfen. Du wirst es sowieso tun."
    - Programmierer waren zu Brooks' Zeiten damit beschäftigt, die Lektion zu lernen, und sind es heute immer noch

- damit im Zusammenhang: Wie geht man mit Veränderungen der Umstände um? Wieviel Veränderung der Umstände ist man bereit, zu akzeptieren?
    - These: Unwillen zur Dokumentation ist Unwillen, sich auf eine Entscheidung festzulegen, von der man schon weiß, dass sie immer nur vorläufig sein kann
    - bis hin zu Teamstruktur: "In einem großen Projekt muss der Manager zwei oder drei Topprogrammierer als technische Kavallerie vorhalten, die zu Rettung galoppieren können, wo immer der Kampf am schwierigsten ist." (Xyrill ist sich nicht sicher, ob er da mitgeht)

- Software-Maintenance laut Brooks besteht hauptsächlich aus "Änderungen, die Designfehler beheben"
    - Xyrill stimmt nur im Teil zu: der meiste Aufwand bei Maintenance heutzutage ist "Bit Rot" (Umgang mit der stetig fortschreitenden Fäulnis im Ökosystem), also nur indirekt Behebung von Designfehlern
    - jede Behebung eines Fehlers kann aber auch neue Fehler erzeugen (sowohl auf handwerklicher Ebene als auch auf Design-Ebene)
    - automatisierte Tests helfen, sofern sie eine hinreichend gute Abdeckung erzielen und nicht so aufwändig sind, dass sie den Arbeitsprozess behindern
    - trotzdem: Programme werden im Verlauf ihres Lebenszyklus immer komplexer, wodurch die Rate von Folgefehlern steigt

- der letzte Satz des Kapitels war bereits im Juni im Pentaradio zitiert: "Das Erbauen eines Programms ist ein entropieverringernder Prozess […] Die Wartung eines Programms ist ein entropieerhöhender Prozess, und selbst die gescheiteste Umsetzung kann das unausweichliche Absacken des Systems in unreparierbare Obsoleszenz höchstens verzögern."

### Kapitel 12: Scharfe Werkzeuge

- Programmierer sind eine der wenigen Berufsstände, die ihre eigenen Werkzeuge bauen
    - da Effizienz aus der Qualität des Instrumentariums folgt, sollen Softwareentwicklungs-Manager Ressourcen abstellen für "Werkzeugmacher"
    - aber: heute anderer Fokus als damals; weniger Bedarf nach z.B. Debugging-Werkzeugen (das bringt die Programmiersprache mit) und mehr Bedarf nach z.B. CI/CD oder Linting

- Brooks bespricht dann Strategien, die nicht auf den allgemeinen Fall übertragbar sind
    - Rationierung der Rechenzeit auf raren Mainframes: heute haben wir überproportionierte PCs für jeden Entwickler, sowie elastische Cloud-Infrastruktur für alles, was da nicht heraufpasst
    - Umgang mit Prerelease-Hardware: betrifft die allermeisten großen Software-Projekte nicht (aber siehe [dieser Vortrag von einem AMD-Chipdesigner](https://media.ccc.de/v/32c3-7171-when_hardware_must_just_work))

- Brooks stellt als offensichtlich wertvolles Werkzeug Hochsprachen auf eine Stufe mit interaktiver Programmierung
    - "Ich bin überzeugt, dass nur Schwerfälligkeit und Faulheit die universelle Annahme dieser Werkzeuge aufhalten; die technischen Schwierigkeiten sind nicht länger valide Ausreden."
    - Vorteile von Hochsprachen: erhöhte Produktivität (aufgrund weniger Code; siehe Kapitel 8/9), weniger Programmierfehler (nicht nur aufgrund weniger Code, sondern weil ganze Fehlerklassen eliminiert werden)
    - bei interaktiver Programmierung Bezug zu MIT Multics (dem Vorgänger von AT&T Unix): Xyrill hatte hier zuerst an Systeme wie Smalltalk gedacht, aber offenbar ist "interaktiv" hier das Gegenwort zu "Batch Processing"

### Kapitel 13: Das Ganze und die Teile

- Wie verhindert man Probleme beim Zusammenfügen mehrerer Teilprogramme zu einem Gesamtsystem?
    - Brooks enumeriert hier viele bereits erwähnte Sachen: strukturierte Programmierung in Hochsprachen, Debugging (entweder interaktiv auf der Zielmaschine, oder post-mortem mit einem Speicherabbild), Tests (sowohl für Einzelkomponenten als auch das Gesamtsystem)

- Missverständnisse bzw. Unklarheiten in der Spezifikation: siehe Kapitel 4-6 für Prozesse zum Erhalt der konzeptionellen Integrität des Gesamtsystems
    - Xyrill beobachtet, dass das Schreiben der Spezifikation oftmals mehr Konzentration erfordert als das Schreiben des Code (hohe Entscheidungsdichte, dichtes Konsequenzgeflecht)
    - Brooks empfiehlt einen Ansatz nach Niklaus Wirth: "Programmentwicklung durch schrittweises Verfeinern" mit verschiedenen Stufen von je nach Stufe möglichst groben Designs, die dann jeweils durch die nächste Stufe verfeinert werden
    - Probleme beim Verfeinern stellen dann evtl. die Designfehler auf der gröberen Ebene bloß -> damit der Versuchung entgehen, ein schlechtes Grunddesign mit immer mehr Epizyklen zuzukleistern

- Empfehlung: klare Markierungen, wenn zum Zwecke der Fehlerbehebung der dokumentierte Zustand des Systems manipuliert wird
    - Brooks beschreibt Verkabelung gemäß Plan mit gelben Kabeln und manuelle Änderungen mit lilanen Kabeln, damit diese ins Auge stechen und dann beim nächsten Update des Kabelplans berücksichtigt werden
    - in Software nicht so sehr formalisiert, wie es sein sollte (am ehesten in Form von Systemen, die manuelle Änderungen am Produktivsystem unterbinden)

- weitere Empfehlungen:
    - immer eine Komponente nach der anderen hinzufügen, nicht mehrere auf einmal
    - Aktualisierungen nicht ständig einspielen, sondern nur ab und zu, damit man zwischendurch den stabilen Zustand des Systems beobachten kann

### Kapitel 14: Wie man eine Katastrophe ausbrütet ("Hatching a Catastrophe")

- Zeitpläne scheitern meist nicht durch ein großes Hindernis (auf welches mit einer großen Anstrengung reagiert würde), sondern durch eine Anhäufung kleiner Hindernisse
    - vgl. politische Reaktion auf Versagen: Einsturz einer einzelnen Brücke mobilisiert sofort Unterstützung, aber schleichender Verfall aller Brücken ist alternativlos
    - Tod durch tausend Schnitte: heute ist der Chefentwickler krank, morgen ist eine Festplatte auf dem Datenbankserver kaputt, übermorgen verspätet sich der Hardwaretest aufgrund Problemen in der Lieferkette

- Erkennen von Verspätungen erfordert einen festen Zeitplan mit spezifischen Meilensteinen
    - Gegenbeispiel: "Implementation 90% fertig" stimmt so gut wie immer :)
    - "Chronisches Überziehen des Zeitplans ist ein Teammoral-Killer."

- "Aber die anderen sind auch zu spät dran."
    - trotzdem ist es ratsam, sich selber ranzuhalten, damit man nicht auf dem [kritischen Pfad](https://de.wikipedia.org/w/index.php?title=Methode_des_kritischen_Pfades&oldid=249068045) landet und dann das ganze Großprojekt ins Stocken bringt
    - Manager sollten den kritischen Pfad kennen, damit sie rechtzeitig wissen, wo ihre Aufmerksamkeit am wertvollsten ist

### Kapitel 15: Das andere Gesicht

- "Ein Computerprogramm ist eine Nachricht des Menschen an die Maschine. \[...] Aber ein geschriebenes Programm hat ein anderes Gesicht, dasjenige, dass seine Geschichte dem menschlichen Benutzer erzählt."
    - selbst für Programme für den eigenen Gebrauch, damit man sich in einem Jahr noch daran erinnert, was das Programm tut

- Brooks möchte, dass Dokumentation aus drei Teilen besteht
    - wie man ein Programm verwendet (klar)
    - wie man dem Programm glaubt (Testfälle zum Ausprobieren)
    - wie man das Programm verändert (Dokumentation über die interne Struktur, das Dateilayout, verwendete Algorithmen und Datenstrukturen)

- Brooks findet Strukturdiagramme super, und Flowcharts scheiße: man solle lieber äquivalenten (Pseudo-)Code hinschreiben
    - Xyrill stimmt zu: die meisten Flowcharts aus der Praxis sind abgrundtief hässlich und unlesbar chaotisch
    - Ergebnis: der Code sollte einfach selbstdokumentierend sein bzw. seine Dokumentation inline enthalten (vgl. Donald Knuth und "Literate Programming")
    - offensichtlich ist das alles aus der Zeit, bevor Software ein einzeln verkauftes Produkt war

- im Gespräch erwähnt: [The Architecture of Open Source Applications](https://aosabook.org/en/index.html); ein Buch darüber, guten Open-Source-Code zu lesen und zu durchdringen

### Epilog

"Die Teergrube der Software-Entwicklung wird noch eine ganze Weile klebrig bleiben. Man kann davon ausgehen, dass die Menschheit auch weiterhin Systeme in Angriff nehmen wird, die gerade so innerhalb oder gerade so außerhalb des Möglichen liegen; und Softwaresysteme sind vielleicht die vertracktesten und komplexesten aller menschlichen Schöpfungen. Das Verwalten dieses komplexen Handwerks wird von uns fordern: den bestmöglichen Einsatz neuer Sprachen und Systeme, die optimale Anwendung bewährter Methoden des Ingenieurmanagements, großzügige Mengen gesunden Menschenverstandes, und die gottgegebene Demut, die eigene Fehlbarkeit und Beschränkungen zu erkennen."

### Podcast-Empfehlungen

* [Gilda con Arne &ndash; Der Politik-Podcast](https://www.acbstories.com/work/gilda-con-arne)
* Dlf Doku Serien
  * einige Trailer für Dokus
  * [Tech Bro Topia](https://www.deutschlandradio.de/tech-bro-topia-100.html)
  * [Rechtsextreme vor Gericht](https://www.deutschlandradio.de/rechtsextreme-vor-gericht-dokuserie-100.html)
  * [Neuland](https://www.deutschlandfunkkultur.de/hoerspielmagazin-gespraech-neuland-sechsteilige-feature-100.html)
* [Die Geschichte geht weiter &ndash; Victor Klemperers Tagebücher 1918-1958](https://www.deutschlandradio.de/die-geschichte-geht-weiter-100.html)
