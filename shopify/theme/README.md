# Shopify Theme Liquid Code Snippets

- [Shopify Theme Liquid Code Snippets](#shopify-theme-liquid-code-snippets)
	- [Multiple Currency](#multiple-currency)
	- [Member only page](#member-only-page)
	- [Member only and special tagged customer only](#member-only-and-special-tagged-customer-only)
	- [Show full featured collection of products](#show-full-featured-collection-of-products)
	- [Find current handle](#find-current-handle)
	- [Find current url](#find-current-url)

>[other snippets](https://github.com/vikrantnegi/shopify-code-snippets)

## Multiple Currency

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

## Show full featured collection of products

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

## Find current handle
```
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
```

## Find current url
```
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
```