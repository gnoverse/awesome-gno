## Frequently Asked Questions

### Why did we choose Go?

Go is a simple, easy-to-learn language with features that are well-suited for
smart contracts:

- All libraries, packages, and app code are designed to be public.
- Importing a package is straightforward due to the public nature of everything.
- Go tooling is top-notch.

Go is also the only language you need to learn to read the code running the
blockchain, the VM, and the contracts themselves.

For more information, see:
- [Gno Philosophy](https://github.com/gnolang/gno/blob/master/PHILOSOPHY.md)

### How will we manage the blockchain size?

This is still under development, but one main idea is to have multiple
blockchains communicating via the IBC protocol. Gno.land aims to be a primary
hub for developing the Gno programming language ecosystem, much like GitHub is
for the open-source world. We anticipate many other chains will rely on the code
developed on Gno.land, so we're not aiming to make Gno.land fit every use-case. 

### How does the company behind Gno.land (AiB) plan to make money?

The company or its members will own 10% of Genesis tokens. This will motivate us
to continue developing and enhancing the ecosystem, increasing its value for
users. We'll also create packages and Realms that will be monetized through
Proof of Contribution.

Furthermore, we plan to invest in the ecosystem through donations, grants, or
VC-style investments. We'll build products on the chain that can have business
plans. We expect gno.land to be a blockchain that people will enjoy securing
through ICS, allowing us to earn rewards from the entire ecosystem. 

For more information, see:
- [Independence Day](https://github.com/gnolang/independence-day/blob/main/mkgenesis/non-airdrop.txt)

### How does Gno differ from other smart contract solutions, like Solidity?

Gno code is interpreted directly from the source code. This is a deliberate
decision to enforce the use of source code directly to execute the code on the
initial Gno.land chain and subsequent chains. This is further enforced by our
code license, the [Gno Network Public
License](https://github.com/gnolang/gno/blob/master/LICENSE.md), a derivative of
the GNU Affero General Public License.

### What is the difference between Gno and Go?

Gno is almost identical to Go, so if you're a Go developer, you can deploy your
first smart contract in no time compared to other languages.

Gno is an interpreted version of Go designed for blockchains. Its main goal is
to write code that looks like Go but is always deterministic, efficient, and
adapted to the blockchain.

For more information, see:
- [From Go to Gno](https://github.com/gnolang/gno/blob/master/docs/concepts/from-go-to-gno.md)
- [Effective Gno](https://github.com/gnolang/gno/blob/master/docs/how-to-guides/effective-gno.md) (coming soon)
