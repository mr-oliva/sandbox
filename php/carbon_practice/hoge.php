<?php
require_once 'vendor/autoload.php';

use Carbon\Carbon;

$now = new DateTime('now');
$dt = Carbon::instance($now);
$dt2 = Carbon::instance($now)->addDay(3);

echo($dt->toDateString());
echo("\n".$dt2->toDateString());


