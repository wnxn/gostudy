package filelisting

import (
	"net/http"
	"io/ioutil"
	"os"
)

func HandlerFileList(writer http.ResponseWriter,
	request *http.Request)error{
	path := request.URL.Path[len("/list/"):]
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
