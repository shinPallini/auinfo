package roles

import (
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

type AuRole string

// セレクトメニューが選ばれたときに送信される値を定数で指定
const (
	// インポスター陣営の役職登録
	BountyHunter    AuRole = "1"
	FireWorks       AuRole = "2"
	Mare            AuRole = "3"
	Puppeteer       AuRole = "4"
	SerialKiller    AuRole = "5"
	ShapeMaster     AuRole = "6"
	Sniper          AuRole = "7"
	TimeThief       AuRole = "8"
	Vampire         AuRole = "9"
	Warlock         AuRole = "10"
	Witch           AuRole = "11"
	Mafia           AuRole = "12"
	Madmate         AuRole = "13"
	MadGuardian     AuRole = "14"
	SidekickMadmate AuRole = "15"

	// クルー陣営の役職登録
	Bait           AuRole = "101"
	Dictator       AuRole = "102"
	Doctor         AuRole = "103"
	Lighter        AuRole = "104"
	SabotageMaster AuRole = "105"
	Sheriff        AuRole = "106"
	SpeedBooster   AuRole = "107"
	Trapper        AuRole = "108"

	// 第3陣営の役職登録
	Arsonist       AuRole = "201"
	Egoist         AuRole = "202"
	Executioner    AuRole = "203"
	Jester         AuRole = "204"
	Opportunist    AuRole = "205"
	SchrodingerCat AuRole = "206"
	Terrorist      AuRole = "207"
)

func NewAuInfo(camp, name, description string) AuInfo {
	return AuInfo{
		camp:        camp,
		name:        name,
		description: description,
	}
}

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

	roles := map[AuRole]AuInfo{
		// インポスター陣営の情報登録
		BountyHunter: NewAuInfo(imposter, "バウンティハンター", "ターゲットをキルした場合、直後のキルクールダウンが半分になる。"),
		FireWorks:    NewAuInfo(imposter, "花火職人", "花火を最大3個設置できる。\n最後のインポスターになったときシェイプシフトのタイミングで一斉に起爆させる。"),
		Mare:         NewAuInfo(imposter, "メアー", "停電の時以外ではキルができないが、キルクールは半分になる。\n停電中のみ移動速度が上昇するが名前が赤く表示される。"),
		Puppeteer:    NewAuInfo(imposter, "パペッティア", "キルした対象キャラクターに次に近づいたクルーをきるさせる。"),
		SerialKiller: NewAuInfo(imposter, "シリアルキラー", "キルクールが短いインポスター。\nその代わり時間までに次のキルをしなければ自爆する。"),
		ShapeMaster:  NewAuInfo(imposter, "シェイプマスター", "変身後のクールダウンを無視して彩度返信ができるインポスター。\n通常は10秒しか変身できないが、設定により変身継続時間を変更できる。"),
		Sniper:       NewAuInfo(imposter, "スナイパー", "遠距離射撃が可能な役職。\nシェイプシフトした地点から解除した地点の延長線上にいる対象をキルできる。\nなお、斜線上にいるクルーには射撃音が聞こえる。\n弾丸を打ち切るまで通常のキルをすることはできない。"),

		// クルー陣営の情報登録
		Bait:     NewAuInfo(crew, "ベイト", "自分をキルしたプレイヤーに強制で自分の死体を通報させることができる。"),
		Dictator: NewAuInfo(crew, "ディクテーター", "会議中に誰かが投票すると会議を強制終了させて投票先を釣ることができる。\n投票したタイミングでディクテーターは死亡する。"),
		Doctor:   NewAuInfo(crew, "ドクター", "プレイヤーの死因を知ることができて、遠隔でバイタルを見ることができる"),

		// 第3陣営の情報登録
		Arsonist: NewAuInfo(third, "アーソニスト", "キルボタンを押して一定時間近くにいれば相手にオイルを塗れる。\n全員にオイルを塗ってベントに入ると起爆して単独勝利となる。"),
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

	discordgox.AddCommandWithComponent(
		&discordgo.ApplicationCommand{
			Name:        "auroles",
			Description: "Town of Hostで使える役職一覧を表示",
		},
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			cmdResponse := discordgox.NewInteractionResponse(
				discordgox.SetType(discordgo.InteractionResponseChannelMessageWithSource),
				discordgox.SetData(discordgox.NewInteractionResponseData(
					discordgox.SetEmbed(
						discordgox.NewList(
							discordgox.NewMessageEmbed(
								discordgox.SetEmbedAuthor(
									s.State.User.Username,
									"https://github.com/shinPallini/auinfo",
									"https://cdn.discordapp.com/embed/avatars/0.png",
								),
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
			s.InteractionRespond(i.Interaction, cmdResponse)
		},
		customId,
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			data := i.MessageComponentData().Values
			embedFields := make([]*discordgo.MessageEmbedField, 0)
			for _, r := range data {
				info, ok := roles[AuRole(r)]
				if ok {
					embedFields = append(embedFields, discordgox.NewMessageEmbedField(
						discordgox.SetEmbedFieldName(info.name),
						discordgox.SetEmbedFieldValue(info.description),
						discordgox.SetEmbedFieldInline(false),
					))
					embedFields = append(embedFields, discordgox.NewMessageEmbedField(
						discordgox.SetEmbedFieldName("\u200B"),
						discordgox.SetEmbedFieldValue("----------------------------------------------"),
						discordgox.SetEmbedFieldInline(false),
					))
				}
			}
			embedFields = embedFields[:len(embedFields)-1]
			respEmbed := discordgox.NewMessageEmbed(
				discordgox.SetEmbedAuthor(
					s.State.User.Username,
					"https://github.com/shinPallini/auinfo",
					"https://cdn.discordapp.com/embed/avatars/0.png",
				),
				discordgox.SetTitle("使用役職一覧"),
				discordgox.SetEmbedField(embedFields),
				discordgox.SetColor(0x21ed43),
			)
			cmpResponse := discordgox.NewInteractionResponse(
				discordgox.SetType(discordgo.InteractionResponseChannelMessageWithSource),
				discordgox.SetData(discordgox.NewInteractionResponseData(
					discordgox.SetEmbed(discordgox.NewList(respEmbed)),
				),
				),
			)
			s.InteractionRespond(i.Interaction, cmpResponse)
		},
	)
}
