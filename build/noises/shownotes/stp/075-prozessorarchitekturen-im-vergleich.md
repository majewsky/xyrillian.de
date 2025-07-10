- Rückbezug STP011: grundsätzliche Beschreibung von [Maschinensprache](https://de.wikipedia.org/w/index.php?title=Maschinensprache&oldid=244058590)
    - siehe Beispiel unter dem obenstehenden Link
    - Maschinensprache ist spezifisch für die Prozessorarchitektur, genauer die [Befehlssatzarchitektur](https://de.wikipedia.org/w/index.php?title=Befehlssatzarchitektur&oldid=252559676) (_instruction set architecture_, ISA)

- heutzutage relevante ISA
    - x86: PCs, Notebooks, Server, [Supercomputer](https://en.wikipedia.org/wiki/File:Processor_families_in_TOP500_supercomputers.svg), Apple Macs bis 2020
    - ARM: Smartphones, Tablets, Apple Macs seit 2020, [Einplatinencomputer (SBC)](https://de.wikipedia.org/w/index.php?title=Einplatinencomputer&oldid=237691993) wie [Raspberry Pi](https://de.wikipedia.org/w/index.php?title=Raspberry_Pi&oldid=252946989)
    - Embedded: AVR, MIPS, etc.
    - Hochleistungsrechnen: neben x86 noch Power und SPARC
    - Neuentwicklungen: [RISC-V](https://de.wikipedia.org/w/index.php?title=RISC-V&oldid=253863999), [LoongArch](https://en.wikipedia.org/w/index.php?title=Loongson&oldid=1277543901#LoongArch)

- Warum ist die Wahl der ISA wichtig?
    - ISA beschreibt nominal, wie ein CPU die Bits in einer ausführbaren Binärdatei (sprich: die kodierte Form der Maschinensprache) zu interpretieren hat
    - daraus als Konsequenz Festlegung von Hardware-Parametern wie der Adressbreite oder der Anzahl und Größe der verfügbaren Register
    - Designentscheidungen in der ISA wirken sich auf die maximal erreichbare Performance und auf die Energieeffizienz aus
    - klassischerweise Gegenüberstellung zweier verschiedener Design-Philosophien: [Complex Instruction Set Computing (CISC)](https://de.wikipedia.org/w/index.php?title=Complex_Instruction_Set_Computer&oldid=252402363) und [Reduced Instruction Set Computing (RISC)](https://de.wikipedia.org/w/index.php?title=Reduced_Instruction_Set_Computer&oldid=250909684) (die Grenzen verschwimmen aber in der Praxis)
    - historischerweise war CISC die Domäne großer CPU-Hersteller, die sich ein kompliziertes Design leisten konnten, um dem Programmierer Arbeit abzunehmen (z.B. Intel); und RISC die Domäne kleiner CPU-Hersteller, die aus möglichst wenig Arbeitseinsatz möglichst viel erreichen wollten

- Was steht in einem Maschinencode-Befehl drin?
    - grundsätzlich eine Operation, die ausgeführt wird (z.B. "Addition zweier 64-Bit-Zahlen")
    - als Quellen für Argumente und Ziel für das Ergebnis: entweder Register, Speicher oder Immediate-Argumente
        - Beispiel CISC: [20 verschiedene Ausführungen des ADD-Befehls in x86-64](https://www.felixcloutier.com/x86/add)
        - Beispiel RISC: [2 verschiedene Ausführungen des ADD- bzw. ADDI-Befehls in RISC-V (RV32I)](https://projectf.io/posts/riscv-arithmetic/#addition)
    - wenn RISC wie bei RISC-V eine [Load/Store-Architektur](https://de.wikipedia.org/w/index.php?title=Load/Store-Architektur&oldid=244001127) verwendet, muss man also im Zweifelsfall Argumente erst aus dem Speicher in Register laden
        - dadurch tendenziell mehr Maschinencode, aber weniger Komplexität in der CPU

- Abwägungen zwischen Energieeffizienz (ARM) und Performance (x86)
    - aktueller ARM-Befehlssatz (A64) hat durchgängig 4 Byte große Befehle, was das Dekodieren vereinfacht; x86-Befehle haben eine variable Breite zwischen 1 und 15 Byte
    - x86 hat eine lange Pipeline, um u.a. die komplexen Befehle in kleinere Teilschritte (Micro-Operations) zu zerlegen oder bestimmte häufige Befehlssequenzen in effizientere Gruppen (Macro-Operations) zu fusionieren und damit die Performance zu steigern; ARM-Prozessoren machen das nicht und sparen dadurch Energieaufwand
    - x86 hat für bestimmte Rollen designierte Spezialregister, die je nach Befehl feste Bedeutungen haben; ARM (und RISC allgemein) verwendet für fast alles 31 General-Purpose-Register
