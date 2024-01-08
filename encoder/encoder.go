package encoder

import (
	"fmt"
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

	
}
