<div align="center">
	<img src="./banner.png" />
</div>

# Awesome Gno [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)

> Do you Gno?

A curated list of awesome apps, frameworks, libraries, tools, and resources for [Gno](https://gno.land) — an interpreted, fully deterministic variation of Go for smart contracts.

**What is Gno.land?** Gno.land is an open-source smart contract platform built on Tendermint2 and GnoVM, using the Gno language to create secure, composable, and transparent decentralized applications.

This repository serves as:
- A launchpad for newcomers to explore the gno.land ecosystem
- A resource hub for developers to find tools, tutorials, and inspiration
- A showcase for community projects built on gno.land
- A space for everyone to publicly share their gno.land projects

> _Contributions welcome! Please read the [contribution guidelines](CONTRIBUTING.md) before submitting a PR._

## Contents

- [Official Links](#official-links)
- [Networks](#networks)
- [Apps](#apps)
- [Community Apps](#community-apps)
- [Developer Tools](#developer-tools)
- [Learning Resources](#learning-resources)
- [SDKs & Clients](#sdks--clients)
- [Frameworks](#frameworks)
- [Social](#social)
- [Miscellaneous](#miscellaneous)
- [Tips & Snippets](#tips--snippets)
- [Archive](#archive)
- [Contributing](#contributing)

## Official Links


- [gno.land](https://gno.land/) - The official gno.land homepage.
- [gno monorepo](https://github.com/gnolang/gno) - The official core repository containing libraries, tools, and example Gno code.
- [docs.gno.land](https://docs.gno.land) - Official documentation for gno.land.
- [gno.land events calendar](https://gno.link/calendar) - Official gno.land Google calendar.
- [Plan](https://github.com/gnolang/gno/blob/master/PLAN.md) - The foundational plan for gno.land.
- [Philosophy](https://github.com/gnolang/gno/blob/master/PHILOSOPHY.md) - The fundamental philosophies of gno.land.

## Networks

_Currently active gno.land networks._

- [Portal Loop (Staging)](https://gno.land/) - A rolling testnet serving as the gno.land homepage. [Learn more](https://docs.gno.land/resources/gnoland-networks#staging-environments).
- [test10.gno.land](https://rpc.test10.testnets.gno.land/) - Latest stable testnet environment (released 2025).

## Apps

_Apps developed by the gno.land team._

- [Gno Playground](https://play.gno.land/) - Online Gno editor for testing, deploying, and sharing code (inspired by the Go Playground).
- [Gno Studio Connect](https://gno.studio/connect) - Simplified interaction with Gno applications.
- [GnoChess](https://github.com/gnoverse/gnochess) - Chess server implementation in Gno with frontend, faucet & tutorial.
- [meme.land](https://github.com/gnoverse/memeland) - Image sharing application built with Gno, React, and Vue.
- [Gno Faucet Hub](https://faucet.gno.land) - Central hub for all gno.land faucets.
- [gno.land Status](https://status.gnoteam.com) - Dashboard showing the status of gno.land services & networks.
- [OpenOcean](https://github.com/Molaryy/openocean) - OpenSea-style NFT marketplace clone in Gno.

## Community Apps

_Apps developed by the gno.land community._

- [Gnoswap](https://github.com/gnoswap-labs/gnoswap) - The first DEX built on gno.land (currently in beta).
- [Adena Wallet](https://adena.app/) - User-friendly wallet for tokens, staking, NFT storage, and dApp connections.
- [Gnolove](https://gnolove.world) - Community leaderboard and contributions analytics for builders of the Gnoland ecosystem.
- [Zenao](https://zenao.io) - Event organization platform for building communities and social organizations.


## Developer Tools

_Tools for developing in Gno._

- [gnodev](https://github.com/gnolang/gno/tree/master/contribs/gnodev) - Your Gno development companion for local development.
- [Gno Debugger](https://gno.land/r/gnoland/blog:p/gno-debugger) - Debugger packaged with the GnoVM.
- [GnoScan](http://gnoscan.io/) - Blockchain explorer for gno.land.
- [tx-indexer](https://github.com/gnolang/tx-indexer) - TM2 blockchain indexer with GraphQL support.
- [Supernova](https://github.com/gnolang/supernova) - Stress testing tool for Gno Tendermint2.


### Editor Support

- [Gno Extension for VS Code](https://marketplace.visualstudio.com/items?itemName=Gnoverse.gnolang) - Language support for Gno in VS Code. ([GitHub](https://github.com/gnoverse/vscode-gno))
- [gno.nvim](https://github.com/x1unix/gno.nvim) - Gno language support for NeoVim.
- [Gno-mode for Emacs](https://gist.github.com/gfanton/6e233656dfeabd7a46f21f7507b6b311) - Major mode for editing Gno files (based on go-mode).
- [Gno for Sublime Text](https://github.com/jdkato/gno-sublime-text) - Gno syntax highlighting for Sublime Text.
- [:GnoFileTest for Vim](https://gist.github.com/grepsuzette/66f5cfaccc1a919c67f52bd7b31a3b09) - `:GnoFileTest` snippet for Vim.

## Learning Resources

_Resources to help you learn gno.land and Gno._

### Getting Started

- [Getting Started](https://github.com/gnolang/getting-started) - Quick-start repo to get your first Gno experience in under 5 minutes.
- [Official Documentation](https://docs.gno.land) - Comprehensive documentation with tutorials and examples.
- [From Zero to gno.land Hero](https://github.com/leohhhn/gno-fzgh) - Complete 0-to-1 tutorial on building your first dApp.

### Tutorials & Guides

- [A gentle introduction to gno.land](https://www.youtube.com/watch?v=hTGeG0z09NU&t=135s) - Intro presentation to gno.land (2024).
- [A Beginner's Guide to the gno.land Testnet](https://medium.com/@onbloc/a-beginners-guide-to-the-gnoland-testnet-6fdc693a48f4) - Visual guide to creating a wallet and receiving testnet tokens.
- [Gno 101](https://github.com/onbloc/gnolang-101) - Course for aspiring smart-contract developers.
- [Gno Basics](https://github.com/moul/gno-basics) - Simple and common examples of Gno realms.
- [Gno Smart Contract Demo](https://www.youtube.com/watch?v=-BlnEXCs0eI) - Short video tutorial on writing and deploying Realms and Packages.

### Presentations & Workshops

- ["go → gno" presentation](https://github.com/gnolang/workshops/tree/main/presentations/2023-06-26--go-to-gno--schollz) - "Things I wish I knew when starting Gno from a Go background" by Zack Scholl.
- [Workshops & Talks](https://github.com/gnolang/workshops) - Slides, videos, and materials from gno.land workshops.

### Community Content

- [Failing In Public](https://proggr.hashnode.dev/gnoland-initial-experience-gonzo-take-on-failing-in-public) - A gonzo journalist take on first gno/CosmosSDK experiences.
- [Peer Dev Learning](https://www.youtube.com/@peerdevlearning) - YouTube channel with Gno development tutorials and learning content.

## SDKs & Clients

_Connecting web2 to the gno.land blockchain._

- [gnoclient](https://github.com/gnolang/gno/tree/master/gno.land/pkg/gnoclient) - a Gno-Go package, allowing you to connect to gno.land chains via Go programs at will.
- [gnolang/blog](https://github.com/gnolang/blog/tree/main/cmd/gnoblog-cli) - `r/gnolang/blog` client, or how to create custom `gnokey` clients, using the [gnoclient](https://github.com/gnolang/gno/tree/master/gno.land/pkg/gnoclient) package.
- [tm2-js-client](https://github.com/gnolang/tm2-js-client) - A TM2 JavaScript client library.
- [gno-js-client](https://github.com/gnolang/gno-js-client) - A Gno JavaScript client library, built upon `tm2-js` with additional Gno-specific functionality.

## Frameworks

_Collections of gno packages providing generic functionality which developers can extend with custom code to create applications._

- [daokit](https://github.com/samouraiworld/gnodaokit) - Create DAOs that can mirror real-world organizations.

## Social

_gno.land's official socials._

- [Discord](https://discord.com/invite/YFtMjWwUN7)
- [Telegram](https://t.me/gnoland)
- [X/Twitter](https://x.com/_gnoland)
- [YouTube](https://www.youtube.com/@_gnoland)

## Miscellaneous

_Miscellaneous._

- [Branding & Assets](https://github.com/gnolang/branding) - Official logo and assets.
- [Peace](https://gno.land/r/gnoland/blog:p/peace) - A call for peace.
- [GitPOAP](https://www.gitpoap.io/gh/gnolang) - Contributors (Git) can mint POAPs.
- [Workshops & Talks](https://github.com/gnolang/workshops) - Slides, videos, and materials to gno.land workshops.
- [Bounty Program](https://github.com/gnolang/bounties) - Official gno.land bounty program for contributors.

## Tips & Snippets

_Note: We'd like to try a section with small tips & snippets and less curation. Please open PRs with just a link to a gist, tweet, screenshot, or discussion._

- [Tendermint2 JSON-RPC Postman Collection](https://gist.github.com/zivkovicmilos/d7b98103f0611ac3b26202a29cee02c4)
- [Goland IDE Intellisense Workaround](https://x.com/leohhhn/status/1836740567541367157) - Developing in the Gno monorepo can provide Gno Intellisense using the builtin Go language support.

## Contributing

Your contributions are always welcome! Please take a look at the 
[contribution guidelines](CONTRIBUTING.md) first.

We will keep some pull requests open if we're not sure whether those libraries 
are awesome. You could [vote for them](https://github.com/gnoverse/awesome-gno/pulls) by adding :+1: to them.

## Archive

_Older, outdated, or archived items._ 

- [Keplr Integration](https://github.com/gnolang/gno/pull/154) - WIP Integration with the Keplr browser extension.
- [Gno to Discord](https://github.com/PoCInnovation/PoCLab) - Send notifications on Discord for new content on https://gno.land/r/demo/boards.
- [test9.gno.land](https://rpc.test9.testnets.gno.land/) - Ninth official testnet environment (archive).
- [test8.gno.land](https://rpc.test8.testnets.gno.land/) - Eighth official testnet environment (archive).
- [test7.gno.land](https://rpc.test7.testnets.gno.land/) - Seventh official testnet environment (archive).
- [test6.gno.land](https://rpc.test6.testnets.gno.land/) - Sixth official testnet environment (archive).
- [test5.gno.land](https://test5.gno.land/) - Fifth official testnet environment (archive).
- [test4.gno.land](https://test4.gno.land/) - Fourth official testnet environment (archive).
- [test3.gno.land](https://test3.gno.land/) - Third official testnet environment (archive).
- [test2.gno.land](https://test2.gno.land/) - Second official testnet environment (archive).
- [test1.gno.land](https://test1.gno.land/) - First official testnet environment (archive).
- [Hello Gno!](https://github.com/xplrz/gnoland-workshop) - A step-by-step workshop on Gno and gno.land's main features.
