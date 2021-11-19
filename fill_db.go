package main

import (
	"ais.com/m/database"
	"ais.com/m/model"
)

func main() {
	pg := database.GetDB()
	err := pg.AutoMigrate(&model.Gun{})
	if err != nil {
		panic(err)
	}

	pg.Create(&model.Gun{
		Model:         "STM-9",
		Company:       "Союз-ТМ",
		Country:       "Россия",
		Year:          2019,
		Type:          "rifle",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        0,
		BarrelLength:  "330/355/400",
		Weight:        4.1,
		Caliber:       "9x19",
	})

	pg.Create(&model.Gun{
		Model:         "Desert Tech MDR",
		Company:       "Desert Tech LLC",
		Country:       "США",
		Year:          2015,
		Type:          "rifle",
		LoadType:      "semi-auto",
		IsBullPup:     true,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        689,
		BarrelLength:  "406",
		Weight:        3.4,
		Caliber:       "5.56x45",
	})

	pg.Create(&model.Gun{
		Model:         "DT-HTI BBA-A",
		Company:       "Desert Tech LLC",
		Country:       "США",
		Year:          2016,
		Type:          "rifle",
		LoadType:      "bolt-action",
		IsBullPup:     true,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        1143,
		BarrelLength:  "736",
		Weight:        8.84,
		Caliber:       ".50 BMG",
	})

	pg.Create(&model.Gun{
		Model:         "Сайга-9",
		Company:       "Калашников",
		Country:       "Россия",
		Year:          2014,
		Type:          "rifle",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        827,
		BarrelLength:  "367",
		Weight:        3.1,
		Caliber:       "9x19",
	})

	pg.Create(&model.Gun{
		Model:         "SR1",
		Company:       "Калашников",
		Country:       "Россия",
		Year:          2018,
		Type:          "rifle",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        940,
		BarrelLength:  "415",
		Weight:        4.2,
		Caliber:       "5.56x45",
	})

	pg.Create(&model.Gun{
		Model:         "Sabatti S.A.R",
		Company:       "Sabatti",
		Country:       "Италия",
		Year:          2020,
		Type:          "rifle",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        880,
		BarrelLength:  "370",
		Weight:        3.15,
		Caliber:       "5.56x45",
	})

	pg.Create(&model.Gun{
		Model:         "Сайга-МК",
		Company:       "Калашников",
		Country:       "Россия",
		Year:          2000,
		Type:          "rifle",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        940,
		BarrelLength:  "415",
		Weight:        3.6,
		Caliber:       "7.62x39",
	})

	pg.Create(&model.Gun{
		Model:         "Тигр исп. 01",
		Company:       "ИЖМАШ",
		Country:       "Россия",
		Year:          1992,
		Type:          "rifle",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        1090,
		BarrelLength:  "530",
		Weight:        4.1,
		Caliber:       "7.62x54",
	})

	pg.Create(&model.Gun{
		Model:         "Вепрь кал. 7,62х51",
		Company:       "Вятско-Полянский машиностроительный завод",
		Country:       "Россия",
		Year:          1995,
		Type:          "rifle",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        1035,
		BarrelLength:  "520",
		Weight:        4.7,
		Caliber:       "7.62x51",
	})

	pg.Create(&model.Gun{
		Model:         "Walther Colt M16",
		Company:       "Walther",
		Country:       "Германия",
		Year:          2009,
		Type:          "rifle",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        1010,
		BarrelLength:  "538",
		Weight:        3.0,
		Caliber:       ".22 LR",
	})

	pg.Create(&model.Gun{
		Model:         "Сайга-12К",
		Company:       "ИЖМАШ",
		Country:       "Россия",
		Year:          1997,
		Type:          "shotgun",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        910,
		BarrelLength:  "430",
		Weight:        3.5,
		Caliber:       "12",
	})

	pg.Create(&model.Gun{
		Model:         "Вепрь-12 \"Молот\"",
		Company:       "Вятско-Полянский машиностроительный завод",
		Country:       "Россия",
		Year:          2003,
		Type:          "shotgun",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        867,
		BarrelLength:  "305",
		Weight:        3.8,
		Caliber:       "12",
	})

	pg.Create(&model.Gun{
		Model:         "Remington 870",
		Company:       "Remington",
		Country:       "США",
		Year:          1951,
		Type:          "shotgun",
		LoadType:      "pomp",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        1060,
		BarrelLength:  "254/770",
		Weight:        3.6,
		Caliber:       "12",
	})

	pg.Create(&model.Gun{
		Model:         "Benelli Supernova",
		Company:       "Benelli",
		Country:       "Италия",
		Year:          2006,
		Type:          "shotgun",
		LoadType:      "pomp",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        1257,
		BarrelLength:  "710",
		Weight:        3.8,
		Caliber:       "12",
	})

	pg.Create(&model.Gun{
		Model:         "Hatsan OPTIMA SR S12",
		Company:       "Hatsan",
		Country:       "Турция",
		Year:          1997, // ?
		Type:          "shotgun",
		LoadType:      "break",
		BreakType:     "vertical",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 2,
		Length:        910, 
		BarrelLength:  "430",
		Weight:        3.5,
		Caliber:       "12",
	})

	pg.Create(&model.Gun{
		Model:         "ИЖ-43",
		Company:       "ИЖМЕХ",
		Country:       "Россия",
		Year:          1986,
		Type:          "shotgun",
		LoadType:      "break",
		BreakType:     "horizontal",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 2,
		Length:        910,
		BarrelLength:  "720",
		Weight:        3.5,
		Caliber:       "12",
	})

	pg.Create(&model.Gun{
		Model:         "ТОЗ-34",
		Company:       "Тульский оружейный завод",
		Country:       "Россия",
		Year:          1964,
		Type:          "shotgun",
		LoadType:      "break",
		BreakType:     "horizontal",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 2,
		Length:        1150,
		BarrelLength:  "711",
		Weight:        3.2,
		Caliber:       "12",
	})

	pg.Create(&model.Gun{
		Model:         "Winchester 1895",
		Company:       "Winchester",
		Country:       "USA",
		Year:          1895,
		Type:          "rifle",
		LoadType:      "lever",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        1160,
		BarrelLength:  "710",
		Weight:        4.1,
		Caliber:       ".405 Winchester",
	})

	pg.Create(&model.Gun{
		Model:         "Marlin Model 336",
		Company:       "Marlin",
		Country:       "USA",
		Year:          1948,
		Type:          "rifle",
		LoadType:      "lever",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        1080,
		BarrelLength:  "508/610",
		Weight:        3.18,
		Caliber:       ".30-30 Winchester",
	})

	pg.Create(&model.Gun{
		Model:         "МР-18М-М",
		Company:       "ИЖМЕХ",
		Country:       "Россия",
		Year:          1964,
		Type:          "shotgun",
		LoadType:      "break",
		BreakType:     "other",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        1230,
		BarrelLength:  "720",
		Weight:        2.6,
		Caliber:       "12",
	})

	pg.Create(&model.Gun{
		Model:         "ВПО Горностай-215",
		Company:       "Вепрь-Молот",
		Country:       "Россия",
		Year:          2018,
		Type:          "rifle",
		LoadType:      "bolt-action",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        1020,
		BarrelLength:  "520",
		Weight:        2.7,
		Caliber:       ".366 TKM",
	})

	pg.Create(&model.Gun{
		Model:         "ATA ARMS Turqua",
		Company:       "ATA ARMS",
		Country:       "Турция",
		Year:          2018,
		Type:          "rifle",
		LoadType:      "bolt-action",
		IsBullPup:     false,
		IsSliced:      true,
		BarrelsNumber: 1,
		Length:        1170,
		BarrelLength:  "610",
		Weight:        3.65,
		Caliber:       ".308 Win",
	})

	pg.Create(&model.Gun{
		Model:         "ИЖ-35",
		Company:       "ИЖМЕХ",
		Country:       "Россия",
		Year:          1978,
		Type:          "sport pistol",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        300,
		BarrelLength:  "156",
		Weight:        1.34,
		Caliber:       ".22 LR",
	})

	pg.Create(&model.Gun{
		Model:         "Glock 17 Gen 4",
		Company:       "Glock Gmbh",
		Country:       "Австрия",
		Year:          2009,
		Type:          "sport pistol",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        186,
		BarrelLength:  "114",
		Weight:        0.62,
		Caliber:       "9x19",
	})

	pg.Create(&model.Gun{
		Model:         "Grand Power T12-FM1",
		Company:       "Grand Power",
		Country:       "Словакия",
		Year:          2015,
		Type:          "traumatic",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        187,
		BarrelLength:  "80",
		Weight:        0.77,
		Caliber:       "10x28",
	})

	pg.Create(&model.Gun{
		Model:         "МР80-13Т",
		Company:       "ИЖМЕХ",
		Country:       "Россия",
		Year:          2008,
		Type:          "traumatic",
		LoadType:      "semi-auto",
		IsBullPup:     false,
		IsSliced:      false,
		BarrelsNumber: 1,
		Length:        162,
		BarrelLength:  "93",
		Weight:        0.73,
		Caliber:       ".45 Rubber",
	})
}
