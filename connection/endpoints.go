package connection

import "fmt"

const chzzkBaseURL = "https://api.chzzk.naver.com"
const chzzkChatURL = "wss://chzzk.naver.com"

// ChzzkLiveStatusEndPoint returns the live status endpoint URL for a given version and channel ID
func ChzzkLiveStatusEndPoint(version string, channelId string) string {
	endPoint := fmt.Sprintf("%s/polling/%s/channels/%s/live-status", chzzkBaseURL, version, channelId)
	return endPoint
}

// LiveDetails returns the live details endpoint URL for a given version and channel ID
func ChzzkLiveDetailsEndPoint(version string, channelId string) string {
	endPoint := fmt.Sprintf("%s/service/%s/channels/%s/live-detail", chzzkBaseURL, version, channelId)
	return endPoint
}

// Setting returns the setting endpoint URL for a given version and channel ID
func ChzzkSettingEndPoint(version string, channelId string) string {
	endPoint := fmt.Sprintf("%s/service/%s/channels/%s/donations/setting", chzzkBaseURL, version, channelId)
	return endPoint
}

// VideoSetting returns the video setting endpoint URL for a given version and channel ID
func ChzzkVideoSettingEndPoint(version string, channelId string) string {
	endPoint := fmt.Sprintf("%s/service/%s/channels/%s/donations/video-setting", chzzkBaseURL, version, channelId)
	return endPoint
}

// ChatSetting returns the chat setting endpoint URL for a given version and channel ID
func ChzzkChatSettingEndPoint(version string, channelId string) string {
	endPoint := fmt.Sprintf("%s/service/%s/channels/%s/donations/chat-setting", chzzkBaseURL, version, channelId)
	return endPoint
}

// ChatRules returns the chat rules endpoint URL for a given version and channel ID
func ChzzkChatRulesEndPoint(version string, channelId string) string {
	endPoint := fmt.Sprintf("%s/service/%s/channels/%s/chat-rules", chzzkBaseURL, version, channelId)
	return endPoint
}

func ChzzkLiveChattingEndPoint(channelId string) string {
	endPoint := fmt.Sprintf("%s/live/%s/chat", chzzkChatURL, channelId)
	return endPoint
}
