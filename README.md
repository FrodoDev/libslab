[toc]

# libslab

try to use variety libs

# list of libs
1. log
    1. zap
    2. xxx
2. proxy





主要试验这几个方面

1. zap 自身使用，核心是参数对结果的影响
2. 生产环境中，如果是写入磁盘的日志，日志切分必不可少
3. 横向比较几个 Go 日志库的性能
4. 分析性能好的原因

## zap 的输出格式，各项配置的不同对应输出结果

默认配置

```go
cfg := zap.NewProductionConfig()
log, err := cfg.Build()
```

输出示例说明

```
// 默认配置 cfg := zap.NewProductionConfig()
{"level":"info","ts":1755160634.812118,"caller":"lzap/lab.go:37","msg":"Lab int","foo":"bar"}
// cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder

// 直接打印在标准输出有颜色，不过zap编码的时候改成 Unicode 编码，就没颜色了
{"level":"\u001b[34minfo\u001b[0m","ts":1755323458.7120051,"caller":"lzap/lab.go:36","msg":"Lab int","foo":"bar"}
```
