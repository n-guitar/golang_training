// どうもうまくいかない・・・

package main

import (
	"bytes"
	"testing"
)

func TestCharcount(t *testing.T) {
	tests := []struct {
		bytes   []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		// TODO: Add test cases.
		{
			[]byte("楽しい楽しいGo言語"),
			map[rune]int{'楽': 2, 'し': 2, 'い': 2, 'G': 1, 'o': 1, '言': 2, '語': 2},
			[]int{2, 0, 0, 8, 0},
			0},
		{
			[]byte("難しい難しい研修"),
			map[rune]int{'難': 2, 'し': 2, 'い': 2, '研': 1, '修': 1},
			[]int{0, 0, 0, 8, 0},
			0},
	}
	for _, tt := range tests {
		counts, utflen, invalid, err := charcount(bytes.NewReader(tt.bytes))

		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}

		if len(counts) != len(tt.counts) {
			t.Errorf("err_counts %d not %d", len(counts), len(tt.counts))
			continue
		}

		for s, t_count := range tt.counts {
			count := counts[s]

			if count != t_count {
				t.Errorf("err_count %d not %d", count, t_count)
				continue
			}
		}

		if len(utflen) != len(tt.utflen) {
			t.Errorf("err_utflen %d not %d", len(utflen), len(tt.utflen))
			continue
		}

		if invalid != tt.invalid {
			t.Errorf("err_invalid %d not %d", invalid, tt.invalid)
		}
	}
}
