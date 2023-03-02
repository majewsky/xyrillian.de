- Einordnung in das Sprachenuniversum
    - Programmiersprachen sind formale Sprachen (d.h. ein Algorithmus kann in endlicher Zeit eindeutig entscheiden, ob ein gegebener Text Teil der Sprache ist) -> Rückbezug zu STP021
    - Programmiersprachen sind kein [Maschinencode](https://de.wikipedia.org/w/index.php?title=Maschinensprache&oldid=223307867 ) (d.h. Programme in Bytefolgen, den ein Prozessor direkt ausführen kann) -> Rückbezug zu STP011
    - Abwägung zwischen Lesbarkeit für den Menschen und Lesbarkeit für den Computer
    - einfachste Form von Programmiersprachen: Maschinensprachen (stellen Maschinencode direkt lesbar dar)

- Beispiel anhand einer frühen Anwendung von Programmierbarkeit: Weben/Stricken/Häkeln
    - Maschinensprache: so etwas wie "zwei links, zwei rechts, zwei fallen lassen"
    - Maschinencode = Lochkarten-Kette für einen [Jacquard-Webstuhl](https://de.wikipedia.org/w/index.php?title=Jacquardwebstuhl&oldid=226034174 )

- darstellendes Beispiel für moderne Maschinensprachen und Programmiersprachen: [Fibonacci-Funktion](https://de.wikipedia.org/w/index.php?title=Fibonacci-Folge&oldid=227809466 )
    - <a href="https://godbolt.org/#g:!((g:!((g:!((h:codeEditor,i:(filename:'1',fontScale:14,fontUsePx:'0',j:1,lang:c%2B%2B,selection:(endColumn:2,endLineNumber:6,positionColumn:2,positionLineNumber:6,selectionStartColumn:2,selectionStartLineNumber:6,startColumn:2,startLineNumber:6),source:'unsigned+int+fibonacci(unsigned+int+num)+%7B%0A++++if+(num+%3D%3D+0+%7C%7C+num+%3D%3D+1)+%7B%0A++++++++return+num%3B%0A++++%7D%0A++++return+fibonacci(num)+%2B+fibonacci(num+%2B+1)%3B%0A%7D'),l:'5',n:'0',o:'C%2B%2B+source+%231',t:'0')),k:50,l:'4',n:'0',o:'',s:0,t:'0'),(g:!((h:compiler,i:(compiler:g122,deviceViewOpen:'1',filters:(b:'0',binary:'1',commentOnly:'0',demangle:'0',directives:'0',execute:'1',intel:'0',libraryCode:'0',trim:'1'),flagsViewOpen:'1',fontScale:14,fontUsePx:'0',j:1,lang:c%2B%2B,libs:!(),options:'',selection:(endColumn:12,endLineNumber:27,positionColumn:12,positionLineNumber:27,selectionStartColumn:12,selectionStartLineNumber:27,startColumn:12,startLineNumber:27),source:1),l:'5',n:'0',o:'+x86-64+gcc+12.2+(Editor+%231)',t:'0')),k:50,l:'4',n:'0',o:'',s:0,t:'0')),l:'2',n:'0',o:'',t:'0')),version:4">interaktiv bei Godbolt</a>
    - reine Textform nachfolgend

```cpp
unsigned int fibonacci(unsigned int x) {
    if (x == 0 || x == 1) {
        return x;
    }
    return fibonacci(x - 1) + fibonacci(x - 2);
}
```

```x86asm
fibonacci(unsigned int):
    push    rbp
    mov     rbp, rsp
    push    rbx
    sub     rsp, 24
    mov     DWORD PTR [rbp-20], edi
    cmp     DWORD PTR [rbp-20], 0
    je      .L2
    cmp     DWORD PTR [rbp-20], 1
    jne     .L3
.L2:
    mov     eax, DWORD PTR [rbp-20]
    jmp     .L4
.L3:
    mov     eax, DWORD PTR [rbp-20]
    sub     eax, 1
    mov     edi, eax
    call    fibonacci(unsigned int)
    mov     ebx, eax
    mov     eax, DWORD PTR [rbp-20]
    sub     eax, 2
    mov     edi, eax
    call    fibonacci(unsigned int)
    add     eax, ebx
.L4:
    mov     rbx, QWORD PTR [rbp-8]
    leave
    ret
```

- Was unterscheidet hier die Programmiersprache von der Maschinensprache?
    - Verwendung von benannten Variablen, um Werte zu halten
        - Maschinensprache: nur Speicheradressen und CPU-Register
    - Variablen können feste Datentypen haben
        - Maschinensprache: alles nur Nullen und Einsen :)
    - Kontrollfluss-Strukturen (Subprozeduren, Abzweigungen, Schleifen)
        - Maschinensprache: nur lineare Ausführung und Sprungbefehle
    - allgemein:
        - bessere Lesbarkeit (immerhin kann Maschinensprache heutzutage auch Freitext-Kommentare)
        - Portabilität (Unabhängigkeit von Details einer bestimmten Maschinen-Architektur)
        - formale Strukturen, die Fehler vermeiden (z.B. Typprüfungen, Kollisionsvermeidung bei nebenläufigen Programmen)

- Arten von Programmiersprachen (["Paradigma"](https://de.wikipedia.org/w/index.php?title=Programmierparadigma&oldid=226550386 ))
    - imperative Programmierung: Programme als Folgen von Befehlen
    - strukturierte Programmierung: Programme als ineinander verschachtelte Blöcke von Befehlsfolgen
    - modulare Programmierung: Programme als Verbund logischer Einheiten aus Prozeduren und Datenstrukturen
    - objektorientierte Programmierung: Programme als Verbund von datenhaltenden Objekten, die miteinander interagieren
    - logische Programmierung: Programme als Verbund von logischen Aussagen und Ableitungsregeln ([Beispiele in der Sprache Prolog](https://de.wikipedia.org/w/index.php?title=Prolog_(Programmiersprache)&oldid=220556833 ); Xyrill hat eine Anekdote zu Prolog)
    - funktionale Programmierung: Programme als Verbund von mathematischen Funktionen (insbesondere von Funktionen, die andere Funktionen als Werte nehmen)
        - Beispiel unten: Berechnung sämtlicher Primzahlen in der Sprache [Haskell](https://de.wikipedia.org/w/index.php?title=Haskell_(Programmiersprache)&oldid=227783196 )
    - reaktive Programmierung: Programme als voneinander abhängige Berechnungen (Beispiel: Tabellenkalkulation, Webanwendungen mit [React](https://reactjs.org/ ))

```haskell
module Primes where

isPrime :: Int -> Bool
isPrime x = all (\y -> mod x y /= 0) $ takeWhile (\y -> y * y <= x) primes

primes :: [Int]
primes = 2 : filter isPrime [3..]
```

- Exkurs: esoterische Programmiersprachen
    - [Brainfuck](https://de.wikipedia.org/w/index.php?title=Brainfuck&oldid=226945132 ): eine der minimalsten turing-vollständigen Programmiersprache; nur 8 verschiedene Befehle
    - [Befunge](https://de.wikipedia.org/w/index.php?title=Befunge&oldid=225691979 ): anschauliches Beispiel für das Konzept einer "virtuellen Maschine" im Kontext von Programmiersprachen
    - [Chef](https://de.wikipedia.org/w/index.php?title=Chef_(Programmiersprache)&oldid=227112450 ): viele esoterische Programmiersprachen sind vor allem clevere Textersetzung

- Im Gespräch erwähnt
    - [Pentaradio 2022-12 zu Komplexität](https://www.c3d2.de/news/pentaradio24-20221227.html)
