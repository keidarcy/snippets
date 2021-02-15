## Regex

- [regexr](https://regexr.com/)

- [Regex](#regex)
  - [Basic](#basic)
  - [Group](#group)
  - [Examples](#examples)

### Basic

- `.` => Matches any character except line breaks.
- `^` => Begining of the string.
- `$` => End of the string.
- `[^]` => Negataed set.
- `[]{3, }` => Match between 3 to unlimited characters.
- `\d <=> [0-9]`, `\D <=> [^0-9]`, `\w <=> [a-zA-Z0-9_]` // (alphanumeric & underscore).
- `\s <=> [\r\n\t\v ]` => Matches any whitespace character (spaces, tabs, line breaks).
- `{0,1} <=> ?` => Match between 0 and 1 times.
- `{0, } <=> *` => Match between 0 to unlimited times.
- `{1, } <=> +` => Match between 1 to unlimited times.

### Group

- use group

```
`xx@gmail.com`.match(/^([a-zA-Z0-9]\w*)@gmail\.com$/)[1] // 'xx'
```

- rename group

```
`xx@gmail.com`.match(/^(?<first>[a-zA-Z0-9]\w*)@gmail\.com$/).groups.first // 'xx'
```

- reuse group in regex

```
^(\d\d)\1$    // match 1212
```

- reuse group with renamed

```
^(?<first>\d+)\k<first>$
```

- `(?=` Positive lookahead

```js
'foobar, foopoo'.replace(/foo(?=bar)/g, 'replaced'); // "replacedbar, foopoo"
```

- `(?!` Negative lookahead

```js
'foobar, foopoo'.replace(/foo(?!bar)/g, 'replaced'); // "foobar, replacedpoo"
```

- `(?<=` Positive lookbehiend

```js
'foobar, foopoo'.replace(/(?<=foo)bar/g, 'replaced'); // "fooreplaced, foopoo"
```

- `(?<!` Negative lookbehiend

```js
'foobar, foopoo'.replace(/(?<=foo)bar/g, 'replaced'); // "foobar, fooreplaced"
```

### Examples

- match email address
  `^[a-zA-Z0-9]\w*@gmail\.com$` => `xx@gmail.com`

- match float string and get fixed
  `'rgba(100,150,200,.72312332)'.replace(/(\.\d{2})[0-9]*/,"$1"); // rgba(100,150,200,.72)`

- match two separate part
  `'/api/xxxx/edit?uu=xxxx&id=1&type=&page='.replace(/\/api|edit.*/g, '') // /xxxx/`