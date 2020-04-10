package Robot

var Init Interface

type Interface interface {
	OnStart()
	OnClose()
	OnMessages([]byte, string)
}

type Base struct {}

func (base *Base) OnStart() {}

func (base *Base) OnClose() {}

func (base *Base) OnMessages(byte []byte, string string) {}