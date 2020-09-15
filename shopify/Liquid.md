 - [Member only page](#member-only-page)
 - [Member only and special tagged customer only](#member-only-and-special-tagged-customer-only)
 - [Show product and add cart](#show-products-add-cart)


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
        <input type="hidden" name="id" value="{{ product.variants.first.id }}" />
        <input min="1" type="number" id="quantity" name="quantity" value="1"/>
        <input type="submit" value="Add to cart" class="btn" />
      </form> 
{% endfor %}
```
