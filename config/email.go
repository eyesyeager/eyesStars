package config

/**
 * 邮箱配置
 * @author eyesYeager
 * @date 2024/7/25 22:28
 */

type Email struct {
	From     string // 发件人
	Addr     string // SMTP服务器的地址
	Identity string // 身份证明
	Username string // 用户名
	Password string // 密码
	Host     string // 主机地址
}
