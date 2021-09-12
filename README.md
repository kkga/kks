# KKS

Handy Kakoune companion.

## Installation

### Download release build

Download the compiled binary for your system from
[Releases](https://github.com/kkga/kks/releases) page and put it somewherr in
your `$PATH`.

### Build from source

Requires [Go](https://golang.org/) installed on your system..

Clone the repository and run `go build`, then copy the compiled binary somewhere
in your `$PATH`.

If Go is [configured](https://golang.org/ref/mod#go-install) to install packages
in `$PATH`, it's also possible to install without cloning the repository: run
`go install github.com/kkga/kks@latest`.

## Manual

### Kakoune configuration

```kak
eval %sh{ kks init }
```

### Shell configuration example

```sh
export EDITOR=`kks edit`

alias k='kks edit'
alias ks='eval (kks-select)'
alias kcd='cd (kks get %sh{pwd})'
alias ka='kks attach'
alias kl='kks list'
```

### Usage

```
USAGE
  kks <command> [-s <session>] [-c <client>] [<args>]

COMMANDS
  new, n         create new session
  edit, e        edit file
  send, s        send command
  attach, a      attach to session
  kill           kill session
  ls             list sessions and clients
  get            get %{val}, %{opt} and friends
  cat            print buffer content
  env            print env
  init           print Kakoune definitions

ENVIRONMENT VARIABLES
  KKS_SESSION    Kakoune session
  KKS_CLIENT     Kakoune client

Use "kks <command> -h" for command usage.
```

### Provided scripts

- [`kks-buffers`](./scripts/kks-buffers) -- pick buffers
- [`kks-files`](./scripts/kks-files) -- pick files
- [`kks-grep`](./scripts/kks-grep) -- search for pattern in working directory
- [`kks-lines`](./scripts/kks-lines) -- jump to line in buffer
- [`kks-mru`](./scripts/kks-mru) -- pick recently opened file
- [`kks-select`](./scripts/kks-select) -- select Kakoune session and client to
  set up environment
