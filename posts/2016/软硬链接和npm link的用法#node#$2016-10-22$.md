## Linux中的Hard Link和Symbolic Link
为了做文件的共享，一般我们会使用硬链接或者软链接（符号链接）来实现。

### 硬链接
通过硬链接创建的文件都具有相同的索引号。看例子

```bash
# 创建文件
touch source_file
# 建立硬链接
ln source_file target_file

# 查看文件列表，并带有inode
ls -li
25262335 -rw-r--r--  2 seed  staff    0 10 22 23:07 source_file
25262335 -rw-r--r--  2 seed  staff    0 10 22 23:07 target_file

# 对源文件写入
echo "hello world" >> source_file

cat source_file
hello world

cat target_file
hello world

# 对目标文件写入
echo "say hi" >> target_file

cat target_file
hello world
say hi

cat source_file
hello world
say hi

# 删除源文件
rm source_file
ls -li
25262335 -rw-r--r--  2 seed  staff    0 10 22 23:07 target_file

cat target_file
hello world
say hi
```

显然，硬链接的索引号是一样的（25262335）。而且无论对哪一个文件进行写入，都会反应在另外一个文件。  
而且无论是删除哪一个文件，其他文件都不会造成影响。

注意的是，硬链接不能作用在目录。  
硬链接不能对目录创建是受限于文件系统的设计。  
现 Linux 文件系统中的目录均隐藏了两个个特殊的目录：当前目录（.）与父目录（..）。查看这两个特殊目录

---

### 软链接（符号链接）
软链接则只是简单的将路径指向了原始文件，因此软链接产生的文件都具有不同的索引号。软链接也因此支持目录

#### 文件软链接

```bash
# 创建源文件
touch source_file
# 建立软链接
ln -s source_file target_file

# 查看文件列表，并带有inode
ls -li
25262824 -rw-r--r--  1 seed  staff    0 10 22 23:41 source_file
25262834 lrwxr-xr-x  1 seed  staff   11 10 22 23:41 target_file -> source_file

# 对源文件写入
echo "hello world" >> source_file

cat source_file
hello world

cat target_file
hello world

# 对目标文件写入
echo "say hi" >> target_file

cat target_file
hello world
say hi

cat source_file
hello world
say hi

# 删除源文件
rm source_file
cat target_file
# 目标文件已经无法访问
cat: target_file: No such file or directory
```
大部分和硬链接是一样的，不同的就是删除源文件，目标文件也无法访问了，因为路径指向的源文件已经不在了。

#### 目录软链接

```bash
# 创建源目录
mkdir source_dir
# 建立软链接
ln -s source_dir target_dir

# 查看文件列表，并带有inode
ls -li
25262953 drwxr-xr-x  2 seed  staff  68 10 22 23:47 source_dir
25262955 lrwxr-xr-x  1 seed  staff  10 10 22 23:47 target_dir -> source_dir

# 在源目录下创建文件
touch source_dir/file1

ls source_dir
file1

ls target_file
file1

# 在目标目录创建一个文件
touch target_dir/file2

ls source_dir
file1 file2

ls target_file
file1 file2

# 对源目录下的一个文件写入
echo "hello world" >> source_dir/file1

cat source_dir/file1
hello world

cat target_dir/file1
hello world

# 对目标目录下的一个文件写入
echo "say hi" >> target_dir/file1

cat source_dir/file1
hello world
say hi

cat target_dir/file1
hello world
say hi
```

---

## npm link
npm link借助了软链接的思想，它可以让我们将系统中某个目录映射到项目的node_modules下。

有时我们的项目依赖一个我们正在开发的模块，我们可以通过这种方式，将对模块的修改映射到我们项目下，提高开发效率。

```bash
# 有个叫redis的模块，放在projects/node-redis目录下
cd ~/projects/node-redis    
# 将该目录加入到全局link下
npm link                    
# 项目的目录
cd ~/projects/node-bloggy   
# 通过npm link安装模块
npm link redis              
```

需要注意的是，npm link后面接的不是目录名，而是包的名字。