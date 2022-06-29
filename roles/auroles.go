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
	// セレクトメニューが選ばれたときに送信される値
	value AuRoles
	// 役職の説明
	description string
}

type AuRoles string

const (
	BountyHunter AuRoles = "1"
	FireWorks    AuRoles = "2"
	Mare         AuRoles = "3"
	Bait         AuRoles = "101"
	Dictator     AuRoles = "102"
	Doctor       AuRoles = "103"
)

func NewAuInfo(camp string, name string, value AuRoles, description string) AuInfo {
	return AuInfo{
		camp:        camp,
		name:        name,
		value:       value,
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
	// third := "第3陣営"

	roles := []AuInfo{
		NewAuInfo(imposter, "バウンティハンター", BountyHunter, "ターゲットをキルした場合、直後のキルクールダウンが半分になる"),
		NewAuInfo(imposter, "花火職人", FireWorks, "花火を最大3個設置できる。\n最後のインポスターになったときシェイプシフトのタイミングで一斉に起爆させる"),
		NewAuInfo(imposter, "メアー", Mare, "停電の時以外ではキルができないが、キルクールは半分になる。\n停電中のみ移動速度が上昇するが名前が赤く表示される"),
		NewAuInfo(crew, "ベイト", Bait, "自分をキルしたプレイヤーに強制で自分の死体を通報させることができる"),
		NewAuInfo(crew, "ディクテーター", Dictator, "会議中に誰かが投票すると会議を強制終了させて投票先を釣ることができる。\n投票したタイミングでディクテーターは死亡する"),
		NewAuInfo(crew, "ドクター", Doctor, "プレイヤーの死因を知ることができて、遠隔でバイタルを見ることができる"),
	}

	// info := map[string]string{
	// 	addpref(prefInposter, "バウンティハンター"): "ターゲットをキルした場合、直後のキルクールダウンが半分になる",
	// 	addpref(prefInposter, "花火職人"):      "花火を最大3個設置できる。\n最後のインポスターになったときシェイプシフトのタイミングで一斉に起爆させる",
	// }
	// var keys []string
	var selectMenuOption = []discordgo.SelectMenuOption{}
	for _, role := range roles {
		selectMenuOption = append(selectMenuOption, *discordgox.NewSelectMenuOption(
			role.name,
			string(role.value),
			discordgox.SetSelectDescription(role.camp),
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
						discordgox.SetTitle("今回使用する役職を選んでください！"),
						discordgox.SetDescription("役職一覧"),
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
	log.Println(cmdResponse.Data)
	log.Println(discordgox.Commands)
}
