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
		fmt.Println(color.RedString("Cannot read data"))
		conn.Close()
		return
	}

	size := binary.LittleEndian.Uint32(size_buffer)

	data_buffer := make([]byte, size)
	_, err = conn.Read(data_buffer)
	if err != nil {
		fmt.Println(color.RedString("Cannot read data"))
		conn.Close()
		return
	}

	var data packets.FileData_Packet
	err = json.Unmarshal(data_buffer, &data)
	if err != nil {
		fmt.Println(color.RedString("Cannot parse data"))
		conn.Close()
		return
	}

	err = os.WriteFile(data.Name, data.Data, 0644)
	if err != nil {
		fmt.Println(color.RedString("%s: %s", err.Error(), ip))
	}

}
