# kks

[![Hits-of-Code](https://hitsofcode.com/github/kkga/kks?branch=main)](https://hitsofcode.com)

Handy Kakoune companion.

## Installation

### From release binaries

Download the compiled binary for your system from
[Releases](https://github.com/kkga/kks/releases) page and put it somewhere in
your `$PATH`.

### From source

Requires [Go](https://golang.org/) installed on your system.

Clone the repository and run `go build`, then copy the compiled binary somewhere
in your `$PATH`.

If Go is [configured](https://golang.org/ref/mod#go-install) to install packages
in `$PATH`, it's also possible to install without cloning the repository: run
`go install github.com/kkga/kks@latest`.

### AUR

`kks` is packaged in the Arch User Repository:
https://aur.archlinux.org/packages/kks/

## Kakoune and shell integration

### Kakoune configuration

Source `kks init` to add `kks-connect` command to Kakoune...

```kak
eval %sh{ kks init }
```

... and use your terminal integration to connect
[provided scripts](#provided-scripts), for example:
`kks-connect terminal kks-files`.

### Shell configuration

You may want to set the `EDITOR` variable to `kks edit` so that connected
programs work as intended:

```sh
export EDITOR='kks edit'
```

Possibly useful aliases:

```sh
alias k="kks edit"
alias ks="eval $(kks-select)"
alias ka="kks attach"
alias kkd="kks kill; unset KKS_SESSION KKS_CLIENT" # kill+detach
alias kcd="cd $(kks get %sh{pwd})"
```

### Kakoune mappings example

```kak
map global normal -docstring 'terminal'         <c-t> ': kks-connect terminal<ret>'
map global normal -docstring 'files'            <c-f> ': kks-connect popup kks-files<ret>'
map global normal -docstring 'buffers'          <c-b> ': kks-connect popup kks-buffers<ret>'
map global normal -docstring 'files by content' <c-g> ': kks-connect popup kks-grep<ret>'
map global normal -docstring 'lines in buffer'  <c-l> ': kks-connect popup kks-lines<ret>'
map global normal -docstring 'recent files'     <c-r> ': kks-connect popup kks-mru<ret>'
map global normal -docstring 'lf'               <c-h> ': kks-connect panel kks-lf<ret>'
map global normal -docstring 'lazygit'          <c-v> ': kks-connect popup lazygit<ret>'
```

For more terminal integrations and for the (quite handy) `popup` command, see:

- [alacritty.kak](https://github.com/alexherbo2/alacritty.kak)
- [foot.kak](https://github.com/kkga/foot.kak)

## Commands

This is the output of `kks -h`. Certain commands take additional flags, see
`kks <command> -h` to learn more.

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
  get            get %val{..}, %opt{..} and friends
  cat            print buffer content
  env            print env
  init           print Kakoune definitions

ENVIRONMENT VARIABLES
  KKS_SESSION    Kakoune session
  KKS_CLIENT     Kakoune client

Use "kks <command> -h" for command usage.
```

### Unknown command

When unknown command is run, `kks` will try to find an executable named
`kks-<command>` in `$PATH`. If the executable is found, `kks` will run it with
all arguments that were provided to the unknown command.

## Configuration

`kks` can be configured through environment variables.

### Automatic sessions based on git directory

```
export KKS_USE_GITDIR_SESSIONS=1
```

When `KKS_USE_GITDIR_SESSIONS` is set to any value and `KKS_SESSION` is empty,
running `kks edit` will do the following:

- if file is inside a git directory, `kks` will search for an existing session
  based on top-level git directory name and connect to it;
- if a session for the directory doesn't exist, `kks` will start a new session
  and connect to it.

### Default session

```
export KKS_DEFAULT_SESSION='mysession'
```

When context is not set (`KKS_SESSION` is empty), running `kks edit` will check
for a session defined by `KKS_DEFAULT_SESSION` variable. If the session is
running, `kks` will connect to it instead of starting a new session.

`kks` will not start the default session if it's not running. You can use the
autostarting mechanism of your desktop to start it with `kak -d -s mysession`.

## Provided scripts

| script                                       | function                                                |
| -------------------------------------------- | ------------------------------------------------------- |
| [`kks-buffers`](./scripts/kks-buffers)       | pick buffers                                            |
| [`kks-fifo`](./scripts/kks-fifo)             | pipe stdin to Kakoune fifo buffer                       |
| [`kks-files`](./scripts/kks-files)           | pick files                                              |
| [`kks-grep`](./scripts/kks-grep)             | search for pattern in working directory                 |
| [`kks-lf`](./scripts/kks-lf)                 | open [lf] with current buffer selected                  |
| [`kks-lines`](./scripts/kks-lines)           | jump to line in buffer                                  |
| [`kks-md-heading`](./scripts/kks-md-heading) | jump to markdown heading                                |
| [`kks-mru`](./scripts/kks-mru)               | pick recently opened file                               |
| [`kks-select`](./scripts/kks-select)         | select Kakoune session and client to set up environment |

[lf]: https://github.com/gokcehan/lf

## Similar projects

- [kakoune.cr](https://github.com/alexherbo2/kakoune.cr)
- [kakoune-remote-control](https://github.com/danr/kakoune-remote-control)
