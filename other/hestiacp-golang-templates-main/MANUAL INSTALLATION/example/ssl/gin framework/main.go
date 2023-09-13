package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")

	r := gin.Default()
	r.SetHTMLTemplate(html)

	//bu makale ile cozdum
	//https://www.vultr.com/docs/secure-a-golang-web-server-with-a-selfsigned-or-lets-encrypt-ssl-certificate/
	//openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout go-server.key -out go-server.crt

	//https://github.com/denji/golang-tls bu da guzel anlatiyor

	//https://stackoverflow.com/questions/41250665/go-https-client-issue-remote-error-tls-handshake-failure
	//cfg := &tls.Config{
	//	CipherSuites: []uint16{
	//		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	//		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	//	},
	//	PreferServerCipherSuites: true,
	//	InsecureSkipVerify:       true,
	//	MinVersion:               tls.VersionTLS11,
	//	MaxVersion:               tls.VersionTLS11,
	//}

	r.SetTrustedProxies([]string{"10.0.0.91"})

	r.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})

	// Listen and Server in https://127.0.0.1:8080
	log.Fatal(r.RunTLS(":8080", "/home/admin/conf/web/test.example.com/ssl/kys.example.com.crt", "/home/admin/conf/web/test.example.com/ssl/kys.example.com.key"))
}
