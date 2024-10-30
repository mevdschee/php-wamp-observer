<?php

class WampObserver
{
    public static $address = 'localhost';
    public static $port = 6666;

    private static $socket = null;
    private static $connected = false;
    private static $connectAt = 0;

    public static function logging()
    {
        if (self::$socket === null) {
            self::$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
            if (!self::$socket) {
                self::$socket = null;
            }
            self::$connected = false;
        }
        if (!self::$connected) {
            $now = time();
            if (self::$connectAt != $now) {
                self::$connectAt = $now;
                self::$connected = @socket_connect(self::$socket, self::$address, self::$port);
            }
        }
        return self::$connected;
    }

    public static function log(string $protocol, string $direction, string $message)
    {
        if (self::logging()) {
            $line = json_encode(array($protocol, $direction, $message));
            if (!@socket_write(self::$socket, $line . "\n", strlen($line) + 1)) {
                self::$socket = null;
                self::$connected = false;
            }
        }
    }
}
