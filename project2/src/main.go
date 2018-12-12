package main

import (
  "fmt"
  "io"
	"os"
	"strings"
)

func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}

	return nil
}

func mod(a int, n int) int {
  if a % n < 0 {
    return (a % n) + n
  } else {
    return a % n
  }
}


// Non-linear FSRs with 0-state
func LFSR2(poly []int, state[]int, n int) (out int, in int) {
  for i := 0; i < len(poly); i++ {
    in = in - poly[i] * state[i]
  }

  if (state[0] != 0 &&
    state[1] == 0 &&
    state[2] == 0 &&
    state[3] == 0) {
    return 1, 0
  } else if (state[1] == 0 &&
    state[2] == 0 &&
    state[3] == 0) {
    return 0, 1
  } else {
    return state[0], mod(in, n)
  }
}

func LFSR5(poly []int, state[]int, n int) (out int, in int) {
  for i := 0; i < len(poly); i++ {
    in = in - poly[i] * state[i]
  }

  if (
    state[0] == 2 &&
    state[1] == 0 &&
    state[2] == 0 &&
    state[3] == 0) {
    return 2, 0
  } else if (
    state[0] == 0 &&
    state[1] == 0 &&
    state[2] == 0 &&
    state[3] == 0) {
    return 0, 1
  } else {
    return state[0], mod(in, n)
  }
}

// Bijective function phi: Z_2 x Z_5 -> Z_10
func phi(x int, y int) int {
  return 5 * x + y
}

func main() {
  p := []int{1, 0, 0, 1}
  q := []int{2, 2, 1, 0}
  state2 := []int{0, 0, 0, 0}
  state5 := []int{0, 0, 0, 0}
  seq := make([]int, 0)

  for i := 0; i < 10003; i++ {
    out2, in2 := LFSR2(p, state2, 2)
    out5, in5 := LFSR5(q, state5, 5)

    state2 = append(state2[1:], in2)
    state5 = append(state5[1:], in5)

    seq = append(seq, phi(out2, out5))
  }

  if err := WriteStringToFile(
    "input",
    strings.Trim(strings.Join(strings.Fields(fmt.Sprint(seq)), ""), "[]")
  );
  err != nil {
    panic(err)
  }
}
