package pkg

type JapanFactory struct {
}

func (cf *JapanFactory) CreateApple() AbstractApple {
	var abstractApple AbstractApple

	abstractApple = new(JapanApple)

	return abstractApple
}

func (cf *JapanFactory) CreateBanana() AbstractBanana {
	var abstractBanana AbstractBanana

	abstractBanana = new(JapanBanana)

	return abstractBanana
}

func (cf *JapanFactory) CreatePear() AbstractPear {
	var abstractPear AbstractPear

	abstractPear = new(JapanPear)

	return abstractPear
}
