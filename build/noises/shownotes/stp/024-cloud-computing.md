* [Cloud Computing erklärt anhand von Pizza](https://blog.pragmaticworks.com/hs-fs/hubfs/Pizza_Cloud.jpg?width=1868&name=Pizza_Cloud.jpg)
    * On-Premise: kein Cloud Computing, Nutzer betreibt seine eigene Hardware und Software
    * IaaS (Infrastructure as a Service): Cloud-Anbieter betreibt die Hardware, Nutzer bringt Betriebssystem und Software
    * PaaS (Platform as a Service): Cloud-Anbieter betreibt die Hardware und das Betriebssystem, Nutzer bringt die Anwendung
    * SaaS (Software as a Service): Cloud-Anbieter betreibt alles, Nutzer bedient die Anwendung nur noch

Fünf Grundanforderungen gemäß [Definition durch NIST](http://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-145.pdf ) (NIST ist dem Aufgabenspektrum nach mit der PTB und dem BSI vergleichbar):

* **On-Demand Self-Service:** Nutzer kann selber Ressourcen klicken, ohne Mitarbeiter beim Dienstanbieter zu involvieren
    * Bsp. SaaS-Anwendungen wie Dropbox, GitHub, Netflix, etc.
    * Bsp. mal kurz ein dutzend VMs klicken, um eine komplizierte Server-Anwendung wie Kubernetes auszuprobieren

* **Broad Network Access:** Ressourcen sind über das Netz mittels "Standardmethoden" erreichbar, ohne durch eine anbieterspezifische Zugriffstechnologie gehen zu müssen
    * Was soll diese Anforderung ausschließen? Xyrills Vermutung: z.B. das Telefonnetz, ttimeless' Vermutung: Mainframes und Universitäts-Rechenzentren (die nur aus dem entsprechenden Intranet zu erreichen sind)
    * zur Konfiguration und Administration jedoch meist anbieterspezifische Schnittstellen und Werkzeuge (entsprechend dem Funktionsumfang des jeweiligen Angebots) -> **Vendor Lock-In**

* **Resource Pooling:** Ressourcen werden unter mehreren/allen Kunden nach Bedarf aufgeteilt
    * sowohl zwischen mehreren Kunden ("Public Cloud") als auch innerhalb einer Firma oder innerhalb eines Teams ("Private Cloud")
    * Bsp. CDN (Content Delivery Network): mehrere Kunden teilen sich ein weltweites Rechenzentrumsnetz, das jeder alleine so nicht bezahlen könnte

* **Rapid Elasticity:** Ressourcen können kurzfristig zugeteilt und zurückgegeben werden
    * Bsp. Internet-Shop am Black Friday, Buchhaltungssystem am Quartalsende
    * Bsp. Cloud-Gaming (Gegenbeispiel: [Error 37](https://knowyourmeme.com/memes/error-37 ))

* **Measured Service:** Messung und Abrechnung der genutzten Ressourcen mindestens zeitnah oder sogar in Echtzeit
    * "wie bei der Telefonrechnung früher"
    * Bsp. sehr kurze Nutzung (z.B. einige Stunden) von leistungsstarker Hardware, z.B. um ein neuronales Netz zu trainieren

* Probleme mit Cloud Computing
    * Zentralisierung: wenn ein Amazon-Rechenzentrum Probleme hat, ist gefühlt das halbe Internet weg
    * Preistransparenz: siehe z.B. ["I got pwned by my cloud costs"](https://www.troyhunt.com/how-i-got-pwned-by-my-cloud-costs/ )
