### 变量的基础
定义一个变量很简单
```bash
var1=test
var2="test it"
```

删除一个变量,使用unset
```bash
unset var1
```

将变量转化为环境变量
```bash
export var1
```

所谓环境变量，即为可以在任何shell中使用的变量。一般的变量无法子shell中使用。

### 变量的删除与替换
#### 变量的删除
linux下的变量很容易删除与取代。直接看例子
```bash
path=/usr/kerberos/sbin:/usr/kerberos/bin:/usr/local/sbin:/usr/local/bin:/sbin:/bin: /usr/sbin:/usr/bin:/root/bin

# 我想将含有kerberos的两个目录删除
path=${path#/*kerberos/bin:}
echo $path
# /usr/local/sbin:/usr/local/bin:/sbin:/bin: /usr/sbin:/usr/bin:/root/bin
```
上述例子，使用了${variable#/*kerberos/bin:}这种格式。

其中*是匹配0到无穷个字符。案例中它会从变量开始向右删除匹配/开头，中间间隔任何字符，最后是kerberos/bin:，因此这里匹配的内容便是：  
/usr/kerberos/sbin:/usr/kerberos/bin:

如果有多个匹配的情况呢，则可以使用##。  
\#：符合取代文字[最短的]一个  
\##：符合取代文字[最长的]一个
```bash
path=/usr/kerberos/sbin:/usr/kerberos/bin:/usr/local/sbin:/usr/local/bin:/sbin:/bin: /usr/sbin:/usr/bin:/root/bin

path=${path##/*:}
echo $path
# /root/bin
```

使用#是从前面开始向后删除，而使用%则可以实现从后向前删除，用法和#是一致的。
```bash
path=/usr/kerberos/sbin:/usr/kerberos/bin:/usr/local/sbin:/usr/local/bin:/sbin:/bin: /usr/sbin:/usr/bin:/root/bin

# 删除最后一个bin目录
path1=${path%*bin}
echo $path1
# path=/usr/kerberos/sbin:/usr/kerberos/bin:/usr/local/sbin:/usr/local/bin:/sbin:/bin: /usr/sbin:/usr/bin:

# 删除到最开始的bin目录
path2=${path%%:*bin}
echo $path2
# /usr/kerberos/sbin:
```

#### 变量的替换
格式如下：  
${变量/旧字符串/新字符串}  替换第一个符合的旧字符串  
${变量//旧字符串/新字符串} 替换全部符合的旧字符串
```bash
path=/usr/kerberos/sbin:/usr/kerberos/bin:/usr/local/sbin:/usr/local/bin:/sbin:/bin: /usr/sbin:/usr/bin:/root/bin

path1=${path/sbin/SBIN}
echo $path1
# /usr/kerberos/SBIN:/usr/kerberos/bin:/usr/local/sbin:/usr/local/bin:/sbin:/bin: /usr/sbin:/usr/bin:/root/bin

path2=${path//sbin/SBIN}
echo $path2
# /usr/kerberos/SBIN:/usr/kerberos/bin:/usr/local/SBIN:/usr/local/bin:/SBIN:/bin: /usr/SBIN:/usr/bin:/root/bin
```

### 变量值的决定
可以通过变量值的情况决定该变量的值

![变量值的决定](https://raw.githubusercontent.com/zhengweikeng/blog/master/posts/2016/images/%E5%8F%98%E9%87%8F%E5%80%BC%E7%9A%84%E5%86%B3%E5%AE%9A.jpg)
```bash
unset user
var=${user-root}
# root

user=''
var=${user-root}
# 此时var是空串

user=test
var=${user-root}
# test

user=''
var=${user:-root}
# root
```
