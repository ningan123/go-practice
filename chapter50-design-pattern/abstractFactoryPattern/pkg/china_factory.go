package pkg

type ChinaFactory struct {
}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var abstractApple AbstractApple

	abstractApple = new(ChinaApple)

	return abstractApple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var abstractBanana AbstractBanana

	abstractBanana = new(ChinaBanana)

	return abstractBanana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var abstractPear AbstractPear

	abstractPear = new(ChinaPear)

	return abstractPear
}
