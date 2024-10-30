<?php

class WampObserver
{
    public static int $port = 6666;

    private static ?Socket $socket = null;
    private static bool $connected = false;
    private static int $connectAt = 0;

    public static function logging(): bool
    {
        if (self::$socket === null) {
            self::$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP) ?: null;
            self::$connected = false;
        }
        if (!self::$connected) {
            $now = time();
            if (self::$connectAt != $now) {
                self::$connectAt = $now;
                self::$connected = @socket_connect(self::$socket, 'localhost', self::$port);
            }
        }
        return self::$connected;
    }

    public static function log(string $protocol, string $direction, string $message)
    {
        if (self::logging()) {
            $line = json_encode([$protocol, $direction, $message]);
            if (!@socket_write(self::$socket, $line . "\n", strlen($line) + 1)) {
                self::$socket = null;
                self::$connected = false;
            }
        }
    }
}
