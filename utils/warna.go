package utils

import "strings"

func AmabilWarna(jurusan string) string {
	jurusan = strings.ToUpper(jurusan)

	switch {
	case strings.Contains(jurusan, "PPLG"):
		return "bg-green-800"
	case strings.Contains(jurusan, "TJKT"):
		return "bg-blue-700"
	case strings.Contains(jurusan, "TBSM"):
		return "bg-red-800"
	case strings.Contains(jurusan, "DKV"):
		return "bg-yellow-500"
	case strings.Contains(jurusan, "TOI"):
		return "bg-black"
	default:
		return "bg-gray-400"
	}
}