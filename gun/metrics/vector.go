package metrics

import (
	"ais.com/m/model"
	"math"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type VectorisedGun struct {
	Model   string `json:"model"`
	Company string
	Country string
	Year    uint

	Type     string
	LoadType string
	// For break-action rifles only
	BreakType     string
	IsBullPup     bool
	IsSliced      bool
	BarrelsNumber uint
	Bullet        string
	Caliber       float32

	Length       uint
	BarrelLength uint
	Weight       float32
}

func modelsToVectorisedGuns(gun1, gun2 model.Gun) (VectorisedGun, VectorisedGun) {
	lengths1 := strings.Split(gun1.BarrelLength, "/")
	lengths2 := strings.Split(gun2.BarrelLength, "/")

	l1, l2 := closestBarrel(lengths1, lengths2)

	g1 := VectorisedGun{
		Model:         gun1.Model,
		Company:       gun1.Company,
		Country:       gun1.Country,
		Year:          gun1.Year,
		Type:          gun1.Type,
		LoadType:      gun1.LoadType,
		BreakType:     gun1.BreakType,
		IsBullPup:     gun1.IsBullPup,
		IsSliced:      gun1.IsSliced,
		BarrelsNumber: gun1.BarrelsNumber,
		Bullet:        gun1.Caliber,
		Caliber:       BulletToCaliber(gun1.Caliber),
		Length:        gun1.Length,
		Weight:        gun1.Weight,
		BarrelLength:  l1,
	}

	g2 := VectorisedGun{
		Model:         gun2.Model,
		Company:       gun2.Company,
		Country:       gun2.Country,
		Year:          gun2.Year,
		Type:          gun2.Type,
		LoadType:      gun2.LoadType,
		BreakType:     gun2.BreakType,
		IsBullPup:     gun2.IsBullPup,
		IsSliced:      gun2.IsSliced,
		BarrelsNumber: gun2.BarrelsNumber,
		Bullet:        gun2.Caliber,
		Caliber:       BulletToCaliber(gun2.Caliber),
		Length:        gun2.Length,
		Weight:        gun2.Weight,
		BarrelLength:  l2,
	}

	return g1, g2
}

func closestBarrel(lengths1, lengths2 []string) (uint, uint) {
	minLen1, _ := strconv.Atoi(lengths1[0])
	minLen2, _ := strconv.Atoi(lengths2[0])

	for i := 0; i < len(lengths1); i += 1 {
		for j := 0; j < len(lengths2); j += 1 {
			l1, _ := strconv.Atoi(lengths1[i])
			l2, _ := strconv.Atoi(lengths2[j])

			if math.Abs(float64(l2- l1)) < math.Abs(float64(minLen2-minLen1)) {
				minLen1 = l1
				minLen2 = l2
			}
		}
	}

	return uint(minLen1), uint(minLen2)
}

func BulletToCaliber(bullet string) float32 {
	if bullet[0] == '.' {
		bullet = bullet[1:]
		for i := 0; i < len(bullet); i += 1 {
			if !unicode.IsDigit(rune(bullet[i]))  {
				bullet = bullet[:i]
				break
			}
		}

		fBullet, _ := strconv.ParseFloat(bullet, 32)
		return float32(fBullet) * .254
	} else {
		for i := 0; i < len(bullet); i += 1 {
			if !unicode.IsDigit(rune(bullet[i])) && bullet[i] != '.' {
				bullet = bullet[:i]
				break
			}
		}

		fBullet, _ := strconv.ParseFloat(bullet, 32)
		return float32(fBullet)
	}
}

func GunsToVectors(gun1 model.Gun, gun2 model.Gun) ([]float32, []float32) {
	g1, g2 := modelsToVectorisedGuns(gun1, gun2)
	gunVal1 := reflect.ValueOf(&g1).Elem()
	gunVal2 := reflect.ValueOf(&g2).Elem()

	vector1 := make([]float32, gunVal1.NumField())
	vector2 := make([]float32, gunVal2.NumField())

	for i := 0; i < gunVal1.NumField(); i += 1 {
		valueField1 := gunVal1.Field(i)
		valueField2 := gunVal2.Field(i)

		switch valueField1.Interface().(type) {
		case string:
			addStringAttribute(&vector1[i], &vector2[i], valueField1.String(), valueField2.String())
		case uint:
			addUintAttribute(&vector1[i], &vector2[i], uint(valueField1.Uint()), uint(valueField2.Uint()))
		case bool:
			addBoolAttribute(&vector1[i], &vector2[i], valueField1.Bool(), valueField2.Bool())
		case float32:
			addFloatAttribute(&vector1[i], &vector2[i], float32(valueField1.Float()), float32(valueField2.Float()))
		}
	}

	return vector1, vector2
}

func addStringAttribute(vector1, vector2 *float32, attribute1, attribute2 string) {
	if attribute1 == attribute2 {
		*vector1 = 1
		*vector2 = 1
	} else {
		*vector1 = 0
		*vector2 = 1
	}
}

func addUintAttribute(vector1, vector2 *float32, attribute1, attribute2 uint) {
	*vector1 = float32(attribute1)
	*vector2 = float32(attribute2)
}

func addFloatAttribute(vector1, vector2 *float32, attribute1, attribute2 float32) {
	*vector1 = attribute1
	*vector2 = attribute2
}

func addBoolAttribute(vector1, vector2 *float32, attribute1, attribute2 bool) {
	if attribute1 {
		*vector1 = 1
	} else {
		*vector1 = 0
	}

	if attribute2 {
		*vector2 = 1
	} else {
		*vector2 = 0
	}
}
