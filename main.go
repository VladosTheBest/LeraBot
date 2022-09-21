package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/Syfaro/telegram-bot-api"
)

var (
	row            []tgbotapi.InlineKeyboardButton
	weatherRequest = "https://api.openweathermap.org/data/2.5/weather?q=Kyiv,ua&units=metric&appid=77b78b59407e0f4951b75d05ed76fa97"
	smiles         = []string{
		"😁",
		"\U0001F90D",
		"😉",
		"😊",
		"😌",
		"😍",
		"😏",
		"😘",
		"😚",
		"😻",
		"",
	}
	buyedWish   = Wish{}
	vladID      = 517223834
	leraID      = 356131381
	compliments = []string{
		"*У тебя Очень прекрасные глазки, а то как они меняют свой цвет это просто Волшебно)*",
		"*Твои ножки просто прекрасны, когда вспоминаю как ты ими хвасталась, аж мурашки по коже идут!*",
		"*Когда мы держимся за руки и ты иногда дергаешься - это так мило, что хочется прижать тебя к себе и не отпускать!*",
		"*Мне с тобой очень комфортно, ты самый прекрасный человек которого я знаю!*",
		"*У тебя очень нежная кожа, когда я к тебе прикасаюсь, то просто не могу оторватся!*",
		"*Прикосновения твоих рук отдаются в моей душе...*",
		"*Никто в этом мире не сможет так бесподобно танцевать как ты... А твои прогибы это просто* :fire",
		"*Мне нравится, что со мною ты такая расслабленная и умиротворенная...*",
		"*Мне нравится, что нам легко быть вместе...*",
		"*Ни один книжный роман не сможет сравниться по яркости с историей НАШЕЙ любви...*",
		"*Мне нравится, как ты готовишь бесподобный кофе... И сырники...*",
		"*Обожаю наши ночные прогулки... И слушать твои увлекательные рассказы про историю всего, что нас окружает*",
		"*Моя жизнь после нашей встречи стала такой прекрасной...*",
		"*Когда ты рядом, мир вокруг наполняется счастьем...*",
		"*Меня сводит с ума сексуальность твоего животика...*",
		"*Мне нравится, что я хочу видеть рядом лишь тебя....*",
		"*Я просто счастлив, когда ты рядом...*",
		"*Фраза:\"To the world you may be one person, but to one person you may be the world. You are the world to me.\" О нас...*",
		"*Ты внимательная, нежная, кошачьи-сексуальная, изобретательная, страстная, ласковая, игривая любовница...*",
		"*Мне нравится просыпаться утром только с тобой...*",
		"*Мне нравится, что ты просто со мной, и никакие лишние слова и обещания не нужны...*",
		"*Мне нравятся ощущения, когда мы делаем друг-другу массаж... Особливо ⭐Зiрочкою⭐  *",
		"*Мне нравится, что ты оставила следы нежных кошачьих лапок в моей душе, зацепившись за нее коготками...*",
		"*После объятий с тобой моя одежда пахнет твоими духами... И когда я одеваю ее на следующий день - ощущение, что Ты рядом и обнимаешь меня...*",
		"*Я Обожаю носить тебя на руках...*",
		"*Мне нравится те ощущения, которые возникают, когда я глажу тебя...*",
		"*Мне нравится, как ты засыпаешь на моем плече...*",
		"*Обожаю твое \"Му-р-р\", когда ты обнимаешь меня...*",
		"*Я покорен твоим нежно-интригующе-таинственным взглядом...*",
		"*Я рад смог завоевать настоящее воплощение женственности - Тебя! *",
		"*Ты стала для меня единственной и неповторимой... Обожаю тебя!*",
		"*Не переставай улыбаться, твоя улыбка обворожительна!*",
		"*Ты очень вкусно пахнешь.*",
		"*Твой смех очаровывает!*",
		"*Я очень сильно ценю тебя.*",
		"*Ты прекрасный слушатель.*",
		"*Ты очень талантливая и креативная девушка*",
		"*Твой голос очень мелодичный и волнующий.*",
		"*У тебя безукоризненные манеры.*",
		"*Мне нравится гулять с тобой.*",
		"*С тобой я могу говорить обо всем.*",
		"*С тобой мне всегда очень комфортно.*",
		"*Я благодарен за то, что ты есть.*",
		"*Просто проводить время с умной, интеллигентной девушкой уже приносит незабываемые минуты радости*",
		"*Ты классная собеседница. Когда общаюсь с тобой, даже не замечаю, как пролетает время*",
		"*У меня мурашки от тебя*",
		"*Ты как магнит притягиваешь к себе*",
		"*Вот смотрю на тебя и понимаю, что в человеке может быть прекрасна душа, тело и мысли*",
		"*Ты важна для меня*",
		"*Ты даришь мне вдохновение в жизни, которое ранее меня не посещало. Классно, что я встретил тебя!*",
		"*Благодаря тебе я поверил в то, что действительно бывают родственные души*",
		"*Мне очень важно выслушать твое мнение*",
		"*Ты – уникальная! В тебя нельзя не влюбиться*",
		"*Всякий раз, когда мне грустно, мысль о тебе сразу поднимает настроение*",
		"*Ты – мой антидепрессант!*",
		"*У тебя огромный творческий потенциал! Ты можешь сотворить шедевр в каждой сфере жизни!*",
		"*Ты очень красива без макияжа, а твоя естественная красота видна сразу. Это большая редкость.*",
		"*У тебя такие тонкие и изящные пальчики, что мне хочется все время держать тебя за руку.*",
		"*Без ума от твоей походки. Каждый шаг, каждое движение бедрами дразнит, манит и чертовски заводит.*",
		"*Хочется целовать каждый сантиметр твоего тела и все время наслаждаться тобой.*",
		"*Обожаю твои изящные изгибы форм. Они будоражат мою кровь.*",
		"*Что я думаю о твоем внешнем виде? Скажем так, если бы существовал измеритель сексуальности, то рядом с тобой он бы уже давно взорвался от перегрузки.*",
		"*Ты — мой личный афродизиак. Волнуешь, возбуждаешь, отключаешь голову, играешь со мной.*",
		"*Я не хочу видеть рядом кого-либо, кроме тебя. Ты — моя единственная, и я твердо в этом уверен.*",
		"*Ты — это самый лучший выбор, который я сделал за свою жизнь. Я рад, что судьба свела нас вместе.*",
		"*Спасибо за то, что делаешь меня счастливым, я еще никогда не испытывал такой радости и блаженства.*",
		"*Я понятия не имею, есть ли там наверху кто-то, кто предназначает людей друг для друга. Но если предположить, что это так, то ты точно моя половина.*",
		"*Я не встречал таких, как ты. Ты действительно особенная и неповторимая.*",
		"*Я хочу, чтобы то, что есть между нами, длилось очень и очень долго. Нет, не так. Хочу, чтобы это было навсегда и именно с тобой.*",
		"*Мне трудно выразить свои чувства к тебе, потому что таких сильных эмоций я никогда не испытывал.*",
		"*С тех пор, как ты появилась в моей жизни, мне стало очень трудно замечать кого-то еще. А сейчас, кроме тебя, в моем мире вообще никого не существует.*",
		"*Ну вот что ты наделала? Работать не могу, развлекаться не могу, с друзьями тоже сам не свой. Почему? Да потому что только о тебе и думаю.*",
		"*Ты такая хрупкая и нежная, что хочется постоянно носить на руках и оберегать от всего плохого.*",
		"*Не хочу отпускать тебя туда одну, вдруг кто-то украдет такое сокровище!*",
		"*Я хочу засыпать и просыпаться с тобой, видеть твою сонную мордашку каждое утро и будить тебя поцелуями.*",
		"*Я знаю, что ты самостоятельная взрослая девушка, но ничего не могу с собой поделать. Хочу заботиться о тебе все время.*",
		"*Каждый раз, когда ты рядом, у меня ощущение, что весь мир становится более красочным и удивительным.*",
		"*Я не умею делать сальто, а вот мое сердце научилось, когда увидело тебя.*",
		"*Ну ты хотя бы огнетушитель с собой носила, раз так выглядишь.*",
		"*Ты вот просто улыбнулась, а у меня капитуляция мозга произошла.*",
		"*Обожаю, когда ты начинаешь неосознанно покусывать (облизывать) губы. Это так сексуально.*",
		"*Потягиваешься, как кошечка. Так и хочется погладить.*",
		"*Благодаря тебе мне хочется быть лучше, сильнее, успешнее.*",
		"*Я действительно счастлив с тобой. И это полностью твоя заслуга.*",
		"*В твоей прекрасной груди бьется жаркое сердце*",
		"*Ты лучше, чем сказочная принцесса*",
		"*Ты красивее, чем можешь представить*",
		"*Ничего в себе не меняй. Ты великолепна*",
		"*Ты всегда будешь поводом моей улыбки*",
		"*Пожалуйста, не переставай улыбаться. Твоя улыбка, даже самый морозный день, сделает для меня теплым.*",
		"*Ты всегда потрясающе выглядишь! Спинным мозгом и затылком чувствую на себе завистливые взгляды мужчин, когда мы с тобой вместе гуляем.*",
		"*Ты красивая и умная, теперь я понял, что мой идеал существует!*",
		"*Знаешь, мне не хватит жизни, чтобы сделать подходящий комплимент тебе, потому что с каждым днем ты становишься все прекраснее и прекраснее*",
		"*Ты – мечта! Ты – наважденье! Ты – воплощенье красоты!*",
		"*Ты ослепляешь и сводишь с ума!\nТы на всем свете такая одна!*",
		"*В твоих глазах – огонь и сила!\nТы – девушка моей мечты!*",
		"*Твои глаза! Хочу в них утонуть...\nКогда ты смотришь на меня,\nЯ забываю, что умею плавать…*",
		"*Твои волосы обалденно пахнут мне этот запах будет теперь повсюду мерещиться!*",
		"*Ты ворвалась в мое сердце, с ноги! Но какой же красивой...*",
		"*Ты самый прекрасный и самый красивый цветок на свете! Я без ума от тебя!*",
		"*Как же я обожаю твои Прекрасные, Густые, Шелковистые волосы!*",
		"*На земле много чудес, но есть одно, самое удивительное и прекрасное, – это ты!*",
		"*С тобой лишь хочу встречать каждый рассвет!*",
		"*Ты для меня – счастье и свет\n Все самое ясное, все самое нежное*",
		"*Какое счастье просто знать, что в любой момент смогу тебя обнять*",
		"*В любом наряде, ты всегда бесподобно красива.*",
		"*В чем смысл жизни и никто до сих пор не нашел ответа. А я все знаю, для меня и ответ, и смысл жизни - это ты!*",
		"*А вообще-то знаешь, Я просто тебя люблю*",
		"*Твои губки это самый прекрасный десерт, который я всегда хочу!*",
		"*Мне нравиться, что с тобой можно говорить о чем угодно и в то же время ниочем, с тобой приятно даже молчать, обожаю тебя!*",
		"*Мне нравиться твой вкус, то как ты одеваешься и даже то, что у тебя частенько развязываются шнурки, это безумно мило!*",
		"*Амбициозная и божественная, великолепная и гордая, добрая и жизнерадостная, замечательная и искренняя. Можно до бесконечности продолжать делать тебе комплименты, ведь в каждой букве алфавита скрыто столько твоих достоинств!*",
		"*В тебе есть та самая изюминка, которая заставляет мое сердце бится чаще!*",
		"*Обожаю слушать твои расказы про историю и другие интересные факты, которые ты мне рассказываешь, когда мы с тобой гуляем!*",
		"*Ты, такая бесконечно нежная, изысканно утонченная и ошеломляюще изумительная!*",
		"*Спасибо тебе, что сразила меня!*",
		"*Скажу только одно, я в восторге от времени проведённого с тобой*",
		"*Ты мой Ангел - не улетай от меня далеко!*",
		"*Не видеть тебя это пытка*",
		"*Ты для меня как сказка, только пусть страницы этой книги будут длинною в жизнь!*",
		"*Видел тебя в восхитительном сне и расстроился от того, что пришла пора проснуться*",
		"*Когда ты рядом, моё сердце стучит так громко, что его слышно невооруженным ухом*",
		"*Когда Стараюсь сосредоточиться и работать. Не получается, все мысли о тебе и совсем не рабочие*",
		"*Ты как конфетка, покрытая глазурью из нежности и ласки!*",
		"*Как тяжело, когда тебя нет рядом, как не хватает твоего тепла*",
		"*Ты улыбаешься, и начинается проливной дождь из положительных эмоций и ярких впечатлений!*",
		"*Губы есть у всех, но твои жгучие губки вне конкуренции!*",
		"*Эта зима будет самой теплой в моей жизни, ведь эта зима будет рядом с тобой!*",
		"*Когда я с тобой, то хочу, чтобы время остановилось. Но оно, увы, летит слишком быстро*",
		"**",
	}
	names = []string{
		"Валерия",
		"Лерочка",
		"Валерочка",
		"🔥Лера Кипяток🔥",
		"Лера",
		"Леруся",
	}
	hbMessage = "Ну и конечно так как у тебя сегодня\U0001F973\n" +
		"*✨✨День Рождения✨✨*\n" +
		"Влад написал небольшое поздравление💌\n\n" +
		"*Лерочка, я хочу поздравить тебя с этим прекрасным днем " +
		"который подарил на свет восьмое чудо мира - ТЕБЯ!!!\n" +
		"Я очень рад, что познакомился с тобой, никогда не думал, что " +
		"найду человека красивого внешне и прекрасного внутри! " +
		"Мне безумно нравится проводить с тобой время, наши ночные прогулки по Киеву\U0001F90D или " +
		"просто лежать в обнимку с тобой - это дает нереальный прилив спокойствия, сил и счастья! " +
		"С тобой всегда есть о чем поговорить, причем на любую тему!) " +
		"Когда я держу тебя за руку и ты иногда дергаешься, это просто нереально мило😊" +
		" и не хочеться отпускать тебя никогда... Я хочу пожелать тебе достижения своих целей, " +
		"своими силами! Ведь самое приятное когда чего-то достигаешь - понимать, что ты этого добилась сама! " +
		"Исполнения всех твоих желаний и мечтаний! Ну и как истинному ипохондрику побольше здоровья и позитва!)\n" +
		"И еще хочу тебе сказать, что ты всегда можешь на меня положится, чтобы не случилось, хорошее или плохое, я " +
		"всегда тебя выслушаю, пойму и постараюсь помочь, если это будет в моих слиах!\n\n" +
		"Спасибо тебе, что ты есть!❤️❤️❤️*"
	isSend   = true // если false отправит поздравление!
	newsSend = false

	leraPath        = "casinoBot/data/lera.json"
	wishPath        = "casinoBot/data/wish.json"
	buyedWishesPath = "casinoBot/data/buyedWishes.json"

	wishList = []Wish{{
		Id:          0,
		Description: "Добавить свое желание: Опиши свое желание которое Влад может исполнить а я сделаю ему оценку и добавлю в список",
		isForever:   true,
		isDone:      false,
		Price:       2,
	},
		{
			Id:          1,
			Description: "Выбор места для свидания \n( формат и все нюансы можно будет описать в заполнении анкеты желания )",
			isForever:   false,
			isDone:      false,
			Price:       5,
		},
		{
			Id:          2,
			Description: "Называй меня `Ваш вариант` целый день",
			isForever:   false,
			isDone:      false,
			Price:       2,
		},
		{
			Id:          3,
			Description: "Называй меня `Ваш вариант` целую неделю",
			isForever:   false,
			isDone:      false,
			Price:       7,
		}}
)

type Horo struct {
	Libra struct {
		Text       string `xml:",chardata"`
		Yesterday  string `xml:"yesterday"`
		Today      string `xml:"today"`
		Tomorrow   string `xml:"tomorrow"`
		Tomorrow02 string `xml:"tomorrow02"`
	} `xml:"libra"`
}

type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed int `json:"speed"`
		Deg   int `json:"deg"`
	} `json:"wind"`
	Rain struct {
		H float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Lera struct {
	Name   string
	Age    int
	Points int
	ID     int64
	Page   int
}

type Wish struct {
	Id              int
	Description     string // описание
	isForever       bool   // Пропадает ли при покупке из магазина желаний
	Price           int    // Цена
	LeraDescription string // Описание Леры
	isDone          bool   //Выполнено ли желание
}

func main() {
	StartBot()
}

func StartBot() {
	//Создаем бота
	bot, err := tgbotapi.NewBotAPI("2049072827:AAHdyBdM7ccGxV4ESBh7Q5X66lJjltwNDL0")
	if err != nil {
		fmt.Println(err)
	}
	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)
	rand.Seed(time.Now().UnixNano())
	go sendWeatherAndAstro(8, *bot)

	for update := range updates {
		//Проверяем что от пользователья пришло именно текстовое сообщение
		if update.Message != nil {
			if update.Message.Chat.ID == 356131381 || update.Message.Chat.ID == 517223834 {
				if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" { // text messages
					update.Message.Text = strings.ToLower(update.Message.Text)
					fmt.Println(update.Message.From, time.Now().String())

					switch update.Message.Text {
					case "/start":
						if isSend == true {
							hbMessage = ""
						}
						//Отправлем сообщение
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Здравствуй \U0001F90DВалерия❤\nМеня создал Влад👉 @vladyur4ik\n"+
							"Я буду говорить тебе😏\n🌟🌟🌟*КОМПЛИМЕНТЫ*🌟🌟🌟 \n"+
							"Когда он занят😔 \nИли когда сам этого захочу 😉\n"+
							"Долгое время ничего не говорит?!😡\nНапиши 👊/kick👊 "+
							"\nИли нажми кнопку 👊Пнуть Влада👊\n"+
							"Я объясню этому кожанному!🙅🙅🙅\nТак делать нельзя!!!😤😤😤\n"+
							"А еще, так как твоя учеба связанна с 🥁\n"+
							"🤓*ИСТОРИЕЙ*🤓\n"+
							"Я помогу тебе в этом!😊😊😊\n"+
							"Я найду информацию про любую дату!😎\n"+
							"Просто напиши «* факт 01/05 *»✍️✍️✍️\n"+
							"📅*  01 - Месяц / 05 - День  * 📅\n"+
							"📠Я пришлю информацию о дате!📠\n"+
							"Правда на английском🇬🇧🇬🇧🇬🇧\n"+
							"Но ты же у нас *Умная девочка*👩 \n"+
							"А еще я буду по утрам присылать тебе📤\n"+
							"Прогноз погоды специально для тебя!☔️\n"+
							"И сегодняшний гороскоп для  весов♎️\n"+hbMessage)
						msg.ParseMode = "markdown"
						msg.ReplyMarkup = getStandartKeyboard()
						bot.Send(msg)
						isSend = true
					case "погода":
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, getWeatherMSG())
						bot.Send(msg)
					case CheckSubstring(update.Message.Text, "fuckt"):
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, getFact(strings.Split(update.Message.Text, " ")[1]))
						bot.Send(msg)
					case CheckSubstring(update.Message.Text, "факт"):
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, getFact(strings.Split(update.Message.Text, " ")[1]))
						bot.Send(msg)
					case CheckSubstring(update.Message.Text, "комплимент"):
						text := strings.ReplaceAll(update.Message.Text, "комплимент", "")
						rndSmile := rand.Intn(len(smiles))
						rndName := rand.Intn(len(names))
						msg := tgbotapi.NewMessage(int64(leraID), fmt.Sprintf("*%v %v %v\n\n✨%v %v*", smiles[1], names[rndName], smiles[1], text, smiles[rndSmile]))
						msg.ParseMode = "markdown"
						msg.ReplyMarkup = getStandartKeyboard()
						bot.Send(msg)
					case "/add":
						lera := GetLeraData("casinoBot/data/lera.json")
						lera.Points++
						WriteJsonData(lera, "casinoBot/data/lera.json")
						msg := tgbotapi.NewMessage(int64(leraID), "🖤             ***Валерия***          🖤\n" +
							"_Вам начислили_ ***1*** 💞💞💞 _lovePoints\n" +
							"Может стоит заглянуть в магазин?\U0001F970_")
						msg.ParseMode = "markdown"
						msg.ReplyMarkup = getStandartKeyboard()
						bot.Send(msg)
					case "/minus":
						lera := Lera(GetLeraData("casinoBot/data/lera.json"))
						lera.Points--
						WriteJsonData(lera, "casinoBot/data/lera.json")
						msg := tgbotapi.NewMessage(lera.ID, LeraToString(lera))
						bot.Send(msg)
					case "/savewish":
						WriteJsonData(wishList, "casinoBot/data/wish.json")
					case CheckSubstring(update.Message.Text, "описание"):
						wished := GetWishData(buyedWishesPath)
						text := strings.ReplaceAll(update.Message.Text, "описание", "")
						buyedWish.LeraDescription = text
						wished = append(wished, buyedWish)
						WriteJsonData(wished, buyedWishesPath)
						msg := tgbotapi.NewMessage(int64(leraID), "***Описание Добавленно!!!***\n\n"+
							"Теперь желание добавилось в список купленных и Влад обязан начать воплощать его в жинь!!!")
						msg.ParseMode = "markdown"
						msg.ReplyMarkup = getStandartKeyboard()
						bot.Send(msg)
						msg.ChatID = int64(vladID)
						msg.Text = "Влад, там Лера купила Желание, нужно его какбы исполнить!!! Так что давай, Шевели Булками!!!"
						bot.Send(msg)

					default:
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "💢                \U0001F90DЛера\U0001F90D                 💢"+
							"\nЯ еще не умею отвечать на такие сообщения...😓😓😓")
						bot.Send(msg)
					}
				}
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Простите, но я выполняю приказы только Леры и Влада...\n"+
					"Найдите себе другого бота...")
				bot.Send(msg)
				return
			}
		}
		if update.CallbackQuery != nil {
			fmt.Println(update.CallbackQuery, time.Now().String())
			command := update.CallbackQuery.Data
			switch command {
			case "kick":
				msg := tgbotapi.NewMessage(int64(vladID), "💢💢💢💢💢*Ало Псина!!!*💢💢💢💢💢"+
					"\n✏     *Напиши Лере, она скучает!!!*    ✏")
				msg.ParseMode = "markdown"
				bot.Send(msg)

				msg.Text = "Я его пнул и проконтролирую,чтобы он тебе написал😉😉😉"
				msg.ParseMode = "markdown"
				msg.ReplyMarkup = getStandartKeyboard()
				msg.ChatID = int64(update.CallbackQuery.From.ID)
				bot.Send(msg)
			case "compliment":
				rnd := rand.Intn(len(compliments))
				rndSmile := rand.Intn(len(smiles))
				rndName := rand.Intn(len(names))
				msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), fmt.Sprintf("*%v %v %v\n\n✨%v %v*", smiles[1], names[rndName], smiles[1], compliments[rnd], smiles[rndSmile]))
				msg.ParseMode = "markdown"
				msg.ReplyMarkup = getStandartKeyboard()
				bot.Send(msg)
			case "weather":
				msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), getWeatherMSG())
				msg.ReplyMarkup = getStandartKeyboard()
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "horo":
				msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), getAstro())
				msg.ReplyMarkup = getStandartKeyboard()
				msg.ParseMode = "markdown"
				bot.Send(msg)
			case "photo":
				files, err := ioutil.ReadDir("D:/Golang program/Project1/tgbot/photos")
				if err != nil {
					log.Fatal(err)
				}
				rnd := rand.Intn(len(files))
				photoBytes, err := ioutil.ReadFile("tgbot/photos/" + files[rnd].Name())
				if err != nil {
					return
				}
				photoFileBytes := tgbotapi.FileBytes{
					Name:  "picture",
					Bytes: photoBytes,
				}
				msg := tgbotapi.NewPhotoUpload(int64(update.CallbackQuery.From.ID), photoFileBytes)

				msg.ReplyMarkup = getStandartKeyboard()
				bot.Send(msg)

				//SHOP - - - - - - -  - - - - - - - -
			case "next":
				lera := GetLeraData(leraPath)
				lera.Page++
				WriteJsonData(lera, leraPath)
				wish := GetWish(lera)
				kb := tgbotapi.InlineKeyboardMarkup{}
				kb = getShopKeyboard()
				edit := tgbotapi.EditMessageTextConfig{
					BaseEdit:              tgbotapi.BaseEdit{
						ChatID: int64(update.CallbackQuery.From.ID),
						ChannelUsername: "",
						MessageID:       update.CallbackQuery.Message.MessageID,
						InlineMessageID: "",
						ReplyMarkup:     &kb,
					},
					Text:                  wishToString(wish, lera),
					ParseMode:             "markdown",
					DisableWebPagePreview: false,
				}
				bot.Send(edit)
			case "prev":
				lera := GetLeraData(leraPath)
				lera.Page--
				WriteJsonData(lera, leraPath)
				wish := GetWish(lera)
				kb := tgbotapi.InlineKeyboardMarkup{}
				kb = getShopKeyboard()
				edit := tgbotapi.EditMessageTextConfig{
					BaseEdit:              tgbotapi.BaseEdit{
						ChatID: int64(update.CallbackQuery.From.ID),
						ChannelUsername: "",
						MessageID:       update.CallbackQuery.Message.MessageID,
						InlineMessageID: "",
						ReplyMarkup:     &kb,
					},
					Text:                  wishToString(wish, lera),
					ParseMode:             "markdown",
					DisableWebPagePreview: false,
				}
				bot.Send(edit)
			case "back":
				kb := tgbotapi.InlineKeyboardMarkup{}
				kb = getStandartKeyboard()
				edit := tgbotapi.EditMessageTextConfig{
					BaseEdit:              tgbotapi.BaseEdit{
						ChatID: int64(update.CallbackQuery.From.ID),
						ChannelUsername: "",
						MessageID:       update.CallbackQuery.Message.MessageID,
						InlineMessageID: "",
						ReplyMarkup:     &kb,
					},
					Text:                  "***🖤              Лерочка              🖤 \n\nТы в главном меню!***\nВыбери действие и нажми кнопочку😘 ",
					ParseMode:             "markdown",
					DisableWebPagePreview: false,
				}
				bot.Send(edit)
			case "shop":
				kb := tgbotapi.InlineKeyboardMarkup{}
				kb = getShopKeyboard()
				edit := tgbotapi.EditMessageTextConfig{
					BaseEdit:              tgbotapi.BaseEdit{
						ChatID: int64(update.CallbackQuery.From.ID),
						ChannelUsername: "",
						MessageID:       update.CallbackQuery.Message.MessageID,
						InlineMessageID: "",
						ReplyMarkup:     &kb,
					},
					Text:                  "Ты зашла в магазин ***Желаний***\nЛюблю тебя и хороших покупок\U0001F970 \n"+wishToString(GetWish(GetLeraData(leraPath)), GetLeraData(leraPath)),
					ParseMode:             "markdown",
					DisableWebPagePreview: false,
				}
				bot.Send(edit)
			case "rules":
				msg :=  "***Лерочка, приветствую тебя в магазине желаний!!!***\n"+
					"🗯🗯                   ***Описание***                  🗯🗯\n\n"+
					"🔊🔊                          ***1***                       🔊🔊\n"+
					"_Здесь ты сможешь выбрать и купить уже существующие желания или же создать свои и потом его купить._\n"+
					"🔊🔊                          ***2***                       🔊🔊\n"+
					"_Все эти желания Влад обязан будет выполнить, список желаний которые он выполнил и которые уже в процессе можно будет посмотреть в моем меню, чтобы"+
					" следить за результатом ваших `_***Торговых отношений***_` ну и для интереса и статистики)))_\n"+
					"🔊🔊                          ***3***                       🔊🔊\n"+
					"_Покупка будет за LovePoints - это очки которые будут тебе начисляться за различные действия и поступки по "+
					"отношению к Владу или же просто в отношениях, за что конкретно тебе начислили очки ты знать не будешь, но будешь получать уведомление,"+
					" когда он их будет начислять!!!!_\n"+
					"🔊🔊                          ***4***                       🔊🔊\n"+
					"_Это создаеться исключительно как какой-то новый элемент который будет давать новые возможности и эмоции в наших отношениях, Если есть пожелания "+
					"по улучшению или изменению правил, свяжись с Владом_ @vladyur4ik \n"+
					"***И хороших покупок!!!***"
				kb := tgbotapi.InlineKeyboardMarkup{}
				kb = getShopKeyboard()
				edit := tgbotapi.EditMessageTextConfig{
					BaseEdit:              tgbotapi.BaseEdit{
						ChatID: int64(update.CallbackQuery.From.ID),
						ChannelUsername: "",
						MessageID:       update.CallbackQuery.Message.MessageID,
						InlineMessageID: "",
						ReplyMarkup:     &kb,
					},
					Text:                  msg,
					ParseMode:             "markdown",
					DisableWebPagePreview: false,
				}
				bot.Send(edit)
			case "buy":
				lera := GetLeraData(leraPath)
				wish := GetWishData(wishPath)
				if lera.Points < wish[lera.Page].Price {
					fmt.Println(lera.Points,wish[lera.Page].Price)
					bot.Send(tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "Прости, но у тебя недостаточно LovePoints, попробуй собрать немного общаясь с Владом)"))
					return
				}
				lera.Points -= GetWish(lera).Price
				msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "Ты купила желание: "+wish[lera.Page].Description+"\n"+LeraToString(lera)+"\n"+
					"Чтобы добавить описание или уточнить '***Только что купленное желание***' напиши '***Описание *** _ Твой текст описания_'")
				WriteJsonData(lera, leraPath)
				buyedWish = wish[lera.Page]
				wish = deleteElementFromArray(wish, lera.Page)
				WriteJsonData(wish, wishPath)
				msg.ReplyMarkup = getStandartKeyboard()
				msg.ParseMode = "markdown"
				bot.Send(msg)

			case "boughtWishes":
				wishes := GetWishData(buyedWishesPath)
				if len(wishes) == 0 || wishes == nil {
					msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), "Валерия вы еще не купили ни одного желания!!!")
					bot.Send(msg)
					return
				}
				msg := tgbotapi.NewMessage(int64(update.CallbackQuery.From.ID), boughtWishesToString(wishes))
				msg.ParseMode = "markdown"
				bot.Send(msg)
			}
		}
	}
}

func CheckSubstring(str string, sub string) string {
	if len(str) == 0 {
		return ""
	}
	if strings.Contains(str, sub) {
		return str
	}
	return ""
}

func getFact(date string) string {
	resp, _ := http.Get(fmt.Sprintf("http://numbersapi.com/%v/date", date))
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return string(body)
}

func getAstro() string {
	data := new(Horo)
	resp, _ := http.Get("https://ignio.com/r/export/utf/xml/daily/com.xml")
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	xml.Unmarshal(body, &data)
	return "\U0001F90D *Валерия* \U0001F90D \n\n💫Твое предсказание на сегодня!💫\n" + data.Libra.Today + "✨🧙"
}

func getWeatherMSG() string {
	data := new(Weather)
	resp, _ := http.Get(weatherRequest)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	json.Unmarshal(body, &data)
	weatherIcon := ""
	switch strings.ToLower(data.Weather[0].Description) {
	case "clear sky":
		weatherIcon = "☀"
	case "few clouds":
		weatherIcon = "⛅"
	case "scattered clouds":
		weatherIcon = "☁"
	case "broken clouds":
		weatherIcon = "☁️☁️"
	case "shower rain":
		weatherIcon = "🌦"
	case "rain":
		weatherIcon = "🌧"
	case "thunderstorm":
		weatherIcon = "⛈"
	case "snow":
		weatherIcon = "🌨☃️"
	case "mist":
		weatherIcon = "🧖‍♂️"
	}

	return fmt.Sprintf("Прогноз для 😘*Лерочки*❤ \n\n"+
		"***Итак, за окном у нас %v %v \nТемпература сейчас %v°C☕️\n"+
		"Ты как истинный ипохондрик должна знать👩‍⚕️💊\nДавление снаружи %v🌡 мм рт. ст.\n"+
		"Ну и конечно влажность снаружи около %v %v💧💧 ***", data.Weather[0].Main, weatherIcon, data.Main.Temp, data.Main.Pressure, data.Main.Humidity, "%")
}

func getStandartKeyboard() tgbotapi.InlineKeyboardMarkup {
	kb := tgbotapi.InlineKeyboardMarkup{}
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("😌Хочу комплимент😌", "compliment"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("👊Пнуть Влада👊", "kick"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("🛍Магазин🛍", "shop"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("❄Что по погоде?❄", "weather"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("🌟Предсказание на Сегодня?🌟", "horo"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("🖤Memories🖤", "photo"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	return kb
}

func getShopKeyboard() tgbotapi.InlineKeyboardMarkup {
	kb := tgbotapi.InlineKeyboardMarkup{}
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("👈Предидущее👈", "prev"))
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("👉Следущее👉", "next"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("💰Купить💰", "buy"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("😈Купленные Желания😈", "boughtWishes"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("🤔Описание🤔", "rules"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	row = []tgbotapi.InlineKeyboardButton{}
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("👣Назад👣", "back"))
	kb.InlineKeyboard = append(kb.InlineKeyboard, row)
	return kb
}

func sendWeatherAndAstro(weather int, bot tgbotapi.BotAPI) {
	for {
		if time.Now().Hour() == weather && !newsSend {
			newsSend = true
			msg := tgbotapi.NewMessage(int64(356131381), getWeatherMSG())
			msg.ParseMode = "markdown"
			bot.Send(msg)
			msg.Text = getAstro()
			msg.ReplyMarkup = getStandartKeyboard()
			bot.Send(msg)
		} else if time.Now().Hour() != weather && newsSend {
			newsSend = false
		}
	}
}

func WriteJsonData(data interface{}, path string) {
	file, _ := json.MarshalIndent(data, "", "")
	_ = ioutil.WriteFile(path, file, 0644)
	fmt.Println("Writed File", data)
}

func GetLeraData(path string) Lera {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error when read file, err", err)
		return Lera{}
	}
	lera := Lera{}
	_ = json.Unmarshal(file, &lera)
	fmt.Println("Readed File", lera)
	return lera
}

func GetWishData(path string) []Wish {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error when read file, err", err)
		return []Wish{}
	}
	wish := []Wish{}
	_ = json.Unmarshal(file, &wish)
	fmt.Println("Readed File", wish)
	return wish
}

func LeraToString(lera Lera) string {
	return fmt.Sprintf("Name : %v \n"+
		"Age : %v \n"+
		"Points : %v \n", lera.Name, lera.Age, lera.Points)
}

func GetWish(lera Lera) Wish {
	wishes := GetWishData("casinoBot/data/wish.json")
	if lera.Page < 0 {
		lera.Page = len(wishes) - 1
	} else if lera.Page > len(wishes)-1 {
		lera.Page = 0
	}
	WriteJsonData(lera, leraPath)
	fmt.Println("WISH: ", wishes[lera.Page])
	return wishes[lera.Page]
}

func boughtWishesToString(wishes []Wish) string {
	text := ""
	for i, wish := range wishes {
		text += fmt.Sprintf("***ID:***                                    _%v_ \n" +
			"***Описание:***    _%v_ \n" +
			"***Описание от Валерии:***   _%v_ \n\n" +
			"***Стоимость:*** 💞_%v_💞 \n" +
			"***Статус:***          _%v_ \n" +
			"➖➖➖➖➖➖➖➖➖➖➖➖\n", i+1, wish.Description, wish.LeraDescription, wish.Price, wishState(wish.isDone))
	}
	return text
}

func wishToString(wish Wish, lera Lera) string {
	wishes := GetWishData(wishPath)
	if lera.Page < 0 {
		lera.Page = len(wishes) - 1
	} else if lera.Page > len(wishes)-1 {
		lera.Page = 0
	}
	return fmt.Sprintf("⬛️              ***Номер Желания: %v***              ⬛️\n"+
		"***Описание:*** _%v_ \n\n"+
		"***Цена желания:***💞 _%v_ 💞 \n"+
		"***У вас на счету:***🖤 _%v_ 🖤\n"+
		"⬛️                   ***Страница: %v/%v***                   ⬛️\n", wish.Id+1, wish.Description, wish.Price, lera.Points, lera.Page+1, len(wishes))
}

func deleteElementFromArray(wishes []Wish, i int) []Wish {

	// Удалить элемент по индексу i из a.

	// 1. Выполнить сдвиг a[i+1:] влево на один индекс.
	copy(wishes[i:], wishes[i+1:])

	// 2. Удалить последний элемент (записать нулевое значение).
	wishes[len(wishes)-1] = Wish{}

	// 3. Усечь срез.
	wishes = wishes[:len(wishes)-1]

	fmt.Println(wishes) // [A B D E]
	return wishes
}

func wishState (isDone bool) string {
	if isDone {
	return "Желание исполнено✅"}
	return "Желание еще в процессе исполнения❌"
}