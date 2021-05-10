package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

var PORT string="12345"
var IP string
//继承MainWindow指针，得到一个自定义窗体
type MyMainWindow struct {
	*walk.MainWindow
}
type Message struct {
	Name          string
	Domesticated  bool
	Remarks       string
}

var chMsg=make(chan string,1)
//var chListen=make(chan net.Listener,1)
func main() {
	// 创建自定义窗体指针
	mw := new(MyMainWindow)
	// 滑动条
	var slvx, slhx *walk.Slider
	var slvy, slhy *walk.Slider
	var maxEditx, minEditx, valueEditx *walk.NumberEdit
	var maxEdity, minEdity, valueEdity *walk.NumberEdit
	// 输出框
	var outTE1,outTE2,outTE3 *walk.TextEdit
	// 滑动条数据
	data := struct{ Min, Max, Value int }{0, 100, 30}
	// 编辑
	message := new(Message)
	message2 := new(Message)


	helpTipMsg:="Debugging ……\r\n\r\n"
	// 主视图


	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "中央任务调度系统",
		Size: Size{
			Width:  800,
			Height: 450,
		},
		MinSize:  Size{400, 200},
		// 排版格式
		Layout:   Grid{Columns: 3,Alignment: AlignHCenterVFar},
		MenuItems: []MenuItem{
			Menu{
				Text: "& 菜单 ",
				Items: []MenuItem{
					Action{
						Text:    "打开服务端",
						OnTriggered: func() {
							cmdUTF:= exec.Command("chcp","65001")
							cmdUTF.Run()
							cmd := exec.Command("ipconfig")
							out, err := cmd.Output()
							if err != nil {
								fmt.Println(err)
							}
							str:=string(out)
							if strings.Contains(str,"WLAN") {
								str=str[strings.Index(str,"WLAN"):]
								if  strings.Contains(str,"IPv4"){
									str=str[strings.Index(str,"192.168"):]
									IP=strings.Trim(str[:15]," \r\n")
								}else {
									IP="err"
								}
							}
							if IP=="err"{
								walk.MsgBox(mw,"警告","请检查局域网IP地址，确认已连接WIFI",walk.MsgBoxIconError)
							}else {
								tcpServerOpen()
							}
						},
					},
					Separator{},
					Action{
						Text:        "E&xit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text: "About",
						OnTriggered: func() {
							walk.MsgBox(mw, "About", helpTipMsg, walk.MsgBoxIconInformation)
						},
					},
				},
			},
		},
		Children: []Widget{
			Label{Text: "TUDO",ColumnSpan: 3},// 第二个参数占位
			Label{Text: "参数1:"},
			TextEdit{
				MaxSize: Size{
					Width:  30,
					Height: 20,
				},
				AssignTo: &outTE1,
				ReadOnly: true,
				Text:     "",

			},
			PushButton{
				Text: "编辑参数1",
				OnClicked: func() {
					if len(chMsg)!=0{
						outTE1.SetText(<-chMsg)
					}else {
						outTE1.SetText("")
					}
				},
			},
			//
			Label{Text: "参数2:"},
			TextEdit{
				MaxSize: Size{Width: 30,Height: 20},
				AssignTo: &outTE2,
				ReadOnly: true,
				Text:     fmt.Sprintf("%+v", message),
			},
			PushButton{
				Text: "编辑参数2",
				OnClicked: func() {
					if cmd, err := RunAnimalDialog(mw, message); err != nil {
						log.Print(err)
					} else if cmd == walk.DlgCmdOK {
						outTE2.SetText(fmt.Sprintf("%+v", message))
					}
				},
			},

			Label{Text: "参数3"},
			TextEdit{
				MaxSize: Size{Height: 20,Width: 300},
				AssignTo: &outTE3,
				ReadOnly: true,
				Text: "",
			},
			PushButton{
				Text: "编辑参数3",
				OnClicked: func() {
					if cmd, err := RunAnimalDialog(mw, message2); err != nil {
						log.Print(err)
					} else if cmd == walk.DlgCmdOK {
						outTE3.SetText(fmt.Sprintln(message2))
					}
				},
			},
			//x
			Slider{
				AssignTo:    &slvx,
				MinValue:    data.Min,
				MaxValue:    data.Max,
				Value:       data.Value,
				Orientation: Vertical,
				OnValueChanged: func() {
					data.Value = slvx.Value()
					valueEditx.SetValue(float64(data.Value))

				},
			},
			Composite{
				Layout:        Grid{Columns: 3},
				StretchFactor: 4,
				Children: []Widget{
					Label{Text: "Min x"},
					Label{Text: "x"},
					Label{Text: "Max x"},
					NumberEdit{
						AssignTo: &minEditx,
						Value:    float64(data.Min),
						OnValueChanged: func() {
							data.Min = int(minEditx.Value())
							slhx.SetRange(data.Min, data.Max)
							slvx.SetRange(data.Min, data.Max)
						},
					},
					NumberEdit{
						AssignTo: &valueEditx,
						Value:    float64(data.Value),
						OnValueChanged: func() {
							data.Value = int(valueEditx.Value())
							slhx.SetValue(data.Value)
							slvx.SetValue(data.Value)
						},
					},
					NumberEdit{
						AssignTo: &maxEditx,
						Value:    float64(data.Max),
						OnValueChanged: func() {
							data.Max = int(maxEditx.Value())
							slhx.SetRange(data.Min, data.Max)
							slvx.SetRange(data.Min, data.Max)
						},
					},
					Slider{
						ColumnSpan: 3,
						AssignTo:   &slhx,
						MinValue:   data.Min,
						MaxValue:   data.Max,
						Value:      data.Value,
						OnValueChanged: func() {
							data.Value = slhx.Value()
							valueEditx.SetValue(float64(data.Value))
						},
					},

					VSpacer{},

				},
			},
			//y
			Slider{
				AssignTo:    &slvy,
				MinValue:    data.Min,
				MaxValue:    data.Max,
				Value:       data.Value,
				Orientation: Vertical,
				OnValueChanged: func() {
					data.Value = slvy.Value()
					valueEdity.SetValue(float64(data.Value))

				},
			},
			Composite{
				Layout:        Grid{Columns: 3},
				StretchFactor: 4,
				Children: []Widget{
					Label{Text: "Min y"},
					Label{Text: "y"},
					Label{Text: "Max y"},
					NumberEdit{
						AssignTo: &minEdity,
						Value:    float64(data.Min),
						OnValueChanged: func() {
							data.Min = int(minEdity.Value())
							slhy.SetRange(data.Min, data.Max)
							slvy.SetRange(data.Min, data.Max)
						},
					},
					NumberEdit{
						AssignTo: &valueEdity,
						Value:    float64(data.Value),
						OnValueChanged: func() {
							data.Value = int(valueEdity.Value())
							slhy.SetValue(data.Value)
							slvy.SetValue(data.Value)
						},
					},
					NumberEdit{
						AssignTo: &maxEdity,
						Value:    float64(data.Max),
						OnValueChanged: func() {
							data.Max = int(maxEdity.Value())
							slhy.SetRange(data.Min, data.Max)
							slvy.SetRange(data.Min, data.Max)
						},
					},
					Slider{
						ColumnSpan: 3,
						AssignTo:   &slhy,
						MinValue:   data.Min,
						MaxValue:   data.Max,
						Value:      data.Value,
						OnValueChanged: func() {
							data.Value = slhy.Value()
							valueEdity.SetValue(float64(data.Value))
						},
					},
					VSpacer{},
					PushButton{
						ColumnSpan: 3,
						Text:       "打印日志",
						OnClicked: func() {
							log.Printf("H: < %d | %d | %d >\n", slhx.MinValue(), slhx.Value(), slhx.MaxValue())
							log.Printf("V: < %d | %d | %d >\n", slvx.MinValue(), slvx.Value(), slvx.MaxValue())

							log.Printf("H: < %d | %d | %d >\n", slhy.MinValue(), slhy.Value(), slhy.MaxValue())
							log.Printf("V: < %d | %d | %d >\n", slvy.MinValue(), slvy.Value(), slvy.MaxValue())
						},
					},
				},
			},
		},

	}.Run()); err != nil {
		log.Fatal(err)
	}

}

func RunAnimalDialog(owner walk.Form, message *Message) (int, error) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var acceptPB, cancelPB *walk.PushButton

	return Dialog{
		AssignTo:      &dlg,
		Title:         Bind("'编辑信息' + (message.Name == '' ? '' : ' - ' + message.Name)"),
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "message",
			DataSource:     message,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		MinSize: Size{300, 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "Name:",
					},
					LineEdit{
						Text: Bind("Name"),
					},

					Label{
						Text: "Domesticated:",
					},
					CheckBox{
						Checked: Bind("Domesticated"),
					},

					VSpacer{
						ColumnSpan: 2,
						Size:       8,
					},

					Label{
						ColumnSpan: 2,
						Text:       "Remarks:",
					},
					TextEdit{
						ColumnSpan: 2,
						MinSize:    Size{100, 50},
						Text:       Bind("Remarks"),
					},

				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}
							//animalPl(animal)
							dlg.Accept()
						},
					},
					PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)
}

// 使用 golang 关键字go开启协程
func tcpServerOpen()  {
	go func() {
		// 服务器在指定端口建立 tcp 监听
		listener,err:=net.Listen("tcp",IP+":"+PORT)
		if err != nil {
			fmt.Println(err, "net.listen()")
			time.Sleep(10*time.Second)
			os.Exit(1)
		}
		// 循环接入所有客户端
		for  {
			fmt.Printf("服务端%s:%s已开启\n",IP,PORT)
			conn,e := listener.Accept()
			if e != nil {
				fmt.Println(e, "listener.Accept()")
				time.Sleep(10*time.Second)
				os.Exit(1)
			}
			// 开协程与当前客户端连接
			go ChatWith(conn)
			break
		}
	}()
}

func ChatWith(conn net.Conn)  {
	// 创建消息缓冲区
	buffer:=make([]byte,1024)
	for  {
		n,err:=conn.Read(buffer)
		if err != nil {
			fmt.Println(err, "conn.Read(buffer)")
			os.Exit(1)
		}
		// 转换为字符串输出
		clientMsg:=string(buffer[0:n])
		fmt.Printf("收到%v的消息:%s\n",conn.RemoteAddr(),clientMsg)

		chMsg<-clientMsg
		if clientMsg!="im off" {
			conn.Write([]byte("已阅 "+clientMsg))
		}else{
			conn.Write([]byte("bye!"))
			break
		}
	}
}