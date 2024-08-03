package template

import (
	"eyesStars/app/model/entity"
	"eyesStars/global"
	"time"
)

/**
 * @author eyesYeager
 * @date 2024/7/30 21:25
 */

// GetMemoryEmailTemplate 回忆邮件模版
func GetMemoryEmailTemplate(content string, publishTime time.Time, name string) entity.Email {
	subject := "星使"
	systemName := global.Config.App.Site
	author := global.Config.App.Author
	url := global.Config.App.Url
	nowTime := time.Now().Format(time.DateTime)
	publishTimeStr := publishTime.Format(time.DateTime)
	var anonymousName string
	if name == "" {
		anonymousName = name
	} else {
		anonymousName = "化名：" + name
	}
	return entity.Email{
		Subject: subject,
		Content: `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>` + systemName + ` | ` + subject + `</title>
  </head>

  <body style="margin: 0; padding: 0;">
    <table border="0" cellpadding="0" cellspacing="0" width="100%">
      <tr>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="330" style="border-collapse: collapse;">
          <tr style="border-bottom: 2px solid;">
            <td style="font-size: 25px;">` + systemName + ` | ` + subject + `</td>
          </tr>
        </table>
      </tr>
      
      <tr>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="330" style="border-collapse: collapse;">
          <tr style="height: 20px;"></tr>
          <tr>
            <td>星空无际，人亦无常。但彼时的星语不会散去，因为星使总会如约而至。</td>
          </tr>
          <tr style="height: 25px;"></tr>
          <tr>
            <td>寄语：` + content + `</td>
          </tr>
          <tr style="height: 5px;"></tr>
          <tr>
            <td>` + anonymousName + `</td>
          </tr>
          <tr style="height: 5px;"></tr>
          <tr>
            <td>时间：` + publishTimeStr + `</td>
          </tr>
        </table>
      </tr>
      <tr>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="330" style="border-collapse: collapse;">
          <tr style="height: 150px;"></tr>
          <tr>
            <td align="right" style="height: 30px;">` + author + `</td>
          </tr>
          <tr>
            <td align="right">` + nowTime + `</td>
          </tr>
        </table>
      </tr>
      <tr>
        <table align="center" border="0" cellpadding="0" cellspacing="0" width="330" style="border-collapse: collapse;">
          <tr style="height: 100px;"></tr>
          <tr style="border-top: 1px dotted">
            <td>此为` + systemName + `（` + url + `） 自动发送，无需回复。</td>
          </tr>
        </table>
      </tr>
    </table>
  </body>
</html>
`,
	}
}
