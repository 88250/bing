// Copyright (c) 2017, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func main() {
	bucket := flag.String("bucket", "", "bucket name")
	ak := flag.String("ak", "", "access key")
	sk := flag.String("sk", "", "secret key")
	flag.Parse()
	if "" == *bucket {
		log.Fatal("please specify bucket with -bucket")
	}
	if "" == *ak {
		log.Fatal("please specify access key with -ak")
	}
	if "" == *sk {
		log.Fatal("please specify secret key with -sk")
	}

	picURL := todayPicURL()
	log.Printf("today's pic URL is [%s]", picURL)
	data := picData(picURL)
	log.Printf("pic length is [%dK]", len(data)/1024)

	key := time.Now().Format("bing/20060102.jpg")
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", *bucket, key), // overwrite if exists
	}
	formUploader := storage.NewFormUploader(nil)
	if err := formUploader.Put(context.Background(), nil, putPolicy.UploadToken(qbox.NewMac(*ak, *sk)),
		key, bytes.NewReader(data), int64(len(data)), nil); nil != err {
		log.Fatal(err)
	}

	log.Println("process sucessfully, exit")
}

func picData(picURL string) (data []byte) {
	var errs []error
	if _, data, errs = gorequest.New().Get(picURL).Timeout(15*time.Second).
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndBytes(); nil != errs {
		log.Fatalf("%s", errs[0])
	}
	return
}

func todayPicURL() string {
	data := map[string]interface{}{}
	if _, _, errs := gorequest.New().Get("https://cn.bing.com/HPImageArchive.aspx?format=js&n=1").
		Timeout(15*time.Second).
		Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError).
		EndStruct(&data); nil != errs {
		log.Fatalf("%s", errs[0])
	}

	images := data["images"].([]interface{})
	image := images[0].(map[string]interface{})

	return "https://www.bing.com" + image["url"].(string)
}
