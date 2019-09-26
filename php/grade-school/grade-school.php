<?php

class School {

    private $studentsByGrade = [];

    public function add(string $student, int $grade): void {
        if (!array_key_exists($grade, $this->studentsByGrade)) {
            $this->studentsByGrade[$grade] = [];
        }
        $this->studentsByGrade[$grade][] = $student;
    }

    public function grade(int $grade): Array {
        return $this->studentsByGrade[$grade] ?? [];
    }

    public function studentsByGradeAlphabetical(): Array {
        ksort($this->studentsByGrade);
        foreach ($this->studentsByGrade as $grade => $students) {
            if (count($students) < 2) {
                continue;
            }
            sort($this->studentsByGrade[$grade]);
        }
        return $this->studentsByGrade;
    }

}

