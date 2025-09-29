- Rückblick: wir hatten schon einige Folgen zu Speicherverwaltung
    - STP007: Speicherhierarchie (Festplatte und SSD, Arbeitsspeicher/RAM, Prozessor-Cache, Registerbank)
    - STP019: Speicherschutz (virtueller Speicher, Auslagerungsspeicher/Swap, Direct Memory Access)
    - Nachtrag/Linktipp: <https://www.linuxatemyram.com>
    - das war alles aus Sicht des Betriebssystems; diesmal Abläufe innerhalb des Prozesses

- Was stellt uns das Betriebssystem bereit? -> [Virtueller Speicher](https://de.wikipedia.org/w/index.php?title=Virtuelle_Speicherverwaltung&oldid=233703410 )
    - nominale Größe: entsprechend der Adressbreite des Prozessors (für 32 Bit sind das 2^32 Bytes = 4 GiB, für 64 Bit sind es 2^64 Bytes = 16 EiB = 16.777.216 TiB)
    - nutzbar aber nur dort, wo das Betriebssystem für uns eine tatsächlich existierende Speicherseite einblendet
    - initiale Belegung (siehe Beispiel unten): hauptsächlich Programmdatei und benötigte Programmbibliotheken mit ihren entsprechenden Segmenten
    - Belegung (und Rückgabe) weiterer Speichersegmente mittels Syscalls (siehe STP019)
    - klassische Struktur des virtuellen Speichers: oben der Stapel, unten ein Haufen

- vorher einige Bemerkungen zu klassischen Schutzmaßnahmen, die man im untenstehenden Beispiel sehen kann
    - die untersten paar KiB des Adressbereiches werden nie vergeben
        - dies verhindert versehentlichen Zugriff auf Nullzeiger mit hoher Wahrscheinlichkeit
    - [Address Space Layout Randomization (ASLR)](https://de.wikipedia.org/w/index.php?title=Address_Space_Layout_Randomization&oldid=227389251 ): Programmdateien und Programmbibliotheken werden an zufällig gewählten Speicheradressen eingeblendet, um es Angreifern zu erschweren, das Programm so zu manipulieren, dass es ungeplant in deren Code hineinspringt
        - klingt nach einer einfach umsetzbaren Schutzmaßnahme
        - war aber lange nicht weit verbreitet, weil der benötigte [Position Independent Code (PIC)](https://en.wikipedia.org/w/index.php?title=Position-independent_code&oldid=1164854994 ) mit einem kleinen Geschwindigkeitsverlust (1-3%) einhergeht

- oberer Teil (d.h. bei hohen Adressen): **Stack** (Stapelspeicher)
    - Programme sind meist als ineinander verschachtelte Unterprogramme (Subroutinen) aufgebaut (siehe STP011)
    - Daten, die nur für die Dauer des jeweiligen Unterprogramms gebraucht werden, können im Stack aufgestapelt werden
    - Idee: Hauptroutine kriegt den ganzen Stack und knapst sich oben soviel ab, wie es selber braucht
    - nächste Unterroutine kriegt den Rest des Stacks und verfährt in derselben Weise
    - wenn die Unterroutine zurückkehrt, kann der von ihr benutzte Stack an die nächste Unterroutine weitergereicht werden -> keine aufwändige Speicherverwaltung nötig
    - bei Prozessen mit mehreren Threads (parallel lauffähigen Teilprozessen) teilen sich alle Threads grundsätzlich denselben virtuellen Speicher, aber kriegen jeweils eigene Stacks
    - Stack ist vom Betriebssystem in einer festen Größe angelegt (im Beispiel unten 132 KiB) &ndash; bei zu vielen verschachtelten Subroutinen-Aufrufen kommt es zum berüchtigten **Stack Overflow**

- unterer Teil (d.h. bei tiefen Adressen): **Heap** (Haufenspeicher)
    - keine Strukturvorgaben durch das Betriebssystem oder die Prozessorarchitektur
    - notwendig für Daten, deren Lebenszeit nicht der Struktur von Unterprogrammen folgt, oder deren Größe die 132 KiB des Stacks übersteigt
    - neue Frage: Wie kann man diesen unstrukturierten Haufen sinnvoll verwalten? -> Teil 2 folgt in STP047

- Randnotiz zu unserem Beispiel unten: [Was sind vvar, vdso und vsyscall?](https://lwn.net/Articles/615809/ )
    - Syscalls sind zeitaufwändig (in der Größenordnung von dutzenden Mikrosekunden), da man vom Prozess in den Kernel umschalten muss und wieder zurück
    - Idee: manche Syscalls kann der Kernel mit geschickter Planung im Prozess selbst beantworten, ohne Kontextwechsel
    - Beispiel: `gettimeofday` (Abfrage der aktuellen Uhrzeit) kann man aus einem vorhergehenden Datenpunkt des Kernels anhand des Taktzählers des Prozessors fortschreiben (z.B. [TSC](https://en.wikipedia.org/w/index.php?title=Time_Stamp_Counter&oldid=1160257061 ) auf x86)
    - vsyscall: erste derartige Implementation, mit zwei großen Problemen (Platzbegrenzung, kein PIC = kein ASLR); nur aus Kompatibilitätsgründen angeboten
    - [vDSO](https://en.wikipedia.org/w/index.php?title=VDSO&oldid=1146593908 ) ("virtual dynamic shared object"): die aktuelle Implementation, verhält sich wie eine dynamisch nachgeladene Programmbibliothek
    - vvar: Datensegment für das vDSO

---

```
$ sleep 30 &
[1] 4970
$ cat /proc/$(pidof sleep)/maps
# address                perms offset   dev   inode       pathname
556ac29f4000-556ac29f6000 r--p 00000000 fe:00 10234368    /usr/bin/sleep
556ac29f6000-556ac29f9000 r-xp 00002000 fe:00 10234368    /usr/bin/sleep
556ac29f9000-556ac29fa000 r--p 00005000 fe:00 10234368    /usr/bin/sleep
556ac29fa000-556ac29fb000 r--p 00006000 fe:00 10234368    /usr/bin/sleep
556ac29fb000-556ac29fc000 rw-p 00007000 fe:00 10234368    /usr/bin/sleep
556ac2f81000-556ac2fa2000 rw-p 00000000 00:00 0           [heap]
7fe4dea00000-7fe4ded42000 r--p 00000000 fe:00 10321904    /usr/lib/locale/locale-archive
7fe4deefc000-7fe4deeff000 rw-p 00000000 00:00 0
7fe4deeff000-7fe4def21000 r--p 00000000 fe:00 10226583    /usr/lib/libc.so.6
7fe4def21000-7fe4df07e000 r-xp 00022000 fe:00 10226583    /usr/lib/libc.so.6
7fe4df07e000-7fe4df0d6000 r--p 0017f000 fe:00 10226583    /usr/lib/libc.so.6
7fe4df0d6000-7fe4df0da000 r--p 001d6000 fe:00 10226583    /usr/lib/libc.so.6
7fe4df0da000-7fe4df0dc000 rw-p 001da000 fe:00 10226583    /usr/lib/libc.so.6
7fe4df0dc000-7fe4df0eb000 rw-p 00000000 00:00 0
7fe4df11e000-7fe4df11f000 r--p 00000000 fe:00 10226556    /usr/lib/ld-linux-x86-64.so.2
7fe4df11f000-7fe4df145000 r-xp 00001000 fe:00 10226556    /usr/lib/ld-linux-x86-64.so.2
7fe4df145000-7fe4df14f000 r--p 00027000 fe:00 10226556    /usr/lib/ld-linux-x86-64.so.2
7fe4df14f000-7fe4df151000 r--p 00031000 fe:00 10226556    /usr/lib/ld-linux-x86-64.so.2
7fe4df151000-7fe4df153000 rw-p 00033000 fe:00 10226556    /usr/lib/ld-linux-x86-64.so.2
7ffe3da15000-7ffe3da36000 rw-p 00000000 00:00 0           [stack]
7ffe3da73000-7ffe3da77000 r--p 00000000 00:00 0           [vvar]
7ffe3da77000-7ffe3da79000 r-xp 00000000 00:00 0           [vdso]
ffffffffff600000-ffffffffff601000 --xp 00000000 00:00 0   [vsyscall]
```

Randbemerkungen: Sofern ich die Ausgabe von `strace sleep 30` richtig lese, wurde `[heap]` hier mit `brk` verwaltet. Die beiden anonymen Segmente gehen höchstwahrscheinlich auf das Konto des Laufzeit-Linkers (hier `ld-linux-x86-64.so.2`), der die Programmbibliotheken (hier nur `libc.so.6`) in den Prozess lädt.

Um das zu bestätigen, habe ich dann ~~mal schnell~~ ein minimales Programm gebaut, dass ausschließlich `sleep` gefolgt von `exit` macht, und definitiv gar nichts anderes. Ich war in der Tat erfolgreich:

```
$ strace ./target/release/minimal-memory-maps
execve("./target/release/minimal-memory-maps", ["./target/release/minimal-memory-"...], 0x7ffd0c104b00 /* 67 vars */) = 0
nanosleep({tv_sec=10, tv_nsec=0}, 0x7ffdaa084800) = 0
exit(0)                                 = ?
+++ exited with 0 +++
```

Und dieses Programm hat nicht mal einen Heap, sondern nur den Stack und die Kernel-Schnittstellen:

```
$ cat /proc/$(pidof minimal-memory-maps)/maps
# address                perms offset   dev   inode       pathname
00400000-00401000         rwxp 00000000 fe:00 9309336     .../target/release/minimal-memory-maps
7ffdaa066000-7ffdaa087000 rw-p 00000000 00:00 0           [stack]
7ffdaa0ea000-7ffdaa0ee000 r--p 00000000 00:00 0           [vvar]
7ffdaa0ee000-7ffdaa0f0000 r-xp 00000000 00:00 0           [vdso]
ffffffffff600000-ffffffffff601000 --xp 00000000 00:00 0   [vsyscall]
```

Hier sieht man übrigens schön, dass ich mein Programm nicht mit PIC kompiliert habe, sodass für die Programmdatei selber keine ASLR vorgenommen werden konnte.

Der Vollständigkeit halber dokumentiere ich hier noch den Code, um mein Testprogramm zu reproduzieren. Der Code funktioniert nur auf Linux und x86-64. Und selbst da habe ich irgendwas nicht ganz richtig gemacht: Das Programm funktioniert nur, wenn man dem Compiler Optimierungen verbietet (?!?). Trotzdem ein Shoutout an [diesen Blogartikel](https://darkcoding.net/software/a-very-small-rust-binary-indeed/) und [diesen anderen Blogartikel](https://sgibala.com/01-06-single-syscall-hello-world-part-2/ ), mit deren Hilfe ich das Kompilieren ohne libc und main zumindest hinreichend weit ans Laufen bekommen habe.

```
$ cat Makefile
build:
    RUSTFLAGS="-Copt-level=0 -Ctarget-cpu=native -Clink-args=-nostartfiles -Clink-args=-Wl,-n,-N,--no-dynamic-linker" cargo build --release

$ cat Cargo.toml
[package]
name = "minimal-memory-maps"
version = "0.1.0"
edition = "2021"
[profile.dev]
panic = "abort"
[profile.release]
panic = "abort"

$ cat src/main.rs
#![no_std]
#![no_main]
use core::arch::asm;

#[repr(C)]
struct Timespec {
    pub tv_sec: i64,
    pub tv_nsec: i64,
}

const NR_NANOSLEEP: usize = 35;
const NR_EXIT: usize = 60;

#[no_mangle]
pub extern "C" fn _start() -> ! {
    let req = Timespec { tv_sec: 10, tv_nsec: 0 };
    let mut resp = Timespec { tv_sec: 0, tv_nsec: 0 };
    unsafe {
        // nanosleep(&req, &mut resp)
        let req_ptr = (&req as *const Timespec) as usize;
        let resp_ptr = (&mut resp as *mut Timespec) as usize;
        let ret: usize;
        asm!(
            "syscall",
            inout("rax") NR_NANOSLEEP => ret,
            in("rdi") req_ptr,
            in("rsi") resp_ptr,
            out("rcx") _,
            out("r11") _,
        );
        let exit_code = if ret == 0 { 0 } else { 1 };
        // exit(0)
        asm!(
            "syscall",
            in("rax") NR_EXIT,
            in("rdi") exit_code,
            options(noreturn)
        )
    }
}

#[panic_handler]
fn my_panic(_info: &core::panic::PanicInfo) -> ! {
    loop {}
}
```
