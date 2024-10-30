package tracking

import (
	"compress/gzip"
	"io"
	"math"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mevdschee/php-observability/metrics"
)

// TestTrack tracks the time spent and count of request response pairs with
// the same messageId and checks that these are properly reported in the metrics.
func TestTrack(t *testing.T) {
	stats := metrics.New()
	track := New(stats, 10*time.Millisecond)
	protocol := "wamp"
	direction := "in"
	request := "[2, \"123\", \"Hello\", {\"message\": \"Hello world?\"}]"
	err := track.Track(protocol+"_"+direction, request)
	if err != nil {
		t.Errorf("error tracking request: %q", err.Error())
	}
	time.Sleep(1 * time.Millisecond)
	response := "[3, \"123\", {\"message\": \"Hello world!\"}]"
	err = track.Track(protocol+"_"+direction, response)
	if err != nil {
		t.Errorf("error tracking response: %q", err.Error())
	}
	w := httptest.NewRecorder()
	stats.Write(w)
	resp := w.Result()
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		t.Errorf("error reading gz: %q", err.Error())
	}
	body, err := io.ReadAll(gz)
	if err != nil {
		t.Errorf("error reading body: %q", err.Error())
	}
	got := string(body)
	want := "wamp_in_responses_seconds_count{message=\"Hello\"} 1\nwamp_in_responses_seconds_sum{message=\"Hello\"} 0.001"
	if !strings.Contains(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestTrackError tracks the time spent and count of request (error) response pairs
// with the same messageId and checks that these are properly reported in the metrics.
func TestTrackError(t *testing.T) {
	stats := metrics.New()
	track := New(stats, 10*time.Millisecond)
	protocol := "wamp"
	direction := "in"
	request := "[2, \"123\", \"Hello\", {\"message\": \"Hello world?\"}]"
	err := track.Track(protocol+"_"+direction, request)
	if err != nil {
		t.Errorf("error tracking request: %q", err.Error())
	}
	time.Sleep(1 * time.Millisecond)
	response := "[4, \"123\", {\"error\": \"Can't say hello.\"}]"
	err = track.Track(protocol+"_"+direction, response)
	if err != nil {
		t.Errorf("error tracking response: %q", err.Error())
	}
	w := httptest.NewRecorder()
	stats.Write(w)
	resp := w.Result()
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		t.Errorf("error reading gz: %q", err.Error())
	}
	body, err := io.ReadAll(gz)
	if err != nil {
		t.Errorf("error reading body: %q", err.Error())
	}
	got := string(body)
	want := "wamp_in_errors_seconds_count{message=\"Hello\"} 1\nwamp_in_errors_seconds_sum{message=\"Hello\"} 0.001"
	if !strings.Contains(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestTrackTimeout tests the response timeout after a request with the same
// messageId and checks that this timeout is properly reported in the metrics.
func TestTrackTimeout(t *testing.T) {
	stats := metrics.New()
	track := New(stats, 1*time.Millisecond)
	protocol := "wamp"
	direction := "in"
	request := "[2, \"123\", \"Hello\", {\"message\": \"Hello world!\"}]"
	err := track.Track(protocol+"_"+direction, request)
	if err != nil {
		t.Errorf("error tracking request: %q", err.Error())
	}
	time.Sleep(10 * time.Millisecond)
	response := "[3, \"123\", {\"message\": \"Hello world!\"}]"
	err = track.Track(protocol+"_"+direction, response)
	if err != nil {
		t.Errorf("error tracking response: %q", err.Error())
	}
	w := httptest.NewRecorder()
	stats.Write(w)
	resp := w.Result()
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		t.Errorf("error reading gz: %q", err.Error())
	}
	body, err := io.ReadAll(gz)
	if err != nil {
		t.Errorf("error reading body: %q", err.Error())
	}
	got := string(body)
	want := "wamp_in_timeouts_seconds_count{message=\"Hello\"} 1\nwamp_in_timeouts_seconds_sum{message=\"Hello\"} 0.001"
	if !strings.Contains(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestBuckets tracks the time spent and count of 2 request response pairs with
// the same messageId and checks that these are properly reported in the metrics
// in the correct buckets.
func TestBuckets(t *testing.T) {
	stats := metrics.New()
	track := New(stats, 100*time.Millisecond)
	protocol := "wamp"
	direction := "in"
	for i := 0; i < 2; i++ {
		request := "[2, \"123\", \"Hello\", {\"message\": \"Hello world?\"}]"
		err := track.Track(protocol+"_"+direction, request)
		if err != nil {
			t.Errorf("error tracking request: %q", err.Error())
		}
		time.Sleep(time.Duration(math.Pow10(i)) * time.Millisecond)
		response := "[3, \"123\", {\"message\": \"Hello world!\"}]"
		err = track.Track(protocol+"_"+direction, response)
		if err != nil {
			t.Errorf("error tracking response: %q", err.Error())
		}
	}
	w := httptest.NewRecorder()
	stats.Write(w)
	resp := w.Result()
	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		t.Errorf("error reading gz: %q", err.Error())
	}
	body, err := io.ReadAll(gz)
	if err != nil {
		t.Errorf("error reading body: %q", err.Error())
	}
	got := string(body)
	want1 := "wamp_in_responses_total_seconds_bucket{le=\"0.005\"} 1"
	want2 := "wamp_in_responses_total_seconds_bucket{le=\"0.025\"} 2"
	if !strings.Contains(got, want1) {
		t.Errorf("got %s, want %s", got, want1)
	}
	if !strings.Contains(got, want2) {
		t.Errorf("got %s, want %s", got, want2)
	}
}
