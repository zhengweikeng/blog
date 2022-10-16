> grep (global search regular expression(RE) and print out the line,全面搜索正则表达式并把行打印出来)是一种强大的文本搜索工具，它能使用正则表达式搜索文本，并把匹配的行打印出来。

###### 常见用法
grep [-acinv] [--color] '搜索字符串' filename
选项与参数  
-a ：将 binary 文件以 text 文件的方式搜寻数据  
-c ：计算找到 '搜寻字符串' 的次数  
-i ：忽略大小写   
-n ：输出行号  
-v ：反向选择，亦即显示出没有 '搜寻字符串' 内容的那一行  
--color=auto ：将找到的内容高亮显示

###### 事例  
1.将/etc/passwd，有出现 root 的行取出来
```
grep root /etc/passwd
# or
cat /etc/passwd |　grep root
```  
2.加上行号
```
grep -n root /etc/passwd
```
3.多条件的情况
```
# 或运算
grep -E "root|nologin" /etc/passwd
# or
cat /etc/passwd | grep -E "root|nologin"

# 与运算
grep nologin /etc/passwd | grep man
# or
cat /etc/passwd | grep nologin | grep man
```
4.输出匹配关键字的前2行后3行
```
grep nologin -A3 -B2 /etc/passwd
```
5.根据文件内容递归查找目录
```
#在当前目录搜索带'container'行的文件
grep container *

#在当前目录及其子目录下搜索'energywise'行的文件
grep -r container *

#在当前目录及其子目录下搜索'energywise'行的文件，但是不显示匹配的行，只显示匹配的文件
greo -r -l container *
```

###### grep与正则表达式
字符类  
1.test和taste均含有t?st
```
grep 't[ea]st' test.sh
```
2.反向选择
```
grep '[^a]st' test.sh
# or
grep '[^a-z]st' test.sh
```
无论[]中有多少个字符，都只会代表一个字符，即它只会搜索test和tast而不会搜素teast  
连续的字符或者数字因为ASCII上编码的顺序是连续的，可以用-来简化使用，[a-z][A-Z][0-9]  

3.行首和行尾匹配
```
# 以te开头
grep '^te' test.sh

# 以e结尾
grep 'e$' test.sh

# 以.结尾
grep '\.$' test.sh
```
由此可知：^ 符号，在字符类符号(括号[])之内与之外是不同的！ 在 [] 内代表『反向选择』，在 [] 之外则代表定位在行首的意义。另外像.这种特殊字符，需要\转义

4.任意一个字节 . 与重复字节 *
. (小数点)：代表『一定有一个任意字节』的意思；  
* (星号)：代表『重复前一个字符， 0 到无穷多次』的意思，为组合形态
```
# 匹配类似tees的值
grep 't..s' test.sh

# 匹配类似teeeeeeees的值
grep 'te*s' test.sh
```
5.限定连续字符范围{}
```
# 2个o
grep 'o\{2\}' test.sh

# 2到5个字符
grep 'o\{2,5\}' test.sh
```

5.搜索所有包含一个或多个3的行
```
# egrep '3+' testfile
# grep -E '3+' testfile
# grep '3\+' testfile 
```

### 数据流导向
一般情况下，我们的每一条命令都是输出到屏幕上，我们也可以通过流导向将数据输出到指定的东西。
1. 标准输入(stdin)：代码为0，使用<或者<<
2. 标准输出(stdout)：代码为1，使用>或者>>
3. 标准错误输出(stderr)：代码为2，使用2>或者2>>

```bash
# 将结果输出到profile文件。此时文件内容是全部被清空后重新写入
ls / > ./profile

# 不同于上方，这可以实现内容追加，而不是清空重写
ls / >> ./profile
```

我们还可以一条命令的错误信息导向文件，将正确信息导向另外的一个文件
```bash
find /home --name .bashrc > list_right 2> list_err
```
这里假设一个用户没有权限查看其他用户的home

#### /etc/null
我们可以将错误信息直接忽略掉，不输出也不写入文件。这时便可以使用垃圾桶黑洞装置
```bash
find /home --name .bashrc > list_right 2> /etc/null
```

#### 将正确的和错误的都输入同一个文件
```bash
# 错误
find /home --name .bashrc > list 2> list

# 正确
find /home --name .bashrc > list 2> &1

# 正确
find /home --name .bashrc &> list
```

### cut
cut -d '分割字符' -f fields
```bash
test=t1:t2:t3:t4

echo $test | cut -d ':' -f 2
# t2
```

cut -c 字符区间
```bash
test=helloworld

echo $test | cut -c 6
# w

echo $test | cut -c 6-
# world

echo $test | cut -c 6-8
# wor
```
