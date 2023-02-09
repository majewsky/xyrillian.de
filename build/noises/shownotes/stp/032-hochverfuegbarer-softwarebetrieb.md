- Schlagwort: [Site Reliability Engineering (SRE)](https://de.wikipedia.org/w/index.php?title=Site_Reliability_Engineering&oldid=218196182 )
    - Hauptziel: "Schaffung skalierbarer und hochzuverlässiger Softwaresysteme"

- Grundkonzept: Service Level
    - [SLO (Service Level Objective)](https://en.wikipedia.org/w/index.php?title=Service-level_objective&oldid=1089474535 ): selbstgewählter Indikator soll bis zu einem bestimmten Grad an 100% herangeführt werden (z.B. Zeiten, zu denen der Dienst erreichbar ist; Anteil der Anfragen, die in weniger als 1 Sekunde beantwortet werden; Anteil der Anfragen, die erfolgreich beantwortet werden)
    - SLA (Service Level Agreement): SLO, der mit dem Kunden vertraglich vereinbart ist
    - Beispiel: 99% Verfügbarkeit = bis zu 3,65 Tage Ausfall; 99,9% = bis zu 8,77 Stunden; 99,99% = bis zu 52 Minuten; 99,999% = bis zu 5,2 Minuten

- Problem: die Unbekannten in der Gleichung
    - **Known Unknowns** (bekanntes Unwissen): ich weiß, dass ich etwas nicht weiß (z.B. wann meine Festplatte ausfallen wird)
    - **Unknown Unknowns** (unbekanntes Unwissen): ich weiß nicht mal, dass da etwas ist, was ich nicht weiß (z.B. [dass meine angebliche 30-TB-Festplatte in Wahrheit nur eine 32-GB-Speicherkarte in einem großen Gehäuse ist](https://www.vice.com/en/article/akek8e/walmart-30tb-ssd-hard-drive-scam-sd-cards ))
    - siehe auch: [zur Wortherkunft der Begriffsgruppe "(un)known (un)knowns"](https://de.wikipedia.org/w/index.php?title=There_are_known_knowns&oldid=222128371 )

- strukturierter Einblick in Applikationen mittels Instrumentierung
    - **Logging:** Ereignisse innerhalb der Applikation werden als textförmiger Bericht (Log) aus der Applikation ausgeleitet und in einem Log-System durchsuchbar gemacht
    - **Monitoring:** Zustand der Applikation wird in Form von Messdaten quantifiziert, in einer [Zeitseriendatenbank](https://de.wikipedia.org/w/index.php?title=Zeitreihendatenbank&oldid=211779199 ) abgelegt und in Dashboards aufbereitet ([Beispielbild](https://en.wikipedia.org/w/index.php?title=Grafana&oldid=1123211485#/media/File:Grafana_dashboard.png )); Messdaten entweder direkt aus der Applikation selbst oder durch ein Extraprogramm von außen erhoben
    - **Alarmierung:** Menschen werden benachrichtigt, wenn ein bestimmter Messwert eine bestimmte Schwelle überschreitet (oder eine Überschreitung projiziert ist, z.B. "die Festplatte läuft gerade voll und in einer Stunde ist sie komplett voll")

- strukturierte Herangehensweisen beim Durchführen von Änderungen
    - Configuration as Code: Änderungen werden möglichst nie manuell vorgenommen, sondern es wird ein Programm geschrieben (bzw. die Konfiguration für jenes Program angepasst), das die Änderungen vornimmt; bei Problemen ist ein Zurückrollen auf den letzten funktionierenden Zustand möglich
    - [Continuous Integration (CI)](https://de.wikipedia.org/w/index.php?title=Kontinuierliche_Integration&oldid=223309917 ): Änderungen an Code und Konfiguration fließen in ein System, das automatisch standardisierte Tests ausführt, um die Änderung auf Fehler zu prüfen; bei Erfolg automatische Übernahme der vorgeschlagenen Änderung in den veröffentlichten Code-Stand
    - [Continuous Delivery (CD)](https://de.wikipedia.org/w/index.php?title=Continuous_Delivery&oldid=218889976 ): automatisches oder stark automatisiertes Ausrollen von Änderungen an Code und Konfiguration von Testsystemen in Verifikationssysteme in Produktivsysteme
    - Rolling Upgrade: bei Systemen mit mehreren gleichartigen Komponenten werden die Komponenten nacheinander durch die neue Version ersetzt, damit das Ausrollen für den Kunden unsichtbar ist und bei Fehlern frühzeitig gestopppt werden kann
    - Canary Deployment: Ausrollen einer Änderung zunächst nur in einem kleinen Teil des Systems; bei Problemen ist nur der kleine Teil betroffen (siehe auch Ankündigungen der Form "diese Funktion ist zurzeit nur für ausgewählte Kunden verfügbar")

- strukturierter Umgang mit Alarmen/Fehlern
    - standardisierte Abläufe ("Playbooks") für Alarme
    - Postmortem-Analyse mindestens nach signifikanten Fehlern (Was ist passiert? Was war der tatsächliche Grund? Was ist bei der Bekämpfung des Problems gut/schlecht gelaufen? Was können wir besser machen?)
    - Fehlerkultur: Das Hauptziel soll nicht sein, einen Schuldigen zu finden, sondern das Problem möglichst schnell zu beheben und dann dafür zu sorgen, dass das Problem nie wieder auftritt.
    - Parallelen zum [Toyota-Produktionssystem](https://de.wikipedia.org/w/index.php?title=Toyota-Produktionssystem&oldid=223213898 )

- im Gespräch erwähnt (im Kontext des Gesetzesvorhabens zur Chatkontrolle):
    - [Pentaradio vom Mai 2022](https://www.c3d2.de/news/pentaradio24-20220524.html) und ganz aktuell [vom Januar 2023](https://c3d2.de/news/pentaradio24-20230124.html)
    - [LNP429: Endlich sind wir am Anfang angekommen](https://logbuch-netzpolitik.de/lnp429-endlich-sind-wir-am-anfang-angekommen)
    - [LNP430: Die größte Koalition aller Zeiten](https://logbuch-netzpolitik.de/lnp430-die-groesste-koalition-aller-zeiten)
