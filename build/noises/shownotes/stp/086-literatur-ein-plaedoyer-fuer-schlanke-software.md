Anhand von ["A Plea for Lean Software" (1995)](https://berthub.eu/articles/LeanSoftware_text.pdf) von Niklaus Wirth besprechen wir zwei wesentliche Fragestellungen. Zuerst einige Vorbeobachtungen:

- Frage 1: Warum ist Software (bzw. Produkte im Allgemeinen) immer weniger langlebig?
    - COBOL-Software aus den 60ern ist immer noch im Einsatz, aber React-Apps von vor 5 Jahren sind hoffnungslos obsolet
        - siehe Vermutung, dass in der Frontend-Entwicklung Schnelllebigkeit von YouTube-Influencern getrieben wird
        - bei COBOL/Fortran-Systemen auch regulatorische Hürden (z.B. Steuerungsoftware für Atombomben in Fortran 77, weil ein [Umstieg auf neuere Versionen Tests der Software erfordern würde](https://infosec.exchange/@david_chisnall/115463730854960344), aber Atomwaffentests verboten sind)
    - vgl. Erfahrung von ttimeless: Fensterbeschläge von vor 30 Jahren noch top, aber neuere von vor 5 Jahren tw. schon abgeranzt (Geplante Obsoleszenz? Survivorship-Bias?)
        - Wie schätzt man Qualität bei selten gekauften Produkten ab?
        - im Gespräch erwähnt: [Backblaze Drive Stats](https://www.backblaze.com/blog/author/drivestatsteam/)
    - sich gegenseitig beschleunigende Marktmechanismen (Qualitätsschwund durch Sparmaßnahmen in der Produktion, Mode)
        - [passend dazu](https://www.reddit.com/r/BuyItForLife/comments/1106bml/why_everything_you_buy_is_worse_now_feels/): "Ich sah neulich eine alte Episode von 'Der Preis ist heiß' aus den 80ern, die UK-Edition mit \[Bruce Forsyth\]. Die Preise waren Dinge wie Mikrowellen, Fernseher, Möbelstücke und so weiter. Die Preise waren weit, weit höher, als was ich je geraten hätte. Wir sind mit der Vorstellung aufgewachsen, dass eine Mikrowelle weniger als 100 Pfund kosten sollte, obwohl sie unter Berücksichtigung der Inflation mittlerweile 400 kosten sollten. Alles ist scheiße günstig geworden."
    - These von ttimeless: mehr Leute machen mehr Zeug machen mehr Altlasten (Xyrill erhöht um eine Coding-KI-Industrie und möchte sehen)

- Frage 2: Warum waren Computer in den 90ern schneller?
    - welch Ironie, dass wir uns nach der Zeit zurücksehnen, in der der Autor unseres Textes sich über verschwenderische Programmierweisen aufregt
    - vgl. Frage im MDR: Wann kriegen wir wieder Akkulaufzeiten von einer Woche wie bei den alten Nokia-Handys?
    - offensichtlicher Einfluss: mehr Leistung, die sich nicht als konkrete Funktion äußert (z.B. höhere Auflösung erfordert mehr Videospeicher, moderne Konnektivität ist objektiv komplexer, als nur auf einen Anruf oder eine SMS zu warten)
    - [Gegenrede](https://news.ycombinator.com/item?id=37538521): "Ich war dabei bei den meisten dieser alten Maschinen und meiner Erinnerung nach waren sie *immer* langsam, weswegen es ein derart unerbittliches Rennen um das Upgrade auf eine schnellere Maschine gab, und warum für eine lange Zeit Leute jedes Jahr ihren alten Computer ersetzten. \[...] Dieser Topos von 'wie konnten Dinge langsamer werden, damals war alles so schnell' ist eine fehlerhafte Erinnerung"
    - [andere Perspektive aus demselben Thread](https://news.ycombinator.com/item?id=37538305): "Weil wir erweiterbare anstatt zusammensetzbarer Software schreiben. Wir liefern nicht Funktionen aus, sondern Apps. Wieviele Instruktionen braucht man, um in Maschinensprache 1+1 auszurechnen? Wieviele Instruktionen braucht man, um die Suchleiste zu öffnen, den Taschenrechner zu finden und die Berechnung auszuführen?"
    - vgl. [Freak Show 273](https://freakshow.fm/fs273-papst-von-china) im Kapitel mit dem Nachruf auf Niklaus Wirth: in den 60ern vorherrschende Vorstellung unmittelbar bevorstehender Ablösung von Programmierern durch Computer, aber mehr Rechenleistung erfordert noch mehr Programmierer (vgl. [Softwarekrise](https://de.wikipedia.org/w/index.php?title=Softwarekrise&oldid=257790161), für den Verweis auf kombinatorische Explosion von Software-Komplexität siehe STP070)

- Komplexitätsquellen laut Wirth
    - "verbesserte Benutzerführung und Funktionalität" laut Wirth keine gute Begründung
        - aber Wirth begründet das mittels seines eigenen extremen Fokus auf Basisfunktionalität ("Diese grundlegenden Aufgaben haben sich \[... angesichts u.&nbsp;a.&nbsp;der\] Ersetzung von aussagekräftigen Befehlswörtern durch hübsche Bilder \[nicht geändert\].")
        - Desktop-Metapher und Skeuomorphismus waren essentiell darin, Computersoftware einem breiten Publikum näherzubringen
        - vgl. Brooks (essentielle vs. versehentliche Komplexität), aber unterschiedliche Grenzziehungen
        - Wirth bevorzugt ziemlich explizit "simple" statt "easy"
    - jede Funktion, die der Kunde haben will, wird eingebaut, ob sie sich nun sinnfällig in das ursprüngliche Design einfügt oder nicht
        - Komplexität durch Komplikationen
    - monolithischer Systementwurf ("Jeder Kunde zahlt für alle Funktionen, aber nutzt nur wenige von ihnen.")
        - vgl. oben "wir liefern nicht Funktionen aus, sondern Apps"
        - Xyrill baut Systeme auch eher monolitisch, weil sich für ihn der Wartungsaufwand verringert, wenn alles in einer Codebase liegt
    - Komplexität wird an sich als Funktionalität missverstanden: "Immer mehr scheinen die Leute Komplexität als Raffinesse zu interpretieren, was rätselhaft ist &mdash; Unverständliches sollte Argwohn erregen, nicht Bewunderung."
        - \*hust\* \*hust\* ✨KI✨ \*hust\* \*hust\*
        - kapitalistischer Anreiz: damit Nutzer in der eigenen Teergrube binden und Support und Beratung verkaufen
    - gutes Design erfordert systematische Iteration auf dem ersten, offensichtlichen, überkomplizierten Entwurf
        - vgl. Antoine de Saint-Exupéry ("Es scheint, dass Perfektion nicht dann erreicht ist, wenn nichts mehr hinzuzufügen ist, sondern wenn nichts mehr wegzunehmen ist.")
        - aber: Zeitdruck ("akribische Ingenieursroutine zahlt sich kurzfristig nicht aus")
        - Wirth drischt hier stark auf die Kunden ein, die ständig neuen Klimbim wollen, was etwas aus der Zeit gefallen wirkt (bzw. nur etwas deplatziert; heute sind "die Kunden" die Werbeabteilung, die noch mehr Datenspuren sammeln will)
    - deswegen auch extrem langsame Einführung von Techniken für methodisches Arbeiten (z.B. stark typisierte Programmiersprachen): "Methodisch korrektes Design \[...\] ist offenbar unerwünscht, weil auf diese Weise entwickelte Produkte zu viel Zeit bis zur Markteinführung brauchen."
        - Wirth beschwert sich, dass in den 90ern die Programmiertheorie der 70er (z.B. abstrakte Datentypen) erst langsam in Produktivumgebungen Einzug fand
        - passend dazu kommt zurzeit die Programmiertheorie der 90er (z.B. algebraische Datentypen wie in STP072, affine Typsysteme) erst langsam in die heutigen Produktivumgebungen, aber C hält sich immer noch hartnäckig

- Wie heraus aus der Teergrube? Wirth bespricht sein eigenes [Oberon System](https://de.wikipedia.org/w/index.php?title=Oberon_(Programmiersprache)&oldid=258424717)
    - ursprüngliches Ziel: ein Betriebssystem inklusive Userspace, das kompakt genug ist, um als Lehrmaterial verwendet und in Lehrbüchern besprochen werden zu können
    - extreme Konzentration auf "Grundfunktionen" (z.B. keine grafische Oberfläche, nur Textinteraktion)
    - Implementation in einer typsicheren Programmiersprache mit Objektorientierung, um starke Modularität auf Codeebene zu ermöglichen
        - Xyrill ist hier etwas skeptisch, dass Objektorientierte Programmierung hier einen großen Unterschied macht, aber für Wirth ist das extrem wichtig
    - keine abgeschlossenen Programme, sondern nur Funktionsaufrufe
        - also vmtl. kein Speicherschutz in Hardware, sondern nur durch Typsicherheit
        - "Wenn ein Modul M eine Prozedur P anbietet, dann kann P aufgerufen werden, indem man einfach auf den Text 'M.P' zeigt, sobald er irgendwo auf dem Bildschirm erscheint, und eine Maustaste drückt"
        - statt einer Leiste mit Icons eine einfache Textbox mit Befehlsnamen
        - vgl. Plan 9, Smalltalk, Forth
    - der Preis des schlanken Systems ist der Aufwand für den sauberen Entwurf

Am Ende neun Lektionen für das Design schlanker Software. Wir zitieren Nummer 7:

> Die Reduktion von Komplexität und Größe muss in jedem Schritt das Ziel sein: im Lastenheft, im Design, und beim detaillierten Programmieren. Die Kompetenz eines Programmierers soll nach der Fähigkeit gemessen werden, einfache Lösungen zu finden, und ganz bestimmt nicht nach Produktivität, gemessen in "Anzahl Codezeilen, die pro Tag ausgeworfen werden". Produktive Programmierer tragen zum sicheren Desaster bei.
