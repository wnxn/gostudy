package filelisting

import (
	"net/http"
	"io/ioutil"
	"os"
	"fmt"
	"strings"
)

const prefix ="/list/"

type userError string

func (e userError)Error()string{
	return e.Message()
}

func(e userError)Message()string{
	return string(e)
}

func HandlerFileList(writer http.ResponseWriter,
	request *http.Request)error{
		fmt.Println()
		// URL prefix error
		if strings.Index(request.URL.Path,prefix) !=0{
			return userError(fmt.Sprintf("path %s must start with %s",
				request.URL.Path,
					prefix))
		}
		path := request.URL.Path[len(prefix):]
		file,err:= os.Open(path)
		if err != nil{
			// If we use error wrapper, not print error details
			//	panic("open URL err")
			//http.Error(
			//	writer,
			//	err.Error(),
			//	http.StatusInternalServerError)
			return err
		}
		defer file.Close()

		all, err:=ioutil.ReadAll(file)
		if err != nil{
			return err
		}
		writer.Write(all)
		return nil
}
