package set

// IComparable interface defines a strict order relation which depends
// on a context c.
//
// 'x <_c y' iff 'x.CompareTo(c, y) < 0'
type IComparable[C any, T any] interface {

	// CompareTo compares receiver and method's argument.
	// result:
	//   < 0: receiver is less than method's argument;
	//   = 0: receiver equals method's argument;
	//   > 0: receiver is greater than method's argument
	CompareTo(C, T) int
}
