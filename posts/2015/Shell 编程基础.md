###### $和$()和${}的区别
当定义一个变量后，即可用$变量名，取得该变量的值。若需要获取一个命令后的结果作为值，可以采用$()，而${}在一些连接字符串时会经常用到
```
#变量名和变量值不能存在空格
name=tom
echo $name

#若变量值存在空格，应该使用单引号或者双引号
who="my name is $name"
#双引号会解释特殊字符的含义
echo $who
who='my name is $name'
#单引号不会解释特殊字符的含义
echo $who

#需要连接字符时，最好双引号或者采用${}
PATH="$PATH":/home/seed/test
PATH=${PATH}:/home/seed/test

#当一串命令需要借助其他命令时，可以采用反引号或者$()
version=`uname -r`
version=$(uname -r)
```

###### 声明一个变量的类型 declare
declare [-aixr] variable
选项与参数：
-a：将变量定义成为array类型  
-i：将变量定义成为integer类型  
-x：用法与export一样，就是将变量变成环境变量；   
-r：将变量配置成为readonly类型，该变量不可被更改内容，也不能unset
```
sum=100+200+300
echo $sum   # 100+200+300
declare -i sum=100+200+300
echo $sum   # 600
```

数值之间的计算，还可以使用$(())
```bash
num1=3
num2=4
res=$(($num1*$num2))
res=$(( $num1 * $num2 ))
res=$((num1*num2))
```
三种方式都会打印出12。推荐使用$(())这种方式，方便记忆。

###### 数组
定义数组，一对括号表示是数组，数组元素用“空格”符号分割开。
`a=(1 2 3 4)`或者逐项设置`a[0]=1`  
读取数组长度  
`echo ${#a[@]}`或者`echo ${#a[*]}`  
读取各项  
`echo ${a[0]}`  
下标是*或者@则得到整个数组  
清除数组
```
# 清除整个数组
unset a
# 清除单个数组项
unset a[0]
```

###### 条件语句
```
if ...; then
   ...
elif ...;then
   ...
else
   ...
fi
```
大多数情况下，可以使用测试命令来对条件进行测试，比如可以比较字符串、判断文件是否存在及是否可读等等……通常用" [ ] "来表示条件测试，注意这里的空格很重要，要确保方括号前后的空格。  
[ -f "somefile" ] ：判断是否是一个文件  
[ -x "/bin/ls" ] ：判断/bin/ls是否存在并有可执行权限  
[ -n "$var" ] ：判断$var变量是否有值  
[ "$a" = "$b" ] ：判断$a和$b是否相等(注意只有一个等号)
```
#!/bin/bash
read -p "please input (y/n):" yn

if [ "$yn" = "y" ] || [ "$yn" = "Y" ]; then
  echo "ok!continue"
  exit 0
elif [ "$yn" = "n" ] || [ "$yn" = "N" ]; then
  echo "oh!interrupt"
  exit 0
fi

echo "I don't know what you are input"
exit 0
```  

###### case语句
case表达式可以用来匹配一个给定的字符串，而不是数字
```
case ... in
   ...
   ;;
esac

case $1 in
  "hello")
    echo "Hello, how are you ?"
    ;;
  "")
    echo "You MUST input parameters, ex> {$0 someword}"
    ;;
  *)   # 其实就相当於万用字节，0~无穷多个任意字节之意！
    echo "Usage $0 {hello}"
    ;;
esac
# $1表示传入的第一个参数
```

###### select语句
一般用于交互
```
select var in ... ; do
  break;
done
  ...

echo "What is your favourite OS?"
select var in "Linux" "Gnu Hurd" "Free BSD" "Other"; do
  break;
done
echo "You have selected $var"

该脚本的运行结果如下：
What is your favourite OS?
1) Linux
2) Gnu Hurd
3) Free BSD
4) Other
#? 1
You have selected Linux
```

###### while和for循环
```
while ...; do
  ...
done

for ... in ...; do
   ...
done

for (( 初始值; 限制值; 运行步阶 ))
do
    程序段
done

for rpmpackage in "$@"; do
   if [ -r "$rpmpackage" ];then
      echo "=============== $rpmpackage =============="
      rpm -qi -p $rpmpackage
   else
      echo "ERROR: cannot read file $rpmpackage"
   fi
done
# $@表示你传入的所有参数

# zsh和dash会报错，建议使用上一种
for (( i=1; i<=$nu; i=i+1 ))
do
    s=$(($s+$i))
done
```

###### 一些特殊符号
$# 是传给脚本的参数个数  
$0 是脚本本身的名字  
$1 是传递给该shell脚本的第一个参数  
$2 是传递给该shell脚本的第二个参数  
$@ 是传给脚本的所有参数的列表   
$* 是以一个单字符串显示所有向脚本传递的参数，与位置变量不同，参数可超过9个  
$$ 是脚本运行的当前进程ID号  
$? 是显示最后命令的退出状态，0表示没有错误，其他表示有错误
