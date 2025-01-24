package proxy

import (
	"encoding/json"
	"slotman/drivers/impl/spi"
	"slotman/services/type/proxy"
	"slotman/utils/log"
)

func (sv *Service) handleSpi(reqBytes []byte) (resBytes []byte, err error) {

	sv.spiDevLock.Lock()
	defer sv.spiDevLock.Unlock()

	req := proxy.Spi{}

	err = json.Unmarshal(reqBytes, &req)
	if err != nil {
		log.Cerror(err)
		return
	}

	//
	// Check for calls w/o device.
	//

	//if req.What == proxy. {
	//	req.Ok, req.Err = gpio.HasGpio()
	//	resBytes, err = json.Marshal(req)
	//	return
	//}

	//
	// Check and create device.
	//

	spiDev := sv.spiDevMap[req.Device]
	if spiDev == nil {
		spiDev = spi.NewDevice(req.Device)
		sv.spiDevMap[req.Device] = spiDev
	}

	switch req.What {

	case proxy.SpiWhatOpen:
		req.Err = spiDev.Open()

	case proxy.SpiWhatClose:
		req.Err = spiDev.Close()

	case proxy.SpiWhatSetMode:
		req.Err = spiDev.SetMode(req.Mode)
		req.Mode = 0

	case proxy.SpiWhatSetBpw:
		req.Err = spiDev.SetBitsPerWord(req.Bpw)
		req.Bpw = 0

	case proxy.SpiWhatSetSpeed:
		req.Err = spiDev.SetSpeed(req.Speed)
		req.Speed = 0

	case proxy.SpiWhatSend:
		req.Recv, req.Err = spiDev.Send(req.Send)
		req.Send = nil
	}

	req.Ok = req.Err == nil

	resBytes, err = json.Marshal(req)

	return
}
