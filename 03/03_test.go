package quest03_test

import (
	"fmt"

	. "github.com/Evokoo/EBC_2026_Go/03"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	{part: 1, file: "inputs/example_I.txt", target: 1478},
	{part: 1, file: "inputs/input_I.txt", target: 12752},
	{part: 2, file: "inputs/example_II.txt", target: 1161463},
	{part: 2, file: "inputs/input_II.txt", target: 30589365515},
	{part: 3, file: "inputs/example_III.txt", target: 50987},
	{part: 3, file: "inputs/input_III.txt", target: 395890},
}

var _ = Describe("EBC_2026_Go - 03", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
