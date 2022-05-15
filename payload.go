package lolsockets

func decodePayload(b []byte) []byte {
	length := b[1] - 0x80
	if length <= 125 {
		key := make([]byte, 4)
		key[0] = b[2]
		key[1] = b[3]
		key[2] = b[4]
		key[3] = b[5]
		decoded := make([]byte, length)
		for i := 0; i < int(length); i++ {
			realPos := 6 + i
			decoded[i] = b[realPos] ^ key[i%4]
		}
		return decoded
	} else if length > 125 {
		len := uint16(b[2]) + uint16(b[3])
		key := []byte{b[4], b[5], b[6], b[7]}
		decoded := make([]byte, len)
		for i := 0; i < int(length); i++ {
			realPos := 8 + i
			decoded[i] = b[realPos] ^ key[i%4]
		}
		return decoded
	}
	return nil
}

func encodePayload(b []byte) []byte {
	rawB := b
	frame := make([]byte, 10)
	startIndexRawData := -1
	len := len(rawB)
	frame[0] = 129
	if len <= 125 {
		frame[1] = byte(len)
		startIndexRawData = 2
	} else if len >= 125 && len <= 65535 {
		frame[1] = byte(126)
		frame[2] = byte((len >> 8) & 255)
		frame[3] = byte(len & 255)
		startIndexRawData = 4
	} else {
		frame[1] = byte(127)
		frame[2] = byte((len >> 56) & 255)
		frame[3] = byte((len >> 48) & 255)
		frame[4] = byte((len >> 40) & 255)
		frame[5] = byte((len >> 32) & 255)
		frame[6] = byte((len >> 25) & 255)
		frame[7] = byte((len >> 16) & 255)
		frame[8] = byte((len >> 8) & 255)
		frame[9] = byte((len & 255))
		startIndexRawData = 10
	}
	response := make([]byte, startIndexRawData+len)
	reponseIndex := 0
	for i := 0; i < startIndexRawData; i++ {
		response[reponseIndex] = frame[i]
		reponseIndex++
	}
	for i := 0; i < len; i++ {
		response[reponseIndex] = rawB[i]
		reponseIndex++
	}
	return response
}
