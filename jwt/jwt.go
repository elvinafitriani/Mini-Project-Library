package jwt

import (
	"library/auth"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func Token(id int, username string) (*auth.TokenDetail, error) {
	errEnv := godotenv.Load()
	if errEnv != nil {
		return nil, errEnv
	}

	form := os.Getenv("FORM")
	secret := os.Getenv("SECRET")

	td := &auth.TokenDetail{}
	exp := time.Now().Add(time.Hour)
	Uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	RefUid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	td.ExpAt = exp.Unix()
	td.AccId = Uid.String()
	td.RefAccID = RefUid.String()

	Claims := auth.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    form,
			ExpiresAt: td.ExpAt,
			Id:        td.AccId,
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
		Uid:  id,
		User: username,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims,
	)

	RefClaims := auth.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    form,
			NotBefore: time.Now().Unix(),
			Id:        td.RefAccID,
			IssuedAt:  time.Now().Unix(),
		},
	}
	refToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		RefClaims,
	)

	signAc, err := token.SignedString([]byte(secret))

	if err != nil {
		return nil, err
	}

	refTokSig, errRef := refToken.SignedString([]byte(secret))

	if errRef != nil {
		return nil, err
	}

	td.RefAccToken = refTokSig
	td.AccsesToken = signAc

	return td, nil
}
