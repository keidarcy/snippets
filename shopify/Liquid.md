 - [slatest](https://github.com/entozoon/slatest)(hot reload with compiler version themekit)
 - [Member only page](#member-only-page)
 - [Member only and special tagged customer only](#member-only-and-special-tagged-customer-only)
 - [Show product and add cart](#show-products-add-cart)
 - [Liquid Objects](https://shopify.dev/docs/themes/liquid/reference/objects)

### Multiple Currency
```
  {{ shop.name }} process all orders in {{ shop.currency }}.
  While the content of your cart is currently displayed in {{ cart.currency.iso_code }}
  <span class="selected-currency"></span>,
  you will checkout using {{ shop.currency }} at the most current exchange rate.
```

```
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

### Member only page
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

### Member only and special tagged customer only
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


### Show products add cart

```liquid
 {% for product in collections['ADIDAS'].products limit: 6 %}
        <a href="{{ product.url }}">
          <img src="{{ product.featured_image | img_url }}" alt="" srcset="">
          <h1>{{ product.title }}</h1>
        </a>
	<form method="post" action="/cart/add">
	    {% unless product.has_only_default_variant %}
	    <div class="selector-wrapper js product-form__item">
		<p>地域</p>
		<label {% if option.name == 'default' %}class="label--hidden"
		    {% endif %}for="SingleOptionSelector-{{ forloop.index0 }}">
		    {{ option.name }}
		</label>
		<select name="id" data-productid="{{ product.id }}" id="ProductSelect-{{ section.id }}">
		    {% for variant in product.variants %}
		    <option value="{{ variant.id }}" {%- if variant == current_variant %} selected="selected"
			{%- endif -%}>
			{{ variant.title }} {%- if variant.available == false %} -
			{{ 'products.product.sold_out' | t }}{% endif %}
		    </option>
		    {% endfor %}
		</select>
	    </div>
	    {% else %}
		<input type="hidden" name="id" value="{{ product.variants.first.id }}" />
	    {% endunless %}


	    <div class="product-form__controls-group">
		<div class="product-form__item">
		    <div>数量</div>
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
	    <a class="btn" href="{{ product.url }}">詳しく見る</a>
	</form>
{% endfor %}
```
