_goose() {
  local cur prev opts
  COMPREPLY=()
  cur="${COMP_WORDS[COMP_CWORD]}"
  prev="${COMP_WORDS[COMP_CWORD-1]}"
  opts="--data --help --outputdir --templatedir --verbose"

  if [[ ${cur} == -* ]] ; then
    COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    return 0
  else
    goose_dir=`ls -d ~/.goose/* | xargs -n1 basename`
    words="$goose_dir"

    COMPREPLY=($(compgen -W "$words" -- $cur))
    return 0
  fi
}

complete -o default -F _goose goose

