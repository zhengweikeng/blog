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
grep nologin /etc/passwd | greo man
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
