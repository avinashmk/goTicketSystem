package server

import (
	"net/http"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/internal/server/handler"
	"github.com/avinashmk/goTicketSystem/logger"
)

const (
	port = ":9908"
)

var (
	lastSessionID    int
	connectionClosed chan int
	serverClosed     chan bool
	//connections      map[int]*Session
	server *http.Server
	//stopMonitorConnections chan bool // re-usable channel to indicate both 'stop' & 'stop complete' signal
)

// Init inits
func Init() bool {
	logger.Info.Println("Init")
	server = &http.Server{Addr: port, Handler: nil}
	lastSessionID = 0
	//connections = make(map[int]*Session)
	//stopMonitorConnections = make(chan bool)
	connectionClosed = make(chan int)
	serverClosed = make(chan bool)
	return true
}

// Finalize Finalizes
func Finalize() {
	logger.Debug.Println("Server Closing...")
	server.Close()
	<-serverClosed
	logger.Debug.Println("Server Closed")
	//stopMonitorConnections <- true
	//<-stopMonitorConnections
	logger.Info.Println("Finalize")
}

// Run starts
func Run() bool {
	logger.Enter.Println("Run")
	defer logger.Leave.Println("Run")

	defer func() { serverClosed <- true }()
	// monitorConnections()
	setupHandlers()
	startServer()
	return true
}

func startServer() {
	logger.Enter.Println("startServer")
	defer logger.Leave.Println("startServer")

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logger.Info.Println(err)
		} else {
			logger.Err.Println(err)
		}
	}
}

func setupHandlers() {
	logger.Enter.Println("setupHandlers()")
	defer logger.Leave.Println("setupHandlers()")

	fileServer := http.FileServer(http.Dir("./web/static"))
	http.Handle("/", fileServer)
	http.HandleFunc(consts.SignInFunc, handler.Signin)
	http.HandleFunc(consts.SignUpFunc, handler.Signup)

	http.HandleFunc(consts.SearchTrainPostAction, handler.SearchTrain)
	http.HandleFunc(consts.MakeReservPostAction, handler.MakeReservation)
	http.HandleFunc(consts.CancelReservPostAction, handler.CancelReservation)
	http.HandleFunc(consts.ViewReservPostAction, handler.ViewReservation)

	http.HandleFunc(consts.AddTrainSchemaFormPostAction, handler.AddTrainSchemaForm)
	http.HandleFunc(consts.AddTrainSchemaPostAction, handler.AddTrainSchema)
	http.HandleFunc(consts.RemoveTrainSchemaPostAction, handler.RemoveTrainSchema)
	http.HandleFunc(consts.ViewTrainSchemaPostAction, handler.ViewTrainSchema)
	http.HandleFunc(consts.UpdateTrainSchemaPostAction, handler.UpdateTrainSchema)
}

// func handleNewConnection() (result bool) {
// 	logger.Enter.Println("handleNewConnection")
// 	defer logger.Leave.Println("handleNewConnection")

// 	sID := nextSessionID()
// 	s := NewSession(sID)
// 	connections[sID] = s
// 	go s.Start()

// 	// TEST CODE BEGIN
// 	// time.Sleep(2 * time.Second)
// 	// sID1 := nextSessionID()
// 	// s1 := NewSession(sID1)
// 	// connections[sID1] = s1
// 	// go s1.Start()
// 	// TEST CODE END

// 	result = true
// 	return
// }

// func monitorConnections() {
// 	logger.Enter.Println("monitorConnections")
// 	defer logger.Leave.Println("monitorConnections")

// 	go func() {
// 		logger.Info.Println("monitorConnections started")
// 		for {
// 			select {
// 			case sessionID := <-connectionClosed:
// 				delete(connections, sessionID)
// 				logger.Info.Println("Closed Session: ", sessionID)
// 			case <-stopMonitorConnections:
// 				logger.Info.Println("monitorConnections stopped")
// 				for k, v := range connections {
// 					logger.Info.Println("Closing session: ", k)
// 					v.Close()
// 				}
// 				stopMonitorConnections <- true
// 				return
// 			}
// 		}
// 	}()
// }

// func nextSessionID() int {
// 	lastSessionID++
// 	logger.Debug.Println("nextSessionID(): ", lastSessionID)
// 	return lastSessionID
// }
