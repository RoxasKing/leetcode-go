package main

import (
	"sort"
	"strconv"
	"strings"
)

// Tags:
// Polynomial Class

func basicCalculatorIV(expression string, evalvars []string, evalints []int) []string {
	evalMap := map[string]int{}
	for i := range evalvars {
		evalMap[evalvars[i]] = evalints[i]
	}
	return Parse(expression).Evaluate(evalMap).ToList()
}

type Polynomial map[string]int

func MakePolynomial(expr string) Polynomial {
	h := map[string]int{}
	if isNumber(expr[0]) {
		num, _ := strconv.Atoi(expr)
		h[""] += num
	} else {
		h[expr]++
	}
	return h
}

func Combine(this, that Polynomial, op byte) Polynomial {
	switch op {
	case '+':
		return this.Add(that)
	case '-':
		return this.Sub(that)
	case '*':
		return this.Mul(that)
	}
	return nil
}

func Parse(expr string) Polynomial {
	n := len(expr)
	polys := []Polynomial{}
	ops := []byte{}
	for i := 0; i < n; i++ {
		c := expr[i]
		if c == '(' {
			pa, j := 0, i
			for ; j < n; j++ {
				if expr[j] == '(' {
					pa++
				} else if expr[j] == ')' {
					pa--
				}
				if pa == 0 {
					break
				}
			}
			polys = append(polys, Parse(expr[i+1:j]))
			i = j
		} else if isLetterOrNumber(c) {
			j := i
			ok := true
			for ; j < n; j++ {
				if expr[j] == ' ' {
					polys = append(polys, MakePolynomial(expr[i:j]))
					ok = false
					break
				}
			}
			if ok {
				polys = append(polys, MakePolynomial(expr[i:]))
			}
			i = j
		} else if c != ' ' {
			ops = append(ops, c)
		}
	}

	if len(polys) == 0 {
		return nil
	}

	for j := len(ops) - 1; j >= 0; j-- {
		if ops[j] == '*' {
			polys[j] = Combine(polys[j], polys[j+1], ops[j])
			copy(polys[j+1:], polys[j+2:])
			polys = polys[:len(polys)-1]
			copy(ops[j:], ops[j+1:])
			ops = ops[:len(ops)-1]
		}
	}

	out := polys[0]
	for j := range ops {
		out = Combine(out, polys[j+1], ops[j])
	}
	return out
}

func isLetterOrNumber(c byte) bool {
	return isLetter(c) || isNumber(c)
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isNumber(c byte) bool {
	return '0' <= c && c <= '9'
}

func (this Polynomial) Add(that Polynomial) Polynomial {
	h := map[string]int{}
	for k, v := range this {
		h[k] += v
	}
	for k, v := range that {
		h[k] += v
	}
	return h
}

func (this Polynomial) Sub(that Polynomial) Polynomial {
	h := map[string]int{}
	for k, v := range this {
		h[k] += v
	}
	for k, v := range that {
		h[k] -= v
	}
	return h
}

func (this Polynomial) Mul(that Polynomial) Polynomial {
	h := map[string]int{}
	for k1, v1 := range this {
		for k2, v2 := range that {
			ks1 := strings.Split(k1, "$")
			ks2 := strings.Split(k2, "$")
			ks := make([]string, 0, len(ks1)+len(ks2))
			for _, k := range ks1 {
				if k == "" {
					continue
				}
				ks = append(ks, k)
			}
			for _, k := range ks2 {
				if k == "" {
					continue
				}
				ks = append(ks, k)
			}
			sort.Strings(ks)
			h[strings.Join(ks, "$")] += v1 * v2
		}
	}
	return h
}

func (this Polynomial) Evaluate(evalMap map[string]int) Polynomial {
	h := map[string]int{}
	for k, v := range this {
		ks := []string{}
		for _, x := range strings.Split(k, "$") {
			if val, ok := evalMap[x]; ok {
				v *= val
			} else {
				ks = append(ks, x)
			}
		}
		h[strings.Join(ks, "$")] += v
	}
	return h
}

func (this Polynomial) ToList() []string {
	out := []string{}
	ks := make([]string, 0, len(this))
	for k := range this {
		ks = append(ks, k)
	}
	sort.Slice(ks, func(i, j int) bool {
		if ks[i] == "" {
			return false
		} else if ks[j] == "" {
			return true
		}
		ss1 := strings.Split(ks[i], "$")
		ss2 := strings.Split(ks[j], "$")
		if len(ss1) != len(ss2) {
			return len(ss1) > len(ss2)
		}
		for i := range ss1 {
			if ss1[i] != ss2[i] {
				return ss1[i] < ss2[i]
			}
		}
		return true
	})
	for _, k := range ks {
		v := this[k]
		if v == 0 {
			continue
		}
		s := strconv.Itoa(v)
		for _, x := range strings.Split(k, "$") {
			if x == "" {
				continue
			}
			s += "*" + x
		}
		out = append(out, s)
	}
	return out
}
