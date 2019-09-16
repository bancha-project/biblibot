package bot

import (
	"github.com/bancha-project/biblibot/infra/env"
	"github.com/nlopes/slack"
	"strings"
)

func CanReply(ev *slack.MessageEvent) bool {
	// 指定したチャンネル以外は対象外
	if ev.Channel != env.GetEnv().ChannelId {
		return false
	}

	// メンションされていなければ対象外
	if !strings.HasPrefix(ev.Msg.Text, "<@" + env.GetEnv().BotId + ">") {
		return false
	}

	// メッセージの変更の場合は返信しない
	if ev.Msg.SubType != "" {
		return false
	}

	return true
}