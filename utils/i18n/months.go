package i18n

import "strings"

var MonthNames = map[string]string{}

func init() {
	for _, n := range januaryNames {
		n := strings.ToLower(n)
		MonthNames[n] = "January"
	}
	for _, n := range februaryNames {
		n := strings.ToLower(n)
		MonthNames[n] = "February"
	}
	for _, n := range marchNames {
		n := strings.ToLower(n)
		MonthNames[n] = "March"
	}
	for _, n := range aprilNames {
		n := strings.ToLower(n)
		MonthNames[n] = "April"
	}
	for _, n := range mayNames {
		n := strings.ToLower(n)
		MonthNames[n] = "May"
	}
	for _, n := range juneNames {
		n := strings.ToLower(n)
		MonthNames[n] = "June"
	}
	for _, n := range julyNames {
		n := strings.ToLower(n)
		MonthNames[n] = "July"
	}
	for _, n := range augustNames {
		n := strings.ToLower(n)
		MonthNames[n] = "August"
	}
	for _, n := range septemberNames {
		n := strings.ToLower(n)
		MonthNames[n] = "September"
	}
	for _, n := range octoberNames {
		n := strings.ToLower(n)
		MonthNames[n] = "October"
	}
	for _, n := range novemberNames {
		n := strings.ToLower(n)
		MonthNames[n] = "November"
	}
	for _, n := range decemberNames {
		n := strings.ToLower(n)
		MonthNames[n] = "December"
	}
}

var januaryNames = []string{
	"January", "يناير", "1月", "1月", "Januari", "Janvier", "Januar",
	"Gennaio", "1月", "1월", "Styczeń", "Janeiro", "Январь", "Enero", "มกราคม",
	"Ocak", "Януари", "Gener", "Siječanj", "Leden", "Januar", "Enero",
	"Tammikuu", "Ιανουάριος", "ינואר", "जनवरी", "Január", "Januari", "Janvāris",
	"Sausis", "Januar", "Ianuarie", "Јануар", "Január", "Januar", "Januari",
	"Січень", "Tháng một", "ژانویه", "Jaanuar", "Januari",
}

var februaryNames = []string{
	"February", "فبراير", "2月", "2月", "Februari", "Février", "Februar",
	"Febbraio", "2月", "2월", "Luty", "Fevereiro", "Февраль", "Febrero",
	"กุมภาพันธ์", "Şubat", "Февруари", "Febrer", "Veljača", "Únor", "Februar",
	"Pebrero", "Helmikuu", "Φεβρουάριος", "פברואר", "फ़रवरी", "Február",
	"Februari", "Februāris", "Vasaris", "Februar", "Februarie", "Фебруар",
	"Február", "Februar", "Februari", "Лютий", "Tháng hai", "فوریه",
	"Veebruar", "Februari",
}

var marchNames = []string{
	"March", "مارس", "3月", "3月", "Maart", "Mars", "März", "Marzo", "3月",
	"3월", "Marzec", "Março", "Март", "Marzo", "มีนาคม", "Mart", "Март", "Març",
	"Ožujak", "Březen", "Marts", "Marso", "Maaliskuu", "Μάρτιος", "מרץ", "मार्च",
	"Március", "Maret", "Marts", "Kovas", "Mars", "Martie", "Март", "Marec",
	"Marec", "Mars", "Березень", "Tháng ba", "مارس", "Märts", "Mac",
}

var aprilNames = []string{
	"April", "أبريل", "4月", "4月", "April", "Avril", "April", "Aprile", "4月",
	"4월", "Kwiecień", "Abril", "Апрель", "Abril", "เมษายน", "Nisan", "Април",
	"Abril", "Travanj", "Duben", "April", "Abril", "Huhtikuu", "Απρίλιος",
	"אפריל", "अप्रैल", "Április", "April", "Aprīlis", "Balandis", "April",
	"Aprilie", "Април", "Apríl", "April", "April", "Квітень", "Tháng tư",
	"آوریل", "Aprill", "April",
}

var mayNames = []string{
	"May", "مايو", "5月", "5月", "Mei", "Mai", "Mai", "Maggio", "5月", "5월",
	"Maj", "Maio", "Май", "Mayo", "พฤษภาคม", "Mayıs", "Май", "Maig", "Svibanj",
	"Květen", "Maj", "Mayo", "Toukokuu", "Μάιος", "מאי", "मई", "Május", "Mei",
	"Maijs", "Gegužė", "Mai", "Mai", "Мај", "Máj", "Maj", "Maj", "Травень",
	"Tháng năm", "مه", "Mai", "Mei",
}

var juneNames = []string{
	"June", "يونيو", "6月", "6月", "Juni", "Juin", "Juni", "Giugno", "6月",
	"6월", "Czerwiec", "Junho", "Июнь", "Junio", "มิถุนายน", "Haziran", "Юни",
	"Juny", "Lipanj", "Červen", "Juni", "Hunyo", "Kesäkuu", "Ιούνιος", "יוני",
	"जून", "Június", "Juni", "Jūnijs", "Birželis", "Juni", "Iunie", "Јун",
	"Jún", "Junij", "Juni", "Червень", "Tháng sáu", "ژوئن", "Juuni", "Jun",
}

var julyNames = []string{
	"July", "يوليو", "7月", "7月", "Juli", "Juillet", "Juli", "Luglio", "7月",
	"7월", "Lipiec", "Julho", "Июль", "Julio", "กรกฎาคม", "Temmuz", "Юли",
	"Juliol", "Srpanj", "Červenec", "Juli", "Hulyo", "Heinäkuu", "Ιούλιος",
	"יולי", "जुलाई", "Július", "Juli", "Jūlijs", "Liepa", "Juli", "Iulie", "Јул",
	"Júl", "Julij", "Juli", "Липень", "Tháng bảy", "ژوئیه", "Juuli", "Julai",
}

var augustNames = []string{
	"August", "أغسطس", "8月", "8月", "Augustus", "Août", "August", "Agosto",
	"8月", "8월", "Sierpień", "Agosto", "Август", "Agosto", "สิงหาคม",
	"Ağustos", "Август", "Agost", "Kolovoz", "Srpen", "August", "Agosto",
	"Elokuu", "Αύγουστος", "אוגוסט", "अगस्त", "Augusztus", "Agustus", "Augusts",
	"Rugpjūtis", "August", "August", "Август", "August", "Avgust", "Augusti",
	"Серпень", "Tháng tám", "اوت", "August", "Ogos",
}

var septemberNames = []string{
	"September", "سبتمبر", "9月", "9月", "September", "Septembre", "September",
	"Settembre", "9月", "9월", "Wrzesień", "Setembro", "Сентябрь",
	"Septiembre", "กันยายน", "Eylül", "Септември", "Setembre", "Rujan", "Září",
	"September", "Setyembre", "Syyskuu", "Σεπτέμβριος", "ספטמבר", "सितम्बर",
	"Szeptember", "September", "Septembris", "Rugsėjis", "September",
	"Septembrie", "Септембар", "September", "September", "September",
	"Вересень", "Tháng chín", "سپتامبر", "September", "September",
}

var octoberNames = []string{
	"October", "أكتوبر", "10月", "10月", "Oktober", "Octobre", "Oktober",
	"Ottobre", "10月", "10월", "Październik", "Outubro", "Октябрь", "Octubre",
	"ตุลาคม", "Ekim", "Октомври", "Octubre", "Listopad", "Říjen", "Oktober",
	"Oktubre", "Lokakuu", "Οκτώβριος", "אוקטובר", "अक्टूबर", "Október",
	"Oktober", "Oktobris", "Spalis", "Oktober", "Octombrie", "Октобар",
	"Október", "Oktober", "Oktober", "Жовтень", "Tháng mười", "اکتبر",
	"Oktoober", "Oktober",
}

var novemberNames = []string{
	"November", "نوفمبر", "11月", "11月", "November", "Novembre", "November",
	"Novembre", "11月", "11월", "Listopad", "Novembro", "Ноябрь", "Noviembre",
	"พฤศจิกายน", "Kasım", "Ноември", "Novembre", "Studeni", "Listopad",
	"November", "Nobyembre", "Marraskuu", "Νοέμβριος", "נובמבר", "नवम्बर",
	"November", "November", "Novembris", "Lapkritis", "November", "Noiembrie",
	"Новембар", "November", "November", "November", "Листопад", "Tháng mười một",
	"نوامبر", "November", "November",
}

var decemberNames = []string{
	"December", "ديسمبر", "12月", "12月", "December", "Décembre", "Dezember",
	"Dicembre", "12月", "12월", "Grudzień", "Dezembro", "Декабрь", "Diciembre",
	"ธันวาคม", "Aralık", "Декември", "Desembre", "Prosinac", "Prosinec",
	"December", "Disyembre", "Joulukuu", "Δεκέμβριος", "דצמבר", "दिसम्बर",
	"December", "Desember", "Decembris", "Gruodis", "Desember", "Decembrie",
	"Децембар", "December", "December", "December", "Грудень", "Tháng mười hai",
	"دسامبر", "Detsember", "Disember",
}
