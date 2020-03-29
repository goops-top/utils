/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:mail_test.go
*Author:BGBiao
*Date:2020年03月29日
*Description:
*
================================================================*/
package mail

import (
	"fmt"
	"testing"
)

func TestPostMail(t *testing.T) {
	maildata := &EmailMetaData{
    Smtp: "smtp.qq.com",
		From: "goops.top@qq.com",
		Pass: "password",
		To: []string{"goops.top@qq.com"},
		Cc: []string{"weichuangxxb@qq.com"},
		Subject: "test",
		ContentType: "text",
		Content: []byte("hahahah"),
	}
	emailErr := maildata.PostEmail()
	fmt.Println(emailErr)
}

