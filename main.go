package main

import (
	"syscall/js"
	"unicode/utf8"
)

func count() {
	document := js.Global().Get("document")

	document.Call("getElementById", "countForm").Call("addEventListener", "submit", js.NewEventCallback(js.PreventDefault, func(e js.Value) {
		labelStr := "文字数"
		charCount := utf8.RuneCountInString(document.Call("getElementById", "textArea").Get("value").String())
		document.Call("getElementById", "resultLabel").Set("textContent", labelStr)
		document.Call("getElementById", "result").Set("textContent", charCount)
	}))
}

func main() {
	c := make(chan struct{}, 0)
	count()
	<-c
}
