


# git 常用的命令

> 可视化给工具 推荐使用：  `sourcetree` 非常方便切换提交版本历史记录 、
> <https://www.sourcetreeapp.com/>


## 项目初始中

### 1.本地创建上次创库

```
git init

```

### 2.远程创库创建项目

```git

git clone 地址
```




## 远程直接覆盖本地

```sh
git fetch --all  #拉去远端最新
git reset --hard origin/master  # 把HEAD指向刚刚下载的最新的版本 
```



## 创建版本



## 创建分析


## 分支合并




```sh

// 本地git 代码管理

管理文件，初始化
$ git init

查看当前文件夹的状态
$ git status  
对指定文件进行版本控制
git add 文件
git add .  对指定文件下的所有文件及子目录进行版本控制
创建提交记录（版本）
git commit -m ‘详细的描述信息’

推送到远程git
git push -u origin master

查看log
git log
查看所有的log
git reflog 

回归到某一个版本
git reset --hard 版本号

删除本地缓存
git rm -r --cached.
重新添加，上传


帮助我们暂时存储已经开发了一些新功能的代码，继续开发
缓存
零时新功能，移走修改过没有提交的
git stash 
git stash pop 合并第一个缓存
git stash list 缓存过的所有的记录
git stash clear  清除
git stash apply  合并哪个一个
git stash drop 清除哪一个

123.com

git branch
创建分支 dev
git branch dev

删除
git branch -d 分支 
切换分支
git checkout 分支

git mast dev 合并2
```


git commit -m "Change repo." # 先把所有为保存的修改打包为一个commit
git remote remove origin # 删掉原来git源
git remote add origin [YOUR NEW .GIT URL] # 将新源地址写入本地版本库配置文件
git push -u origin master # 提交所有代码





git remote add origin git@github.com:YouAge/vue_uniapp_shoppp.git
git push -u origin master