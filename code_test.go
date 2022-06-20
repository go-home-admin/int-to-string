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
