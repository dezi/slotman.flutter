package tacho

import (
	"slotman/things/mcp/mcp23017"
	"slotman/utils/log"
)

func (sv *Service) DoControlTask() {
	sv.checkSensors()
}

func (sv *Service) checkSensors() {

	if sv.speedSensor == nil {

		sensors, err := mcp23017.ProbeThings(nil, []byte{0x21})

		if err != nil {
			log.Cerror(err)
		} else {

			for _, sensor := range sensors {

				sensor.SetHandler(sv)

				err = sensor.Open()
				if err != nil {
					log.Cerror(err)
					continue
				}

				err = sensor.Start()
				if err != nil {
					log.Cerror(err)
					_ = sensor.Close()
					continue
				}

				err = sensor.SetPinDirections(0x000f)
				if err != nil {
					log.Cerror(err)
					_ = sensor.Close()
					return
				}

				log.Printf("Registered sensor MCP23017 path=%s uuid=%s",
					sensor.DevicePath, sensor.GetUuid()[:8])

				sv.speedSensor = sensor

				sv.waitGroup.Add(2)

				go sv.speedRead()
				go sv.speedEval()
			}
		}
	}
}
