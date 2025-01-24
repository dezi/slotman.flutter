package proxy

import (
	"github.com/gorilla/websocket"
	"net/http"
	"slotman/services/iface/proxy"
	"slotman/services/impl/provider"
	"slotman/utils/log"
	"sync"
	"time"
)

type Service struct {
	httpMux     *http.ServeMux
	httpServer  *http.Server
	httpRunning bool

	webServer *websocket.Conn

	webClients map[string]*websocket.Conn
	mapsLock   sync.Mutex
}

var (
	singleTon *Service
)

func StartService() (err error) {

	if singleTon != nil {
		return
	}

	singleTon = &Service{}

	singleTon.webClients = make(map[string]*websocket.Conn)

	provider.SetProvider(singleTon)

	return
}

func StopService() (err error) {

	if singleTon == nil {
		return
	}

	provider.UnsetProvider(singleTon)

	log.Printf("Stopping service...")

	_ = singleTon.stopServers()

	log.Printf("Stopped service.")

	singleTon = nil

	return
}

func (sv *Service) GetName() (name provider.Service) {
	return proxy.Service
}

func (sv *Service) GetControlOptions() (interval time.Duration) {
	interval = time.Second * 60
	return
}
