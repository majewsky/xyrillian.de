- Rückbezüge:
    - fundamentale Datenstrukturen siehe STP071: Felder/Listen, ~~assoziative Datenfelder (Maps), Graphen~~
    - Frage: Wie stellt man solche Datenstrukturen im Speicher dar? Gibt es darauf überhaupt die eine richtige Antwort?
    - algorithmische Komplexität siehe STP029: Liste mit `n` Elementen ausdrucken in `O(n)`, aber sortieren in `O(n log(n))` bis `O(n^2)`
    - Speicherallokation siehe STP047 und Speicherschutz siehe STP019: Bezug wird gleich klar werden

- Listen kann man als [verkettete Liste darstellen](https://de.wikipedia.org/w/index.php?title=Liste_(Datenstruktur)&oldid=253638160)
    - klassisches Studienobjekt in Erstsemester-Datenstrukturen-Vorlesungen
    - intuitiv verständlich: Parallele zu segmentierten Halsketten
    - wahlweise einfach oder doppelt verkettet
    - effiziente Operationen: Einfügen am Ende, Entfernen am Ende
    - ineffiziente Operationen: Einfügen in der Mitte, Wahlzugriff/Suche
    - [Vergleichstabelle mit Darstellung der Zeiteffizienz](https://en.wikipedia.org/w/index.php?title=Template:List_data_structure_comparison&oldid=1251607720)
    - alles in allem durchwachsene Performance -> geht es besser?

- alternative Strategie: interne Darstellung der Liste als [balancierter Baum](https://de.wikipedia.org/w/index.php?title=Balancierter_Baum&oldid=243407080) (oder evtl. "ausgeglichener Baum")
    - außerdem Link auf [die englische Wikipedia](https://en.wikipedia.org/w/index.php?title=Self-balancing_binary_search_tree&oldid=1273477858), die nicht nur unbalancierte, sondern auch balancierte Bäume zeigt
    - kann nur sortierte Listen darstellen
    - Idee: Wurzelknoten hat das Median-Element, dann der linke Ast alle kleineren und der rechte Ast alle größeren Elemente
    - im Grunde alle gängigen Operationen mittelschnell: Wahlzugriff/Suche, Einfügen an beliebigen Stellen, Entfernen von beliebigen Stellen (Änderungen erfordern im Allgemeinen ein Austarieren des Baumes)
    - große Variation von Implementationsstrategien für dieses Balancieren -> hier nicht

- verkettete Listen und balancierte Bäume sehen auf dem Papier ziemlich effizient aus, haben aber in ihrer reinen Form pathologisch schlechtes Speicherverhalten
    - hoher Platzverbrauch: z.B. bei einfach bzw. doppelt verketteten Listen muss zu jedem Element müssen noch eine bzw. zwei Speicheradressen abgelegt werden
    - hohe Allokationslast: wenn nicht eine Arena oder ein vergleichbarer Small Object Allocator verwendet wird
    - schlechte Lokalität: beim Durchlaufen nachfolgender Elemente werden im schlimmsten Fall ständig unterschiedliche Speicherseiten getroffen, was fortlaufend Seitenfehler verursachen kann

- in der Praxis mit Abstand dominante Implementationsstrategie: [dynamische Felder](https://en.wikipedia.org/w/index.php?title=Dynamic_array&oldid=1268522792)
    - Beobachtung: Einfügen oder Löschen an beliebigen Stellen wird kaum gemacht; man hängt eher mehrmals ans Ende an und sortiert dann, falls nötig
    - Idee: Optimieren auf Einfügen am Ende bei möglichst optimalen Speicherverhalten
    - Umsetzung: einfaches Feld (ein fortlaufendes Stück Speicher, in dem mehrere Elemente hintereinander abgelegt werden) mit aktuellem Füllstand N und Kapazität K
    - Einfügen ans Ende: normalerweise einfach N erhöhen; wenn N nicht in K passt, größeren Speicher reservieren, alles hinüberkopieren und die alten Speicherallokation verwerfen
    - Löschen vom Ende: einfach N reduzieren, keine Deallokation erforderlich
    - Einfügen am Anfang oder in die Mitte: alle Elemente dahinter nach hinten verschieben

- nahezu optimales Speicherverhalten von dynamischen Feldern begründet ihre Dominanz
    - Platzverbrauch: neben der Speichergröße der Elemente selber nur zwei Zahlen (N und K)
    - Allokationslast: Vergrößern geht für gewöhnlich in exponentiellen Schritten und wird damit für wachsende Listen immer seltener nötig
    - Lokalität: lineare Suche durch die Liste geht linear durch den Speicher hindurch -> folgende Speicherseiten können vom Prozessor oft schon auf Verdacht vorgeladen werden
    - trotzdem: die anderen Datenstrukturen haben auch ihre Berechtigung (z.B. modifizierte balancierte Bäume als Basis für Datenbank-Indizes)

- Xyrill will auch noch auf analoge Weise über Maps und Graphen reden -> weiter in STP079
