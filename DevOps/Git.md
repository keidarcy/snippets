- [Basic](#basic)
  - [settings](#settings)
      - [basic settings](#basic-settings)
      - [help messages](#help-messages)
- [creating snapshots](#creating-snapshots)
  - [status](#status)
  - [diff, difftool](#diff-difftool)
  - [show, ls-tree, restore, clean](#show-ls-tree-restore-clean)
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