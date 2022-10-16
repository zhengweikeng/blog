* `git checkout  [-q][<commit>] [--] <path>`
*  `git checkout [<branch>]`
* `git checkout [-m] [[-b|--orphan] <new_branch>] [start_point]`  

1.第一种用法
省略commit，则会用暂存区的文件覆盖工作区的文件；否则用指定提交的文件覆盖工作区中对应的文件
```
vim README.md  (编辑后为 “aaa”)
git add README.md
vim README.md (编辑后为 "aaa bbb")
git checkout -- README.md
# 此时README.md为 "aaa"
# 相当于取消上次执行git add以来（如果执行过）的本地修改
git commit -m "commit"
vim README.md (编辑后为 "aaa ccc")
git checkout -- README.md
# 此时README.md 仍为 "aaa"
```  
2.第二种用法
切换到一个分支 `git checkout dev`

3.第三种用法
用于创建并切换到该分支，新的分支从<start-point>指定的提交开始创建  
`git checkout -b dev`

##### 获取远程分支
```
# 获取远程分支
git branch -r
# 假设打印出如下信息
  origin/HEAD -> origin/master
  origin/dev
  origin/master
# 查看本地分支
git branch
# 假设只有master，若要获取远程最新的dev分支
git checkout -b dev origin/dev  # 分支dev不存在，所以需要加上-b创建分支
```

##### 更新远程分支
```
# 未来分支会更新，需要先从远程更新远程分支信息

# 拉取远程所有分支的更新
git fetch <远程主机名>
# 拉取远程指定分支的更新
git fetch <远程主机名> <分支名>
# 所取回的更新，在本地主机上要用”远程主机名/分支名”的形式读取

# 本地分支不存在，则参考《获取远程分支》方式，checkout分支
git checkout -b dev origin/dev

# 本地合并远程分支
git merge origin/master
# or
git rebase origin/master
```
