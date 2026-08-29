package inc

type ManagerInterface[T any] interface {
	Register(entity T)
	GetByName(name string) (T, bool)
	GetAll() []T
}

var (
	_ ManagerInterface[Background] = BackgroundManager
	_ ManagerInterface[Character]  = CharacterManager
	_ ManagerInterface[Dialog]     = DialogManager
	_ ManagerInterface[Item]       = ItemManager
	_ ManagerInterface[Script]     = ScriptManager
	_ ManagerInterface[Scene]      = SceneManager
	_ ManagerInterface[Theme]      = ThemeManager
)
