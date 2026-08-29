package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenIceCreamShop,
		Title:      "ICE CREAM SHOP — WHERE THE BAR WAS",
		Background: BgIceCreamShop,
		OnSelector: true,
	})
}
