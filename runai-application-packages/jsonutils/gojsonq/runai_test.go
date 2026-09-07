package gojsonq

import (
	"encoding/json"
	"testing"
)

// TestCleanMap covers the CleanMap behaviors beyond the basic nil/nested cases:
// falsy scalars must be kept, empty maps/slices are pruned (and emptiness
// cascades upward), and slice elements that are nil or clean down to empty are
// dropped.
func TestCleanMap(t *testing.T) {
	cases := []struct {
		name     string
		src      map[string]interface{}
		expected map[string]interface{}
	}{
		{
			// top-level nil scalar is dropped, real scalar kept.
			name: "scalar",
			src: map[string]interface{}{
				"image": "ubuntu",
				"tag":   nil,
			},
			expected: map[string]interface{}{
				"image": "ubuntu",
			},
		},
		{
			// nils are pruned at every depth: inside a map nested in a slice.
			name: "deep map",
			src: map[string]interface{}{
				"storage": map[string]interface{}{
					"git": []interface{}{
						map[string]interface{}{
							"repository": "https://github.com/go-jsonq/go-jsonq.git",
							"branch":     nil,
							"secret": map[string]interface{}{
								"token":    "secret",
								"password": nil,
							},
						},
					},
				},
			},
			expected: map[string]interface{}{
				"storage": map[string]interface{}{
					"git": []interface{}{
						map[string]interface{}{
							"repository": "https://github.com/go-jsonq/go-jsonq.git",
							"secret": map[string]interface{}{
								"token": "secret",
							},
						},
					},
				},
			},
		},
		{
			// nil is removed, but falsy-but-real values (false, 0, "") are kept.
			name: "keeps falsy scalars, drops nil",
			src: map[string]interface{}{
				"emptyStr": "",
				"zero":     0,
				"false":    false,
				"nilVal":   nil,
			},
			expected: map[string]interface{}{
				"emptyStr": "",
				"zero":     0,
				"false":    false,
			},
		},
		{
			// an already-empty map value is dropped, siblings survive.
			name: "drops empty nested map",
			src: map[string]interface{}{
				"env":             map[string]interface{}{},
				"imagePullPolicy": "Always",
			},
			expected: map[string]interface{}{
				"imagePullPolicy": "Always",
			},
		},
		{
			// a map that only holds nils becomes empty and is removed; emptiness
			// then cascades all the way up, leaving an empty top-level map.
			name: "emptiness cascades to top",
			src: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"c": nil,
					},
				},
			},
			expected: map[string]interface{}{},
		},
		{
			// scalar slices are preserved as-is.
			name: "keeps scalar slice",
			src: map[string]interface{}{
				"args": []interface{}{"a", "b", "c"},
			},
			expected: map[string]interface{}{
				"args": []interface{}{"a", "b", "c"},
			},
		},
		{
			// nil elements are dropped from a slice, order of the rest preserved.
			name: "drops nil slice elements",
			src: map[string]interface{}{
				"args": []interface{}{"a", nil, "b"},
			},
			expected: map[string]interface{}{
				"args": []interface{}{"a", "b"},
			},
		},
		{
			// an empty slice is dropped entirely.
			name: "drops empty slice",
			src: map[string]interface{}{
				"args": []interface{}{},
				"keep": "x",
			},
			expected: map[string]interface{}{
				"keep": "x",
			},
		},
		{
			// a slice whose only elements clean down to empty maps is removed.
			name: "drops slice of emptied maps",
			src: map[string]interface{}{
				"git": []interface{}{
					map[string]interface{}{},
					map[string]interface{}{"branch": nil},
				},
			},
			expected: map[string]interface{}{},
		},
		{
			// mixed slice: keep scalars and non-empty maps, drop nils and emptied maps.
			name: "prunes mixed slice",
			src: map[string]interface{}{
				"items": []interface{}{
					"keep",
					map[string]interface{}{"token": "secret", "password": nil},
					map[string]interface{}{},
					nil,
				},
			},
			expected: map[string]interface{}{
				"items": []interface{}{
					"keep",
					map[string]interface{}{"token": "secret"},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			CleanMap(tc.src)
			assertInterface(t, tc.expected, tc.src, tc.name)
		})
	}
}

// assertJSONNumber fails unless v is a json.Number whose literal equals want.
// It checks the concrete type on purpose: assertInterface compares via
// json.Marshal, which renders json.Number("1") and float64(1) identically and
// therefore cannot verify the conversion NormalizeNumbers performs.
func assertJSONNumber(t *testing.T, v interface{}, want string) {
	t.Helper()
	n, ok := v.(json.Number)
	if !ok {
		t.Errorf("expected json.Number, got %T (%v)", v, v)
		return
	}
	if n.String() != want {
		t.Errorf("expected json.Number(%q), got %q", want, n.String())
	}
}

// TestNormalizeNumbers verifies that every numeric kind is rewritten as a
// json.Number (matching decoder.UseNumber()), at every depth, while
// non-numbers and existing json.Numbers are left untouched.
func TestNormalizeNumbers(t *testing.T) {
	src := map[string]interface{}{
		"float":    float64(0.5),
		"wholeFlt": float64(8),
		"float32":  float32(1.5),
		"int":      int(3),
		"int64":    int64(4),
		"uint":     uint(5),
		"already":  json.Number("2"),
		"str":      "x",
		"bool":     true,
		"nilVal":   nil,
		"nested":   map[string]interface{}{"n": float64(1)},
		"slice":    []interface{}{float64(2), "y", map[string]interface{}{"m": int(3)}},
	}

	out, ok := NormalizeNumbers(src).(map[string]interface{})
	if !ok {
		t.Fatalf("expected map result, got %T", NormalizeNumbers(src))
	}

	assertJSONNumber(t, out["float"], "0.5")
	assertJSONNumber(t, out["wholeFlt"], "8")
	assertJSONNumber(t, out["float32"], "1.5")
	assertJSONNumber(t, out["int"], "3")
	assertJSONNumber(t, out["int64"], "4")
	assertJSONNumber(t, out["uint"], "5")
	assertJSONNumber(t, out["already"], "2")

	// non-numbers are left as-is.
	if out["str"] != "x" {
		t.Errorf("expected string %q, got %v", "x", out["str"])
	}
	if out["bool"] != true {
		t.Errorf("expected bool true, got %v", out["bool"])
	}
	if out["nilVal"] != nil {
		t.Errorf("expected nil, got %v", out["nilVal"])
	}

	// numbers are converted at every depth.
	nested := out["nested"].(map[string]interface{})
	assertJSONNumber(t, nested["n"], "1")

	slice := out["slice"].([]interface{})
	assertJSONNumber(t, slice[0], "2")
	if slice[1] != "y" {
		t.Errorf("expected string %q, got %v", "y", slice[1])
	}
	assertJSONNumber(t, slice[2].(map[string]interface{})["m"], "3")

	// scalar input is converted and returned.
	assertJSONNumber(t, NormalizeNumbers(float64(1.25)), "1.25")

	// idempotent: a second pass keeps the json.Number.
	assertJSONNumber(t, NormalizeNumbers(NormalizeNumbers(float64(7))), "7")
}
