package main

// Wine is wine
type Wine struct {
	Name      string
	ASIN      string
	Producer  string
	Origin    string
	Vintage   string
	Sweetness string
	Grape     string
	Price     string
	Quantity  string
	Text      string
	Image     string
}

// http://www.amazon.de/s/ref=sr_nr_p_72_0?fst=as%3Aoff&rh=n%3A340846031%2Ck%3Amerlot%2Cp_n_feature_four_browse-bin%3A2317930031%2Cp_n_feature_thirteen_browse-bin%3A5715460031%2Cp_72%3A419117031&keywords=merlot&ie=UTF8&qid=1460750487&rnid=419116031

// Products is the list of products
var Products = []Wine{
	{
		"Sweet Reeh",
		"B00OB9L0NM",
		"Reeh Hannes",
		"Österreich",
		"2014",
		"süß",
		"Merlot, Zweigelt",
		"9,09",
		"1 x 0.75 l (EUR 12,12 / l)",
		"Ein leuchtendes Kirschrot sticht ins Auge. In der Nase ganz typisch ein Merlotgeruch. Am Gaumen zeigt er sich weich, mild mit leichter Würze und Beeren. Abgerundet wird der Geschmack durch eine zarte Restsüße sowie mit einem Abgang der an reife Brombeeren erinnert.",
		"31KLOh3UlvL.jpg",
	},
	{
		"Black Print",
		"B017KOW3TO",
		"Markus Schneider",
		"Deutschland",
		"2014",
		"trocken",
		"Merlot",
		"14,90",
		"1 x 0.75 l (EUR 19,87 / l)",
		"Black Print - mit ihm begann alles! Eine traditionelle Maischegärung im Holzbottich und mit Liebe gepresst. Im Glas ein faszinierender Wein - dunkel Purpur duftet wie ein Rotwein aus Frankreichs Süden. Faszinierend, geschmeidig und natürlich am Gaumen. Der Jahrgang 2014 ist ein Cuvée aus Merlot, Cabernet Sauvignon und Cabernet Dorsa. Hier ist nur das allerbeste Traubenmaterial zum Einsatz gekommen, kombiniert mit deutscher Handwerkskunst. Das ist ganz grosses Weinkino aus Deutschland! Primärfruchtig mit Cassis, Pflaumen und Rumtopf.",
		"31f6DJEyRRL.jpg",
	},
	{
		"Merlot Terroir",
		"B003VAMICE",
		"Miolo",
		"Brasilien",
		"2009",
		"trocken",
		"Merlot",
		"22,47",
		"1 x 0.75 l (EUR 29,96 / l)",
		"Rotwein aus dem Vale dos Vinhedos. Selektierte Merlot-Reben aus speziellen Jahrgängen. 14% vol. Limitierte Auflage!",
		"316EKa7E0XL.jpg",
	},
	{
		"Paul Mas 2014er Merlot",
		"B00P1JH0BW",
		"Paul Mas",
		"Frankreich",
		"2014",
		"trocken",
		"Merlot",
		"7,50",
		"1 x 0.75 l (EUR 10,00 / l)",
		"Klares, kräftiges Granatrot,aromatisches Bouquet nach dunklen Beeren u. Früchten wie Kirsche, Pflaume u. Heidelbeere, vollmundig u. geschmeidig am Gaumen mit zarten Tanninen u. feinen Vanillenoten, langer eleganter Abgang. Fruchtig und Fruchtbetont. Samtig und weich. Elegant. Kräftig und Körperreich. Vollmundig. Holzig.",
		"31C5gnojzaL.jpg",
	},
	{
		"Jean d'Alibert",
		"B0169NX666",
		"Chantovent",
		"Frankreich",
		"2014",
		"trocken",
		"Merlot",
		"29,70",
		"6 x 0.75 l (EUR 6,60 / l)",
		"Der Merlot „Jean d’Alibert“ ist einer der absoluten Lieblingsprodukte vieler Kunden. Warum? Weil sich weich, rund und fruchtig präsentiert, dabei nicht zu schwer ist und er daher so viel Trinkfreude bereitet. Traditionelle Maischegärung mit anschließendem Holzfasslager für neun Monate. Am Gaumen weich, sehr klare Beerenfrucht und leichte Aromen von Bitterschokolade, rund und ausgewogen, animierend, ordentlicher Abgang",
		"31VOG7241VL.jpg",
	},
}
