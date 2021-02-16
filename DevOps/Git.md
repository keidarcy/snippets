### Some useages
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