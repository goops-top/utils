/*
Copyright 2020 BGBiao Ltd. All rights reserved.
@File   : elasticsearch_test.go
@Time   : 2021/04/06 15:26:28
@Update : 2021/04/06 15:26:28
@Author : BGBiao
@Version: 1.0
@Contact: weichaungxxb@qq.com
@Desc   : None
*/
package v6

import (
	"fmt"
	"testing"
)

func TestElasticSearch(t *testing.T) {
	// clientMetadata := NewClientInfo([]string{"http://10.0.202.140:9200", "http://10.0.202.134:9200"})
	clientMetadata := NewClientInfo([]string{"http://10.0.32.222:9200", "http://10.0.34.183:9200"})
	clientMetadata.SetClientAuth("username", "password")

	_, err := clientMetadata.NewClient()
	if err != nil {
		panic(err)
	}

	indexNames, err := clientMetadata.GetIndexNames()

	if err != nil {
		panic(err)
	}

	for _, index := range indexNames {
		fmt.Println(index)
	}

}
