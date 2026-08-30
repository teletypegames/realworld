package inc

func init() {
	TapeManager.Register(Tape{
		Name:     TapeMystery,
		Item:     ItemMysteryTape,
		Dialogue: DlgMysteryTape,
	})
}
