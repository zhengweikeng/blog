> 从去年年初开始便将编辑器由sublime换成了vscode，觉得非常好用，故做一些使用记录

**注意：以下所有操作均在mac环境下执行**

# 常用功能和快捷键
vscode的界面如下所示：
![userinterface_hero](./images/userinterface_hero.png)
## 布局和侧边栏

打开和关闭侧边栏(side bar)：`cmd + b`

在编辑区（edit group）打开一个新的栏：`cmd + \`

不在编辑区的当前激活栏打开文件，而是在第二个栏打开文件，有如下几种方式：  
1. 按住cmd，之后点击文件
1. `cmd + p`，选择文件后，`cmd + enter` 

切换编辑区的栏目: `cmd + 1|2|3`

为了尽可能的使用屏幕，一般我们都会将活动栏和侧边栏收起来，通过快捷键快速调出  

其中vscode提供了一种Zen Mode，进入该模式会隐藏所有栏目，并进入全屏，让我们充分利用屏幕  
进入Zen Mode：`cmd + k, z`
侧边栏显示资源管理器：`cmd + shift + e`  
侧边栏显示搜索区：`cmd + shift + f`  
侧边栏显示版本管理区：`ctrl + shift + g`  
侧边栏显示调试区：`ctrl + shift + d`  
侧边栏显示扩展区：`ctrl + shift + x`

默认情况下编辑区中多个栏目是水平方向的，我们也可以变成垂直方向：`alt + cmd + 1`

关闭编辑区文件：`cmd + w`  
关闭所有在编辑区打开的文件：`cmd + k + w`

## 命令面板
快速选择和打开文件：`cmd + p`，输入要打开的文件名的几个字母

显示最近打开的几个文件，并选择和打开它：`ctrl + shift + tab`，通过tab切换

输入一些vscode内置或者插件提供的命令：`cmd + shift + p`

跳转到当前鼠标所在文件中的某个函数或者变量的定义中：`cmd + shift + o`

跳转到指定行数：`ctrl + g`

## 基本功能
删除一行：`cmd + shift + k`

向后插入一行空行：`cmd + enter`  
向前插入一行空行：`cmd + shift + enter`

向下移动一行：`alt + ↓`  
向上移动一行：`alt + ↑`

拷贝当前行到下一行：`alt + shift + ↓`  
拷贝当前行到上一行：`alt + shift + ↑`  

选中当前单词，并逐步选中后续相同的单词：`cmd + d`  
选择所有出现的当前选择：`cmd + shift + l`

逐步撤销鼠标的操作：`cmd + u`






# 主题
vscode默认提供了一些主题，使用`cmd + k` + `cmd + t`选择主题

另外vscode也提供了icon的主题，使用`cmd + shift + p`，输入icon，选择文件图标主题

更多的主题到[vscode官网](https://marketplace.visualstudio.com/VSCode)下载

# 配置
vscode有许多配置，通过`cmd + ,`查看所有默认配置项

同时vscode提供了两种方式修改这些配置项：  
1. user setting，即用户修改的配置只要打开vscode就生效
1. workspace setting，即只在当前项目下剩下，它会在当前目录下生成一个`.vscode/setting.json`的文件，在文件中修改配置

修改配置无需去复制原始值，鼠标靠近配置项会出现一个“✏️”，点击会出现该配置项支持的修改值，点击修改值即可将配置添加到user setting或者workspace setting中。  
到user setting或者workspace setting也可以通过这种方式修改配置。  
这种方式解决了你不知道配置项不知道有哪些备选值得问题。
