package int_to_string

import (
	"testing"
)

func TestCodeService_ToUserCode(t *testing.T) {
	server := New(4)
	code := server.ToCode(1234567890)
	if server.ToId(code) != 1234567890 {
		t.Error("加密解密不一致")
	}
}

func TestCodeService_ToUserCodeAll(t *testing.T) {
	server := New(6)

	for i := 1; i < 100000000; i++ {
		code := server.ToCode(i)
		if server.ToId(code) != i {
			t.Error("加密解密不一致 ", code, " != ", i)
		}
	}
}

func TestCodeService_ToUserCodeGen8(t *testing.T) {
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
}
