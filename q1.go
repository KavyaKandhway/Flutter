package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/twinj/uuid"
)

const SENDGRID_API_KEY = "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"

var hmacSampleSecret []byte = []byte("jdnfksdmfksd")
var (
	router = gin.Default()
)

// func getToken(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()
// 	query.Add("token", token["acsess_token"])
// }
func main() {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Unauthorised Access.")
	})
	router.POST("/login", Login)
	log.Fatal(router.Run(":8080"))
	router.HandleFunc("/signup", ExtractToken)

}

type User struct {
	ID       uint64 `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// This will come from our Atlas Cluster
var user = User{
	ID:       1,
	Email:    "email",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}

func createLink(tokens map[string]string) string {

	var host = "http://localhost:8080"
	var endpoint = "/linksignup"
	var finalURL string
	finalURL = host + endpoint + "?token=" + tokens["access_token"]
	return finalURL

}
func sendMail(u User, tokens map[string]string) {
	var linkSignup string = createLink(tokens)
	from := mail.NewEmail("Example User", "test@furrble.com")
	subject := "Link Verificationn"
	to := mail.NewEmail("Example User", u.Email)
	plainTextContent := "Link---" + linkSignup
	htmlContent := "link " + linkSignup
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	ts, err := CreateToken(u.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
	u.ID = 100
	c.JSON(http.StatusOK, u)
	sendMail(u, tokens)
}

func CreateToken(userid uint64) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("token")
	//normally Authorization the_token_xxx
	log.Println(bearToken)
	return bearToken
}

// // sample token string taken from the New example
// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"
func VerifyToken(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}
