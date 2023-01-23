## MEmo for datascienceatthecommandline

```bash
< alice.txt | tr '[:upper:]' '[:lower:]' | grep -oE '\w{2,}' |\ngrep -E '^a.*e$' | sort | uniq -c | sort -nr |\nawk '{print $2" ," $1}' | header -a char,number | csvlook|head -n 12
```

```bash
echo '  hello   world! ' | sed -re 's/\shello/bye/;s/\s+/ /g;s/\s+//''\n'
```

## Difference between `""`, `''`, `backtil`

```bash

$ export foo=uname
$ echo '$foo' # $foo
$ echo "$foo" # uname
$ echo `$foo` # Linux
$ echo $($foo) # Linux

```

## Useful aliaes

#### global installed npm packages

```
npm list -g --depth=0

```

#### local ip

```
ifconfig | grep 'inet 192'| awk '{ print $2}'
```

## practical solutions

#### `node (eval):1: command not found: _node` zsh problem

- update omz `upgrade_oh_my_zsh`
- delete all caches `rm ~/.zcompdump*`

## jq