package inc

func init() {
	ScreenManager.Register(Screen{
		Name:       ScreenBVKBranch,
		Title:      "BVK — SAN FRANCISCO BRANCH",
		Background: BgBVKBranch,
		OnSelector: true,
	})
}
