- Pure CSS slidedown / slideup animation using transform translateY

```css
.slide-up,
.slide-down {
  overflow: hidden;
}
.slide-up > div,
.slide-down > div {
  transform: translateY(-100%);
  transition: 0.4s ease-in-out;
}
.slide-down > div {
  transform: translateY(0);
}
```

- aspect-ratio

```html
<iframe
  src="https://www.youtube.com/embed/vQAvjof1oe4"
  frameborder="0"
  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
  allowfullscreen
></iframe>

<style>
  body {
    min-height: 100vh;
    display: grid;
    place-items: center;
  }

  iframe {
    width: 75%;
    aspect-ratio: 16 / 9;
  }
</style>
```

- outline vs border

```css
.border {
  border: 2px solid red;
  outline: 2px solid black;
  border-radius: 10px;
}

.outline {
  border: 30px solid green;
  outline: 20px solid #fff;
  outline-offset: -10px;
  background: #333;
  color: white;
}
```

- rem vs em

`rem` is based on `html` tag, without setting is 16px.
`em` is based on parent tag size.

- selector

|                    |                  |                                                                                     |
| :----------------: | :--------------: | :---------------------------------------------------------------------------------: |
|   element.class    |     p.intro      |                     Selects all <p> elements with class="intro"                     |
|  element,element   |      div, p      |                   Selects all <div> elements and all <p> elements                   |
|  element element   |      div p       |                   Selects all <p> elements inside <div> elements                    |
|  element>element   |     div > p      |            Selects all <p> elements where the parent is a <div> element             |
|  element+element   |     div + p      |   Selects the first <p> element that are placed immediately after <div> elements    |
| element1~element2  |      p ~ ul      |            Selects every <ul> element that are preceded by a <p> element            |
|    [attribute]     |     [target]     |                    Selects all elements with a target attribute                     |
| [attribute=value]  | [target=_blank]  |                     Selects all elements with target="\_blank"                      |
| [attribute~=value] | [title~=flower]  |      Selects all elements with a title attribute containing the word "flower"       |
| [attribute=value]  |    [lang =en]    |         Selects all elements with a lang attribute value starting with "en"         |
| [attribute^=value] | a[href^="https"] |      Selects every <a> element whose href attribute value begins with "https"       |
| [attribute$=value] | a[href$=".pdf"]  |        Selects every <a> element whose href attribute value ends with ".pdf"        |
| [attribute*=value] | a[href*="darcy"] | Selects every <a> element whose href attribute value contains the substring "darcy" |
