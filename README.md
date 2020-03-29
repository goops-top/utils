# Golang项目中可能常用到的工具库




**邮件发送**

```
# install the package
$ go get -v github.com/goops-top/utils/mail

# update the mail data
$ cat test-mail.go

package main

import (
	"fmt"

	"github.com/goops-top/utils/mail"
)

func main() {
	maildata := &mail.EmailMetaData{
		Smtp:        "smtp.qq.com",
		From:        "goops@qq.com",
		Pass:        "passwoed",
		To:          []string{"goops@qq.com"},
		Cc:          []string{"weichuangxxb@qq.com"},
		Subject:     "test",
		ContentType: "text",
		Content:     []byte("hahahah"),
	}
	emailErr := maildata.PostEmail()
	if emailErr != nil {
		fmt.Printf("邮件发送失败:%v\n", emailErr)
	} else {
		fmt.Println("邮件已发送")
	}

}


# test mail
$ go run test-mail.go
邮件已发送

```
