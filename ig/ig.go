package ig
import (
  "net/http"
  "github.com/labstack/echo/v4"
)
func IgAuth(c echo.Context) error {
  return c.String(http.StatusOK, "IgAuth")
}
