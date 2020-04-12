package SocketService

var Channel = &SocketChannel{}

type SocketChannel struct {
	Channel SocketMessage `json:"channel"`
}

type SocketMessage struct {
	MessageType 	string 			`json:"message_type"`
	SerialMessage 	serialMessage 	`json:"serial_message"`
	RemoteMessage	remoteMessage 	`json:"remote_message"`
	ServiceMessage	serviceMessage 	`json:"service_message"`
	TcpMessage		tcpMessage 	`json:"tcp_message"`
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