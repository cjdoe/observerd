

## How to launch

`todo: describe database access`

### Lamport One Time Signature back-end
For proper work with signatures, compiled binary of 
[Lamport One Time Signature back-end](https://github.com/vTCP-Foundation/lamport-crypto-back) 
must be placed alongside with `observerd` binary 
(this document assumes, that `observerd` is located in `./bin` of the sources). <br> 
Please, ensure it is called `lamportc` and has corresponding permissions to be executed. 

**Note:** instead of using the copy of the `lamportc`, 
you can just use a symlink to it from the corresponding sources repo.  

Structure of the `./bin` directory is expected to be like this
```
> ls
  
assets  conf.yaml  keys  observerd  lamportc
```