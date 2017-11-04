// biguint implements the standard arithmetic operators - addition, subtraction, multiplication, division - for
// arbitrary-length unsigned integers.
// Arguments are passed and results are returned as strings of integers.
package biguint

import (
	"fmt"
)

// Add returns the sum of any number of unsigned integers.
func Add(strn ...string) (string, error) {
	nums, err := cnvargs(strn...)
	if err != nil {
		return "", err
	}
	res := nums[0].add(nums[1:]...)
	return res.String(), nil
}

// Subtract returns the absolute difference between two unsigned integers.
func Subtract(str1, str2 string) (string, error) {
	nums, err := cnvargs(str1, str2)
	if err != nil {
		return "", err
	}
	res := nums[0].subtract(nums[1])
	return res.String(), nil
}

// Multiply returns the product of two unsigned integers.
func Multiply(str1, str2 string) (string, error) {
	nums, err := cnvargs(str1, str2)
	if err != nil {
		return "", err
	}
	res := nums[0].times(nums[1])
	return res.String(), nil
}

// Divide returns the result of integer division of two unsigned integers.
// The smaller number is the divisor.
func Divide(str1, str2 string) (string, error) {
	nums, err := cnvargs(str1, str2)
	if err != nil {
		return "", err
	}
	ans, ok := nums[0].divby(nums[1])
	if !ok {
		return "", fmt.Errorf("Attempted division by zero")
	}
	return ans.String(), nil
}

// cnvargs converts the arguments - slice of strings (of integers) - into a slice of biguints.
func cnvargs(strn ...string) ([]biguint, error) {
	if len(strn) < 2 {
		return nil, fmt.Errorf("Need at least 2 numbers")
	}
	errs := make([]error, len(strn))
	ans := make([]biguint, len(strn))
	for indx, str := range strn {
		ans[indx], errs[indx] = strToBig(str)
	}
	var err error
	for indx, e := range errs {
		if e != nil {
			if err != nil {
				err = fmt.Errorf("%v not recognised as an unsigned integer: %v \n", strn[indx], err)
			} else {
				err = fmt.Errorf("%v not recognised as an unsigned integer: \n", strn[indx])
			}
		}
	}
	if err != nil {
		return nil, err
	}
	return ans, nil
}
