package parser

import (
	"kidstales/internal/model"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListPage(t *testing.T) {
	f, err := os.Open("testdata/list.html")
	require.NoError(t, err)

	values, err := new(BooksListPageParser).Parse(f)
	require.NoError(t, err)

	books := values["Books"].([]*model.Book)

	require.Equal(t, []*model.Book{
		{
			Name:       "Русалочка",
			Author:     "Disney",
			PageURL:    "https://skazkiwsem.fun/kniga-rusalochka-disney/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2020/06/img303-650x650.jpg",
		},
		{
			Name:       "Моана",
			Author:     "Disney",
			PageURL:    "https://skazkiwsem.fun/kniga-moana-disney/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2020/06/img108-650x650.jpg",
		},
		{
			Name:       "Три кота. Дружные котята",
			Author:     "Мои любимые сказки",
			PageURL:    "https://skazkiwsem.fun/detskaya-kniga-tri-kota-druzhnye-kotyata/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2019/07/1-5-650x650.jpg",
		},
		{
			Name:       "Три кота. Котята-помощники",
			Author:     "Мои любимые сказки",
			PageURL:    "https://skazkiwsem.fun/kniga-tri-kota-kotyata-pomoshhniki-moi-lyubimye-skazki/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2024/07/img053-650x650.jpg",
		},
		{
			Name:       "Фрижель и Флуффи. Затерянный остров. Выпуск 5",
			Author:     "Фрижель",
			PageURL:    "https://skazkiwsem.fun/kniga-frizhel-i-fluffi-zateryannyj-ostrov-vypusk-5-frizhel/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2024/07/img001-650x650.jpg",
		},
		{
			Name:       "Щенячий патруль. На суше и на море",
			Author:     "Nickelodeon",
			PageURL:    "https://skazkiwsem.fun/kniga-shhenyachij-patrul-na-sushe-i-na-more-nickelodeon/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2024/06/0-0001-650x650.jpg",
		},
		{
			Name:       "Щенячий патруль. В гостях у друзей",
			Author:     "Nickelodeon",
			PageURL:    "https://skazkiwsem.fun/kniga-shhenyachij-patrul-v-gostyax-u-druzej-nickelodeon/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2024/05/0-0001-650x650.jpg",
		},
		{
			Name:       "Сказка о потерянном времени",
			Author:     "Евгений Шварц",
			PageURL:    "https://skazkiwsem.fun/kniga-skazka-o-poteryannom-vremeni-evgenij-shvarc/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2024/03/0-0001-650x650.jpg",
		},
		{
			Name:       "Сказка о рыбаке и рыбке",
			Author:     "Пушкин Александр Сергеевич",
			PageURL:    "https://skazkiwsem.fun/kniga-skazka-o-rybake-i-rybke-pushkin-aleksandr-sergeevich/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2024/02/0-0001-650x650.jpg",
		},
		{
			Name:       "Волшебник страны Оз",
			Author:     "Лаймен Фрэнк Баум",
			PageURL:    "https://skazkiwsem.fun/kniga-volshebnik-strany-oz-lajmen-frenk-baum/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2024/01/0-0001-650x650.jpg",
		},
		{
			Name:       `Новогодний дневник агента "Сказочного патруля`,
			Author:     "Олег Рой",
			PageURL:    "https://skazkiwsem.fun/kniga-novogodnij-dnevnik-agenta-skazochnogo-patrulya-oleg-roj/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/12/0-0001-2-650x650.jpg",
		},
		{
			Name:       "Сказки-изобреталки от кота Потряскина",
			Author:     "Анатолий Гин",
			PageURL:    "https://skazkiwsem.fun/kniga-skazki-izobretalki-ot-kota-potryaskina-anatolij-gin/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/12/0-0001-1-650x650.jpg",
		},
		{
			Name:       "Может, Нуль не виноват?",
			Author:     "Ирина Токмакова",
			PageURL:    "https://skazkiwsem.fun/kniga-mozhet-nul-ne-vinovat-irina-tokmakova/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/12/0-0001-650x650.jpg",
		},
		{
			Name:       "Школа динозавров. Диплодок становится героем",
			Author:     "Пьер Жемм",
			PageURL:    "https://skazkiwsem.fun/kniga-shkola-dinozavrov-diplodok-stanovitsya-geroem-per-zhemm/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/11/0-0001-6-650x650.jpg",
		},
		{
			Name:       "Маленький Мук",
			Author:     "Вильгельм Гауф",
			PageURL:    "https://skazkiwsem.fun/kniga-malenkij-muk-vilgelm-gauf/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/11/0-0001-5-650x650.jpg",
		},
		{
			Name:       "Дюймовочка",
			Author:     "Ханс Кристиан Андерсен",
			PageURL:    "https://skazkiwsem.fun/kniga-dyujmovochka-xans-kristian-andersen/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/11/0-0001-4-650x650.jpg",
		},
		{
			Name:       "Пеппи маленький детектив. Куда исчезла гусеница?",
			Author:     "Рик ДеДонато",
			PageURL:    "https://skazkiwsem.fun/kniga-peppi-malenkij-detektiv-kuda-ischezla-gusenica-rik-dedonato/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/11/0-0001-3-650x650.jpg",
		},
		{
			Name:       "Пеппи маленький детектив. Кто похитил обед в лесу?",
			Author:     "Рик ДеДонато",
			PageURL:    "https://skazkiwsem.fun/kniga-peppi-malenkij-detektiv-kto-poxitil-obed-v-lesu-rik-dedonato/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/11/0-0001-2-650x650.jpg",
		},
		{
			Name:       "Не буду просить прощения",
			Author:     "Софья Прокофьева",
			PageURL:    "https://skazkiwsem.fun/kniga-ne-budu-prosit-proshheniya-sofya-prokofeva/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/11/0-0001-1-650x650.jpg",
		},
		{
			Name:       "Хомячок Фрош: Космические приключения во сне и наяву",
			Author:     "Елена Никитина",
			PageURL:    "https://skazkiwsem.fun/kniga-xomyachok-frosh-kosmicheskie-priklyucheniya-vo-sne-i-nayavu-elena-nikitina/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/11/0-0001-650x650.jpg",
		},
		{
			Name:       "Сказки Феи Дружного королевства",
			Author:     "Терентьева Ирина",
			PageURL:    "https://skazkiwsem.fun/kniga-skazki-fei-druzhnogo-korolevstva-terenteva-irina/",
			PictureURL: "https://skazkiwsem.fun/wp-content/uploads/2023/10/0-0001-7-650x650.jpg",
		},
	}, books)
}

func TestGetBookNameAuthor(t *testing.T) {
	const line = "Книга: «Хомячок Фрош: Космические приключения во сне и наяву» Елена Никитина"

	name, author, err := getBookNameAuthor(line)
	require.NoError(t, err)
	require.Equal(t, "Хомячок Фрош: Космические приключения во сне и наяву", name)
	require.Equal(t, "Елена Никитина", author)
}
