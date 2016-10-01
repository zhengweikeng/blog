# 安装入门
## virtualbox
需要先安装[virtualbox](https://www.virtualbox.org/)

安装完后，可以选择性的安装它的扩展，因为在后续步骤中，你安装的box可能是需要virtualbox装有该插件

[插件下载页](https://www.virtualbox.org/wiki/Downloads)  
安装  
VirtualBox 5.1.6 Oracle VM VirtualBox Extension Pack

下载后将其添加到virtualbox即可

Mac的话，在`偏好设置/扩展`这里添加

## 安装vagrant
[vagrant](https://www.vagrantup.com/downloads.html)

如果已经安装过需要升级的，也是直接下载安装，安装时会直接覆盖旧的文件

## 生成Vagrantfile
选择一个目录，作为工程的根目录，在该目录生成Vagrantfile

```bash
vagrant init
```

## 安装Box
### 方法1
Box可以认为就是一个环境镜像，这个环境可以是ubuntu，可以是centos等等。

因此根据自己的需求，安装不同的Box。

可以到 [HashiCorp's Atlas box catalog](https://atlas.hashicorp.com/boxes/search?utf8=%E2%9C%93&sort=&provider=&q=)这里查找需要的Box

```
vagrant box add ubuntu/trusty64
```

### 方法2
方法1在国内的话，在终端下载镜像会非常慢

所以我们一般不会采取这种方式安装。

而是直接去网上搜索镜像，利用迅雷将box下载下来，再进行安装。

可以到这里搜索box  
[vagrantbox.es](http://www.vagrantbox.es/)

这里我们下载了debian-8.1-lxc-puppet.box，接下来往vagrant添加box

```
# 查看box add的帮助文档
vagrant box add --help

# 文档提示：
# vagrant box add [options] <name, url, or path>
# 即可以接具体的文件路径，name是添加到vagrant后的命名，便于管理

vagrant box add debian /Users/seed/Public/box/debian-8.1-lxc-puppet.box

vagrant box list
# 显示，说明已经添加进来
debian      (virtualbox, 0)
```

做完上述操作后，需要修改Vagrantfile。在该文件中找到下面的文字

```vagrantfile
Vagrant.configure("2") do |config|
  config.vm.box = "base"
end
```
将`config.vm.box = "base"`  
修改为  
`config.vm.box = "debian"`

## 启动和进入

```
vagrant up
# 一切正常后
vagrant ssh

cd /vagrant
ls
```

在/vagrant目录下可以看到，我们的Vagrantfile所在目录的所有文件，即它被共享到/vagrant这个目录了

`CTRL+D`可以退出回到宿主机器

# 其他配置
## 修改同步目录
默认情况下，vagrant会将Vagrantfile所在目录的文件同步到/vagrant目录下，我们可以修改这个目录

修改Vagrantfile文件中的`config.vm.synced_folder`

```vagrantfile
config.vm.synced_folder "." "/seed"
```
之后

```
vagrant reload
vagrant ssh
cd /seed
ls
```
便可以看到同步的目录了

更多关于同步的，参考官网[文档](https://www.vagrantup.com/docs/synced-folders/)

## PROVISIONING
有时我们会在多台机器使用vagrant安装一些linux环境，同时安装一些工具，如apache

如果每台机器安装完环境后，都需要逐个去安装工具的话就太费事了

vagrant提供了PROVISIONING机制，在vagrant安装完环境后自动安装其他的工具

首先编写一个安装脚本bootstrap.sh

```shell
#!/usr/bin/env bash

apt-get update
apt-get install -y apache2
if ! [ -L /var/www ]; then
  rm -rf /var/www
  ln -fs /vagrant /var/www
fi

```

修改Vagrantfile文件

```vagrantfile
Vagrant.configure("2") do |config|
  config.vm.box = "debian"
  config.vm.provision :shell, path: "bootstrap.sh"
end
```

这次在vagrant启动的时候，我们指定了脚本bootstrap.sh

```
vagrant up
```

这时在启动完后便会去执行脚本

如果已经执行过`vagrant up`操作启动了虚拟机，可以使用reload命令来再次执行脚本

```vagrantfile
vagrant reload --provision
```

更多关于PROVISIONING的，参考官网[文档](https://www.vagrantup.com/docs/provisioning/)

## 网络
有时我们经常需要使用网络的方式访问我们的虚拟机

如我们在虚拟机中测试部署了一个网站，监听80端口，我们想在宿主机器中访问它。

修改Vagrantfile

```vagrantfile
Vagrant.configure("2") do |config|
  config.vm.network "forwarded_port", guest: 80, host: 8080
  config.vm.network "private_network", ip: "192.168.33.10"
  # config.vm.network "public_network"
end
```

我们可以通过如下两种方式访问到

http://localhost:8080  
http://192.168.33.10

基本上配置这两个已经足够满足开发人员的需求了。

如果需要让其他机器也能访问，则需要将网络设置为public_network

详情可参考[文档](https://www.vagrantup.com/docs/networking/)

## 关闭虚拟机
vagrant可以让我们根据不同的情况关闭虚拟机

### vagrant suspend
执行后将会保存当前虚拟机的状态，简单点理解就是将系统睡眠。

启动时依旧执行`vagrant up`

这种情况下的启动会很快，并且会恢复到suspend之前的系统状态。

这种方式的缺点就是，suspend之后虚拟机依旧会占有宿主机器的内存和硬盘空间。

### vagrant halt
执行后会完全关闭虚拟机

启动时依旧执行`vagrant up`

这种方式的好处就是关闭后不占据宿主机器的内存和硬盘空间。

缺点就是启动的需要会花费一些时间

### vagrant destroy
这种方式和vagrant halt很像，但是区别就是它还会释放所有资源，宿主机器恢复到一个干净没有任何虚拟机资源的环境下。

缺点之一是在启动的时候也需要花费一些时间。  
缺点之二就是需要重新re-provision一遍，就是重新运行之前我们说的脚本。  
因为destroy后的状态已经是运行脚本前的状态了