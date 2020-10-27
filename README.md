## NPP

`npp` is a tool that designed to help musicians.

### Usage

```
$ npp --help
Command line tool for N++ music library.

Usage:
  npp [command]

Available Commands:
  get         Get music elements such as chord triads and scales
  help        Help about any command
  list        List all music elements such as chord triads and scales

Flags:
  -h, --help   help for npp
```

### Get
```
$ npp get chord --root C --type maj
[C E G]
$ npp get scale --root C --type maj
[C D E F G A B]
```

### List
```
$ npp list chord
     A#maj: [A# D F]
     A#min: [A# C# F]
      Amaj: [A C# E]
      Amin: [A C E]
      Bmaj: [B D# F#]
      Bmin: [B D F#]
     C#maj: [C# F G#]
     C#min: [C# E G#]
      Cmaj: [C E G]
      Cmin: [C D# G]
     D#maj: [D# G A#]
     D#min: [D# F# A#]
      Dmaj: [D F# A]
      Dmin: [D F A]
      Emaj: [E G# B]
      Emin: [E G B]
     F#maj: [F# A# C#]
     F#min: [F# A C#]
      Fmaj: [F A C]
      Fmin: [F G# C]
     G#maj: [G# C D#]
     G#min: [G# B D#]
      Gmaj: [G B D]
      Gmin: [G A# D]
```
