# Command-line executables

will take a png representing a cell pattern (non-black=live cell) run a number of ticks, and save result.

note: a glider, or other such, will cause the creation of a very large png file after a large number of ticks, since here the origin is always in the png file.

set 'movie' to a folder, and 'wrap' to something and get snapshot images saved every 'frameticks' cycles.


|  sys/arch     |   file suffix      |           details                                                                         |    notes       |
|---------------|--------------------|-------------------------------------------------------------------------------------------|----------------|
| linux/amd64   | [SYSV64].elf       | ELF 64-bit LSB executable: x86-64: version 1 (SYSV): statically linked: not stripped      |                |
| linux/386     | [SYSV32].elf       | ELF 32-bit LSB executable: Intel 80386: version 1 (SYSV): statically linked: not stripped |                |
| linux/arm64   | [arm64A].elf       | ELF 64-bit LSB executable: ARM aarch64: version 1 (SYSV): statically linked: not stripped |   Cortex A     |
| linux/arm32   | [arm32v5].elf      | ELF 32-bit LSB executable: ARM: EABI5 version 1 (SYSV): statically linked: not stripped   |   no HW F-P    |
| linux/arm32   | [arm32v6].elf      | ELF 32-bit LSB executable: ARM: EABI5 version 1 (SYSV): statically linked: not stripped   |   		      |
| linux/arm32   | [arm32v7].elf      | ELF 32-bit LSB executable: ARM: EABI5 version 1 (SYSV): statically linked: not stripped   |  	          |
| windows/amd64 | [PE32+].exe        | PE32+ executable (console) x86-64 (stripped to external PDB): for MS Windows              |                |
| windows/386   | [PE32].exe         | PE32 executable (console) Intel 80386 (stripped to external PDB): for MS Windows          |                |
| darwin/amd64  | [macho64]          | Mach-O 64-bit x86_64 executable                                                           |                |
| darwin/386    | [macho32]          | Mach-O i386 executable                                                                    |                |



Usage
```
   -f uint
    	ticks for each snapshot image. (default 1)
  -frameTicks uint
    	ticks for each snapshot image. (default 1)
  -h	display help/usage.
  -help
    	display help/usage.
  -i value
    	source for the starting cell pattern, encoded in PNG image.(default:<Stdin>)
  -input value
    	source for the starting cell pattern, encoded in PNG image.(default:<Stdin>)
  -interval duration
    	time between log status reports (default 1s)
  -log value
    	progress log destination.(default:Stderr)
  -m value
    	directory for snapshot frames, PNG images.
  -movie value
    	directory for snapshot frames, PNG images.
  -o value
    	file for encoding result cell pattern, PNG image.(default:Stdout)
  -output value
    	file for encoding result cell pattern, PNG image.(default:Stdout)
  -p	send snapshot images to Stdout.
  -pipeMovie
    	send snapshot images to Stdout.
  -s uint
    	size of snapshots. (default 32)
  -size uint
    	size of snapshots. (default 32)
  -ticks uint
    	Ticks/Cycles (default 1)
  -w	sets arena to (s)ize.
  -wrap
    	sets arena to (s)ize.
```

