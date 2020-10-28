# prototool 
proto 辅助工具，根据配置文件的依赖描述，编译，格式化proto，还可以对proto进行拼写检查

# 安装
```shell script
go get -u github.com/dotamixer/doom/tool/doom-protoc
```

# 命令
请查看help

```shell script
prototool -h
Usage:
  prototool [command]

Available Commands:
  config      生成配置文件
  fmt         格式化proto
  generate    编译proto
  help        Help about any command
  lint        对proto进行拼写检查

Flags:
  -h, --help   help for prototool

Use "prototool [command] --help" for more information about a command.
```

# 依赖

格式化功能依赖clang-format;请下在clang-format 后将clang-format 放到PATH 环境变量下；

# 参考
###  同类工具

[uber prototool](https://github.com/uber)

### 拼写检查手册

[Google API Linter](https://linter.aip.dev/)

### llvm 

[clang-format 下载](https://github.com/llvm/llvm-project)

[clang-format 说明](https://clang.llvm.org/docs/ClangFormat.html)
