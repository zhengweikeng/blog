> 近日看到一个相对于grep更加强大的命令----awk，在需要对数据进行格式化输出时，使用它会更加方便。

###### 简要说明
awk其名称得自于它的创始人 Alfred Aho 、Peter Weinberger 和 Brian Kernighan 姓氏的首个字母。  
实际上awk的拥有自己的语言：AWK 程序设计语言，三位创建者已将它正式定义为“样式扫描和处理语言”。  
它允许您创建简短的程序，这些程序读取输入文件、为数据排序、处理数据、对输入执行计算以及生成报表，还有无数其他的功能。

###### 使用方式
`awk '{pattern + action}'`  
或者  
`awk 'pattern {action}'`  
其中 pattern 表示 AWK 在数据中查找的内容，而 action 是在找到匹配内容时所执行的一系列命令。花括号 ({}) 不需要在程序中始终出现，但它们用于根据特定的模式对一系列指令进行分组

###### 调用方式
1.命令行方式
awk [-F  field-separator]  'commands'  input-file(s)
其中，commands 是真正awk命令，[-F域分隔符]是可选的。 input-file(s) 是待处理的文件。
在awk中，文件的每一行中，由域分隔符分开的每一项称为一个域。通常，在不指名-F域分隔符的情况下，默认的域分隔符是空格。

2.shell脚本方式
将所有的awk命令插入一个文件，并使awk程序可执行，然后awk命令解释器作为脚本的首行，一遍通过键入脚本名称来调用。
相当于shell脚本首行的：#!/bin/sh
可以换成：#!/bin/awk

3.将所有的awk命令插入一个单独文件，然后调用：
awk -f awk-script-file input-file(s)
其中，-f选项加载awk-script-file中的awk脚本，input-file(s)跟上面的是一样的。

###### 简单实例
1.输出前五行，只显示登陆账号
```
last -n 5 | awk '{print $1}'
``` 
awk工作流程是这样的：读入有'\n'换行符分割的一条记录，然后将记录按指定的域分隔符划分域，填充域，$0则表示所有域,$1表示第一个域,$n表示第n个域。默认域分隔符是"空白键" 或 "[tab]键",所以$1表示登录用户，$3表示登录用户ip,以此类推。

2./etc/passwd账户，使用:分割
```
cat /etc/passed | awk -F ':' '{print $1}'
# 以tab隔离
cat /etc/passwd | awk -F ':' '{print $1"\t"print $7}'
```

3.加入前后header
```
cat /etc/passwd | awk -F ':' 'BEGIN  {print "name,shell"} {print $1","$7} END {print "the end"}'
```

前面几个都是awk+action的例子，以下展示awk+pattern的例子

4.搜索/etc/passwd有root关键字的所有行
```
awk -F ':' '/root/' /etc/passwd
```
这种是pattern的使用示例，匹配了pattern(这里是root)的行才会执行action(没有指定action，默认输出每行的内容)
只显示对应的shell
```
awk -F ':' '/root/{print $7}' /etc/passwd
```
该例子则为awk+pattern+action的示范

###### awk内置变量
ARGC               命令行参数个数
ARGV               命令行参数排列
ENVIRON            支持队列中系统环境变量的使用
FILENAME           awk浏览的文件名
FNR                浏览文件的记录数
FS                 设置输入域分隔符，等价于命令行 -F选项
NF                 浏览记录的域的个数
NR                 已读的记录数
OFS                输出域分隔符
ORS                输出记录分隔符
RS                 控制记录分隔符

统计/etc/passwd:文件名，每行的行号，每行的列数，对应的完整行内容
```
awk -F ':' '{print "filename:"FILENAME",linenumber:"NR",columns:"NF",linecontent:"$0}' /etc/passwd
#or
awk -F ':' '{printf("filename:%10s,linenumber:%s,columns:%s,linecontent:%s",FILENAME,NR,NF,$0)}' /etc/passwd
```
关于print和printf的用法，参照[此文](http://lnwayne.duapp.com/awkzhong-guan-yu-printfde-xiang-guan-yong-fa/)  
更多关于awk的用法，请访问[该处](http://www.gnu.org/software/gawk/manual/gawk.html)
