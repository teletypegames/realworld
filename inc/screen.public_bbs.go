package inc

// publicBBS is one of the wiki's 💡 suggestions: the canon has BBS terminals
// standing in cafés, post offices and bus stops — live infrastructure, not
// relics. An alternative way to read the boards, in a place where somebody can
// read over your shoulder (wiki: entities#potenciális-további-helyszínek).
func screenPublicBBS() screen {
	return screen{
		name:       ScreenPublicBBS,
		title:      "PUBLIC BBS TERMINAL",
		background: BgPublicBBS,
		onSelector: true,
	}
}
