# errno

业务域错误可以统一通过 proto 定义业务原因，然后通过 protoc 生成枚举定义和相关 helper functions。

在 errno 包中，错误模型主要跟 http 状态码一致，并且实现相关接口将 Error 与 GRPCStatus 互转。

errno 中的 Error 本身不携带堆栈信息，如果想要在 error 中携带堆栈信息，可以在业务代码中借助 github.com/pkg/errors 实现。
被 errors 库包装过的 Error 依然可以使用内置的 errno.IsXXX 等大错误类型判方法。

### 依赖项

安装新版 protoc-gen-go ，protoc-gen-go-grpc