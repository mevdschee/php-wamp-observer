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

    wamp_in_response_count{message="ALL"} 5747
    wamp_in_response_count{message="hello1"} 641
    wamp_in_response_count{message="hello2"} 608
    wamp_in_response_count{message="hello3"} 631
    wamp_in_response_count{message="hello4"} 607
    wamp_in_response_count{message="hello5"} 651
    wamp_in_response_count{message="hello6"} 637
    wamp_in_response_count{message="hello7"} 643
    wamp_in_response_count{message="hello8"} 651
    wamp_in_response_count{message="hello9"} 678
    wamp_in_response_duration{message="ALL"} 432.056
    wamp_in_response_duration{message="hello1"} 48.496
    wamp_in_response_duration{message="hello2"} 45.984
    wamp_in_response_duration{message="hello3"} 47.966
    wamp_in_response_duration{message="hello4"} 44.981
    wamp_in_response_duration{message="hello5"} 48.651
    wamp_in_response_duration{message="hello6"} 47.589
    wamp_in_response_duration{message="hello7"} 47.854
    wamp_in_response_duration{message="hello8"} 49.659
    wamp_in_response_duration{message="hello9"} 50.877
    wamp_in_timeout_count{message="ALL"} 1361
    wamp_in_timeout_count{message="hello1"} 173
    wamp_in_timeout_count{message="hello2"} 148
    wamp_in_timeout_count{message="hello3"} 157
    wamp_in_timeout_count{message="hello4"} 163
    wamp_in_timeout_count{message="hello5"} 153
    wamp_in_timeout_count{message="hello6"} 125
    wamp_in_timeout_count{message="hello7"} 146
    wamp_in_timeout_count{message="hello8"} 149
    wamp_in_timeout_count{message="hello9"} 147
    wamp_in_timeout_duration{message="ALL"} 4083.620
    wamp_in_timeout_duration{message="hello1"} 519.076
    wamp_in_timeout_duration{message="hello2"} 444.069
    wamp_in_timeout_duration{message="hello3"} 471.073
    wamp_in_timeout_duration{message="hello4"} 489.073
    wamp_in_timeout_duration{message="hello5"} 459.073
    wamp_in_timeout_duration{message="hello6"} 375.056
    wamp_in_timeout_duration{message="hello7"} 438.066
    wamp_in_timeout_duration{message="hello8"} 447.070
    wamp_in_timeout_duration{message="hello9"} 441.062

Enjoy!