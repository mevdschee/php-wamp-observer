# php-wamp-observer

A code base to showcase high frequency websocket (WAMP RPC) message logging in PHP and aggregating into metrics using Go.

### Usage

To run the server:

    go run .

In bash run:

    for run in {1..100}; do php writer.php & done

And to stop:

    killall php

Now observe the metrics:

http://localhost:8080/

NB: The metrics are Prometheus compatible and follow the [OpenMetrics specification](https://github.com/OpenObservability/OpenMetrics/).

### Example metrics

Here is an example of published metrics:

    # HELP wamp_in_errors_seconds A summary of the wamp in errors.
    # TYPE wamp_in_errors_seconds summary
    wamp_in_errors_seconds_count{message="hello1"} 82
    wamp_in_errors_seconds_sum{message="hello1"} 0.752
    wamp_in_errors_seconds_count{message="hello2"} 97
    wamp_in_errors_seconds_sum{message="hello2"} 1.300
    wamp_in_errors_seconds_count{message="hello3"} 106
    wamp_in_errors_seconds_sum{message="hello3"} 1.307
    wamp_in_errors_seconds_count{message="hello4"} 91
    wamp_in_errors_seconds_sum{message="hello4"} 1.175
    wamp_in_errors_seconds_count{message="hello5"} 106
    wamp_in_errors_seconds_sum{message="hello5"} 1.636
    wamp_in_errors_seconds_count{message="hello6"} 95
    wamp_in_errors_seconds_sum{message="hello6"} 0.875
    wamp_in_errors_seconds_count{message="hello7"} 76
    wamp_in_errors_seconds_sum{message="hello7"} 0.652
    wamp_in_errors_seconds_count{message="hello8"} 95
    wamp_in_errors_seconds_sum{message="hello8"} 1.580
    wamp_in_errors_seconds_count{message="hello9"} 90
    wamp_in_errors_seconds_sum{message="hello9"} 0.855
    # HELP wamp_in_errors_total_seconds A histogram of the wamp in errors.
    # TYPE wamp_in_errors_total_seconds histogram
    wamp_in_errors_total_seconds_bucket{le="0.005"} 227
    wamp_in_errors_total_seconds_bucket{le="0.01"} 741
    wamp_in_errors_total_seconds_bucket{le="0.025"} 766
    wamp_in_errors_total_seconds_bucket{le="0.05"} 787
    wamp_in_errors_total_seconds_bucket{le="0.1"} 820
    wamp_in_errors_total_seconds_bucket{le="0.25"} 837
    wamp_in_errors_total_seconds_bucket{le="0.5"} 838
    wamp_in_errors_total_seconds_bucket{le="1"} 838
    wamp_in_errors_total_seconds_bucket{le="2.5"} 838
    wamp_in_errors_total_seconds_bucket{le="5"} 838
    wamp_in_errors_total_seconds_bucket{le="10"} 838
    wamp_in_errors_total_seconds_bucket{le="+Inf"} 838
    wamp_in_errors_total_seconds_sum 10.133
    wamp_in_errors_total_seconds_count 838
    # HELP wamp_in_responses_seconds A summary of the wamp in responses.
    # TYPE wamp_in_responses_seconds summary
    wamp_in_responses_seconds_count{message="hello1"} 9096
    wamp_in_responses_seconds_sum{message="hello1"} 140.204
    wamp_in_responses_seconds_count{message="hello2"} 8967
    wamp_in_responses_seconds_sum{message="hello2"} 140.017
    wamp_in_responses_seconds_count{message="hello3"} 9042
    wamp_in_responses_seconds_sum{message="hello3"} 138.221
    wamp_in_responses_seconds_count{message="hello4"} 9068
    wamp_in_responses_seconds_sum{message="hello4"} 138.337
    wamp_in_responses_seconds_count{message="hello5"} 9146
    wamp_in_responses_seconds_sum{message="hello5"} 140.650
    wamp_in_responses_seconds_count{message="hello6"} 8983
    wamp_in_responses_seconds_sum{message="hello6"} 139.571
    wamp_in_responses_seconds_count{message="hello7"} 9026
    wamp_in_responses_seconds_sum{message="hello7"} 142.527
    wamp_in_responses_seconds_count{message="hello8"} 9026
    wamp_in_responses_seconds_sum{message="hello8"} 138.335
    wamp_in_responses_seconds_count{message="hello9"} 8907
    wamp_in_responses_seconds_sum{message="hello9"} 130.873
    # HELP wamp_in_responses_total_seconds A histogram of the wamp in responses.
    # TYPE wamp_in_responses_total_seconds histogram
    wamp_in_responses_total_seconds_bucket{le="0.005"} 21483
    wamp_in_responses_total_seconds_bucket{le="0.01"} 69704
    wamp_in_responses_total_seconds_bucket{le="0.025"} 72428
    wamp_in_responses_total_seconds_bucket{le="0.05"} 74558
    wamp_in_responses_total_seconds_bucket{le="0.1"} 78781
    wamp_in_responses_total_seconds_bucket{le="0.25"} 80598
    wamp_in_responses_total_seconds_bucket{le="0.5"} 81261
    wamp_in_responses_total_seconds_bucket{le="1"} 81261
    wamp_in_responses_total_seconds_bucket{le="2.5"} 81261
    wamp_in_responses_total_seconds_bucket{le="5"} 81261
    wamp_in_responses_total_seconds_bucket{le="10"} 81261
    wamp_in_responses_total_seconds_bucket{le="+Inf"} 81261
    wamp_in_responses_total_seconds_sum 1248.733
    wamp_in_responses_total_seconds_count 81261
    # HELP wamp_in_timeouts_seconds A summary of the wamp in timeouts.
    # TYPE wamp_in_timeouts_seconds summary
    wamp_in_timeouts_seconds_count{message="hello1"} 1760
    wamp_in_timeouts_seconds_sum{message="hello1"} 527.778
    wamp_in_timeouts_seconds_count{message="hello2"} 1750
    wamp_in_timeouts_seconds_sum{message="hello2"} 525.126
    wamp_in_timeouts_seconds_count{message="hello3"} 1854
    wamp_in_timeouts_seconds_sum{message="hello3"} 556.331
    wamp_in_timeouts_seconds_count{message="hello4"} 1723
    wamp_in_timeouts_seconds_sum{message="hello4"} 516.831
    wamp_in_timeouts_seconds_count{message="hello5"} 1674
    wamp_in_timeouts_seconds_sum{message="hello5"} 502.152
    wamp_in_timeouts_seconds_count{message="hello6"} 1762
    wamp_in_timeouts_seconds_sum{message="hello6"} 528.338
    wamp_in_timeouts_seconds_count{message="hello7"} 1721
    wamp_in_timeouts_seconds_sum{message="hello7"} 516.419
    wamp_in_timeouts_seconds_count{message="hello8"} 1724
    wamp_in_timeouts_seconds_sum{message="hello8"} 517.313
    wamp_in_timeouts_seconds_count{message="hello9"} 1719
    wamp_in_timeouts_seconds_sum{message="hello9"} 515.787
    # HELP wamp_in_timeouts_total_seconds A histogram of the wamp in timeouts.
    # TYPE wamp_in_timeouts_total_seconds histogram
    wamp_in_timeouts_total_seconds_bucket{le="0.005"} 0
    wamp_in_timeouts_total_seconds_bucket{le="0.01"} 0
    wamp_in_timeouts_total_seconds_bucket{le="0.025"} 0
    wamp_in_timeouts_total_seconds_bucket{le="0.05"} 1
    wamp_in_timeouts_total_seconds_bucket{le="0.1"} 3
    wamp_in_timeouts_total_seconds_bucket{le="0.25"} 7
    wamp_in_timeouts_total_seconds_bucket{le="0.5"} 15687
    wamp_in_timeouts_total_seconds_bucket{le="1"} 15687
    wamp_in_timeouts_total_seconds_bucket{le="2.5"} 15687
    wamp_in_timeouts_total_seconds_bucket{le="5"} 15687
    wamp_in_timeouts_total_seconds_bucket{le="10"} 15687
    wamp_in_timeouts_total_seconds_bucket{le="+Inf"} 15687
    wamp_in_timeouts_total_seconds_sum 4706.074
    wamp_in_timeouts_total_seconds_count 15687


Enjoy!