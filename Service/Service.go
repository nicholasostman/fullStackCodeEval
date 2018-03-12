package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"fullStackCodeEval/Joke"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sort"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Result I would normally just take len(result) and not waste memory on total, but it's a good demonstration of a more complex data struct
type Result []struct {
	Category []string
	Value    string
}
type jokeStruct struct {
	Total int
	Result
}

func main() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals)

	go func() {
		msg := <-signals
		log.Printf("Received Signal: %s", msg)
		// Would call a method to handle appCleanup as needed. This may include letting Docker know
		os.Exit(1)
	}()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello/{name}", helloEndPoint)
	router.HandleFunc("/category/", categoryEndPoint)
	router.HandleFunc("/joketerm/{joketerm}", joketermEndPoint)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Please use an appropriate endpoint \n /hello/ \n /category/ \n /joketerm/"))
	})

	handler := cors.Default().Handler(router)

	IP := getOutboundIP()
	log.Printf("creating listener on %s:%d", IP, 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:8080", IP), handler))
}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal("Encountered an error within net.Dial call: ", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func categoryEndPoint(w http.ResponseWriter, r *http.Request) {
}
func joketermEndPoint(w http.ResponseWriter, r *http.Request) {
	log.Println("Responding to /joketerm request")

	//handle comms to client
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	log.Println("request:", vars)

	rvars := r.URL.Query()
	log.Println("query string", rvars)
	joketerm := vars["joketerm"]
	if joketerm == "" {
		joketerm = "kite"
	}

	record := joke.Send(joketerm)

	// Take only pieces of record that we care about into joke to send: sendJoke
	var sendJoke jokeStruct
	sendJoke.Total = record.Total
	sendJoke.Result = make(Result, len(record.Result))

	for i := 0; i < len(record.Result); i++ {
		for j := 0; j < len(sendJoke.Result[i].Category); j++ {
			sendJoke.Result[i].Category = make([]string, len(record.Result[i].Category[j]))
			sendJoke.Result[i].Category[j] = record.Result[i].Category[j]
		}
		sendJoke.Result[i].Value = record.Result[i].Value
	}
	// garbage collection happens automatically with Go

	// sort server side for easier pagination
	// Sort alpha
	sort.Slice(sendJoke.Result[:], func(i, j int) bool {
		return sendJoke.Result[i].Value < sendJoke.Result[j].Value
	})

	// Send response
	body := "joke success"
	okResp := http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Request:       r,
		Header:        make(http.Header, 0),
	}

	// Convert to string
	jsonStr, err := json.Marshal(sendJoke)
	if err != nil {
		fmt.Println("error in converting joke struct to JSON: ", err)
	}

	buffer := bytes.NewBuffer(nil)
	okResp.Write(buffer)
	w.Write(jsonStr)
}

func helloEndPoint(w http.ResponseWriter, r *http.Request) {
	log.Println("Responding to /hello request")
	// log.Println(r.UserAgent())

	//handle comms to client
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	name := vars["name"]
	if name == "" {
		name = "Dave"
	}

	body := "Hello world"
	okResp := http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Request:       r,
		Header:        make(http.Header, 0),
	}
	buffer := bytes.NewBuffer(nil)
	okResp.Write(buffer)

	fmt.Fprintf(w, "Hello %s\n", name)
}
