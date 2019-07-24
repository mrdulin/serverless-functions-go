package utils

import (
	"reflect"
	"testing"
)

type googleAccount struct {
	ClientCustomerId int
	RefreshToken     string
}

func TestUnique(t *testing.T) {
	type args struct {
		googleAccounts []googleAccount
	}

	tests := []struct {
		name string
		args args
		want []googleAccount
	}{
		{
			name: "should remove duplicated element from struct slice correctly",
			args: args{
				[]googleAccount{{1, "a"}, {2, "b"}, {1, "a"}},
			},
			want: []googleAccount{{1, "a"}, {2, "b"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var interfaceSlice = make([]interface{}, len(tt.args.googleAccounts))
			for i, d := range tt.args.googleAccounts {
				interfaceSlice[i] = d
			}

			got := Unique(interfaceSlice)

			var gotGoogleAccount = make([]googleAccount, len(got))
			for j, val := range got {
				gotGoogleAccount[j] = val.(googleAccount)
			}
			if !reflect.DeepEqual(gotGoogleAccount, tt.want) {
				t.Errorf("unique got = %#v, want = %#v", gotGoogleAccount, tt.want)
			}
		})
	}
}

func TestUnique_v2(t *testing.T) {

	type args struct {
		googleAccounts []googleAccount
	}

	tests := []struct {
		name string
		args args
		want []googleAccount
	}{
		{
			name: "should remove duplicated struct from slice correctly",
			args: args{
				googleAccounts: []googleAccount{{1, "a"}, {2, "b"}, {1, "a"}},
			},
			want: []googleAccount{{1, "a"}, {2, "b"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UniqueV2(tt.args.googleAccounts)
			gotGoogleAccounts := make([]googleAccount, len(got))
			for i, val := range got {
				gotGoogleAccounts[i] = val.(googleAccount)
			}
			if !reflect.DeepEqual(gotGoogleAccounts, tt.want) {
				t.Errorf("unique v2 got = %#v, want = %#v", gotGoogleAccounts, tt.want)
			}
		})
	}
}
