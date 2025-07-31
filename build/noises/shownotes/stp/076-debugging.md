[Wikipedia definiert den Begriff](https://de.wikipedia.org/w/index.php?title=Debuggen&oldid=254167673):

> Als Debuggen (dt. Entwanzen) oder Fehlerbehebung bezeichnet man in der Informatik den Vorgang, in einem Computerprogramm Fehler oder unerwartetes Verhalten zu diagnostizieren und zu beheben. Die Suche von Programmfehlern (sogenannten Bugs) ist eine der wichtigsten und anspruchsvollsten Aufgaben der Softwareentwicklung und nimmt einen großen Teil der Entwicklungszeit in Anspruch. Ein guter Programmierer muss daher auch das Debuggen beherrschen und umgekehrt: Nicht jeder gute Programmierer ist auch ein guter Debugger.

- Xyrill hat eine Anekdote: Firefox im Hotel-WLAN hängt beim Starten vier Minuten
    - nicht dramatisch, könnte man auch hinnehmen (oder halt auf einen anderen Browser wechseln)
    - Xyrill war aber in der Stimmung, zu debuggen

- Kategorie 1: Werkzeuge zum Mitlauschen
    - [traceroute](https://de.wikipedia.org/w/index.php?title=Traceroute&oldid=236437154) (bzw. `tracert` unter Windows): Nachverfolgung der Netzwerkroute zu einem bestimmten Ziel anhand von Ping-Paketen mit begrenzter TTL; stark variable Pings oder plötzliche Anstiege des Pings deuten auf ein Problem an dem entsprechenden Hop hin
    - [tcpdump](https://de.wikipedia.org/w/index.php?title=Tcpdump&oldid=235025527) und [Wireshark](https://de.wikipedia.org/w/index.php?title=Wireshark&oldid=238943114): Mitschneiden aller Netzwerkpakete, die über ein Netzwerk-Interface versendet oder empfangen werden (evtl. mit Suchmuster, z.B. "nur auf Port 80" oder "nur für Ziel-/Quelladresse 192.168.0.1")
    - [mitmproxy](https://mitmproxy.org/) et al: Fangnetz für HTTP-Anfragen und -Antworten, die von einem Prozess geschickt und empfangen werden (quasi wie tcpdump, aber einige Netzwerkschichten höher); super zum Reverse-Engineering von webbasierten APIs
    - mal nicht Netzwerk: [strace](https://en.wikipedia.org/w/index.php?title=Strace&oldid=1268124388) listet alle Syscalls (Systemrufe) auf, die ein Prozess während seiner Lebenszeit macht (siehe Beispiel unten)
    - zurück zum Hotel-WLAN-Beispiel: strace zeigt Timeout bei DNS-Anfrage zum eigenen Hostnamen jeweils 2x nach 120s
    - Grund für diese Anfragen bleibt unklar, aber die Lösung ist eindeutig: expliziter Eintrag für den eigenen Hostnamen in `/etc/hosts`, um die DNS-Abfrage zu umgehen

- Kategorie 2: Werkzeuge zum Nachverfolgen
    - [printf](https://de.wikipedia.org/w/index.php?title=Printf&oldid=243456170)-Debugging: im Programmcode vorrübergehend Befehle einbauen, die zu sinnfälligen Zeitpunkten Logmeldungen ausdrucken, insb. beinhaltend den aktuellen Zustand von bestimmten Variablen (sprich: transiente Form von Logging, siehe STP032)
    - Debugger wie [gdb](https://de.wikipedia.org/w/index.php?title=GNU_Debugger&oldid=237899188): Ausführung eines Programms bis zum Moment des Crashs oder bis zu benutzerdefinierten Haltepunkten; dann schrittweise Ausführung bzw. Inspektion von Variablen (flexibler als printf-Debugging, aber erfordert direkten Zugriff auf den Prozess)
    - Profiler wie [valgrind](https://de.wikipedia.org/w/index.php?title=Valgrind&oldid=246184662): Ausführung eines Programms und währenddessen Erfassung von Statistiken zu Laufzeit- und Speichernutzung, um Speicherzugriffsfehler oder Ineffizienzen aufzudecken

- Kategorie 3: Werkzeuge zum Zurückverfolgen
    - [Bisektion](https://de.wikipedia.org/w/index.php?title=Bisektion&oldid=194714676) z.B. mit [git bisect](https://git-scm.com/docs/git-bisect): Aufspüren derjenigen Änderung in einem Versionsverwaltungssystem, die eine bestimmte Verhaltensänderung (z.B. einen Bug oder einen Performance-Einbruch) ausgelöst hat

---

Beispiel einer strace-Ausgabe für den Befehl `mv /tmp/foo /tmp/bar` (stark verkürzt):

```
$ strace mv /tmp/foo /tmp/bar
execve("/usr/bin/mv", ["mv", "/tmp/foo", "/tmp/bar"], 0x7fff820357b0 /* 67 vars */) = 0
... (ca. 60 Zeilen an Setup-Phase der libc und anderer Shared Libraries ausgelassen) ...
ioctl(0, TCGETS, {c_iflag=ICRNL|IXON|IUTF8, c_oflag=NL0|CR0|TAB0|BS0|VT0|FF0|OPOST|ONLCR, c_cflag=B38400|CS8|CREAD, c_lflag=ISIG|ICANON|ECHO|ECHOE|ECHOK|IEXTEN|ECHOCTL|ECHOKE, ...}) = 0
renameat2(AT_FDCWD, "/tmp/foo", AT_FDCWD, "/tmp/bar", RENAME_NOREPLACE) = -1 EEXIST (File exists)
openat(AT_FDCWD, "/tmp/bar", O_RDONLY|O_PATH|O_DIRECTORY) = -1 ENOTDIR (Not a directory)
newfstatat(AT_FDCWD, "/tmp/foo", {st_mode=S_IFREG|0644, st_size=0, ...}, AT_SYMLINK_NOFOLLOW) = 0
newfstatat(AT_FDCWD, "/tmp/bar", {st_mode=S_IFREG|0644, st_size=0, ...}, AT_SYMLINK_NOFOLLOW) = 0
geteuid()                               = 1001
faccessat2(AT_FDCWD, "/tmp/bar", W_OK, AT_EACCESS) = 0
renameat2(AT_FDCWD, "/tmp/foo", AT_FDCWD, "/tmp/bar", 0) = 0
lseek(0, 0, SEEK_CUR)                   = -1 ESPIPE (Illegal seek)
close(0)                                = 0
close(1)                                = 0
close(2)                                = 0
exit_group(0)                           = ?
+++ exited with 0 +++
```
