<?php

function sample(string $dt) {
    echo $dt;    
}

$dt = new DateTime();
sample($dt->format("Y-m-d H:i:s"));
