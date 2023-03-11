package dynamic_array

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDynamicArray(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dynamic Array Test Suite")
}

var _ = Describe("DynamicArray", func() {
	var (
		da DynamicArray
	)

	Context("Initialization", func() {
		da = NewDynamicArray(1, 2, 3)

		It("it creates an empty linked list", func() {
			Expect(da).To(Equal(DynamicArray{
				Size:     3,
				Capacity: 4,
				Elements: []int{1, 2, 3},
			}))
		})
	})

	Context("Insert", func() {
		When("inserts one element in array", func() {
			BeforeEach(func() {
				da = Append(da, 4)
			})

			It("grows capacity", func() {
				Expect(da).To(Equal(DynamicArray{
					Size:     4,
					Capacity: 6,
					Elements: []int{1, 2, 3, 4},
				}))
			})
		})

		When("double the size of the array", func() {
			BeforeEach(func() {
				da = Append(da, 5, 6, 7, 8)
			})

			It("grows capacity based on new size", func() {
				Expect(da).To(Equal(DynamicArray{
					Size:     8,
					Capacity: 12,
					Elements: []int{1, 2, 3, 4, 5, 6, 7, 8},
				}))
			})
		})

		When("insert two elements", func() {
			BeforeEach(func() {
				da = Append(da, 9, 10)
			})

			It("does not grow capacity", func() {
				Expect(da).To(Equal(DynamicArray{
					Size:     10,
					Capacity: 12,
					Elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				}))
			})
		})
	})
})
