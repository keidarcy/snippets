
## Version Control System

- [Version Control System](#version-control-system)
	- [concepts](#concepts)
	- [settings](#settings)
		- [basic settings](#basic-settings)
		- [help messages](#help-messages)
		- [Push commit to different remote branch](#push-commit-to-different-remote-branch)
		- [Dry run push](#dry-run-push)
		- [Delete branch](#delete-branch)
		- [Update forked repository to original repository latest](#update-forked-repository-to-original-repository-latest)
		- [Clean current branch](#clean-current-branch)


### concepts

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
#### basic settings

```bash
git config --global user.name "xyh"
git config --global user.email ""
git config --global core.editor ""
git config --global -e # open global settings
git config --global core.autocrlf "" # windows - true, mac - input
```

#### help messages

```bash
git config --help # details
git config -h # brief
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