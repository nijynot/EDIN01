package main

import (
  "fmt"
  "strings"
  // "reflect"
  "strconv"
  "os"
  "io"
  // "math"
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

func hamming(u []int, z []int) int {
  if len(u) != len(z) {
    panic("Error: `u` and `z` not of the same length")
    return -1
  }

  d := 0
  for i := 0; i < len(u); i++ {
    if u[i] != z[i] {
      d = d + 1
    }
  }

  return d
}

func LFSR(poly []int, state *[]int, n int) (out int, in int) {
  for i := 0; i < len(poly); i++ {
    in = in - poly[i] * (*state)[i]
  }

  out = (*state)[0]
  in = mod(in, n)
  *state = append((*state)[1:], in)

  return out, in
}

func SeqSplit(s string) []int {
  a := strings.Split(s, "")
  b := make([]int, len(a))
  for i, v := range a {
    b[i], _ = strconv.Atoi(v)
  }
  return b
}

func Cycle(p []int, state *[]int, clock int) []int {
  seq := make([]int, 0)

  for i := 0; i < clock; i++ {
    out, _ := LFSR(p, state, 2)
    seq = append(seq, out)
  }

  return seq
}

func Min(s []int) int {
  var min int
  for i, v := range s {
    if i == 0 || v < min {
      min = v
    }
  }

  return min
}

func Generator(p []int, init []int, size int) [][]int {
  zero := make([]int, len(p))
  generator := init
  trials := [][]int{zero}
  for i := 0; i < size; i++ {
    trials = append(trials, generator)
    LFSR(p, &generator, 2)
  }
  return trials
}

// func MinimizeP(trials [][]int) key []int {
//
// }

func main() {
  z := SeqSplit("1100100111111111010110110100001101001101001101111000011100111001011101111000100110011110010101011010011100110001010010100000101101010010011001001101101110110001010101010110100110100111010011011")
  C1 := []int{1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1}
  C2 := []int{1, 0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 0}
  C3 := []int{1, 1, 0, 0, 1, 0, 0, 1, 0, 8, 0, 0, 1, 1, 0, 1, 0}

  N := len(z)

  // p(x) = x^13 + x^4 + x^3 + x^1 + 1
  // p13 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1}
  // generator13 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
  // trials13 := Generator(p13, generator, 8191)

  // p(x) = x^15 + x^1 + 1
  // p15 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
  // generator := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
  // trials := Generator(p15, generator, 32767)

  // p(x) = x^17 + x^3 + 1
  // p17 := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}
  // generator := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
  // trials := Generator(p17, generator, 131071)

  // // Find K1 with the least hamming distance
  // var minDistance int
  // var minU []int
  // for i := 0; i < len(trials); i++ {
  //   init := trials[i]
  //   u := Cycle(C1, &(trials[i]), N)
  //
  //   if i == 0 || hamming(u, z) < minDistance {
  //     minDistance = hamming(u, z)
  //     minU = init
  //   }
  // }
  // fmt.Println(minDistance, minU)

  // // Find K2 with the least hamming distance
  // var minDistance int
  // var minU []int
  // for i := 0; i < len(trials); i++ {
  //   init := trials[i]
  //   u := Cycle(C3, &(trials[i]), N)
  //
  //   fmt.Println(hamming(u, z))
  //
  //   if i == 0 || hamming(u, z) < minDistance {
  //     minDistance = hamming(u, z)
  //     minU = init
  //   }
  // }
  // fmt.Println(minDistance, minU)

  // trash
  // K2 := []int{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 0, 1}
  // K3 := []int{1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0, 0, 1}

  // K = (K1, K2, K3) with lowest hamming distance
  K1 := []int{1, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0}
  K2 := []int{1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1}
  K3 := []int{1, 1, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1}

  u_1 := Cycle(C1, &K1, N)
  u_2 := Cycle(C2, &K2, N)
  u_3 := Cycle(C3, &K3, N)

  fmt.Println(hamming(u_1, z))
  fmt.Println(hamming(u_2, z))
  fmt.Println(hamming(u_3, z))

  seq := make([]int, 0)

  for i := 0; i < 193; i++ {
    out1, _ := LFSR(C1, &K1, 2)
    out2, _ := LFSR(C2, &K2, 2)
    out3, _ := LFSR(C3, &K3, 2)

    // fmt.Println(out1 + out2 + out3)

    if out1 + out2 + out3 > 1 {
      seq = append(seq, 1)
    } else {
      seq = append(seq, 0)
    }
  }

  fmt.Println(seq)
  fmt.Println(hamming(seq, z))

  // if err := WriteStringToFile(
  //   "input",
  //   strings.Trim(strings.Join(strings.Fields(fmt.Sprint(seq)), ""), "[]")); err != nil {
  //   panic(err)
  // }

  // Testing
  // C1 := []int{1, 0, 0, 1}
  // state := []int{0, 0, 0, 1}
  // N := 16

  // u := make([]int, 0)
  //
  // for i := 0; i < N; i++ {
  //   out, _ := LFSR(C1, &state, 2)
  //   u = append(u, out)
  //   // fmt.Println(state)
  // }
}
