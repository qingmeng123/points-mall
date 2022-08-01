/*******
* @Author:qingmeng
* @Description:
* @File:issue
* @Date:2022/7/30
 */

package auth

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

func GetCreds() credentials.TransportCredentials{
	//申请证书，双向认证
	//从证书相关文件中读取和解析信息，得到证书公钥，密钥对
	cert,_:=tls.LoadX509KeyPair("pbfile/cert/client.pem","pbfile/cert/client.key")
	//创建一个新的，空的CertPool
	certPool:=x509.NewCertPool()
	ca,_:=ioutil.ReadFile("pbfile/cert/ca.crt")
	//解析所传入的pem编码的证书。成功则加入certPool中
	certPool.AppendCertsFromPEM(ca)
	//构建基于TLS的TransportCredentials选项
	creds:=credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		//客户端必须验证域名
		ServerName: "*.duryun.xyz",
		RootCAs: certPool,
	})
	return creds
}