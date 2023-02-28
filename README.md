## GPG-Chat Notes
- The chat app should make use of p2p and DHT to connect to peers.
- Messages sent to other peers should be encrypted with GPG.
- You must have all other users GPG keys to be able to send/receive messages from them.
- Messages you can not decrypt will not display in message window
- Must use TLS for communication between nodes
- Client mode will display TUI
- Server mode will act as relay/bootstrap for p2p service
  - Maybe store and forward
  - football protocol
- Default to starting client

### Libraries 
- https://github.com/rs/zerolog - pretty logs
- https://github.com/libp2p/go-libp2p - p2p library
- https://github.com/charmbracelet/bubbletea - TUI library
- https://github.com/charmbracelet/bubbles - Components for TUI
- https://github.com/elewis787/boa - Cobra + Bubbletea help enhancement
- https://github.com/ProtonMail/gopenpgp - Golang OpenGPG implementation

### Reference
- https://github.com/manishmeganathan/peerchat - p2p chat based on libp2p
- https://github.com/elewis787/rkl - cli uses bubbletea
- https://github.com/pterm/pterm - Another TUI
- https://github.com/manifoldco/promptui - Another TUI
- https://github.com/nvim-telescope/telescope.nvim - Another TUI


### Other / Interesting
-  [How to build an interactive CLI app with Go, Cobra & promptui](https://dev.to/divrhino/building-an-interactive-cli-app-with-go-cobra-promptui-346n)
- [Charming Cobras with Bubbletea - Part 1](https://elewis.dev/charming-cobras-with-bubbletea-part-1)
- https://github.com/charmbracelet/charm - Easy backend for CLI projects
- https://www.youtube.com/@charmcli/videos - Charm Vidoes
- https://github.com/mudler/edgevpn - VPN based on libp2p