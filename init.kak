define-command -override kaks-connect -params 1.. -command-completion  %{
  %arg{1} sh -c %{
    export KAKS_SESSION=$1
    export KAKS_CLIENT=$2
    shift 3

    [ $# = 0 ] && set "$SHELL"

    "$@"
  } -- %val{session} %val{client} %arg{@}
}
