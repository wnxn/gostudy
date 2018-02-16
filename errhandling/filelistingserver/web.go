package main

import (
	"github.com/wnxn/gostudy/errhandling/filelistingserver/filelisting"
	"net/http"
	"os"
	//"github.com/gpmgo/gopm/modules/log"
	"log"
	_ "net/http/pprof"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(
	http.ResponseWriter,
	*http.Request) {
	return func(writer http.ResponseWriter,
		request *http.Request) {
		// panic error
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)

		if err != nil {
			log.Printf("Error occured %s\n", err.Error())
			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			// server error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,
				http.StatusText(code),
				code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandlerFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
