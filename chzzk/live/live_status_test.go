package live

import (
	"log"
	"testing"
)

func TestGet_live_status(t *testing.T) {
	t.Run("라이브 상태 러너 테스트", func(t *testing.T) {
		apiVersion := "v2"
		channelId := "bee4b42475b5937226b8b7ccbe2eb2dc"

		res, err := GetLiveStatus(apiVersion, channelId)
		if err != nil {
			t.Errorf("Request Error ::: %+v", err)
		}

		if res == nil {
			t.Errorf("Response Error ::: %+v", res)
		}
		log.Printf("%+v\n", res)
	})

	t.Run("라이브 상태 소풍왔니 테스트", func(t *testing.T) {
		apiVersion := "v2"
		channelId := "089185efc29a8fbe14ea294dc85f9661"

		res, err := GetLiveStatus(apiVersion, channelId)
		if err != nil {
			t.Errorf("Request Error ::: %+v", err)
		}

		if res == nil {
			t.Errorf("Response Error ::: %+v", res)
		}
		log.Printf("%+v\n", res)
	})

	t.Run("라이브 상태 강퀴88 테스트", func(t *testing.T) {
		apiVersion := "v2"
		channelId := "1a1dd9ce56fb61a37ffb6f69f6d5b978"

		res, err := GetLiveStatus(apiVersion, channelId)
		if err != nil {
			t.Errorf("Request Error ::: %+v", err)
		}

		if res == nil {
			t.Errorf("Response Error ::: %+v", res)
		}
		log.Printf("%+v\n", res)
	})

	t.Run("라이브 상태 한동숙 테스트", func(t *testing.T) {
		apiVersion := "v2"
		channelId := "75cbf189b3bb8f9f687d2aca0d0a382b"

		res, err := GetLiveStatus(apiVersion, channelId)
		if err != nil {
			t.Errorf("Request Error ::: %+v", err)
		}

		if res == nil {
			t.Errorf("Response Error ::: %+v", res)
		}
		log.Printf("%+v\n", res)
	})
}
