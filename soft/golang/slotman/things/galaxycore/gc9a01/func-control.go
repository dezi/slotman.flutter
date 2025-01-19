package gc9a01

func (se *GC9A01) Initialize() (err error) {

	//_ = se.writeCommandBytes(0xEF)
	//_ = se.writeCommandBytes(0xEB, 0x14)
	//_ = se.writeCommandBytes(0xFE)
	//_ = se.writeCommandBytes(0xEF)
	//_ = se.writeCommandBytes(0xEB, 0x14)
	//_ = se.writeCommandBytes(0x84, 0x40)
	//_ = se.writeCommandBytes(0x85, 0xFF)
	//_ = se.writeCommandBytes(0x86, 0xFF)
	//_ = se.writeCommandBytes(0x87, 0xFF)
	//_ = se.writeCommandBytes(0x88, 0x0A)
	//_ = se.writeCommandBytes(0x89, 0x21)
	//_ = se.writeCommandBytes(0x8A, 0x00)
	//_ = se.writeCommandBytes(0x8B, 0x80)
	//_ = se.writeCommandBytes(0x8C, 0x01)
	//_ = se.writeCommandBytes(0x8D, 0x01)
	//_ = se.writeCommandBytes(0x8E, 0xFF)
	//_ = se.writeCommandBytes(0x8F, 0xFF)
	//_ = se.writeCommandBytes(0xB6, 0x00, 0x00)
	//_ = se.writeCommandBytes(0x36, 0x18)
	//_ = se.writeCommandBytes(COLOR_MODE, COLOR_MODE_18_BIT)
	//_ = se.writeCommandBytes(0x90, 0x08, 0x08, 0x08, 0x08)
	//_ = se.writeCommandBytes(0xBD, 0x06)
	//_ = se.writeCommandBytes(0xBC, 0x00)
	//_ = se.writeCommandBytes(0xFF, 0x60, 0x01, 0x04)
	//_ = se.writeCommandBytes(0xC3, 0x13)
	//_ = se.writeCommandBytes(0xC4, 0x13)
	//_ = se.writeCommandBytes(0xC9, 0x22)
	//_ = se.writeCommandBytes(0xBE, 0x11)
	//_ = se.writeCommandBytes(0xE1, 0x10, 0x0E)
	//_ = se.writeCommandBytes(0xDF, 0x21, 0x0c, 0x02)
	//_ = se.writeCommandBytes(0xF0, 0x45, 0x09, 0x08, 0x08, 0x26, 0x2A)
	//_ = se.writeCommandBytes(0xF1, 0x43, 0x70, 0x72, 0x36, 0x37, 0x6F)
	//_ = se.writeCommandBytes(0xF2, 0x45, 0x09, 0x08, 0x08, 0x26, 0x2A)
	//_ = se.writeCommandBytes(0xF3, 0x43, 0x70, 0x72, 0x36, 0x37, 0x6F)
	//_ = se.writeCommandBytes(0xED, 0x1B, 0x0B)
	//_ = se.writeCommandBytes(0xAE, 0x77)
	//_ = se.writeCommandBytes(0xCD, 0x63)
	//_ = se.writeCommandBytes(0x70, 0x07, 0x07, 0x04, 0x0E, 0x0F, 0x09, 0x07, 0x08, 0x03)
	//_ = se.writeCommandBytes(0xE8, 0x34)
	//_ = se.writeCommandBytes(0x62, 0x18, 0x0D, 0x71, 0xED, 0x70, 0x70, 0x18, 0x0F, 0x71, 0xEF, 0x70, 0x70)
	//_ = se.writeCommandBytes(0x63, 0x18, 0x11, 0x71, 0xF1, 0x70, 0x70, 0x18, 0x13, 0x71, 0xF3, 0x70, 0x70)
	//_ = se.writeCommandBytes(0x64, 0x28, 0x29, 0xF1, 0x01, 0xF1, 0x00, 0x07)
	//_ = se.writeCommandBytes(0x66, 0x3C, 0x00, 0xCD, 0x67, 0x45, 0x45, 0x10, 0x00, 0x00, 0x00)
	//_ = se.writeCommandBytes(0x67, 0x00, 0x3C, 0x00, 0x00, 0x00, 0x01, 0x54, 0x10, 0x32, 0x98)
	//_ = se.writeCommandBytes(0x74, 0x10, 0x85, 0x80, 0x00, 0x00, 0x4E, 0x00)
	//_ = se.writeCommandBytes(0x98, 0x3e, 0x07)
	//_ = se.writeCommandBytes(0x35)
	//_ = se.writeCommandBytes(0x21)
	//_ = se.writeCommandBytes(0x11)
	//_ = se.writeCommandBytes(0x29)

	_ = se.writeCommand(0xEF)

	_ = se.writeCommand(0xEB)
	_ = se.writeByte(0x14)

	_ = se.writeCommand(0xFE)
	_ = se.writeCommand(0xEF)

	_ = se.writeCommand(0xEB)
	_ = se.writeByte(0x14)

	_ = se.writeCommand(0x84)
	_ = se.writeByte(0x40)

	_ = se.writeCommand(0x85)
	_ = se.writeByte(0xFF)

	_ = se.writeCommand(0x86)
	_ = se.writeByte(0xFF)

	_ = se.writeCommand(0x87)
	_ = se.writeByte(0xFF)

	_ = se.writeCommand(0x88)
	_ = se.writeByte(0x0A)

	_ = se.writeCommand(0x89)
	_ = se.writeByte(0x21)

	_ = se.writeCommand(0x8A)
	_ = se.writeByte(0x00)

	_ = se.writeCommand(0x8B)
	_ = se.writeByte(0x80)

	_ = se.writeCommand(0x8C)
	_ = se.writeByte(0x01)

	_ = se.writeCommand(0x8D)
	_ = se.writeByte(0x01)

	_ = se.writeCommand(0x8E)
	_ = se.writeByte(0xFF)

	_ = se.writeCommand(0x8F)
	_ = se.writeByte(0xFF)

	_ = se.writeCommand(0xB6)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x00)

	_ = se.writeCommand(0x36)
	_ = se.writeByte(0x18)

	_ = se.writeCommand(COLOR_MODE)
	_ = se.writeByte(COLOR_MODE_18_BIT)

	_ = se.writeCommand(0x90)
	_ = se.writeByte(0x08)
	_ = se.writeByte(0x08)
	_ = se.writeByte(0x08)
	_ = se.writeByte(0x08)

	_ = se.writeCommand(0xBD)
	_ = se.writeByte(0x06)

	_ = se.writeCommand(0xBC)
	_ = se.writeByte(0x00)

	_ = se.writeCommand(0xFF)
	_ = se.writeByte(0x60)
	_ = se.writeByte(0x01)
	_ = se.writeByte(0x04)

	_ = se.writeCommand(0xC3)
	_ = se.writeByte(0x13)
	_ = se.writeCommand(0xC4)
	_ = se.writeByte(0x13)

	_ = se.writeCommand(0xC9)
	_ = se.writeByte(0x22)

	_ = se.writeCommand(0xBE)
	_ = se.writeByte(0x11)

	_ = se.writeCommand(0xE1)
	_ = se.writeByte(0x10)
	_ = se.writeByte(0x0E)

	_ = se.writeCommand(0xDF)
	_ = se.writeByte(0x21)
	_ = se.writeByte(0x0c)
	_ = se.writeByte(0x02)

	_ = se.writeCommand(0xF0)
	_ = se.writeByte(0x45)
	_ = se.writeByte(0x09)
	_ = se.writeByte(0x08)
	_ = se.writeByte(0x08)
	_ = se.writeByte(0x26)
	_ = se.writeByte(0x2A)

	_ = se.writeCommand(0xF1)
	_ = se.writeByte(0x43)
	_ = se.writeByte(0x70)
	_ = se.writeByte(0x72)
	_ = se.writeByte(0x36)
	_ = se.writeByte(0x37)
	_ = se.writeByte(0x6F)

	_ = se.writeCommand(0xF2)
	_ = se.writeByte(0x45)
	_ = se.writeByte(0x09)
	_ = se.writeByte(0x08)
	_ = se.writeByte(0x08)
	_ = se.writeByte(0x26)
	_ = se.writeByte(0x2A)

	_ = se.writeCommand(0xF3)
	_ = se.writeByte(0x43)
	_ = se.writeByte(0x70)
	_ = se.writeByte(0x72)
	_ = se.writeByte(0x36)
	_ = se.writeByte(0x37)
	_ = se.writeByte(0x6F)

	_ = se.writeCommand(0xED)
	_ = se.writeByte(0x1B)
	_ = se.writeByte(0x0B)

	_ = se.writeCommand(0xAE)
	_ = se.writeByte(0x77)

	_ = se.writeCommand(0xCD)
	_ = se.writeByte(0x63)

	_ = se.writeCommand(0x70)
	_ = se.writeByte(0x07)
	_ = se.writeByte(0x07)
	_ = se.writeByte(0x04)
	_ = se.writeByte(0x0E)
	_ = se.writeByte(0x0F)
	_ = se.writeByte(0x09)
	_ = se.writeByte(0x07)
	_ = se.writeByte(0x08)
	_ = se.writeByte(0x03)

	_ = se.writeCommand(0xE8)
	_ = se.writeByte(0x34)

	_ = se.writeCommand(0x62)
	_ = se.writeByte(0x18)
	_ = se.writeByte(0x0D)
	_ = se.writeByte(0x71)
	_ = se.writeByte(0xED)
	_ = se.writeByte(0x70)
	_ = se.writeByte(0x70)
	_ = se.writeByte(0x18)
	_ = se.writeByte(0x0F)
	_ = se.writeByte(0x71)
	_ = se.writeByte(0xEF)
	_ = se.writeByte(0x70)
	_ = se.writeByte(0x70)

	_ = se.writeCommand(0x63)
	_ = se.writeByte(0x18)
	_ = se.writeByte(0x11)
	_ = se.writeByte(0x71)
	_ = se.writeByte(0xF1)
	_ = se.writeByte(0x70)
	_ = se.writeByte(0x70)
	_ = se.writeByte(0x18)
	_ = se.writeByte(0x13)
	_ = se.writeByte(0x71)
	_ = se.writeByte(0xF3)
	_ = se.writeByte(0x70)
	_ = se.writeByte(0x70)

	_ = se.writeCommand(0x64)
	_ = se.writeByte(0x28)
	_ = se.writeByte(0x29)
	_ = se.writeByte(0xF1)
	_ = se.writeByte(0x01)
	_ = se.writeByte(0xF1)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x07)

	_ = se.writeCommand(0x66)
	_ = se.writeByte(0x3C)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0xCD)
	_ = se.writeByte(0x67)
	_ = se.writeByte(0x45)
	_ = se.writeByte(0x45)
	_ = se.writeByte(0x10)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x00)

	_ = se.writeCommand(0x67)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x3C)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x01)
	_ = se.writeByte(0x54)
	_ = se.writeByte(0x10)
	_ = se.writeByte(0x32)
	_ = se.writeByte(0x98)

	_ = se.writeCommand(0x74)
	_ = se.writeByte(0x10)
	_ = se.writeByte(0x85)
	_ = se.writeByte(0x80)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x00)
	_ = se.writeByte(0x4E)
	_ = se.writeByte(0x00)

	_ = se.writeCommand(0x98)
	_ = se.writeByte(0x3e)
	_ = se.writeByte(0x07)

	_ = se.writeCommand(0x35)
	_ = se.writeCommand(0x21)

	_ = se.writeCommand(0x11)
	//time.Sleep(time.Millisecond * 120)

	_ = se.writeCommand(0x29)
	//time.Sleep(time.Millisecond * 20)

	return
}
