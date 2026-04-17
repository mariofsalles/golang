package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// Configure initializes the securecookie with environment variables
func Configure() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encoded, err := s.Encode("cookie_1", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "cookie_1",
		Value:    encoded,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

// Read return the cookie storage
func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("cookie_1")
	if err != nil {
		return nil, err
	}
	
	cookieValues := make(map[string]string)
	if err = s.Decode("cookie_1", cookie.Value, &cookieValues); err != nil {
		return nil, err
	}
	return cookieValues, nil
}

// Delete removes the cookie from the user's browser	
func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "cookie_1",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:   time.Unix(0, 0),
	})
}