<div align="center">

<pre>
 ┌────────────────────────────────────────────────────┐
 │  user@dev:~/projects$ ./converter                  │
 │                                                    │
 │      ██████╗   ██████╗   ███╗   ██╗  ██╗   ██╗     │
 │     ██╔════╝  ██╔═══██╗  ████╗  ██║  ██║   ██║     │
 │     ██║       ██║   ██║  ██╔██╗ ██║  ██║   ██║     │
 │     ██║       ██║   ██║  ██║╚██╗██║  ╚██╗ ██╔╝     │
 │     ╚██████╗  ╚██████╔╝  ██║ ╚████║   ╚████╔╝      │
 │      ╚═════╝   ╚═════╝   ╚═╝  ╚═══╝    ╚═══╝       │
 │                                                    │
 │  > terminal unit converter · go 1.23 stdlib        │
 └────────────────────────────────────────────────────┘
</pre>

![type](https://img.shields.io/badge/type-study%20project-2ea043?style=for-the-badge&labelColor=0d1117)
![version](https://img.shields.io/badge/version-1.0.0-58a6ff?style=for-the-badge&labelColor=0d1117)
![license](https://img.shields.io/badge/license-mit-8b949e?style=for-the-badge&labelColor=0d1117)
![stack](https://img.shields.io/badge/stack-go%201.23-d29922?style=for-the-badge&labelColor=0d1117)

</div>

---

## `$ cat about.txt`

```
converter is an interactive terminal menu that converts six kinds of
units: length, weight, temperature, time, area and volume.

the options live in a map[int]string, the choice and the value are read
from stdin with bufio and strconv, and a switch hands the number to one
exported function in conversions/, each one a single line of arithmetic.

built as a hands-on study of go: modules, packages, bufio and strconv.
```

## `$ ls features/`

```
[x] six conversions, one exported function each in conversions/
[x] meters -> kilometers, kilograms -> pounds, celsius -> fahrenheit
[x] seconds -> minutes, square meters -> square km, liters -> us gallons
[x] menu loop that keeps going until you pick 0
[x] unparsable input and unknown options are asked again, no crash
[x] colored ansi output, the value with 2 decimals, the result with 3
[x] a 370ms pause after each menu line and before each result or error
[ ] one direction only — no kilometers -> meters, no pounds -> kilograms
[ ] the menu starts at a random option on each loop (it ranges over a map)
[ ] no range check — negative values, inf and nan are converted as-is
[ ] piped input that ends without a 0 loops forever on "INVALID INPUT"
[ ] go.mod names the module github.com/miguelcrwz, not .../converter
[ ] no tests
```

## `$ cat stack.lock`

```
runtime     go 1.23.3 (go.mod)
framework   none — stdlib only (bufio, fmt, os, strconv, strings, time)
data        none, every factor is a constant in conversions/
infra       none, runs locally in a terminal
tests       none yet
tooling     go run / go build
```

## `$ ./install.sh`

```bash
# clone
git clone https://github.com/miguelcrwz/converter.git
cd converter

# dependencies (stdlib only — just needs go 1.23.3+)
go version
```

## `$ ./run`

```bash
# development (from the repo root)
go run .

# production
go build -o converter . && ./converter
```

```
WELCOME TO UNIT CONVERTER

 CHOOSE A CONVERSION:
1. METERS to KILOMETERS
2. KILOGRAMS to POUNDS
3. CELSIUS to FAHRENHEIT
4. SECONDS to MINUTES
5. SQUARE METERS to SQUARE KILOMETERS
6. LITERS to GALLONS
0. PROGRAM EXIT
 Enter your CHOICE: 3

Enter TEMPERATURE in CELSIUS: 100

100.00 CELSIUS = 212.000 FAHRENHEIT
```

## `$ cat menu.flow`

```
stdin
  │
  ├─ print the menu (map order, so it can start at any option)
  │
  └─ read the choice
       ├─ not an integer, or EOF ─────> INVALID INPUT, menu again
       ├─ 0 ──────────────────────────> GOODBYE, exit code 0
       ├─ not 1-6 ────────────────────> INVALID CHOICE, menu again
       └─ 1-6 ─> read the value
            ├─ not a float, or EOF ───> INVALID INPUT, ask again
            └─ ok ─> conversions.Xxx ─> print, menu again
```

## `$ tree`

```
converter/
├── main.go              # menu loop, input parsing and colored output
├── conversions/
│   ├── lenght.go        # meters / 1000 = kilometers
│   ├── weight.go        # kilograms * 2.20462 = pounds
│   ├── temperature.go   # celsius * 9 / 5 + 32 = fahrenheit
│   ├── time.go          # seconds / 60 = minutes
│   ├── area.go          # square meters / 1_000_000 = square km
│   └── volume.go        # liters * 0.264172 = us gallons
├── preview/
│   └── gopher.gif       # the gopher that celebrates exit 0
├── go.mod               # module github.com/miguelcrwz, go 1.23.3
├── .gitignore
├── LICENSE
└── README.md
```

## `$ git log --oneline roadmap`

```
a1b2c3d  print the menu in a fixed order (sorted keys or a slice)
e4f5g6h  stop on EOF instead of asking again forever
i7j8k9l  add the reverse of every conversion (km -> m, lb -> kg, ...)
m0n1o2p  table-driven tests for the conversions package
q3r4s5t  rename the module to github.com/miguelcrwz/converter
```

## `$ cat LICENSE | head -1`

```
MIT © 2025 Miguel Cruz
```

## `$ whoami`

```
miguel cruz
├── github    https://github.com/miguelcrwz
├── linkedin  https://linkedin.com/in/miguelcrwz
└── email     miguwlcrwz@gmail.com
```

<div align="center">

<pre>
$ exit
[process exited with code 0]
</pre>

<img src="preview/gopher.gif" alt="dancing gopher" width="128">

</div>
