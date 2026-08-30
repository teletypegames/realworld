package inc

func registerContent() {
	BackgroundManager.Each(World.G.AssetManager.Register)
	CharacterManager.Each(World.G.CharacterManager.Register)
	ItemManager.Each(World.G.ItemManager.Register)
	DialogManager.Each(World.G.DialogueManager.Register)
	ScriptManager.Each(World.G.ScriptManager.Register)
	registerScene()
}
