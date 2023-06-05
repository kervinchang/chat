package service

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kervinchang/chat/config"
	"github.com/sashabaranov/go-openai"
)

var SystemMessage = openai.ChatCompletionMessage{Role: "system", Content: config.BotDesc}

// CreateChatCompletion 对话补全
func CreateChatCompletion(ctx *gin.Context, messages []openai.ChatCompletionMessage) (string, error) {
	request := openai.ChatCompletionRequest{
		Model:    config.Model,
		Messages: messages,
	}

	gptConfig := openai.DefaultConfig(config.ApiKey)

	transport := &http.Transport{}

	if strings.HasPrefix(config.Proxy, "socks5h://") {
		// 创建一个 DialContext 对象，并设置代理服务器
		dialContext, err := NewDialContext(config.Proxy[10:])
		if err != nil {
			panic(err)
		}
		transport.DialContext = dialContext
	} else {
		// 创建一个 HTTP Transport 对象，并设置代理服务器
		proxyUrl, err := url.Parse(config.Proxy)
		if err != nil {
			panic(err)
		}
		transport.Proxy = http.ProxyURL(proxyUrl)
	}
	// 创建一个 HTTP 客户端，并将 Transport 对象设置为其 Transport 字段
	gptConfig.HTTPClient = &http.Client{
		Transport: transport,
	}

	client := openai.NewClientWithConfig(gptConfig)
	if request.Messages[0].Role != "system" {
		newMessage := append([]openai.ChatCompletionMessage{
			{Role: "system", Content: config.BotDesc},
		}, request.Messages...)
		request.Messages = newMessage
	}

	request.Model = config.Model
	resp, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
