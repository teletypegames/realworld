package inc

func registerContent() {
	registerAll(BackgroundManager, World.G.AssetManager.Register)
	registerAll(CharacterManager, World.G.CharacterManager.Register)
	registerAll(ItemManager, World.G.ItemManager.Register)
	registerAll(DialogManager, World.G.DialogueManager.Register)
	registerAll(ScriptManager, World.G.ScriptManager.Register)
	registerScreen()
}
