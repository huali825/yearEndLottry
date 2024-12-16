package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"time"
)

func main() {
	// 创建一个新的应用程序
	myApp := app.New()
	// 创建一个新的窗口
	myWin := myApp.NewWindow("Hello")

	// 创建一个新的Label组件，用于显示时间
	clock := widget.NewLabel("")
	// 将Label组件设置为窗口的内容
	myWin.SetContent(clock)

	// 启动一个goroutine，每秒更新一次时间
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	// 设置窗口内容为标签
	//myWin.SetContent(widget.NewLabel("Hello Fyne!"))
	// 设置窗口大小为200x200
	myWin.Resize(fyne.NewSize(200, 200))
	// 显示窗口并运行应用程序
	myWin.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	// 获取当前时间，并按照指定格式进行格式化
	formatted := time.Now().Format("Time: 03:04:05")
	// 将格式化后的时间设置为Label组件的文本
	clock.SetText(formatted)
}
