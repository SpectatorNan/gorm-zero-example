import (
	"database/sql"
	"fmt"
	"strings"
	{{if .time}}"time"{{end}}
    "github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"gorm.io/gorm"
)
