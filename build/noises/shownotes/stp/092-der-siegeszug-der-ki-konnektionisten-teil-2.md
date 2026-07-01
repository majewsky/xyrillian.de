- Nochmal zurück zum ersten KI-Winter (siehe STP091): Warum skalieren reine Perzeptrone schlecht? Bsp. Bildererkennung
    - Bild in Thumbnail-Größe (z.B. 256 Pixel im Quadrat, 3 Farbkanäle) besteht aus `256 * 256 * 3 = 196608` Einzelwerten
    - Abbildung im ersten Layer des Perzeptrons auf z.B. 1000 Neuronen würde das Training von 196 Millionen Gewichten erfordern, und das ist nur der erste Layer!
    - so viele Gewichte können nur mit sehr vielen Eingabedaten sinnvoll trainiert werden -> enormer Rechenaufwand für Training und Inferenz

- [Faltende neurale Netzwerke (Convolutional Neural Network, CNN, ConvNet)](https://de.wikipedia.org/w/index.php?title=Convolutional_Neural_Network&oldid=264196705)
    - Idee 1: pro Layer nicht einfach alle Eingabewerte mit allen Ausgabeneuronen verknüpfen, sondern eine bestimmte Struktur vorgeben (und alle nicht dazu passenden Verbindungen verbieten), um die Menge von Gewichten zu reduzieren
    - Idee 2: Mustererkennung basiert auf Merkmalen (*Features*), die unabhängig von ihrer genauen Position im Bild immer gleich erkannt werden sollen
    - **Faltungslayer** (*Convolution Layer*): aus einer Eingabe mit `N * N` Werten fasst jedes Feature-Neuron eine Teilregion der Größe `K * K` zusammen (z.B. K = 11: 121 Gewichte pro Neuron); Anzahl von Ausgabeneuronen basierend auf der gewünschten Anzahl `F` unterschiedlicher Features entweder `N * N * F` (bei Abtastrate 1) oder geringer (bei weniger dichter Abtastung)
        - alle Neuronen desselben Features verwenden dieselben Gewichte -> enorme Reduktion des Trainings- und Inferenzaufwands
    - **Zusammenfassungslayer** (*Pooling Layer*): Reduktion einer Eingabe mit `N * N * F` Werten in z.B. `N/2 * N/2 * F` Werte durch Zusammenfassen von nichtüberlappenden Regionen (z.B. `2x2`)
        - hier keine trainierbaren Gewichte, sondern Vorauswahl einer Zusammenfassungsfunktion wie `max` oder `avg`
    - am Ende meist einige Ebenen klassische Perzeptron-Layer mit voll trainierbaren Gewichten zwischen allen Neuronen
    - [Beispiel: Architekturen früher Bilderkennungssysteme](https://en.wikipedia.org/w/index.php?title=Convolutional_neural_network&oldid=1337481938#/media/File:Comparison_image_neural_networks.svg)
    - Tuning-Entscheidungen: Anordnung der Layer, Anzahl der Features, Faltungsgrößen, Abtastraten, Aktivierungsfunktion (heute meist [ReLU](https://de.wikipedia.org/w/index.php?title=Rectifier_%28neuronale_Netzwerke%29&oldid=245277150))

- ab 2015: [Diffusionsmodelle](https://en.wikipedia.org/w/index.php?title=Diffusion_model&oldid=1337976238)
    - grundsätzlicher Ansatz: durch den Datenraum gerichtete Diffusion von Daten von einem Startpunkt zu einem höherwertigen Endpunkt
        - während des Trainings fängt man mit hochwertigen Daten an und fügt immer mehr Rauschen hinzu, wodurch die Daten von ihrem Ausgangszustand "wegdiffundieren"
        - Modell wird darin trainiert, diesen Diffusionsprozess umzukehren
    - im Bereich der Bildverarbeitung z.B. [Super-resolution imaging](https://en.wikipedia.org/w/index.php?title=Super-resolution_imaging&oldid=1327501054) (mit einem niedrig aufgelösten Bild als Startpunkt) oder Bilderzeugung (mit reinem Rauschen als Startpunkt)
        - Problem: Wenn man mit nur Rauschen anfängt, wie entscheidet man, wo man hin will?
    - 2021: [CLIP (Contrastive Language-Image Pre-Training)](https://en.wikipedia.org/w/index.php?title=Contrastive_Language-Image_Pre-training&oldid=1338246766)
        - Trainingsdatensatz: Bilder mit möglichst akkuraten Textbeschreibungen ihrer Inhalte
        - gleichzeitiges Training von zwei Encoder-Transformern, die einmal Bilder und einmal Textbeschreibungen in Embeddings übersetzen
        - Training der Encoder versucht, die Ähnlichkeit der Embeddings zwischen passenden Bildern und Textbeschreibungen zu maximieren und zwischen unpassenden Paaren zu minimieren
    - Bilderzeugung aus Diffusion: Text-Encoder übersetzt den Prompt in ein Embedding, Diffusionsmodell denoised damit ein initial komplett verrauschtes Bild in das Ausgabebild
        - beim Training des Diffusionsmodells muss zu den Trainingsbildern das Embedding der entsprechenden Beschreibung präsentiert werden
        - mögliche Realisierung: Diffusionsmodell als Transformer mit Self-Attention innerhalb des Bildes und Cross-Attention zum Embedding der Textbeschreibung
    - Videotipp: [3Blue1Brown zu Diffusionsmodellen](https://www.youtube.com/watch?v=iv-5mZ_9CPY); dort zum Beispiel zur Kodierung von "Negative Prompts" (z.B. "no extra fingers", "no JPEG compression residue")

Der gespielte O-Ton ist aus [Deutschlandfunk Hintergrund: "10 Jahre OpenAI: Wie ChatGPT mit KI die Welt verändert"](https://www.deutschlandfunk.de/10-jahre-openai-wie-chatgpt-mit-ki-die-welt-veraendert-100.html).

Xyrills Abendgedanken: Nach dieser Ausarbeitung verstehe ich jetzt, warum die Konnektionisten Transformer als den Durchbruch schlechthin und als Allheilmittel einsetzen. Aber, nochmal aus [Gary Marcus's Blogartikel](https://de.wikipedia.org/w/index.php?title=Induktion_(Philosophie)&oldid=261188485), der in STP089 referenziert war:

> Für immer gefangen im Land der Korrelationen, können neurale Netzwerke niemals (egal, wie viel Daten oder Rechenleistung ihnen zur Verfügung stehen) in der Lage sein, kausale Beziehungen zu verstehen -- warum Dinge so sind, wie sie sind -- und daraus folgend kausale Schlussfolgerungen zu ziehen. Dieser kritische Teil der menschlichen Kognition ist, warum Menschen nur einmal die Straßenverkehrsregeln in einer Stadt lernen müssen, um dann in vielen anderen Städten sicher fahren zu können, argumentiert Marcus. Teslas Autopilot hingegen kann Milliarden von Kilometern an Fahrten aufzeichnen und trotzdem crashen, wenn er in ein unbekanntes Szenario gerät oder mit ein paar strategisch plazierten Aufklebern überlistet wird. Marcus plädiert dafür, stattdessen Konnektionismus und Symbolismus miteinander zu verbinden.

Und weiter:

> Viele der Fortschritte heutiger KI-Modelle resultieren wahrscheinlich aus der Ausweitung des Einsatzes symbolischer Werkzeuge, und nicht aus mehr Skalierung. Riesige Infrastruktur-Investitionen wie Stargate basieren wahrscheinlich auf falschen Annahmen darüber, was tatsächlich den Fortschritt vorantreibt.

Als Schlusswort dieselbe Kritik an der Skalierungshypothese, aber weniger technisch ausgedrückt, [von Quinn Norton in "What we talk about when we talk about AI"](https://emptywheel.net/what-we-talk-about-when-we-talk-about-ai-part-one/):

> Menschen sind anders. Trotz der Entlehnung von Nomenklatur aus der Biologie haben die neuralen Netzwerke, die im Training von KI verwendet werden, keine Neuronen der menschlichen Art. Der Unterschied ist erkennbar. Wir lernen Sprache, Lesen und Schreiben anhand eines winzigen Datensatzes, und dieser Prozess umfasst Mimikry, Gefühle, Kognition und Liebe. Da mag auch statistische Gewichtung vorkommen, aber wenn dem so ist, dann haben wir den entsprechenden Mechanismus noch nicht in unserem Verstand und unseren Gehirnen gefunden. Es erscheint unwahrscheinlich, dass es dort in einer ähnlichen Form existieren würde, da diese KI-Systeme so viel mehr Informationen und Rechenleistung aufwenden müssen, um zu tun, was ein Erstsemesterstudent mit ein bisschen Motivation hinbekommt. Motivation ist unser Problem, aber es ist niemals ein Problem für KI-Systeme. Sie machen immer weiter, bis ihre Befehle einen Endpunkt erreichen, und dann erstarren sie. KI-Systeme sind leblos zu Beginn, leblos währenddessen, und leblos am Ende.

