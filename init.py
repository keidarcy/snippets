import os

files = os.listdir('./')
folders = []
for file in files:
  if not '.' in file:
    folders.append(file)

for folder in folders:
  readme_list = os.listdir('./' + folder) 
  print(readme_list)
