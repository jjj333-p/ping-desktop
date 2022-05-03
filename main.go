package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type messageObj struct {
	content, name                                     *string
	renderedContent, renderedName, renderedMessageObj fyne.CanvasObject
	timestamp                                         time.Time
}

//returns msg object configured with the input data
func NewMessageObject(content string, authorname string) messageObj {

	//create rendered text
	rn := widget.NewLabel(authorname)
	rc := widget.NewLabel(content)

	//export all the configured data feilds
	return messageObj{
		content:            &content,
		name:               &authorname,
		renderedContent:    rc,
		renderedName:       rn,
		timestamp:          time.Now().UTC(),
		renderedMessageObj: container.New(layout.NewBorderLayout(rn, nil, nil, nil), rn, rc),
	}

}

//struct to pair chatUI to its id
type chatUI struct {
	id             *int
	ui, sendbutton fyne.CanvasObject
	renderedmsgs   *fyne.Container
	msgbox         *container.Scroll
	msgs           []messageObj
	mainentry      *widget.Entry
}

//handling new messages
func (chat chatUI) append(content string, authorname string) {

	//make object from what is currently in the text entry box
	msg := NewMessageObject(content, authorname)

	//clear the text input box
	chat.mainentry.SetText("")

	//add msg object to the array of msgs
	chat.msgs = append(chat.msgs, msg)

	//add msg to rendered obj
	chat.renderedmsgs.Add(msg.renderedMessageObj)

	//scroll box to bottom to see new msg
	chat.msgbox.ScrollToBottom()

}

//creat the chat box for a particular chat
func NewChatUI(id int) chatUI {

	//create object and add the id to it.
	chat := chatUI{
		id: &id,
	}

	//create typing box
	chat.mainentry = widget.NewMultiLineEntry()

	//that background text thingy before you type
	chat.mainentry.SetPlaceHolder("eeee...")

	//create object for the rendered messages
	chat.renderedmsgs = container.New(layout.NewVBoxLayout())

	//add scroll to the rendered text
	chat.msgbox = container.NewVScroll(chat.renderedmsgs)

	//array to store msgs
	chat.msgs = []messageObj{}

	//button to send message in mainentry, and add message to chat on send button
	chat.sendbutton = widget.NewButton("send ^", func() { chat.append(chat.mainentry.Text, "you") })

	//put together entry and send button
	input := container.NewBorder(nil, nil, nil, chat.sendbutton, chat.mainentry)

	//attaching entry and button to the msgbox
	chat.ui = container.NewBorder(nil, input, nil, nil, chat.msgbox)

	//return out object to use.
	return chat

}

func main() {

	//make objects for the app
	a := app.New()
	w := a.NewWindow("pain")

	divider := container.NewHSplit(widget.NewTextGridFromString("aaaaaaa"), NewChatUI(0).ui)

	w.SetContent(divider)

	//default size of the window
	w.Resize(fyne.NewSize(100, 100))

	//just posting
	w.ShowAndRun()

}
