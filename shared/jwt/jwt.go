package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// var privateKEY = `-----BEGIN RSA PRIVATE KEY-----
// MIIEoQIBAAKCAQEAvrRxA8DBr91o1fyiCI4TjaEv/+FbWDvNqiBUUhagKC2MGzzC
// 3TFxUKoyd52pnc9QP2DEV4wezKlesnhVPQF914VZ2VRxiN4dztAiJxLDRawPdXay
// EcPcy0e0rx0QKlURfxT056a2VPhV6k70FFW/pSPNzJ2rORwW2kO1lKGx7dr/NzN8
// HD3uoHULCeWap+u7/WgQA9aUwa0xVSXuRylcrpIMbYzzwKKcygnWA+LSugqsQSRi
// 0Iu+GIiNffN596mrYx35P8aprFkBZsMXG7u1qNO3KXjj5eGxhGGW4tCL0FDTK4On
// 4dGYORoqJyY1k5df58dcyuF1yb9APU8KvPmI/wIDAQABAoH/SzzlWZWy1K+FrOb/
// BmEHVdlCFrHSV/1AJt6aZciHZ1KQ9C4Esz/OdQSw5IBPavNftaF31RJzFEu+EKto
// 9aig/WafDB9Eq7r2B6IV11SPte0sLCuoFVowwgKIRo4w0oK9ZGXgOYSMavmO6+I6
// actq0LxPdWq/IsjyOyp7eShSCoStye1/hGmSqAJ8li1EjT7s65TemDfXGOhBsRDe
// B1ZM9i+YNEtNEuxmigGgn5uoMIn7hGW1PZ2BiIRwjIwUovtHGI8MGa6xKdDcAP6Q
// hSDxE5wCjd4xqzjLP6ktt8usEQrfMY68AYKOU83N5Fk/FOYvZUqO6NSGcQ9gWXf9
// KNV5AoGBAOGTiSrLMMswrZH9hvKT8K93/LHV6FHxx9lMiRVPkfnZrBsncxMhZPqi
// /WHCISI6clI+VeKunP6dMY/BKbJfbNcRA6STfgXmMZtCTLXtbraNnbVf2hsxF4Xw
// Ev2Mi0IbZZ+iOWqXnKHRCsUwRcgySzPRCQAsHkxyyJOSFmWwR/0dAoGBANhs5mwR
// 6WzEexImkbf5VRY9Kf/zaVtb8R7A0zFwvJI7lfVaqRJh1wtvhbIG3NwZtlTMncV8
// XGFSuelKwx1OkzmCrPpKHhOwlt26YKlhxX1Rk5HZIzlnsi90vQpvrsEvTjV4/mHh
// qqrHdvtHmRsuPEnB5H+YtWLyoPUivh+nwa/LAoGADPnpuBZ8NhQUyAMnkijEfbOP
// S8OcW6pm7q8ia6FqKk9FQUKhsgYHwwtBPDBoq4llLin70tBso5DzWuuntGUc47pM
// 1VjOtRQq4l2MACMqbUH0QozDBTFrwv6uePtuv1zIGcjBOMqD7iMSVYmTWhLalJ5/
// wAzJqWgo9aQ/uZXMblkCgYBngVWGE97uLN2rJZUFRpJh62idx1z423Tqv0+B0qfs
// y+CBEhXP+8jr6C8poEyaWxWvYpiF7V8FEJpnL2E3L/ALTirKHQ5bXlYYvt0hxOe9
// cnlABHfrKWO3fH44codCTwx5WF9YkqObv39w16IqtKcSo09TksYVB3LhKfeBjip8
// lwKBgQDVpQPBhPfOOtohm70FYnL72kw4Bf5Ie8woBtNEVMoFcv3RD2COXbsPYz5x
// 3YWEYBEPOGdLZxpmDvdp8wvB/pH+QzjudRQk2xXjWSmviQM50PTJJKJc87P/pJW7
// lWxhPS976mjUWPtwhrHOC46mPT6pWYoSOzN+kGomRk5yWDb2Tg==
// -----END RSA PRIVATE KEY-----`

var publicKEY = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvrRxA8DBr91o1fyiCI4T
jaEv/+FbWDvNqiBUUhagKC2MGzzC3TFxUKoyd52pnc9QP2DEV4wezKlesnhVPQF9
14VZ2VRxiN4dztAiJxLDRawPdXayEcPcy0e0rx0QKlURfxT056a2VPhV6k70FFW/
pSPNzJ2rORwW2kO1lKGx7dr/NzN8HD3uoHULCeWap+u7/WgQA9aUwa0xVSXuRylc
rpIMbYzzwKKcygnWA+LSugqsQSRi0Iu+GIiNffN596mrYx35P8aprFkBZsMXG7u1
qNO3KXjj5eGxhGGW4tCL0FDTK4On4dGYORoqJyY1k5df58dcyuF1yb9APU8KvPmI
/wIDAQAB
-----END PUBLIC KEY-----
`

// type jwtAuthentication struct {
// 	privateKey *rsa.PrivateKey
// 	PublicKey  *rsa.PublicKey
// }

type MaropostClaim struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	AccountID string `json:"accountID"`
	jwt.StandardClaims
}

var PublicKey *rsa.PublicKey
var PrivateKey *rsa.PrivateKey

// var PublicKeyCERT string

//var authBackendInstance *jwtAuthentication

// init Parse Private and Public Keys for JWT Verification
// func init() {

// 	// fmt.Println("Logging service initiated :")
// 	// privateKey, privateErr := parsePrivateKey()
// 	// if privateErr != nil {
// 	// 	fmt.Println("unable to Parse Private Key :", privateErr)
// 	// 	//fmt.Println("unable to Parse Private Key :", privateErr)
// 	// 	os.Exit(1)
// 	// }
// 	// publicKey, publicErr := ParsePublicKey()
// 	// if publicErr != nil {
// 	// 	fmt.Println("unable to Parse Private Key :", publicErr)
// 	// 	//fmt.Println("unable to Parse Private Key :", publicErr)
// 	// 	os.Exit(1)
// 	// }

// 	// authBackendInstance = &jwtAuthentication{
// 	// 	privateKey: privateKey,
// 	// 	PublicKey:  publicKey,
// 	// }

// }

//GenerateJWT check jwt valid or not
//func GenerateJWT(userName, accountID string) (string, error) {
func GenerateJWT(email, firstName, lastName, accountID string) (string, error) {
	//https://www.madboa.com/geek/openssl/#key-rsa

	// Create the Claims
	claims := MaropostClaim{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		AccountID: accountID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "maropost-relay",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(PrivateKey)

}

//ValidateJWTToken check jwt valid or not
func ValidateJWTToken(jwtToken string) (bool, error) {

	token, err := jwt.ParseWithClaims(jwtToken, &MaropostClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return PublicKey, nil
	})

	//Malformed token,
	if err != nil || !token.Valid {
		return false, errors.New("Token is not valid")
	}

	claim, ok := token.Claims.(*MaropostClaim)
	if !ok {
		fmt.Printf("%v %v %v\n", claim.Email, claim.StandardClaims.ExpiresAt, claim.Issuer)
		return false, errors.New("Invalid JWT Token Claim")
	}

	if claim.Issuer != "maropost-relay" {
		return false, errors.New("Invaid JWT Issuer")
	}

	if claim.ExpiresAt < time.Now().Unix() {
		return false, errors.New("JWT Token Expired")
	}

	return true, nil
}

//GetClaimFromJWTToken check jwt valid or not
func GetClaimFromJWTToken(jwtToken string) (*MaropostClaim, error) {

	token, err := jwt.ParseWithClaims(jwtToken, &MaropostClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return PublicKey, nil
	})

	//Malformed token,
	if err != nil || !token.Valid {
		return nil, errors.New("Invalid JWT Token")
	}

	claim, ok := token.Claims.(*MaropostClaim)
	if !ok {
		fmt.Printf("%v %v %v\n", claim.Email, claim.StandardClaims.ExpiresAt, claim.Issuer)
		return nil, errors.New("Invalid JWT Token ")
	}

	return claim, nil
}

//ParsePublicKey ...
func ParsePublicKey(publicKeys string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(string(debug.Stack()), "execption in ParsePublicKey : ", err)
		}
	}()

	fmt.Println("publicKeyCERT : ", publicKEY)
	data, _ := pem.Decode([]byte(publicKEY))
	fmt.Println("data : ", data)
	fmt.Println("data.Bytes : ", string(data.Bytes))

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)
	if err != nil {
		fmt.Println("parsePublicKey x509 ParsePKIXPublicKey error: ", err)
		return
	}
	var ok bool
	PublicKey, ok = publicKeyImported.(*rsa.PublicKey)
	if !ok {
		return
	}
}

//GenerateIds ...
func GenerateIds() string {
	uuids, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("There is problem to generate uuid")
	}
	return uuids.String()
}
