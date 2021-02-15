- [VSCODE](#vscode)
    - [regex replace in file](#regex-replace-in-file)
- [VIM](#vim)
  - [vim-surround](#vim-surround)
  - [vim-sneak](#vim-sneak)
  - [vim-easymotion](#vim-easymotion)
  - [CamelCaseMotion](#camelcasemotion)

## VSCODE

#### regex replace in file

- `command + option + f`
- example 1 - add "" to every lines >> `.+` => `"$0"`
- stripe "" from every lines >> `"(.+)"` => `$1`
- swap order of words >> `(\w+) (\w+)` => `$2 $1`


## VIM

- `<C-o>` - go to last edited place backward
- `<C-i>` - go to last edited place forward

### [vim-surround](https://github.com/tpope/vim-surround)

- `S <desired char>` - Surround when in visual modes (surrounds full selection)
- `d s <existing char>` - Delete existing surround

### [vim-sneak](https://github.com/justinmk/vim-sneak)

- `s<char><char>` - Move forward to the first occurrence of `<char><char>`
- `S<char><char>` - Move backward to the first occurrence of `<char><char>`

### [vim-easymotion](https://github.com/easymotion/vim-easymotion)

- `<leader>s <char><char>` - Search character

### [CamelCaseMotion](https://github.com/bkad/CamelCaseMotion)

- `<leader>w` - Move forward to the start of the next camelCase or snake_case word segment.
- `<leader>b` - Move back to the prior beginning of a camelCase or snake_case word segment.
- `<leader>e` - Move forward to the next end of a camelCase or snake_case word segment.
