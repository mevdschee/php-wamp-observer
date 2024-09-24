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
  usleep(random_int(50, 100) * 100);
  // 1 out of 9 gets delayed for the timeout period
  if (random_int(1, 9) == 5) {
    usleep(300 * 1000);
  }
  // send response
  if (WampObserver::logging()) {
    // 1 out of 9 does not get answered
    if (random_int(1, 9) != 5) {
      // 1 out of 99 is an error
      if (random_int(1, 99) == 5) {
        WampObserver::log($protocol, $direction, $errorMsg);
      } else {
        WampObserver::log($protocol, $direction, $outMsg);
      }
    }
  }
  usleep(25 * 100);
}

class WampObserver
{
  public static string $address = 'localhost';
  public static int $port = 6666;

  private static ?Socket $socket = null;
  private static bool $connected = false;
  private static int $connectAt = 0;

  public static function logging(bool $connect = true): bool
  {
    if (!self::$socket) {
      self::$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP) ?: null;
      socket_set_nonblock(self::$socket);
      self::$connected = false;
    }
    if (!self::$connected) {
      $now = time();
      if ($connect && self::$connectAt != $now) {
        self::$connectAt = $now;
        self::$connected = @socket_connect(self::$socket, self::$address, self::$port);
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
