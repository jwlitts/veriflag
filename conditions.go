package veriflag

func InListint(list []int) func(x int) bool {
	return func(x int) bool {
		for _, elem := range list {
			if elem == x {
				return true
			}
		}
		return false
	}
}
func InRangeint(start int, end int) func(x int) bool {
	return func(x int) bool {
		return x > start && x < end
	}
}
func InListint32(list []int32) func(x int32) bool {
	return func(x int32) bool {
		for _, elem := range list {
			if elem == x {
				return true
			}
		}
		return false
	}
}
func InRangeint32(start int32, end int32) func(x int32) bool {
	return func(x int32) bool {
		return x > start && x < end
	}
}
func InListint64(list []int64) func(x int64) bool {
	return func(x int64) bool {
		for _, elem := range list {
			if elem == x {
				return true
			}
		}
		return false
	}
}
func InRangeint64(start int64, end int64) func(x int64) bool {
	return func(x int64) bool {
		return x > start && x < end
	}
}
func InListstring(list []string) func(x string) bool {
	return func(x string) bool {
		for _, elem := range list {
			if elem == x {
				return true
			}
		}
		return false
	}
}
func InRangestring(start string, end string) func(x string) bool {
	return func(x string) bool {
		return x > start && x < end
	}
}
func InListfloat32(list []float32) func(x float32) bool {
	return func(x float32) bool {
		for _, elem := range list {
			if elem == x {
				return true
			}
		}
		return false
	}
}
func InRangefloat32(start float32, end float32) func(x float32) bool {
	return func(x float32) bool {
		return x > start && x < end
	}
}
func InListfloat64(list []float64) func(x float64) bool {
	return func(x float64) bool {
		for _, elem := range list {
			if elem == x {
				return true
			}
		}
		return false
	}
}
func InRangefloat64(start float64, end float64) func(x float64) bool {
	return func(x float64) bool {
		return x > start && x < end
	}
}
func InListuint(list []uint) func(x uint) bool {
	return func(x uint) bool {
		for _, elem := range list {
			if elem == x {
				return true
			}
		}
		return false
	}
}
func InRangeuint(start uint, end uint) func(x uint) bool {
	return func(x uint) bool {
		return x > start && x < end
	}
}
func InListuint32(list []uint32) func(x uint32) bool {
	return func(x uint32) bool {
		for _, elem := range list {
			if elem == x {
				return true
			}
		}
		return false
	}
}
func InRangeuint32(start uint32, end uint32) func(x uint32) bool {
	return func(x uint32) bool {
		return x > start && x < end
	}
}
func InListuint64(list []uint64) func(x uint64) bool {
	return func(x uint64) bool {
		for _, elem := range list {
			if elem == x {
				return true
			}
		}
		return false
	}
}
func InRangeuint64(start uint64, end uint64) func(x uint64) bool {
	return func(x uint64) bool {
		return x > start && x < end
	}
}
