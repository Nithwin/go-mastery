package main

import (
	"fmt"
)

type AIProvider interface {
	Chat(prompt string) (string, error)
}

type Gemini struct{}

func (Gemini) Chat(prompt string)(string,error){
	return "Gemini", nil
}

type OpenAI struct{}

func (OpenAI) Chat(prompt string)(string,error){
	return "GPT", nil
}

type Claude struct{}
func (Claude) Chat(prompt string)(string,error){
	return "Claude", nil
}


func ask(provider AIProvider, prompt string){
	res , err := provider.Chat(prompt);
	if err == nil {
		fmt.Println(res);
	}
}
func main(){
	gemini := Gemini{}
	claude := Claude{}
	ask(gemini, "Hi");
	ask(claude, "Hi-lo");
}
