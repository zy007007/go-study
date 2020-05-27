func maxArea(height []int) int {
    i, j := 0, len(height)-1
    res := 0
    for i < j {
        h := min(height[i], height[j])
        res = max(res, h*(j-i))
        if height[i] < height[j] {
            i ++
        } else {
            j --
        }
    }
    return res
}

func min(x, y int) int {
    if x > y {
        return y
    }
    return x
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
