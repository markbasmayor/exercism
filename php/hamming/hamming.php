<?php

/*
This is only a SKELETON file for the "Hamming" exercise. It's been provided as a
convenience to get you started writing code faster.

Remove this comment before submitting your exercise.
*/

function distance(string $strandA, string $strandB) : int
{
    if (mb_strlen($strandA) != mb_strlen($strandB)) {
        throw new InvalidArgumentException('DNA strands must be of equal length.');
    }

    $a = str_split($strandA);
    $b = str_split($strandB);

    return count(array_diff_assoc($a,$b));
}
