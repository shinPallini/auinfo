package roles

import (
	"log"
	"sort"

	"github.com/bwmarrin/discordgo"
	"github.com/shinPallini/discordgox"
)

type AuInfo struct {
	// インポスターやクルー陣営についての情報
	camp string
	// セレクトメニューに追加する表示名
	name string

	// 役職の説明
	description string
}

type AuRoles string

// セレクトメニューが選ばれたときに送信される値を定数で指定
const (
	BountyHunter AuRoles = "1"
	FireWorks    AuRoles = "2"
	Mare         AuRoles = "3"
	Bait         AuRoles = "101"
	Dictator     AuRoles = "102"
	Doctor       AuRoles = "103"
	Arsonist     AuRoles = "201"
)

func NewAuInfo(camp, name, description string) AuInfo {
	return AuInfo{
		camp:        camp,
		name:        name,
		description: description,
	}
}

// func (a *AuInfo) displayName() string {
// 	return fmt.Sprintf("%s: %s", a.displayName, a.camp)
// }

var (
	// 複数選択できるMenuの選択最小値
	min_value = 1
	// 値受け渡し用のcustomID
	customId = "select-roles"
)

func init() {
	imposter := "インポスター陣営"
	crew := "クルー陣営"
	third := "第3陣営"

	roles := map[AuRoles]AuInfo{
		BountyHunter: NewAuInfo(imposter, "バウンティハンター", "ターゲットをキルした場合、直後のキルクールダウンが半分になる"),
		FireWorks:    NewAuInfo(imposter, "花火職人", "花火を最大3個設置できる。\n最後のインポスターになったときシェイプシフトのタイミングで一斉に起爆させる"),
		Mare:         NewAuInfo(imposter, "メアー", "停電の時以外ではキルができないが、キルクールは半分になる。\n停電中のみ移動速度が上昇するが名前が赤く表示される"),
		Bait:         NewAuInfo(crew, "ベイト", "自分をキルしたプレイヤーに強制で自分の死体を通報させることができる"),
		Dictator:     NewAuInfo(crew, "ディクテーター", "会議中に誰かが投票すると会議を強制終了させて投票先を釣ることができる。\n投票したタイミングでディクテーターは死亡する"),
		Doctor:       NewAuInfo(crew, "ドクター", "プレイヤーの死因を知ることができて、遠隔でバイタルを見ることができる"),
		Arsonist:     NewAuInfo(third, "アーソニスト", "キルボタンを押して一定時間近くにいれば相手にオイルを塗れる。\n全員にオイルを塗ってベントに入ると起爆して単独勝利となる"),
	}

	selectMenuOption := make([]discordgo.SelectMenuOption, 0)
	for role, info := range roles {
		selectMenuOption = append(selectMenuOption, *discordgox.NewSelectMenuOption(
			info.name,
			string(role),
			discordgox.SetSelectDescription(info.camp),
		))
	}

	sort.SliceStable(selectMenuOption, func(i, j int) bool { return selectMenuOption[i].Description < selectMenuOption[j].Description })

	cmdResponse := discordgox.NewInteractionResponse(
		discordgox.SetType(discordgo.InteractionResponseChannelMessageWithSource),
		discordgox.SetData(discordgox.NewInteractionResponseData(
			//discordgox.SetContent("今回使用する役職を選んでください！"),
			discordgox.SetEmbed(
				discordgox.NewList(
					discordgox.NewMessageEmbed(
						// SetAuthorでこのボットの名前を表示したい
						discordgox.SetTitle("今回使用する役職を選んでください！"),
						discordgox.SetColor(0x21ed43),
					),
				),
			),
			discordgox.SetComponent(
				discordgox.NewList[discordgo.MessageComponent](
					discordgox.NewActionsRow(
						discordgox.SetMultiSelectMenu(
							customId,
							selectMenuOption,
							&min_value,
							len(selectMenuOption),
						),
					),
				),
			),
		),
		),
	)

	// cmpResponse := discordgox.NewInteractionResponse(
	// 	discordgox.SetType(discordgo.InteractionResponseChannelMessageWithSource),
	// 	discordgox.SetData(discordgox.NewInteractionResponseData(
	// 		discordgox.SetContent("今回使用する役職を選んでください！"),
	// 		discordgox.SetComponent(
	// 			discordgox.NewList[discordgo.MessageComponent](
	// 				discordgox.NewActionsRow(
	// 					discordgox.SetMultiSelectMenu(
	// 						"select-roles",
	// 						selectMenuOption,
	// 						&min_value,
	// 						len(selectMenuOption),
	// 					),
	// 				),
	// 			),
	// 		),
	// 	),
	// 	),
	// )

	discordgox.AddCommand(
		&discordgo.ApplicationCommand{
			Name:        "auroles",
			Description: "Town of Hostで使える役職一覧を表示",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, cmdResponse)
		},
	)

	// discordgox.AddCommandWithComponent(
	// 	&discordgo.ApplicationCommand{
	// 		Name:        "auroles",
	// 		Description: "Town of Hostで使える役職一覧を表示",
	// 	},
	// 	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 		s.InteractionRespond(i.Interaction, cmdResponse)
	// 	},
	// 	customId,
	// 	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// 		data := i.MessageComponentData().Values
	// 		for _, d := range data {
	// 			v, ok :=
	// 		}
	// 		respEmbed := discordgox.NewMessageEmbed(
	// 			// SetAuthorしたい
	// 			discordgox.SetTitle("使用役職一覧"),
	// 			discordgox.SetEmbedField()
	// 			)
	// 		)
	// 	},
	// )
	log.Println(cmdResponse.Data)
	log.Println(discordgox.Commands)
}
