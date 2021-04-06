/*
Copyright 2020 BGBiao Ltd. All rights reserved.
@File   : elasticsearch.go
@Time   : 2021/04/06 14:50:55
@Update : 2021/04/06 14:50:55
@Author : BGBiao
@Version: 1.0
@Contact: weichaungxxb@qq.com
@Desc   : None
*/
package v6

import (
	"fmt"

	"github.com/olivere/elastic/v6"
)

// elastic 中的参数均采用 function options 的方式传入
// type ClientOptionFunc func(*Client) error
// https://pkg.go.dev/github.com/olivere/elastic/v6#ClientOptionFunc

type clientInfo struct {
	URLs              []string        `json:"urls"`
	IsBasicAuth       bool            `json:"isBasicAuth"`
	BasicAuthUser     string          `json:"basicAuthUser"`
	BasicAuthPassword string          `json:"basicAuthPassword"`
	Gzip              bool            `json:"gzip"`
	Client            *elastic.Client `json:"client"`
}

func NewClientInfo(urls []string) *clientInfo {
	return &clientInfo{
		URLs:        urls,
		IsBasicAuth: false,
	}
}

// enable the basic auth with the username and password
func (c *clientInfo) SetClientAuth(username, password string) {
	c.IsBasicAuth = true
	c.BasicAuthUser = username
	c.BasicAuthPassword = password

}

func (c *clientInfo) NewClient() (client *elastic.Client, err error) {
	// urls := strings.Join(c.URLs, " ")
	// notice: elastic.SetURl(args ...string)
	var url1, url2 string
	if len(c.URLs) > 1 {
		url1 = c.URLs[0]
		url2 = c.URLs[1]
	} else {
		url1, url2 = c.URLs[0], c.URLs[0]
	}
	urls := c.URLs[0]

	fmt.Println(urls, c.IsBasicAuth, c.Gzip)
	if c.IsBasicAuth {
		client, err = elastic.NewClient(elastic.SetURL(url1, url2), elastic.SetBasicAuth(c.BasicAuthUser, c.BasicAuthPassword))
	} else {
		client, err = elastic.NewClient(elastic.SetURL(url1, url2), elastic.SetGzip(c.Gzip))
	}

	fmt.Println(client, err)
	if err == nil {
		c.Client = client
	}
	return client, err
}

// GetINdexNames wrap to the  IndexNames that  returns the names of all indices in the cluster.
func (c *clientInfo) GetIndexNames() ([]string, error) {
	return c.Client.IndexNames()
}
