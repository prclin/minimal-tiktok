# minimal tiktok
字节跳动第六届青训营大项目，极简版抖音
## 1. 目录结构
config: 配置结构体包

controller: 控制器包

core: 核心包，存放项目启动所需初始化项

dao: 数据访问包

global：全局包，存放全局需要用到的变量或函数

model: 模型包，存放实体，请求，响应模型

service: 服务包，包含业务逻辑

util: 工具包，包含与业务无关的工具

application.yaml: 配置文件

main.go: 入口文件
## 2. 提交规范
正式开发时，将仓库pull下来后，不允许在master分支修改、提交、push，按如下规范操作:
1. git add . (如果新建了文件则跟踪文件，以免下一步操作丢失新增的文件)
2. git stash (将本次修改存入stash栈内，操作之后分支会回到最近一次提交的状态)
3. git checkout master (切换到master分支)
4. git pull (获取远程最新更新)
5. git checkout dev (切换到自己的开发分支)
6. git stash pop (恢复暂存)
7. git commit -m "" (提交)
8. git rebase master (变基并处理冲突)
9. git push origin dev (只能提交到自己的远程开发分支)
10. 发起pull request
> 其中1-7步都是为了预先处理冲突而为
> 
> 其中第4，如果你在master分支作了修改，可以直接使用以下操作代替以产生不必要的merge:
>
> git fetch origin master (将远程分支所有新提交获取到本地)
> 
> git rebase origin/master (以远程master分支为基分支执行rebase，有冲突就处理)