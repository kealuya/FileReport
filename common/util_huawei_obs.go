package common

//
//func putObject() {
//	input := &obs.PutObjectInput{}
//	input.Bucket = bucketName
//	input.Key = objectKey
//	input.Metadata = map[string]string{"meta": "value"}
//	input.Body = strings.NewReader("Hello OBS")
//	output, err := getObsClient().PutObject(input)
//	if err == nil {
//		fmt.Printf("StatusCode:%d, RequestId:%s\n", output.StatusCode, output.RequestId)
//		fmt.Printf("ETag:%s, StorageClass:%s\n", output.ETag, output.StorageClass)
//	} else {
//		if obsError, ok := err.(obs.ObsError); ok {
//			fmt.Println(obsError.StatusCode)
//			fmt.Println(obsError.Code)
//			fmt.Println(obsError.Message)
//		} else {
//			fmt.Println(err)
//		}
//	}
//}
