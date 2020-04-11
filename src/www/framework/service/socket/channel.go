package SocketService

var Channel = &SocketChannel{}

type SocketChannel struct {
	Channel SocketMessage `json:"channel"`
}

type SocketMessage struct {
	MessageType 	string 			`json:"message_type"`
	SystemMessage 	systemMessage 	`json:"system_message"`
	SerialMessage 	serialMessage 	`json:"serial_message"`
	RemoteMessage	remoteMessage 	`json:"remote_message"`
	ServiceMessage	serviceMessage 	`json:"service_message"`
	TcpMessage		tcpMessage 	`json:"tcp_message"`
}

type systemMessage struct {
	Time    string 		`json:"time"`
	Memory 	float64 	`json:"memory"`
	Cpu 	[]float64 	`json:"cpu"`
}

type serialMessage struct {
	Port    string 		`json:"port"`
	Content string 		`json:"content"`
}

type remoteMessage struct {
	Content string 		`json:"content"`
}

type serviceMessage struct {
	Content string 		`json:"content"`
}

type tcpMessage struct {
	Content string 		`json:"content"`
}