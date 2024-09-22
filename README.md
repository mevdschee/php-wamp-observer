# php-wamp-observer

To run the server:

    go run .

In bash run:

    for run in {1..100}; do php writer.php & done

And to stop:

    killall php

Now observe the stats:

http://localhost:4000/

Example stats:

    # HELP wamp_in_errors_seconds A summary of the wamp in errors.
    # TYPE wamp_in_errors_seconds summary
    wamp_in_errors_count{message="hello1"} 2
    wamp_in_errors_sum{message="hello1"} 0.137
    wamp_in_errors_count{message="hello2"} 1
    wamp_in_errors_sum{message="hello2"} 0.073
    wamp_in_errors_count{message="hello3"} 4
    wamp_in_errors_sum{message="hello3"} 0.308
    wamp_in_errors_count{message="hello4"} 4
    wamp_in_errors_sum{message="hello4"} 0.348
    wamp_in_errors_count{message="hello5"} 3
    wamp_in_errors_sum{message="hello5"} 0.231
    wamp_in_errors_count{message="hello6"} 2
    wamp_in_errors_sum{message="hello6"} 0.125
    wamp_in_errors_count{message="hello7"} 1
    wamp_in_errors_sum{message="hello7"} 0.083
    wamp_in_errors_count{message="hello8"} 5
    wamp_in_errors_sum{message="hello8"} 0.376
    wamp_in_errors_count{message="hello9"} 3
    wamp_in_errors_sum{message="hello9"} 0.188
    # HELP wamp_in_errors_seconds A histogram of the wamp in errors.
    # TYPE wamp_in_errors_seconds histogram
    wamp_in_errors_seconds_bucket{le="0.001"} 0
    wamp_in_errors_seconds_bucket{le="0.01"} 0
    wamp_in_errors_seconds_bucket{le="0.1"} 23
    wamp_in_errors_seconds_bucket{le="1"} 25
    wamp_in_errors_seconds_bucket{le="10"} 25
    wamp_in_errors_seconds_bucket{le="100"} 25
    wamp_in_errors_seconds_bucket{le="+Inf"} 25
    wamp_in_errors_seconds_sum 1.871
    wamp_in_errors_seconds_count 25
    # HELP wamp_in_responses_seconds A summary of the wamp in responses.
    # TYPE wamp_in_responses_seconds summary
    wamp_in_responses_count{message="hello1"} 381
    wamp_in_responses_sum{message="hello1"} 28.446
    wamp_in_responses_count{message="hello2"} 318
    wamp_in_responses_sum{message="hello2"} 23.833
    wamp_in_responses_count{message="hello3"} 353
    wamp_in_responses_sum{message="hello3"} 26.253
    wamp_in_responses_count{message="hello4"} 348
    wamp_in_responses_sum{message="hello4"} 25.849
    wamp_in_responses_count{message="hello5"} 299
    wamp_in_responses_sum{message="hello5"} 22.116
    wamp_in_responses_count{message="hello6"} 302
    wamp_in_responses_sum{message="hello6"} 22.547
    wamp_in_responses_count{message="hello7"} 334
    wamp_in_responses_sum{message="hello7"} 24.926
    wamp_in_responses_count{message="hello8"} 318
    wamp_in_responses_sum{message="hello8"} 24.257
    wamp_in_responses_count{message="hello9"} 317
    wamp_in_responses_sum{message="hello9"} 23.906
    # HELP wamp_in_responses_seconds A histogram of the wamp in responses.
    # TYPE wamp_in_responses_seconds histogram
    wamp_in_responses_seconds_bucket{le="0.001"} 0
    wamp_in_responses_seconds_bucket{le="0.01"} 0
    wamp_in_responses_seconds_bucket{le="0.1"} 2912
    wamp_in_responses_seconds_bucket{le="1"} 2970
    wamp_in_responses_seconds_bucket{le="10"} 2970
    wamp_in_responses_seconds_bucket{le="100"} 2970
    wamp_in_responses_seconds_bucket{le="+Inf"} 2970
    wamp_in_responses_seconds_sum 222.133
    wamp_in_responses_seconds_count 2970
    # HELP wamp_in_timeouts_seconds A summary of the wamp in timeouts.
    # TYPE wamp_in_timeouts_seconds summary
    wamp_in_timeouts_count{message="hello1"} 68
    wamp_in_timeouts_sum{message="hello1"} 204.035
    wamp_in_timeouts_count{message="hello2"} 79
    wamp_in_timeouts_sum{message="hello2"} 237.039
    wamp_in_timeouts_count{message="hello3"} 72
    wamp_in_timeouts_sum{message="hello3"} 216.036
    wamp_in_timeouts_count{message="hello4"} 74
    wamp_in_timeouts_sum{message="hello4"} 222.032
    wamp_in_timeouts_count{message="hello5"} 77
    wamp_in_timeouts_sum{message="hello5"} 231.037
    wamp_in_timeouts_count{message="hello6"} 78
    wamp_in_timeouts_sum{message="hello6"} 234.037
    wamp_in_timeouts_count{message="hello7"} 81
    wamp_in_timeouts_sum{message="hello7"} 243.043
    wamp_in_timeouts_count{message="hello8"} 82
    wamp_in_timeouts_sum{message="hello8"} 246.041
    wamp_in_timeouts_count{message="hello9"} 95
    wamp_in_timeouts_sum{message="hello9"} 285.045
    # HELP wamp_in_timeouts_seconds A histogram of the wamp in timeouts.
    # TYPE wamp_in_timeouts_seconds histogram
    wamp_in_timeouts_seconds_bucket{le="0.001"} 0
    wamp_in_timeouts_seconds_bucket{le="0.01"} 0
    wamp_in_timeouts_seconds_bucket{le="0.1"} 0
    wamp_in_timeouts_seconds_bucket{le="1"} 0
    wamp_in_timeouts_seconds_bucket{le="10"} 706
    wamp_in_timeouts_seconds_bucket{le="100"} 706
    wamp_in_timeouts_seconds_bucket{le="+Inf"} 706
    wamp_in_timeouts_seconds_sum 2118.343
    wamp_in_timeouts_seconds_count 706

Enjoy!