package api

import (
	"github.com/ditsuke/go-amizone/amizone"
	"github.com/ditsuke/go-amizone/server/handlers"
	"github.com/go-logr/logr"
	"net/http"
)

var apiHandlers handlers.Cfg

func GetSchedule(rw http.ResponseWriter, req *http.Request) {
	apiHandlers.ClassScheduleHandler(rw, req)
}

// init initializes the api handlers
func init() {
	apiHandlers = handlers.Cfg{
		L: logr.Discard(),
		A: func(cred amizone.Credentials, httpClient *http.Client) (amizone.ClientInterface, error) {
			return amizone.NewClient(cred, httpClient)
		},
	}
}
