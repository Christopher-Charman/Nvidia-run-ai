package gojsonq

import "testing"

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
