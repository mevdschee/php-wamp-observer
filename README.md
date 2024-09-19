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

wamp_in_error_count{tag="ALL"} 88
wamp_in_error_count{tag="hello1"} 10
wamp_in_error_count{tag="hello2"} 9
wamp_in_error_count{tag="hello3"} 11
wamp_in_error_count{tag="hello4"} 9
wamp_in_error_count{tag="hello5"} 14
wamp_in_error_count{tag="hello6"} 8
wamp_in_error_count{tag="hello7"} 6
wamp_in_error_count{tag="hello8"} 9
wamp_in_error_count{tag="hello9"} 12
wamp_in_error_duration{tag="ALL"} 6.446
wamp_in_error_duration{tag="hello1"} 0.686
wamp_in_error_duration{tag="hello2"} 0.646
wamp_in_error_duration{tag="hello3"} 0.934
wamp_in_error_duration{tag="hello4"} 0.675
wamp_in_error_duration{tag="hello5"} 0.976
wamp_in_error_duration{tag="hello6"} 0.565
wamp_in_error_duration{tag="hello7"} 0.448
wamp_in_error_duration{tag="hello8"} 0.618
wamp_in_error_duration{tag="hello9"} 0.897
wamp_in_response_count{tag="ALL"} 8415
wamp_in_response_count{tag="hello1"} 928
wamp_in_response_count{tag="hello2"} 944
wamp_in_response_count{tag="hello3"} 918
wamp_in_response_count{tag="hello4"} 901
wamp_in_response_count{tag="hello5"} 942
wamp_in_response_count{tag="hello6"} 928
wamp_in_response_count{tag="hello7"} 930
wamp_in_response_count{tag="hello8"} 962
wamp_in_response_count{tag="hello9"} 962
wamp_in_response_duration{tag="ALL"} 634.884
wamp_in_response_duration{tag="hello1"} 69.936
wamp_in_response_duration{tag="hello2"} 70.757
wamp_in_response_duration{tag="hello3"} 69.322
wamp_in_response_duration{tag="hello4"} 67.647
wamp_in_response_duration{tag="hello5"} 70.822
wamp_in_response_duration{tag="hello6"} 70.273
wamp_in_response_duration{tag="hello7"} 70.030
wamp_in_response_duration{tag="hello8"} 73.255
wamp_in_response_duration{tag="hello9"} 72.843
wamp_in_timeout_count{tag="ALL"} 2128
wamp_in_timeout_count{tag="hello1"} 225
wamp_in_timeout_count{tag="hello2"} 241
wamp_in_timeout_count{tag="hello3"} 232
wamp_in_timeout_count{tag="hello4"} 256
wamp_in_timeout_count{tag="hello5"} 228
wamp_in_timeout_count{tag="hello6"} 216
wamp_in_timeout_count{tag="hello7"} 266
wamp_in_timeout_count{tag="hello8"} 241
wamp_in_timeout_count{tag="hello9"} 223
wamp_in_timeout_duration{tag="ALL"} 6385.460
wamp_in_timeout_duration{tag="hello1"} 675.146
wamp_in_timeout_duration{tag="hello2"} 723.150
wamp_in_timeout_duration{tag="hello3"} 696.166
wamp_in_timeout_duration{tag="hello4"} 768.205
wamp_in_timeout_duration{tag="hello5"} 684.158
wamp_in_timeout_duration{tag="hello6"} 648.146
wamp_in_timeout_duration{tag="hello7"} 798.177
wamp_in_timeout_duration{tag="hello8"} 723.160
wamp_in_timeout_duration{tag="hello9"} 669.152

Enjoy!