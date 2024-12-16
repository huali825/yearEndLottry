package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("抽奖程序")
	myWindow.Resize(fyne.NewSize(600, 400))

	// 创建按钮
	startButton := widget.NewButton("开始抽奖", func() {
		// 开始抽奖的逻辑
	})
	resultButton := widget.NewButton("结果", func() {
		// 结果按钮的逻辑
	})
	helpButton := widget.NewButton("帮助", func() {
		// 帮助按钮的逻辑
	})

	// 创建设置下拉菜单的回调函数
	showSettingMenu := func() {
		popup := dialog.NewCustom("设置", "关闭", container.NewVBox(
			widget.NewButton("导入抽奖名单", func() {
				// 导入抽奖名单的逻辑
			}),
			widget.NewButton("设置抽奖规则", func() {
				// 设置抽奖规则的逻辑
			}),
			widget.NewButton("设置中奖模式", func() {
				// 设置中奖模式的逻辑
			}),
			widget.NewButton("清空抽奖数据", func() {
				// 清空抽奖数据的逻辑
			}),
			widget.NewButton("清空所有数据", func() {
				// 清空所有数据的逻辑
			}),
		), myWindow)
		popup.Show()
	}

	// 为设置按钮创建一个点击事件来显示下拉菜单
	settingButton := widget.NewButton("设置", showSettingMenu)

	// 创建按钮容器
	buttonContainer := container.NewHBox(
		startButton,
		settingButton,
		resultButton,
		helpButton,
	)

	// 将按钮容器添加到窗口
	myWindow.SetContent(buttonContainer)

	myWindow.ShowAndRun()
}
