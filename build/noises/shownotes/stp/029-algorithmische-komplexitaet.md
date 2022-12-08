- **Komplexität**: Ressourcenbedarf eines Algorithmus (Zeitkomplexität, Platzkomplexität; vielleicht auch Kostenkomplexität etc.)
    - Forschungsgegenstand der Komplexitätstheorie
    - hier hauptsächlich Zeitkomplexität

- einfaches Beispielproblem: "Gegeben ist eine sortierte Liste von Wörtern (Schlagwörter im Wörterbuch). Finde ein bestimmtes Wort."
    - naive Idee: jedes Wort nacheinander anschauen -> lineare Laufzeit (Verdopplung der Wörtermenge verdoppelt die zu erwartende Suchzeit)
    - kluge Idee: [Bisektion](https://de.wikipedia.org/w/index.php?title=Bisektion&oldid=194714676 ) -> logarithmische Laufzeit (Verdopplung der Wörtermenge erfordert nur einen zusätzlichen Schritt)

- Beschreibung von Komplexität: [**Landau-Symbole**](https://de.wikipedia.org/w/index.php?title=Landau-Symbole&oldid=217959469 )
    - wichtigste Form: `f = O(g)` heißt, dass `f` "nicht wesentlich schneller als `g` wächst (sprich: `f(n)/g(n)` bleibt beschränkt, wenn `n -> ∞`)
    - Beispiel: naive Suche über eine Liste von `n` Elementen hat eine Laufzeit in `O(n)`, Bisektion hat eine Laufzeit in `O(log n)`

- etwas komplexeres Beispiel: Sortieralgorithmen ("Gegeben ist eine Liste von Zahlen/Wörtern/etc. Sortiere diese Liste.")
    - Live-Beispiel: Skatkarten mischen und dann sortieren -> Auf welche Art und Weise sortieren wir intuitiv?
    - Sortieralgorithmen auf Computern: zwei Grundbausteine
        - Vergleich von zwei Elementen
        - Tauschen von zwei Elementen
    - anhand von Skatkarten diskutierbar:
        - [Insertionsort](https://de.wikipedia.org/w/index.php?title=Insertionsort&oldid=222921729 ): `O(n^2)`
        - [Quicksort](https://de.wikipedia.org/w/index.php?title=Quicksort&oldid=224184691 ): `O(n log n)`
        - [Mergesort](https://de.wikipedia.org/w/index.php?title=Mergesort&oldid=223738956 ): `O(n log n)`
    - wünschenswerte Eigenschaften:
        - Lokalität ([Timsort](https://de.wikipedia.org/w/index.php?title=Timsort&oldid=217156701 )!, siehe STP007)
        - Stabilität (Mergesort!)
        - Codegröße (Quicksort!)

- Schauempfehlung (mit Epilepsie-Warnung): [Visualisierung verschiedener Sortieralgorithmen](https://www.youtube.com/watch?v=8MsTNqK3o_w )

- im Gespräch erwähnt:
    - [Bacon-Zahl](https://de.wikipedia.org/w/index.php?title=Bacon-Zahl&oldid=214690502 )
    - [Erdős-Zahl](https://de.wikipedia.org/w/index.php?title=Erd%C5%91s-Zahl&oldid=222414205 )
    - [How large is 52 factorial?](https://czep.net/weblog/52cards.html )
