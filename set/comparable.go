package set

// IComparable interface defines a strict order relation.
// 'x < y' iff 'x.CompareTo(y) < 0'
type IComparable[T any] interface {

	// CompareTo compares receiver and method's argument.
	// result:
	//   < 0: receiver is less than method's argument;
	//   = 0: receiver equals method's argument;
	//   > 0: receiver is greater than method's argument
	CompareTo(T) int
}
