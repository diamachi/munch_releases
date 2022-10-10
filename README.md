# munch_releases
You can find builds for Poco F4 (munch) here


### Some notes to build Awaken OS

#### Environment
  - Ubuntu 18.04
  - 8vCPU 
  - 32GB RAM
  - GCP Spot Instance
  - 400GB SSD
  - AkhilNarang Scripts
  
#### Specific to AwakenOS
  - [For Bison/Flex error](./config.go) - change in build soong ui paths
  - ```export USE_GAPPS=true``` for gapps build 
  - [For 5G/NR support](./vendor.prop) - add in sm8250-common