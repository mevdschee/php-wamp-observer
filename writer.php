<?php

$i = 0;
while (true) {
  $id = uniqid();
  $i++;
  $t = random_int(1, 9);
  $inMsg = '[2,"' . $id . '","hello' . $t . '",{"msg":"hello world' . $i . '"}]';
  $outMsg = '[3,"' . $id . '","hello' . $t . '","msg":"hello world' . $i . ' reply"}]';
  $protocol = 'wamp';
  $direction = 'in';
  if (WampObserver::logging()) {
    WampObserver::log($protocol, $direction, $inMsg);
  }
  usleep(random_int(500, 1000) * 1000);
  if (WampObserver::logging()) {
    WampObserver::log($protocol, $direction, $outMsg);
  }
  usleep(250 * 1000);
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
