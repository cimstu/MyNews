package card

const (
	CARD_STATE_ON_TABLE_UNSHOWN = iota
	CARD_STATE_ON_TABLE_SENT
	CARD_STATE_IN_HAND_EAST
	CARD_STATE_IN_HAND_SORTH
	CARD_STATE_IN_HAND_WEST
	CARD_STATE_IN_HAND_NORTH
)

const (
	CARD_CODE_BING_BEGIN = 0
	CARD_CODE_WAN_BEGIN  = 9
	CARD_CODE_TIAO_BEGIN = 18
	CARD_CODE_FENG_BEGIN = 27
	CARD_CODE_ZI_BEGIN   = 31
)

const (
	CARD_TYPE_NONE = iota
	CARD_TYPE_BING
	CARD_TYPE_WAN
	CARD_TYPE_TIAO
	CARD_TYPE_FENG
	CARD_TYPE_ZI
)

const (
	EGG_TYPE_NONE  = 0
	EGG_TYPE_FENG  = 1
	EGG_TYPE_ZI    = 2
	EGG_TYPE_1     = 4
	EGG_TYPE_9     = 8
	EGG_TYPE_YAOJI = EGG_TYPE_FENG | EGG_TYPE_ZI | EGG_TYPE_1 | EGG_TYPE_9
)


type CardSlice []*Card

//for sort
func (s CardSlice) Len() int {
	return len(s)
}

func (s CardSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s CardSlice) Less(i, j int) bool {
	return s[i].Code < s[j].Code
}
