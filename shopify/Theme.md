# Shopify Theme Liquid Code Snippets

- [Shopify Theme Liquid Code Snippets](#shopify-theme-liquid-code-snippets)
  - [New theme helper](#new-theme-helper)
  - [Add link loop](#add-link-loop)
  - [Add custom fileds](#add-custom-fileds)
    - [Add fields to product form](#add-fields-to-product-form)
    - [Add fields to cart form](#add-fields-to-cart-form)
    - [Add fields to the customer registration form](#add-fields-to-the-customer-registration-form)
  - [Product form](#product-form)
    - [Minimal product form](#minimal-product-form)
    - [Show full featured collection of products](#show-full-featured-collection-of-products)
  - [Show theme information in console with theme.liquid](#show-theme-information-in-console-with-themeliquid)
  - [Multiple currency selector](#multiple-currency-selector)
  - [Member only page](#member-only-page)
  - [Member only and special tagged customer only](#member-only-and-special-tagged-customer-only)
  - [Add recommend section in product page with alphinejs](#add-recommend-section-in-product-page-with-alphinejs)
  - [Use money formatter with multiple currenies](#use-money-formatter-with-multiple-currenies)
  - [Cart attribute](#cart-attribute)
    - [Add cart attribute](#add-cart-attribute)
    - [Render added cart attribute](#render-added-cart-attribute)
  - [Add item to cart with fetch()](#add-item-to-cart-with-fetch)
  - [Create product handles array and render product list](#create-product-handles-array-and-render-product-list)
  - [Nested for loop get index](#nested-for-loop-get-index)
  - [Show tagged articles with limited number](#show-tagged-articles-with-limited-number)
  - [Storefront search sytax](#storefront-search-sytax)
    - [search](#search)
    - [collection filter](#collection-filter)
  - [Show vendor list](#show-vendor-list)
  - [Show multiple options selector in product page](#show-multiple-options-selector-in-product-page)
  - [Toggle variant options](#toggle-variant-options)

> [offical liquid code examples](https://shopify.github.io/liquid-code-examples/)

> [extra snippets](https://github.com/vikrantnegi/shopify-code-snippets)

> [extra snippets](https://github.com/freakdesign/Shopify-code-snippets)

## New theme helper

```html
{{ 'main.min.css' | asset_url | stylesheet_tag }} {{ 'main.min.js' | asset_url |
script_tag }}
<script>
  // TODO: remove this helper
  {% assign current_handle = '' %}
  {% case template %}
    {% when 'page' %}
      {% assign current_handle = page.handle %}
    {% when 'blog' %}
      {% assign current_handle = blog.handle %}
    {% when 'article' %}
      {% assign current_handle = blog.handle %}
    {% when 'collection' %}
      {% assign current_handle = collection.handle %}
    {% when 'product' %}
      {% assign current_handle = product.handle %}
  {% endcase %}
  {% assign current_url = '' %}

  {% case template %}
    {% when 'page' %}
      {% assign current_url = page.url %}
    {% when 'blog' %}
      {% assign current_url = blog.url %}
    {% when 'article' %}
      {% assign current_url = blog.url %}
    {% when 'collection' %}
      {% assign current_url = collection.url %}
    {% when 'product' %}
      {% assign current_url = product.url %}
  {% endcase %}
    console.log('template: {{ template }}, theme.name: {{ theme.name }}');
    console.log('current_handle: {{ current_handle }}');
    console.log('current_url: {{ current_url }}');
</script>
```

## Add link loop

```
{% for category in linklists.category.links %}
  {{ category.title }}
{% endfor %}
```

## Add custom fileds

### Add fields to product form

```html
<p class="line-item-property__field">
  <label for="your-name">Your name</label>
  <input id="your-name" type="text" name="properties[Your name]" />
</p>
```

### Add fields to cart form

```html
<p class="cart-attribute__field">
  <label for="your-name">Your name</label>
  <input id="your-name" type="text" name="attributes[Your name]" value="{{
  cart.attributes["Your name"] }}">
</p>
```

### Add fields to the customer registration form

```html
<label for="CustomerFormAllergies">Allergies</label>
<input
  type="text"
  id="CustomerFormAllergies"
  name="customer[note][Allergies]"
  placeholder="Allergies"
/>
```

## Product form

### Minimal product form

```liquid
<form action="/cart/add" method="post">
  <select name="id">
    {% for variant in product.variants %}
    {% if variant.available %}
    <option value="{{ variant.id }}">
      {{ variant.title }}
    </option>
    {% else %}
    <option disabled="disabled">{{ variant.title }} - {{ 'products.product.sold_out' | t }}</option>
    {% endif %}
    {% endfor %}
  </select>
  <input type="number" name="quantity" value="1" min="1" class="QuantityInput">
  <button type="submit">カートにテスト追加</button>
</form>

```

### Show full featured collection of products

```liquid
 {% for product in collections['ADIDAS'].products limit: 6 %}
		<a href="{{ product.url }}">
			<img src="{{ product.featured_image | img_url }}" alt="" srcset="">
			<h1>{{ product.title }}</h1>
		</a>
		<div>
				<a href="{{ product.url }}">
						<img src="{{ product.featured_image | img_url :'x600'}}" alt="" srcset="">
				</a>
		</div>
		<div>
			<h1>{{ product.title }}</h1>
			<div class="plan">
					<h3 id="_product_price_{{ product.id }}">{{ product.price | money }}</h3>
			</div>
			{{ product.description }}

			{% form 'product', product, data-productid: product.id %}
			<div class="selector-wrapper js product-form__item" style="display: {% if product.has_only_default_variant %} none;{% endif %}">
					<p>Variants</p>
					<label {% if option.name == 'default' %}class="label--hidden"
							{% endif %}for="SingleOptionSelector-{{ forloop.index0 }}">
							{{ option.name }}
					</label>
					<select name="id" data-productid="{{ product.id }}" id="ProductSelect-{{ product.id }}">
							{% for variant in product.variants %}
							<option price="{{ variant.price | money }}" value="{{ variant.id }}"
									{%- if variant == current_variant %} selected="selected" {%- endif -%}>
									{{ variant.title }} {%- if variant.available == false %} -
									{{ 'products.product.sold_out' | t }}{% endif %}
							</option>
							{% endfor %}
					</select>
			</div>


			<div class="product-form__controls-group">
					<div class="product-form__item">
							<div>Quantity</div>
							<input type="number" id="Quantity-{{ section.id }}" name="quantity" value="1" min="1"
									pattern="[0-9]*" class="product-form__input product-form__input--quantity"
									data-quantity-input>
					</div>
			</div>

			<button class="btn product-form__cart-submit btn--secondary-accent" type="submit" name="add"
					{% if product.available %}{% else %}disabled{% endif %}>
					{% unless product.available %}
					{{ 'products.product.sold_out' | t }}
					{% else %}
					{{ 'products.product.add_to_cart' | t }}
					{% endunless %}
			</button>
			{{ form | payment_button }}
			{% endform %}
	</div>
	<script>
		if(document.querySelector('#ProductSelect-{{ product.id }}')){
			const select_{{ product.id }} = document.querySelector('#ProductSelect-{{ product.id }}');
			select_{{ product.id }}.addEventListener('change', () => {
					const price = select_{{ product.id }}.selectedOptions[0].getAttribute('price')
					document.querySelector('#_product_price_{{ product.id }}').innerHTML = price;
			});
		}
	</script>
{% endfor %}
```

## Show theme information in console with theme.liquid

## Multiple currency selector

```liquid
  {% form 'currency' %}
    {{ form | currency_selector }}
  {% endform %}
```

```liquid
{% form 'currency' %}
  <select name="currency">
    {% for currency in shop.enabled_currencies %}
{% if currency == cart.currency %}
  <option selected="true" value="{{ currency.iso_code }}">{{currency.iso_code}} {{currency.symbol}}</option>
  {% else %}
  <option value="{{ currency.iso_code }}">{{currency.iso_code}} {{currency.symbol}}</option>
{% endif %}
    {% endfor %}
  </select>
{% endform %}
```

```js
$('.shopify-currency-form select').on('change', function () {
  $(this).parents('form').submit();
});
```

## Member only page

```liquid
{% unless customer %}
    {% if template contains 'customers' %}
        {% assign send_to_login = false %}
    {% else %}
        {% assign send_to_login = true %}
    {% endif %}
{% endunless %}

{% if send_to_login %}
<meta content="0; url=/account/login?checkout_url=/" http-equiv="refresh" />
{% else %}
CONTENT
{% endif %}
```

## Member only and special tagged customer only

```liquid
{% if customer %}
  {% for tag in customer.tags %}
    {% unless tag contains 'multipass'  %}
      <script type="text/javascript">
        location.href = "/";
      </script>
    {% endunless %}
  {% endfor %}
  {% if customer.tags.size == 0 %}
	  <script type="text/javascript">
        location.href = "/";
      </script>
  {% endif %}
{% else %}
  <script type="text/javascript">
    location.href = "/";
  </script>
{% endif %}


```

## Add recommend section in product page with alphinejs

```html
<ul class="brandList" x-data="recommendData()" x-init="init()">
  <template x-for="recommend in recommendedProducts" :key="recommend.id">
    <li>
      <a :href="recommend.url">
        <div class="photoBox">
          <img :src="recommend.featured_image" :alt="recommend.title" />
        </div>
        <p class="ttl" x-text="recommend.title"></p>
        <p class="price" x-text="recommend.price"></p>
        <p class="tag" x-text="recommend.vendor"></p>
        {% render 'sake-stars' %}
      </a>
    </li>
  </template>
</ul>
<script>
  const URL =
    '{{ routes.product_recommendations_url }}' +
    '.json?product_id=' +
    '{{ product.id }}' +
    '&limit=4';
  function recommendData() {
    return {
      recommendedProducts: [],
      init() {
        fetch(URL)
          .then((response) => response.json())
          .then((response) => {
            this.recommendedProducts = response.products;
          });
      }
    };
  }
</script>
```

## Use money formatter with multiple currenies

- `{{ "{{ this " }}}}` => `{{ this }}` in liquid.

```html
<script>
  var theme = {
  	moneyFormat: {{ shop.money_format | json }}
  }
  theme.moneyFormat.replace(/{{ "{{[a-zA-Z0-9_]*" }}}}/, p.price)
</script>
```

## Cart attribute

### Add cart attribute

- cart note

```html
<div>
  <label>Add a note to your order</label>
  <textarea name="note"></textarea>
</div>
```

```html
<p class="cart-attribute__field">
  <label for="your-name">memo1</label>
  <textarea id="your-name" name="attributes[memo1]">
{{ cart.attributes["memo1"] }}</textarea
  >
</p>

<p class="cart-attribute__field">
  <label for="your-name">memo2</label>
  <textarea id="your-name" name="attributes[memo2]">
{{ cart.attributes["memo2"] }}</textarea
  >
</p>
```

### Render added cart attribute

```html
{% if order.attributes %}
<ul>
  {% for attribute in order.attributes %}
  <li><strong>{{ attribute | first }}</strong>: {{ attribute | last }}</li>
  <script>
    console.log('{{ attribute }}');
  </script>
  {% endfor %}
</ul>
{% endif %}
```

```
deliveryDate2020-10-31
deliveryTime16～18
deliveryCode1618

```

## Add item to cart with fetch()

```js
(function () {
  var addData = {
    id: 21373873027 /* for testing, change this to a variant ID on your store */,
    quantity: 1
  };

  fetch('/cart/add.js', {
    body: JSON.stringify(addData),
    credentials: 'same-origin',
    headers: {
      'Content-Type': 'application/json',
      'X-Requested-With':
        'xmlhttprequest' /* XMLHttpRequest is ok too, it's case insensitive */
    },
    method: 'POST'
  })
    .then(function (response) {
      return response.json();
    })
    .then(function (json) {
      /* we have JSON */
      console.log(json);
    })
    .catch(function (err) {
      /* uh oh, we have error. */
      console.error(err);
    });
})();
```

## Create product handles array and render product list

```liquid
{% assign prodlist = '' %}
{% for item in items limit: limit %}
    {% prodlist = prodlist | append: item.handle | append: ';' %}
{% endfor %}
{% assign myproducts = prodlist | remove_last: ';' | split: ';' %}
{% for handle in myproducts %}
  {{ handle }}
  {{ all_products[ handle ].title }}
{% endfor%}

```

## Nested for loop get index

```liquid
{% for outerItem in outerItems %}
    {% assign outer_forloop = forloop %}
    {% for item in items%}
        <div>{{ outer_forloop.counter }}.&nbsp;{{ item }}</div>
    {% endfor %}
{% endfor %}
```

## Show tagged articles with limited number

```liquid
{% assign count = 1 %}
{% for article in blogs["handle"].articles %}
  {% if article.tags contains 'tag name' and count < 9 %}
  {% assign count = count | plus: 1 %}
  <li><a href="{{ article.url }}"><div class="photo"><img src="{{ article | img_url: 'master' }}" alt="{{ article.title }}"></div><div class="txtInner">
    <p class="ttl">{{ article.title }}</p>
    <p>{{ article.metafields.global.products }}</p>
    <p class="link"><span>detail</span></p>
  </div></a></li>
  {% endif %}
{% endfor %}
```

## Storefront search sytax

### search

- [Storefront search](https://shopify.dev/tutorials/use-advanced-query-parameters-in-shopify-online-store-search)
- [Searching your store](https://help.shopify.com/en/manual/online-store/os/storefront-search)

`https://xxx.myshopify.com/search?q=classic+tag%3A女性%2C+adidas+product_type%3AACCESSORIES+vendor%3Aadidas&options%5Bprefix%5D=last`
`classic tag:女性, adidas product_type:ACCESSORIES vendor:adidas`

### collection filter

- [customize-theme-filter-collections-with-product-tags](https://shopify.dev/tutorials/customize-theme-filter-collections-with-product-tags)
- [URL filters](https://shopify.dev/docs/themes/liquid/reference/filters/url-filters)
- [demo](https://turbo-theme.myshopify.com/collections/all/)
- [action](https://github.com/discolabs/deploy-shopify-theme-action)

`https://xxx.myshopify.com/collections/all/SCA_STOREPICKUP_PRODUCT+男性+man`

- [Product Filter & Search](https://apps.shopify.com/product-filter-search)

## Show vendor list

```liquid
{% for product_vendor in shop.vendors limit: 6 %}
  {%- assign brand = collections[product_vendor] -%}
  {% if brand %}
    <a class="py-3 mt-1 t-style-icon-size" href="{{ product_vendor | url_for_vendor }}">
      <img class="w-full" src="{{ brand.image | img_url: 'master' }}" alt="{{ product_vendor }}"/>
    </a>
  {% endif %}
{% endfor %}
```

## Show multiple options selector in product page

```html
<form action="/cart/add" method="post">
  {% if product.variants.size > 1 %} {% if product.options[0] %} {% assign used = '' %}
  <label for="select-one">{{ product.options[0] }}</label>
  <select id="select-one" onchange="letsDoThis()">
    {% for variant in product.variants %} {% unless used contains variant.option1 %}
    <option value="{{ variant.option1 }}">{{ variant.option1 }}</option>
    {% capture used %} {{ used }} {{ variant.option1 }} {% endcapture %} {% endunless %}
    {% endfor %}
  </select>
  {% endif %} {% if product.options[1] %} {% assign used = '' %}
  <label for="select-one">{{ product.options[1] }}</label>
  <select id="select-two" onchange="letsDoThis()">
    {% for variant in product.variants %} {% unless used contains variant.option2 %}
    <option value="{{ variant.option2 }}">{{ variant.option2 }}</option>
    {% capture used %} {{ used }} {{ variant.option2 }} {% endcapture %} {% endunless %}
    {% endfor %}
  </select>
  {% endif %} {% if product.options[2] %} {% assign used = '' %}
  <label for="select-one">{{ product.options[2] }}</label>
  <select id="select-three" onchange="letsDoThis()">
    {% for variant in product.variants %} {% unless used contains variant.option3 %}
    <option value="{{ variant.option3 }}">{{ variant.option3 }}</option>
    {% capture used %} {{ used }} {{ variant.option3 }} {% endcapture %} {% endunless %}
    {% endfor %}
  </select>
  {% endif %} {% endif %}
  <input
    type="hidden"
    name="id"
    id="product-select"
    value="{{ product.variants.first.id }}"
  />
</form>
<script>
  function letsDoThis() {
    {% if product.options[0] %}
      var opt1 = document.getElementById("select-one").value;{% endif %}{% if product.options[1] %}
      var opt2 = document.getElementById("select-two").value;{% endif %}{% if product.options[2] %}
      var opt3 = document.getElementById("select-three").value;{% endif %}var id = "";{% for v in product.variants %}
      if (
        opt1 == "{{ v.option1 }}"{% if product.options[1] %} && opt2 == "{{ v.option2 }}"{% endif %}{% if product.options[2] %} && opt3 == "{{ v.option3 }}"{% endif %}
      ) {
        var id = {{ v.id }};
        var price = "{{ v.price | money }}";
      }
    {% endfor %}
    if (id != "") {
      document.getElementById("product-select").value = id;
      document.getElementById("price").innerHTML = price;
    } else {
      document.getElementById("product-select").value = "";
      document.getElementById("price").innerHTML = "Unavailable";
    }
  }
</script>
```

## Toggle variant options

```js
document.addEventListener('DOMContentLoaded', () => {
  const stockObj = Array.from(document.querySelector('select[name="id"]').children).map(
    (ele, index) => {
      return {
        name: ele.innerText.replace(/\s-.+/, ''),
        disabled: ele.getAttribute('disabled') === 'disabled'
      };
    }
  );

  const toggleEnable = (ele, isEnable) => {
    if (isEnable) {
      ele.style.pointerEvents = '';
      ele.style.cursor = 'pointer';
    } else {
      ele.style.pointerEvents = 'none';
      ele.style.cursor = 'not-allowed';
    }
  };

  // 0-color, 1-size

  Array.from(
    document.querySelector('ul.SingleOptionSelectorJS-0').querySelectorAll('label')
  ).forEach((ele) => {
    ele.addEventListener('click', function () {
      const sizes = stockObj.filter((option) => option.name.includes(this.innerText));
      Array.from(
        document.querySelector('ul.SingleOptionSelectorJS-1').querySelectorAll('label')
      ).forEach((element, index) => {
        if (sizes[index].disabled) {
          toggleEnable(element, false);
        } else {
          toggleEnable(element, true);
        }
      });
    });
  });

  Array.from(
    document.querySelector('ul.SingleOptionSelectorJS-1').querySelectorAll('label')
  ).forEach((ele) => {
    ele.addEventListener('click', function () {
      const colors = stockObj.filter((option) => option.name.includes(this.innerText));
      const sizes = stockObj.filter((option) => option.name.includes(this.innerText));
      Array.from(
        document.querySelector('ul.SingleOptionSelectorJS-0').querySelectorAll('label')
      ).forEach((element, index) => {
        if (
          colors[index].disabled &&
          stockObj
            .filter((option) => option.name.includes(sizes[index].name.split('/')[0]))
            .every((option) => option.disabled)
        ) {
          toggleEnable(element, false);
        } else {
          toggleEnable(element, true);
        }
      });
    });
  });
});
```
