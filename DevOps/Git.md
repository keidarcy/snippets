- [Basic](#basic)
  - [settings](#settings)
      - [basic settings](#basic-settings)
      - [help messages](#help-messages)
- [creating snapshots](#creating-snapshots)
  - [status](#status)
  - [diff, difftool](#diff-difftool)
  - [show, ls-tree, restore, clean](#show-ls-tree-restore-clean)
- [Browsing history](#browsing-history)
  - [`--state`, `--patch`](#--state---patch)
  - [filter commits](#filter-commits)
  - [format git log](#format-git-log)
  - [viewing a commit](#viewing-a-commit)
  - [viewing the changes between commits](#viewing-the-changes-between-commits)
  - [checking out a commit](#checking-out-a-commit)
  - [finding bugs using bisect](#finding-bugs-using-bisect)
  - [finding contributors using shortlog](#finding-contributors-using-shortlog)
  - [viewing history of a file](#viewing-history-of-a-file)
  - [restoring a deleting file](#restoring-a-deleting-file)
  - [finding the author of line using blame](#finding-the-author-of-line-using-blame)
  - [tagging](#tagging)
    - [Push commit to different remote branch](#push-commit-to-different-remote-branch)
    - [Dry run push](#dry-run-push)
    - [Delete branch](#delete-branch)
    - [Update forked repository to original repository latest](#update-forked-repository-to-original-repository-latest)
    - [Clean current branch](#clean-current-branch)


## Basic

- why vcs
  - track history
  - work together
- types
  - centralized (subversion, team foundation server )
  - distributed (git, mercurial)
- why git
  - free
  - open source
  - super fast
  - scalable
  - cheap branching and merging

### settings

- levers
  - system
  - global
  - local
##### basic settings

```bash
git config --global user.name "xyh"
git config --global user.email ""
git config --global core.editor ""
git config --global -e # open global settings
git config --global core.autocrlf "" # windows - true, mac - input
```

##### help messages

```bash
git config --help # details
git config -h # brief
```

## creating snapshots

- staging area(last snapshot version)

once first commit be created staging area exists.

```bash
git add file1 # add file1 to staging area
```

```bash
git commit -m 'first commit' # every commit git store fill content not diff
```

- add and commit to staging at same time

```bash
git commit -ma 'Refactor code.'
```

- remove from staging area

```bash
git rm --cached -r bin/
```

### status

```bash
git status -sb
```

### diff, difftool

```bash
git diff # the diff of working directory and staging area
git diff --stages   # the diff of staging area and last commit
```

### show, ls-tree, restore, clean

- `show` will show git objects
  - commits
  - blobs(files)
  - trees(directories)
  - tags

```bash
git show HEAD # HEAD is last commit
git show HEAD~1 # HEAD~n is nth commit before last commit
git show HEAD~1:.gitignore # show the whole file
```

```bash
git ls-tree HEAD~1 # show all
```

- replace `reset` with `restore`
- `restore` move the version to the last version

```bash
git restore --staged file2.js
```

```bash
git clean -fd
```

## Browsing history

### `--state`, `--patch`

```bash
git log --oneline --stat # the diff line numbers
git log --oneline --patch # the diff detail
```

### filter commits

```bash
git log --oneline -3 # last 3 commits
git log --oneline --author="XING YAHAO"
git log --oneline --after="2020-08-17"
git log --oneline --after="yesterday"
git log --oneline --after="one week ago"
git log --oneline --grep="commit   lomessage term"
git log --oneline -S "git" # search all commit include the search contents and show commit
git log --oneline -S "git" --patch # show the diff it self
git log --oneline fb0d..dad47 # show between range of two commits
git log --oneline netlify.toml # all the commits that modified specific file
git log --oneline -- netlify.toml # file name is ambiguous
git log --oneline --patch -- netlify.toml # show commit content
```

### format git log

```bash
git log --pretty=format:"%an committed %h on %cd" # %an -> author name, %h -> hash value, %cd -> committed date
git log --pretty=format:"%Cgreen%an%Creset committed %h on %cd" # %Cgreen %Creset  change color
```

### viewing a commit

```bash
git show HEAD~2
git show HEAD~2:packages/react-reconciler/src/ReactFiberCommitWork.new.js # see the final version of the commit
git show HEAD~2 --name-only
git show HEAD~2 --name-status
```

### viewing the changes between commits

```bash
git diff HEAD~2 HEAD # find the all differences between two commits
git diff HEAD~2 HEAD fixtures/concurrent/time-slicing/src/index.js # find the differences in specific file
git diff HEAD~2 HEAD --name-only
git diff HEAD~2 HEAD --name-status
```
### checking out a commit

```bash
git checkout c0a77029c # You are in 'detached HEAD' state
git log --oneline -all
git checkout master
```

### finding bugs using bisect

```bash
git bisect bad
git bisect good c0a77029c
git bisect good
git bisect bad
git bisect reset
```

### finding contributors using shortlog

```bash
git shortlog -n -s -e  # author, commits number, email
git shortlog -nse --after="" --before=""
```
### viewing history of a file

```bash
git log --oneline --stat --patch .gitignore
```

### restoring a deleting file

```bash
git rm toc.txt
git commit -m "Removed toc.txt"
git log --oneline -- toc.txt
git checkout a642e12 toc.txt
git commit -m "Restored toc.txt"
```

### finding the author of line using blame

```bash
git blame .gitignore
git blame -e .gitignore
git blame -e -L 1,3 .gitignore
```
### tagging

```bash
git tag v1.0 {hash} // lightweight tag
git checkout v1.0
git tag -a v1.1 -m "My version 1.1"  // annotate tag
git tag -n # show tag message
```










#### Push commit to different remote branch

```bash
git push origin local-name:remote-name
```

#### Dry run push

```sh
git push -nu origin xxx
# -n --dry-run
# -u --set-upstream
```

#### Delete branch

```bash
# delete branch locally
git branch -d localBranchName

# delete branch remotely
git push origin --delete remoteBranchName
```

#### Update forked repository to original repository latest

```bash
git remote add upstream https://github.com/original/repository.git

git fetch upstream

git rebase upstream/master

git push origin master --force
```

#### Clean current branch

```bash
alias nah='git reset --hard;git clean -df;'
```