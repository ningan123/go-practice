package pkg

type AmericaFactory struct {
}

func (cf *AmericaFactory) CreateApple() AbstractApple {
	var abstractApple AbstractApple

	abstractApple = new(AmericaApple)

	return abstractApple
}

func (cf *AmericaFactory) CreateBanana() AbstractBanana {
	var abstractBanana AbstractBanana

	abstractBanana = new(AmericaBanana)

	return abstractBanana
}

func (cf *AmericaFactory) CreatePear() AbstractPear {
	var abstractPear AbstractPear

	abstractPear = new(AmericaPear)

	return abstractPear
}
