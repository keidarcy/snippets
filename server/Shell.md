#### `node (eval):1: command not found: _node` zsh problem

- update omz `upgrade_oh_my_zsh`
- delete all caches `rm ~/.zcompdump*`

- global installed npm packages

```
npm list -g --depth=0

```

- local ip

```
ifconfig | grep 'inet 192'| awk '{ print $2}'
```

## Difference between `""`, `''`, `backtil`

```bash

$ export foo=uname
$ echo '$foo' # $foo
$ echo "$foo" # uname
$ echo `$foo` # Linux
$ echo $($foo) # Linux

```
