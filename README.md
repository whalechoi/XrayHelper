English | [简体中文](README_zh_CN.md)

# XrayHelper
A unified helper for Android to control system proxy, some scripts in [Xray4Magisk](https://github.com/Asterisk4Magisk/Xray4Magisk) rewritten with golang, provide arm64 and amd64 binary.

# Config
XrayHelper use yml format configuration file, default is `/data/adb/xray/xrayhelper.yml`, and you can customize the path with the `-c` option.  
[Example of xrayhelper config](config.yml)

# Commands
## Control Core Service
`xrayhelper service start`, start core service  
`xrayhelper service stop`, stop core service  
`xrayhelper service restart`, restart core service  
`xrayhelper service status`, show core status  

## Control System Proxy
`xrayhelper proxy enable`, enable system proxy  
`xrayhelper proxy disable`, disable system proxy  
`xrayhelper proxy refresh`, refresh system proxy rule  

## Update Components
- update core  
  `xrayhelper update core`, should configure **xrayHelper.coreType** first
- update adghome  
  `xrayhelper update adghome`, update adghome from [AdguardTeam/AdGuardHome](https://github.com/AdguardTeam/AdGuardHome)
- update tun2socks  
  `xrayhelper update tun2socks`, update tun2socks from [heiher/hev-socks5-tunnel](https://github.com/heiher/hev-socks5-tunnel)
- update geodata  
  `xrayhelper update geodata`, update geodata from [Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat)
- update subscribe  
  `xrayhelper update subscribe`, update your subscribe, should configure **xrayHelper.subList** first
- update yacd-meta  
  `xrayhelper update yacd-meta`, update yacd-meta for mihomo, dest path is `${xrayHelper.coreConfig}/Yacd-meta-gh-pages`
- update metacubexd  
  `xrayhelper update metacubexd`, update metacubexd for mihomo, dest path is `${xrayHelper.coreConfig}/Yacd-meta-gh-pages`

## Switch Proxy Node
### xray, sing-box, hysteria2
- switch subscribe nodes  
  `xrayhelper switch`, should configure **xrayHelper.proxyTag** and update subscribe first, **warning: it will replace your outbounds configuration which has the same proxy tag**
- switch custom nodes  
  `xrayhelper switch custom`, put custom nodes share link into `${xrayHelper.dataDir}/custom.txt` file, then you can find them use this command

### mihomo
- switch subscribe config  
  `xrayhelper switch`, should update subscribe first
- switch custom config  
  `xrayhelper switch example.yaml`, use `${xrayHelper.coreConfig}/example.yaml` file as config

**notice: ${xrayHelper.clash.template} will overwrite(or inject) selected config above**

## License
[Mozilla Public License Version 2.0 (MPL)](https://raw.githubusercontent.com/Asterisk4Magisk/XrayHelper/master/LICENSE)

## Credits
- [@Loyalsoldier/v2ray-rules-dat](https://github.com/Loyalsoldier/v2ray-rules-dat)
- [@2dust/v2rayNG](https://github.com/2dust/v2rayNG)
- [@AdguardTeam/AdGuardHome](https://github.com/AdguardTeam/AdGuardHome)
- [@heiher/hev-socks5-tunnel](https://github.com/heiher/hev-socks5-tunnel)
- ~~[@haishanh/yacd](https://github.com/haishanh/yacd)~~
- [@MetaCubeX/Yacd-meta](https://github.com/MetaCubeX/Yacd-meta)
- [@MetaCubeX/metacubexd](https://github.com/MetaCubeX/metacubexd)
