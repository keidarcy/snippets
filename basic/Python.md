## Connect google spread sheet with colab

### read google sheet in google colab

```python
!pip install --upgrade -q gspread

from google.colab import auth
auth.authenticate_user()

import gspread
from oauth2client.client import GoogleCredentials

gc = gspread.authorize(GoogleCredentials.get_application_default())
worksheet = gc.open_by_url(URL).sheet1

# get_all_values gives a list of rows.
rows = worksheet.get_all_values()

# Convert to a DataFrame and render.
import pandas as pd
pdrows=pd.DataFrame.from_records(rows)

pdrows
```

### save sheet from colab

```python
from gspread_dataframe import set_with_dataframe

pdnew = pd.DataFrame.from_records(pdrows)

sheet = gc.open_by_url(NEW_URL).sheet1
set_with_dataframe(sheet, pdnew)
```

## Range function

> [1, 2, 3, 4, 5]

- python3

```python
list(range(1,6))
```

- php

```php
range(1,5)
```

- js

```
[...Array(5).keys()].splice(1)
```

### lambda

```py
getBoolean = lambda x: x == '○'
```

### Useful functions

```py
import json
import math
import random
import string


get_boolean = lambda x: x == '○'

def is_not_NaN(num):
    return num == num

def add_if_not_none(list_val):
    tags = []
    for val in list_val:
        if(is_not_NaN(val)):
            tags.append(val)
    return tags
get_image_path_if_not_NaN = lambda x: 'https://xxxx' + x if(is_not_NaN(x)) else ''

def get_price_number(price):
    return price.replace('¥','').replace(',','')

get_value_if_string = lambda x: x if(type(x) == str) else ''


def get_random_string(length):
    letters = string.ascii_lowercase
    result_str = ''.join(random.choice(letters) for i in range(length))
    return result_str

```

### if elif else

```py
num = 7
if num in [1,2]:
    print(num)
elif num in [3,4]:
    print(num)
elif num in [5,6,7]:
    print(num)
else
    print(num)
```

### Dictionaries

```py
student = {'name': 'John', 'age': 25, 'courses': ['Math', 'Compci']}
print(student['courses'])
print(student['phone'])
print(student.get('phone', 'Not Found'))

student['phone'] = 55555
student['name'] = 'Mary'

student.update({'name': 'Smith', 'age': 26})

age = student.pop('age')

print(student.items())
print(student.values())

for key, value in student.items():
    print(key, value)
    
people = {1: {'Name': 'John', 'Age': '27', 'Sex': 'Male'},
          2: {'Name': 'Marie', 'Age': '22', 'Sex': 'Female'}}

for p_id, p_info in people.items():
    print("\nPerson ID:", p_id)
    
    for key in p_info:
        print(key + ':', p_info[key])

```

### list, tuple, dic, set

```
| method | initialize | how to get value | add or remove value | get value  |
|--------|------------|------------------|---------------------|------------|
| list   | []         | index            | OK                  | OK         |
| tuple  | ()         | index            | NO                  | OK         |
| dic    | {}         | key              | OK                  | OK         |
| set    | set()      | none             | OK                  | OK(random) |
```
