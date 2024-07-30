#!/bin/bash

# <swiftbar.hideAbout>true</swiftbar.hideAbout>
# <swiftbar.hideRunInTerminal>true</swiftbar.hideRunInTerminal>
# <swiftbar.hideLastUpdated>true</swiftbar.hideLastUpdated>
# <swiftbar.hideDisablePlugin>true</swiftbar.hideDisablePlugin>
# <swiftbar.hideSwiftBar>true</swiftbar.hideSwiftBar>

cat $1 | ${HOME}/bin/totp | pbcopy

echo TOTP
echo ---
echo "Okta| bash=$0 param1=${HOME}/.ssh/okta_key.txt terminal=false"
echo "Github| bash=$0 param1=${HOME}/.ssh/github_key.txt terminal=false"

# add more echo lines like that ^^^ - for other keys, as needed
