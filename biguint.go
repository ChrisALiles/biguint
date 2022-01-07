package biguint

// An exercise to help me to learn the basics of Go.
// The arithmetic uses the algorithms I learned in primary school - long division etc.
// I resisted the urge to call it "bigunit" with great difficulty.

import (
	"fmt"
	"strconv"
)

// biguints are implemnted as slices of ints.
type biguint []int

// String is the standard string function for biguint.
func (b biguint) String() string {
	ret := ""
	for _, i := range b {
		ret += strconv.Itoa(i)
	}
	return ret
}

// add adds biguints.
func (b1 biguint) add(bn ...biguint) biguint {
	work := b1
	for _, s := range bn {
		work = add2(work, s)
	}
	return work
}

// add2 adds 2 biguints (not a method).
func add2(s1, s2 biguint) biguint {
	var arg1, arg2 biguint
	// Make the 2 slices equal length to simplify the addition.
	switch {
	case len(s1) == len(s2):
		arg1 = s1
		arg2 = s2
	case len(s1) < len(s2):
		arg1 = append(make([]int, len(s2)-len(s1)), s1...)
		arg2 = s2
	case len(s1) > len(s2):
		arg1 = s1
		arg2 = append(make([]int, len(s1)-len(s2)), s2...)
	}
	carry := 0
	ans := make(biguint, len(arg1))
	for i := len(arg1) - 1; i >= 0; i-- {
		sum := arg1[i] + arg2[i] + carry
		ans[i] = sum % 10
		carry = sum / 10
	}
	if carry == 1 {
		return append([]int{1}, ans...)
	} else {
		return ans
	}
}

// z9 removes leading zeros from a biguint (z9 from the COBOL PICTURE clause).
func (b biguint) z9() biguint {
	nzi := -1
	for i, dig := range b {
		if dig != 0 {
			nzi = i
			break
		}
	}
	if nzi == -1 {
		nzi = len(b) - 1
	}
	return b[nzi:]
}

// compare compares two biguints for <, >, =.
func (b1 biguint) compare(b2 biguint) string {
	v1 := b1.z9()
	v2 := b2.z9()
	switch {
	case len(v1) < len(v2):
		return "<"
	case len(v1) > len(v2):
		return ">"
	}
	for i, d1 := range v1 {
		switch {
		case d1 < v2[i]:
			return "<"
		case d1 > v2[i]:
			return ">"
		}
	}
	return "="
}

// greater returns true if b1 > b2.
func (b1 biguint) greater(b2 biguint) bool {
	if b1.compare(b2) == ">" {
		return true
	} else {
		return false
	}
}

// ge returns true if b1 >= b2.
func (b1 biguint) ge(b2 biguint) bool {
	if b1.compare(b2) == ">" || b1.compare(b2) == "=" {
		return true
	} else {
		return false
	}
}

// less returns true if b1 < b2.
func (b1 biguint) less(b2 biguint) bool {
	if b1.compare(b2) == "<" {
		return true
	} else {
		return false
	}
}

// le returns true if b1 <= b2.
func (b1 biguint) le(b2 biguint) bool {
	if b1.compare(b2) == "<" || b1.compare(b2) == "=" {
		return true
	} else {
		return false
	}
}

// equal returns true if b1 = b2.
func (b1 biguint) equal(b2 biguint) bool {
	if b1.compare(b2) == "=" {
		return true
	} else {
		return false
	}
}

// subtract subtracts biguints.
func (a1 biguint) subtract(a2 biguint) biguint {
	var arg1, arg2 biguint
	b1 := a1.z9()
	b2 := a2.z9()
	// Make the 2 slices equal length to simplify the subtraction.
	// Also make arg1 the larger if the biguints are not equal.
	switch {
	case len(b1) == len(b2):
		if b1.greater(b2) {
			arg1 = b1
			arg2 = b2
		} else {
			arg1 = b2
			arg2 = b1
		}
	case len(b1) < len(b2):
		arg2 = append(make([]int, len(b2)-len(b1)), b1...)
		arg1 = b2
	case len(b1) > len(b2):
		arg1 = b1
		arg2 = append(make([]int, len(b1)-len(b2)), b2...)
	}
	carry := 0
	ans := make(biguint, len(arg1))
	for i := len(arg1) - 1; i >= 0; i-- {
		sum := arg2[i] + carry
		if sum > arg1[i] {
			carry = 1
			ans[i] = 10 + arg1[i] - sum
		} else {
			carry = 0
			ans[i] = arg1[i] - sum
		}
	}
	return ans.z9()
}

// times multiplies two biguints.
func (b1 biguint) times(b2 biguint) biguint {
	var m1, m2 biguint // m2 the multiplier, m1 the multiplicand
	if len(b1.z9()) < len(b2.z9()) {
		m2 = b1.z9()
		m1 = b2.z9()
	} else {
		m1 = b1.z9()
		m2 = b2.z9()
	}
	ans := make(biguint, len(m1)+len(m2))
	for mx := len(m2) - 1; mx >= 0; mx-- {
		work := make(biguint, len(m1)+len(m2))
		carry := 0
		wx := len(work) - (len(m2) - mx) // index into work - start one place further left each time.
		for dx := len(m1) - 1; dx >= 0; dx-- {
			prod := m1[dx]*m2[mx] + carry
			work[wx] = prod % 10
			carry = prod / 10
			wx--
		}
		work[wx] = carry
		ans = ans.add(work)
	}
	return ans.z9()
}

// divby divides two biguints.
func (b1 biguint) divby(b2 biguint) (biguint, bool) {
	var dor, dend biguint // divisor and dividend
	var work biguint
	if b1.less(b2) {
		dor = b1.z9()
		dend = b2.z9()
	} else {
		dend = b1.z9()
		dor = b2.z9()
	}
	if len(dor) == 1 && dor[0] == 0 {
		return nil, false // division by zero
	}
	ans := make(biguint, len(dend))
	start := len(dor) // start point of dividend the loop begins at
	wk := dend[:len(dor)]
	if dor.le(wk) {
		start -= 1
	}
	if start > 0 {
		work = dend[:start]
	}
	for _, dig := range dend[start:] {
		var prod biguint
		work = append(work, dig)
		mult := biguint{0} // here we do division by multiplication (because the divisor might be too big)
		for {
			mult[0] += 1
			prod = mult.times(dor)
			if prod.greater(work) {
				break
			}
		}
		mult[0] -= 1
		ans = append(ans, mult...)
		work = work.subtract(mult.times(dor))
	}
	return ans.z9(), true
}

// exp is biguint exponentiation.
func (b biguint) exp(e biguint) biguint {
	fmt.Println("Into exp")
	bigzero := biguint{0}
	bigone := biguint{1}
	rslt := biguint{1}
	count := e
	//	if e.equal(bigzero) {
	//		return bigone
	//	} else {
	//		return b.times(b.exp(e.subtract(bigone)))
	//	}
	for !count.equal(bigzero) {
		rslt = rslt.times(rslt)
		count = count.subtract(bigone)
	}
	return rslt
}

// strToBig converts a string of digits to a biguint.
func strToBig(s string) (biguint, error) {
	bi := make(biguint, len(s))
	for i, _ := range s {
		var err error
		bi[i], err = strconv.Atoi(s[i : i+1])
		if err != nil {
			return nil, err
		}
	}
	return bi, nil
}
