package main

import (
	"net"
	"net/http"

	"github.com/alanadiprastyo/learn-by-code/kube"
	"github.com/c-robinson/iplib"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/ips", func(c *gin.Context) {
		// ips := new(kube.IpAddress)

		// ips.Ip = c.PostForm("ip")
		// ips.Netmask, _ = strconv.Atoi(c.PostForm("netmask"))
		// ips.Count, _ = strconv.Atoi(c.PostForm("count"))

		var ips kube.IpAddress
		c.BindJSON(&ips)

		//debug trace data
		//fmt.Println(ips.Ip, ips.Netmask, count)

		//initial variable slice interface
		var totalIp []interface{}

		for i := 1; i <= ips.Count; i++ {
			ipa := net.ParseIP(ips.Ip)
			ip4 := iplib.NewNet4(ipa, ips.Netmask)
			bAddr := ip4.BroadcastAddress()
			newIp := iplib.NextIP(bAddr)
			newNetwork := iplib.NewNet4(newIp, ips.Netmask)
			ips.Ip = newIp.String()
			//fmt.Println(newNetwork.String())
			totalIp = append(totalIp, newNetwork.String())
		}

		response := gin.H{
			"Network": totalIp,
		}

		c.JSON(http.StatusOK, response)
	})

	router.Run("localhost:8080")
}
