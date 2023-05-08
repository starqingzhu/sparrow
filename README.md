# sparrow
项目基础建设

在 Go 语言中，通常会遵循以下的目录规范来组织项目结构：  
cmd：该目录用于存放主要的应用程序入口文件，相当于 main 包。   
pkg：该目录用于存放可导出的库代码，这些代码可以供其他项目使用。   
internal：该目录用于存放不可导出的库代码，这些代码只能在当前项目中使用。   
api：该目录用于存放与外部服务交互的代码，如 REST 和 gRPC 接口等。   
config：该目录用于存放配置文件，如 YAML、JSON 和 TOML 等格式的配置文件。  
test：该目录用于存放测试代码。  
vendor：该目录用于存放第三方库的依赖项。   
此外，还可以在根目录下添加 README 文件、LICENSE 文件和 docs 目录用于存放文档。最终项目的目录结构如下所示：

├── cmd  
│   ├── main.go  
│   └── other.go  
├── pkg  
│   ├── lib1  
│   │   ├── lib1.go  
│   │   └── lib_test.go  
│   └── lib2  
│       ├── lib2.go  
│       └── lib_test.go  
├── internal  
│   ├── api1  
│   ├── api2  
│   └── some  
│       └── module  
├── api  
├── config  
├── test  
├── vendor  
├── README.md  
├── LICENSE  
└── docs  

需要注意的是，具体的目录结构可能因项目类型、功能和个人习惯等原因而有所不同，在组织目录结构时需要具体问题具体分析，以适应项目的需求。