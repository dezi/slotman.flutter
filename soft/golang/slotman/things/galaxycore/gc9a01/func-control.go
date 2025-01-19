package gc9a01

import "errors"

func (se *GC9A01) BlipFullRawImage(image []byte) (err error) {

	if len(image) != screenWidth*screenHeight*3 {
		err = errors.New("invalid image size")
		return
	}

	err = gc9a01.SetFrame(Frame{X0: 0, Y0: 0, X1: screenWidth - 1, Y1: screenHeight - 1})
	if err != nil {
		return
	}

	chunkSize := screenWidth * 4 * 3

	for chunkPos := 0; chunkPos < len(image); chunkPos += chunkSize {
		if chunkPos == 0 {
			_ = gc9a01.writeMem(image[chunkPos : chunkPos+chunkSize])
		} else {
			_ = gc9a01.writeMemCont(image[chunkPos : chunkPos+chunkSize])
		}
	}

	return
}

func (se *GC9A01) SetOrientation(orientation int) (err error) {

	switch orientation {
	case 0:
		err = se.writeCommandBytes(0x36, 0x18)
	case 1:
		err = se.writeCommandBytes(0x36, 0x28)
	case 2:
		err = se.writeCommandBytes(0x36, 0x48)
	case 3:
		err = se.writeCommandBytes(0x36, 0x88)
	default:
		err = errors.New("wrong orientation")
	}

	return
}

func (se *GC9A01) SetFrame(frame Frame) (err error) {

	var data [4]byte

	data[0] = byte(frame.X0 >> 8)
	data[1] = byte(frame.X0)
	data[2] = byte(frame.X1 >> 8)
	data[3] = byte(frame.X1)

	_ = se.writeCommandBytes(COL_ADDR_SET, data[:]...)

	data[0] = byte(frame.Y0 >> 8)
	data[1] = byte(frame.Y0)
	data[2] = byte(frame.Y1 >> 8)
	data[3] = byte(frame.Y1)

	_ = se.writeCommandBytes(ROW_ADDR_SET, data[:]...)

	return
}

func (se *GC9A01) Initialize() (err error) {

	_ = se.writeCommandBytes(0xEF)
	_ = se.writeCommandBytes(0xEB, 0x14)
	_ = se.writeCommandBytes(0xFE)
	_ = se.writeCommandBytes(0xEF)
	_ = se.writeCommandBytes(0xEB, 0x14)
	_ = se.writeCommandBytes(0x84, 0x40)
	_ = se.writeCommandBytes(0x85, 0xFF)
	_ = se.writeCommandBytes(0x86, 0xFF)
	_ = se.writeCommandBytes(0x87, 0xFF)
	_ = se.writeCommandBytes(0x88, 0x0A)
	_ = se.writeCommandBytes(0x89, 0x21)
	_ = se.writeCommandBytes(0x8A, 0x00)
	_ = se.writeCommandBytes(0x8B, 0x80)
	_ = se.writeCommandBytes(0x8C, 0x01)
	_ = se.writeCommandBytes(0x8D, 0x01)
	_ = se.writeCommandBytes(0x8E, 0xFF)
	_ = se.writeCommandBytes(0x8F, 0xFF)
	_ = se.writeCommandBytes(0xB6, 0x00, 0x00)
	_ = se.writeCommandBytes(0x36, 0x18)
	_ = se.writeCommandBytes(COLOR_MODE, COLOR_MODE_18_BIT)
	_ = se.writeCommandBytes(0x90, 0x08, 0x08, 0x08, 0x08)
	_ = se.writeCommandBytes(0xBD, 0x06)
	_ = se.writeCommandBytes(0xBC, 0x00)
	_ = se.writeCommandBytes(0xFF, 0x60, 0x01, 0x04)
	_ = se.writeCommandBytes(0xC3, 0x13)
	_ = se.writeCommandBytes(0xC4, 0x13)
	_ = se.writeCommandBytes(0xC9, 0x22)
	_ = se.writeCommandBytes(0xBE, 0x11)
	_ = se.writeCommandBytes(0xE1, 0x10, 0x0E)
	_ = se.writeCommandBytes(0xDF, 0x21, 0x0c, 0x02)
	_ = se.writeCommandBytes(0xF0, 0x45, 0x09, 0x08, 0x08, 0x26, 0x2A)
	_ = se.writeCommandBytes(0xF1, 0x43, 0x70, 0x72, 0x36, 0x37, 0x6F)
	_ = se.writeCommandBytes(0xF2, 0x45, 0x09, 0x08, 0x08, 0x26, 0x2A)
	_ = se.writeCommandBytes(0xF3, 0x43, 0x70, 0x72, 0x36, 0x37, 0x6F)
	_ = se.writeCommandBytes(0xED, 0x1B, 0x0B)
	_ = se.writeCommandBytes(0xAE, 0x77)
	_ = se.writeCommandBytes(0xCD, 0x63)
	_ = se.writeCommandBytes(0x70, 0x07, 0x07, 0x04, 0x0E, 0x0F, 0x09, 0x07, 0x08, 0x03)
	_ = se.writeCommandBytes(0xE8, 0x34)
	_ = se.writeCommandBytes(0x62, 0x18, 0x0D, 0x71, 0xED, 0x70, 0x70, 0x18, 0x0F, 0x71, 0xEF, 0x70, 0x70)
	_ = se.writeCommandBytes(0x63, 0x18, 0x11, 0x71, 0xF1, 0x70, 0x70, 0x18, 0x13, 0x71, 0xF3, 0x70, 0x70)
	_ = se.writeCommandBytes(0x64, 0x28, 0x29, 0xF1, 0x01, 0xF1, 0x00, 0x07)
	_ = se.writeCommandBytes(0x66, 0x3C, 0x00, 0xCD, 0x67, 0x45, 0x45, 0x10, 0x00, 0x00, 0x00)
	_ = se.writeCommandBytes(0x67, 0x00, 0x3C, 0x00, 0x00, 0x00, 0x01, 0x54, 0x10, 0x32, 0x98)
	_ = se.writeCommandBytes(0x74, 0x10, 0x85, 0x80, 0x00, 0x00, 0x4E, 0x00)
	_ = se.writeCommandBytes(0x98, 0x3e, 0x07)
	_ = se.writeCommandBytes(0x35)
	_ = se.writeCommandBytes(0x21)
	_ = se.writeCommandBytes(0x11)
	_ = se.writeCommandBytes(0x29)

	return
}
