- [Proxy](https://de.wikipedia.org/w/index.php?title=Proxy_(Rechnernetz)&oldid=267567396): ein Vermittler zwischen Client und Server
    - von engl. _proxy representative_ (Stellvertreter), [Kognat](https://de.wikipedia.org/w/index.php?title=Kognat&oldid=264251978) zur Prokurist:in via lat. _procurator_ (Stellvertreter, Besorger, Verwalter)
    - vgl. Kellner als Proxy zwischen den Gästen und der Küche
    - vgl. Callcenter-Mitarbeiter als Proxy zwischen dem Kunden und der Firma
    - im Netzwerksinne ein Vermittler nicht auf Netzwerkebene (wie ein Router auf Layer 3, vgl. STP005), sondern auf Applikationsebene

- im engeren Sinne verwaltet ein Proxy Anfragen eines Clients an mehrere Server, aber heute vielleicht am weitesten verbreitet: [Reverse Proxy](https://de.wikipedia.org/w/index.php?title=Reverse_Proxy&oldid=261820436)
    - steht vor einem Server und stellt die Schnittstelle nach außen dar
    - Anwendungsfälle: Lastverteilung, Caching, Authentifizierung/Authorisierung (siehe STP039), DDoS-Schutz, TLS-Terminierung (siehe STP052) bzw. Vermittlung zwischen TLS nach außen und mTLS nach innen (Wikipedia ist nostalgisch und redet hier noch von "Beschleunigungshardware" für SSL)

- hingegen der vom Client bestellte Vermittler: Anwendungsfälle für Forward Proxies (insb. im Kontext von HTTP)
    - Sicherheitsgrenze in Firmen- oder anderen geschützten Netzwerken: heute kaum noch verbreitet aufgrund der weiten Verbreitung von TLS (stattdessen eher DNS-Level-Filterung, oder halt TLS-MitM-Proxying mit In-House-CA; evtl. Einsatz auf Servern, die nur mit wohldefinierten Zielen reden dürfen, aber nicht mit z.B. C&C-Servern von Botnets)
    - Umgehung von Zugangsbeschränkungen (Ratelimits oder Geoblocking): heute eher mit VPN (siehe STP098)
    - Anonymität: heute eher mit Tor (siehe STP098)
    - Verbindungspersistenz: früher bei IRC relevant, heutzutage (siehe STP022) ist Chat einfach nicht mehr dezentral (bestenfalls noch föderiert)
