# int-to-string
对纯数字混淆加密, 用于想要自增字段的性能, 又不想公开自增id

# 使用 

````go
    // 生成长度>=5的code
	server := New(5)
	code := server.ToCode(1234567890)
	if server.ToId(code) != 1234567890 {
		t.Error("加密解密不一致")
	}
	
	// 生成新的加密模版
    str := New(5).BuildConfig()
    server := New(5, str)
````