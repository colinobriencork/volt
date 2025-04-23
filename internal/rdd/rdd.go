package rdd

// RDD represents a Resilient Distributed Dataset with generic type T.
type RDD[T any] struct {
	data []T
}

// NewRDD creates a new RDD from a slice of data.
func NewRDD[T any](data []T) *RDD[T] {
	return &RDD[T]{data: data}
}

// Map applies a transformation function to each element in the RDD.
func (r *RDD[T]) Map(f func(T) T) *RDD[T] {
	result := make([]T, len(r.data))
	for i, v := range r.data {
		result[i] = f(v)
	}
	return NewRDD(result)
}

// Filter filters elements in the RDD based on a predicate function.
func (r *RDD[T]) Filter(f func(T) bool) *RDD[T] {
	result := []T{}
	for _, v := range r.data {
		if f(v) {
			result = append(result, v)
		}
	}
	return NewRDD(result)
}

// Collect returns the data in the RDD as a slice.
func (r *RDD[T]) Collect() []T {
	return r.data
}
