
 - Delete branch

```bash
// delete branch locally
git branch -d localBranchName

// delete branch remotely
git push origin --delete remoteBranchName
```

- Update forked repository to original repository latest

```
git remote add upstream https://github.com/original/repository.git

git fetch upstream

git rebase upstream/master

git push origin master --force
```

- Clean current branch

```
nah='git reset --hard;git clean -df;'
```
