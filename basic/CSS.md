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
