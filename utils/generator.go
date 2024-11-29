package utils

import "github.com/speps/go-hashids/v2"

type GeneratorHash struct {
	Hash *hashids.HashID
}

func InitHash(salt string, length int) HashInterface {
	if length == 0 {
		length = 10
	}

	if salt == "" {
		salt = "bbd24f22-f4d8-4af4-b2e4-4eb9852e5351"
	}
	hd := hashids.NewData()
	// let's make Salt and MinLength as constant do not change
	hd.Salt = salt
	hd.MinLength = length
	h, _ := hashids.NewWithData(hd)
	return GeneratorHash{
		Hash: h,
	}
}

type HashInterface interface {
	EncodePublicID(id int64) string
	DecodePublicID(encodedStr string) int64
}

func (g GeneratorHash) EncodePublicID(id int64) string {
	e, _ := g.Hash.EncodeInt64([]int64{id})
	return e
}

func (g GeneratorHash) DecodePublicID(encodedStr string) int64 {
	d, _ := g.Hash.DecodeInt64WithError(encodedStr)
	return d[0]
}
