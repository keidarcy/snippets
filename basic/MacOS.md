#### ssh in pi from mac without password

```
# from mac
scp pi_rsa.pub pi@0.0.0.000:/home/pi/.ssh

# in pi
cat pi_rsa.pub >> authorized_keys

scp -i {{keypath}} {{filename}} {{user}}@{{host}}:{{targetpath}}
```

## MacOs

- Install broken app

```
sudo xattr -r -d com.apple.quarantine /Applications/XXX.app
```

- [Change](https://github.com/stuartcryan/custom-iterm-applescripts-for-alfred) Alfred default terminal to iterm2

## Hotkeys

- mutiple cursor different rows
  `command` + `option` + arrow key
