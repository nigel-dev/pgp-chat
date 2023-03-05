package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
)

type PublicKey struct {
	key         string
	userName    string
	keyId       string
	fingerprint string
	ownerTrust  string
	validity    string
	active      bool
}

func (p PublicKey) Title() string {
	if p.active {
		return "✅ " + p.userName
	}
	return p.userName
}
func (p PublicKey) Description() string {
	return p.fingerprint
}
func (p PublicKey) FilterValue() string {
	if p.active {
		return "FFF " + p.userName
	}
	return p.userName
}

type Client struct {
	input      textarea.Model
	username   string
	users      []string
	viewport   viewport.Model
	userList   list.Model
	keyList    list.Model
	publicKeys []PublicKey
}
