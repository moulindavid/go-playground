package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type SummaryResponse struct {
	Id      string `json:"id"`
	Summary string `json:"summary"`
	Meta    string `json:"meta"` //TODO
}

type SummaryRequest struct {
	Text string `json:"text"`
}

func SummarizeText(content string) SummaryResponse {
	url := "https://api.cohere.ai/v1/summarize"
	request := SummaryRequest{Text: content}
	string_request, _ := json.Marshal(request)

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(string_request)))
	api_token := os.Getenv("COHERE_TOKEN")
	fmt.Println(os.Getenv("HOME"))
	fmt.Println(api_token)
	fmt.Println(os.Getenv("COHERE_TOKEN"))
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", api_token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var summary SummaryResponse
	json.Unmarshal(body, &summary)
	return summary
}
