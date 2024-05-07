package chat

import (
	"encoding/json"
	"fmt"
	"github.com/jongsangkuun/go-chzzk/chzzk/models"
	"github.com/jongsangkuun/go-chzzk/connection"

	"log"
)

func GetChatRules(apiVersion string, channelId string) (*models.ContentChatRules, error) {
	response := models.Response{}
	// 실시간 상태 정보를 담을 데이터 구조체를 선언합니다.
	chatRules := models.ContentChatRules{}

	// chatRulesUrl API 엔드포인트로, 입력된 apiVersion과 channelId를 사용하여 생성합니다.
	chatRulesUrl := connection.ChzzkChatRulesEndPoint(apiVersion, channelId)
	// 생성된 API 엔드포인트를 출력합니다.
	log.Println(chatRulesUrl)

	// 해당 엔드포인트로 HTTP 요청을 보내고 응답 데이터를 body에 저장합니다.
	body, err := connection.GetHTTPDataFast(chatRulesUrl)
	if err != nil {
		return nil, err
	}

	// 응답받은 데이터를 response 구조체로 변환 (Unmarshal) 합니다.
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	// HTTP 응답 코드가 200이 아닌 경우 에러를 반환합니다.
	if response.Code != 200 {
		return nil, fmt.Errorf("응답 코드가 %d입니다", response.Code)
	}

	// response의 Content 부분을 liveStatus 구조체로 변환합니다.
	err = json.Unmarshal(response.Content, &chatRules)
	if err != nil {
		return nil, err
	}

	return &chatRules, nil
}
