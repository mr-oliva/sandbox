<?php

class P {
    private $name;
    public function __construct(string $name) {
        $this->name = $name;
        echo ($this->name."\n");
    }
}

class C extends P{
}
class D extends P{
}

$c = new C("c");
$d = new D("d");

