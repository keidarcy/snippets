### Bring reactivity to JsDom

```js
let root = document.querySelector('[x-data]') as HTMLElement;

let directives = {
  'x-text': (el: HTMLElement, value: string) => {
    el.innerText = value;
  },
  'x-show': (el: HTMLElement, value: boolean) => {
    el.style.display = value ? 'block' : 'none';
  }
};

let rawData = getInitailData();

let data = observe(rawData);

registerListeners();

function registerListeners() {
  walkDom(root, (el: HTMLElement) => {
    Array.from(el.attributes).forEach((attribute) => {
      if (!attribute.name.startsWith('@')) return;

      let event = attribute.name.replace('@', '');

      el.addEventListener(event, () => {
        eval(`with (data) { (${attribute.value}) }`);
      });
    });
  });
}

function observe(data: any) {
  return new Proxy(data, {
    set(target, key, value) {
      target[key] = value;
      refreshDom();
      return true;
    }
  });
}

function getInitailData(): {} {
  let dataString = root.getAttribute('x-data');
  return eval(`(${dataString})`);
}

type VoidFunc = (el: HTMLElement) => void;

function walkDom(el: HTMLElement, callback: VoidFunc) {
  callback(el);

  el = el.firstElementChild as HTMLElement;

  while (el) {
    walkDom(el, callback);
    el = el.nextElementSibling as HTMLElement;
  }
}

function refreshDom() {
  walkDom(root, (el) => {
    if (el.hasAttribute('x-text')) {
      const expression = el.getAttribute('x-text');
      el.innerText = eval(`with (data){ ${expression} }`);
    }
    Array.from(el.attributes).forEach((attribute) => {
      if (!Object.keys(directives).includes(attribute.name)) return;

      directives[attribute.name](el, eval(`with (data) {(${attribute.value})}`));
    });
  });
}

refreshDom();
console.log(data);

```
