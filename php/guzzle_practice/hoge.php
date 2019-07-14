<?php

//require_once 'vendor/autoload.php';

use GuzzleHttp\Client;

class Hoge
{
    private $client;

    function __construct() {
        $this->client = new Client();
    }

    function hoge() {
    }

}

$hoge = new Hoge();
