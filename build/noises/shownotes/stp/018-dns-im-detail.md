Nicht die Desoxyribonukleinsäure, sondern das [Domain Name System](https://de.wikipedia.org/wiki/Domain_Name_System ).

- Rückbezüge
    - zu STP002 (Aufruf einer Webseite): Webserver haben Domain-Namen, die aufgelöst werden müssen
    - zu STP012 (Datenbanken): DNS ist eine globale, verteilte (aber nicht dezentrale!), hierarchische Datenbank

- Grundaufgabe: Namensauflösung
    - menschenlesbare und stabile Server-Namen anstatt unlesbarer und möglicherweise wechselnder IP-Adressen
    - aber auch andere Abfragen möglich, siehe unten
    - historischer Vorgänger: `/etc/hosts`

- Aufteilung des Internet in **Domains** (Domänen) und verschachtelte **Subdomains** (Unterdomänen)
    - [Beispiel:](https://commons.wikimedia.org/wiki/File:Dns-raum.svg ) Wurzeldomain "." enthält Domain "org" enthält Domain "wikipedia.org" enthält Domain "de.wikipedia.org" -> deswegen _hierarchische_ Datenbank
    - wenn die Subdomain einer anderen Person oder Organisation gehört als die Domain darüber, hat man eine neue **Zone**
    - jeder Inhaber einer Zone hält auf einem (oder mehreren) Webserver(n) die Inhalte seiner Zone (also Domain-Daten und Zonenreferenzen) in einem **Nameserver** -> deswegen _verteilte_ Datenbank

- Wurzelzone ist unter der Kontrolle der [ICANN](https://de.wikipedia.org/wiki/Internet_Corporation_for_Assigned_Names_and_Numbers ) -> deswegen _nicht dezentrale_ Datenbank
    - Verteilung der Wurzelzone durch die [Root-Nameserver](https://de.wikipedia.org/wiki/Root-Nameserver )
    - finanzieller Anreiz zum Erlauben von immer mehr Top-Level-Domains (TLD)

- Einträge in einer DNS-Zone heißen **Records** (Datensätze), mehrere Record Types möglich:
    - A, AAAA: IP-Adresse zu einer Domain (A = IPv4, AAAA = IPv6)
    - CNAME: Domain-Alias (anstatt direkt eine IP zu erhalten, erhält man eine andere Domain, die man stattdessen nach der finalen IP fragen muss)
    - PTR: Rückwärtsabfrage, Domain zu einer IP-Adresse
    - MX: Mail-Server zu einer Domain
    - SRV: ähnlich wie MX, aber für beliebige Dienste (z.B. XMPP)
    - TXT: beliebige Textinformationen (wird zum Beispiel für Mail-Empfangsregeln verwendet)
    - NS: Zonendelegation, die entsprechende Domain ist der Ausgangspunkt einer Unterzone und der Record enthält den autoritativen Nameserver für diese Zone

- Abfrage einer bestimmten Domain
    - Grundidee: zuerst Root-Nameserver fragen, der delegiert an den Nameserver der TLD, der delegiert an den Nameserver der Second-Level-Domain, etc.
    - Problem: absurd hohe Last auf die Root-Nameserver
    - Lösung 1: jeder Eintrag in einer Nameserver-Datenbank hat eine TTL ("Time to Live")
    - Lösung 2: Endnutzer reden nicht direkt mit den autoritativen Nameservern, sondern mit Vermittlern (**Resolver**), die sie sich mit anderen Endnutzern teilen

- Komplexbeispiel: Abfrage von `de.wikipedia.org` über die Root-Nameserver mit dem Unix-Tool `dig` (Ausgaben stark gekürzt)

```
$ dig A de.wikipedia.org @m.root-servers.net

;; QUESTION SECTION:
;de.wikipedia.org.              IN      A
;; AUTHORITY SECTION:
org.                    172800  IN      NS      a0.org.afilias-nst.info.
...
;; ADDITIONAL SECTION:
a0.org.afilias-nst.info. 172800 IN      A       199.19.56.1
...

$ dig A de.wikipedia.org @199.19.56.1

;; QUESTION SECTION:
;de.wikipedia.org.              IN      A
;; AUTHORITY SECTION:
wikipedia.org.          86400   IN      NS      ns0.wikimedia.org.
...
;; ADDITIONAL SECTION:
ns0.wikimedia.org.      86400   IN      A       208.80.154.238
...

$ dig A de.wikipedia.org @208.80.154.238

;; QUESTION SECTION:
;de.wikipedia.org.              IN      A
;; ANSWER SECTION:
de.wikipedia.org.       86400   IN      CNAME   dyna.wikimedia.org.

$ dig A dyna.wikimedia.org @199.19.56.1

;; QUESTION SECTION:
;dyna.wikimedia.org.            IN      A
;; AUTHORITY SECTION:
wikimedia.org.          86400   IN      NS      ns0.wikimedia.org.
...
;; ADDITIONAL SECTION:
ns0.wikimedia.org.      86400   IN      A       208.80.154.238
...

$ dig A dyna.wikimedia.org @208.80.154.238

;; QUESTION SECTION:
;dyna.wikimedia.org.            IN      A
;; ANSWER SECTION:
dyna.wikimedia.org.     600     IN      A       91.198.174.192
```

- zum Vergleich: Abfrage von `de.wikipedia.org` über einen öffentlichen Resolver (beachte die unterschiedlichen TTL-Werte)

```
$ dig A de.wikipedia.org @9.9.9.9

;; QUESTION SECTION:
;de.wikipedia.org.              IN      A
;; ANSWER SECTION:
de.wikipedia.org.       32533   IN      CNAME   dyna.wikimedia.org.
dyna.wikimedia.org.     227     IN      A       91.198.174.192
```

- Beispiel: Rückwärtsabfrage der so erhaltenen IP mittels PTR-Record unter der Pseudo-Domain `in-addr.arpa`

```
dig PTR 192.174.198.91.in-addr.arpa
;; QUESTION SECTION:
;192.174.198.91.in-addr.arpa.   IN      PTR

;; ANSWER SECTION:
192.174.198.91.in-addr.arpa. 1104 IN    PTR     text-lb.esams.wikimedia.org.
```

- Woher kommt der Resolver?
    - normalerweise über DHCP (auf demselben Wege, über den das eigene Gerät eine IP-Adresse vom Router bekommt; oder über den der Router eine IP-Adresse vom ISP bekommt), unter Unix siehe `/etc/resolv.conf`
    - somit meist ein Resolver unter Kontrolle des ISP
    - alternative Resolver: z.B. `1.1.1.1` (Cloudflare), `8.8.8.8` (Google), `9.9.9.{9,10,11}` ([Quad9](https://de.wikipedia.org/wiki/Quad9 )), `5.9.164.112` ([Digitalcourage](https://digitalcourage.de/support/zensurfreier-dns-server ))
    - Resolverwahl: Implikationen für Privatsphäre und für Vertrauenswürdigkeit

- DNS selbst ist unverschlüsselt -> neuere Entwicklung: DNS-over-TLS und DNS-over-HTTP
    - eigentlich würde DNS-over-TLS reichen, aber kann dann vom ISP geblockt werden; HTTP hingegen ist aus praktischen Gründen quasi unblockbar
    - Unterstützung in Betriebssystemen noch lückenhaft, aber Browser können DoT und DoH am Betriebssystem vorbei machen (Browsereinstellungen prüfen!)

- alternative DNS-Roots: meist in privaten Netzen (z.B. Firmen-Intranet mit `.corp`-Domains), aber es gibt auch alternative Roots mit globalem Anspruch (z.B. [Namecoin](https://en.wikipedia.org/wiki/Namecoin ))
