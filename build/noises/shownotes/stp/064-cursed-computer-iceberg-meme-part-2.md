Wir setzen unsere Betrachtung des [Cursed Computer Iceberg Meme](https://suricrasia.online/iceberg/) aus STP042 fort. Das letzte Mal waren wir bis "Below the Water" gekommen. Xyrill hat aber noch ein paar Sachen in den oberen Ebenen gesehen, die besprechenswert sind.

* "Above the Iceberg" - Oberhalb des Eisberges
    * [Little Bobby Tables](https://xkcd.com/327/): das Mem-Beispiel einer SQL-Injektion; dieses wiederum eine klassische Art von Sicherheitslücken (siehe STP049)

* "On the Iceberg" - Auf dem Eisberg
    * [Quine](https://de.wikipedia.org/w/index.php?title=Quine_(Computerprogramm)&oldid=240777383): ein Programm, dass seinen eigenen Programmcode ausdruckt (Beispiele siehe Link)
        * siehe auch: [quine-relay](https://github.com/mame/quine-relay)
    * [Intel Management Engine](https://en.wikipedia.org/w/index.php?title=Intel_Management_Engine&oldid=1235943064)

* "Below the Water" - Unterhalb der Wasserlinie
    * [`rm -rf "$STEAM_ROOT"/*`](https://github.com/ValveSoftware/steam-for-linux/issues/3671): wie eine einzige uninitialisierte Variable unzählige Computer außer Gefecht setzt (heutzutage würde man sagen, diese Computer [wurden ge-crowdstriket](https://c3d2.de/news/pentaradio24-20240827.html))
    * ["lp0 on fire"](https://en.wikipedia.org/w/index.php?title=Lp0_on_fire&oldid=1236781886): heute eher ein Mem als alles andere, aber wie im Artikel beschrieben gab es früher Drucker mit Heizelementen, die bei Papierstau in Flammen aufgingen (und trotzdem weiter Papier nachschoben); Xyrill fühlt sich an die Szene aus "Fantasia" mit den belebten Besen erinnert
    * [Illegale Zahlen](https://de.wikipedia.org/w/index.php?title=Illegale_Zahl&oldid=245251846): eine konkrete Art, wie Redefreiheit in der Praxis durch Urheberrecht etc. eingeschränkt werden kann
        * siehe auch ["Welche Farbe haben deine Bits?"](https://ansuz.sooke.bc.ca/entry/23)
        * [Economics of everyday things - Happy birthday to you](https://freakonomics.com/podcast/happy-birthday-to-you/)
    * [Kaltstartattacke](https://de.wikipedia.org/w/index.php?title=Kaltstartattacke&oldid=240863372): ein sehr handfester Seitenkanalangriff
        * Es gibt auch die Möglichkeit, RAM einzufrieren, damit darin enthaltene Schlüssel forensisch ausgelesen werden können.
        * Manchmal sieht man auch Hausdurchsuchungen, bei denen die Polizei eine USV und Werkzeug zum Kabelspleißen mitbringt.

* "Middle of the Iceberg" - Mitte des Eisbergs
    * [42.zip](https://de.wikipedia.org/w/index.php?title=Archivbombe&oldid=246391169): gepackt nur 42 KiB; "beim Entpacken wächst das Datenvolumen allerdings um das Hundertmilliardenfache auf \[4 PiB]"
        * Rückbezug zu Quines: [Zip Files All The Way Down](https://research.swtch.com/zip)
    * ["OpenOffice kann dienstags nicht drucken"](https://bugs.launchpad.net/ubuntu/+source/cupsys/+bug/255161/comments/28): ein legendärer Fehlerbericht, fast so gut wie die 500-Meilen-Mail aus STP042

* "Bottom of the Iceberg" - Unterseite des Eisbergs
    * ["You fool. You absolute, unmitigated, unadulterated, complete and utter, fool."](https://gist.github.com/rjhansen/67ab921ffb4084c865b3618d6955275f): Manöverkritik zum Keypair-Spam-Angriff auf das GPG-Keyserver-Netzwerk (GPG hatten wir bereits am Rande in STP048 erwähnt)

Es gibt noch jede Menge im tiefen Wasser. Wir sehen uns wieder in STP084.
