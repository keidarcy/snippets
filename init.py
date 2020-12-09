import os

files = os.listdir('./')
folders = []
for file in files:
  if not '.' in file:
    folders.append(file)

links = {}
for folder in folders:
  file_list = os.listdir('./' + folder) 
  links[folder] = file_list

print(links)

print('--------------OO-----')


with open("test.txt",'w',encoding = 'utf-8') as f:
   f.write("my first file\n")
   f.write("This file\n\n")
   f.write("contains three lines\n")