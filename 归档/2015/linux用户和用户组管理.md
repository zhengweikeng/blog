##### 与用户管理相关文件
1. /etc/passwd
   User account information  
   该账户不记录用户密码，密码用x代替
2. /etc/shadow
   Secure user account information  
   文件内容和1类似，不过该文件记录了密码（密文）
3. /etc/group
   Group account information  
   文件内容中“x”代表无密码
4. /etc/gshadow
   Secure group account information  
   文件内容和1类似，不过该文件记录了密码（密码）
5. /etc/default/useradd
   Default values for accout creation  
   该文件记录了创建用户时不加参数情况下参数的默认值。  
   该文件有一行是`SKEL=/etc/skel`,指的便是用户home目录下存在的文件，这些文件即`/etc/skel/`目录下的文件
6. /etc/skel/
   Directory containing default files
   5中有解释
7. /etc/login.defs
   System Management Commands  
   设置用户帐号限制的文件，在这里我们可配置密码的最大过期天数，密码的最大长度约束等内容。该文件里的配置对root用户无效。如果/etc/shadow文件里有相同的选项，则以/etc/shadow里的设置为准，也就是说/etc/shadow的配置优先级高于/etc/login.defs 

##### 用户管理
1. 添加用户
 `useradd [options] username`  
 常用选项  
 -e 用户账户的有效日期，格式YYYY-MM-DD  
 -g 用户所属主用户组  
 -G 用户附加用户组  
 -s 用户默认shell  
 -u:直接给出userID
2. 删除用户
 `userdel [options] username`  
 常用选项  
 -r 连同用户的家目录一同删除
3. 查看用户信息
  `id username`和`groups username`
4. 修改用户信息 
 `usermod [options] LOGIN`  
 常用选项  
 -a 将用户添加到新的附加组，需要和-G一起使用`usermod -aG group1 testuser`  
 -G 修改用户附加组  
 -g 修改用户所属主组
 -l 修改用户名  
 其他选项和添加用户时相同
5. 用户口令
 `passwd [options] username`  
 常用选项  
 -l 锁定口令，即禁用账号。  
 -u 口令解锁。  
 -d 使账号无口令。  
 -f 强迫用户下次登录时修改口令。  
 usermod 命令禁用和启用账号通过在 /etc/shadow 中相应用户密码位之前添加和删除 "!" 实现的

##### 用户组管理
1. 添加组
 `groupadd [options] groupname`  
 常用选项
 -g 指定组id
 -p 组密码
2. 删除组
 `groupdel groupname`
3. 修改组
 `groupmod [options] groupname`
4. 查看组
 `cat /etc/group`

##### gpasswd
gpasswd命令是Linux下工作组文件/etc/group和/etc/gshadow管理工具。
`gpasswd [options] groupname`
-a：添加用户到组；  
-d：从组删除用户； 
-A：指定管理员；  
-M：指定组成员和-A的用途差不多；  
-r：删除密码；  
-R：限制用户登入组，只有组中的成员才可以用newgrp加入该组。
```
 gpasswd -A peter group1
 # 这样peter就成为组管理员，可进行如下操作
 gpasswd -a user1 group1
 gpasswd -a user2 group1

 # 将用户假如sudoer
 gpasswd -a user3 wheel
```
注意：添加用户到某一个组 可以使用`usermod -G group_name user_name`这个命令可以添加一个用户到指定的组，但是以前添加的组就会清空掉。  
所以想要添加一个用户到一个组，同时保留以前添加的组时，请使用`gpasswd -a`这个命令来添加操作用户

