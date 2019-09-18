package card

import (
	"fmt"
	"log"
)


func GetEggTypeByCode(code int) int {
	res := EGG_TYPE_NONE

	if CardType(code) == CARD_TYPE_FENG {
		return EGG_TYPE_FENG
	} else if CardType(code) == CARD_TYPE_ZI {
		return EGG_TYPE_ZI
	} else if code == CARD_CODE_BING_BEGIN || code == CARD_CODE_WAN_BEGIN {
		return EGG_TYPE_1
	} else if code == CARD_CODE_WAN_BEGIN-1 || code == CARD_CODE_TIAO_BEGIN-1 || code == CARD_CODE_FENG_BEGIN-1 {
		return EGG_TYPE_9
	} else if code == CARD_CODE_TIAO_BEGIN {
		res = EGG_TYPE_YAOJI
	}

	return res
}

func GetEggType(egg CardSlice) int {
	res := EGG_TYPE_NONE

	for _, card := range egg {
		t := GetEggTypeByCode(card.Code)
		if t == EGG_TYPE_NONE {
			continue
		} else if t == EGG_TYPE_YAOJI {
			res = EGG_TYPE_YAOJI
		}

		return t
	}

	return res
}

func CardType(code int) int {
	if code < CARD_CODE_WAN_BEGIN {
		return CARD_TYPE_BING
	} else if code < CARD_CODE_TIAO_BEGIN {
		return CARD_TYPE_WAN
	} else if code < CARD_CODE_FENG_BEGIN {
		return CARD_TYPE_TIAO
	} else if code < CARD_CODE_ZI_BEGIN {
		return CARD_TYPE_FENG
	} else if code >= CARD_CODE_ZI_BEGIN {
		return CARD_TYPE_ZI
	}

	log.Fatal("Card Type Wrong:", code)
	return CARD_TYPE_NONE
}

func CodeToDesp(code int) string {
	if code >= 0 && code <= 8 {
		return fmt.Sprintf("%d饼", code+1)
	} else if code >= 9 && code <= 17 {
		return fmt.Sprintf("%d万", code-9+1)
	} else if code >= 18 && code <= 26 {
		return fmt.Sprintf("%d条", code-18+1)
	} else if code >= 27 && code <= 30 {
		switch code {
		case 27:
			return "东风"
		case 28:
			return "南风"
		case 29:
			return "西风"
		case 30:
			return "北风"
		}
	} else if code >= 31 && code <= 33 {
		switch code {
		case 31:
			return "红中"
		case 32:
			return "发财"
		case 33:
			return "白板"
		}
	}

	log.Fatal("CardCodeToDesp got wrong code:", code)
	return ""
}

func GetEggCards(slice CardSlice) (cards CardSlice) {
	for _, card := range slice {
		if GetEggTypeByCode(card.Code) & EGG_TYPE_YAOJI > 0 {
			cards = append(cards, card)
		}
	}

	return cards
}

func GetCardsByEggType(slice CardSlice, tp int) (cards CardSlice) {
	for _, card := range slice {
		if GetEggTypeByCode(card.Code) == tp {
			cards = append(cards, card)
		}
	}

	return cards
}

func Combine2Cards(slice CardSlice) (res []CardSlice) {
	for i, card1 := range slice {
		for j := i+1; j < slice.Len(); j++ {
			res = append(res, CardSlice{card1, slice[j]})
		}
	}

	return res
}

func Combine3Cards(slice CardSlice) (res []CardSlice) {
	for i, card1 := range slice {
		combine2 := Combine2Cards(slice[i+1:])
		for _, combine2_slice := range combine2 {
			res = append(res, append(combine2_slice, card1))
		}
	}

	return res
}

func IsEgg(slice CardSlice) bool {
	if slice.Len() != 3 {
		return false
	}

	if GetEggTypeByCode(slice[0].Code) & GetEggTypeByCode(slice[1].Code) & GetEggTypeByCode(slice[2].Code) == 0 {
		return false
	}

	isRepeat := func(slice CardSlice, code int) bool {
		for _, c := range slice {
			if c.Code == code {
				return true
			}
		}

		return false
	}

	for i, c := range slice {
		if c.Code == CARD_CODE_TIAO_BEGIN {
			continue
		}

		if isRepeat(slice[:i], c.Code) {
			return false
		}
	}

	return true
}