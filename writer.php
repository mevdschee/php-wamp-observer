<?php

$i = 0;
while (true) {
  $id = uniqid();
  $i++;
  $t = random_int(1, 9);
  $inMsg = '[2,"' . $id . '","hello' . $t . '",{"msg":"hello world' . $i . ' details"}]';
  $outMsg = '[3,"' . $id . '","{msg":"hello world' . $i . ' reply details"}]';
  $errorMsg = '[4,"' . $id . ',"{msg":"hello world' . $i . ' error details"}]';
  $protocol = 'wamp';
  $direction = 'in';
  // send request
  if (WampObserver::logging()) {
    WampObserver::log($protocol, $direction, $inMsg);
  }
  usleep(random_int(50, 100) * 1000);
  // 1 ot of 9 gets delayed for the timeout period
  if (random_int(1, 9) == 5) {
    sleep(3);
  }
  // send response
  if (WampObserver::logging()) {
    // 1 ot of 9 does not get answered
    if (random_int(1, 9) != 5) {
      // 1 ot of 99 is an error
      if (random_int(1, 99) == 5) {
        WampObserver::log($protocol, $direction, $errorMsg);
      } else {
        WampObserver::log($protocol, $direction, $outMsg);
      }
    }
  }
  usleep(25 * 1000);
}

class WampObserver
{
  private static ?Socket $socket = null;
  private static bool $connected = false;

  public static function logging(bool $connect = true): bool
  {
    if (!self::$socket) {
      self::$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP) ?: null;
      self::$connected = false;
    }
    if (!self::$connected) {
      if ($connect) {
        self::$connected = @socket_connect(self::$socket, 'localhost', '6666');
      }
    }
    return self::$connected;
  }

  public static function log(string $protocol, string $direction, string $message)
  {
    if (self::$connected) {
      $line = "$protocol:$direction:$message\n";
      if (!@socket_write(self::$socket, $line, strlen($line))) {
        self::$socket = null;
        self::$connected = false;
      }
    }
  }
}
