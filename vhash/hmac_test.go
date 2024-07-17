package vhash

import (
	"crypto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHMacSHA1(t *testing.T) {
	assert.Equal(t, "c65e87811bdd32e8cdb738819a193bf606e40994", HMacSHA1("vazhenina", "ILoveYou"))
}

func TestHMacSHA256(t *testing.T) {
	assert.Equal(t, "020bf262934d06f3bceefc544b0dcd341abbf0cba13e93b6a281e96bb750c22b", HMacSHA256("vazhenina", "ILoveYou"))
}

func TestHMac(t *testing.T) {
	type args struct {
		hash crypto.Hash
		key  string
		s    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "md5",
			args: args{hash: crypto.MD5, key: "vazhenina", s: "ILoveYou"},
			want: "6781269dfb1c216e86563d3357cb1c3d",
		},
		{
			name: "sha1",
			args: args{hash: crypto.SHA1, key: "vazhenina", s: "ILoveYou"},
			want: "c65e87811bdd32e8cdb738819a193bf606e40994",
		},
		{
			name: "sha224",
			args: args{hash: crypto.SHA224, key: "vazhenina", s: "ILoveYou"},
			want: "8ed7b8a323146a28286aa24eb3ba494d605833a186a7eecf1b327bd8",
		},
		{
			name: "sha256",
			args: args{hash: crypto.SHA256, key: "vazhenina", s: "ILoveYou"},
			want: "020bf262934d06f3bceefc544b0dcd341abbf0cba13e93b6a281e96bb750c22b",
		},
		{
			name: "sha384",
			args: args{hash: crypto.SHA384, key: "vazhenina", s: "ILoveYou"},
			want: "c635f82ee90fd0e1f0fe36a4b668f65e57775eb5dd5a88c28f4e83a5d6e9ca20c258259603e34081067fe1a6782dfa2d",
		},
		{
			name: "sha512",
			args: args{hash: crypto.SHA512, key: "vazhenina", s: "ILoveYou"},
			want: "f2f6578415fee4cbfa0a5c495f60f45af1a673efa7a0d4581aa671de2cd6995f7c671442607bb3f97b3a7952eb936522851a29ebeeaef1512d7f34177df9bb12",
		},
	}
	for _, tt := range tests {
		v, err := HMac(tt.args.hash, tt.args.key, tt.args.s)

		assert.Nil(t, err)
		assert.Equal(t, tt.want, v)
	}
}
