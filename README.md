# gogad
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
![Go Version](https://img.shields.io/github/go-mod/go-version/zhiwei-Feng/gogad)

Available GPUs Detection tool

> 用于解决平时查看各个服务器上可用的GPU的重复工作

## Get Started
### 准备服务器配置文件
按照`machines-template.csv`的内容创建自己可用的服务器的配置，名称自拟，但需要符合csv文件格式，用逗号隔开
### 运行
#### go run的方式
1. `git clone https://github.com/zhiwei-Feng/gogad.git`
2. `go run gogad.go <path-of-csv>`

#### exe文件 的方式
下载latest release的exe文件直接运行即可
`./gogad.exe <path-of-csv>`
   

