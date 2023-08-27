package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"thresher/infra/model"
	"thresher/utils/config"

	"github.com/gin-gonic/gin"
)

type DetailMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Choices struct {
	Index        int16         `json:"index"`
	FinishReason string        `json:"finish_reason"`
	Message      DetailMessage `json:"message"`
	Usage        interface{}   `json:"usage"`
}

type Message struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created int32     `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
}

func GetOpenAiChat(ctx *gin.Context, msgs []model.Message, target model.Users) (*string, error) {
	conf := config.LoadConfig().OpenAi
	endpoint := fmt.Sprintf("%sopenai/deployments/%s/chat/completions?api-version=2023-05-15", conf.Endpoint, conf.Model)
	classifiedMessages := classifyMessages(msgs, target)
	requestData := map[string]interface{}{
		"messages": classifiedMessages,
	}

	requestDataJSON, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		"POST",
		endpoint,
		bytes.NewBuffer([]byte(requestDataJSON)),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", conf.Secret)

	postClient := &http.Client{}
	res, err := postClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var responseMessages Message
	if err := json.Unmarshal(body, &responseMessages); err != nil {
		return nil, err
	}
	msg := string(responseMessages.Choices[0].Message.Content)
	return &msg, nil
}
func classifyMessages(messages []model.Message, target model.Users) []map[string]string {
	classifiedMessages := make([]map[string]string, len(messages)+1)
	description := target.Description
	content := ""
	if description != "" {
		content = fmt.Sprintf("あなたは%sという特徴があります。", target.Description)
	}
	content += "会話をしてください。あなたがAIであるということは秘密です。人間を装ってください。\n {content:string}というフォーマットで送ってください #出力"
	classifiedMessages[0] = map[string]string{
		"role":    "system",
		"content": content,
	}

	for i, msg := range messages {
		label := "assistant"
		if msg.SenderID != target.ID {
			label = "user"
		}

		// 分類されたメッセージを作成
		classifiedMsg := map[string]string{
			"role":    label,
			"content": msg.Content,
		}

		classifiedMessages[i+1] = classifiedMsg
	}

	return classifiedMessages
}
