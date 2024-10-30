<?php

class WampObserver
{
    public static string $address = 'localhost';
    public static int $port = 6666;

    private static ?Socket $socket = null;
    private static bool $connected = false;
    private static int $connectAt = 0;

    public static function logging(): bool
    {
        if (self::$socket === null) {
            self::$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP) ?: null;
            socket_set_option(self::$socket, SOL_SOCKET, SO_SNDTIMEO, ['sec' => 0, 'usec' => 1]);
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
            $line = json_encode([$protocol, $direction, $message]);
            if (!@socket_write(self::$socket, $line . "\n", strlen($line) + 1)) {
                self::$socket = null;
                self::$connected = false;
            }
        }
    }
}
