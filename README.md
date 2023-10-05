# Golang Chat TUI with OpenPGP encryption

This is a simple terminal user interface (TUI) chat application written in 
Golang that uses OpenPGP encryption to secure messages sent and received. The 
application has peer-to-peer (P2P) capabilities, which means that you can 
chat directly with other users without any centralized server.

The application uses the OpenPGP standard to encrypt messages and ensures that 
only the intended recipient(s) can read them. This is done by generating a 
unique public/private key pair for each user, and exchanging public keys with 
other users to establish a trust relationship. Once the trust relationship is 
established, users can exchange messages that are encrypted and signed with 
their private key.

The TUI interface is simple and intuitive, with a list of contacts on the right 
and a chat window on the left. Users can add new contacts by sharing their 
public key, and can start a chat by selecting a contact and typing a message. The 
application also supports basic commands, such as /help to show the list of available 
commands, /quit to exit the application, and /clear to clear the chat history.

This project is open-source and welcomes contributions from anyone interested in 
improving it. Feel free to clone the repository and submit pull requests with bug 
fixes, new features, or improvements to the documentation.

## Getting started

To run the application, you need to have Golang installed on your system. Once you have 
Golang, you can clone the repository and build the application by running the following 
commands:

```bash
git clone https://github.com/nbazzeghin/pgp-chat.git
cd pgp-chat
go build .
./pgp-chat
```

## GPG-Chat Notes
- The chat app should make use of p2p and DHT to connect to peers.
- Messages sent to other peers should be encrypted with GPG.
- You must have all other users GPG keys to be able to send/receive messages from them.
  - Search `keys.openpgp.org` for keys
  - Ability to configure other key servers
  - Maybe impliment WKD searching (https://wiki.gnupg.org/WKD)
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
- https://goreleaser.com/ - release cli for golang!
- https://golangci-lint.run/ - golang linter w/ IDE support

## Contributing
Contributions are always welcome! If you want to contribute to this project, 
please follow these steps:

1. Fork the repository
2. Create a new branch for your changes
3. Make your changes and commit them
4. Push your changes to your fork
5. Submit a pull request to the main repository

Please make sure to follow the code of conduct and the contribution guidelines when submitting pull requests.

## License
This project is licensed under the MIT License - see the [LICENSE]() file for details.
