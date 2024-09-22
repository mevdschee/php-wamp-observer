# php-wamp-observer

To run the server:

    go run .

In bash run:

    for run in {1..100}; do php writer.php & done

And to stop:

    killall php

Now observe the metrics:

http://localhost:4000/

NB: The metrics are Prometheus compatible and follow the [OpenMetrics specification](https://github.com/OpenObservability/OpenMetrics/).

Example metrics:

    # HELP wamp_in_errors_seconds A summary of the wamp in errors.
    # TYPE wamp_in_errors_seconds summary
    wamp_in_errors_seconds_count{message="hello1"} 9
    wamp_in_errors_seconds_sum{message="hello1"} 0.723
    wamp_in_errors_seconds_count{message="hello2"} 7
    wamp_in_errors_seconds_sum{message="hello2"} 0.502
    wamp_in_errors_seconds_count{message="hello3"} 6
    wamp_in_errors_seconds_sum{message="hello3"} 0.426
    wamp_in_errors_seconds_count{message="hello4"} 7
    wamp_in_errors_seconds_sum{message="hello4"} 0.566
    wamp_in_errors_seconds_count{message="hello5"} 13
    wamp_in_errors_seconds_sum{message="hello5"} 0.971
    wamp_in_errors_seconds_count{message="hello6"} 3
    wamp_in_errors_seconds_sum{message="hello6"} 0.257
    wamp_in_errors_seconds_count{message="hello7"} 4
    wamp_in_errors_seconds_sum{message="hello7"} 0.307
    wamp_in_errors_seconds_count{message="hello8"} 3
    wamp_in_errors_seconds_sum{message="hello8"} 0.271
    wamp_in_errors_seconds_count{message="hello9"} 5
    wamp_in_errors_seconds_sum{message="hello9"} 0.411
    # HELP wamp_in_errors_total_seconds A histogram of the wamp in errors.
    # TYPE wamp_in_errors_total_seconds histogram
    wamp_in_errors_total_seconds_bucket{le="0.001"} 0
    wamp_in_errors_total_seconds_bucket{le="0.01"} 0
    wamp_in_errors_total_seconds_bucket{le="0.1"} 56
    wamp_in_errors_total_seconds_bucket{le="1"} 57
    wamp_in_errors_total_seconds_bucket{le="10"} 57
    wamp_in_errors_total_seconds_bucket{le="100"} 57
    wamp_in_errors_total_seconds_bucket{le="+Inf"} 57
    wamp_in_errors_total_seconds_sum 4.435
    wamp_in_errors_total_seconds_count 57
    # HELP wamp_in_responses_seconds A summary of the wamp in responses.
    # TYPE wamp_in_responses_seconds summary
    wamp_in_responses_seconds_count{message="hello1"} 591
    wamp_in_responses_seconds_sum{message="hello1"} 43.989
    wamp_in_responses_seconds_count{message="hello2"} 576
    wamp_in_responses_seconds_sum{message="hello2"} 42.666
    wamp_in_responses_seconds_count{message="hello3"} 624
    wamp_in_responses_seconds_sum{message="hello3"} 46.508
    wamp_in_responses_seconds_count{message="hello4"} 568
    wamp_in_responses_seconds_sum{message="hello4"} 43.037
    wamp_in_responses_seconds_count{message="hello5"} 632
    wamp_in_responses_seconds_sum{message="hello5"} 48.004
    wamp_in_responses_seconds_count{message="hello6"} 660
    wamp_in_responses_seconds_sum{message="hello6"} 49.111
    wamp_in_responses_seconds_count{message="hello7"} 645
    wamp_in_responses_seconds_sum{message="hello7"} 48.806
    wamp_in_responses_seconds_count{message="hello8"} 632
    wamp_in_responses_seconds_sum{message="hello8"} 47.228
    wamp_in_responses_seconds_count{message="hello9"} 644
    wamp_in_responses_seconds_sum{message="hello9"} 48.047
    # HELP wamp_in_responses_total_seconds A histogram of the wamp in responses.
    # TYPE wamp_in_responses_total_seconds histogram
    wamp_in_responses_total_seconds_bucket{le="0.001"} 0
    wamp_in_responses_total_seconds_bucket{le="0.01"} 0
    wamp_in_responses_total_seconds_bucket{le="0.1"} 5454
    wamp_in_responses_total_seconds_bucket{le="1"} 5572
    wamp_in_responses_total_seconds_bucket{le="10"} 5572
    wamp_in_responses_total_seconds_bucket{le="100"} 5572
    wamp_in_responses_total_seconds_bucket{le="+Inf"} 5572
    wamp_in_responses_total_seconds_sum 417.396
    wamp_in_responses_total_seconds_count 5572
    # HELP wamp_in_timeouts_seconds A summary of the wamp in timeouts.
    # TYPE wamp_in_timeouts_seconds summary
    wamp_in_timeouts_seconds_count{message="hello1"} 113
    wamp_in_timeouts_seconds_sum{message="hello1"} 339.047
    wamp_in_timeouts_seconds_count{message="hello2"} 128
    wamp_in_timeouts_seconds_sum{message="hello2"} 384.053
    wamp_in_timeouts_seconds_count{message="hello3"} 126
    wamp_in_timeouts_seconds_sum{message="hello3"} 378.055
    wamp_in_timeouts_seconds_count{message="hello4"} 104
    wamp_in_timeouts_seconds_sum{message="hello4"} 312.045
    wamp_in_timeouts_seconds_count{message="hello5"} 108
    wamp_in_timeouts_seconds_sum{message="hello5"} 324.042
    wamp_in_timeouts_seconds_count{message="hello6"} 127
    wamp_in_timeouts_seconds_sum{message="hello6"} 381.049
    wamp_in_timeouts_seconds_count{message="hello7"} 107
    wamp_in_timeouts_seconds_sum{message="hello7"} 321.049
    wamp_in_timeouts_seconds_count{message="hello8"} 101
    wamp_in_timeouts_seconds_sum{message="hello8"} 303.042
    wamp_in_timeouts_seconds_count{message="hello9"} 120
    wamp_in_timeouts_seconds_sum{message="hello9"} 360.048
    # HELP wamp_in_timeouts_total_seconds A histogram of the wamp in timeouts.
    # TYPE wamp_in_timeouts_total_seconds histogram
    wamp_in_timeouts_total_seconds_bucket{le="0.001"} 0
    wamp_in_timeouts_total_seconds_bucket{le="0.01"} 0
    wamp_in_timeouts_total_seconds_bucket{le="0.1"} 0
    wamp_in_timeouts_total_seconds_bucket{le="1"} 0
    wamp_in_timeouts_total_seconds_bucket{le="10"} 1034
    wamp_in_timeouts_total_seconds_bucket{le="100"} 1034
    wamp_in_timeouts_total_seconds_bucket{le="+Inf"} 1034
    wamp_in_timeouts_total_seconds_sum 3102.429
    wamp_in_timeouts_total_seconds_count 1034

Enjoy!