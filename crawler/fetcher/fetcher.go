package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func determineEncoding(r *bufio.Reader) encoding.Encoding{
	bytes, err:=r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error %v", err)
		return unicode.UTF8
	}
	e,_,_:=charset.DetermineEncoding(bytes, "")
	return e
}

func Fetch(url string) ([]byte, error){
	log.Printf("Fetching: %s", url)
	resp, err:=http.Get(url)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	// any encode convert to UTF-8
	bodyReader :=bufio.NewReader(resp.Body)
	e:=determineEncoding(bodyReader)
	utf8Reader:=transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}