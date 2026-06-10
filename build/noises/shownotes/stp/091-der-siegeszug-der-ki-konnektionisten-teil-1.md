Damit wir uns nicht komplett verzetteln, beschränken wir uns in diesem Zweiteiler auf alle Entwicklungen, die zu den aktuellen Buzzword-Technologien geführt haben: Transformer und Diffusionsmodelle.

- [KI-Winter](https://de.wikipedia.org/w/index.php?title=KI-Winter&oldid=261879950): Flaute in der KI-Forschung nach dem Scheitern eines bestimmten Ansatzes
    - frühe 1970er: erster KI-Winter nach Problemen mit dem Skalieren von Perzeptronen (Scheitern konnektionistischer Ansätze)
    - späte 1980er: zweiter KI-Winter nach allgemeiner Ernüchterung über Expertensysteme (Scheitern symbolischer Ansätze)
    - danach neuer Aufwind für konnektionistische Ansätze durch neue Optimierungen und Moore's Law ([Skalierungshypothese](https://en.wikipedia.org/w/index.php?title=Neural_scaling_law&oldid=1338998332))

- [Rekurrente neurale Netzwerke (Recurrent Neural Network, RNN)](https://de.wikipedia.org/w/index.php?title=Rekurrentes_neuronales_Netz&oldid=257632743)
    - Grundproblem: Wie geht man mit zeitlich geordneten Daten um, die sich entlang des Zeitstrahls kausal beeinflussen (z.B. Sprache, Aktienkurse)?
    - Idee: neurales Netzwerk verarbeitet immer die Eingaben aus einem Zeitschritt und gibt zusätzlich zur Ausgabe einen internen Zustand aus, der beim nächsten Zeitschritt zur Eingabe wird
    - Problem: Wie trainiert man das? Backpropagation funktioniert nicht, wenn Signale zu früheren Layern rückwärts geschickt werden
        - frühe Ansätze hatten Probleme mit hohem Trainingsaufwand, weil Backpropagation in Schleifen laufen muss
        - hier aus Zeitgründen nur der Ansatz, der heutzutage die größte Bedeutung hat

- 2014: [seq2seq](https://en.wikipedia.org/w/index.php?title=Seq2seq&oldid=1326493865) zum Übersetzen natürlicher Sprache
    - zwei separate aufeinanderfolgende Schritte
        - Encoder: Lesen jeweils eines Eingabesymbols (**Token**); hierbei keine Ausgabe, sondern nur Aktualisierung eines internen Zustands
        - Decoder: Extraktion eines Ausgabetokens aus dem internen Zustand unter Präsentation des zuletzt generierten Tokens, dabei ebenfalls Aktualisierung des internen Zustandes
    - Abwägung
        - Vorteil: damit keine Probleme mit rückwärts laufenden Signalen, Training mit Backpropagation möglich
        - Nachteil: interner Zustand muss die gesamte Eingabe darstellen können, bei langen Eingabesequenzen "vergisst" der interne Zustand Teile der Eingabe
        - im Gespräch erwähnt: [Hutter Prize](https://en.wikipedia.org/w/index.php?title=Hutter_Prize&oldid=1326108826)
    - Lösung: [Attention](https://youtube.com/watch?v=eMlx5fFNoYc); im Falle von seq2seq:
        - interner Zustand nach jedem Encoding wird gemerkt
        - Decoding aus dem aktuellsten internen Zustand (wie zuvor) ergibt ein vorläufiges Ergebnis, das mit einem Satz trainierbarer Gewichte in eine **Query** (Abfrage) überführt wird (anschauliche Interpretation z.B. "Wir generieren jetzt das Adjektiv, das vor 'Garten' steht. Wer weiß was?")
        - analog dazu werden die gemerkten internen Zustände des Encodings mit einem anderen Satz trainierbarer Gewichte in einen **Key** (Schlüssel) überführt (anschauliche Interpretation z.B. "Ich bin ein Adjektiv.")
        - ...und außerdem mit einem dritten Satz trainierbarer Gewichte in einen **Value** (Wert) (anschauliche Interpretation z.B. "Das zuletzt gesehene Adjektiv ist 'groß'.")
        - Values, bei denen Query und Key gut aufeinanderpassen, werden mit Gewichtung aufaddiert zu einem **Kontext**: `C = softmax(Q * K^T) * V` mit [Softmax-Funktion](https://de.wikipedia.org/w/index.php?title=Softmax-Funktion&oldid=263971912)
        - tatsächliches Ausgabezeichen aus Kombination des Kontext mit dem ursprünglichen Ergebnis des Decoding

- 2017: [Transformer](https://de.wikipedia.org/w/index.php?title=Transformer_%28Maschinelles_Lernen%29&oldid=260701397) verwerfen die rekurrente Struktur und arbeiten nur mit Attention, um die Parallelisierbarkeit zu verbessern (daher der berühmte Paper-Name "Attention is *all* you need")
    - jedes Eingabesymbol (**Token**) wird in ein **Embedding** überführt, einen hochdimensionalen Vektor (z.B. GPT-3: je nach Modellvariante zwischen 768 und 12288 Dimensionen)
        - Ziel des Embedding: Darstellung von Konzepten in einem Vektorraum ([Beispiel](https://commons.wikimedia.org/wiki/File:Word_embedding_illustration.svg))
    - bei autoregressiver Texterzeugung (autoregressiv = "aus sich selbst heraus"):
        - **Self-Attention** auf den Embeddings der bis jetzt vorliegenden Tokenfolge (eventuell begrenzt durch das **Context Window**), Resultat ist ein aktualisierter Embedding-Vektor
        - aber: Attention ist eine lineare Funktion und kann nicht als universeller Funktionsapproximator fungieren
        - deswegen dahinter ein klassisches dicht verschaltetes Feed-Forward-Netzwerk, in dem die überwältigende Masse der trainierbaren Gewichte sitzt
        - Ergebnis: der finale Embedding-Vektor, der in ein Token zurückübersetzt wird
    - bei Texterzeugung anhand einer separaten Eingabesequenz (z.B. maschinelle Übersetzung, Spracherkennung, [RAG](https://de.wikipedia.org/w/index.php?title=Retrieval-Augmented_Generation&oldid=264408330)) außerdem **Cross-Attention** zum Transfer von Informationen aus der Eingabesequenz in den Dekoder

Abendgedanken aus ["Large Language Models, Model Collapse and the Conservation of Information"](https://www.youtube.com/watch?v=ShusuVq32hc): Es ist überraschend, wieviele Probleme sich auf "das nächste Wort vorhersagen" reduzieren lassen (vgl. Meldungen der Form "Goldmedaille in der Matheolympiade"). Aber:

1. LLMs wägen nicht ab, sie verarbeiten.
    - statistische Aussagen auf Basis von Embeddings, nicht mehr und nicht weniger
    - ungeeignet für symbolische Aufgaben wie Arithmetik
2. LLMs überlegen nicht, sie rationalisieren.
    - die statistische Natur der Sache ist gut für [Induktion](https://de.wikipedia.org/w/index.php?title=Induktion_%28Philosophie%29&oldid=261188485), aber schlecht für Deduktion
3. LLMs können nicht endlose Informationen erzeugen.
