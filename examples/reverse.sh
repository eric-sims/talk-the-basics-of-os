#!/bin/bash

A=0
B=0

while true; do
  clear
  for (( j=0; j<628; j+=7 )); do
    for (( i=0; i<628; i+=2 )); do
      c=$(echo "s($i/100)" | bc -l)
      d=$(echo "c($j/100)" | bc -l)
      e=$(echo "s($A/100)" | bc -l)
      f=$(echo "s($j/100)" | bc -l)
      g=$(echo "c($A/100)" | bc -l)
      h=$(echo "$d + 2" | bc -l)
      D=$(echo "1/($c * $h * $e + $f * $g + 5)" | bc -l)
      l=$(echo "c($i/100)" | bc -l)
      m=$(echo "c($B/100)" | bc -l)
      n=$(echo "s($B/100)" | bc -l)
      t=$(echo "$c * $h * $g - $f * $e" | bc -l)
      x=$((40 + 30 * $D * ($l * $h * $m - $t * $n) ))
      y=$((12 + 15 * $D * ($l * $h * $n + $t * $m) ))
      o=$((x + 80 * y))
      N=$((8 * (($f * $e - $c * $d * $g) * $m - $c * $d * $e - $f * $g - $l * $d * $n) ))
      if [[ $y -gt 0 && $y -lt 22 && $x -gt 0 && $x -lt 80 ]]; then
        b[$o]="${b[$o]}${chars:$((N>0?N:0)):1}"
      fi
    done
  done

  for (( k=0; k<1760; k++ )); do
    if [[ $((k % 80)) -eq 0 ]]; then
      echo
    fi
    echo -n "${b[$k]:- }"
  done

  A=$(echo "$A + 0.04" | bc -l)
  B=$(echo "$B + 0.02" | bc -l)
  sleep 0.05
done