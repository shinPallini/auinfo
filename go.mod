module auinfo

go 1.18

require (
	github.com/bwmarrin/discordgo v0.25.0
	github.com/joho/godotenv v1.4.0
	github.com/shinPallini/discordgox v0.0.5
	roles v0.0.0-00010101000000-000000000000
)

require (
	github.com/gorilla/websocket v1.4.2 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/sys v0.0.0-20201119102817-f84b799fce68 // indirect
)

replace roles => ./roles
