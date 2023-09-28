# int-to-string
对纯数字混淆加密, 用于想要自增字段的性能, 又不想公开自增id

# 使用 

强烈推荐, 每个项目都先生成独立的配置后再使用。防止不同项目生成了同一个数字

````go
// 生成一份配置
server := New(8)
config := server.BuildConfig() // 需要把这个配置保存起来

// 下次启动要用同一份配置

// 重新启动, 加载生成的配置
server = New(8, config)
for i := 1; i < 1000000000; i++ {
    code := server.ToCode(i)
    if server.ToId(code) != i {
        t.Errorf("加密前(%v) 加密后(%v) 解析后(%v)", i, code, server.ToId(code))
    }
}
````