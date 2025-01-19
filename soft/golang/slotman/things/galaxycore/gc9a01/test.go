package gc9a01

import (
	"math/rand"
	"slotman/utils/log"
	"time"
)

var gc9a01 *GC9A01

func TestDisplay() {

	log.Printf("Display GC8A01 test started...")

	gc9a01 = NewGC9A01("/dev/spidev0.0", 25)

	err := gc9a01.OpenSensor()
	if err != nil {
		return
	}

	log.Printf("Display GC8A01 device SPI0-0 opened.")

	_ = gc9a01.Initialize()

	_ = gc9a01.SetFrame(Frame{
		X0: 0,
		Y0: 0,
		X1: 239,
		Y1: 239,
	})

	log.Printf("Display GC8A01 test patterns.")

	color := make([]byte, 3)
	line := make([]byte, 3*240*240)

	for {

		color[0] = byte(rand.Int31())
		color[1] = byte(rand.Int31())

		//for x := 0; x < 240; x++ {
		//	for y := 0; y < 240; y++ {
		//		if x < y {
		//			color[2] = 0xFF
		//		} else {
		//			color[2] = 0x00
		//		}
		//		if x == 0 && y == 0 {
		//			_ = gc9a01.WriteMem(color)
		//		} else {
		//			_ = gc9a01.WriteMemCont(color)
		//		}
		//	}
		//}

		off := 0

		for x := 0; x < 240; x++ {

			for y := 0; y < 240; y++ {
				if x < y {
					color[2] = 0xFF
				} else {
					color[2] = 0x00
				}

				line[off] = color[0]
				off++
				line[off] = color[1]
				off++
				line[off] = color[2]
				off++
			}

			//if x == 0 {
			//	_ = gc9a01.WriteMem(line)
			//} else {
			//	_ = gc9a01.WriteMemCont(line)
			//}
		}

		_ = gc9a01.WriteMem(line)

		time.Sleep(time.Millisecond * 250)
	}
}
