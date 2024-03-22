package mysql

import (
	"log"
	"time"

	"github.com/Syahreza-Ferdian/heal-in/entity"
	"github.com/Syahreza-Ferdian/heal-in/pkg/bcrypt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func questionSeeder(db *gorm.DB) error {
	questions := []entity.JournalingQuestion{
		{
			Question: "3 hal yang aku alami hari ini",
		},
		{
			Question: "3 hal yang membuatku bersyukur hari ini",
		},
		{
			Question: "3 hal positif yang aku terima dari orang lain hari ini",
		},
		{
			Question: "3 harapan yang aku inginkan akan terjadi",
		},
		{
			Question: "3 hal yang ingin kulakukan agar besok bisa lebih baik lagi",
		},
		{
			Question: "Pesan yang ingin kusampaikan untuk dirimu sendiri dan orang lain yang sudah berbuat baik kepadaku",
		},
	}

	err := db.Create(&questions).Error

	if err != nil {
		return err
	}

	return nil
}

func moodSeeder(db *gorm.DB) error {
	moods := []entity.JournalingMood{
		{
			Mood: "Happy",
		},
		{
			Mood: "Sad",
		},
		{
			Mood: "Angry",
		},
		{
			Mood: "Neutral",
		},
	}

	err := db.Create(&moods).Error

	if err != nil {
		return err
	}

	return nil

}

func userSeeder(db *gorm.DB, bcrypt bcrypt.BcryptInterface) error {
	pass, err := bcrypt.HashPassword("syahreza")

	if err != nil {
		return err
	}

	user := []entity.User{
		{
			ID:               uuid.New(),
			Name:             "Syahreza",
			Email:            "superadmin@admin.com",
			Password:         pass,
			IsEmailVerified:  true,
			VerificationCode: "",
			IsSubscribed:     true,
		},
		{
			ID:               uuid.New(),
			Name:             "Also Syahreza",
			Email:            "me@syahreza.com",
			Password:         pass,
			IsEmailVerified:  true,
			VerificationCode: "",
			IsSubscribed:     false,
		},
	}

	err = db.Create(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func afirmationWordSeeder(db *gorm.DB) error {
	words := []entity.AfirmationWord{
		{
			ID:     uuid.New(),
			MoodID: 1,
			Word:   "Afirmation Word Happy 1",
		},
		{
			ID:     uuid.New(),
			MoodID: 1,
			Word:   "Afirmation Word Happy 2",
		},
		{
			ID:     uuid.New(),
			MoodID: 1,
			Word:   "Afirmation Word Happy 3",
		},
		{
			ID:     uuid.New(),
			MoodID: 1,
			Word:   "Afirmation Word Happy 4",
		},
		{
			ID:     uuid.New(),
			MoodID: 2,
			Word:   "Afirmation Word Sad 1",
		},
		{
			ID:     uuid.New(),
			MoodID: 2,
			Word:   "Afirmation Word Happy 2",
		},
		{
			ID:     uuid.New(),
			MoodID: 2,
			Word:   "Afirmation Word Happy 3",
		},
		{
			ID:     uuid.New(),
			MoodID: 2,
			Word:   "Afirmation Word Happy 4",
		},
		{
			ID:     uuid.New(),
			MoodID: 3,
			Word:   "Afirmation Word Angry 1",
		},
		{
			ID:     uuid.New(),
			MoodID: 3,
			Word:   "Afirmation Word Angry 2",
		},
		{
			ID:     uuid.New(),
			MoodID: 3,
			Word:   "Afirmation Word Angry 3",
		},
		{
			ID:     uuid.New(),
			MoodID: 4,
			Word:   "Afirmation Word Neutral 1",
		},
		{
			ID:     uuid.New(),
			MoodID: 4,
			Word:   "Afirmation Word Neutral 2",
		},
		{
			ID:     uuid.New(),
			MoodID: 4,
			Word:   "Afirmation Word Neutral 3",
		},
	}

	err := db.CreateInBatches(&words, len(words)).Error
	if err != nil {
		return err
	}

	return nil
}

func artikelSeeder(db *gorm.DB) error {
	var articles []entity.Artikel

	artikel1 := entity.Artikel{
		ID:    uuid.New(),
		Title: "Meningkatkan Kesadaran tentang Kesehatan Mental: Pentingnya Perawatan Diri",
		Body:  "Kesehatan mental adalah bagian integral dari kesejahteraan kita secara keseluruhan. Sayangnya, masih ada stigma yang melekat pada masalah kesehatan mental, yang dapat mencegah banyak orang untuk mencari bantuan atau dukungan yang mereka butuhkan. Oleh karena itu, penting bagi kita untuk meningkatkan kesadaran tentang pentingnya merawat kesehatan mental kita dengan serius. Salah satu hal yang perlu dipahami adalah bahwa kesehatan mental bukanlah sesuatu yang statis, tetapi merupakan spektrum yang bergerak. Sama seperti kesehatan fisik kita, kesehatan mental kita juga bisa berubah dari waktu ke waktu. Mungkin ada saat-saat ketika kita merasa baik-baik saja, tetapi juga ada saat-saat ketika kita merasa tegang, cemas, atau sedih. Ini adalah hal yang normal, tetapi jika perasaan-perasaan tersebut terus berlanjut atau mengganggu kehidupan sehari-hari, penting untuk mencari bantuan. Ada banyak cara yang dapat kita lakukan untuk merawat kesehatan mental kita. Salah satunya adalah dengan menjaga keseimbangan antara pekerjaan, waktu luang, dan istirahat. Kita juga perlu mengakui dan menghargai perasaan kita, tanpa merasa bersalah atau malu atas apa yang kita rasakan. Berbicara dengan orang-orang yang kita percayai atau mencari bantuan dari profesional kesehatan mental juga merupakan langkah yang sangat penting. Tidak hanya itu, gaya hidup sehat juga berperan penting dalam menjaga kesehatan mental. Olahraga teratur, pola makan yang seimbang, cukup tidur, dan mengelola stres dengan baik dapat membantu menjaga keseimbangan kimiawi dalam otak kita, yang merupakan kunci untuk kesehatan mental yang baik. Terakhir, penting bagi kita semua untuk menghilangkan stigma seputar kesehatan mental. Kita harus menganggapnya sama pentingnya dengan kesehatan fisik dan memberikan dukungan kepada mereka yang membutuhkannya. Dengan begitu, kita dapat menciptakan lingkungan yang lebih inklusif dan peduli terhadap kesejahteraan mental setiap individu.",
	}
	artikel1.ArtikelImage = append(artikel1.ArtikelImage, entity.ArtikelImage{
		ID:        uuid.New(),
		ArtikelID: artikel1.ID,
		Image:     "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Artikel/christina-wocintechchat-com-rCyiK4_aaWw-unsplash.jpg",
	})
	artikel1.ArtikelImage = append(artikel1.ArtikelImage, entity.ArtikelImage{
		ID:        uuid.New(),
		ArtikelID: artikel1.ID,
		Image:     "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Artikel/tim-mossholder-8R-mXppeakM-unsplash.jpg",
	})

	artikel2 := entity.Artikel{
		ID:    uuid.New(),
		Title: "Mengatasi Stres: Tips untuk Mengelola Stres Sehari-hari",
		Body:  "Stres adalah bagian alami dari kehidupan sehari-hari, tetapi jika tidak dikelola dengan baik, dapat berdampak negatif pada kesehatan fisik dan mental kita. Oleh karena itu, penting untuk memiliki strategi yang efektif dalam mengatasi stres. Dengan meningkatkan kesadaran akan cara mengelola stres, kita dapat meningkatkan kesejahteraan kita secara keseluruhan. Pertama-tama, penting untuk mengidentifikasi pemicu stres kita. Setiap orang memiliki pemicu stres yang berbeda-beda, dan dengan mengenali apa yang memicu stres dalam hidup kita, kita dapat mencari cara untuk mengatasinya. Misalnya, mungkin stres disebabkan oleh tumpukan pekerjaan, konflik interpersonal, atau perubahan hidup. Salah satu strategi yang efektif dalam mengelola stres adalah dengan menggunakan teknik relaksasi. Meditasi, pernapasan dalam, dan yoga adalah contoh teknik relaksasi yang dapat membantu mengurangi tingkat stres dan meningkatkan kesejahteraan secara keseluruhan. Meluangkan waktu untuk melakukan aktivitas yang menenangkan pikiran dapat membantu mengurangi tekanan dan kecemasan yang kita rasakan. Selain itu, penting untuk menjaga keseimbangan antara pekerjaan dan waktu luang. Terkadang, stres disebabkan oleh kelebihan beban kerja atau kurangnya waktu untuk diri sendiri. Mengatur jadwal yang seimbang antara kerja, istirahat, dan rekreasi dapat membantu mengurangi stres dan meningkatkan produktivitas. Mengelola stres juga melibatkan mengubah pola pikir kita. Mengubah pola pikir negatif menjadi positif dapat membantu kita menghadapi tantangan dengan lebih baik dan mengurangi tingkat stres yang kita rasakan. Berlatihlah untuk melihat situasi dari sudut pandang yang berbeda dan fokus pada hal-hal yang dapat kita kendalikan. Terakhir, jangan ragu untuk mencari dukungan dari orang-orang terdekat atau profesional jika merasa kesulitan mengatasi stres. Berbicara dengan orang lain tentang apa yang kita rasakan dapat membantu mengurangi beban yang kita rasakan dan memberikan kita pandangan baru tentang cara menghadapi stres. Dengan menerapkan tips-tips ini dalam kehidupan sehari-hari, kita dapat mengelola stres dengan lebih efektif dan meningkatkan kesejahteraan kita secara keseluruhan. Ingatlah bahwa mengatasi stres adalah proses, dan dengan kesabaran dan ketekunan, kita dapat mengatasi tantangan dengan lebih baik dan hidup dengan lebih sejahtera.",
	}
	artikel2.ArtikelImage = append(artikel2.ArtikelImage, entity.ArtikelImage{
		ID:        uuid.New(),
		ArtikelID: artikel2.ID,
		Image:     "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Artikel/luis-villasmil-mlVbMbxfWI4-unsplash.jpg",
	})
	artikel2.ArtikelImage = append(artikel2.ArtikelImage, entity.ArtikelImage{
		ID:        uuid.New(),
		ArtikelID: artikel2.ID,
		Image:     "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Artikel/jared-rice-NTyBbu66_SI-unsplash.jpg",
	})

	artikel3 := entity.Artikel{
		ID:    uuid.New(),
		Title: "Meningkatkan Kesejahteraan Emosional: Pentingnya Koneksi Sosial",
		Body:  "Koneksi sosial adalah salah satu aspek penting dalam meningkatkan kesejahteraan emosional kita. Manusia adalah makhluk sosial, dan memiliki hubungan yang kuat dengan orang lain dapat memberikan dukungan yang sangat dibutuhkan dalam menghadapi tantangan hidup. Dalam artikel ini, kita akan menjelajahi mengapa koneksi sosial begitu penting dan bagaimana kita dapat memperkuatnya untuk meningkatkan kesejahteraan emosional kita. Pertama-tama, memiliki hubungan yang positif dengan orang lain dapat memberikan rasa dukungan dan pengertian. Ketika kita merasa terhubung dengan orang lain, kita memiliki tempat untuk berbagi perasaan, pikiran, dan pengalaman kita. Ini dapat membantu kita merasa lebih didengar, dipahami, dan diterima, yang pada gilirannya dapat meningkatkan kepercayaan diri dan kesejahteraan emosional kita. Selain itu, koneksi sosial juga dapat memberikan kita kesempatan untuk belajar dan tumbuh. Dalam hubungan yang sehat, kita dapat belajar dari pengalaman orang lain, mendapatkan dukungan dan dorongan untuk mencapai tujuan kita, serta menemukan kekuatan dan kelemahan kita sendiri. Hal ini dapat membantu kita mengembangkan keterampilan sosial, meningkatkan keterampilan komunikasi, dan mengembangkan rasa empati terhadap orang lain. Tidak hanya itu, koneksi sosial juga dapat memberikan kita rasa tujuan dan makna dalam hidup. Ketika kita merasa terhubung dengan orang lain dan merasa bahwa kontribusi kita dihargai, kita merasa lebih terikat pada masyarakat dan lebih termotivasi untuk berkontribusi secara positif. Ini dapat memberikan kita perasaan kepuasan dan kebahagiaan yang mendalam, serta membantu kita mengatasi rasa kesepian dan isolasi. Meskipun penting untuk memiliki koneksi sosial yang kuat, terkadang membangun dan memelihara hubungan bisa menjadi tantangan. Oleh karena itu, penting untuk mengambil langkah-langkah aktif dalam memperkuat koneksi sosial kita. Ini bisa termasuk mengambil inisiatif untuk bertemu dengan orang baru, menjaga komunikasi terbuka dengan teman dan keluarga, serta mencari kelompok atau komunitas yang memiliki minat dan nilai yang sama dengan kita. Dengan memprioritaskan koneksi sosial dalam hidup kita, kita dapat meningkatkan kesejahteraan emosional kita dan merasa lebih terhubung dengan dunia di sekitar kita. Ingatlah bahwa hubungan tidak hanya memberi kita dukungan dan pengertian, tetapi juga memberikan makna dan tujuan dalam hidup kita. Dengan menjaga dan merawat hubungan kita dengan orang lain, kita dapat menciptakan kehidupan yang lebih memuaskan dan berarti bagi diri kita sendiri dan orang lain.",
	}
	artikel3.ArtikelImage = append(artikel3.ArtikelImage, entity.ArtikelImage{
		ID:        uuid.New(),
		ArtikelID: artikel3.ID,
		Image:     "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Artikel/helena-lopes-e3OUQGT9bWU-unsplash.jpg",
	})

	artikel4 := entity.Artikel{
		ID:    uuid.New(),
		Title: "Mendukung Kesehatan Emosional: Pentingnya Hobi dan Kreativitas",
		Body:  "Dalam kehidupan yang sibuk dan sering kali penuh tekanan, menemukan waktu untuk mengekspresikan diri melalui hobi dan kreativitas dapat menjadi salah satu kunci untuk mendukung kesehatan emosional kita. Dalam artikel ini, kita akan menjelajahi mengapa hobi dan kreativitas penting, serta bagaimana kita dapat mengintegrasikannya ke dalam rutinitas sehari-hari untuk meningkatkan kesejahteraan emosional kita. Hobi dan kreativitas memungkinkan kita untuk mengekspresikan diri dan mengalami perasaan positif yang mendalam. Ketika kita terlibat dalam kegiatan yang kita nikmati, seperti melukis, menulis, memasak, atau berkebun, kita dapat merasa lebih bahagia, tenang, dan puas. Aktivitas ini memungkinkan kita untuk menyatu dengan diri kita sendiri, meningkatkan keseimbangan emosional, dan mengurangi tingkat stres yang mungkin kita rasakan. Selain itu, hobi dan kreativitas juga dapat menjadi bentuk terapi yang efektif. Mengekspresikan emosi melalui seni atau kegiatan kreatif lainnya dapat membantu kita mengatasi konflik internal, mengurangi kecemasan, dan meningkatkan rasa kesejahteraan secara keseluruhan. Aktivitas kreatif juga dapat membantu kita mengalami perasaan pencapaian dan kepuasan, yang dapat meningkatkan kepercayaan diri dan harga diri kita. Tidak hanya itu, melibatkan diri dalam hobi dan kreativitas juga dapat membantu kita mengembangkan keterampilan baru dan meningkatkan rasa pencapaian. Ketika kita menantang diri kita sendiri untuk belajar sesuatu yang baru atau meningkatkan kemampuan kita dalam bidang tertentu, kita dapat merasa lebih kompeten dan berdaya. Ini dapat meningkatkan motivasi dan kepercayaan diri kita, serta memberikan kita kesempatan untuk terus tumbuh dan berkembang sebagai individu. Meskipun hidup sibuk, penting untuk membuat waktu untuk mengejar hobi dan kreativitas dalam kehidupan sehari-hari. Ini bisa termasuk mengatur jadwal khusus untuk kegiatan kreatif, menemukan waktu di antara kesibukan untuk mengekspresikan diri, atau bahkan mengambil cuti singkat untuk fokus pada proyek kreatif yang lebih besar. Dengan membuat hobi dan kreativitas sebagai bagian penting dari rutinitas kita, kita dapat meningkatkan kesehatan emosional kita dan merasa lebih bersemangat tentang hidup. Dengan demikian, mendukung kesehatan emosional kita melalui hobi dan kreativitas dapat menjadi salah satu cara yang efektif untuk mencapai keseimbangan dan kesejahteraan dalam hidup kita. Dengan menemukan waktu untuk mengekspresikan diri dan mengejar minat kita, kita dapat merasa lebih bahagia, tenang, dan terhubung dengan diri kita sendiri dan dunia di sekitar kita.",
	}
	artikel4.ArtikelImage = append(artikel4.ArtikelImage, entity.ArtikelImage{
		ID:        uuid.New(),
		ArtikelID: artikel4.ID,
		Image:     "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Artikel/jeshoots-com-eCktzGjC-iU-unsplash.jpg",
	})
	artikel4.ArtikelImage = append(artikel4.ArtikelImage, entity.ArtikelImage{
		ID:        uuid.New(),
		ArtikelID: artikel4.ID,
		Image:     "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Artikel/steve-johnson-A2OL6S9zB7o-unsplash.jpg",
	})

	artikel5 := entity.Artikel{
		ID:    uuid.New(),
		Title: "Menghargai Diri Sendiri: Pentingnya Praktik Penerimaan Diri",
		Body:  "Menghargai diri sendiri adalah langkah penting dalam perjalanan menuju kesejahteraan emosional yang seimbang. Dalam artikel ini, kita akan membahas mengapa praktik penerimaan diri sangat penting dan bagaimana kita dapat mengintegrasikannya ke dalam kehidupan sehari-hari untuk meningkatkan kesejahteraan kita secara keseluruhan. Pertama-tama, penting untuk menyadari bahwa menerima diri sendiri bukanlah hal yang mudah, terutama di dunia yang sering kali menekankan standar yang tidak realistis dan gambaran diri yang sempurna. Namun, ketika kita mempraktikkan penerimaan diri, kita membangun fondasi yang kuat untuk kesejahteraan emosional kita. Ini berarti menerima semua bagian dari diri kita, baik yang positif maupun negatif, tanpa penilaian atau kritik yang berlebihan. Praktik penerimaan diri juga melibatkan memahami bahwa kita semua memiliki kelemahan dan kekurangan, dan itu tidak mengurangi nilai atau martabat kita sebagai individu. Ketika kita belajar menerima dan merangkul keunikan kita, kita dapat merasa lebih percaya diri dan berdaya dalam menghadapi tantangan hidup. Selain itu, penerimaan diri juga melibatkan menghargai dan memperlakukan diri kita sendiri dengan penuh kasih sayang dan penghargaan. Ini berarti memberi diri kita waktu untuk istirahat dan pemulihan, menghargai batas-batas kita, dan melakukan apa pun yang diperlukan untuk merawat diri kita sendiri dengan baik. Ketika kita menghargai diri sendiri dengan baik, kita menciptakan pondasi yang kuat untuk kesejahteraan emosional kita. Praktik penerimaan diri juga memungkinkan kita untuk lebih terhubung dengan orang lain secara lebih autentik. Ketika kita menerima dan mencintai diri sendiri sepenuhnya, kita dapat membawa kehadiran yang lebih otentik dan terbuka ke dalam hubungan kita dengan orang lain. Ini membantu kita membangun hubungan yang lebih dalam dan lebih bermakna, serta merasa lebih terhubung dengan dunia di sekitar kita. Meskipun mungkin sulit pada awalnya, penting untuk terus berlatih penerimaan diri setiap hari. Ini bisa melibatkan praktik seperti meditasi, jurnal, atau terapi yang membantu kita mengeksplorasi dan menerima diri kita sendiri dengan lebih dalam. Dengan waktu dan kesabaran, kita dapat membangun hubungan yang lebih sehat dengan diri kita sendiri dan merasakan kesejahteraan emosional yang lebih besar. Dengan demikian, menghargai diri sendiri melalui praktik penerimaan diri adalah langkah penting dalam perjalanan menuju kesejahteraan emosional yang seimbang. Dengan menerima diri kita sendiri sepenuhnya dan merawat diri kita sendiri dengan penuh kasih sayang, kita dapat menciptakan kehidupan yang lebih memuaskan dan bermakna bagi diri kita sendiri dan orang lain di sekitar kita.",
	}
	artikel5.ArtikelImage = append(artikel5.ArtikelImage, entity.ArtikelImage{
		ID:        uuid.New(),
		ArtikelID: artikel5.ID,
		Image:     "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Artikel/annie-spratt-9R3izhP3rtI-unsplash.jpg",
	})

	articles = append(articles, artikel1)
	articles = append(articles, artikel2)
	articles = append(articles, artikel3)
	articles = append(articles, artikel4)
	articles = append(articles, artikel5)

	if err := db.CreateInBatches(&articles, len(articles)).Error; err != nil {
		return err
	}

	return nil
}

func videoSeeder(db *gorm.DB) error {
	var videos []entity.Video

	video1 := entity.Video{
		ID:          uuid.New(),
		Title:       "Belajar Psikologi - Kesehatan Mental",
		Description: "Seiring berkembangnya zaman, tentunya tekanan dalam kehidupan sehari-hari juga ikut berkembang. Pastinya, beragam tekanan ini dapat mempengaruhi kesehatan mental kita semua. Pada kenyataannya, memang benar angka penderita gangguan mental terus meningkat setiap tahunnya, lebih jauh lagi, ternyata gangguan mental juga dapat dirasakan oleh anak-anak yang berusia dibawah 18 tahun. Akan tetapi, kita tidak boleh putus asa, tentunya kita juga harus memanfaatkan perkembangan zaman ini untuk memperdalam pengetahuan kita mengenai kesehatan dan gangguan mental.",
		Link:        "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Videos/Kesehatan%20Mental.mp4",
	}

	video2 := entity.Video{
		ID:          uuid.New(),
		Title:       "Bagaimana Stress Mempengaruhi Otak",
		Description: "Stres tidak selalu merupakan hal yang buruk; ini bisa berguna untuk mengeluarkan energi dan fokus ekstra, seperti saat Anda bermain olahraga kompetitif atau harus berbicara di depan umum. Namun bila terus menerus, hal itu sebenarnya mulai mengubah otak Anda. Madhumita Murgia menunjukkan bagaimana stres kronis dapat memengaruhi ukuran otak, strukturnya, dan fungsinya, hingga ke tingkat gen Anda.",
		Link:        "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Videos/How%20stress%20affects%20your%20brain.mp4",
	}

	videos = append(videos, video1)
	videos = append(videos, video2)

	if err := db.CreateInBatches(&videos, len(videos)).Error; err != nil {
		return err
	}
	return nil
}

func podcastSeeder(db *gorm.DB) error {
	var podcasts []entity.Podcast

	podcast1 := entity.Podcast{
		ID:          uuid.New(),
		Title:       "Mental Health di Kalangan Milenial",
		Description: "Podcast ini membahas tentang kesehatan mental di kalangan milenial. Dalam podcast ini, kita akan membahas tentang berbagai aspek kesehatan mental, mulai dari pentingnya merawat kesehatan mental kita, hingga cara mengatasi stres dan kecemasan yang sering kita alami sehari-hari. Dengan mendengarkan podcast ini, diharapkan kita dapat lebih memahami pentingnya kesehatan mental dan bagaimana kita dapat merawatnya dengan baik.",
		Link:        "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Podcast/Mental%20Health%20di%20Kalangan%20Milenial.mp3",
		Thumbnail:   "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Podcast/dan-meyers-hluOJZjLVXc-unsplash.jpg",
	}

	podcasts = append(podcasts, podcast1)

	if err := db.CreateInBatches(&podcasts, len(podcasts)).Error; err != nil {
		return err
	}
	return nil
}

func eventSeeder(db *gorm.DB) error {
	var events []entity.Event

	event1 := entity.Event{
		ID:               uuid.New(),
		Title:            "Seminar Membangun Kesehatan Mental: Menjelajahi Keseimbangan Emosional",
		Body:             "Event ini adalah sebuah seminar yang bertujuan untuk meningkatkan kesadaran dan pemahaman tentang kesehatan mental di kalangan masyarakat. Dalam seminar ini, kami akan menyajikan berbagai data dan informasi terkini mengenai prevalensi masalah kesehatan mental, faktor risiko, serta strategi untuk membangun keseimbangan emosional. Peserta akan diajak untuk memahami pentingnya menjaga kesehatan mental, mengidentifikasi tanda-tanda gangguan mental, dan belajar teknik-teknik yang dapat membantu mengelola stres dan meningkatkan kesejahteraan emosional. Melalui diskusi interaktif dan sharing pengalaman, kami berharap peserta dapat merasa lebih berdaya dalam menjaga kesehatan mental mereka sendiri serta mendukung orang-orang di sekitar mereka.",
		Location:         "Universitas Brawijaya, Malang",
		StartDate:        time.Date(2024, 03, 22, 8, 00, 00, 00, time.Local),
		EndDate:          time.Date(2024, 03, 22, 12, 00, 00, 00, time.Local),
		IsRequirePayment: false,
		PaymentAmount:    0,
	}
	event1.EventImage = append(event1.EventImage, entity.EventImage{
		ID:        uuid.New(),
		EventID:   event1.ID,
		ImageLink: "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Event/austin-distel-rxpThOwuVgE-unsplash.jpg",
	})

	event2 := entity.Event{
		ID:               uuid.New(),
		Title:            "Walk for Wellness: Langkah Pertama Menuju Kesehatan Mental yang Lebih Baik",
		Body:             "\"Walk for Wellness\" adalah acara komunitas yang bertujuan untuk menggalang kesadaran tentang pentingnya kesehatan mental sambil mengutamakan kebugaran fisik. Acara ini menawarkan kesempatan bagi peserta untuk berpartisipasi dalam kegiatan jalan kaki bersama dengan teman-teman, keluarga, dan masyarakat sekitar. Selama perjalanan, peserta akan diberi kesempatan untuk berinteraksi, berbagi pengalaman, dan belajar lebih lanjut tentang strategi-strategi sederhana untuk merawat kesehatan mental mereka. Dengan menggabungkan gerakan fisik dengan kesadaran mental, \"Walk for Wellness\" bertujuan untuk mempromosikan perubahan gaya hidup yang lebih seimbang dan membangun komunitas yang peduli akan kesejahteraan mental bersama-sama.",
		Location:         "Taman Kota, Surabaya",
		StartDate:        time.Date(2024, 04, 15, 7, 00, 00, 00, time.Local),
		EndDate:          time.Date(2024, 04, 16, 10, 00, 00, 00, time.Local),
		IsRequirePayment: true,
		PaymentAmount:    50000,
	}
	event2.EventImage = append(event2.EventImage, entity.EventImage{
		ID:        uuid.New(),
		EventID:   event2.ID,
		ImageLink: "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Event/maxime-bhm-icyZmdkCGZ0-unsplash.jpg",
	})

	event3 := entity.Event{
		ID:               uuid.New(),
		Title:            "Mewujudkan Keseimbangan Mental: Yoga dan Meditasi untuk Kesejahteraan Emosional",
		Body:             "Event \"Mewujudkan Keseimbangan Mental\" adalah sebuah acara yang dirancang untuk memperkenalkan pentingnya yoga dan meditasi dalam meningkatkan kesehatan mental. Dalam event ini, peserta akan diajak untuk mengalami sesi yoga yang difokuskan pada pernapasan dan gerakan yang menenangkan, serta praktik meditasi yang membantu mengurangi stres dan meningkatkan konsentrasi. Data akan disajikan mengenai manfaat ilmiah dari praktik-praktik ini terhadap kesejahteraan emosional, termasuk penelitian terbaru tentang efeknya terhadap gangguan kecemasan dan depresi. Melalui pengalaman langsung dan panduan praktis, kami berharap peserta akan merasakan manfaat segera dari praktik-praktik ini dan dapat mengintegrasikannya ke dalam kehidupan sehari-hari mereka untuk mencapai keseimbangan mental yang lebih baik.",
		Location:         "Studio Yoga Sehat, Jakarta",
		StartDate:        time.Date(2024, 05, 20, 9, 00, 00, 00, time.Local),
		EndDate:          time.Date(2024, 05, 20, 12, 00, 00, 00, time.Local),
		IsRequirePayment: true,
		PaymentAmount:    75000,
	}
	event3.EventImage = append(event3.EventImage, entity.EventImage{
		ID:        uuid.New(),
		EventID:   event3.ID,
		ImageLink: "https://lhchzzctoliiqnjoibcp.supabase.co/storage/v1/object/public/heal_in/Event/kaylee-garrett-GaprWyIw66o-unsplash.jpg",
	})

	events = append(events, event1)
	events = append(events, event2)
	events = append(events, event3)

	if err := db.CreateInBatches(&events, len(events)).Error; err != nil {
		return err
	}

	return nil
}

func SeedData(db *gorm.DB, bcrypt *bcrypt.BcryptInterface) {
	var totalQuestion int64
	var totalMood int64
	var totalUser int64
	var totalAfirmationWord int64
	var totalArtikel int64
	var totalVideo int64
	var totalPodcast int64
	var totalEvent int64

	err := db.Model(&entity.JournalingQuestion{}).Count(&totalQuestion).Error
	if err != nil {
		log.Fatalf("Error counting question data: %v", err)
	}

	err = db.Model(&entity.JournalingMood{}).Count(&totalMood).Error
	if err != nil {
		log.Fatalf("Error counting mood data: %v", err)
	}

	err = db.Model(&entity.User{}).Count(&totalUser).Error
	if err != nil {
		log.Fatalf("Error counting user data: %v", err)
	}

	err = db.Model(&entity.AfirmationWord{}).Count(&totalAfirmationWord).Error
	if err != nil {
		log.Fatalf("Error counting affirmation word data: %v", err)
	}

	err = db.Model(&entity.Artikel{}).Count(&totalAfirmationWord).Error
	if err != nil {
		log.Fatalf("Error counting affirmation word data: %v", err)
	}

	err = db.Model(&entity.Video{}).Count(&totalVideo).Error
	if err != nil {
		log.Fatalf("Error counting video data: %v", err)
	}

	err = db.Model(&entity.Podcast{}).Count(&totalPodcast).Error
	if err != nil {
		log.Fatalf("Error counting podcast data: %v", err)
	}

	err = db.Model(&entity.Event{}).Count(&totalEvent).Error
	if err != nil {
		log.Fatalf("Error counting event data: %v", err)
	}

	// seed data if there is no data in the table
	if totalQuestion == 0 {
		err := questionSeeder(db)

		if err != nil {
			log.Fatalf("Error seeding question data: %v", err)
			return
		}
		log.Println("Question data seeded")
	}

	if totalMood == 0 {
		err := moodSeeder(db)

		if err != nil {
			log.Fatalf("Error seeding mood data: %v", err)
			return
		}
		log.Println("Mood data seeded")
	}

	if totalUser == 0 {
		err := userSeeder(db, *bcrypt)

		if err != nil {
			log.Fatalf("Error seeding user data: %v", err)
			return
		}
		log.Println("User data seeded")
	}

	if totalAfirmationWord == 0 {
		err := afirmationWordSeeder(db)

		if err != nil {
			log.Fatalf("Error seeding affirmation word data: %v", err)
			return
		}
		log.Println("Afirmation words data seeded")
	}

	if totalArtikel == 0 {
		if err := artikelSeeder(db); err != nil {
			log.Fatalf("Error seeding artikel data: %v", err)
			return
		}
		log.Println("Artikel data seeded")
	}

	if totalVideo == 0 {
		if err := videoSeeder(db); err != nil {
			log.Fatalf("Error seeding video data: %v", err)
			return
		}
		log.Println("Video data seeded")
	}

	if totalPodcast == 0 {
		if err := podcastSeeder(db); err != nil {
			log.Fatalf("Error seeding podcast data: %v", err)
			return
		}
		log.Println("Podcast data seeded")
	}

	if totalEvent == 0 {
		if err := eventSeeder(db); err != nil {
			log.Fatalf("Error seeding event data: %v", err)
			return
		}
		log.Println("Event data seeded")
	}
}
