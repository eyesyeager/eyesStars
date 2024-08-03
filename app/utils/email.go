package utils

import (
	"errors"
	"eyesStars/global"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

/**
 * @author eyesYeager
 * @date 2024/7/25 21:51
 */

type emailUtils struct{}

var EmailUtils = emailUtils{}

// SendEmailWithText 发送邮件(Text格式)
func (emailUtils *emailUtils) SendEmailWithText(subject string, to []string, text string) error {
	return emailUtils.SendEmailWithTextAndCC(subject, to, []string{}, text)
}

// SendEmailWithTextAndCC 发送邮件(Text格式，附CC)
func (emailUtils *emailUtils) SendEmailWithTextAndCC(subject string, to []string, cc []string, text string) error {
	err := sendEmail(subject, to, cc, text, "")
	if err != nil {
		global.Log.Error(fmt.Sprintf("'SendEmailWithText' method error, failed to send email!subject:%s;to:%s;cc:%s;error:,%s",
			subject, to, cc, err.Error()))
	}
	return err
}

// SendEmailWithHTML 发送邮件(HTML格式)
func (emailUtils *emailUtils) SendEmailWithHTML(subject string, to []string, html string) error {
	return emailUtils.SendEmailWithHTMLAndCC(subject, to, []string{}, html)
}

// SendEmailWithHTMLAndCC 发送邮件(HTML格式，附CC)
func (emailUtils *emailUtils) SendEmailWithHTMLAndCC(subject string, to []string, cc []string, html string) error {
	err := sendEmail(subject, to, cc, "", html)
	if err != nil {
		global.Log.Error(fmt.Sprintf("'SendEmailWithHTML' method error, failed to send email!subject:%s;to:%s;cc:%s;error:,%s",
			subject, to, cc, err.Error()))
	}
	return err
}

// sendEmail 发送邮件
func sendEmail(subject string, to []string, cc []string, text string, html string) error {
	e := email.NewEmail()
	e.Subject = subject
	e.From = global.Config.Email.From + " <" + global.Config.Email.Username + ">"
	if len(to) == 0 {
		return errors.New("email address cannot be empty")
	}
	e.To = to
	if len(cc) != 0 {
		e.Cc = cc
	}
	if text == "" && html == "" {
		return errors.New("the email content cannot be empty")
	}
	if text != "" {
		e.Text = []byte(text)
	}
	if html != "" {
		e.HTML = []byte(html)
	}
	return e.Send(global.Config.Email.Addr, smtp.PlainAuth(global.Config.Email.Identity, global.Config.Email.Username, global.Config.Email.Password, global.Config.Email.Host))
}
