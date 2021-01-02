module github.com/Jecosine/obm-back-end

go 1.15

require (
	github.com/go-ini/ini v1.62.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.16
)

replace github.com/Jecosine/obm-back-end/pkg/setting => ./pkg/setting
replace github.com/Jecosine/obm-back-end/pkg/e => ./pkg/e
replace github.com/Jecosine/obm-back-end/models => ./models
