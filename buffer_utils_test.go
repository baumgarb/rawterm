package rawterm

import "testing"

func TestRemoveCharAt(t *testing.T) {
	type test struct {
		xr     []rune
		pos    int
		expect []rune
	}

	tt := []test{
		{[]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 3, []rune{'a', 'b', 'c', 'e', 'f'}},
		{[]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 0, []rune{'b', 'c', 'd', 'e', 'f'}},
		{[]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 5, []rune{'a', 'b', 'c', 'd', 'e'}},
	}

	for i, td := range tt {
		result := removeCharAt(td.xr, td.pos)
		if len(td.expect) != len(result) {
			t.Errorf("#%v: expected %v (len %v), got %v (len %v)", i, td.expect, len(td.expect), result, len(result))
			continue
		}
		for j := 0; j < len(result); j++ {
			if td.expect[j] != result[j] {
				t.Errorf("#%v: expected %v (len %v), got %v (len %v)", i, td.expect, len(td.expect), result, len(result))
				break
			}
		}
	}
}

func TestInsertCharAt(t *testing.T) {
	type test struct {
		xr     []rune
		pos    int
		r      rune
		expect []rune
	}

	tt := []test{
		{[]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 3, 'z', []rune{'a', 'b', 'c', 'z', 'd', 'e', 'f'}},
		{[]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 0, 'z', []rune{'z', 'a', 'b', 'c', 'd', 'e', 'f'}},
		{[]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 5, 'z', []rune{'a', 'b', 'c', 'd', 'e', 'z', 'f'}},
		{[]rune{'a', 'b', 'c', 'd', 'e', 'f'}, 6, 'z', []rune{'a', 'b', 'c', 'd', 'e', 'f', 'z'}},
	}

	for i, td := range tt {
		result := insertCharAt(td.xr, td.pos, td.r)
		if len(td.expect) != len(result) {
			t.Errorf("#%v: expected %v (len %v), got %v (len %v)", i, td.expect, len(td.expect), result, len(result))
			continue
		}
		for j := 0; j < len(result); j++ {
			if td.expect[j] != result[j] {
				t.Errorf("#%v: expected %v (len %v), got %v (len %v)", i, td.expect, len(td.expect), result, len(result))
				break
			}
		}
	}
}
