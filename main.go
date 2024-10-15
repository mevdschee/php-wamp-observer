package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"runtime/pprof"
	"time"

	"github.com/mevdschee/php-observability/metrics"
	"github.com/mevdschee/php-wamp-observer/tracking"
)

var stats = metrics.New()
var track = tracking.New()

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	memprofile := flag.String("memprofile", "", "write mem profile to file")
	listenAddress := flag.String("listen", "localhost:6666", "address to listen for high frequent events over TCP")
	metricsAddress := flag.String("metrics", ":8080", "address to listen for Prometheus metric scraper over HTTP")
	binaryAddress := flag.String("binary", ":9999", "address to listen for Gob metric scraper over HTTP")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	go serve(*memprofile, *metricsAddress)
	go serveGob(*binaryAddress)
	wampListener(*listenAddress)
}

func serve(memprofile, metricsAddress string) {
	err := http.ListenAndServe(metricsAddress, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		stats.Write(&writer)
		if memprofile != "" {
			f, err := os.Create(memprofile)
			if err != nil {
				log.Fatal(err)
			}
			pprof.WriteHeapProfile(f)
			f.Close()
		}
	}))
	log.Fatal(err)
}

func serveGob(metricsAddress string) {
	err := http.ListenAndServe(metricsAddress, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		stats.WriteGob(&writer)
	}))
	log.Fatal(err)
}

func wampListener(listenAddress string) {
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to start listener: %v", err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Printf("failed to accept conn: %v", err)
			continue
		}
		go handleWampConn(conn)
	}
}

func handleWampConn(conn net.Conn) {
	defer conn.Close()
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		input := scan.Text()
		var fields []string
		json.Unmarshal([]byte(input), &fields)
		if len(fields) != 3 {
			log.Printf("malformed input: %v", input)
			continue
		}
		protocol := fields[0]
		direction := fields[1]
		messageString := fields[2]
		var message []any
		err := json.Unmarshal([]byte(messageString), &message)
		if err != nil {
			log.Printf("malformed message: %v", messageString)
			continue
		}
		msgType := int(message[0].(float64))
		msgId := message[1].(string)
		if msgType == 2 {
			msgName := message[2].(string)
			track.Add(msgId, msgName, time.Now(), 300*time.Millisecond, func() {
				start, msgName, ok := track.Del(msgId)
				if ok {
					duration := time.Since(start).Seconds()
					stats.Add(protocol+"_"+direction+"_timeouts", "message", msgName, duration)
				}
			})
		}
		if msgType == 3 {
			start, msgName, ok := track.Del(msgId)
			if ok {
				duration := time.Since(start).Seconds()
				stats.Add(protocol+"_"+direction+"_responses", "message", msgName, duration)
			}
		}
		if msgType == 4 {
			start, msgName, ok := track.Del(msgId)
			if ok {
				duration := time.Since(start).Seconds()
				stats.Add(protocol+"_"+direction+"_errors", "message", msgName, duration)
			}
		}
		//log.Printf("track length: %v", track.Len())
		log.Printf("received input: %v", input)
	}
}
