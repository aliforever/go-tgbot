package tgbot

type menu struct {
	name         string
	backMenuName string
	buttons      [][]string
	handler      menuHandler
}

func NewMenu(name string, handler menuHandler) *menu {
	return &menu{
		name:    name,
		handler: handler,
	}
}

func (m *menu) SetBackMenu(backMenu string) *menu {
	m.backMenuName = backMenu
	m.buttons = append(m.buttons, []string{"Back"})
	return m
}

func (m *menu) SetButtons(buttons [][]string) *menu {
	m.buttons = buttons
	return m
}

func (m *menu) AppendButtonsRow(buttons []string) *menu {
	m.buttons = append(m.buttons, buttons)
	return m
}

func (m *menu) PrependButtonsRow(buttons []string) *menu {
	btns := m.buttons
	m.buttons = append(m.buttons, buttons)
	m.buttons = append(m.buttons, btns...)
	return m
}
