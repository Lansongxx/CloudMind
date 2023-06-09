package {{.pkg}}
{{if .withCache}}
import (
   "gorm.io/gorm"
)
{{else}}
import "github.com/zeromicro/go-zero/core/stores/sqlx"
{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
   // {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
   // and implement the added methods in custom{{.upperStartCamelObject}}Model.
   {{.upperStartCamelObject}}Model interface {
      {{.lowerStartCamelObject}}Model
   }

   custom{{.upperStartCamelObject}}Model struct {
      *default{{.upperStartCamelObject}}Model
   }
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn *gorm.DB) {{.upperStartCamelObject}}Model {
   return &custom{{.upperStartCamelObject}}Model{
      default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn),
   }
}
