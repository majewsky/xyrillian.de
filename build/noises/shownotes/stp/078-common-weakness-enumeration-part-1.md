- Rückbezug STP049: [CVE-Datenbank](https://de.wikipedia.org/w/index.php?title=Common_Vulnerabilities_and_Exposures&oldid=255367540) für bekannte Sicherheitslücken in bestimmten Programmen (z.B. [Heartbleed](https://de.wikipedia.org/w/index.php?title=Heartbleed&oldid=253665585) ist CVE-2014-0160)
    - betrieben von dem privaten Forschungsinstitut [Mitre Corporation](https://de.wikipedia.org/w/index.php?title=Mitre_Corporation&oldid=252423314) auf Basis von Fördermitteln der US-Regierung
    - April 2025: kurzzeitige Verwirrung nach Aussetzen des Fördervertrages, dann einen Tag später Kommando zurück ([Pentaradio berichtete](https://c3d2.de/news/pentaradio24-20250422.html))
    - neben CVE pflegt Mitre Corp. eine weniger bekannte Datenbank: [CWE (Common Weakness Enumeration)](https://en.wikipedia.org/w/index.php?title=Common_Weakness_Enumeration&oldid=1276415996), ein Klassifizierungssystem für Arten von Sicherheitslücken
    - wir besprechen die jährliche Rangliste der [CWE Top 25](https://cwe.mitre.org/top25/archive/2024/2024_cwe_top25.html) mit Stand von 2024
    - wertvoll, weil man so als interessierter Laie oder Programmierer-Generalist ein Gefühl für wichtige Angriffsvektoren bekommt

- Platz 1: [CWE-79](https://cwe.mitre.org/data/definitions/79.html); fehlerhafte Neutralisierung von Eingaben während der Erzeugung von Webseiten ("Cross-Site Scripting", XSS)
    - siehe Beispielbild im Link
    - Webseiten können JavaScript-Code enthalten, der auf dem Browser des Endanwenders ausgeführt wird
    - wenn Webseiten aus Eingaben anderer Benutzer zusammengebaut werden, können diese Code einschleusen
    - Lösung: Escaping (Entfernen oder Ersetzen von Zeichen mit syntaktischer Bedeutung), kann durch Template Engines teilweise automatisiert werden
    - passend dazu Platz 7: [CWE-78](https://cwe.mitre.org/data/definitions/78.html); fehlerhafte Neutralisierung von Eingaben während der Erzeugung von Befehlszeilen
    - außerdem Platz 13: [CWE-77](https://cwe.mitre.org/data/definitions/77.html); dasselbe für alle anderen Befehlssprachen, die nicht JavaScript oder die System-Shell sind
    - allgemeiner als Platz 12: [CWE-20](https://cwe.mitre.org/data/definitions/20.html); unzureichende Überprüfung von Eingaben
    - Lösung: "assume all input is malicious"; nur Eingaben akzeptieren, die strikten Formatvorgaben folgen (z.B. unter Einsatz von regulären Ausdrücken, siehe STP021)

- Platz 2, 6 und 20: [CWE-787](https://cwe.mitre.org/data/definitions/787.html) und [CWE-125](https://cwe.mitre.org/data/definitions/125.html); Schreiben bzw. Lesen außerhalb des allokierten Speicherbereichs ("Out-of-Bounds Write/Read") sowie als übergeordnete Kategorie [CWE-119](https://cwe.mitre.org/data/definitions/119.html); allgemeiner Pufferüberlauf ("Buffer Overflow")
    - und passend dazu Platz 8: [CWE-416](https://cwe.mitre.org/data/definitions/416.html); Benutzung nach dem Freigeben ("Use after Free")
    - alle bereits besprochen in STP047
    - Lösung: speichersichere Programmiersprachen, Speicherprofiling-Tools wie Valgrind (siehe STP076)

- Platz 3: [CWE-89](https://cwe.mitre.org/data/definitions/89.html); SQL-Injektion
    - SQL siehe STP012: Abfragesprache für relationale Datenbanksysteme
    - Beispiel siehe [XKCD 327](https://xkcd.com/327/)
    - Lösung: Datenwerte von der Abfrage trennen (z.B. `SELECT * FROM users WHERE name = 'Alice'` -> `SELECT * FROM users WHERE name = $1` mit Datensatz `["Alice"]`)
    - passend dazu Platz 11: [CWE-94](https://cwe.mitre.org/data/definitions/94.html); Code-Injektion (dasselbe für Programmcode); Überlapp zu CWE-77

- Platz 4: [CWE-352](https://cwe.mitre.org/data/definitions/352.html); Cross-Site Request Forgery (CSRF)
    - Angreifer möchte im Namen des Opfers eine Aktion ausführen (z.B. auf Rechnung des Opfers in einem Webshop bestellen)
    - Angreifer schiebt dem Opfer ein manipuliertes HTML-Formular unter, dass beim Abschicken die gewünschte Aktion auslöst, aber mit den Cookies (und damit der Identität des Opfers)
    - Lösung: CSRF-Tokens (echte HTML-Formulare erhalten ein einmalig nutzbares Token, dass nicht vom Angreifer gefälscht werden kann)
    - passend dazu Platz 19: [CWE-918](https://cwe.mitre.org/data/definitions/918.html); Server-Site Request Forgery (SSRF)
    - dabei manipuliert der Angreifer nicht ein menschliches Opfer oder einen Webclient, sondern einen Webserver, mit seinem privilegierten Zugriff andere Webserver abzufragen

- Platz 5: [CWE-22](https://cwe.mitre.org/data/definitions/22.html); Pfaddurchquerung ("Path Traversal")
    - z.B. bei Webservern mit statischen Inhalten: entsprechende Dateien liegen im Dateisystem in einem bestimmten Ordner
    - Dateien außerhalb dieses Ordners sollen eigentlich nicht zugreifbar sein, aber manchmal geht <http://example.com/../../../etc/passwd> oder dergleichen
    - Lösung: magische Pfadelemente wie `..` (bei Pfaden) oder `/` (bei Dateinamen) ablehnen

#### Feedback zu STP063

Alex schreibt:

> danke für den Podcast. Ich bin wie immer weit im Rückstand mit Hören, habe zu Folge 63 aber mal die Turing-Maschine für die Subtraktion neu gebastelt. \[...] Den Input muss man mit folgender 0 eingeben, also 1110110 z.B.
> Gebaut habe ich das auf <https://turingmachinesimulator.com/>.

```
//----------------------------------------------------------------------
//Syntax:

//-------CONFIGURATION
name: sub
init: q0
accept: qe

//-------DELTA FUNCTION:
//[current_state],[read_symbol]
//[new_state],[write_symbol],[>|<|-]

// < = left
// > = right
// - = hold
// use underscore for blank cells

q0,1 // 1en 1.Op überspringen
q0,1,>

q0,0
q1,0,>

q1,0 // 0en überspringen
q1,0,>

q1,1
q2,1,>

q2,0 // mehr als eine 1 im 2. Op? Ja->q5, nein->q3
q3,0,<

q2,1
q5,1,>

q5,1 // 1en 2.Op überspringen, dann auf letzte 1 des 2.Op
q5,1,>

q5,0
q6,0,<

q6,1 // letzte 1 des 2. Op durch 0 ersetzen
q7,0,<

q7,1 // zurück, 1en 2. Op überspringen
q7,1,<

q7,0
q8,0,<

q8,0 // zurück, 0en überspringen
q8,0,<

q8,1 // letzte 1 des 1. Op auf 0 setzen, zurück zu q1
q1,0,>

q3,1 // letzte 1 des 2. Op auf 0 setzen
q4,0,<

q4,0 // zurück, 0en überspringen
q4,0,<

q4,1 // letzte 1 des 1.Op auf Null setzen, zum Ende
qe,0,>
```
