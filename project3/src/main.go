package main

import (
  "fmt"
  "strings"
  // "reflect"
  "strconv"
  "os"
  "io"
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

func main() {
  z := SeqSplit("1100100111111111010110110100001101001101001101111000011100111001011101111000100110011110010101011010011100110001010010100000101101010010011001001101101110110001010101010110100110100111010011011")
  C1 := []int{1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1}
  state := []int{1, 1, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1}
  N := len(z)

  // C1 := []int{1, 0, 0, 1}
  // state := []int{0, 0, 0, 1}
  // N := 16

  u := make([]int, 0)

  for i := 0; i < N; i++ {
    out, _ := LFSR(C1, &state, 2)
    u = append(u, out)
    // fmt.Println(state)
  }

  // fmt.Println(z[0])
  // fmt.Println(hamming(u, z))
  fmt.Println(u)
}
