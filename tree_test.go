package radixtree

import (
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	tests := []struct {
		name   string
		labels []string
		want   *Tree
	}{
		{
			name: "tree",
			labels: []string{
				"PLAN",
				"PLAY",
				"POLL",
				"POST",
			},
			want: &Tree{
				label: "",
				edges: []*Tree{
					&Tree{
						label: "P",
						edges: []*Tree{
							&Tree{
								label: "LA",
								edges: []*Tree{
									&Tree{
										label: "N",
									},
									&Tree{
										label: "Y",
									},
								},
							},
							&Tree{
								label: "O",
								edges: []*Tree{
									&Tree{
										label: "LL",
									},
									&Tree{
										label: "ST",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "tree",
			labels: []string{
				"romane",
				"romanus",
				"romulus",
				"rubens",
				"ruber",
				"rubicon",
				"rubicundus",
			},
			want: &Tree{
				label: "",
				edges: []*Tree{
					&Tree{
						label: "r",
						edges: []*Tree{
							&Tree{
								label: "om",
								edges: []*Tree{
									&Tree{
										label: "an",
										edges: []*Tree{
											&Tree{
												label: "e",
											},
											&Tree{
												label: "us",
											},
										},
									},
									&Tree{
										label: "ulus",
									},
								},
							},
							&Tree{
								label: "ub",
								edges: []*Tree{
									&Tree{
										label: "e",
										edges: []*Tree{
											&Tree{
												label: "ns",
											},
											&Tree{
												label: "r",
											},
										},
									},
									&Tree{
										label: "ic",
										edges: []*Tree{
											&Tree{
												label: "on",
											},
											&Tree{
												label: "undus",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRadixTree()
			for _, label := range tt.labels {
				got.Insert(label)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Lookup(t *testing.T) {
	tests := []struct {
		name   string
		labels []string
	}{
		{
			name: "tree",
			labels: []string{
				"PLAN",
				"PLAY",
				"POLL",
				"POST",
			},
		},
		{
			name: "tree",
			labels: []string{
				"romane",
				"romanus",
				"romulus",
				"rubens",
				"ruber",
				"rubicon",
				"rubicundus",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewRadixTree()
			for _, label := range tt.labels {
				tree.Insert(label)
			}
			for _, label := range tt.labels {
				got := tree.Lookup(label)
				if !got {
					t.Errorf("Tree.Lookup(%v) = %v, want %v", label, got, true)
				}
			}

		})
	}
}

func TestTree_Delete(t *testing.T) {

	tests := []struct {
		name   string
		labels []string
		delete string
		want   bool
	}{
		{
			name: "tree",
			labels: []string{
				"PLAN",
				"PLAY",
				"POLL",
				"POST",
			},
			delete: "PLAN",
			want:   true,
		},
		{
			name: "tree",
			labels: []string{
				"romane",
				"romanus",
				"romulus",
				"rubens",
				"ruber",
				"rubicon",
				"rubicundus",
			},
			delete: "rub",
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewRadixTree()
			for _, label := range tt.labels {
				tree.Insert(label)
			}

			got := tree.Delete(tt.delete)
			if got != tt.want {
				t.Errorf("Tree.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_HasPrefix(t *testing.T) {

	tests := []struct {
		name   string
		labels []string
		prefix string
		want   []string
	}{
		{
			name: "tree",
			labels: []string{
				"PLAN",
				"PLAY",
				"POLL",
				"POST",
			},
			prefix: "PLAN",
			want:   []string{"PLAN"},
		},
		{
			name: "tree",
			labels: []string{
				"romane",
				"romanus",
				"romulus",
				"rubens",
				"ruber",
				"rubicon",
				"rubicundus",
			},
			prefix: "rub",
			want: []string{
				"rubens",
				"ruber",
				"rubicon",
				"rubicundus",
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewRadixTree()
			for _, label := range tt.labels {
				tree.Insert(label)
			}

			got := tree.HasPrefix(tt.prefix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.HasPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
