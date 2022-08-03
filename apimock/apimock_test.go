package apimock

import "testing"

func TestGenerateAPIDocument(t *testing.T) {
	var apiMockList = []*APIMock{
		{
			Title:       "title test",
			Description: "description test",
			Route:       "url test",
		},
	}
	if result, err := GenerateAPIDocument(APIDDocumentFormatMarkdown, apiMockList); err != nil || len(result) <= 0 {
		t.Fail()
		return
	}

	if result, err := GenerateAPIDocument(APIDDocumentFormatHtml, apiMockList); err == nil || len(result) > 0 {
		t.Fail()
		return
	}
}
