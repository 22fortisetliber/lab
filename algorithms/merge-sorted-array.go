func merge(nums1 []int, m int, nums2 []int, n int) {
	var out []int
	n1Index := m
	n2Index := 0
	for n1Index < m && n2Index < n {
		if nums1[n1Index] < nums2[n2Index] {
			if nums1[n1Index] != 0 {
				out = append(out, nums1[n1Index])
			}
			n1Index++
		} else {
			if nums2[n2Index] != 0 {
				out = append(out, nums2[n2Index])
			}
			n2Index++
		}
	}

    
	for n1Index < m {
		if nums1[n1Index] != 0 {
			out = append(out, nums1[n1Index])
		}
		n1Index++
	}
	for n2Index < n {
		if nums2[n2Index] != 0 {
			out = append(out, nums2[n2Index])
		}
		n2Index++
	}
    nums1 = out
}