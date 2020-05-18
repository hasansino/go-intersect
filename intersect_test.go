package intersect

import (
	"reflect"
	"testing"
)

func TestInt64(t *testing.T) {
	var (
		a = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b = []int64{5, 6, 7, 11, 12, 13, 14, 15}
		x []int64
	)

	x = Int64(a, b)
	if !reflect.DeepEqual(x, []int64{5, 6, 7}) {
		t.Errorf("Int32() returned invalid values of intersection %v", x)
	}

	x = Int64(a, nil)
	if len(x) > 0 {
		t.Errorf("Int32() returned invalid values of intersection %v", x)
	}

	x = Int64(nil, nil)
	if len(x) > 0 {
		t.Errorf("Int32() returned invalid values of intersection %v", x)
	}

	x = Int64([]int64{}, []int64{})
	if len(x) > 0 {
		t.Errorf("Int32() returned invalid values of intersection %v", x)
	}
}

func BenchmarkInt64(b *testing.B) {
	var (
		x = []int64{1, 2, 3, 4, 5, 11, 6, 7, 8, 9, 10}
		y = []int64{11, 12, 13, 14, 15, 5, 6, 7}
	)

	for i := 0; i < b.N; i++ {
		_ = Int64(x, y)
	}

	b.ReportAllocs()
}

func TestString(t *testing.T) {
	var (
		a = []string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i",
			"j", "k", "l", "m", "o", "p", "q", "r", "s",
			"t", "u", "v", "w", "x", "y", "z"}
		b = []string{
			"xa", "xb", "xc",
			"w", "x", "y", "z",
			"qa", "qb", "qc",
		}
		x []string
	)

	x = String(a, b)
	if !reflect.DeepEqual(x, []string{"w", "x", "y", "z"}) {
		t.Errorf("String() returned invalid values of intersection %v", x)
	}

	x = String(a, nil)
	if len(x) > 0 {
		t.Errorf("String() returned invalid values of intersection %v", x)
	}

	x = String(nil, nil)
	if len(x) > 0 {
		t.Errorf("String() returned invalid values of intersection %v", x)
	}

	x = String([]string{}, []string{})
	if len(x) > 0 {
		t.Errorf("String() returned invalid values of intersection %v", x)
	}
}

func BenchmarkString(b *testing.B) {
	var (
		x = []string{
			"a", "b", "c", "d", "e", "f", "g", "h", "i",
			"j", "k", "l", "m", "o", "p", "q", "r", "s",
			"t", "u", "v", "w", "x", "y", "z"}
		y = []string{
			"xa", "xb", "xc",
			"w", "x", "y", "z",
			"qa", "qb", "qc",
		}
	)

	for i := 0; i < b.N; i++ {
		_ = String(x, y)
	}

	b.ReportAllocs()
}
