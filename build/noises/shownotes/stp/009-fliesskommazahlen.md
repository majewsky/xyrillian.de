* Rückbezug zu STP001: Ganzzahlen
    * z.B. 8-Bit-Zahlen: ohne Vorzeichen 0 bis 255 oder mit Vorzeichen -128 bis 127
    * Problem 1: Wie stellt man Zahlen mit Nachkommastellen dar?
    * Problem 2: Wie stellt man sehr große Zahlen effizient dar?
    * zwei Optionen: Festkommazahlen und Fließkommazahlen

* **Festkommazahlen**
    * im Prinzip dasselbe wie Ganzzahlen, aber zählt in Schritten von 0,1 oder 0,01 etc.
    * Beispiel: Währungsrechnungen oft mit 4 Nachkommastellen (also auf 0,01 Eurocent genau)
    * gegenüber Fließkommazahlen: schneller und genauer, aber nicht so flexibel (löst nur Problem 1, nicht Problem 2)

* **Fließkommazahlen**
    * grundsätzliche Idee: analog zur [wissenschaftlichen Zahlennotation](https://de.wikipedia.org/wiki/Wissenschaftliche_Notation)
        * z.B. "1,2 Milliarden": nicht als `1.200.000.000`, sondern `1,2 x 10^9`
        * z.B. "2 Hunderttausendstel": nicht als `0,000002`, sondern `2 x 10^-6`
    * Vorteil: unabhängig von der Größenordnung kann immer die gleiche Zahl an Ziffern dargestellt werden
    * in Computern natürlich nicht zur Basis 10, sondern zur Basis 2

* Bestandteile einer Fließkommazahl
    * **Exponent** (E): die Anzahl der Zweierpotenzen
    * **Mantisse** (M): der Teil der Zahl, der nicht in der Potenz steht
    * insgesamt also `M * 2^E` für positive oder `-M * 2^E` für negative Zahlen

* Rechenoperationen ungefähr analog zum Rechenschieber
    * z.B. Addition: erst die Exponenten angleichen, dann die Mantissen addieren, dann eventuell nochmal den Exponenten anpassen, damit die finale Darstellung wieder schön kompakt ist
    * hier kommt das Überlaufbit aus STP003 wieder ins Spiel

* Kodierung: fast immer gemäß [IEEE-754-Standard](https://de.wikipedia.org/wiki/IEEE_754 )
        * [IEEE](https://de.wikipedia.org/wiki/Institute_of_Electrical_and_Electronics_Engineers )
    * Problem: wir wollen uns nicht merken müssen, wo in der Mantisse das Komma steht
    * Trick: wir schieben das Komma so hin, dass die Mantisse mit `1,` anfängt, und speichern dann nur den Teil dahinter (**normalisierte Zahlen**)
    * Ausnahme: beim kleinsten Wert des Exponenten fängt die Mantisse mit `0,` an, damit wir auch die Zahl 0 und sehr kleine Zahlen nahe 0 darstellen können (**subnormale Zahlen**)
    * z.B. für 32 Bit: [einfache Genauigkeit](https://de.wikipedia.org/wiki/Einfache_Genauigkeit ) ("Single Precision")
        * 1 Bit Vorzeichen
        * 8 Bit Exponent (Ganzzahl zwischen -127 und 128)
        * 23 Bit Signifikand (Mantisse ohne das führende `1,` oder `0,`)
    * Werte je nach Exponent:
        * Exponent = -127 -> subnormale Zahlen
        * Exponent = 128 -> spezielle Werte
        * ansonsten -> normalisierte Zahlen
    * im Exponent 128 werden einige spezielle Werte kodiert, die man braucht, damit jede Rechenoperation immer eine gültige Fließkommazahl erzeugt:
        * **plus Unendlich**: z.B. `1 / 0 = +∞` oder `∞ + ∞ = ∞`
        * **minus Unendlich**: z.B. `-1 / 0 = -∞` oder `log2(0) = -∞`
        * **keine Zahl** (not a number, NaN): z.B. `∞ - ∞ = NaN` oder `0 / 0 = NaN` oder `log2(-1) = NaN`

* Implementation von Fließkommarechnungen
    * initial in Software auf Basis von Ganzzahlen
    * seit den 1980ern zunehmend in Hardware (seit den 1990ern auch in PCs)

* Fallstricke bei Fließkommarechnungen
    * nicht alle kompakten Dezimalzahlen sind kompakte Binärzahlen: z.B. `0.3/0.2 = 1.4999999999999998`
    * bei komplexen Berechnungen können sich kleine Rundungsfehler zu großen Fehlern aufschaukeln, siehe zum Beispiel [Kahan-Summenalgorithmus](https://en.wikipedia.org/wiki/Kahan_summation_algorithm ) (leider in Wikipedia nur auf englisch)

* Anwendungsbereiche üblicher Fließkommaformate
    * doppelte Genauigkeit (64 Bit): für wissenschaftliche Berechnungen
    * einfache Genauigkeit (32 Bit): für 3D-Grafik (siehe nächste Folge)
    * halbe Genauigkeit (16 Bit), Viertelgenauigkeit (8 Bit): für Training von neuronalen Netzwerken
