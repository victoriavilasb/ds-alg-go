package dynamic_array

const GROWTH_FACTOR float64 = 1.5

type DynamicArray struct {
	Size     int
	Capacity int
	Elements []int
}

// A DynamicArray grows based on the initial elements provided. Initially we're
// going to start the array with the elements provided by the user
func NewDynamicArray(elems ...int) DynamicArray {
	cap := int(float64(len(elems)) * GROWTH_FACTOR)
	daElements := make([]int, 0, cap)
	daElements = append(daElements, elems...)

	return DynamicArray{
		Size:     len(elems),
		Capacity: int(float64(len(elems)) * GROWTH_FACTOR),
		Elements: daElements,
	}
}

func Append(da DynamicArray, elems ...int) DynamicArray {
	newSize := da.Size + len(elems)
	newCap := int(float64(newSize) * GROWTH_FACTOR)

	if newSize >= da.Capacity {
		newArr := make([]int, 0, newCap)

		newArr = append(newArr, da.Elements...)
		newArr = append(newArr, elems...)
		da.Size = newSize
		da.Capacity = newCap
		da.Elements = newArr

		return da
	}

	da.Size = newSize
	da.Elements = append(da.Elements, elems...)
	return da
}
