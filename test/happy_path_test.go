package test

import (
	"cronParser/expression_parser"
	"cronParser/output_formatter"
	"reflect"
	"testing"
)

func TestRunTests(t *testing.T) {
	var tests = []struct {
		testName string
		fieldExpr      string
		expectedOutput [][]string
		errorExpected bool
	}{
		{
			"test1",
			"*/15 0 1,15 * 1-5 /usr/bin/find",
			[][]string{
				{"minute        0 15 30 45 "},
				{"hour          0 "},
				{"day of month  1 15 "},
				{"month         1 2 3 4 5 6 7 8 9 10 11 12 "},
				{"day of week   1 2 3 4 5 "},
				{"command       /usr/bin/find"},
			},
			false,
		},
		{
			"test2",
			"*/15 ? 1,15 * 1-5 /usr/bin/find",
			[][]string{
				{"minute        0 15 30 45 "},
				{"hour          "},
				{"day of month  1 15 "},
				{"month         1 2 3 4 5 6 7 8 9 10 11 12 "},
				{"day of week   1 2 3 4 5 "},
				{"command       /usr/bin/find"},
			},
			false,
		},
		{
			"test3",
			"*/* 0 1,15 * 5-1 /usr/bin/find",
			[][]string{},
			true,
		},
		{
			"test4",
			"*/15 ? 1,15 * 1-50 /usr/bin/find",
			[][]string{},
			true,
		},
		{
			"test5",
			"*/150 ? 1,15 * 1-5 /usr/bin/find",
			[][]string{},
			true,
		},
		{
			"test6",
			"*/15 ? 1,15 * 5-1 /usr/bin/find",
			[][]string{},
			true,
		},
		{
			"test7",
			"*/a ? 1,15 * 1-5 /usr/bin/find",
			[][]string{},
			true,
		},
		{
			"test8",
			"*/15 ?* 1,15 * 1-5 /usr/bin/find",
			[][]string{},
			true,
		},
		{
			"test9",
			"*/15 1? 1,15 * 1-5 /usr/bin/find",
			[][]string{},
			true,
		},
		{
			"test10",
			"*/15 1 1,15 * 1-5",
			[][]string{},
			true,
		},
		{
			"test11",
			"*/15 1 1,15 * /usr/bin/find",
			[][]string{},
			true,
		},
		{
			"test12",
			"*/15 abc 1,15 * /usr/bin/find",
			[][]string{},
			true,
		},
		{
			"test13",
			"*/59 ? 1,15 * 1-5 /usr/bin/find",
			[][]string{
				{"minute        0 59 "},
				{"hour          "},
				{"day of month  1 15 "},
				{"month         1 2 3 4 5 6 7 8 9 10 11 12 "},
				{"day of week   1 2 3 4 5 "},
				{"command       /usr/bin/find"},
			},
			false,
		},
		{
			"test14",
			"*/15 ? 1,15 * MON-FRI /usr/bin/find",
			[][]string{
				{"minute        0 15 30 45 "},
				{"hour          "},
				{"day of month  1 15 "},
				{"month         1 2 3 4 5 6 7 8 9 10 11 12 "},
				{"day of week   1 2 3 4 5 "},
				{"command       /usr/bin/find"},
			},
			false,
		},
		{
			"test15",
			"*/15 ? 1,15 JAN-FEB MON-FRI /usr/bin/find",
			[][]string{
				{"minute        0 15 30 45 "},
				{"hour          "},
				{"day of month  1 15 "},
				{"month         1 2 "},
				{"day of week   1 2 3 4 5 "},
				{"command       /usr/bin/find"},
			},
			false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			out, err := expression_parser.ParseCronExpression(tt.fieldExpr)
			if err != nil {
				if tt.errorExpected {
					t.Log(err)
					return
				}
				t.Error(err)
			}
			if tt.errorExpected && err == nil {
				t.Errorf("Error expected but was nil")
			}
			fout := output_formatter.FormatOutput(out)
			for i, op:= range fout {
				if !reflect.DeepEqual(op, tt.expectedOutput[i]) {
					t.Errorf("expect val %+v, inistead got %+v", tt.expectedOutput[i], op)
				}
			}

		})
	}

}
