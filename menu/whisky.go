package menu

type whisky struct {
	Name       string
	PricePart1 int
	PricePart2 int
}

var Whiskies []*whisky = []*whisky{
	&whisky{"윈저12", 200_000, 160_000},
	&whisky{"윈저아이스", 200_000, 160_000},
	&whisky{"펜텀", 200_000, 160_000},
	&whisky{"임페12", 200_000, 160_000},
	&whisky{"스카치", 200_000, 160_000},
	&whisky{"골든블루12", 210_000, 170_000},
	&whisky{"앱솔루트", 210_000, 170_000},
	&whisky{"몬테스알파", 210_000, 170_000},
	&whisky{"윈저17", 380_000, 190_000},
	&whisky{"윈저W17", 380_000, 190_000},
	&whisky{"임페17", 380_000, 190_000},
	&whisky{"골든17", 390_000, 200_000},
	&whisky{"티나", 390_000, 200_000},
	&whisky{"윈저W19", 420_000, 230_000},
	&whisky{"멈", 420_000, 230_000},
	&whisky{"퀀텀", 440_000, 250_000},
	&whisky{"멈올라프 로제", 440_000, 250_000},
	&whisky{"룩펠레어 로제", 440_000, 250_000},
	&whisky{"임페21", 460_000, 270_000},
	&whisky{"윈저21", 460_000, 270_000},
	&whisky{"맥켈란12(700ml)", 460_000, 270_000},
	&whisky{"글렌피딕12(700ml)", 460_000, 270_000},
	&whisky{"헤네시VSOP", 510_000, 320_000},
	&whisky{"발렌타인17(500ml)", 510_000, 320_000},
	&whisky{"발렌타인17(700ml)", 560_000, 370_000},
	&whisky{"로얄샬루트21(500ml)", 560_000, 370_000},
	&whisky{"발렌타인21(500ml)", 560_000, 370_000},
	&whisky{"죠니블루21(500ml)", 620_000, 430_000},
	&whisky{"맥켈란15(700ml)", 620_000, 430_000},
	&whisky{"헤네시XO(500ml)(150)", 660_000, 470_000},
	&whisky{"발렌타인21(700ml)", 660_000, 470_000},
	&whisky{"로얄샬루트21(700ml)", 660_000, 470_000},
	&whisky{"페리에쥬에", 680_000, 490_000},
	&whisky{"죠니블루(700ml)", 740_000, 550_000},
	&whisky{"발렌타인30", 1_640_000, 1_450_000},
}
