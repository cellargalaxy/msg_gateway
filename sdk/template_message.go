package sdk

import (
	"context"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/templateMessage/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/templateMessage/response"
	"github.com/cellargalaxy/go_common/util"
	"github.com/sirupsen/logrus"
)

func GetTemplateId(ctx context.Context, name string) (string, error) {
	value, err := GetTemplate(ctx, name)
	if err != nil {
		return "", err
	}
	if value == nil {
		return "", err
	}
	return value.TemplateID, nil
}
func GetTemplate(ctx context.Context, name string) (*response.Template, error) {
	list, err := GetTemplates(ctx, nil)
	if err != nil {
		return nil, err
	}
	for i := range list {
		if list[i].Title == name {
			return list[i], nil
		}
	}
	return nil, nil
}
func GetTemplates(ctx context.Context, req any) ([]*response.Template, error) {
	resp, err := ReTry(ctx, getTemplates, req)
	return resp, err
}
func getTemplates(ctx context.Context, req any) ([]*response.Template, error) {
	sdk, err := GetSdk(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.TemplateMessage.GetPrivateTemplates(ctx)
	if resp == nil || err != nil {
		return nil, err
	}
	return resp.TemplateList, err
}

func SendTemplateText(ctx context.Context, name string, url string, text string) {
	data := make(map[string]string)
	data["logid"] = util.GetLogIdString(ctx)
	data["sn"] = util.GetEnv("sn")
	data["text"] = text
	SendTemplateData(ctx, name, url, data)
}
func SendTemplateData(ctx context.Context, name string, url string, data map[string]string) {
	templateId, err := GetTemplateId(ctx, name)
	if templateId == "" || err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{"err": err}).Error("发送模板消息，查询模板为空")
		return
	}

	hash := make(map[string]interface{})
	for key, value := range data {
		hash[key] = map[string]string{"value": value}
	}
	hashMap := power.HashMap(hash)
	var req request.RequestTemlateMessage
	req.TemplateID = templateId
	req.URL = url
	req.Data = &hashMap

	users := GetUsers()
	for _, user := range users {
		req.ToUser = user
		SendTemplateMessage(ctx, req)
	}
}
func SendTemplateMessage(ctx context.Context, req request.RequestTemlateMessage) (*response.ResponseTemplateSend, error) {
	resp, err := ReTry(ctx, sendTemplateMessage, req)
	return resp, err
}
func sendTemplateMessage(ctx context.Context, req request.RequestTemlateMessage) (*response.ResponseTemplateSend, error) {
	sdk, err := GetSdk(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.TemplateMessage.Send(ctx, &req)
	if resp == nil || err != nil {
		return nil, err
	}
	return resp, err
}
