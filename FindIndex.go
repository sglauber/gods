import "fmt"

const size = 5;

func FindIndexSlice(haystack []int, needle int) int {
	for i, value range haystack {
		if value == needle {
			return i;
		}
	}

	return -1;
}

/* an example using a fixed size for arrays since 
slices are dynamic structures that points to an underlying array
*/

func FindIndexArray(haystack [size]int, needle int) int {
	for i := 0; i < size; i++ {
		if haystack[i] == needle {
			return i;
		}
	}

	return -1;
}

func FindIndexAddress(h *[size]int, needle int) (int, *int) {
	for i := 0; i < size; i++ {
		if h[i] == needle {
			return i, &h[i];
		}
	}
	
	return -1, nil;
}

/*
	h := [size]int{1,2,3,4,5}
	v := findIndex(h, 5)
	
	fmt.Printf("h[v]: %v \n", v)
	
	h := []int{0,9,8,7,6,55555,4444,333,22,1}
	v := findIndex2(h, 7)
	
	fmt.Printf("h[v]: %v", v)

	h := [size]int{1,2,3,4,5}
	val := findIndexPointer(h, 7)
*/
