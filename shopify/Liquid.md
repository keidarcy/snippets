 - [member only page](#member-only-page)


### Member only page
```
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
```
