import glob
import re
import os

mdDics = {}
folders = []
for folder in os.listdir('./'):
    if os.path.isdir(folder) and not '.' in folder:
        folders.append(folder)

for filepath in glob.iglob('./' + '**/*.md', recursive=True):
    if not 'node_modules' in filepath and not 'README' in filepath:
        mdDics[os.path.basename(filepath).replace(
            '.md', '').lower()] = filepath

mdStr = """
# Snippets

run this before push new commit

```bash
python3 init.py
```

---

"""

for folder in folders:
    mdStr += ("- {0}\n".format(folder))
    for name, path in mdDics.items():
        if folder in path:
            fileStr = """  - [{0}]({1})\n""".format(name, path)
            mdStr += (fileStr)

mdStr += """
---

### My External Links

- [Xpath](http://xpather.com/)
- [Figma](https://www.figma.com/file/GAMKg6zWYqYId04ICOHOPq/funny?node-id=1%3A2)

"""
with open("README.md", 'w', encoding='utf-8') as f:
    f.write(mdStr)

print(mdStr)

print('-------------NEW README LINK GENERATE SUCCESSFULLY------------------')
