
##### Range function
>[1, 2, 3, 4, 5]
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

## Connect google spread sheet with colab

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
print(rows)

# Convert to a DataFrame and render.
import pandas as pd
pdrows=pd.DataFrame.from_records(rows)

pdrows
```

