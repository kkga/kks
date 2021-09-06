define-command -override kks-connect -params 1.. -command-completion  %{
  %arg{1} sh -c %{
    export EDITOR='kks edit'
    export KKS_SESSION=$1
    export KKS_CLIENT=$2
    shift 3

    [ $# = 0 ] && set "$SHELL"

    "$@"
  } -- %val{session} %val{client} %arg{@}
}
