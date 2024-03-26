package sdkConfig

var Testing = New(&SDKConfig{
	SignCertificate: map[string]string{
		"default": `-----BEGIN CERTIFICATE-----
MIIEOjCCAyKgAwIBAgIFEAJkAUkwDQYJKoZIhvcNAQEFBQAwWDELMAkGA1UEBhMC
Q04xMDAuBgNVBAoTJ0NoaW5hIEZpbmFuY2lhbCBDZXJ0aWZpY2F0aW9uIEF1dGhv
cml0eTEXMBUGA1UEAxMOQ0ZDQSBURVNUIE9DQTEwHhcNMTUxMjA0MDMyNTIxWhcN
MTcxMjA0MDMyNTIxWjB5MQswCQYDVQQGEwJjbjEXMBUGA1UEChMOQ0ZDQSBURVNU
IE9DQTExEjAQBgNVBAsTCUNGQ0EgVEVTVDEUMBIGA1UECxMLRW50ZXJwcmlzZXMx
JzAlBgNVBAMUHjA0MUBaMTJAMDAwNDAwMDA6U0lHTkAwMDAwMDA2MjCCASIwDQYJ
KoZIhvcNAQEBBQADggEPADCCAQoCggEBAMUDYYCLYvv3c911zhRDrSWCedAYDJQe
fJUjZKI2avFtB2/bbSmKQd0NVvh+zXtehCYLxKOltO6DDTRHwH9xfhRY3CBMmcOv
d2xQQvMJcV9XwoqtCKqhzguoDxJfYeGuit7DpuRsDGI0+yKgc1RY28v1VtuXG845
fTP7PRtJrareQYlQXghMgHFAZ/vRdqlLpVoNma5C56cJk5bfr2ngDlXbUqPXLi1j
iXAFb/y4b8eGEIl1LmKp3aPMDPK7eshc7fLONEp1oQ5Jd1nE/GZj+lC345aNWmLs
l/09uAvo4Lu+pQsmGyfLbUGR51KbmHajF4Mrr6uSqiU21Ctr1uQGkccCAwEAAaOB
6TCB5jAfBgNVHSMEGDAWgBTPcJ1h6518Lrj3ywJA9wmd/jN0gDBIBgNVHSAEQTA/
MD0GCGCBHIbvKgEBMDEwLwYIKwYBBQUHAgEWI2h0dHA6Ly93d3cuY2ZjYS5jb20u
Y24vdXMvdXMtMTQuaHRtMDgGA1UdHwQxMC8wLaAroCmGJ2h0dHA6Ly91Y3JsLmNm
Y2EuY29tLmNuL1JTQS9jcmw0NDkxLmNybDALBgNVHQ8EBAMCA+gwHQYDVR0OBBYE
FAFmIOdt15XLqqz13uPbGQwtj4PAMBMGA1UdJQQMMAoGCCsGAQUFBwMCMA0GCSqG
SIb3DQEBBQUAA4IBAQB8YuMQWDH/Ze+e+2pr/914cBt94FQpYqZOmrBIQ8kq7vVm
TTy94q9UL0pMMHDuFJV6Wxng4Me/cfVvWmjgLg/t7bdz0n6UNj4StJP17pkg68WG
zMlcjuI7/baxtDrD+O8dKpHoHezqhx7dfh1QWq8jnqd3DFzfkhEpuIt6QEaUqoWn
t5FxSUiykTfjnaNEEGcn3/n2LpwrQ+upes12/B778MQETOsVv4WX8oE1Qsv1XLRW
i0DQetTU2RXTrynv+l4kMy0h9b/Hdlbuh2s0QZqlUMXx2biy0GvpF2pR8f+OaLuT
AtaKdU4T2+jO44+vWNNN2VoAaw0xY6IZ3/A1GL0x
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
	FrontTransPath: "https://gateway.test.95516.com/gateway/api/frontTransReq.do",
	//后台请求地址
	BackTransPath: "https://gateway.test.95516.com/gateway/api/backTransReq.do",
	//批量交易
	BatchTransPath: "https://gateway.test.95516.com/gateway/api/batchTrans.do",
	//单笔查询请求地址
	SingleQueryPath: "https://gateway.test.95516.com/gateway/api/queryTrans.do",
	//有卡交易地址
	CardRequestPath: "https://gateway.test.95516.com/gateway/api/cardTransReq.do",
	//App交易地址
	AppRequestPath: "https://gateway.test.95516.com/gateway/api/appTransReq.do",
	//文件传输请求地址
	FileQueryPath: "https://filedownload.test.95516.com/",
	//issBackTransReq
	IssBackTransPath: "https://c2c.95516.com/qrc/api/issBackTransReq.do",
})
