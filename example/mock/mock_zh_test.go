package mock

import (
	"github.com/WGrape/apimock/apimock"
	"os"
	"testing"
)

func TestGetUserInZh(t *testing.T) {
	var apiMockList []*apimock.APIMock

	apiMock := GetUserInZh()
	apiMockList = append(apiMockList, apiMock)

	result, err := apimock.GenerateAPIDocument(apimock.APIDDocumentFormatMarkdown, apiMockList)
	if err != nil {
		t.Fail()
		return
	}

	if err = os.WriteFile("../apidoc.md", result, 0666); err != nil {
		t.Fail()
		return
	}
}
