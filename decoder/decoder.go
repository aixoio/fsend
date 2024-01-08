package decoder

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/aixoio/fsend/packets"
	"github.com/fatih/color"
)

func Start(ip string) {
	conn, err := net.Dial("tcp", ip + ":2328")
	if err != nil {
		fmt.Println(color.RedString("%s: %s", err.Error(), ip))
		return
	}

	defer conn.Close()

	size_buffer := make([]byte, 4)

	_, err = conn.Read(size_buffer)
	if err != nil {
		fmt.Println(color.RedString("Cannot read data: %s", err.Error()))
		conn.Close()
		return
	}

	size := binary.LittleEndian.Uint32(size_buffer)

	data_buffer := make([]byte, size)
	byte_buff := make([]byte, 1)
	for i := 0; i < int(size); i++ {
		_, err = conn.Read(byte_buff)
		if err != nil {
			fmt.Println(color.RedString("Cannot read data: %s", err.Error()))
			conn.Close()
			return
		}
		data_buffer[i] = byte_buff[0]
		fmt.Printf("%.2f%% - %d/%d\n", ((float64(size) / float64((i + 1))) * 100), (i + 1), size)
	}

	var data packets.FileData_Packet
	err = json.Unmarshal(data_buffer, &data)
	if err != nil {
		fmt.Println(color.RedString("Cannot parse data: %s", err.Error()))
		conn.Close()
		return
	}

	err = os.WriteFile(data.Name, data.Data, 0644)
	if err != nil {
		fmt.Println(color.RedString("%s: %s", err.Error(), ip))
	}

}
