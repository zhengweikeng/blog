linux下从压缩文件的扩展名可以看出它提供了很多的压缩技术，如*.tar, *.tar.gz, *.tgz, *.gz, *.Z, *.bz2。

其中*.Z是由compress技术提供，这个技术已经过时，也很少见到有人使用，因此便不浪费精力介绍。

### gzip
使用gzip建立的压缩文件的格式为：*.gz。并且使用gzip还能解开comress,zip,gzip等软件压缩过的文件。

gzip的压缩比会比comress的好的多。
```bash
gzip [-cdtv#] 文件名
-c --stdout          write to stdout, keep original files
   --to-stdout
-d --decompress      uncompress files
   --uncompress
-t --test            test compressed file
-v --verbose         print extra statistics
-# set compression level, default 6

# 例子
gzip -v test.conf
```
上述例子会生成一个test.conf.gz的文件，并且源文件不存在了。如果需要保留源文件，可以如下做法。

```bash
gzip -cv test.conf > test.conf.gz
```
gzip解压缩命令如下所示：
```bash
gzip -d test.conf.gz
```

另外可以使用zcat读取压缩包的内容，而不需要解压缩出来
```bash
zcat test.conf.gz
```

### bzip2
使用bzip2压缩后的文件名为：*.bz2。用法与zip是一致的。只是它提供了更好的压缩比。
```bash
bzip2 -z test.conf
```
同理，会生成一个test.conf.bz2的压缩文件，并且源文件会消失。

若需要保留源文件，则如下：
```bash
bzip2 -c test.conf > test.conf.bz2
```

bzip2可以解压缩.bz,.bz2,.tbz,tbz2这些压缩文件。
```bash
bzip2 -d test.conf.bz2
```

也可以查看压缩文件内容，而不解压缩出来
```bash
bzcat test.conf.bz2
```

### tar
linux下的tar提供了上述的所有功能，既可以压缩和解压缩gz，也可以是bz2。根据配置不同的参数决定压缩和解压缩的方式。

先看看压缩
```bash
# 压缩成bz2，参数为-j
tar -jcv -f test.conf.tar.bz2 test.conf

# 压缩成gz，参数为-c
tar -zcv -f test.conf.tar.gz test.conf
```
使用tar时需要我们手动定义压缩文件名，我们发现尾缀都带有tar，这是为了让人们看出我们使用的是tar的压缩。

如果不加-j或者-z的话，尾缀可以是*.tar。

再来看看解压缩
```bash
# 解压缩bz2文件
tar -jxv -f test.conf.tar.bz2

# 解压缩gz文件
tar -zxv -f test.conf.tar.gz
```

同理，tar也可以查看压缩文件内容，而不解压。
```bash
tar -jtv -f test.conf
```
