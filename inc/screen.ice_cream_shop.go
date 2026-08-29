package inc

// iceCreamShop stands where Paul's old bar used to be, and they insist there
// never was a bar. The BVK-registered terminal in the corner is the one the
// candy-crush guy will not give up (wiki: story, the Fagyizó beat).
//
// The search he eventually allows is what names the club — a lead, not a door:
// the way there is the trinket shop's address, so this screen has no exit of
// its own.
func screenIceCreamShop() screen {
	return screen{
		name:       ScreenIceCreamShop,
		title:      "ICE CREAM SHOP — WHERE THE BAR WAS",
		background: BgIceCreamShop,
		onSelector: true,
	}
}
