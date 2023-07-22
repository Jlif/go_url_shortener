package util

import "github.com/spaolacci/murmur3"

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ShortUrl32(original string) string {
	return murmurHash32ToBase62(getMurmurHash32(original))
}

func getMurmurHash32(originalUrl string) uint32 {
	if originalUrl == "" {
		return 0
	}
	hash := murmur3.Sum32([]byte(originalUrl))
	return hash
}

func murmurHash32ToBase62(hash uint32) string {
	base := uint32(len(base62Chars))
	var result string

	for hash > 0 {
		remainder := hash % base
		result = string(base62Chars[remainder]) + result
		hash /= base
	}

	return result
}

func ShortUrl64(original string) string {
	return murmurHash64ToBase62(getMurmurHash64(original))
}

func getMurmurHash64(originalUrl string) uint64 {
	if originalUrl == "" {
		return 0
	}
	hash := murmur3.Sum64([]byte(originalUrl))
	return hash
}

func murmurHash64ToBase62(hash uint64) string {
	base := uint64(len(base62Chars))
	var result string

	for hash > 0 {
		remainder := hash % base
		result = string(base62Chars[remainder]) + result
		hash /= base
	}

	return result
}
