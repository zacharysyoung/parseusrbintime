# Parse /usr/bin/time -l

Parse the output of '/usr/bin/time -l'; includes a simple command to output a parsed record as CSV (with a header).

To install, clone/download the repo and build and install [parseusrbintime/main.go](parseusrbintime/main.go) with `go install`.

Time a command (it's output must be redirected away from the command):

```none
{ /usr/bin/time -l wc /usr/share/dict/words >/dev/null ; } 2>&1 | parseusrbintime
Real(s),User(s),Sys(s),Memory(B)
0.03,0.01,0.01,2327552
```
