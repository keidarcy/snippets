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
