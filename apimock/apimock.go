package apimock

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// RequestExplainItem the explanation of request
type RequestExplainItem struct {
	Field  string `json:"field"`
	Type   string `json:"type"`
	IsMust bool   `json:"is_must"`
	Mark   string `json:"mark"`
}

// ResponseExplainItem the explanation of response
type ResponseExplainItem struct {
	Field string `json:"field"`
	Type  string `json:"type"`
	Mark  string `json:"mark"`
}

// APIMock the core struct of APIMock
type APIMock struct {
	Title           string                `json:"title"`
	Description     string                `json:"description"`
	Route           string                `json:"route"`
	Request         interface{}           `json:"request"`
	RequestExplain  []RequestExplainItem  `json:"request_explain"`
	Response        interface{}           `json:"response"`
	ResponseExplain []ResponseExplainItem `json:"response_explain"`
}

// some constants
const (
	APIDDocumentFormatMarkdown = 1
	APIDDocumentFormatHtml     = 2
)

// GenerateAPIDocument generate the api document of different formats
func GenerateAPIDocument(docType int, apiMockList []*APIMock) ([]byte, error) {
	switch docType {
	case APIDDocumentFormatMarkdown:
		return generateMarkdownAPIDocument(apiMockList)
	}
	return nil, errors.New("not support document format")
}

// generateMarkdownAPIDocument generate the api document of markdown format
func generateMarkdownAPIDocument(apiMockList []*APIMock) ([]byte, error) {
	var (
		index        int
		apiMock      *APIMock
		contentList  []string
		headLinkList []string
	)
	for index, apiMock = range apiMockList {
		headLinkList = append(headLinkList, fmt.Sprintf("- [%d. %s](#%d-%s)\n", index+1, apiMock.Title, index+1, apiMock.Title))
		contentList = append(contentList, "## "+fmt.Sprintf("%d. ", index+1)+apiMock.Title+"\n")

		contentList = append(contentList, "\n### (1) Description\n")
		contentList = append(contentList, apiMock.Description+"\n")

		contentList = append(contentList, "\n### (2) URL\n")
		contentList = append(contentList, apiMock.Route+"\n")

		contentList = append(contentList, "\n### (3) Request\n")
		contentList = append(contentList, "| Field | Type | Required | Mark |\n| :----: | :----: | :----: | :----: |\n")
		for _, explain := range apiMock.RequestExplain {
			isMust := "yes"
			if explain.IsMust {
				isMust = "no"
			}
			contentList = append(contentList, fmt.Sprintf("|%s|%s|%s|%s|\n", explain.Field, explain.Type, isMust, explain.Mark))
		}
		bytes, err := json.MarshalIndent(apiMock.Request, "", "    ")
		if err != nil {
			return nil, err
		}
		contentList = append(contentList, "```json\n"+string(bytes)+"\n```\n")

		contentList = append(contentList, "\n### (4) Response\n")
		contentList = append(contentList, "| Field | Type | Mark |\n| :----: | :----: | :----: |\n")
		for _, explain := range apiMock.ResponseExplain {
			contentList = append(contentList, fmt.Sprintf("|%s|%s|%s|\n", explain.Field, explain.Type, explain.Mark))
		}
		bytes, err = json.MarshalIndent(apiMock.Response, "", "    ")
		if err != nil {
			return nil, err
		}
		contentList = append(contentList, "```json\n"+string(bytes)+"\n```\n")
	}
	if len(headLinkList) > 0 {
		headLinkList = append([]string{"#### Content\n"}, headLinkList...)
	}

	var (
		lastUpdateTime = time.Now().Format("2006-01-02 15:04:05")
		headerTemplate = "# API Document\nThis [API document](https://github.com/WGrape/apimock) can be automatically generated and quickly released in the CI stage without manual editing.\n\n#### Version\nLast update time at ```%s```"
		header         = fmt.Sprintf(headerTemplate, lastUpdateTime)
	)
	result := fmt.Sprintf("%s\n%s\n%s", header, strings.Join(headLinkList, ""), strings.Join(contentList, ""))

	return []byte(result), nil
}
