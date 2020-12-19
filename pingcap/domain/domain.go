package domain

type Codec interface {
	Encode() []byte
	Decode([]byte)
}
