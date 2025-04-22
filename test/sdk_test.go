package test

import (
	"github.com/cellargalaxy/go_common/util"
	"github.com/cellargalaxy/msg_gateway/sdk"
	"testing"
)

func TestSendTemplateMessageWithUrl(t *testing.T) {
	ctx := util.GenCtx()
	sdk.SendTemplateMessageWithUrl(ctx, "通用消息", "https://baidu.com", map[string]string{"logid": "logid", "sn": "sn", "text": "text"})
}
