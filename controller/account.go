package account
import (
	"github.com/labstack/echo/v4"
	"net/http"
)
func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Home")
}
func Login(c echo.Context) error {
	return c.String(http.StatusOK, "Login")
}
