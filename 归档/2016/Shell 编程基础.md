### $和$()和${}的区别
当定义一个变量后，即可用$变量名，取得该变量的值。若需要获取一个命令后的结果作为值，可以采用$()，而${}在一些连接字符串时会经常用到
```bash
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

### 声明一个变量的类型 declare
declare [-aixr] variable
选项与参数：
-a：将变量定义成为array类型  
-i：将变量定义成为integer类型  
-x：用法与export一样，就是将变量变成环境变量；   
-r：将变量配置成为readonly类型，该变量不可被更改内容，也不能unset
```bash
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

### 数组
定义数组，一对括号表示是数组，数组元素用“空格”符号分割开。
`a=(1 2 3 4)`或者逐项设置`a[1]=1`  
读取数组长度  
`echo ${#a[@]}`或者`echo ${#a[*]}`  
读取各项  
`echo ${a[2]}`  
下标是*或者@则得到整个数组  
清除数组
```
# 清除整个数组
unset a
# 清除单个数组项
unset a[1]
```

### 判断
shell有两种方式进行判断，不仅能做变量之间的比较（如空与非空，相等，大于，小于，不等于等等），还可以判断文件或者目录是否存在。  
不仅如此还能做文件或者目录权限的判断。

第一种方式，采取test的方式。
```bash
read -p "username:" username
test -z $username && echo "username should not empty!" && exit 0
```
这里做了一个变量的非空判断。语法便是：`test -primaries expression`

再来看一个例子，用于判断文件是否存在。
```bash
echo your name is ${username}

read -p "filename:" filename
test ! -e $filename && echo The file ${filename} is not exist || echo The file ${filename} is exist
```
当然，test还有更多配置项，如-o，即为o语句，-a，即为and语句。具体可以`man test`

第二种方式，[]的方式
```bash
echo your name is ${username}
[ ! -e $filename ] && echo The file ${filename} is not exist || echo The file ${filename} is exist
```
上述例子也是判断文件是否存在，这里要注意的便是，[]里的两端都要留出空格，不然会报错。

再看一个例子
```bash
param1="test what"
param2="what"
# [ $param1 == $param2 ] && echo equal || echo not equal
[ "$param1" == "$param2" ] && echo equal || echo not equal
```
可以发现上面如果注释的地方如果让它运行是会报错的，因为会解析后会变成`[ test what == what ]`，这是不符合语法规定的。  
因此，在变量和常量的比较里，最好还是用引号包裹起来。

[]这种方式会经常用在if这类的条件语句中。

最后看一个or的例子
```bash
read -p "input (Y/N):" yn
[ "$yn" == "Y" -o "$yn" == "y" ] && echo "OK" && exit 0
[ "$yn" == "N" -o "$yn" == "n" ] && echo "wrong!" && exit 0
```

### 条件语句
```
if ...; then
   ...
elif ...;then
   ...
else
   ...
fi
```
一般情况下，我们都会使用上述测试命令，即[]中写表达式的方式来写条件语句，不过我们一般不会在[]里面写多个表达式，而是会放在外部，写多个[]的方式。
```bash
#!/bin/bash
read -p "please input (y/n):" yn

if [ "$yn" = "y" ] || [ "$yn" = "Y" ]; then
   echo "ok!continue"
   exit 0
elif [ "$yn" = "n" ] || [ "$yn" = "N" ]; then
   echo "oh!interrupt"
   exit 0
eles
   echo "I don't know what you are input"
   exit 0
fi
```  

### case语句
case表达式可以用来匹配一个给定的字符串，而不是数字
```bash
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

### select语句
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

### while和for循环
第一种循环方式：while do done, until do done (不定循环)
```bash
while [ "$yn" != "yes" ] && [ "$yn" != "YES" ]
do
  read -p "Please input yes/YES to stop this program: " yn
done

echo "go next"

until [ "$yn" != "yes" ] && [ "$yn" != "YES" ]
do 
  read -p "Please input yes/YES to stop this program: " yn
done
```
第二种循环： for...do...done (固定循环)
形态1：
```bash
# exp1
for animal in dog cat elephant
do
  echo "The annimal is $animal"
done

# exp2
network="192.168.19.1"
for site in $(seq 1 100)
do
  ping -c 1 -w ${network}.${site} &> /etc/null && result=0 || result=1

  if [ "$result" == 0 ]; then
    echo "Server ${network}.${site} is UP"
  else
    echo "Server ${network}.${site} is DOWN"
  fi
done
```
形态2
```bash
for (( 初始值; 限制值; 运行步阶 ))
do
    程序段
done

s=0
for ((i=1; i<=100; i=i+1))
do
  echo $(( $s+ $i ))
done
```

### 一些特殊符号
$# 是传给脚本的参数个数  
$0 是脚本本身的名字  
$1 是传递给该shell脚本的第一个参数  
$2 是传递给该shell脚本的第二个参数  
$@ 是传给脚本的所有参数的列表   
$* 是以一个单字符串显示所有向脚本传递的参数，与位置变量不同，参数可超过9个  
$$ 是脚本运行的当前进程ID号  
$? 是显示最后命令的退出状态，0表示没有错误，其他表示有错误

有如下脚本文件，test.sh
```bash
#!/bin/bash

echo "The script name is ==> $0"
echo "Total parameter number is ==> $#"
echo "Your whole parameter is ==> $@"

echo "The 1st param ===> $1"
echo "The 2nd param ===> $2"

shift
echo "Your whole parameter is ==> $@"

shift 3
echo "Your whole parameter is ==> $@"
```
运行`./test.sh aa bb cc dd ee ff gg`  
结果如下：
```
The script name is ==> ./test.sh
Total parameter number is ==> 7
Your whole parameter is ==> aa bb cc dd ee ff gg
The 1st param ===> aa
The 2nd param ===> bb
Your whole parameter is ==> bb cc dd ee ff gg
Your whole parameter is ==> ee ff gg
```
注意到这里有个shift的命令，它可以拿掉最前面的n个参数

### 函数
shell也有函数的概念
```bash
function printIt(){
   echo "print it: $1"
}
```
函数也是有内建变量，只不过是相对于函数的
