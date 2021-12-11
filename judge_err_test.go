package jerr

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestJudegErr(t *testing.T) {
	req, err := http.NewRequest("GET1", "http://www.baidu.com1", nil)
	JudgeErr(err, false, "err = ")

	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4595.0 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(req)
	JudgeErr(err, false, "err = ")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	JudgeErr(err, false, "err = ")

	t.Log(string(body))
}

// go test -v judge_err_test.go judge_err.go -run TestJudegErr
