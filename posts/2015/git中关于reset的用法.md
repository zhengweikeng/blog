> 用于回滚错误的提交    

`git reset [--mixed | --soft | --hard | --merge | --keep] [-q] [<commit>或HEAD]`  
一般在`git add`之后控制台会打印如下提示： 
```
On branch master
Your branch is up-to-date with 'origin/master'.

Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

	modified:   README.md
``` 
我们可以使用git reset回滚操作，将当前的分支重设（reset）到指定的<commit>或者HEAD（默认，如果不显示指定commit，默认是HEAD，即最新的一次提交），并且根据参数的不同有可能更新暂存区和工作区。 
 
1.--hard
该参数会将暂存区和工作区的修改全部丢弃，HEAD指向commit
```
commit3: add test3.c
commit2: add test2.c
commit1: add test1.c
```  
假设最新的提交为commit3，此时执行`git reset --hard HEAD^1`，控制台显示`HEAD is now at dd2d2f1 commit2`，说明已经成功会滚到commit2的状态  

2.--soft
该参数会将指定的commit撤销，此时状态处于待提交状态，可使用`git commit`进行提交。使用soft只会更改引用的指向，不改变暂存区和工作区
```
git reset --soft HEAD~1
# 将会发现暂存区和工作区的文件均为最后一次提交时的样子
git status
-> On branch master
   Your branch is up-to-date with 'origin/master'.

   Changes to be committed:
      (use "git reset HEAD <file>..." to unstage)

	     modified:   README.md
git commit -m "commit3"
```  

3.--mixed或者不使用参数（默认为--mixed）
该参数会改变引用的指向，并将暂存区丢弃，但是不改变工作区
```
git reset --mixed HEAD~1
# OR
git reset HEAD~1
->  Unstaged changes after reset:
    M	README.md
```
