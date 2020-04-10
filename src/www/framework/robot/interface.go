package Robot

var Init Interface

type Interface interface {
	OnStart()
	OnClose()
	OnReMessages([]byte, string)
}

type Base struct {}

func (base *Base) OnStart() {}

func (base *Base) OnClose() {}

func (base *Base) OnReMessages(byte []byte, string string) {}