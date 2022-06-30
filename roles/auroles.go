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

type CustomSelectMenuOption []discordgo.SelectMenuOption

func (c CustomSelectMenuOption) extract(camp string) []discordgo.SelectMenuOption {
	l := []discordgo.SelectMenuOption{}
	for _, v := range c {
		if v.Description == camp {
			l = append(l, v)
		}
	}
	return l
}

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
	Mayor          AuRole = "105"
	SabotageMaster AuRole = "106"
	Sheriff        AuRole = "107"
	SpeedBooster   AuRole = "108"
	Trapper        AuRole = "109"

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
	imposter  = "インポスター陣営"
	crew      = "クルー陣営"
	third     = "第3陣営"
	customIDs = map[string]string{
		imposter: "custom-imposter",
		crew:     "custom-crew",
		third:    "custom-third",
	}
)

func init() {

	roles := map[AuRole]AuInfo{
		// インポスター陣営の情報登録
		BountyHunter:    NewAuInfo(imposter, "バウンティハンター", "ターゲットをキルした場合、直後のキルクールダウンが半分になる。"),
		FireWorks:       NewAuInfo(imposter, "花火職人", "花火を最大3個設置できる。\n最後のインポスターになったときシェイプシフトのタイミングで一斉に起爆させる。"),
		Mare:            NewAuInfo(imposter, "メアー", "停電の時以外ではキルができないが、キルクールは半分になる。\n停電中のみ移動速度が上昇するが名前が赤く表示される。"),
		Puppeteer:       NewAuInfo(imposter, "パペッティア", "キルした対象キャラクターに次に近づいたクルーをきるさせる。"),
		SerialKiller:    NewAuInfo(imposter, "シリアルキラー", "キルクールが短いインポスター。\nその代わり時間までに次のキルをしなければ自爆する。"),
		ShapeMaster:     NewAuInfo(imposter, "シェイプマスター", "変身後のクールダウンを無視して彩度返信ができるインポスター。\n通常は10秒しか変身できないが、設定により変身継続時間を変更できる。"),
		Sniper:          NewAuInfo(imposter, "スナイパー", "遠距離射撃が可能な役職。\nシェイプシフトした地点から解除した地点の延長線上にいる対象をキルできる。\nなお、斜線上にいるクルーには射撃音が聞こえる。\n弾丸を打ち切るまで通常のキルをすることはできない。\n「精密射撃モード」 が OFF の場合、扇状にキル範囲が展開される。\n「精密射撃モード」 が ON の場合、直線上にキル範囲が展開される。"),
		TimeThief:       NewAuInfo(imposter, "タイムシーフ", "プレイヤーをキルすると会話時間が減少していく。\nタイムシーフが追放またはキルされれば失われた会議時間が元に戻る。"),
		Vampire:         NewAuInfo(imposter, "ヴァンパイア", "キルしてから一定時間後、キル対象がテレポートせずに死亡する。\nただし、キル対象がベイトだった場合ヴァンパイアの能力は発動せず、普通のキルが発生する。\nキルが発動してから死亡するまでに会議が発生すると、会議が発生した瞬間に対象は死亡する。"),
		Warlock:         NewAuInfo(imposter, "ウォーロック", "ウォーロックが変身する前にキルすると相手に呪いがかかる。\n呪いをかけた後に変身すると、呪った人に一番近い人が死ぬようになる。\n呪いキルが成功するか、会議を挟むと呪いはリセットされる。"),
		Witch:           NewAuInfo(imposter, "ウィッチ", "キルボタンを押すと「キルモード」と「スペルモード」が入れ替わり、スペルモードの時にキルボタンを押すと、その対象に魔術を描けることができる。\n魔術をかけられたプレイヤーには会議で特殊なマークがつくようになり、その会議中に魔女を追放できなければ死亡する。"),
		Mafia:           NewAuInfo(imposter, "マフィア", "初期状態ではベントやサボタージュ、変身は可能だがキルすることはできない。\nマフィアではないインポスターが全員死亡するとマフィアもキルできるようになる。"),
		Madmate:         NewAuInfo(imposter, "マッドメイト", "キルもサボタージュもできないが、インポスターの味方をする役職。\nすべてのタスクを完了するとインポスターが誰かわかるようになる。"),
		MadGuardian:     NewAuInfo(imposter, "マッドガーディアン", "キルもサボタージュも出来ないが、インポスターの見方をする役職。\nすべてのタスクを完了するとバリアを入手することができる。\nバリア入手後はインポスターにキルされなくなる。"),
		SidekickMadmate: NewAuInfo(imposter, "サイドキックマッドメイト", "シェイプシフトの能力を持つ役職がシェイプシフトした際に作られる。"),

		// クルー陣営の情報登録
		Bait:           NewAuInfo(crew, "ベイト", "自分をキルしたプレイヤーに強制で自分の死体を通報させることができる。"),
		Dictator:       NewAuInfo(crew, "ディクテーター", "会議中に誰かが投票すると会議を強制終了させて投票先を釣ることができる。\n投票したタイミングでディクテーターは死亡する。"),
		Doctor:         NewAuInfo(crew, "ドクター", "プレイヤーの死因を知ることができて、遠隔でバイタルを見ることができる"),
		Lighter:        NewAuInfo(crew, "ライター", "タスクを完了させると自分の視界が広がるようになり、停電の視界減少の影響をうけなくなる"),
		Mayor:          NewAuInfo(crew, "メイヤー", "票を複数所持しており、まとめて1人のプレイヤーまたはスキップに投票できる。"),
		SabotageMaster: NewAuInfo(crew, "サボタージュマスター", "サボタージュを素早く治すことができる\nリアクター:片方を手をかざす/番号を入力するだけで直る\nO2:片方に番号を入力するだけで直る\n停電:1回のクリックですべて直る\nコミュニケーションサボタージュ:変更なし"),
		Sheriff:        NewAuInfo(crew, "シェリフ", "人外をキルすることができるが、クルーメイトをキルすると自爆してしまう。"),
		SpeedBooster:   NewAuInfo(crew, "スピードブースター", "タスクを完了させると生存しているランダムなプレイヤーの速度を上げる。"),
		Trapper:        NewAuInfo(crew, "トラッパー", "キルされると、キルした人を数秒間の間移動不可能にする。"),

		// 第3陣営の情報登録
		Arsonist:       NewAuInfo(third, "アーソニスト", "キルボタンを押して一定時間近くにいれば相手にオイルを塗れる。\n全員にオイルを塗ってベントに入ると起爆して単独勝利となる。"),
		Egoist:         NewAuInfo(third, "エゴイスト", "インポスターはエゴイストを認識している。\nエゴイストもインポスターを認識している。\nインポスターとエゴイストは斬りあうことができない。\nエゴイストは他のインポスターが全滅すると勝利する。\nエゴイストが勝利するとインポスターは敗北となる。"),
		Executioner:    NewAuInfo(third, "エクスキューショナー", "ターゲットに対してこちらからのみ視認できるダイヤマークがついている。\nこのダイヤマークがついている人を投票で追放すれば単独勝利となる。\nもし対象がキルされた場合は役職が変化する。\nまたターゲットがジェスターの場合追加勝利する。"),
		Jester:         NewAuInfo(third, "ジェスター", "投票で追放されたとき単独勝利となる。追放されずにゲームが終了するかキルされると敗北となる。"),
		Opportunist:    NewAuInfo(third, "オポチュニスト", "ゲーム終了時に生き残っていれば追加勝利。\nなおタスクはありません。"),
		SchrodingerCat: NewAuInfo(third, "シュレディンガーの猫", "デフォルトでは勝利条件を持たず、条件を満たすことで初めて勝利条件を持つようになる。\n\n- インポスターにキルされると、キルを防いでインポスター陣営になる\n- シェリフにキルされると、キルを防いでクルー陣営になる。\n- 第3陣営にキルされると、キルを防いで第3陣営になる。\n- 追放された場合、役職は変化せずそのまま死亡する。\n- ウォーロックの能力でキルされるとそのまま死亡する。\n- 自殺系キル(ヴァンパイア以外)でキルされると、そのまま死亡する。\n\nまたタスクは存在しない。"),
		Terrorist:      NewAuInfo(third, "テロリスト", "第3陣営であり、タスクがある。\nタスクをすべて完了させてからなんらかの要因で死亡すれば単独勝利となる。\n死因はキル/投票どちらでも単独勝利可能。\nテロリストのタスクはクルーメイトのタスク総数にカウントされず、テロリスト以外のクルー全員がタスクを終わらせればクルーメイトの勝利となる。"),
	}

	// selectMenuOption := make([]discordgo.SelectMenuOption, 0)
	selectMenuOption := CustomSelectMenuOption{}
	for role, info := range roles {
		selectMenuOption = append(selectMenuOption, *discordgox.NewSelectMenuOption(
			info.name,
			string(role),
			discordgox.SetSelectDescription(info.camp),
		))
	}

	selectMenuOptionImposter := selectMenuOption.extract(imposter)
	selectMenuOptionCrew := selectMenuOption.extract(crew)
	selectMenuOptionThird := selectMenuOption.extract(third)

	// ort.SliceStable(selectMenuOption, func(i, j int) bool { return selectMenuOption[i].Description < selectMenuOption[j].Description })
	sort.SliceStable(selectMenuOptionImposter, func(i, j int) bool { return selectMenuOption[i].Description < selectMenuOption[j].Description })

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
									customIDs[imposter],
									selectMenuOptionImposter,
									&min_value,
									len(selectMenuOptionImposter),
									imposter,
								),
							),
							discordgox.NewActionsRow(
								discordgox.SetMultiSelectMenu(
									customIDs[crew],
									selectMenuOptionCrew,
									&min_value,
									len(selectMenuOptionCrew),
									crew,
								),
							),
							discordgox.NewActionsRow(
								discordgox.SetMultiSelectMenu(
									customIDs[third],
									selectMenuOptionThird,
									&min_value,
									len(selectMenuOptionThird),
									third,
								),
							),
						),
					),
				),
				),
			)
			s.InteractionRespond(i.Interaction, cmdResponse)
		},
		customIDs[imposter],
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
