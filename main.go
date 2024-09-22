package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/mevdschee/php-wamp-observer/statistics"
	"github.com/mevdschee/php-wamp-observer/tracking"
)

var stats = statistics.New()
var track = tracking.New()

func main() {
	listenAddress := flag.String("listen", ":6666", "address to listen for high frequent events over TCP")
	metricsAddress := flag.String("metrics", ":4000", "address to listen for Prometheus metric scraper over HTTP")
	flag.Parse()
	go serve(*metricsAddress)
	wampListener(*listenAddress)
}

func serve(metricsAddress string) {
	http.ListenAndServe(metricsAddress, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		stats.Write(&writer)
	}))
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
		fields := strings.SplitN(input, ":", 3)
		if len(fields) != 3 {
			continue
		}
		protocol := fields[0]
		direction := fields[1]
		message := fields[2]
		msgFields := strings.SplitN(strings.Trim(message, "[]"), ",", 4)
		msgType := msgFields[0]
		msgId := strings.Trim(msgFields[1], "\"")
		if strings.TrimSpace(msgType) == "2" {
			msgName := strings.Trim(msgFields[2], "\"")
			track.Add(msgId, msgName, time.Now(), 3*time.Second, func() {
				start, msgName, ok := track.Del(msgId)
				if ok {
					duration := time.Since(start).Seconds()
					stats.Add(protocol+"_"+direction+"_timeouts", "message", msgName, duration)
				}
			})
		}
		if strings.TrimSpace(msgType) == "3" {
			start, msgName, ok := track.Del(msgId)
			if ok {
				duration := time.Since(start).Seconds()
				stats.Add(protocol+"_"+direction+"_responses", "message", msgName, duration)
			}
		}
		if strings.TrimSpace(msgType) == "4" {
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
