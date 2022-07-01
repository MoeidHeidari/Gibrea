
<img title="" src="images/gibrea-logo.png" alt="" width="209" data-align="center">
### Overview

Gibrea is a git batch repository aggregator platform that can take an input file containing thousands of repositories inside and spread the the migrations among several cores.

---

### commands

```bash
go run main.go -h
        .d8888b.  8888888 888888b.   8888888b.  8888888888        d8888 
        d88P  Y88b   888   888  "88b  888   Y88b 888              d88888 
        888    888   888   888  .88P  888    888 888             d88P888 
        888          888   8888888K.  888   d88P 8888888        d88P 888 
        888  88888   888   888  "Y88b 8888888P"  888           d88P  888 
        888    888   888   888    888 888 T88b   888          d88P   888 
        Y88b  d88P   888   888   d88P 888  T88b  888         d8888888888 
         "Y8888P88 8888888 8888888P"  888   T88b 8888888888 d88P     888


    
Gibrea is a batch aggregator program that aggregates a bunch of git repositories from different sources and migrate them to a git repository system

Usage:
  main [flags]
  main [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  logs        shows the the logs of running program
  pause       pauses the prrogramm execution temorarily with last state persisted
  resum       reads the last state of the program and resumes the execution
  start       starts the gibrea programm
  status      shows the current status of the programm
  stop        stops the program with state persisted

Flags:
  -h, --help     help for main
  -t, --toggle   Help message for toggle

Use "main [command] --help" for more information about a command.
```

### source code

```bash
git clone https://github.com/MoeidHeidari/Gibrea.git
cd Gibrea
```

---

### Build

```bash
go get
go build
```

### Test

```bash
go test ./...
```

### Documentation

```bash
godoc -http=:6060
```

http://localhost:6060