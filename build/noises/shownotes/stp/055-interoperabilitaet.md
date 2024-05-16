- historischer Einstieg: initial nur einzelne monolithische Programme
    - Computer startet, sobald das Programm eingegeben ist, zusammen mit dem Programm
    - mehrere Programme möglich, aber diese alle komplett voneinander unabhängig; bzw. jeweils nur abhängig von der Maschine selbst
    - dann wurden Programme so komplex, dass sie nicht mehr von einer Person verfasst werden konnten
    - 1975: akademische Unterscheidung zwischen ["Programming in the large" und "Programming in the small"](https://en.wikipedia.org/w/index.php?title=Programming_in_the_large_and_programming_in_the_small&oldid=1184277192)
    - Unterteilung der Software eines Computers in kleinere Einheiten erfordert definierte **Schnittstellen**
    - Analogie: Unternehmen mit mehreren Mitarbeiterinnen erfordern meist definierte Abläufe

- Schnittstellen auf der großen Ebene (zwischen Prozessen): [Kommunikationsprotokolle](https://de.wikipedia.org/w/index.php?title=Kommunikationsprotokoll&oldid=225866784) und [Datenformate](https://de.wikipedia.org/w/index.php?title=Dateiformat&oldid=240585393)
    - Kommunikationsprotokolle: regeln Sende-/Empfangsreihenfolge, Inhalte der einzelnen Datenpakete und das durch sie ausgelöste Verhalten, Fehlerbehandlung etc.
    - Datenformate: definieren die Struktur eines Datenpaketes; kann in einer Datei gespeichert sein; kann Teil eines Kommunikationsprotokolls sein; kann Baustein in einem größeren Datenformat sein (z.B. Videodatei besteht aus einem Audiodatenformat, einem Bildstromdatenformat und einem umschließenden Containerformat)

- Schnittstellen auf der kleinen Ebene (innerhalb eines Prozesses): [API (Application Programming Interface)](https://en.wikipedia.org/w/index.php?title=API&oldid=1196841027) und [ABI (Application Binary Interface)](https://en.wikipedia.org/w/index.php?title=Application_binary_interface&oldid=1192371815)
    - z.B. eine Komponente, die einen Sortieralgorithmus bereitstellt, hat eine Schnittstelle, die eine Liste von Zahlen entgegennimmt und eine sortierte Liste von Zahlen zurückgibt
    - dies ungefähr die geringste Ebene von Komplexität, Skala nach oben offen
    - API: auf der Ebene von Programmcode
    - ABI: auf der Ebene von Maschinencode
    - Beispiel API: [OpenGL](https://en.wikipedia.org/w/index.php?title=OpenGL&oldid=1198092106) vs. [DirectX](https://en.wikipedia.org/w/index.php?title=DirectX&oldid=1197353025) vs. [Vulkan](https://en.wikipedia.org/w/index.php?title=Vulkan&oldid=1199515232)
    - [XKCD 927](https://xkcd.com/927/)

- Interoperabilität mit der Brechstange: [Reverse Engineering](https://de.wikipedia.org/w/index.php?title=Reverse_Engineering&oldid=241382265)
    - siehe auch STP040 (§69e UrhG: Dekompilierung und Reverse Engineering sind unter bestimmten Umständen auch ohne Zustimmung des Urheberrechteinhabers erlaubt)

- Interoperabilität mit etwas sanfterem Hebel: [Emulation](https://de.wikipedia.org/w/index.php?title=Emulator&oldid=240982806)
    - siehe auch STP023

- siehe auch: Berufsbezeichnung [Systemintegrator](https://de.wikipedia.org/w/index.php?title=Systemintegrator&oldid=240254359)
    - Unterdisziplin des Ausbildungsberufes "Fachinformatiker", in Abgrenzung zum Anwendungsentwickler
    - Tätigkeitsschwerpunkt: Anpassung bestehender Systeme zum Zwecke der Interoperabilität
    - vor allem im Umfeld von Firmensoftware verbreitet (Warenwirtschaftssystem des Herstellers A muss mit dem Buchhaltungssystem des Herstellers B verbunden werden)
