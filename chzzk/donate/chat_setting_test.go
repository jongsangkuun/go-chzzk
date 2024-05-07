package donate

import (
	"log"
	"testing"
)

func TestGet_chat_setting(t *testing.T) {
	t.Run("채팅 규칙 러너 테스트", func(t *testing.T) {
		apiVersion := "v1"
		channelId := "bee4b42475b5937226b8b7ccbe2eb2dc"

		res, err := GetChatSetting(apiVersion, channelId)
		if err != nil {
			t.Errorf("Request Error ::: %+v", err)
		}

		if res == nil {
			t.Errorf("Response Error ::: %+v", res)
		}
		log.Printf("%+v\n", res)
	})

	t.Run("채팅 규칙 소풍왔니 테스트", func(t *testing.T) {
		apiVersion := "v1"
		channelId := "089185efc29a8fbe14ea294dc85f9661"

		res, err := GetChatSetting(apiVersion, channelId)
		if err != nil {
			t.Errorf("Request Error ::: %+v", err)
		}

		if res == nil {
			t.Errorf("Response Error ::: %+v", res)
		}
		log.Printf("%+v\n", res)
	})

	t.Run("채팅 규칙 강퀴88 테스트", func(t *testing.T) {
		apiVersion := "v1"
		channelId := "1a1dd9ce56fb61a37ffb6f69f6d5b978"

		res, err := GetChatSetting(apiVersion, channelId)
		if err != nil {
			t.Errorf("Request Error ::: %+v", err)
		}

		if res == nil {
			t.Errorf("Response Error ::: %+v", res)
		}
		log.Printf("%+v\n", res)
	})

	t.Run("채팅 규칙 한동숙 테스트", func(t *testing.T) {
		apiVersion := "v1"
		channelId := "75cbf189b3bb8f9f687d2aca0d0a382b"

		res, err := GetChatSetting(apiVersion, channelId)
		if err != nil {
			t.Errorf("Request Error ::: %+v", err)
		}

		if res == nil {
			t.Errorf("Response Error ::: %+v", res)
		}
		log.Printf("%+v\n", res)
	})

	t.Run("채팅 규칙 시라유키히나 테스트", func(t *testing.T) {
		apiVersion := "v1"
		channelId := "b044e3a3b9259246bc92e863e7d3f3b8"

		res, err := GetChatSetting(apiVersion, channelId)
		if err != nil {
			t.Errorf("Request Error ::: %+v", err)
		}

		if res == nil {
			t.Errorf("Response Error ::: %+v", res)
		}
		log.Printf("%+v\n", res)
	})

}
