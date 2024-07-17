package vhash

import (
	"crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	assert.Equal(t, "17d3f69a94e615e55221ca0948c0c0bb", MD5("vazhenina"))
}

func TestSHA1(t *testing.T) {
	assert.Equal(t, "bec5d20a9632bf0a57bd3ef45181bd5fca37ea0e", SHA1("vazhenina"))
}

func TestSHA256(t *testing.T) {
	assert.Equal(t, "b912d156a6fbc6e6922c41b3d926eeb09991e1195f84989baa98de62b3a5a83e", SHA256("vazhenina"))
}

func TestHash(t *testing.T) {
	type args struct {
		hash crypto.Hash
		s    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "md5",
			args: args{hash: crypto.MD5, s: "vazhenina"},
			want: "17d3f69a94e615e55221ca0948c0c0bb",
		},
		{
			name: "sha1",
			args: args{hash: crypto.SHA1, s: "vazhenina"},
			want: "bec5d20a9632bf0a57bd3ef45181bd5fca37ea0e",
		},
		{
			name: "sha224",
			args: args{hash: crypto.SHA224, s: "vazhenina"},
			want: "42a1b6dfe2d93e82719bc8687152cc19e5260a7100cdc71ed3265b8f",
		},
		{
			name: "sha256",
			args: args{hash: crypto.SHA256, s: "vazhenina"},
			want: "b912d156a6fbc6e6922c41b3d926eeb09991e1195f84989baa98de62b3a5a83e",
		},
		{
			name: "sha384",
			args: args{hash: crypto.SHA384, s: "vazhenina"},
			want: "d0714d3621756d58d44e8bfcb11606bd856361d6af418527aba82db9538953695590413ed8d887a49fd3c7dd71521e26",
		},
		{
			name: "sha512",
			args: args{hash: crypto.SHA512, s: "vazhenina"},
			want: "eae2861fbc441c23176d7185a67a777449d9b80ce03a16feca7c1e499405c00d457a11bbcf1c2cedf22729786460a552a0ad24ba529d63b9f1dc6fcd503814b1",
		},
	}
	for _, tt := range tests {
		v, err := Hash(tt.args.hash, tt.args.s)

		assert.Nil(t, err)
		assert.Equal(t, tt.want, v)
	}
}
