package main

func encryptIt(input, key string) string {
	var byteInput = []byte(input)
	var byteKey = []byte(key)
	var output = []byte("")

	for i := 0; i < len(byteInput); i++ {

		output = append(output, (byteInput[i] ^ byteKey[i%len(byteKey)]%26))
	}

	return string(output)
}

func decryptIt(input, input2 string) string {
	var byteInput = []byte(input)
	var byteInput2 = []byte(input2)
	var output = []byte("")

	for i := 0; i < len(byteInput); i++ {

		output = append(output, (byteInput[i] ^ (byteInput2[i] % 26)))

	}

	return string(output)
}
