## Check hash of smart contract [testnet]
1. Create new multisig on https://multisig.money/
```
Instantiate contract
 type
 /terra.wasm.v1beta1.MsgInstantiateContract
 sender
 terra1cqfk9jyec2la2zf78x5lv0a39ue44z8qpd7t6t
 admin
 terra1cqfk9jyec2la2zf78x5lv0a39ue44z8qpd7t6t
 code_id
 15690
 init_msg
 {
   "max_voting_period": {
     "time": 604800
   },
   "required_weight": 1,
   "voters": [
     {
       "addr": "terra1cqfk9jyec2la2zf78x5lv0a39ue44z8qpd7t6t",
       "weight": 1
     },
     {
       "addr": "terra1ftj4ly5ql78wvvdxtkfyjnleqclmrs3skfkh7t",
       "weight": 1
     }
   ]
```
2. Go to https://bombay-fcd.terra.dev/terra/wasm/v1beta1/codes/15690 - get json 
receive -> 
```
{
  "code_info": {
    "code_id": "15690",
    "code_hash": "e5u0h1FbG7HWVuw1+w35sznQGbycP5sDWliwKqhukSc=",
    "creator": "terra1z75w0td9urxkh254jwqd5m8pd6ulqjfdjsxpj5"
  }
}
```

contract - https://finder.terra.money/testnet/address/terra12h5knfe4snzzf430azcurrwuexaa0jfce25tph

transaction - https://finder.terra.money/testnet/tx/e65858f1452995b98a7da7e08c51e934da83e216d3bbb673b891d69f5425f5bd

```
>>> import base64
>>> base64.b64decode("e5u0h1FbG7HWVuw1+w35sznQGbycP5sDWliwKqhukSc=").hex()
'7b9bb487515b1bb1d656ec35fb0df9b339d019bc9c3f9b035a58b02aa86e9127' - hash of real contract in net
```

3. Try to find hash in different versions in repo https://github.com/CosmWasm/cw-plus/ (check hash only `cw3-fixed-multisig`)
````
====================================================
v = v0.13.0
file size =  286914
hash = 23e1cdc8e3389ebf77b101b462a16c8098f73fc33e3bfb449ec04082a86fc6e8
hash intermediate = 0d16a2eca911272f712f9196f4febc1f703bd16bef0fcfede7f614e818a0370a
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.12.1
file size =  287120
hash = 3d64e3013fa7c23275fb0463b4510043d0c75ac6da152a318d7675eae6e77604
hash intermediate = b9f845420ff33a9efba5ed8b24de67d5a1fb4d5c48399f107ce8a172f2c0c7a5
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.12.0-alpha3
file size =  287558
hash = e3b55df3a1b0c4916d548fcc71345c5ffaa38ad510f29c33e8073673956a1e50
hash intermediate = f1b271b4bc55d289c13f7b0fe3a62dc3265f0d2db31f1204123941adf004b91d
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.12.0
file size =  287552
hash = 422bed258388895b3b04313777501fe7a5d5886a66f1543ef680afdad264a908
hash intermediate = 9a8522a1f4d39970f9250cef1bccf1f991cc544bf4c948fb65f9e11a2317eaa1
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.12.0-alpha2
file size =  287558
hash = e3b55df3a1b0c4916d548fcc71345c5ffaa38ad510f29c33e8073673956a1e50
hash intermediate = f1b271b4bc55d289c13f7b0fe3a62dc3265f0d2db31f1204123941adf004b91d
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.12.0-alpha1
file size =  284663
hash = ed050ce84b500ad646af08f3bea68abb47107485db67e599703c3f1331d01220
hash intermediate = 8cd6b48255044245bc7179ee8eb5f670f94c4a6b4bec981ce39e5d1fcb9398db
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.11.1
file size =  284663
hash = ac1f2327f3c80f897110f0fca0369c7022586e109f856016aef91f3cd1f417c1
hash intermediate = 4d6cd5e7b632271c94d5480ba73821da2cfd61a73710e4def7e6df626628b79c
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.11.0-patch1
file size =  284663
hash = c6485a518d624a39b0406f99e99c10433c6f756c2bd210d0b0d625f2660a3df6
hash intermediate = d5af0e91113b5522a6917319497e04ef34227a934402603141dcf374719258ac
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.11.0
file size =  285828
hash = 228c759b56a19a898a7b74ad88b6baa1e484634f965f3c7d2c356966fbd9aa7f
hash intermediate = d5af0e91113b5522a6917319497e04ef34227a934402603141dcf374719258ac
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.10.2
file size =  255643
hash = dba3463e637eac408613e28183e731d87dd11c3260de2caccedfd55d6860719d
hash intermediate = 9b82479e5669d5b7b58fd74dbb082bca2973e200afcef6a29a44942d9aba3b89
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.10.1
file size =  255643
hash = 1b757ea7b5e4fb9bee326e8955e3f000da8c89e1f489d06419fac757bd7f8a75
hash intermediate = 0b85356553ef233dfd5feadaf6ca603ed1b9c58caf3be1c4e662b753ee27c760
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.10.0
file size =  255643
hash = 640d02bb00e7fba1aa703b48a02816e1a0da7dad3b407d50a363f3a38c0af8d9
hash intermediate = ba33fde6530b10f5cf6c3aa296457519511d75a02708799f75e462d1f0342345
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.10.0-soon4
file size =  255655
hash = 4e36f8c26a316196cc0004b173891daf6f24382e489e3c87f8ae5f043221dc0e
hash intermediate = ae400c9b6bd93b39fb1d41a612cb416208e4b8b4e79ceb55badb15a9c6cff241
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.10.0-soon3
file size =  257224
hash = 1efb1090c5e580aaf6ff769625db4153e4ef596a88bc53c42aebb6251b918444
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.10.0-soon2
file size =  257224
hash = 7e1203f80dbb7de9fcd4e7f5e367d71eda6485f2a6b0a445eb42ef3b6a397374
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.9.1
file size =  255218
hash = 31af69d166595f658b115d2c4cd43351e9a146ccdebfabcf615e275ca7453ca6
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.10.0-soon
file size =  257224
hash = 35cdd1f05aad177fe29b512b6de713215dfcfb24ed0ee423e069e7fa29664b39
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.9.0
file size =  255218
hash = 1436552287aed5e7c23e3030a00eb9ba7dd10e045d1a43cb2cc07970e84f45a8
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.8.1
file size =  255241
hash = c692c6be347cef6cbdf6f525f8834571e1f6a7bc05278e1ca36349d974bcad86
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.8.0-rc3
file size =  255241
hash = db6db31f5a55ea172e563ebbed39a97bb5354495ca734971f2cd20d5e39e8775
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.8.0
file size =  255241
hash = e1db7ab1a74d216fc1e8c90987a9638d5856633159078f9a735e21747a9f669e
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.8.0-rc2
file size =  255241
hash = d7e97b1ac89685bbc07a39a54f48a5b2246116b87ecb4a2546b24d9d8f321b5b
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.8.0-rc1
file size =  288088
hash = 203ceb9ffa61fdcb34b859a4d1249bec517112ecbaf9a9092b3787dc8199dde6
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.7.0
file size =  287547
hash = 363040ffa1928b0f22753bdd2e5153372d14326fc4bb0913a6ee25f0977f0993
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.6.2
file size =  285974
hash = 01ba5adf32deedb9d57072a8e571612cef2908e626cb877026dc6eed7ae0fe29
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.6.1
file size =  285974
hash = 9505bbe4500122f9b68d6ebd248bc04ef3845f4ff0779a70fe9c36a2990fdf6e
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.6.0
file size =  285974
hash = 3e3ea92c4ec35d90aa0bef344bdb0a94939120fce57265d5ff81bf8e8caec7e9
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.6.0-beta3
file size =  283308
hash = 8615889900edcd131531ffa85bb6706ece78ebde6b1bbf378ce0a813d49dac5f
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.6.0-beta2
file size =  278487
hash = 701eb64afaf00c5758d4f373c610f820a8cbf856995fad2ec904b6633b35e47a
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.6.0-beta1
file size =  278487
hash = d8ad8afede2012fc55b10c93c77ae54632dbc36a3911dc83f4a28937a56a1158
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.6.0-alpha3
file size =  264135
hash = d5d97faaa283ed6c6268244acefb7c2f2b2142d6136041410a5e66d0fb97f111
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.6.0-alpha2
file size =  264135
hash = 46caede7a7f610aec7e539d924f5b2e177c05764f595495d8094cdad8cf415cc
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.6.0-alpha1
file size =  264135
hash = 704f809dbb04028e7d8c6b3cfff46c4b384e2b8b9822043fdc273ee6db5bc366
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.5.0
file size =  247262
hash = 447335378d46e3e7487c7def68bab4e4ba4b1841a30feb113d5ed6e5d0db684f
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.4.0
file size =  246062
hash = 23ed7b20d31a8c9ae6881b6858992cd0e9468ef85ad032533e59c32c5a9fa537
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.3.2
file size =  221914
hash = 3659b5e6ead89608a00d696dfc503634108b4cc05a00f460a9c21ef9e2ef83f7
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.3.1
file size =  261632
hash = 9c1d757f2052fc12538be3e01b8cae8a0e69b03ac53a8e686a15052557d8bf61
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.3.0
file size =  260762
hash = 8f089f4e639313b943b62ff75a14044a78e14b58db2f50c50eb8cf74b1687d2f
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.2.3
file size =  255150
hash = 5038c6f516d376a694f54cd2d491f0f6c90ab97d467fc6becd0fe67b450039f7
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.2.2
file size =  255166
hash = f9c328a32cf292784a6bd9d9d90b9ac804726c70ce63113acf90fc60056644d1
filename = "cw3_fixed_multisig.wasm"
====================================================
v = v0.2.1
NOT FOUND CONTRACT v0.2.1
====================================================
v = v0.2.0
NOT FOUND CONTRACT v0.2.0
====================================================
v = v0.1.1
NOT FOUND CONTRACT v0.1.1
====================================================
v = v0.1.0
NOT FOUND CONTRACT v0.1.0

````
Result - hash not found
## Question
How can I verify  the contract in the network ?
