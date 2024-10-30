<?php

include 'WampObserver.php';

$i = 0;
while (true) {
  $id = uniqid();
  $i++;
  $t = random_int(1, 9);
  $inMsg = json_encode([2, $id, "hello$t", ["msg" => "hello world$i request details"]]);
  $outMsg = json_encode([3, $id, ["msg" => "hello world$i reply details"]]);
  $errorMsg = json_encode([4, $id, ["msg" => "hello world$i error details"]]);
  $protocol = 'wamp';
  $direction = 'in';
  // send request
  WampObserver::log($protocol, $direction, $inMsg);
  usleep(random_int(50, 100) * 100);
  // 1 out of 9 gets delayed for the timeout period
  if (random_int(1, 9) == 5) {
    usleep(300 * 1000);
  }
  // send response
  // 1 out of 9 does not get answered
  if (random_int(1, 9) != 5) {
    // 1 out of 99 is an error
    if (random_int(1, 99) == 5) {
      WampObserver::log($protocol, $direction, $errorMsg);
    } else {
      WampObserver::log($protocol, $direction, $outMsg);
    }
  }
  usleep(25 * 100);
}
