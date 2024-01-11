> Die drei wichtigsten Dinge, die Experten über Software-Sicherheit wissen:
>
> 1. Software ist unglaublich unzuverlässig und unsicher.
> 2. Nein, wirklich, ihr könnt es euch nicht vorstellen.
> 3. Es ist sogar noch schlimmer als das.
>
> [Matt Blaze](https://en.wikipedia.org/w/index.php?title=Matt_Blaze&oldid=1177352443) [auf Twitter](https://twitter.com/mattblaze/status/1192503673664016385) (Achtung: zweiter Link geht zu einer Domain von Elon Musk)

- Vorbemerkung: für offensive und defensive Computernutzung wird regelmäßig der Wortteil "Cyber-" verwendet (z.B. "Cyberangriff", "Cybersicherheit")
    - in IT-Kreisen, insb. im CCC-Umfeld, gilt "Cyber-" als [Schibboleth](https://de.wikipedia.org/w/index.php?title=Schibboleth&oldid=237089412) für Leute, die von IT keine tiefgehende Ahnung haben (siehe z.B. ["Unsere Cyber-Cyber-Regierung"](https://www.youtube.com/watch?v=WY6KkRsS26M))
    - wir bleiben deswegen sicherheitshalber bei "IT-Angriff" :)

- Grundlage: [Sicherheitslücken](https://de.wikipedia.org/w/index.php?title=Sicherheitsl%C3%BCcke&oldid=229656948) ("Fehler in einer Software oder einer Hardware, durch \[die] ein Programm mit Schadwirkung oder ein Angreifer in ein Computersystem eindringen kann")
    - siehe STP037: heutige berühmte Softwarefehler sind meist Sicherheitslücken (z.B. Heartbleed, wie in STP037 besprochen)
    - siehe STP039: handwerkliche Fehler in der Authentifizierung oder Autorisierung ergeben meist besonders gut ausnutzbare Sicherheitslücken (siehe gerade neulich [Azure AD](https://msrc.microsoft.com/blog/2023/09/results-of-major-technical-investigations-for-storm-0558-key-acquisition/))
    - siehe STP045/STP047: Grundlage vieler Sicherheitslücken sind Speicherverwaltungsfehler (z.B. Pufferüberlauf, wie in STP047 besprochen)
    - Sicherheitslücken werden aktiv gesucht, entweder von "Whitehats" (Sicherheitsforschern oder den Red-Teams der Softwarehersteller selbst) oder von "Blackhats" (die die Lücken für tatsächliche Angriffe ausnutzen wollen oder an Angreifer verkaufen wollen)
    - bekannte Sicherheitslücken werden in der [CVE-Datenbank](https://de.wikipedia.org/w/index.php?title=Common_Vulnerabilities_and_Exposures&oldid=227659039) ("Common Vulnerabilities and Exposures", dt. "Bekannte Schwachstellen und Anfälligkeiten") geführt (z.B. Heartbleed ist CVE-2014-0160)

- zweiter Schritt: [Exploit](https://de.wikipedia.org/w/index.php?title=Exploit&oldid=229895345) (eine systematische Methode, um eine Sicherheitslücke auszunutzen)
    - auch für Whitehats sind Exploits wichtig, um die Kritikalität einer Sicherheitslücke zu demonstrieren
    - für Blackhats sind Exploits die handelbare Form von Sicherheitslücken
    - Preise für Exploits orientieren sich an Angebot und Nachfrage, z.B. für RCE-Exploits in iOS oder Android [bis zu 20 Mio. $](https://techcrunch.com/2023/10/05/zero-days-for-hacking-whatsapp-are-now-worth-millions-of-dollars/?guccounter=1) (RCE = Remote Code Execution)
    - Angriffe auf der Grundlage von Exploits haben meist eines von zwei Zielen: Ausführung von Schadcode oder Ausschleusung von Informationen (z.B. Passwörter)

- Schadcode: der große Nachteil von Computern ist auch ihr Vorteil; sie tun genau all das, was man ihnen sagt :)
    - [Computervirus](https://de.wikipedia.org/w/index.php?title=Computervirus&oldid=237356220): ein sich selbst verbreitendes Programm, dass sich in einen Computer einnistet (biologisches Analog: [Pilze, die die Gehirne von Insekten infizieren](https://www.youtube.com/watch?v=NdaYRSW76Mg))
    - [Computerwurm](https://de.wikipedia.org/w/index.php?title=Computerwurm&oldid=236407314): wie ein Virus, aber nistet sich nicht ein, sondern verbreitet sich nur weiter (z.B. über E-Mails oder USB-Sticks)
    - [Trojaner](https://de.wikipedia.org/w/index.php?title=Trojanisches_Pferd_(Computerprogramm)&oldid=233931219): ein Virus, der sich nicht mittels Exploit einer Sicherheitslücke verbreitet, sondern sich als nützliches Programm tarnt
        - in der netzpolitischen Debatte vor allem relevant in Form von [Staatstrojanern](https://de.wikipedia.org/w/index.php?title=Staatstrojaner&oldid=236566592)

- Vorschau: zweiter Teil in STP051 (Ablauf eines IT-Angriffs)

- im Gespräch erwähnt: [Vortrag von HonkHase zu Hackback und Digitalen Waffen](https://media.ccc.de/v/2019-218-wie-hackback-mit-der-gesellschaft-spielt)
