package constants

// SmsScene defines constants for different SMS scenes.
type SmsScene string

const (
	SmsSceneUserRegister      SmsScene = "user.register"      // 短信场景：注册
	SmsSceneUserLogin         SmsScene = "user.login"         // 短信场景：登录
	SmsSceneUserResetPassword SmsScene = "user.resetPassword" // 短信场景：重置密码
	SmsSceneUserChangeMobile  SmsScene = "user.changeMobile"  // 短信场景：修改手机号
	SmsSceneUserChangeEmail   SmsScene = "user.changeEmail"   // 短信场景：修改邮箱
)
