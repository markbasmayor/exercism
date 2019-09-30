<?php

function from(DateTimeInterface $date): DateTimeInterface {
    $gigasecondInterval = 'PT1000000000S';
    return $date->add(new DateInterval($gigasecondInterval));
}

