package webhook

import "os"

type Webhooker interface {
	SendMessage(msg string) error
	SendMessageWithAt(msg string, atMobiles ...string) error
}

type Webhook struct {
	// 钉钉机器人webhook地址
	DingDingWebhook string `json:"dingDingWebhook"`
	// 企业微信机器人webhook地址
	WeChatWebhook string `json:"weChatWebhook"`
	// 飞书机器人webhook地址
	FeiShuWebhook string `json:"feiShuWebhook"`
	// 企业微信应用webhook地址
	WeChatAppWebhook string `json:"weChatAppWebhook"`
	// 企业微信应用secret
	WeChatAppSecret string `json:"weChatAppSecret"`
}

func NewWebhook() *Webhook {
	return &Webhook{
		DingDingWebhook:  os.Getenv("DING_DING_WEBHOOK"),
		WeChatWebhook:    os.Getenv("WE_CHAT_WEBHOOK"),
		FeiShuWebhook:    os.Getenv("FEI_SHU_WEBHOOK"),
		WeChatAppWebhook: os.Getenv("WE_CHAT_APP_WEBHOOK"),
		WeChatAppSecret:  os.Getenv("WE_CHAT_APP_SECRET"),
	}
}

func (w *Webhook) SendMessage(msg string) error {
	switch {
	case w.DingDingWebhook != "":
		return w.sendDingDingMessage(msg)
	case w.WeChatWebhook != "":
		return w.sendWeChatMessage(msg)
	case w.FeiShuWebhook != "":
		return w.sendFeiShuMessage(msg)
	case w.WeChatAppWebhook != "":
		return w.sendWeChatAppMessage(msg)
	default:
		return nil
	}
}

func (w *Webhook) SendMessageWithAt(msg string, atMobiles ...string) error {
	switch {
	case w.DingDingWebhook != "":
		return w.sendDingDingMessageWithAt(msg, atMobiles...)
	case w.WeChatWebhook != "":
		return w.sendWeChatMessageWithAt(msg, atMobiles...)
	case w.FeiShuWebhook != "":
		return w.sendFeiShuMessageWithAt(msg, atMobiles...)
	case w.WeChatAppWebhook != "":
		return w.sendWeChatAppMessageWithAt(msg, atMobiles...)
	default:
		return nil
	}
}

func (w *Webhook) sendDingDingMessage(msg string) error {
	// 发送钉钉消息
	return nil
}

func (w *Webhook) sendWeChatMessage(msg string) error {
	// 发送微信消息
	return nil
}

func (w *Webhook) sendFeiShuMessage(msg string) error {
	// 发送飞书消息
	return nil
}

func (w *Webhook) sendWeChatAppMessage(msg string) error {
	// 发送企业微信消息
	return nil
}

func (w *Webhook) sendWeChatMessageWithAt(msg string, atMobiles ...string) error {
	// 发送微信消息
	return nil
}

func (w *Webhook) sendFeiShuMessageWithAt(msg string, atMobiles ...string) error {
	// 发送飞书消息
	return nil
}

func (w *Webhook) sendWeChatAppMessageWithAt(msg string, atMobiles ...string) error {
	// 发送企业微信消息
	return nil
}

func (w *Webhook) sendDingDingMessageWithAt(msg string, atMobiles ...string) error {
	// 发送钉钉消息
	return nil
}