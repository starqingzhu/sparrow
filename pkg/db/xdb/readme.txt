README.md: SQLX 的简要介绍。
LICENSE.txt: SQLX 的许可证信息。
benchmark: 包含对 SQLX 进行基准测试的代码和数据文件。
examples: 包含使用 SQLX 的示例代码。
sqlx.go: SQLX 的核心代码，包含了 DB 和 Tx 结构体以及所有的公共方法。
reflectx: 提供了对反射类型的支持，用来处理数据库中的数据和结构体之间的映射关系。
types: 提供了一些类型的定义，例如 NullInt64 和 NullString，用来处理数据库中的 Null 值。
另外，SQLX 还有一些子模块，例如 mysql、postgres、sqlite3 等，分别用于支持不同的数据库类型。你可以进入对应的子模块目录中查看该子模块的源代码。

总的来说，SQLX 的源代码结构清晰，注释详细，对于想要学习和使用 SQLX 的人来说，十分友好。