package main

import (
	"bufio"
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
	go serve()
	wampListener()
}

func serve() {
	http.ListenAndServe(":4000", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		stats.Write(&writer)
	}))
}

func wampListener() {
	lis, err := net.Listen("tcp", ":6666")
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
		msgName := strings.Trim(msgFields[2], "\"")
		if strings.TrimSpace(msgType) == "2" {
			track.Add(msgId, time.Now(), 3*time.Second, func() {
				start, ok := track.Del(msgId)
				if ok {
					duration := time.Since(start).Seconds()
					stats.Add(protocol+"_"+direction+"_timeout", msgName, duration)
					stats.Add(protocol+"_"+direction+"_timeout", "ALL", duration)
				}
			})
		}
		if strings.TrimSpace(msgType) == "3" {
			start, ok := track.Del(msgId)
			if ok {
				duration := time.Since(start).Seconds()
				stats.Add(protocol+"_"+direction+"_response", msgName, duration)
				stats.Add(protocol+"_"+direction+"_response", "ALL", duration)
			}
		}
		if strings.TrimSpace(msgType) == "4" {
			start, ok := track.Del(msgId)
			if ok {
				duration := time.Since(start).Seconds()
				stats.Add(protocol+"_"+direction+"_error", msgName, duration)
				stats.Add(protocol+"_"+direction+"_error", "ALL", duration)
			}
		}
		//log.Printf("track length: %v", track.Len())
		log.Printf("received input: %v", input)
	}
}
