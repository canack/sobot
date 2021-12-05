package comm

// import (
// 	"bufio"
// 	"log"
// 	"net"
// 	"strings"

// 	sobot "github.com/canack/sobot/src"
// )

// var username, password, platform, file, comment string

// func StartServer(connHost, connPort string) {
// 	log.Println("Starting tcp server on " + connHost + ":" + connPort)
// 	l, err := net.Listen("tcp", connHost+":"+connPort)
// 	if err != nil {
// 		log.Fatalln("Error listening:", err.Error())
// 	}
// 	defer l.Close()

// 	for {
// 		c, err := l.Accept()
// 		if err != nil {
// 			log.Println("Error connecting:", err.Error())
// 			return
// 		}

// 		go handleConnection(c)

// 	}
// }

// func handleConnection(conn net.Conn) {

// 	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

// 	if err != nil {
// 		log.Println("Client left.")
// 		conn.Close()
// 		return
// 	}

// 	message := string(buffer[:len(buffer)-1])
// 	seperatedMessage := strings.Split(message, " ")
// 	// log.Println("| Address:", conn.RemoteAddr(), " | Client message:", message)
// 	if len(seperatedMessage) < 5 {
// 		// BASIT KULLANIM VERILERINI GONDER
// 		conn.Write([]byte("Yeterli argüman yok\n\n"))
// 		conn.Write([]byte("Kullanım şekli:\n"))
// 		conn.Write([]byte("kullanıcı_adı parola platformadı fotoğraf_yolu/fotoğraf.jpg açıklama_yazısı\n"))
// 		conn.Write([]byte("\nÖrnek kullanım şekli:\n"))
// 		conn.Write([]byte("alanturing 123456 instagram turing_machine.jpg Bugün turing makinesini icat ettim.\n"))
// 		// BASIT KULLANIM VERILERINI GONDER

// 	} else {
// 		conn.Write([]byte("Veriler alındı\n"))

// 		username = seperatedMessage[0]
// 		password = seperatedMessage[1]
// 		platform = seperatedMessage[2]
// 		file = seperatedMessage[3]
// 		comment = strings.Join(seperatedMessage[4:], " ")

// 		// Buradan aşağısı gelen cevaba göre programın işleyişini belirliyor.
// 		if platform == "instagram" {
// 			sobot.Instagram(username, password, &conn).SetFile(file).SetCaption(comment).Share(true)
// 			conn.Write([]byte("Başka işlem kalmadı.\n"))

// 		} else if platform == "twitter" {
// 			sobot.Twitter(username, password, &conn).SetFile(file).SetCaption(comment).Share(true)
// 			conn.Write([]byte("Başka işlem kalmadı.\n"))

// 		} else {
// 			conn.Write([]byte("\nÖrnek kullanım şekli:\n"))
// 			conn.Write([]byte("alanturing 123456 instagram turing_machine.jpg Bugün turing makinesini icat ettim.\n"))
// 			log.Println("Platform seçimi hatalı")
// 		}
// 	}

// 	conn.Close()
// 	// conn.Write(buffer) for reply
// 	//	handleConnection(conn)
// }
