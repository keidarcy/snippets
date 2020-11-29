## Jupyter notebook

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
df = pd.DataFrame(np.arange(12).reshape((4,3)),columns=['A', 'B', 'C'],index=["1行目", '2行目', '3行目', '4行目'])
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
df.loc['1行目', :]
df.loc[['1行目', '2行目'], :]
```

```py
# iloc method use index of index and columns
df.iloc[1,1]
df.iloc[1:,1]
df.iloc[1:, :2]
```

```py
sample = pd.read_csv('./csvs/excelify_sample.csv')
product_list0 = pd.read_csv('./csvs/product_list.csv')
product_csv0 = pd.read_csv('./csvs/product_csv.csv')

product_list = pd.DataFrame(product_list0[1:])
product_list.index = list(product_list0.iloc[1:, 1])
product_list.columns = list(product_list0.iloc[0, :])
product_list.sort_values(by="商品コード", ascending=False)
product_list

product_csv = pd.DataFrame(product_csv0)

product_csv.index = product_csv0.loc[:,'商品コード']


product_csv.index.names = ['index']
product_csv.sort_values(by="商品コード")
product_csv

pro = product_list.join(product_csv.add_suffix('_2'), how='outer')
print(product_list.shape)
print(product_csv.shape)
# pro.to_csv ('new_excelify_sunday.csv', index = False, header=True)
pro.shape
```
