module gorm-zero-example

go 1.17

require (
	github.com/SpectatorNan/gorm-zero v1.0.0
	gorm.io/gorm v1.22.4
)

replace (
	github.com/SpectatorNan/gorm-zero v1.0.0 => ../gorm-zero
)