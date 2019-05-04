package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type minio_obj struct {
	service, token, bucketname, filename, function string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var minioobj minio_obj
	params := strings.Split(r.RequestURI, "/")
	for val, ind := range params {
		log.Println(string(val))
	}
	minioobj.service, minioobj.function, minioobj.bucketname, minioobj.filename = params[1], params[2], params[3], params[4]
	fmt.Fprintf(w, "%s", minioobj)
}

func main() {
	// ############################# Server part ###############################################################
	port := ":9999"
	http.HandleFunc("/", handler)
	log.Printf("Server started at port%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
	// ##########################################################################################################
}

//############################## Get credentials #############################################################
// nossl := false
// client, err := minio.New("127.0.0.1:1111", "V2XOJRNTZFZOKZ80LP8N", "DANIriky5Z08LM8gfTrVKVnD1feSAZpPBUufA1le", nossl)
// if err != nil {
// 	fmt.Println("Can't connect to Minio server")
// }
//##########################################################################################################

//############################## Downloading object ######################################################
// object, err := client.GetObject("simple", "stats.json", minio.GetObjectOptions{})
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// localfile, err := os.Create("./stat_minio.json")
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// if _, err = io.Copy(localfile, object); err != nil {
// 	fmt.Println(err)
// 	return
// }
//#########################################################################################################

//############################# List of buckets ###########################################################
// buckets, err := client.ListBuckets()
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// for _, bucket := range buckets {
// 	fmt.Println(bucket.Name)
// }
//#########################################################################################################
