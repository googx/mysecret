/*
----------------------------------------------
 File Name: main.go
 Author: hanxu
 AuthorSite: http://www.googx.top/
 GitSource: https://github.com/googx/
 Created Time: 2024-12-21-18:16:26
-------------------功能说明-------------------

----------------------------------------------
*/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type QuestionFile struct {
	QuestionSha256 string
	QuestionList   []struct {
		Question string
		Answer   string
	}
}

func main() {
	//密保问题的json文件
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("read question file failed:%s ", err.Error())
		return
	}
	f := &QuestionFile{}
	if e := json.Unmarshal(file, f); e != nil {
		fmt.Printf("Unmarshal question file failed:%s ", e.Error())
		return
	}
	// 检查问题的hash是否匹配
	var (
		questionHasher = sha256.New()
		answerHasher   = sha256.New()
	)
	for i, s := range f.QuestionList {
		questionHasher.Write([]byte(s.Question))
		answerHasher.Write([]byte(s.Answer))
		fmt.Printf("%d. Question: [%s] , Answer: [%s] \n", i+1, s.Question, s.Answer)
	}
	questionHash := hex.EncodeToString(questionHasher.Sum(nil))
	if !strings.EqualFold(questionHash, f.QuestionSha256) {
		fmt.Printf("question hash not match")
		return
	}
	answerHash := hex.EncodeToString(answerHasher.Sum(nil))
	// 打印输出hash结果
	fmt.Println()
	fmt.Println("QuestionHash: ", questionHash)
	fmt.Println("AnswerHash: ", answerHash)
}
