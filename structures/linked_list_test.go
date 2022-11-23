package structures

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LinkedList", func() {
	var (
		ll *LinkedList
	)

	Context("Initialization", func() {
		ll = NewLinkedList(2)

		It("it creates an empty linked list", func() {
			Expect(ll).To(BeComparableTo(&LinkedList{
				Head:     nil,
				Tail:     nil,
				Size:     0,
				Max_size: 2,
			}))
		})
	})

	Context("Insert in list with max 3 elements", func() {
		BeforeEach(func() {
			ll = NewLinkedList(3)
		})

		When("inserts one element in list", func() {
			BeforeEach(func() {
				ll.InsertInFront(1)
			})

			It("inserts element in head", func() {
				Expect(ll.Head.elem).To(Equal(1))
			})
		})

		When("inserts two elements in list", func() {
			BeforeEach(func() {
				ll.InsertInFront(1)
				ll.InsertInFront(2)
			})

			It("transforms 1 in last element", func() {
				Expect(ll.Tail.elem).To(Equal(1))
			})
		})

		When("inserts 4 elements in list", func() {
			var (
				err error
			)

			BeforeEach(func() {
				ll.InsertInFront(1)
				ll.InsertInFront(2)
				ll.InsertInFront(3)
				err = ll.InsertInFront(4)
			})

			It("transforms 1 in last element", func() {
				Expect(err).NotTo(BeNil())
			})
		})
	})

	Context("Remove elements", func() {
		BeforeEach(func() {
			ll = NewLinkedList(2)
			ll.InsertInFront(1)
			ll.InsertInFront(2)
		})

		When("removes element in head", func() {
			It("changes head to tail", func() {
				ll.Remove(2)
				Expect(ll.Head.elem).To(Equal(1))
			})
		})

		When("insert node in front", func() {
			It("changes head to tail", func() {
				ll.Remove(2)
				ll.InsertNodeInFrom(&Node{
					elem: 5,
				})
				Expect(ll.Head.elem).To(Equal(5))
			})
		})
	})

	Context("Insert, remove and compare results", func() {
		BeforeEach(func() {
			ll = NewLinkedList(5)
			ll.InsertInFront(1)
			ll.InsertInFront(2)
		})

		When("transform elements in slice", func() {
			It("returns a slice with the correct order of elements", func() {
				elems := ll.GetElements()

				Expect(elems).To(Equal([]int{2, 1}))
			})
		})

		When("element is removed from list", func() {
			It("returns an error when trying to find element", func() {
				ll.Remove(2)
				ok := ll.Find(2)

				Expect(ok).To(BeFalse())
			})
		})
	})
})
