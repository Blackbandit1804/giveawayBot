 <div align="center">
     <h1>~ giveawayBot ~</h1>
     <strong>A little Discord bot focusing on creating giveaways engagable with reactions.</strong><br><br>
     <img src="https://forthebadge.com/images/badges/made-with-go.svg" height="25"/>
     <br>
     <br>
     <a href="https://travis-ci.org/zekroTJA/giveawayBot"><img src="https://travis-ci.org/zekroTJA/giveawayBot.svg?branch=master"/></a>&nbsp;
     <a class="badge-align" href="https://www.codacy.com/app/zekroTJA/giveawayBot?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=zekroTJA/serverManager2&amp;utm_campaign=Badge_Grade"><img src="https://api.codacy.com/project/badge/Grade/ab73367da2154d6db55c0efb4f44907f"/></a>&nbsp;
     <a href="https://github.com/zekroTJA/serverManager2/releases"><img src="https://img.shields.io/github/release/zekroTJA/giveawayBot/all.svg"/></a>
 </div>

---

# Description

This is a little Discord bot to easily create giveaways joinable by a single click on a reaction. This bot is easy to set up and fast to use.

---

# Installation

Simply use the `installer.bash` script to download all required files:
```bash
$ curl https://raw.githubusercontent.com/zekroTJA/giveawayBot/master/installer.bash >> installer.bash && installer.bash 0.4.1
```
> If you want an other version, just enter the version you want instead of `0.4.1` at the end of the line. You can see all versions in [Releases Tab](https://github.com/zekroTJA/giveawayBot/releases).

Or if you want to compile it by yourself:

> Go installation is required for this!
> You can download it [here](https://golang.org/dl/).

```
$ mkdir -p giveawayBot && cd giveawayBot
$ export GOPATH=$PWD
$ go get github.com/zekroTJA/giveawayBot
$ curl https://raw.githubusercontent.com/zekroTJA/giveawayBot/master/config_example.yaml > bin/config.yaml
```

Now, you have the compiled binary in **`bin/giveawayBot`**. Also there is a config preset. Fill in the config and copy **both files together** somewhere else. Then, just start the binary and Go! (*haha because it's written in Go... Okay, I thought this would be funny...*)

---

# Used 3rd-Party-Packages

- [discordgo](https://github.com/bwmarrin/discordgo)
- [yaml](https://github.com/go-yaml/yaml)

---

© 2018 - present Ringo Hoffmann (zekro Development)  
contact[at]zekro.de | https://zekro.de
