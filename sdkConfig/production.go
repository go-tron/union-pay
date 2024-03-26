package sdkConfig

var Production = New(&SDKConfig{
	SignCertificate: map[string]string{
		"default": `-----BEGIN CERTIFICATE-----
MIIEIDCCAwigAwIBAgIFEDRVM3AwDQYJKoZIhvcNAQEFBQAwITELMAkGA1UEBhMC
Q04xEjAQBgNVBAoTCUNGQ0EgT0NBMTAeFw0xNTEwMjcwOTA2MjlaFw0yMDEwMjIw
OTU4MjJaMIGWMQswCQYDVQQGEwJjbjESMBAGA1UEChMJQ0ZDQSBPQ0ExMRYwFAYD
VQQLEw1Mb2NhbCBSQSBPQ0ExMRQwEgYDVQQLEwtFbnRlcnByaXNlczFFMEMGA1UE
Aww8MDQxQDgzMTAwMDAwMDAwODMwNDBA5Lit5Zu96ZO26IGU6IKh5Lu95pyJ6ZmQ
5YWs5Y+4QDAwMDE2NDkzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
tXclo3H4pB+Wi4wSd0DGwnyZWni7+22Tkk6lbXQErMNHPk84c8DnjT8CW8jIfv3z
d5NBpvG3O3jQ/YHFlad39DdgUvqDd0WY8/C4Lf2xyo0+gQRZckMKEAId8Fl6/rPN
HsbPRGNIZgE6AByvCRbriiFNFtuXzP4ogG7vilqBckGWfAYaJ5zJpaGlMBOW1Ti3
MVjKg5x8t1/oFBkpFVsBnAeSGPJYrBn0irfnXDhOz7hcIWPbNDoq2bJ9VwbkKhJq
Vz7j7116pziUcLSFJasnWMnp8CrISj52cXzS/Y1kuaIMPP/1B0pcjVqMNJjowooD
OxID3TZGfk5V7S++4FowVwIDAQABo4HoMIHlMB8GA1UdIwQYMBaAFNHb6YiC5d0a
j0yqAIy+fPKrG/bZMEgGA1UdIARBMD8wPQYIYIEchu8qAQEwMTAvBggrBgEFBQcC
ARYjaHR0cDovL3d3dy5jZmNhLmNvbS5jbi91cy91cy0xNC5odG0wNwYDVR0fBDAw
LjAsoCqgKIYmaHR0cDovL2NybC5jZmNhLmNvbS5jbi9SU0EvY3JsMjI3Mi5jcmww
CwYDVR0PBAQDAgPoMB0GA1UdDgQWBBTEIzenf3VR6CZRS61ARrWMto0GODATBgNV
HSUEDDAKBggrBgEFBQcDAjANBgkqhkiG9w0BAQUFAAOCAQEAHMgTi+4Y9g0yvsUA
p7MkdnPtWLS6XwL3IQuXoPInmBSbg2NP8jNhlq8tGL/WJXjycme/8BKu+Hht6lgN
Zhv9STnA59UFo9vxwSQy88bbyui5fKXVliZEiTUhjKM6SOod2Pnp5oWMVjLxujkk
WKjSakPvV6N6H66xhJSCk+Ref59HuFZY4/LqyZysiMua4qyYfEfdKk5h27+z1MWy
nadnxA5QexHHck9Y4ZyisbUubW7wTaaWFd+cZ3P/zmIUskE/dAG0/HEvmOR6CGlM
55BFCVmJEufHtike3shu7lZGVm2adKNFFTqLoEFkfBO6Y/N6ViraBilcXjmWBJNE
MFF/yA==
-----END CERTIFICATE-----`,
		"new": `-----BEGIN CERTIFICATE-----
MIIEKzCCAxOgAwIBAgIFEpVGRCEwDQYJKoZIhvcNAQEFBQAwITELMAkGA1UEBhMC
Q04xEjAQBgNVBAoTCUNGQ0EgT0NBMTAeFw0yMDA3MTYwOTM4MzRaFw0yNTA3MTYw
OTM4MzRaMIGWMQswCQYDVQQGEwJjbjESMBAGA1UEChMJQ0ZDQSBPQ0ExMRYwFAYD
VQQLEw1Mb2NhbCBSQSBPQ0ExMRQwEgYDVQQLEwtFbnRlcnByaXNlczFFMEMGA1UE
Aww8MDQxQDgzMTAwMDAwMDAwODMwNDBA5Lit5Zu96ZO26IGU6IKh5Lu95pyJ6ZmQ
5YWs5Y+4QDAwMDE2NDk0MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
r50XGgVgM+8NnK3fDoMkqy0E+KcnnA6lQflB0Oet1zemVIzzn+76tPS0vV02OcpV
u9dPt5iq83pMKBLY9isUuyRUWz8fn8Z7o3KvBoCRK4edtui/ihUt5vysJ920s8aG
CbBRAdRmdIa44ha6W61KEJqrhw5iI2QkDK6OgVxs7imXgYiMc5lxLQL+9bRRGbKq
zCAidolds633dQC58GZCtKIGvnwuDo8GGVTtjci7OU4c+54vtss2aDnE4QfLY4OY
1y+YXqy0D8Pax9T8ZnX7op8rCcO7FyH+0xgYA6gGnFlE3puiqxCFXCD7QI0np/bA
XuZ6tIoBrqKGvsUobVO3swIDAQABo4HzMIHwMB8GA1UdIwQYMBaAFNHb6YiC5d0a
j0yqAIy+fPKrG/bZMEgGA1UdIARBMD8wPQYIYIEchu8qAQEwMTAvBggrBgEFBQcC
ARYjaHR0cDovL3d3dy5jZmNhLmNvbS5jbi91cy91cy0xNC5odG0wOAYDVR0fBDEw
LzAtoCugKYYnaHR0cDovL2NybC5jZmNhLmNvbS5jbi9SU0EvY3JsMjQ5NjMuY3Js
MAsGA1UdDwQEAwID6DAdBgNVHQ4EFgQUQP9Yqy8KJGuiHVVGrE1k+OryQyYwHQYD
VR0lBBYwFAYIKwYBBQUHAwIGCCsGAQUFBwMEMA0GCSqGSIb3DQEBBQUAA4IBAQBk
OfvkzBq1GSgBCED+ETg15xpIz/Zujb9PkgNY0UinywYIjkn6dfluMIk2cNiCOMfM
Rg6LhtFi01Fnn3qwHe2vCEVBPJlazSsFE61tRCBTTWm8p/zfZKI9wGyir5aYBiPC
TRPgXaQ4cYqSAh1n98a4ONBy2/StBl+TfKvCIoXARUSp12lOVY/aKg+8Jk4MIvEw
8WCL98tTVxXe1nWPlpFDS9y0ivMyfYlWkTb6+0gMrYA2nzrfFGS1KZNRBS7p3Bh5
tdBPIgSd5gLZpAun8d0C3CcRZhcIof9hmxIc9ieQoWas52oVZDzsaGTo9rsTo9nU
3N3BThugW+P/koUnIFRG
-----END CERTIFICATE-----`,
	},
	EncryptCertId: "69042905377",
	EncryptPublicKeyPem: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3cGxf+TInYFGHYhbgbJh
8WziDlFwgQ+HMZ+jQjO0N6ozxx/BEaqSVq9ge5l8Ubpc9vU3+nWtdCXDIEnpRDCC
dW4ALJZr34Kp/rrhc2n68hXH2Cuqiv2XOskrqNM+t3nsAk26GkUYBbR1ECN8XlkB
2lnnqBiJYWCnbNMhcaNegDQwepgoEYwxh0VJncSRGGwnSPIl5oF6nZWawUO04OWJ
bRflP5xKA+fX7POUfC7Wy+YFjGHdmkRjeETBHwpDCNrl3lvSRRnl4J6mD07IHzL4
ro/lXEI3xgfBWxcVjPWukSaManao5s7YD6/IlpoJ20HcB6mmwYvAYIhdD+3nDKM6
oQIDAQAB
-----END PUBLIC KEY-----`,
	//前台请求地址
	FrontTransPath: "https://gateway.95516.com/gateway/api/frontTransReq.do",
	//后台请求地址
	BackTransPath: "https://gateway.95516.com/gateway/api/backTransReq.do",
	//批量交易
	BatchTransPath: "https://gateway.95516.com/gateway/api/batchTrans.do",
	//单笔查询请求地址
	SingleQueryPath: "https://gateway.95516.com/gateway/api/queryTrans.do",
	//有卡交易地址
	CardRequestPath: "https://gateway.95516.com/gateway/api/cardTransReq.do",
	//App交易地址
	AppRequestPath: "https://gateway.95516.com/gateway/api/appTransReq.do",
	//文件传输请求地址
	FileQueryPath: "https://filedownload.95516.com/",
	//issBackTransReq
	IssBackTransPath: "https://c2c.95516.com/qrc/api/issBackTransReq.do",
})
