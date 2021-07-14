package ig
import (
	"github.com/labstack/echo/v4"
	"net/http"
)
func IgAuth(c echo.Context) error {
  return c.String(http.StatusOK, "IgAuth")
}
