/*================================================================
*Copyright (C) 2020 BGBiao Ltd. All rights reserved.
*
*FileName:mail.go
*Author:Xuebiao Xu
*Date:2020年03月29日
*Description:
*
================================================================*/
package mail
import (
	"net/smtp"

	"github.com/jordan-wright/email"
)


type EmailMetaData struct {
  Smtp    string    // smtp邮件服务器地址
	From		string		// 发送方
	Pass		string		// 发送密码
	To		[]string	// 接受方
	Cc		[]string	// 抄送
	Bcc		[]string	// 秘密抄送

	Subject		string		// 邮件主题
	Content		[]byte		// 邮件内容(可支持e.Text和e.HTML)
	ContentType	string		// 邮件内容形式:text|html
}

func (emData *EmailMetaData) PostEmail() (error) {
	emailObj := email.NewEmail()
        emailObj.From = emData.From
        emailObj.To = emData.To
        emailObj.Cc = emData.Cc
        emailObj.Bcc = emData.Bcc

        emailObj.Subject = emData.Subject

	if emData.ContentType == "html" {
		emailObj.HTML = emData.Content
	}else {
		emailObj.Text = emData.Content
	}

	err := emailObj.Send(emData.Smtp+":25", smtp.PlainAuth("", emData.From, emData.Pass, emData.Smtp))
        return err
}


