package main

import (
	"testing"
)

func TestDuplicated(t *testing.T) {
	testCases := []struct {
		input       [10]int
		expect      bool
		description string
	}{
		{
			input:       [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expect:      false,
			description: "all count 0",
		},
		{
			input:       [10]int{0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
			expect:      false,
			description: "all count under 1",
		},
		{
			input:       [10]int{2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			expect:      false,
			description: "0 count duplicated",
		},
		{
			input:       [10]int{0, 0, 2, 0, 0, 0, 0, 0, 0, 0},
			expect:      true,
			description: "1~9 count duplicated",
		},
	}
	for _, testCase := range testCases {
		t.Run(
			testCase.description,
			func(t *testing.T) {
				result := isDuplicated(testCase.input)
				expect := testCase.expect
				if expect != result {
					t.Errorf("test case %v failed. expect result is %v but received %v", testCase.description, expect, result)
				}
			},
		)
	}
}

func TestVerified(t *testing.T) {
	testCases := []struct {
		input       Board
		expect      bool
		description string
	}{
		{
			input: Board{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expect:      true,
			description: "all matrix element value is 0",
		},
		{
			input: Board{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expect:      true,
			description: "matrix row each element value use 1~9 at once",
		},
		{
			input: Board{
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{2, 0, 0, 0, 0, 0, 0, 0, 0},
				{3, 0, 0, 0, 0, 0, 0, 0, 0},
				{4, 0, 0, 0, 0, 0, 0, 0, 0},
				{5, 0, 0, 0, 0, 0, 0, 0, 0},
				{6, 0, 0, 0, 0, 0, 0, 0, 0},
				{7, 0, 0, 0, 0, 0, 0, 0, 0},
				{8, 0, 0, 0, 0, 0, 0, 0, 0},
				{9, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expect:      true,
			description: "matrix column each element value use 1~9 at once",
		},
		{
			input: Board{
				{1, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expect:      false,
			description: "matrix row element value duplicated",
		},
		{
			input: Board{
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expect:      false,
			description: "matrix column element value duplicated",
		},
	}
	for _, testCase := range testCases {
		t.Run(
			testCase.description,
			func(t *testing.T) {
				result := verify(testCase.input)
				expect := testCase.expect
				if !expect == result {
					t.Errorf("test case %v failed. expect result is %v but received %v", testCase.description, expect, result)
				}
			},
		)
	}
}

func TestBackTrack(t *testing.T) {
	testCases := []struct {
		input       Board
		expect      bool
		description string
	}{
		{
			input: Board{
				{0, 0, 0, 9, 0, 0, 7, 4, 0},
				{8, 0, 1, 0, 0, 0, 0, 0, 0},
				{6, 0, 7, 0, 0, 0, 5, 0, 3},
				{0, 3, 0, 0, 4, 0, 0, 5, 0},
				{5, 0, 0, 0, 8, 0, 3, 2, 0},
				{4, 0, 0, 0, 0, 1, 0, 0, 0},
				{0, 5, 0, 0, 0, 6, 0, 0, 0},
				{0, 0, 0, 0, 3, 0, 0, 1, 9},
				{1, 7, 8, 0, 0, 0, 0, 0, 0},
			},
			expect:      true,
			description: "shold be solved",
		}
	}
	for _, testCase := range testCases {
		t.Run(
			testCase.description,
			func(t *testing.T) {
				result := backtrack(&testCase.input)
				expect := testCase.expect
				if expect != result {
					t.Errorf("test case %v failed. expect result is %v but received %v", testCase.description, expect, result)
				}
			},
		)
	}
}
