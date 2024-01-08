package encoder

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/aixoio/fsend/packets"
	"github.com/fatih/color"
)

func Start(filename string) {
	filedata, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(color.RedString("%s: %s", err.Error(), filename))
		return
	}

	packet_data := packets.FileData_Packet{
		Name: filename,
		Data: filedata,
	}

	server, err := net.Listen("tcp", ":2328")
	if err != nil {
		fmt.Println(color.RedString("Cannot start TCP/IP server"))
		return
	}

	defer server.Close()

	fmt.Println("TCP/IP server is waiting for connections...")

	conn, err := server.Accept()
	if err != nil {
		fmt.Println(color.RedString("Cannot connect to client"))
		server.Close()
		return
	}

	defer conn.Close()

	fmt.Println("Client connected")

	json_data, err := json.Marshal(packet_data)
	if err != nil {
		fmt.Println(color.RedString("Cannot encode data"))
		conn.Close()
		server.Close()
		return
	}

	size_buff := make([]byte, 4)
	binary.LittleEndian.PutUint32(size_buff, uint32(len(json_data)))
	_, err = conn.Write(size_buff)
	if err != nil {
		fmt.Println(color.RedString("Cannot send data: %s", err.Error()))
		conn.Close()
		server.Close()
		return
	}
	length := len(json_data)
	for i := 0; i < length; i++ {
		_, err = conn.Write([]byte{json_data[i]})
		if err != nil {
			fmt.Println(color.RedString("Cannot send data: %s", err.Error()))
			conn.Close()
			server.Close()
			return
		}
		fmt.Printf("%.4f%% - %d/%d\n", ((float64((i + 1)) / float64(length)) * 100), (i + 1), length)
	}

	fmt.Println("Data sent")

}
