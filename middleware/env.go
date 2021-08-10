package middleware

import "gopractice/middleware/security"

func Env()  {
	security.Security()
}