package massage

import (
  "fmt"
 // "encoding/json"
//	"io/ioutil"
  "net/smtp"
//  "strconv"
  )

func EmailOTP(toEmail string, id int) error {
  subject := "Bli-Bli lite"
  body := "selamat datag di Web kami, kami akan berusahan menjaggan akun anda deggan baik, dan apakah anda baru saja membuat akun?, jaggan memberikan code verivikasi ke orang lain, jika itu bukan anda abaikan saja email ini CODE :"
  
  
    smtpHost := "smtp.gmail.com"
    smtpPort := 587
    senderEmail := "naimmmmab@gmail.com"
    
    senderPassword := "cniv hqso rmmi qppt"

   
    auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

   
    // Membentuk pesan email
    	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n%d", toEmail, subject, body, id))

    // Mengirim email melalui SMTP server Gmail
   smtp.SendMail(smtpHost+":"+fmt.Sprint(smtpPort), auth, senderEmail, []string{toEmail}, message)
   
   return nil
}