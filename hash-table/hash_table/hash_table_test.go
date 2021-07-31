package hash_table

import (
	"testing"
)

func TestHashTable_Add(t *testing.T) {
	type fields struct {
		hashValue [10]HashValue
		size      int
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
	}{
		{
			"add one value",
			fields{hashValue: [10]HashValue{}, size: 10},
			[]args{{"TEST", 6}},
		}, {
			"add two value",
			fields{hashValue: [10]HashValue{}, size: 10},
			[]args{
				{"TEST", 6},
				{"TEST2", "99999"},
			},
		}, {
			"duplicate key",
			fields{hashValue: [10]HashValue{}, size: 10},
			[]args{
				{"TEST", 6},
				{"TEST", "99999"},
			},
		}, {
			"add nine values",
			fields{hashValue: [10]HashValue{}, size: 10},
			[]args{
				{"TEST", 6},
				{"TEST2", "99999"},
				{"TEST3", 6},
				{"TEST4", "99999"},
				{"TEST5", 6},
				{"TEST6", "99999"},
				{"TEST7", 6},
				{"TEST8", "99999"},
				{"TEST9", 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				hashValue: tt.fields.hashValue,
				size:      tt.fields.size,
			}
			for _, arg := range tt.args {
				h.Add(arg.key, arg.value)
				result := h.Get(arg.key)
				if result != arg.value {
					t.Errorf("%v was incorrect, key: %s , got: %v, expect: %v.", tt.name, arg.key, result, arg.value)
				}
			}

		})
	}
}

func TestHashTable_Get(t *testing.T) {
	type fields struct {
		hashValue [10]HashValue
		size      int
	}
	type args struct {
		key     string
		value   interface{}
		testKey string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"Get by key GET",
			fields{
				hashValue: [10]HashValue{}, size: 10},
			args{"GET", "GET THIS", "GET"},
		},
		{
			"Get by key GET expected panic",
			fields{
				hashValue: [10]HashValue{}, size: 10},
			args{"GET", "GET THIS", "AB001"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				hashValue: tt.fields.hashValue,
				size:      tt.fields.size,
			}
			h.Add(tt.args.testKey, tt.args.value)
			result := h.Get(tt.args.key)
			if result != tt.args.value {
				t.Errorf("Value of hash value was incorrect, got: %v, want: %v.", result, tt.args.value)
			}
		})
	}
}

func TestHashTable_Exist(t *testing.T) {
	type fields struct {
		hashValue [10]HashValue
		size      int
	}
	type args struct {
		key     string
		value   interface{}
		testKey string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedResult interface{}
	}{
		{
			"Check key TEST1 exist or not expect true",
			fields{
				hashValue: [10]HashValue{}, size: 10},
			args{"TEST1", "1", "TEST1"},
			true,
		},
		{
			"Check key TEST1 exist or not expect false",
			fields{
				hashValue: [10]HashValue{}, size: 10},
			args{"TEST1", "1", "TEST10"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				hashValue: tt.fields.hashValue,
				size:      tt.fields.size,
			}
			h.Add(tt.args.testKey, tt.args.value)
			result := h.Exist(tt.args.key)
			if result != tt.expectedResult {
				t.Errorf("Value of hash value was incorrect, got: %v, want: %v.", result, tt.expectedResult)
			}
		})
	}
}

func TestHashTable_Remove(t *testing.T) {
	type fields struct {
		hashValue [10]HashValue
		size      int
	}
	type args struct {
		key     string
		value   interface{}
		testKey string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		expectedResult interface{}
	}{
		{
			"Remove key TEST1 expect false",
			fields{
				hashValue: [10]HashValue{
					{"TEST", "6"},
					{"TEST2", "99999"},
					{"TEST3", "6"},
					{"TEST4", "99999"},
					{"TEST5", "6"},
					{"TEST6", "99999"},
					{"TEST7", "6"},
					{"TEST8", "99999"},
					{"TEST9", "6"},
				}, size: 10,
			},
			args{"TEST1", "1", "TEST1"},
			false,
		},
		{
			"Remove key TEST10 expect true",
			fields{
				hashValue: [10]HashValue{}, size: 10},
			args{"TEST1", "1", "TEST10"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HashTable{
				hashValue: tt.fields.hashValue,
				size:      tt.fields.size,
			}
			h.Add(tt.args.key, tt.args.value)
			h.Remove(tt.args.testKey)
			result := h.Exist(tt.args.key)
			if result != tt.expectedResult {
				t.Errorf("Value of hash value was incorrect, got: %v, want: %v.", result, tt.expectedResult)
			}
		})
	}
}
