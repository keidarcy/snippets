import string
import random
import math
import json
import pandas as pd

sample = pd.read_csv('./csvs/excelify_sample.csv')
product_list0 = pd.read_csv('./csvs/product_list.csv')
product_csv0 = pd.read_csv('./csvs/product_csv.csv')

product_list = pd.DataFrame(product_list0[1:])
product_list.index = list(product_list0.iloc[1:, 1])
product_list.columns = list(product_list0.iloc[0, :])
product_list = product_list.sort_values(
    by="", ascending=True, na_position='first')
product_list = product_list.groupby(product_list.index).first()

product_csv = pd.DataFrame(product_csv0)
product_csv.index = product_csv0.loc[:, '']
product_csv.index.names = ['index']
product_csv = product_csv.sort_values(
    by="", ascending=True, na_position='first')
product_csv = product_csv.groupby(product_csv.index).first()

product_data = product_list.join(product_csv.add_suffix('_2'), how='outer')
product_data.dropna(subset=[''], inplace=True)
product_data.shape


def get_boolean(x): return x == '○'


def is_not_NaN(num):
    return num == num


def add_if_not_none(list_val):
    tags = []
    for val in list_val:
        if(is_not_NaN(val)):
            tags.append(val)
    return tags


def get_image_path_if_not_NaN(x): return 'PATH' + x if(is_not_NaN(x)) else ''


def get_price_number(price):
    return price.replace('¥', '').replace(',', '')


def get_value_if_string(x): return x if(type(x) == str) else ''


def get_random_string(length):
    letters = string.ascii_lowercase
    result_str = ''.join(random.choice(letters) for i in range(length))
    return result_str


newColumns = sample.columns
newColumnsList = list(newColumns)
newColumnsList.remove('Metafield: specs.range [integer]')
newColumnsList.remove('Variant Metafield: something [string]')
newColumnsList.append('Metafield: information [string]')
new_excelify = pd.DataFrame(columns=newColumnsList)
new_excelify

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
        new_excelify.loc[index]['Title'] = row['']
        new_excelify.loc[index]['Body HTML'] = get_value_if_string(
            row['']) + '\n' + get_value_if_string(row[''])
        new_excelify.loc[index]['Vendor'] = row['']
        new_excelify.loc[index]['Type'] = row['']
        new_excelify.loc[index]['Tags'] = ','.join(add_if_not_none([row['']]))
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
        new_excelify.loc[index]['Image Src'] = get_image_path_if_not_NaN(
            row[''])
        new_excelify.loc[index]['Image Command'] = 'MERGE'
        new_excelify.loc[index]['Image Position'] = '1'
        new_excelify.loc[index]['Image Width'] = ''
        new_excelify.loc[index]['Image Height'] = ''
        new_excelify.loc[index]['Image Alt Text'] = row['']
        new_excelify.loc[index]['Variant ID'] = ''
        new_excelify.loc[index]['Variant Command'] = 'MERGE'
        if is_not_NaN(row['']):
            new_excelify.loc[index]['Option1 Name'] = 'color'
            new_excelify.loc[index]['Option1 Value'] = get_random_string(
                2) + '.' + row['']
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
        new_excelify.loc[index]['Variant Price'] = get_price_number(row[''])
        new_excelify.loc[index]['Variant Cost'] = ''
        new_excelify.loc[index]['Variant Requires Shipping'] = 'TRUE'
        new_excelify.loc[index]['Variant Taxable'] = 'TRUE'
        new_excelify.loc[index]['Variant Tax Code'] = ''
        new_excelify.loc[index]['Variant Barcode'] = ''
        new_excelify.loc[index]['Variant Image'] = get_image_path_if_not_NaN(
            row[""])
        new_excelify.loc[index]['Variant Inventory Tracker'] = 'shopify'
        new_excelify.loc[index]['Variant Inventory Policy'] = 'deny'
        new_excelify.loc[index]['Variant Fulfillment Service'] = 'manual'
        new_excelify.loc[index]['Variant Inventory Adjust'] = ''
        new_excelify.loc[index]['Variant Inventory Qty'] = row['']
        new_excelify.loc[index]['Metafield: description_tag'] = row['']
        new_excelify.loc[index]['Metafield: information [string]'] = row['']
        new_excelify.loc[index]['Metafield: title_tag'] = row['']
        metafiledDic = {
            'textable': get_boolean(row['']),
            'miniMark': get_boolean(row['']),
            'sisyuNormal': get_boolean(row['']),
            'sisyuDisney': get_boolean(row['']),
            'sisyuStarwars': get_boolean(row['']),
            'giftbox': get_boolean(row['']),
            'giftboxSet': get_boolean(row['']),
            'wrapping': get_boolean(row['']),
            'giftbox': get_boolean(row['']),
            'yahoo': get_boolean(row['']),
            'color': get_boolean(row[''])
        }
        new_excelify.loc[index]['Metafield: custom.json [json_string]'] = json.dumps(
            metafiledDic)

    else:
        new_excelify.loc[index]['Handle'] = row['']
        new_excelify.loc[index]['Command'] = 'UPDATE'
        new_excelify.loc[index]['Vendor'] = row['']
        new_excelify.loc[index]['Type'] = row['']
        new_excelify.loc[index]['Tags'] = ','.join(add_if_not_none(
            [row[''], row[''], row[''], row[''], row[''], row[''], row[''], row[''], row[''], row[''], row[''], row[''], row[''], row['']]))
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
        new_excelify.loc[index]['Image Src'] = get_image_path_if_not_NaN(
            row[''])
        new_excelify.loc[index]['Image Command'] = 'MERGE'
        new_excelify.loc[index]['Image Position'] = '1'
        new_excelify.loc[index]['Image Width'] = ''
        new_excelify.loc[index]['Image Height'] = ''
        new_excelify.loc[index]['Image Alt Text'] = row['']
        new_excelify.loc[index]['Variant ID'] = ''
        new_excelify.loc[index]['Variant Command'] = 'MERGE'
        if is_not_NaN(row['']):
            new_excelify.loc[index]['Option1 Name'] = 'color'
            new_excelify.loc[index]['Option1 Value'] = get_random_string(
                2) + '.' + row['']
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
        new_excelify.loc[index]['Variant Price'] = get_price_number(row[''])
        new_excelify.loc[index]['Variant Cost'] = ''
        new_excelify.loc[index]['Variant Requires Shipping'] = 'TRUE'
        new_excelify.loc[index]['Variant Taxable'] = 'TRUE'
        new_excelify.loc[index]['Variant Tax Code'] = ''
        new_excelify.loc[index]['Variant Barcode'] = ''
        new_excelify.loc[index]['Variant Image'] = get_image_path_if_not_NaN(
            row[''])
        new_excelify.loc[index]['Variant Inventory Tracker'] = 'shopify'
        new_excelify.loc[index]['Variant Inventory Policy'] = 'deny'
        new_excelify.loc[index]['Variant Fulfillment Service'] = 'manual'
        new_excelify.loc[index]['Variant Inventory Adjust'] = ''
        new_excelify.loc[index]['Variant Inventory Qty'] = row['']
    if (is_not_NaN(row[''])):
        index_inline += 1
        new_excelify.loc[index + index_inline] = ''
        new_excelify.loc[index + index_inline]['Handle'] = row['']
        new_excelify.loc[index + index_inline]['Command'] = 'UPDATE'
        new_excelify.loc[index +
                         index_inline]['Image Src'] = get_image_path_if_not_NaN(row[''])
        new_excelify.loc[index + index_inline]['Image Command'] = 'MERGE'
    if (is_not_NaN(row[''])):
        index_inline += 1
        new_excelify.loc[index + index_inline] = ''
        new_excelify.loc[index + index_inline]['Handle'] = row['']
        new_excelify.loc[index + index_inline]['Command'] = 'UPDATE'
        new_excelify.loc[index +
                         index_inline]['Image Src'] = get_image_path_if_not_NaN(row[''])
        new_excelify.loc[index + index_inline]['Image Command'] = 'MERGE'
    if (is_not_NaN(row[''])):
        index_inline += 1
        new_excelify.loc[index + index_inline] = ''
        new_excelify.loc[index + index_inline]['Handle'] = row['']
        new_excelify.loc[index + index_inline]['Command'] = 'UPDATE'
        new_excelify.loc[index +
                         index_inline]['Image Src'] = get_image_path_if_not_NaN(row[''])
        new_excelify.loc[index + index_inline]['Image Command'] = 'MERGE'
    if (is_not_NaN(row[''])):
        index_inline += 1
        new_excelify.loc[index + index_inline] = ''
        new_excelify.loc[index + index_inline]['Handle'] = row['']
        new_excelify.loc[index + index_inline]['Command'] = 'UPDATE'
        new_excelify.loc[index +
                         index_inline]['Image Src'] = get_image_path_if_not_NaN(row[''])
        new_excelify.loc[index + index_inline]['Image Command'] = 'MERGE'

new_excelify.to_csv('excelify_formatted.csv', index=False, header=True)
