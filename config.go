package main


var (
	
	state          = "anystring"                                // 授权服务器会原封不动的返回,可以作为定位用户的字段
	redirect_uri   = "http://localhost:5001/code"               // 重定向网址，用于接收授权code响应
	code_challenge = genCodeChallenge()                         // 生成code_challenge
	authUrlFormat  = "https://login.microsoftonline.com/organizations/oauth2/v2.0/authorize?client_id=%s&response_type=code&redirect_uri=%s&response_mode=query&scope=%s&state=%s&code_challenge=%s&code_challenge_method=S256"
	tokenUri       = "https://login.microsoftonline.com/organizations/oauth2/v2.0/token"
	tokenResponse  = "" // 请求token响应
)
