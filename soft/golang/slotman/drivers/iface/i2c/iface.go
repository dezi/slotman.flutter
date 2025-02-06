package i2c

type I2c interface {
	Open() (err error)
	Close() (err error)

	GetDevice() (device string)
	GetAddr() (addr uint8)

	Write(data []byte) (xfer int, err error)
	Read(data []byte) (xfer int, err error)

	WriteBytes(data []byte) (xfer int, err error)
	ReadBytes(data []byte) (xfer int, err error)

	ReadRegByte(reg byte) (value byte, err error)
	ReadRegBytes(reg byte, n int) (data []byte, xfer int, err error)

	WriteReg(reg byte) (err error)
	WriteRegByte(reg byte, value byte) (err error)
	WriteRegBytes(reg byte, data []byte) (err error)

	ReadRegU16BE(reg byte) (value uint16, err error)
	ReadRegU16LE(reg byte) (value uint16, err error)
	ReadRegS16BE(reg byte) (value int16, err error)
	ReadRegS16LE(reg byte) (value int16, err error)

	WriteRegU16BE(reg byte, value uint16) (err error)
	WriteRegU16LE(reg byte, value uint16) (err error)
	WriteRegS16BE(reg byte, value int16) (err error)
	WriteRegS16LE(reg byte, value int16) (err error)
}
