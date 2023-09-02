# P2PComet

A lightweight CLI tool to download torrent files!

BitTorrent is a peer-to-peer (P2P) file sharing protocol that allows users to distribute and download files over the internet. It's designed to efficiently distribute large files by breaking them into smaller pieces. Instead of relying on a central server to host the file, BitTorrent relies on a decentralized network of peers, where each peer can upload and download pieces of the file to and from other peers.

## Installation

Install

```bash
  go get github.com/Jay179-sudo/P2PComet
```

## Usage and Examples

General Use

```bash
P2PComet.exe <command> <flag(s)>
```

Example

```bash
P2PComet.exe debian-10.2.0-amd64-netinst.iso.torrent debian.iso
```

## Acknowledgements

- [Cobra-CLI](https://github.com/spf13/cobra)
- [BitTorrent Whitepaper](https://www.bittorrent.org/beps/bep_0003.html)
