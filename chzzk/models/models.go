package models

import (
	"encoding/json"
	"time"
)

// ContentLiveStatus represents the content field in the /liveStatus API response
type ContentLiveStatus struct {
	LiveTitle              string `json:"liveTitle"`              // Title of the live stream
	Status                 string `json:"status"`                 // Status of the live stream (OPEN / CLOSE)
	ConcurrentUserCount    int    `json:"concurrentUserCount"`    // Number of current concurrent viewers
	AccumulateCount        int    `json:"accumulateCount"`        // Total accumulated viewer count
	PaidPromotion          bool   `json:"paidPromotion"`          // Unknown field related to paid promotion
	Adult                  bool   `json:"adult"`                  // Indicates if the stream is set as Adult content
	ChatChannelId          string `json:"chatChannelId"`          // UUID of the chat channel
	CategoryType           string `json:"categoryType"`           // Category of the stream (e.g., GAME, ETC)
	LiveCategory           string `json:"liveCategory"`           // Subject of the stream (e.g., Lost_Ark, talk)
	LiveCategoryValue      string `json:"liveCategoryValue"`      // Value of the live category (e.g., Lost_Ark, talk)
	LivePollingStatusJson  string `json:"livePollingStatusJson"`  // Unknown field related to live polling status
	FaultStatus            string `json:"faultStatus,omitempty"`  // Unknown field related to fault status
	UserAdultStatus        string `json:"userAdultStatus"`        // Indicates the adult status of the logged in user
	ChatActive             bool   `json:"chatActive"`             // Indicates if chat is active
	ChatAvailableGroup     string `json:"chatAvailableGroup"`     // Indicates the group that can chat
	ChatAvailableCondition string `json:"chatAvailableCondition"` // Indicates the conditions for the user to chat
	MinFollowerMinute      int    `json:"minFollowerMinute"`      // Minimum follower minutes
}

// ChannelDetails represents the channel details in the /liveDetail API response
type ChannelDetails struct {
	ChannelId       string `json:"channelId"`       // UUID of the channel
	ChannelName     string `json:"channelName"`     // Name of the channel
	ChannelImageUrl string `json:"channelImageUrl"` // URL of the channel image
	VerifiedMark    bool   `json:"verifiedMark"`    // Indicates if the channel is verified
}

// ContentLiveDetail represents the content field in the /liveDetail API response
type ContentLiveDetail struct {
	LiveId                   int            `json:"liveId"`                   // Unique ID of the live stream
	LiveTitle                string         `json:"liveTitle"`                // Title of the live stream
	Status                   string         `json:"status"`                   // Status of the live stream (OPEN / CLOSE)
	LiveImageUrl             string         `json:"liveImageUrl"`             // URL of the live stream image
	DefaultThumbnailImageUrl string         `json:"defaultThumbnailImageUrl"` // URL of the default thumbnail image
	ConcurrentUserCount      int            `json:"concurrentUserCount"`      // Number of current concurrent viewers
	AccumulateCount          int            `json:"accumulateCount"`          // Total accumulated viewer count
	OpenDate                 string         `json:"openDate"`                 // Starting time of the live stream
	CloseDate                string         `json:"closeDate"`                // Closing time of the live stream
	Adult                    bool           `json:"adult"`                    // Indicates if the stream is set as Adult content
	ChatChannelId            string         `json:"chatChannelId"`            // UUID of the chat channel
	CategoryType             string         `json:"categoryType"`             // Category of the stream (e.g., GAME, ETC)
	LiveCategory             string         `json:"liveCategory"`             // Subject of the stream (e.g., Lost_Ark, talk)
	LiveCategoryValue        string         `json:"liveCategoryValue"`        // Value of the live category (e.g., Lost_Ark, talk)
	ChatActive               bool           `json:"chatActive"`               // Indicates if chat is active
	ChatAvailableGroup       string         `json:"chatAvailableGroup"`       // Indicates the group that can chat
	PaidPromotion            bool           `json:"paidPromotion"`            // Unknown field related to paid promotion
	ChatAvailableCondition   string         `json:"chatAvailableCondition"`   // Indicates the conditions for the user to chat
	UserAdultStatus          string         `json:"userAdultStatus"`          // Indicates the adult status of the logged in user
	Channel                  ChannelDetails `json:"channel"`                  // Details about the channel
}

// ContentSetting represents the content field in the /setting API response
type ContentSetting struct {
	ChatDonationActive  bool `json:"chatDonationActive"`  // Indicates if chat donation is active
	VideoDonationActive bool `json:"videoDonationActive"` // Indicates if video donation is active

}

// ContentVideoSetting represents the content field in the /videoSetting API response
type ContentVideoSetting struct {
	DonationActive       bool `json:"donationActive"`       // Indicates if donation is active
	PayAmountPerSecond   int  `json:"payAmountPerSecond"`   // Amount paid per second for video donation
	MinCurrencyPayAmount int  `json:"minCurrencyPayAmount"` // Minimum donation amount
	MaxDurationLength    int  `json:"maxDurationLength"`    // Maximum length of video donation
	IsYoutubeVideoAllow  bool `json:"isYoutubeVideoAllow"`  // Indicates if Youtube videos are allowed for video donation
	IsChzzkClipAllow     bool `json:"isChzzkClipAllow"`     // Indicates if Chzzk Clip videos are allowed for video donation
	IsAllowForSubscriber bool `json:"isAllowForSubscriber"` // Indicates if donation is only allowed for subscribers
}

// ContentChatSetting represents the content field in the /chatSetting API response
type ContentChatSetting struct {
	DonationActive         bool `json:"donationActive"`       // Indicates if donation is active
	MinCurrencyPayAmount   int  `json:"minCurrencyPayAmount"` // Minimum donation amount
	ExposureDonationAmount bool `json:"exposureDonationAmount"`
}

type ContentChatRules struct {
	Agree        bool   `json:"agree"`
	ChannelId    string `json:"channelId"`
	Rule         string `json:"rule"`
	UpdatedDate  string `json:"updatedDate"`
	ServiceAgree bool   `json:"serviceAgree"`
}

// ResponseContent represents the content field in the /chatRules API response
type ResponseContent struct {
	Agree        bool      `json:"agree"`        // Unknown field
	ChannelId    string    `json:"channelId"`    // UUID of the channel
	Rule         string    `json:"rule"`         // Unknown field
	UpdatedDate  time.Time `json:"updatedDate"`  // Time the chat rules were updated
	ServiceAgree bool      `json:"serviceAgree"` // Unknown field
}

// Response represents the API response
type Response struct {
	Code    int             `json:"code"`              // HTTP status code of the response
	Message string          `json:"message,omitempty"` // Message of the response (if any)
	Content json.RawMessage `json:"content"`           // Content of the response, can be any of the previously defined content types
}
