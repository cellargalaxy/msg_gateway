package main

import (
	"github.com/cellargalaxy/go_common/util"
	"testing"
)

func TestSendTemplateMessageWithUrl(t *testing.T) {
	ctx := util.GenCtx()
	SendTemplateMessageWithUrl(ctx, "通用消息", "https://baidu.com", map[string]string{"logid": "logid", "sn": "sn", "text": "text"})
}
