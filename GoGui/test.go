package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"log"
)

//var isSpecialMode = walk.NewMutableCondition()
//
//type MyMainWindow struct {
//	*walk.MainWindow
//}
//
//func main() {
//	MustRegisterCondition("isSpecialMode", isSpecialMode)
//
//	mw := new(MyMainWindow)
//
//	var openAction, showAboutBoxAction *walk.Action
//	var recentMenu *walk.Menu
//	var toggleSpecialModePB *walk.PushButton
//
//	if err := (MainWindow{
//		AssignTo: &mw.MainWindow,
//		Title:    "Walk Actions Example",
//		MenuItems: []MenuItem{
//			Menu{
//				Text: "&File",
//				Items: []MenuItem{
//					Action{
//						AssignTo:    &openAction,
//						Text:        "&Open",
//						//Image:       "../img/open.png",
//						Enabled:     Bind("enabledCB.Checked"),
//						Visible:     Bind("!openHiddenCB.Checked"),
//						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
//						OnTriggered: mw.openAction_Triggered,
//					},
//					Menu{
//						AssignTo: &recentMenu,
//						Text:     "Recent",
//					},
//					Separator{},
//					Action{
//						Text:        "E&xit",
//						OnTriggered: func() { mw.Close() },
//					},
//				},
//			},
//			Menu{
//				Text: "&View",
//				Items: []MenuItem{
//					Action{
//						Text:    "Open / Special Enabled",
//						Checked: Bind("enabledCB.Visible"),
//					},
//					Action{
//						Text:    "Open Hidden",
//						Checked: Bind("openHiddenCB.Visible"),
//					},
//				},
//			},
//			Menu{
//				Text: "&Help",
//				Items: []MenuItem{
//					Action{
//						AssignTo:    &showAboutBoxAction,
//						Text:        "About",
//						OnTriggered: mw.showAboutBoxAction_Triggered,
//					},
//				},
//			},
//		},
//		ToolBar: ToolBar{
//			ButtonStyle: ToolBarButtonImageBeforeText,
//			Items: []MenuItem{
//				ActionRef{&openAction},
//				Menu{
//					Text:  "New A",
//					//Image: "../img/document-new.png",
//					Items: []MenuItem{
//						Action{
//							Text:        "A",
//							OnTriggered: mw.newAction_Triggered,
//						},
//						Action{
//							Text:        "B",
//							OnTriggered: mw.newAction_Triggered,
//						},
//						Action{
//							Text:        "C",
//							OnTriggered: mw.newAction_Triggered,
//						},
//					},
//					OnTriggered: mw.newAction_Triggered,
//				},
//				Separator{},
//				Menu{
//					Text:  "View",
//					//Image: "../img/document-properties.png",
//					Items: []MenuItem{
//						Action{
//							Text:        "X",
//							OnTriggered: mw.changeViewAction_Triggered,
//						},
//						Action{
//							Text:        "Y",
//							OnTriggered: mw.changeViewAction_Triggered,
//						},
//						Action{
//							Text:        "Z",
//							OnTriggered: mw.changeViewAction_Triggered,
//						},
//					},
//				},
//				Separator{},
//				Action{
//					Text:        "Special",
//					//Image:       "../img/system-shutdown.png",
//					Enabled:     Bind("isSpecialMode && enabledCB.Checked"),
//					OnTriggered: mw.specialAction_Triggered,
//				},
//			},
//		},
//		ContextMenuItems: []MenuItem{
//			ActionRef{&showAboutBoxAction},
//		},
//		MinSize: Size{300, 200},
//		Layout:  VBox{},
//		Children: []Widget{
//			CheckBox{
//				Name:    "enabledCB",
//				Text:    "Open / Special Enabled",
//				Checked: true,
//				Accessibility: Accessibility{
//					Help: "Enables Open and Special",
//				},
//			},
//			CheckBox{
//				Name:    "openHiddenCB",
//				Text:    "Open Hidden",
//				Checked: true,
//			},
//			PushButton{
//				AssignTo: &toggleSpecialModePB,
//				Text:     "Enable Special Mode",
//				OnClicked: func() {
//					isSpecialMode.SetSatisfied(!isSpecialMode.Satisfied())
//
//					if isSpecialMode.Satisfied() {
//						toggleSpecialModePB.SetText("Disable Special Mode")
//					} else {
//						toggleSpecialModePB.SetText("Enable Special Mode")
//					}
//				},
//				Accessibility: Accessibility{
//					Help: "Toggles special mode",
//				},
//			},
//		},
//	}.Create()); err != nil {
//		log.Fatal(err)
//	}
//
//	addRecentFileActions := func(texts ...string) {
//		for _, text := range texts {
//			a := walk.NewAction()
//			a.SetText(text)
//			a.Triggered().Attach(mw.openAction_Triggered)
//			recentMenu.Actions().Add(a)
//		}
//	}
//
//	addRecentFileActions("Foo", "Bar", "Baz")
//
//	mw.Run()
//}
//
//func (mw *MyMainWindow) openAction_Triggered() {
//	walk.MsgBox(mw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
//}
//
//func (mw *MyMainWindow) newAction_Triggered() {
//	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
//}
//
//func (mw *MyMainWindow) changeViewAction_Triggered() {
//	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
//}
//
//func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
//	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
//}
//
//func (mw *MyMainWindow) specialAction_Triggered() {
//	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
//}

type MyWindow struct {
	*walk.MainWindow
	hWnd        win.HWND
	minimizeBox *walk.CheckBox
	maximizeBox *walk.CheckBox
	closeBox    *walk.CheckBox
	sizeBox     *walk.CheckBox
	ni          *walk.NotifyIcon
}

func (mw *MyWindow) SetMinimizeBox() {
	if mw.minimizeBox.Checked() {
		mw.addStyle(win.WS_MINIMIZEBOX)
		return
	}
	mw.removeStyle(^win.WS_MINIMIZEBOX)
}

func (mw *MyWindow) SetMaximizeBox() {
	if mw.maximizeBox.Checked() {
		mw.addStyle(win.WS_MAXIMIZEBOX)
		return
	}
	mw.removeStyle(^win.WS_MAXIMIZEBOX)
}

func (mw *MyWindow) SetSizePersistent() {
	if mw.sizeBox.Checked() {
		mw.addStyle(win.WS_SIZEBOX)
		return
	}
	mw.removeStyle(^win.WS_SIZEBOX)
}

func (mw *MyWindow) addStyle(style int32) {
	currStyle := win.GetWindowLong(mw.hWnd, win.GWL_STYLE)
	win.SetWindowLong(mw.hWnd, win.GWL_STYLE, currStyle|style)
}

func (mw *MyWindow) removeStyle(style int32) {
	currStyle := win.GetWindowLong(mw.hWnd, win.GWL_STYLE)
	win.SetWindowLong(mw.hWnd, win.GWL_STYLE, currStyle&style)
}

func (mw *MyWindow) SetCloseBox() {
	if mw.closeBox.Checked() {
		win.GetSystemMenu(mw.hWnd, true)
		return
	}
	hMenu := win.GetSystemMenu(mw.hWnd, false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
}

func (mw *MyWindow) AddNotifyIcon() {
	var err error
	mw.ni, err = walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}

	//icon, err := walk.Resources.Image("img/show.ico")
	if err != nil {
		log.Fatal(err)
	}
	//mw.SetIcon(icon)
	//mw.ni.SetIcon(icon)
	mw.ni.SetVisible(true)

	mw.ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button == walk.LeftButton {
			mw.Show()
			win.ShowWindow(mw.Handle(), win.SW_RESTORE)
		}
	})

}

func main() {
	mw := new(MyWindow)
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "notify icon",
		Size:     Size{550, 380},
		Layout:   VBox{MarginsZero: true},
		OnSizeChanged: func() {
			if win.IsIconic(mw.Handle()) {
				mw.Hide()
				mw.ni.SetVisible(true)
			}
		},
		Children: []Widget{
			CheckBox{
				AssignTo:            &mw.minimizeBox,
				Text:                "显示最小化按钮",
				Checked:             true,
				OnCheckStateChanged: mw.SetMinimizeBox,
			},
			CheckBox{
				AssignTo:            &mw.maximizeBox,
				Text:                "显示最大化按钮",
				Checked:             true,
				OnCheckStateChanged: mw.SetMaximizeBox,
			},
			CheckBox{
				AssignTo:            &mw.closeBox,
				Text:                "显示关闭按钮",
				Checked:             true,
				OnCheckStateChanged: mw.SetCloseBox,
			},
			CheckBox{
				AssignTo:            &mw.sizeBox,
				Text:                "允许修改大小",
				Checked:             true,
				OnCheckStateChanged: mw.SetSizePersistent,
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}
	mw.hWnd = mw.Handle()
	mw.AddNotifyIcon()

	mw.Run()
}
