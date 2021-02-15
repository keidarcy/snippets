- [Jupyter notebook](#jupyter-notebook)
- [Pandas to fomat csv file](#pandas-to-fomat-csv-file)

## Jupyter notebook

- [Jupyter notebook](#jupyter-notebook)
- [Pandas to fomat csv file](#pandas-to-fomat-csv-file)

```bash
python3 -m venv csv
```

```bash
source bin/activate
```

```bash
pip3 install jupyter
pip3 install pandas
```

```py
import pandas as pd

ser = pd.Series([12,20,30,40])
ser
```

```py
df = pd.DataFrame([[10, 'a', True], [20, False, False], [30, 'c', False], [40, 'd', True]])
df
```

```py
import numpy as np
df = pd.DataFrame(np.arange(100).reshape((25,4)))
df.head()
df.tail()
df.shape

df = pd.DataFrame(np.arange(6).reshape(3,2))
df.index = ['01', '02', '03']
df.columns = ['A', 'B']
```

```py
df = pd.DataFrame(np.arange(12).reshape((4,3)),columns=['A', 'B', 'C'],index=["1  ", '2  ', '3  ', '4  '])
df['A']
df[['A', 'B']]
df[:2]
```

```py
# loc method use name of index and columns
df.loc[:,:]
df.loc[:,'A']
# df = pd.DataFrame(np.arange(12).reshape((4,3)),columns=['A', 'B', 'C'],index=list(df.loc[:,'A']))
df.loc[:, ['A', 'B']]
df.loc['1  ', :]
df.loc[['1  ', '2  '], :]
```

```py
# iloc method use index of index and columns
df.iloc[1,1]
df.iloc[1:,1]
df.iloc[1:, :2]
```

## Pandas to fomat csv file

```py
import pandas as pd

sample = pd.read_csv('./csvs/excelify_sample.csv')
product_list0 = pd.read_csv('./csvs/product_list.csv')
product_csv0 = pd.read_csv('./csvs/product_csv.csv')
```

```py
product_list = pd.DataFrame(product_list0[1:])
product_list.index = list(product_list0.iloc[1:, 1])
product_list.columns = list(product_list0.iloc[0, :])
product_list = product_list.sort_values(by="", ascending=True, na_position='first')
product_list = product_list.groupby(product_list.index).first()
```

```py
product_csv = pd.DataFrame(product_csv0)
product_csv.index = product_csv0.loc[:,'']
product_csv.index.names = ['index']
product_csv = product_csv.sort_values(by="", ascending=True, na_position='first')
product_csv = product_csv.groupby(product_csv.index).first()

```

```py
product_data = product_list.join(product_csv.add_suffix('_2'), how='outer')
print(product_list.shape)
print(product_csv.shape)

product_data.dropna(subset=[''],inplace=True)
product_data.shape

```

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
get_image_path_if_not_NaN = lambda x: 'https:xx' + x if(is_not_NaN(x)) else ''

def get_price_number(price):
    return price.replace('¥','').replace(',','')

get_value_if_string = lambda x: x if(type(x) == str) else ''


def get_random_string(length):
    letters = string.ascii_lowercase
    result_str = ''.join(random.choice(letters) for i in range(length))
    return result_str

```

```py
newColumns = sample.columns
newColumnsList = list(newColumns)
newColumnsList.remove('Metafield: specs.range [integer]')
newColumnsList.remove('Variant Metafield: something [string]')
newColumnsList.append('Metafield: information [string]')
new_excelify = pd.DataFrame(columns=newColumnsList)
new_excelify
```

```py
index = -1
index_inline = 0
index_product = 0
# for sku, row in product_data.iterrows():
for sku, row in product_data[0:25].iterrows():
    index += 1
    index += index_inline
    index_inline = 0
#     print(index)
    new_excelify.loc[index] = ''
    if index == 0 or not new_excelify.loc[index - 1]['Handle'] == row['']:
        new_excelify.loc[index]['Handle'] = row['']
        new_excelify.loc[index]['Command'] = 'UPDATE'
        new_excelify.loc[index]['Title'] = row['   ']
        new_excelify.loc[index]['Body HTML'] = get_value_if_string(row['  -       _2']) + '\n' + get_value_if_string(row['      ー_2'])
        new_excelify.loc[index]['Vendor'] = row['    ']
        new_excelify.loc[index]['Type'] = row['  ']
        new_excelify.loc[index]['Tags'] = ','.join(add_if_not_none([row['  ー１'], row['  ー２'], row[' ー ー   ( ・   )'], row['    '], row['   '], row['   ー     '], row['        '], row['  '], row['   '], row['  ー   '], row['  '], row['    '], row['    '], row['  ']]))
        new_excelify.loc[index]['Tags Command'] = 'REPLACE'
        new_excelify.loc[index]['Updated At'] = ''
        new_excelify.loc[index]['Published'] = 'TRUE'
        new_excelify.loc[index]['Published At'] = ''
        new_excelify.loc[index]['Published Scope'] = 'global'
        new_excelify.loc[index]['Template Suffix'] = ''
        new_excelify.loc[index]['Gift Card'] = 'FALSE'
        new_excelify.loc[index]['Row'] = ''
        new_excelify.loc[index]['Top Row'] = 'TRUE'
        new_excelify.loc[index]['Custom Collections'] = ''
        new_excelify.loc[index]['Image Src'] = get_image_path_if_not_NaN(row['  -       _2'])
        new_excelify.loc[index]['Image Command'] = 'MERGE'
        new_excelify.loc[index]['Image Position'] = '1'
        new_excelify.loc[index]['Image Width'] = ''
        new_excelify.loc[index]['Image Height'] = ''
        new_excelify.loc[index]['Image Alt Text'] = row['   ']
        new_excelify.loc[index]['Variant ID'] = ''
        new_excelify.loc[index]['Variant Command'] = 'MERGE'
        if is_not_NaN(row['  ー１']):
            new_excelify.loc[index]['Option1 Name'] = 'color'
            new_excelify.loc[index]['Option1 Value'] = get_random_string(2) + '.' +row['  ー１']
        new_excelify.loc[index]['Option2 Name'] = ''
        new_excelify.loc[index]['Option2 Value'] = ''
        new_excelify.loc[index]['Option3 Name'] = ''
        new_excelify.loc[index]['Option3 Value'] = ''
        new_excelify.loc[index]['Variant Generate From Options'] = 'FALSE'
        new_excelify.loc[index]['Variant Position'] = '1'
        new_excelify.loc[index]['Variant SKU'] = row['']
        new_excelify.loc[index]['Variant Weight'] = ''
        new_excelify.loc[index]['Variant Weight Unit'] = ''
        new_excelify.loc[index]['Variant HS Code'] = ''
        new_excelify.loc[index]['Variant Country of Origin'] = 'JP'
        new_excelify.loc[index]['Variant Compare At Price'] = ''
        new_excelify.loc[index]['Variant Price'] = get_price_number(row['    '])
        new_excelify.loc[index]['Variant Cost'] = ''
        new_excelify.loc[index]['Variant Requires Shipping'] = 'TRUE'
        new_excelify.loc[index]['Variant Taxable'] = 'TRUE'
        new_excelify.loc[index]['Variant Tax Code'] = ''
        new_excelify.loc[index]['Variant Barcode'] = ''
        new_excelify.loc[index]['Variant Image'] = get_image_path_if_not_NaN(row["  -      (1)_2"])
        new_excelify.loc[index]['Variant Inventory Tracker'] = 'shopify'
        new_excelify.loc[index]['Variant Inventory Policy'] = 'deny'
        new_excelify.loc[index]['Variant Fulfillment Service'] = 'manual'
        new_excelify.loc[index]['Variant Inventory Adjust'] = ''
        new_excelify.loc[index]['Variant Inventory Qty'] = row['   _2']
        new_excelify.loc[index]['Metafield: description_tag'] = row['  -       _2']
        new_excelify.loc[index]['Metafield: information [string]'] = row['  -       _2']
        new_excelify.loc[index]['Metafield: title_tag'] = row['   ']
        metafiledDic = {
            'textable': get_boolean(row['  ']),
            'miniMark': get_boolean(row['   ー ']),
            'sisyuNormal': get_boolean(row['  ー   （  ）']),
            'sisyuDisney': get_boolean(row['  ー   （    ー）']),
            'sisyuStarwars': get_boolean(row['  ー   （  ー  ー ）']),
            'giftbox': get_boolean(row['        ']),
            'giftboxSet': get_boolean(row['            ']),
            'wrapping': get_boolean(row['      ']),
            'giftbox': get_boolean(row['          ']),
            'yahoo': get_boolean(row['Yahoo    ']),
            'color': get_boolean(row['  ー ー '])
        }
        new_excelify.loc[index]['Metafield: custom.json [json_string]'] = json.dumps(metafiledDic)

    else:
        new_excelify.loc[index]['Handle'] = row['']
        new_excelify.loc[index]['Command'] = 'UPDATE'
        new_excelify.loc[index]['Vendor'] = row['    ']
        new_excelify.loc[index]['Type'] = row['  ']
        new_excelify.loc[index]['Tags'] = ','.join(add_if_not_none([row['  ー１'], row['  ー２'], row[' ー ー   ( ・   )'], row['    '], row['   '], row['   ー     '], row['        '], row['  '], row['   '], row['  ー   '], row['  '], row['    '], row['    '], row['  ']]))
        new_excelify.loc[index]['Tags Command'] = 'REPLACE'
        new_excelify.loc[index]['Updated At'] = ''
        new_excelify.loc[index]['Published'] = 'TRUE'
        new_excelify.loc[index]['Published At'] = ''
        new_excelify.loc[index]['Published Scope'] = 'global'
        new_excelify.loc[index]['Template Suffix'] = ''
        new_excelify.loc[index]['Gift Card'] = 'FALSE'
        new_excelify.loc[index]['Row'] = ''
        new_excelify.loc[index]['Top Row'] = ''
        new_excelify.loc[index]['Custom Collections'] = ''
        new_excelify.loc[index]['Image Src'] = get_image_path_if_not_NaN(row['  -       _2'])
        new_excelify.loc[index]['Image Command'] = 'MERGE'
        new_excelify.loc[index]['Image Position'] = '1'
        new_excelify.loc[index]['Image Width'] = ''
        new_excelify.loc[index]['Image Height'] = ''
        new_excelify.loc[index]['Image Alt Text'] = row['   ']
        new_excelify.loc[index]['Variant ID'] = ''
        new_excelify.loc[index]['Variant Command'] = 'MERGE'
        if is_not_NaN(row['  ー１']):
            new_excelify.loc[index]['Option1 Name'] = 'color'
            new_excelify.loc[index]['Option1 Value'] = get_random_string(2) + '.' +row['  ー１']
        new_excelify.loc[index]['Option2 Name'] = ''
        new_excelify.loc[index]['Option2 Value'] = ''
        new_excelify.loc[index]['Option3 Name'] = ''
        new_excelify.loc[index]['Option3 Value'] = ''
        new_excelify.loc[index]['Variant Generate From Options'] = 'FALSE'
        new_excelify.loc[index]['Variant Position'] = '1'
        new_excelify.loc[index]['Variant SKU'] = row['']
        new_excelify.loc[index]['Variant Weight'] = ''
        new_excelify.loc[index]['Variant Weight Unit'] = ''
        new_excelify.loc[index]['Variant HS Code'] = ''
        new_excelify.loc[index]['Variant Country of Origin'] = 'JP'
        new_excelify.loc[index]['Variant Compare At Price'] = ''
        new_excelify.loc[index]['Variant Price'] = get_price_number(row['    '])
        new_excelify.loc[index]['Variant Cost'] = ''
        new_excelify.loc[index]['Variant Requires Shipping'] = 'TRUE'
        new_excelify.loc[index]['Variant Taxable'] = 'TRUE'
        new_excelify.loc[index]['Variant Tax Code'] = ''
        new_excelify.loc[index]['Variant Barcode'] = ''
        new_excelify.loc[index]['Variant Image'] = get_image_path_if_not_NaN(row['  -      (1)_2'])
        new_excelify.loc[index]['Variant Inventory Tracker'] = 'shopify'
        new_excelify.loc[index]['Variant Inventory Policy'] = 'deny'
        new_excelify.loc[index]['Variant Fulfillment Service'] = 'manual'
        new_excelify.loc[index]['Variant Inventory Adjust'] = ''
        new_excelify.loc[index]['Variant Inventory Qty'] = row['   _2']
    if (is_not_NaN(row['  -      (2)_2'])):
        index_inline += 1
        new_excelify.loc[index + index_inline] = ''
        new_excelify.loc[index + index_inline]['Handle'] = row['']
        new_excelify.loc[index + index_inline]['Command'] = 'UPDATE'
        new_excelify.loc[index + index_inline]['Image Src'] = get_image_path_if_not_NaN(row['  -      (2)_2'])
        new_excelify.loc[index + index_inline]['Image Command'] = 'MERGE'
    if (is_not_NaN(row['  -      (3)_2'])):
        index_inline += 1
        new_excelify.loc[index + index_inline] = ''
        new_excelify.loc[index + index_inline]['Handle'] = row['']
        new_excelify.loc[index + index_inline]['Command'] = 'UPDATE'
        new_excelify.loc[index + index_inline]['Image Src'] = get_image_path_if_not_NaN(row['  -      (3)_2'])
        new_excelify.loc[index + index_inline]['Image Command'] = 'MERGE'
    if (is_not_NaN(row['  -      (4)_2'])):
        index_inline += 1
        new_excelify.loc[index + index_inline] = ''
        new_excelify.loc[index + index_inline]['Handle'] = row['']
        new_excelify.loc[index + index_inline]['Command'] = 'UPDATE'
        new_excelify.loc[index + index_inline]['Image Src'] = get_image_path_if_not_NaN(row['  -      (4)_2'])
        new_excelify.loc[index + index_inline]['Image Command'] = 'MERGE'
    if (is_not_NaN(row['  -      (5)_2'])):
        index_inline += 1
        new_excelify.loc[index + index_inline] = ''
        new_excelify.loc[index + index_inline]['Handle'] = row['']
        new_excelify.loc[index + index_inline]['Command'] = 'UPDATE'
        new_excelify.loc[index + index_inline]['Image Src'] = get_image_path_if_not_NaN(row['  -      (5)_2'])
        new_excelify.loc[index + index_inline]['Image Command'] = 'MERGE'
```

```py
new_excelify.to_csv ('xcelify_formatted.csv', index = False, header=True)
```

```py
df.sort(['a', 'b'], ascending=[True, False])
```
