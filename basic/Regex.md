## Regex

- [regexr](https://regexr.com/)
- [regex101](https://regex101.com/)

- [Regex](#regex)
- [Basic](#basic)
- [Group](#group)
  - [name group](#name-group)
  - [use group in regex](#use-group-in-regex)
      - [use named group](#use-named-group)
  - [advance group useage](#advance-group-useage)
    - [`(?=...)` Positive lookahead](#-positive-lookahead)
    - [`(?!...)` Negative lookahead](#-negative-lookahead)
    - [`(?<=...)` Positive lookbehind](#-positive-lookbehind)
    - [`(?<!...)` Negative lookbehiend](#-negative-lookbehiend)
  - [Examples](#examples)

## Basic

- `^`     >> asserts position at start of a line.
- `$`     >> asserts position at the end of a line.
- `[^]`   >> Negataed set.
- `.`     >> matches any character (except for line terminators).
- `{3, }` >> Matches between 3 and unlimited times, as many times as possible, giving back as needed.
- `\d`    >> matches a digit (equal to [0-9])
- `\D`    >> matches any character that's not a digit (equal to [^0-9])
- `\w`    >> matches any word character (equal to [a-zA-Z0-9_])
- `\W`    >> matches any non-word character (equal to [^a-zA-Z0-9_])
- `\s`    >> matches any whitespace character (equal to [\r\n\t\f\v \u00a0\u1680\u2000-\u200a\u2028\u2029\u202f\u205f\u3000\ufeff])
- `\S`    >> matches any non-whitespace character (equal to [^\r\n\t\f\v \u00a0\u1680\u2000-\u200a\u2028\u2029\u202f\u205f\u3000\ufeff])
- `?`     >> Matches between zero and one times, as many times as possible, giving back as needed(equal to `{0,1}`)
- `*`     >> Matches between zero and unlimited times, as many times as possible, giving back as needed(equal to `{0,}`)
- `+`     >> Matches between zero and unlimited times, as many times as possible, giving back as needed(equal to `{1,}`)

## Group

### name group

- `(?<NAME>)`

- `^(?<first>\w+@\w+\.\w+` >> Named Capture Group first.

```js
`fjwia@fdsfss.a`.match(/^(?<first>\w+)@\w+\.\w+/)
```

### use group in regex

- `(SOMETHING)\1`

- `^\d+\1`

##### use named group

- `(?<NAME>SOMTHING)\k<NAME>`

- `^(?<first>\d)\k<first>`


### advance group useage

#### `(?=...)` Positive lookahead

- `foobar foobaz` match the `foo` before `bar`

- `foo(?=bar)`

```js
'foobar foobaz'.replace(/foo(?=bar)/g, 'replaced'); // "replacedbar, foobaz"
```

#### `(?!...)` Negative lookahead

- `foobar foobaz` match the `foo` not before `bar`

- `foo(?!bar)`

```js
'foobar foobaz'.replace(/foo(?!bar)/g, 'replaced'); // "foobar, replacedbaz"
```

#### `(?<=...)` Positive lookbehind


- `foobar foobaz` match the `foo` behind `bar`

- `/(?<=foo)bar/`

```js
'foobar foobaz'.replace(/(?<=foo)bar/g, 'replaced'); // "fooreplaced, foobaz"
```

#### `(?<!...)` Negative lookbehiend

- `foobar foobaz` match the `foo` not behind `bar`

- `(?<!foo)bar`

```js
'foobar foobaz'.replace(/(?<!foo)bar/g, 'replaced'); // "foobar, fooreplaced"
```

### Examples

- match email address
  `^[a-zA-Z0-9]\w*@gmail\.com$` => `xx@gmail.com`

- match float string and get fixed
  `'rgba(100,150,200,.72312332)'.replace(/(\.\d{2})[0-9]*/,"$1"); // rgba(100,150,200,.72)`

- match two separate part
  `'/api/xxxx/edit?uu=xxxx&id=1&type=&page='.replace(/\/api|edit.*/g, '') // /xxxx/`

- match `aabc`
  `^(?<a>.)\k<a>(?!\k<a>)(?<b>.)(?!\k<b>|\k<a>).$`
